package constdefine

const (
	REDIS_STATE_INFO = "StateInfo"
)

// 工作状态枚举
type WorkStatusEnum int

const (
	WORK_STAT_DEFAULT  WorkStatusEnum = iota //默认
	WORK_STAT_IDLE                           // 空闲
	WORK_STAT_INSERT                         //插枪
	WORK_STAT_PAUSED                         //暂停
	WORK_STAT_CHARGING                       //充电中
	WORK_STAT_FULL                           //充满
	WORK_STAT_OFFLINE                        //离线
	WORK_STAT_FAULT                          //故障
)

// 操作类型枚举
type OpTypeEnum int

const (
	OPEN_CHARGE OpTypeEnum = 1
	STOP_CHARGE OpTypeEnum = 2
)
