package chargeservice

import (
	"context"
	"github.com/njylll/thirdparty_auxiliary_tool_go/constdefine"
	"github.com/njylll/thirdparty_auxiliary_tool_go/dto"
	"github.com/njylll/thirdparty_auxiliary_tool_go/middleware/redis"
	"github.com/njylll/thirdparty_auxiliary_tool_go/tools/charge"
	"github.com/njylll/thirdparty_auxiliary_tool_go/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

const START_WAY_REMOTE = 1

const (
	OP_TYPE_START = 1
	OP_TYPE_STOP  = 2
)

type ChargeCtrlService struct {
}

func (c *ChargeCtrlService) CtrlCharge(param *dto.ChargeCtrlParam) bool {
	// ====== 校验设备状态 =====

	//从redis中取出状态信息
	canIndex := utils.GunCode2CanIndex(param.DevType, param.GunCode)

	key := utils.JoinKey(constdefine.REDIS_STATE_INFO, param.CtrlAddr, ":", strconv.Itoa(canIndex))
	str := []interface{}{key}
	str = append(str, "StateType", "StateValue")
	ret, err := redis.DB10.HmGetString(context.Background(), str)
	if err != nil {
		logrus.Errorf("获取[%s]redis状态信息失败, err: %v", key, err)
		return false
	}

	//解析状态信息
	stat, err := utils.PlatStat2DevStat(ret[0], ret[1])
	if err != nil {
		logrus.Errorf("解析平台状态信息失败, err: %v", err)
		return false
	}

	valid := utils.CheckOperationType(stat, constdefine.OpTypeEnum(param.OperationType))
	if !valid {
		logrus.Infof("[%s]操作类型和当前状态不匹配", param.CtrlAddr)
		return false
	}

	// ===== 发送控制指令到mq =====
	if param.OperationType != OP_TYPE_START && param.OperationType != OP_TYPE_STOP {
		logrus.Errorf("无效的操作类型: %d", param.OperationType)
		return false
	}

	// 发送启动指令
	if param.OperationType == OP_TYPE_START {
		err = charge.SendStartCharge(param.CtrlAddr, canIndex, START_WAY_REMOTE)
		if err != nil {
			logrus.Errorf("发送启动指令到mq失败, %v", err)
			return false
		}
		return true
	}

	// 发送结束指令
	err = charge.SendStopCharge(param.CtrlAddr, canIndex, param.BillCode)
	if err != nil {
		logrus.Errorf("发送启动指令到mq失败, %v", err)
		return false
	}
	return true
}
