package xiamenprotocol

import (
	stob "struct2bytes"
	"thirdprotocol"
)

var _ thirdprotocol.ThirdProtocol = &Codec{}

// Codec 编解码器
type Codec struct{}

// NewCodec 编解码器构造函数
func NewCodec() thirdprotocol.ThirdProtocol {
	return &Codec{}
}

func (k *Codec) Marshal(obj interface{}) ([]byte, error) {
	return stob.Marshal(obj, ByteOrder)
}

func (k *Codec) UnMarshal(src []byte, obj interface{}) error {
	return stob.Unmarshal(obj, src, ByteOrder)
}

//0x11
//桩每次启动后必须首先发送此包给中心进行认证。 桩根据中心返回值决定是否处于正常工作状态。
type LoginReq struct {
	DevAddr     string       `json:"桩的序列号" stob:"len:26"`
	DevType     string       `json:"桩型号" stob:"len:21"`
	VendorID    byte         `json:"厂商标识" `
	SoftVersion string       `json:"软件版本号" stob:"len:51"`
	GunNum      byte         `json:"枪个数"`
	FixKw       stob.Float16 `json:"桩的额定功率"`
	ICCID       string       `json:"内置 SIM 卡 ICCID"  stob:"len:21"`
	IMSI        string       `json:"内置 模块 IMSI"  stob:"len:21"`
}

//0x11
//应答
type LoginAns struct {
	Code              LoginAnsCode `json:"响应编码"`
	HeartbeatInterval int32        `json:"心跳间隔 秒"`
	Source            uint32       `json:"桩 id。"`
}

//0x15
//心跳
type HeartbeatReq struct {
}

//应答
type HeartbeatAns struct {
	Code HeartbeatAnsCode `json:"响应编码"`
	Time int32            `json:"当前时间"`
}

//0x16
//上报枪的状态信息
type StateReq struct {
	GunCode       byte        `json:"枪号 0:上报桩的"`
	Status        StatusCode  `json:"枪状态"`
	ConnectStatus ConnectCode `json:"插枪状态"`
	ErrorCode     ErrorCode   `json:"故障码"`
	Time          int32       `json:"错误发生时间"`
}

//响应
type StateAns struct {
	Code StateAnsCode `json:"响应编码"`
}

//0x02
//启动充电指令
type StartChargeReq struct {
	GunCode     byte       `json:"枪号"`
	AccountID   string     `json:"用户标识" stob:"len:16"`
	ChargeType  ChargeType `json:"充电类型"`
	ChargeParam int32      `json:"充电参数"`
}

//响应
type StartChargeAns struct {
	Code StartChargeAnsCode `json:"响应编码"`
}

//0x13
//开始充电订单
type StartChargeBillReq struct {
	AccountID     string     `json:"用户标识" stob:"len:16"`
	GunCode       byte       `json:"枪号"`
	BeginTime     int32      `json:"开始充电时间"`
	ExpectEndTime int32      `json:"预计结束时间"`
	BillCode      int32      `json:"交易流水号"`
	ChargeType    ChargeType `json:"充电类型"`
	ChargeParam   int32      `json:"充电参数"`
}

type StartChargeBillAns struct {
	BillCode int32 `json:"交易流水号"`
}

//0x0C
type GetChargeInfoReq struct {
	GunCode byte `json:"枪号"`
}

type GetChargeInfoAns struct {
	Code             DataValidCode `json:"响应码"`
	BillCode         int32         `json:"交易流水号"`
	Soc              byte          `json:"soc"`
	ChargedTime      int16         `json:"已充时间"`
	RemainChargeTime int16         `json:"剩余充电时间"`
	T                stob.Float16  `json:"温度" stob:"prec:1"`
	P                stob.Float16  `json:"功率" stob:"prec:2"`
	V                stob.Float16  `json:"当前电压" stob:"prec:1"`
	I                stob.Float16  `json:"当前电流" stob:"prec:1"`
	MaxV             stob.Float16  `json:"最大单体电压" stob:"prec:2"`
	MaxT             stob.Float16  `json:"最大单体温度" stob:"prec:1"`
	ChargedKwh       stob.Float32  `json:"已充电量" stob:"prec:2"`
}

//0x03
//停止充电
type StopChargeReq struct {
	BillCode int32 `json:"交易流水号"`
}

type StopChargeAns struct {
	Code StopChargeAnsCode `json:"响应码"`
}

//0x14
//结束订单上送
type StopChargeBillReq struct {
	BillCode   int32   `json:"交易流水号"`
	AccountID  string  `json:"用户标识" stob:"len:16"`
	Emergency  byte    `json:"停止充电方式"`
	MeterCount int8    `json:"电量时刻数"`
	Meters     []Meter `json:"电量表" stob:"len:MeterCount"`
}

type Meter struct {
	TimeStamp int32        `json:"时刻"`
	Kwh       stob.Float32 `json:"电量" stob:"prec:2"`
}

type StopChargeBillAns struct {
	Code  StopBillAnsCode `json:"响应码"`
	Kwh   stob.Float32    `json:"充电电量" stob:"prec:2"`
	Money stob.Float32    `json:"充电金额" stob:"prec:2"`
}
