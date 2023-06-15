package utils

import "github.com/njylll/thirdparty_auxiliary_tool_go/constdefine"

// 校验操作类型是否有效
func CheckOperationType(workStat constdefine.WorkStatusEnum, needOperation constdefine.OpTypeEnum) bool {
	if needOperation == constdefine.OPEN_CHARGE {
		return workStat == constdefine.WORK_STAT_INSERT
	}

	if needOperation == constdefine.STOP_CHARGE {
		return workStat == constdefine.WORK_STAT_CHARGING
	}
	return false
}
