package charge

import (
	"fmt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/charge_protobuf"
	. "github.com/njylll/thirdparty_auxiliary_tool_go/pkg/constdefine"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/feslib"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/protoparse"
	"math/rand"
	"time"
)

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
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 12; i++ {
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
