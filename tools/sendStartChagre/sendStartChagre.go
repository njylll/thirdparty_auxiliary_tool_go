package sendStartChagre

import (
	"fmt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/charge_protobuf"
	. "github.com/njylll/thirdparty_auxiliary_tool_go/pkg/constdefine"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/feslib"
	. "github.com/njylll/thirdparty_auxiliary_tool_go/pkg/msgtype"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/protoparse"
	"github.com/njylll/thirdparty_auxiliary_tool_go/setting"
	"log"
	"math/rand"

	"time"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func Run(ctrlAddrStr string, canIndex, chargeType int) {
	//连接mq
	conn, err := amqp.Dial("amqp://" + setting.SendCmdSetting.ConnStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		setting.SendCmdSetting.Topic, // name
		"topic",                      // type
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	failOnError(err, "Failed to declare an exchange")
	sendStartCharge(ch, setting.SendCmdSetting.Topic, ctrlAddrStr, canIndex, chargeType)

}

// 平台侧启动充电指令.下发（0x11）
func sendStartCharge(ch *amqp.Channel, topic string, ctrlAddrStr string, canIndex, chargetype int) {
	startChargeCmdReq := &charge_protobuf.StartChargeCmdReq{}
	startChargeCmdReq.CanIndex = int32(canIndex)
	startChargeCmdReq.BillCode = genBillCode()
	startChargeCmdReq.Type = charge_protobuf.StartChargeType(chargetype)
	startChargeCmdReq.StartWay = int32(charge_protobuf.StartChargeWay_ClouldStart)
	startChargeCmdReq.Strategy = createAccountingStrategy()

	data, err := proto.Marshal(startChargeCmdReq)
	if nil != err {
		fmt.Println("Marshaling StartChargeCmdReq error:", err)
		return
	}

	buffer := PacketToRtp(StartChargeCmdReq, []byte(ctrlAddrStr), data)
	sendMsgToRabbitMq(ch, topic, buffer)
	log.Printf("StartChargeCmdReq %v", startChargeCmdReq)
}

func genBillCode() string {
	now := time.Now()
	billCode := now.Format("20060102150405") + "0" + fmt.Sprintf("%03d", now.Nanosecond()/1000000)
	return billCode
}

func createAccountingStrategy() []*charge_protobuf.AccountingStrategy {
	var acc []*charge_protobuf.AccountingStrategy
	beginTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).UnixNano() / 1e6

	add := 2 * 60 * 60 * 1000

	a := []float64{0.56, 0.76, 0.33, 0.23}
	b := []float64{0.62, 0.12, 0.34, 0.22}
	//mark 硬编码添加费率
	for i := 0; i < 12; i++ {
		rand.Seed(time.Now().UnixNano())

		idx := rand.Int() % 4
		temp := &charge_protobuf.AccountingStrategy{
			BeginTime:    &charge_protobuf.DateTime{Time: beginTime},
			EndTime:      &charge_protobuf.DateTime{Time: beginTime + int64(add)},
			KWhPrice:     a[idx],
			ServicePrice: b[idx],
		}
		beginTime = beginTime + int64(add)
		acc = append(acc, temp)
	}
	return acc
}

// 发送给实时数据中心报文
func PacketToRtp(msgType byte, ctrlAddrBytes []byte, message []byte) []byte {
	data := make([]byte, FesToRtpHearderlength)
	data[0] = ConstHeader
	data[1] = FesToRtpHearderlength
	data[2] = ProtocolType
	data[3] = EncodingType

	payloadlen := feslib.IntToBytes(len(message))
	data[4] = payloadlen[0]
	data[5] = payloadlen[1]
	data[6] = payloadlen[2]
	data[7] = payloadlen[3]
	data[MessageTypeIndex] = msgType
	data[9] = 0

	DRDCSendTimeBytes := protoparse.Int64ToBytes(time.Now().UnixNano() / 1e6)
	for i := 0; i < FesSendTimeLen; i++ { //前置发送报文的时间(UTC格式)，精确到毫秒
		data[DRDCSendTimeIndex+i] = DRDCSendTimeBytes[i]
	}

	actualLen := len(ctrlAddrBytes)
	for i := 0; i < actualLen; i++ { //集控器地址
		data[FesWithRtpCtrlAddrIndex+i] = ctrlAddrBytes[i]
	}

	return append(data[:], message...)
}

func sendMsgToRabbitMq(ch *amqp.Channel, topic string, data []byte) {
	err := ch.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
