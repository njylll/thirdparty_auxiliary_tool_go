package xiamenprotocol

const (
	DefaultHeartbeatInterval = 20 //S
	DefaultIdNum             = "999999"
	DefaultBeginMeter        = 1.0 ///默认开始电表读数

	DefaultV = 200
	DefaultI = 50
)

//登录响应编码
type LoginAnsCode byte

const (
	Ban     LoginAnsCode = 0x00 //0x00: 禁止本桩工作； 0x01:允许本桩工作； 0x02: 服务器在线，本桩正常工作
	Allow   LoginAnsCode = 0x01
	Success LoginAnsCode = 0x02
)

//心跳响应编码
type HeartbeatAnsCode byte

const (
	HeartbeatSuccess = 0x01 //0x01: 成功； 0x02: 校验失败，重新发送；
	HeartbeatNo      = 0x02
)

//充电类型
type ChargeType byte

const (
	Full     ChargeType = 0x00 //0:充满为止；1:按时间充（单位：秒）； 2：按百分比冲；3:按金额充(单位：元)；
	ByTime   ChargeType = 0x01
	BySoc    ChargeType = 0x02
	ByAmount ChargeType = 0x03
)

//启动充电响应编码
type StartChargeAnsCode byte

//0x00: 远程启动充电失败，未知原因 0x01: 远程启动充电成功 0x02: 校验失败，重新发送；
//0x03: 远程启动充电失败，枪未连接 0x04: 充电模块没有剩余，等待空闲充电模块 0x05: 远程启动充电失败，枪处于预约状态
const (
	StartFailure StartChargeAnsCode = 0x00
	StartSuccess StartChargeAnsCode = 0x01
	InValid      StartChargeAnsCode = 0x02
	NotConnected StartChargeAnsCode = 0x03
	NoIdleModule StartChargeAnsCode = 0x04
	Reserving    StartChargeAnsCode = 0x05
)

//数据有效性编码
type DataValidCode byte

const (
	InValidData = 0x01 //0x00: 数据无效 0x01: 数据有效
	ValidData   = 0x00
)

//结束充电响应编码
type StopChargeAnsCode byte

//0x00: 远程停止充电失败 0x01: 远程停止充电成功 0x02: 校验失败，重新发送；
const (
	StopFailure = 0x00
	StopSuccess = 0x01
	StopInvalid = 0x02
)

//结束充电响应
type StopBillAnsCode byte

const (
	StopBillSuccess = 0x00 //0x01: 成功； 0x02: 失败，重新发送；
	StopBillFailure = 0x01
)

//充电桩状态编码
type StatusCode byte

const (
	DevIdle      StatusCode = 0x00 //0x00: 桩或枪状态，空闲 0x01: 桩或枪状态，正在使用 0x02: 桩或枪状态，出错 0x03: 桩或枪状态，不可用 0x04: 桩或枪状态，预约状态
	DevUsing     StatusCode = 0x01
	DevError     StatusCode = 0x02
	DevFault     StatusCode = 0x03
	DevReserving StatusCode = 0x04
)

//枪连接编码
type ConnectCode byte

const (
	StandBy ConnectCode = 0x00 //0x00 枪处于拔出状态； 0x01 枪处于连接状态； 0x02 枪处于已归位状态;
	Plug    ConnectCode = 0x01
	Over    ConnectCode = 0x02
)

//错误原因编码
type ErrorCode byte

//0x01: 锁枪错误 0x02: 对地绝缘错误 0x03: 温度过高 0x04: 无错误,所有错误已经排除
//0x05: 电流保护装置故障 0x06: 电能表接触故障 0x07: 功率接触器故障 0x08: 读卡器故障
//0x09: 软件复位错误 0x10: 低压故障 0x11: 无线通信的信号错误 0x12: 过流保护错误 0xFF: 未知错误
const (
	LockGunError          ErrorCode = 0x01
	GroundInsulationError ErrorCode = 0x02
	HighTemp              ErrorCode = 0x03
	NoError               ErrorCode = 0x04
	IError                ErrorCode = 0x05
	MeterError            ErrorCode = 0x06
	KwError               ErrorCode = 0x07
	CardReaderError       ErrorCode = 0x08
	SoftResetError        ErrorCode = 0x09
	LowVError             ErrorCode = 0x10
	WirelessError         ErrorCode = 0x11
	OverIError            ErrorCode = 0x12
	UnknownError          ErrorCode = 0xFF
)

//状态上送响应编码
type StateAnsCode byte

//0x00: 上报失败； 0x01: 成功；
//0x02: 校验失败，重新发送； 0x03: 忙，拒绝；
const (
	StateSuccess StateAnsCode = 0x00
	StateFailure StateAnsCode = 0x01
	StateInvalid StateAnsCode = 0x02
	StateBusy    StateAnsCode = 0x03
)
