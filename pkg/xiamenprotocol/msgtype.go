package xiamenprotocol

// 消息类型定义
const (
	ANS              = 255
	XiaMenMaxMsgType = 500

	LoginReqType         = 0x11
	StateReqType         = 0x16
	HeartbeatReqType     = 0x15
	StartChargeReqType   = 0x02
	StartBillReqType     = 0x13
	ConfigReqType        = 0x07
	GetChargeInfoReqType = 0x0C
	StopChargeReqType    = 0x03
	StopBillReqType      = 0x20
)

var TypeMeans [XiaMenMaxMsgType]string

func init() {
	InitTypeMeans()
}
func InitTypeMeans() {
	TypeMeans[LoginReqType] = "登陆请求"
	TypeMeans[LoginReqType+ANS] = "登陆应答"
	TypeMeans[StateReqType] = "状态报文上送"
	TypeMeans[StateReqType+ANS] = "状态报文上送应答"
	TypeMeans[HeartbeatReqType] = "心跳请求"
	TypeMeans[HeartbeatReqType+ANS] = "心跳应答"
	TypeMeans[StartChargeReqType] = "开始充电请求"
	TypeMeans[StartChargeReqType+ANS] = "开始充电请求应答"
	TypeMeans[StopChargeReqType] = "停止充电请求"
	TypeMeans[StopChargeReqType+ANS] = "停止充电请求应答"
	TypeMeans[ConfigReqType] = "充电设置"
	TypeMeans[ConfigReqType+ANS] = "充电设置应答"
	TypeMeans[StartBillReqType] = "开始订单上送"
	TypeMeans[StartBillReqType+ANS] = "开始订单上送响应"
	TypeMeans[StopBillReqType] = "结束订单上送"
	TypeMeans[StopBillReqType+ANS] = "结束订单上送响应"
	TypeMeans[GetChargeInfoReqType] = "获取充电信息请求"
	TypeMeans[GetChargeInfoReqType+ANS] = "获取充电信息请求响应"
}
