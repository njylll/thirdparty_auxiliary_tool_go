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

// SendStopCharge 平台侧停止充电指令.下发（0x13）
func SendStopCharge(ctrlAddrStr string, canIndex int, code string) error {
	stopChargeCmdReq := &charge_protobuf.StopChargeCmdReq{}
	stopChargeCmdReq.CanIndex = int32(canIndex)
	stopChargeCmdReq.BillCode = code
	stopChargeCmdReq.ChargeType = charge_protobuf.ChargeDischarge_StopChargeType

	data, err := proto.Marshal(stopChargeCmdReq)
	if nil != err {
		return fmt.Errorf("marshaling stopChargeCmdReq error", err)
	}

	buffer := PacketToRtp(StopChargeCmdReq, []byte(ctrlAddrStr), data)
	err = mq.RbClient.SendMsgToRabbitMq(setting.SendCmdSetting.Topic, buffer)
	if err != nil {
		return fmt.Errorf("发送消息至mq失败, err: %v", err)
	}
	logrus.Infof("stopChargeCmdReq %v", stopChargeCmdReq)
	return nil
}
