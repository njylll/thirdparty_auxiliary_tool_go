package charge

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/njylll/thirdparty_auxiliary_tool_go/middleware/mq"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/charge_protobuf"
	. "github.com/njylll/thirdparty_auxiliary_tool_go/pkg/msgtype"
	"github.com/njylll/thirdparty_auxiliary_tool_go/setting"
	"github.com/sirupsen/logrus"
)

// SendStartCharge 平台侧启动充电指令.下发（0x11）
func SendStartCharge(ctrlAddrStr string, canIndex, chargetype int) error {
	startChargeCmdReq := &charge_protobuf.StartChargeCmdReq{}
	startChargeCmdReq.CanIndex = int32(canIndex)
	startChargeCmdReq.BillCode = genBillCode()
	startChargeCmdReq.Type = charge_protobuf.StartChargeType(chargetype)
	startChargeCmdReq.StartWay = int32(charge_protobuf.StartChargeWay_ClouldStart)
	startChargeCmdReq.Strategy = createAccountingStrategy()

	data, err := proto.Marshal(startChargeCmdReq)
	if nil != err {
		return fmt.Errorf("marshaling StartChargeCmdReq error: %v", err)
	}

	buffer := PacketToRtp(StartChargeCmdReq, []byte(ctrlAddrStr), data)
	err = mq.RbClient.SendMsgToRabbitMq(setting.SendCmdSetting.Topic, buffer)
	if err != nil {
		return err
	}
	logrus.Infof("下发StartChargeCmdReq[%s]至mq %v", ctrlAddrStr, startChargeCmdReq)
	return nil
}
