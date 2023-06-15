package utils

import (
	"fmt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/constdefine"
	"github.com/njylll/thirdparty_auxiliary_tool_go/dto"
	"strconv"
	"strings"
)

// guncode和canindex的转换
func GunCode2CanIndex(devType, gunCode int) int {
	canIndex := 0
	switch devType {
	case 1:
		canIndex = gunCode + 180
	case 2:
		canIndex = gunCode
	case 3:
		canIndex = gunCode + 150
	}
	return canIndex
}

// JoinKey 拼接缓存key
func JoinKey[T ~string](a, b T, sep string, ex ...T) string {
	tmp := make([]string, 2, 2+len(ex))
	tmp[0] = string(a)
	tmp[1] = string(b)
	for _, v := range ex {
		tmp = append(tmp, string(v))
	}
	return strings.Join(tmp, sep)
}

// PlatStat2DevStat 平台状态字符串转换为设备状态
func PlatStat2DevStat(statType, statValue string) (constdefine.WorkStatusEnum, error) {
	types := strings.Split(statType, ",")
	values := strings.Split(statValue, ",")

	if len(types) == 0 || len(types) != len(values) {
		return constdefine.WORK_STAT_DEFAULT, fmt.Errorf("redis中的设备状态信息错误")
	}

	//解析状态信息切片
	stateType := new(dto.StateType)
	for i := 0; i < len(types); i++ {
		typeInt, _ := strconv.Atoi(types[i])
		valInt, _ := strconv.Atoi(values[i])
		err := stateType.SetValue(typeInt, valInt)
		if err != nil {
			return constdefine.WORK_STAT_DEFAULT, err
		}
	}

	//连接确认开关 断开状态
	if stateType.Connect == 0 {
		// 充电机实时状态 故障
		if stateType.Charger == 2 {
			return constdefine.WORK_STAT_FAULT, nil
		}
		return constdefine.WORK_STAT_IDLE, nil
	}

	// 输出继电器状态 闭合
	if stateType.Relay == 1 {
		return constdefine.WORK_STAT_CHARGING, nil
	}

	// 充电机实时状态 暂停
	if stateType.Charger == 4 {
		return constdefine.WORK_STAT_PAUSED, nil
	}

	// 输出继电器状态 断开
	if stateType.Relay == 0 {
		return constdefine.WORK_STAT_INSERT, nil
	}

	return constdefine.WORK_STAT_DEFAULT, nil
}
