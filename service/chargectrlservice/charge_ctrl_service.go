package chargectrlservice

import (
	"github.com/njylll/thirdparty_auxiliary_tool_go/constdefine"
	"github.com/njylll/thirdparty_auxiliary_tool_go/dto"
	"github.com/njylll/thirdparty_auxiliary_tool_go/middleware/redis"
	util2 "github.com/njylll/thirdparty_auxiliary_tool_go/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ChargeCtrlService struct {
}

func (c *ChargeCtrlService) CtrlCharge(param dto.ChargeCtrlParam) bool {
	// ====== 校验设备状态 =====

	//从redis中取出状态信息
	canIndex := util2.GunCode2CanIndex(param.DevType, param.GunCode)
	key := util2.JoinKey(constdefine.REDIS_STATE_INFO, param.CtrlAddr, strconv.Itoa(canIndex))
	str := []interface{}{key}
	str = append(str, "StateType", "StateValue")
	ret, err := redis.DB10.HGetSomeString(nil, str)
	if err != nil {
		logrus.Errorf("获取redis信息失败, err: %v", err)
		return false
	}

	//解析状态信息
	stat, err := util2.PlatStat2DevStat(ret[0], ret[1])
	if err != nil {
		logrus.Errorf("解析平台状态信息失败, err: %v", err)
		return false
	}

	valid := util2.CheckOperationType(stat, constdefine.OpTypeEnum(param.OperationType))
	if !valid {
		return false
	}

	// ===== 发送控制指令到mq =====

}
