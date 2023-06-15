/*
 * Created by 丁德鑫  on 2017/10/25 9:09:51
 * ==============================================
 * 包说明：消息类型定义
 * ==============================================
 * 南京德睿能源研究院版权所有
 * ==============================================
 */
package msgtype

//  空余字段{0x14,0x44,0x73,0x7B,0x7D,0x7E,0x7F,0x81,0x83,0x89,0x8D,0x8E,0x8F,
// 0x98,0xA7,0xBE,0xC0,0xC4,0xCE,0xCF,0xEE,0xF1,0xF3,0xF5,0xF8,0xFE}
const (
	//	消息缩写		消息代码	//	消息名称	发送方
	ErrorAns = 0xFF //	错误.应答	双方
	//设备管理
	HeartbeatReq = 0x00 //	心跳.请求	集控器
	HeartbeatAns = 0x80 //	心跳.应答	云平台
	LoginReq     = 0x02 //	上线申请.请求	集控器
	LoginAns     = 0x82 //	上线申请.应答	云平台
	OfflineReq   = 0x04 //	离线通知.请求	集控器
	OfflineAns   = 0x84 //	离线通知.应答*	云平台

	ShutDownRtuLink = 0x03 //断开链路
	//充电
	StartChargeCmdReq = 0x11 //	平台侧启动充电指令.下发
	StartChargeCmdAns = 0x91 //	平台侧启动充电指令.响应
	StopChargeCmdReq  = 0x13 //	平台侧停止充电指令.下发
	StopChargeCmdAns  = 0x93 //	平台侧停止充电指令.响应
	StartChargeReq    = 0x10 //	设备侧启动充电申请.请求
	StartChargeAns    = 0x90 //	设备侧启动充电申请.响应
	StopChargeReq     = 0x12 //	设备侧停止充电申请.请求
	StopChargeAns     = 0x92 //	设备侧停止充电申请.响应
	SuspendChargeReq  = 0x15 //	平台侧暂停充电指令.下发
	SuspendChargeAns  = 0x95 //	平台侧暂停充电指令.响应
	ResumeChargeReq   = 0x17 //	平台侧恢复充电指令.下发
	ResumeChargeAns   = 0x97 //	平台侧恢复充电指令.响应

	BillUploadReq            = 0x16 //	订单上传.请求
	BillUploadAns            = 0x96 //	订单上传.响应
	ChargerBrgStateChangeReq = 0x18 //	前置自己使用的充电机状态变化通知.请求 原始的没有处理过的充电机状态，发出来给bridge_monitor处理，能把调度中和已完成显示到能量云的终端状态页面上

	ChargerStateChangeReq = 0x20 //	充电机状态变化通知.请求
	ChargerStateChangeAns = 0xA0 //	充电机状态变化通知.响应
	ChargerStateReq       = 0x21 //	充电机状态召唤.请求
	ChargerStateAns       = 0xA1 //	充电机状态召唤.响应
	AccountDetailReq      = 0x22 //	账户详情获取.请求
	AccountDetailAns      = 0xA2 //	账户详情获取.响应
	AccountSyncReq        = 0x24 //	账户信息同步申请.请求
	AccountSyncAns        = 0xA4 //	账户信息同步申请.响应

	CanPileNumReq = 0x25 //设备侧申请CAN地址和终端编号对应关系.请求
	CanPileNumAns = 0xA5 //设备侧申请CAN地址和终端编号对应关系.响应

	PowerCruveCmdReq  = 0x23 //	充电功率计划曲线指令.下发
	PowerCruveCmdAns  = 0xA3 //	充电功率计划曲线指令.响应
	BmsParamNotifyReq = 0x26 // BMS参数通知.请求 D
	BmsParamNotifyAns = 0xA6 //BMS参数通知.响应 P

	TelemetryReq         = 0x31 //	遥测数据召唤.请求
	TelemetryAns         = 0xB1 //	遥测数据召唤.响应
	StateReq             = 0x33 //	状态数据召唤.请求
	StateAns             = 0xB3 //	状态数据召唤.响应
	BmsReq               = 0x35 //	BMS数据召唤.请求
	BmsAns               = 0xB5 //	BMS数据召唤.响应
	TelemetryNotifyReq   = 0x32 //	遥测数据通知.请求
	TelemetryNotifyAns   = 0xB2 //	遥测数据通知.响应
	StateNotifyReq       = 0x34 //	状态数据通知.请求
	StateNotifyAns       = 0xB4 //	状态数据通知.响应
	BmsNotifyReq         = 0x36 //	BMS数据通知.请求
	BmsNotifyAns         = 0xB6 //	BMS数据通知.响应
	BillsNotConfirmedReq = 0x41 //	未被确认所有订单召唤.请求
	BillsNotConfirmedAns = 0xC1 //	未被确认所有订单召唤.响应
	SpecificBillReq      = 0x43 //	指定编号订单召唤.请求
	SpecificBillAns      = 0xC3 //	指定编号订单召唤.响应

	TurnToEmergencyCmdReq = 0x47 //平台要求集控主动转为本地应急充电指令.下发
	TurnToEmergencyCmdAns = 0xC7 //平台要求集控主动转为本地应急充电指令.响应
	ClearAccountCmdReq    = 0x49 //集控账户信息清空命令.下发
	ClearAccountCmdAns    = 0xC9 //集控账户信息清空命令.响应
	GetCarInfoReq         = 0x3D //集控查询车辆信息.请求
	GetCarInfoAns         = 0xBD //集控查询车辆信息请求.响应
	CreateV2GCurveCmdReq  = 0x76 //V2G计划曲线生成指令.下发
	CreateV2GCurveCmdAns  = 0xF6 //V2G计划曲线生成指令响应.上传
	V2GCurveResultReq     = 0x72 //V2G计划曲线确认结果.下发
	V2GCurveResultAns     = 0xF2 //V2G计划曲线确认结果响应.上传

	CallAllDataReq = 0x75 //平台下发总召.请求

	ChargingExceptionReq  = 0x3E //充电异常信息上报
	ThirdBinaryDevInfoReq = 0x40 //第三方A模式设备信息上报

	ReserveCmdReq = 0x7D //预约状态指令下发（0x7D）
	ReserveCmdAns = 0xFD //预约状态指令响应（0xFD）

	// 虚负荷检定
	VirtualLoadCheckReq = 0x0B // 虚负荷检定.指令下发(0x0B)
	VirtualLoadCheckAns = 0x8B // 虚负荷检定.指令下发确认(0x8B)
	VirtualLoadMeterReq = 0x0D // 虚负荷电量上送(0x0D)
	VirtualLoadMeterAns = 0x8D // 虚负荷电量上送确认(0x8D)

	/////////////////////////智能运维////////////////////
	MonitorTelemetryReq  = 0x51 //	遥测数据召唤.请求
	MonitorTelemetryAns  = 0xD1 //	遥测数据召唤.响应
	MonitorSignalReq     = 0x53 //	遥信数据召唤.请求
	MonitorSignalAns     = 0xD3 //	遥信数据召唤.响应
	MonitorStateReq      = 0x55 //	状态数据召唤.请求
	MonitorStateAns      = 0xD5 //	状态数据召唤.响应
	MonitorDynamicStrReq = 0x59 //动态字符数据召唤.请求

	MonitorDynamicStrNotifyReq = 0xD9 //动态字符数据通知.请求
	MonitorTelemetryNotifyReq  = 0x52 //遥测数据通知.请求
	MonitorTelemetryNotifyAns  = 0xD2 //	遥测数据通知.响应
	MonitorSignalNotifyReq     = 0x54 //	遥信数据通知.请求
	MonitorSignalNotifyAns     = 0xD4 //	遥信数据通知.响应
	MonitorStateNotifyReq      = 0x56 //	状态数据通知.请求
	MonitorStateNotifyAns      = 0xD6 //	状态数据通知.响应

	AlarmReq           = 0x58 //	告警数据通知.请求
	AlarmAns           = 0xD8 //	告警数据通知.响应
	ModelConfCmdReq    = 0x5A //模型配置召唤命令.请求
	ModelConfCmdAns    = 0xDA //召唤模型配置命令.响应
	DevConfInfoReq     = 0x5B //设备模型配置信息召唤.请求
	DevConfInfoAns     = 0xDB //设备模型配置信息召唤.响应
	MeasureConfInfoReq = 0x5C //量测模型配置信息召唤.请求
	MeasureConfInfoAns = 0xDC //量测模型配置信息召唤.响应

	SetGunGroupInfoReq = 0x45 //平台下发多枪分组信息.下发
	SetGunGroupInfoAns = 0xC5 //平台下发多枪分组信息.响应
	GunGroupInfoReq    = 0x42 //集控上传主副枪信息通知.请求
	GunGroupInfoAns    = 0xC2 // 集控上传主副枪信息通知.响应

	SetChargeModReq     = 0x1E //设置充电模式指令.下发	0x1E	P ->D
	SetChargeModAns     = 0x9E //置充电模式指令.响应	0x9E	D->P
	SetGroupStrategyReq = 0x1F //设置群充策略指令.下发	0x1F	P->D
	SetGroupStrategyAns = 0x9F //设置群充策略指令.响应	0x9F	D->P
	CallAlarmCmd        = 0x6E //告警召唤命令.下发（0x6E），P -> D (0x6E)

	GetChargeLimitReq   = 0x37 //集控请求新的余额信息（0x37），D ->P
	GetChargeLimitAns   = 0xB7 //集控请求新的余额信息.响应（0xB7），P ->D
	GetChargeFeeReq     = 0x38 //充电过程中计费数据.召唤（0x38），P -> D
	GetChargeFeeAns     = 0xB8 //充电过程中计费数据.上传（0xB8），D-> P
	GetDeviceFileReq    = 0x39 //设备文件获取指令下发0x39，P ->D
	GetDeviceFileAns    = 0xB9 //设备文件获取响应(0xB9)，D ->P
	SendFileToDeviceReq = 0x3A //下发文件到设备0x3A，P ->D
	SendFileToDeviceAns = 0xBA // 下发文件到设备响应0xBA，D ->P

	RemoteControlReq = 0x2C // 远程控制请求0x2C，P -> D
	RemoteControlAns = 0xAC // 远程控制响应0xAC，D -> P

	SetDeviceConfigInfoReq = 0x3B //设备配置信息下发（0x3B），P ->D
	SetDeviceConfigInfoAns = 0xBB //设备配置信息下发.响应 ，D ->P
	RestartCmdReq          = 0x3C //平台下发重启指令 0x3C，P ->D
	RestartCmdAns          = 0xBC // 平台下发重启指令确认(0xBC)，D ->P
	/////////////////////mgc///////////////////////////
	CreateDockerReq        = 0x5E //创建容器请求
	UpdateDockerReq        = 0x5F //更新容器请求
	DestroyDockerReq       = 0x6A //删除容器请求
	CreateDockerAns        = 0xDE //创建容器响应
	UpdateDockerAns        = 0xDF //更新容器响应
	DestroyDockerAns       = 0xEA //删除容器响应
	CallMgcAllDataReq      = 0x65 //数据总召唤.请求
	CallMgcAllDataAns      = 0xE5 // 数据总召唤.响应
	MGCTelemetryNotifyReq  = 0x66 //遥测数据通知.请求
	MGCTelemetryNotifyAns  = 0xE6 //遥测数据通知.响应
	MGCSignalNotifyReq     = 0x68 //遥信数据通知.请求
	MGCSignalNotifyAns     = 0xE8 // 遥信数据通知.响应
	TeleSetlReq            = 0x69 //遥调指令.请求P ->D
	TeleSetlAns            = 0xE9 //遥调指令.响应D ->P
	MGCSOENotifyReq        = 0x70 //SOE数据通知.请求
	MGCSOENotifyAns        = 0xF0 //SOE数据通知.响应
	TelecontrolPreReq      = 0x57 //遥控预置指令.请求
	TelecontrolPreAns      = 0xD7 //遥控反校结果.响应
	TelecontrolExecReq     = 0x59 //遥控执行/撤销指令.请求
	TelecontrolExecAns     = 0xD9 //遥控执行/撤销指令.响应
	CallAmmeterReq         = 0x67 //电表数据召唤.请求
	CallAmmeterAns         = 0xE7 //电表数据召唤.响应
	ReadFileReq            = 0x61 //读文件激活.请求
	ReadFileAns            = 0xE1 // 读文件激活.确认
	ReadFileDataNotifyReq  = 0x62 //读文件数据传输.请求
	ReadFileDataFinishAns  = 0xE2 //读文件数据传输完成.确认
	WriteFileReq           = 0x63 //写文件激活.请求
	WriteFileAns           = 0xE3 //写文件激活.确认
	WriteFileDataNotifyReq = 0x64 //写文件数据传输.请求
	WriteFileDataFinishAns = 0xE4 // 写文件数据传输完成.确认
	FixedValueSelectReq    = 0x6C //定值修改选择.请求
	FixedValueSelectAns    = 0xEC // 定值修改选择.确认
	FixedValueExecuteReq   = 0x6D //定值修改执行.请求
	FixedValueExecuteAns   = 0xED // 定值修改执行.确认
	FixedValueQuaryReq     = 0x6F // 定值查询.请求
	FixedValueQuaryAns     = 0xEF // 定值查询.确认
	// 断线补招相关
	OffDataRecallStateCallReq  = 0x41 // 断线补招存储状态.召唤
	OffDataRecallStateCallAns  = 0xC1 // 断线补招存储状态.上送
	OffDataRecallTaskSendReq   = 0x42 // 断线补招任务指令.下发
	OffDataRecallTaskSendAns   = 0xC2 // 断线补招任务指令.响应
	OffDataRecallTaskCancelReq = 0x43 // 断线补招任务取消指令.下发
	OffDataRecallTaskCancelAns = 0xC3 // 断线补招任务取消指令.响应
	OffDataRecallTaskDataReq   = 0x44 // 断线补招任务离线数据.上送
	OffDataRecallTaskDataAns   = 0xC4 // 断线补招任务离线数据上送.确认

	ShortAddrMappingPushReq        = 0x46 //短地址映射列表推送.下发
	ShortAddrMappingPushAns        = 0xC6 //短地址映射列表推送.响应
	ShortAddrMappingCheckResultReq = 0xBF //短地址映射列表校验结果.上送
	MgcEventReq                    = 0x3F //微网控制器事件上报

	BatteryStatusReq = 0x0D // 储能电池状态数据.上送(0x0D) D->P
	BatteryConfigReq = 0x0F // 储能电池电池配置数据.上送(0x0F) D->P
	BatterySingleReq = 0x14 // 储能电池单体数据.上送(0x14) D->P
	BatteryInfoReq   = 0x71 // 大数据梯次电池基础信息.上送(0x71) D->P

	EconomicPerformanceCurveReq = 0x5D //经济运行曲线下发指令
	EconomicPerformanceCurveAns = 0xDD //经济运行曲线下发指令.反校

	BillUploadFailReq = 0x48 // 订单上传失败通知.请求
	BillUploadFailAns = 0xC8 // 订单上传失败通知.响应

	LockCmdReq               = 0x74 // 平台下发地锁遥控指令.请求 P-> D
	LockCmdAck               = 0xF4 // 设备回复地锁遥控指令.确认 D-> P
	LockCmdAns               = 0x06 // 设备上报地锁遥控指令执行结果.请求 D-> P
	LockCmdAnsAck            = 0x86 //平台回复地锁遥控指令执行结果.确认 P-> D
	LockStateChangeReq       = 0x08 // 地锁状态变化通知.请求 D-> P
	LockStateChangeAns       = 0x88 // 地锁状态变化通知.响应 P-> D
	LockParaSetReq           = 0x77 // 平台下发地锁参数设置指令.请求		P-->D
	LockParaSetAck           = 0xF7 //设备回复地锁参数设置指令.确认 	D-->P
	LockParaSetAns           = 0x0A // 设备回复地锁参数设置结果.请求	0x0A	D-->P
	LockParaSetAnsAck        = 0x8A //平台回复地锁参数设置结果.确认		P-->D
	OutputRelayYkCmqReq      = 0x79 //平台下发输出继电器控制指令.请求	P-->D
	OutputRelayYkAck         = 0xF9 //设备回复继电器控制指令.确认		D-->P
	OutputRelayYkAns         = 0x0C //设备上报输出继电器控制执行结果.请求 D-->P
	OutputRelayYkAnsAck      = 0x8C //平台回复输出继电器控制执行结果.确认	P-->D
	MeterCallReq             = 0x7A //平台下发召唤电表数据.请求	P-->D
	MeterCallAck             = 0xFA // 设备回复召唤电表数据.确认		D-->P
	MeterNotifyReq           = 0x0E //设备上报电表数据通知.请求    D-> P
	MeterCallAns             = 0xFB // 平台召唤电表数据.响应	D-->P
	MeterParaSetReq          = 0x7C // 平台下发电表地址参数设置指令.请求	P-->D
	MeterParaSetAns          = 0xFC // 平台下发电表地址参数设置指令.响应	D-->P
	EmergencyConfigReq       = 0x1D //应急充电配置.请求	P-->D
	EmergencyConfigAns       = 0x9D //应急充电配置.响应     D-->P
	EmergencyDetailConfigReq = 0x1C //应急充电详细配置.请求	P-->D
	EmergencyDetailConfigAns = 0x9C //应急充电详细配置.响应   D-->P
	GetDevConfigReq          = 0x1B //获取集控配置.请求，P ->D
	GetDevConfigAns          = 0x9B //获取集控配置.响应   D-->P
	RebootDevReq             = 0x19 //重启设备指令.请求（0x19） P ->D
	RebootDevAns             = 0x99 //重启设备指令.响应   D-->P
	GetDevVersionReq         = 0x1A //获取设备版本.请求（0x1A），P ->D
	GetDevVersionAns         = 0x9A //获取设备版本.响应   D-->P
	VinPlateNotiftReq        = 0x28 //设备侧上报VIN/车牌号.通知 D-P
	VinPlateNotiftAns        = 0xA8 // 设备侧上报VIN/车牌号.响应		P->D
	DevUpdateCmdReq          = 0x6B // 设备升级指令.请求 P-->D
	DevUpdateCmdAns          = 0xEB // 设备升级指令.确认 D-->P

	SendShutDownLinkWithReasonReq = 0x07 //平台主动断开通信链路和原因(0x07)，P ->D
	SendShutDownLinkWithReasonAns = 0x87 //平台主动断开通信链路和原因.响应(0x87)，D ->P
	ControllerSwitchReq           = 0x05 //平台要求电站切换引流的消息体(0x05)，P ->D
	ControllerSwitchAns           = 0x85 //平台要求电站切换引流的应答消息体  通信--->后台
	QueryButtonStatusReq          = 0x2A //平台查询检修按钮状态.下发(0x2A)，P ->D
	QueryButtonStatusAns          = 0xAA //设备上送检修按钮状态(0xAA)，D ->P
	SendClosePowerFailureReq      = 0x2B //平台关闭开门断电指令.下发(0x2B)，P ->D
	SendClosePowerFailureAns      = 0xAB //设备响应平台关闭开门断电指令(0xAB)，D ->P
	SwitchAuxPowerReq             = 0x2C //平台下发切换辅源指令(0x2C),P ->D
	SwitchAuxPowerAns             = 0xAC //设备响应切换辅源指令(0xAC) D ->P
	CtrlEmergencyReq              = 0x09 //3.0公交站应急指令  SG--->通信

	ActiveProtecParametersAns = 0x94 // 主动防护参数.下发，P ->D
	CMSDataReq                = 0x27 //主动防护数据上报，D ->P
	ModelChangeUploadReq      = 0x4A // 模型变化通知上送（0x4A）
	ModelChangeUploadAns      = 0xCA //模型变化通知上送（0xCA）

	FileDataTransSchedule = 0x2E //读文件数据传输进度.确认（0x2E），D<->P
	V2GRealDataReq        = 0x50 //V2G实时数据上传(0x50)，D -> P
	V2GRecordReq          = 0x60 //V2G充放电记录上传 （0x60），D ->P
	V2GRecordAns          = 0xE0 //V2G充放电记录上传.响应（0xE0） P ->D

	MeterActiveUploadReq = 0x29 //设备侧主动上报电表数据.请求（0x29），D-> P
	CtrlConfigInfoCall   = 0x2D //集控配置信息召唤 （0x2D），P ->D
	CtrlConfigInfoAns    = 0xAD //集控配置信息召唤.响应（0xAD），D ->P
	DevCodeInfoReq       = 0x2F // 设备终端编号请求（0x2F），D ->P
	DevCodeInfoAns       = 0xAF // 设备终端编号请求.响应（0xAF），P ->D

	SetQRCode          = 0x4E // 下发二维码，P ->D
	SendTask           = 0x4B //平台下发任务（0x4B）
	SendTaskAns        = 0xCB //设备响应远程升级指令（0xCB）
	SendTaskStatus     = 0x4C //设备上送升级结果（0x4C）
	SendTaskStatusAns  = 0xCC //平台确认升级结果（0xCC）
	QueryDevVer        = 0x4D //平台查询设备版本（0x4D）这个A模式不用
	QueryTaskStatus    = 0x4C //平台查询任务状态（0x4C）这个A模式不用
	DevUpdateResultAns = 0xAE //平台确认升级结果
	LoginReqUpdata     = 0x01 //A模式转为升级协议的 设备升级登录申请
	SendDevVer         = 0xCD // 设备上报版本信息（远程升级模型）

	SendCtrlConfigReq = 0x0B //平台下发/清空集控上线配置指令（0x0B）
	SendCtrlConfigAns = 0x8B //设备上报下发/清空集控上线配置结果（0x8B）

	HLHTMessageReq = 0x39 //互联互通信息上传请求（0x39），D-> P

	SetOfflineLoadCtrlStrategyReq = 0x30 // 离网负荷约束下发.请求(0x30) P -> D
	SetOfflineLoadCtrlStrategyAns = 0xB0 // 离网负荷约束下发.响应(0x30) D -> P

	DevVersionAns       = 0xCD // 设备上报版本信息（0xCD）
	UpgradeTaskStateAns = 0xCC // 设备上报任务状态（0xCC） // 与A模式协议号有歧义冲突
	OvpnReq             = 0x78 // 设备提交VPN上线凭证请求（0x78）
	OvpnAns             = 0xF8 // 设备提交VPN上线凭证响应（0xF8）
)

var TypeMeans [256]string

func InitTypeMeans() {
	TypeMeans[CreateDockerReq] = "创建容器.请求"
	TypeMeans[UpdateDockerReq] = "更新容器.请求"
	TypeMeans[DestroyDockerReq] = "删除容器.请求"
	TypeMeans[CreateDockerAns] = "创建容器.响应"
	TypeMeans[UpdateDockerAns] = "更新容器.响应"
	TypeMeans[DestroyDockerAns] = "删除容器.响应"
	TypeMeans[ShutDownRtuLink] = "断开链路"
	TypeMeans[ErrorAns] = "错误.响应"
	TypeMeans[HeartbeatReq] = "心跳.请求"
	TypeMeans[HeartbeatAns] = "心跳.应答"
	TypeMeans[LoginReq] = "上线申请.请求"
	TypeMeans[LoginAns] = "上线申请.响应"
	TypeMeans[OfflineReq] = "离线通知.请求"
	TypeMeans[OfflineAns] = "离线通知.相应"
	TypeMeans[StartChargeCmdReq] = "平台侧启动充电指令.下发"
	TypeMeans[StartChargeCmdAns] = "平台侧启动充电指令.响应"
	TypeMeans[StopChargeCmdReq] = "平台侧停止充电指令.下发"
	TypeMeans[StopChargeCmdAns] = "平台侧停止充电指令.响应"
	TypeMeans[StartChargeReq] = "设备侧启动充电申请.请求"
	TypeMeans[StartChargeAns] = "设备侧启动充电申请.响应"
	TypeMeans[StopChargeReq] = "设备侧停止充电申请.请求"
	TypeMeans[StopChargeAns] = "设备侧停止充电申请.响应"
	TypeMeans[SuspendChargeReq] = "平台侧暂停充电指令.下发"
	TypeMeans[SuspendChargeAns] = "平台侧暂停充电指令.响应"
	TypeMeans[ResumeChargeReq] = "平台侧恢复充电指令.下发"
	TypeMeans[ResumeChargeAns] = "平台侧恢复充电指令.响应"

	TypeMeans[BillUploadReq] = "订单上传.请求"
	TypeMeans[BillUploadAns] = "订单上传.响应"
	TypeMeans[ChargerStateChangeReq] = "充电机状态变化通知.请求"
	TypeMeans[ChargerStateChangeAns] = "充电机状态变化通知.响应"
	TypeMeans[ChargerBrgStateChangeReq] = "前置使用的充电机变化通知.请求"
	TypeMeans[ChargerStateReq] = "充电机状态召唤.请求"
	TypeMeans[ChargerStateAns] = "充电机状态召唤.响应"
	TypeMeans[AccountDetailReq] = "账户详情获取.请求"
	TypeMeans[AccountDetailAns] = "账户详情获取.响应"
	TypeMeans[AccountSyncReq] = "账户信息同步申请.请求"
	TypeMeans[AccountSyncAns] = "账户信息同步申请.响应"
	TypeMeans[CanPileNumReq] = "设备侧申请CAN地址和终端编号对应关系.请求"
	TypeMeans[CanPileNumAns] = "设备侧申请CAN地址和终端编号对应关系.响应"
	TypeMeans[PowerCruveCmdReq] = "充电功率计划曲线指令.下发"
	TypeMeans[PowerCruveCmdAns] = "充电功率计划曲线指令.响应"
	TypeMeans[TelemetryReq] = "遥测数据召唤.请求"
	TypeMeans[TelemetryAns] = "遥测数据召唤.响应"
	TypeMeans[StateNotifyReq] = "状态数据通知.请求"
	TypeMeans[StateNotifyAns] = "状态数据通知.响应"
	TypeMeans[BmsReq] = "BMS数据召唤.请求"
	TypeMeans[BmsAns] = "BMS数据召唤.响应"
	TypeMeans[TelemetryNotifyReq] = "遥测数据通知.请求"
	TypeMeans[TelemetryNotifyAns] = "遥测数据通知.响应"

	TypeMeans[StateReq] = "状态数据召唤.请求"
	TypeMeans[StateAns] = "状态数据召唤.响应"

	TypeMeans[BillsNotConfirmedReq] = "未被确认所有订单召唤.请求"
	TypeMeans[BillsNotConfirmedAns] = "未被确认所有订单召唤.响应"
	TypeMeans[SpecificBillReq] = "指定编号订单召唤.请求"
	TypeMeans[SpecificBillAns] = "指定编号订单召唤.响应"
	TypeMeans[BmsNotifyReq] = "BMS数据通知.请求"
	TypeMeans[BmsNotifyAns] = "BMS数据通知.响应"

	TypeMeans[BmsParamNotifyReq] = "BMS参数通知.请求"
	TypeMeans[BmsParamNotifyAns] = "BMS参数通知.响应"

	TypeMeans[SpecificBillReq] = "指定编号订单召唤.请求"
	TypeMeans[SpecificBillAns] = "指定编号订单召唤.响应"

	TypeMeans[TurnToEmergencyCmdReq] = "平台要求集控主动转为本地应急充电指令.下发"
	TypeMeans[TurnToEmergencyCmdAns] = "平台要求集控主动转为本地应急充电指令.响应"

	TypeMeans[ClearAccountCmdReq] = "集控账户信息清空命令.下发"
	TypeMeans[ClearAccountCmdAns] = "集控账户信息清空命令.响应"

	TypeMeans[GetCarInfoReq] = "集控查询车辆信息.请求"
	TypeMeans[GetCarInfoAns] = "集控查询车辆信息请求.响应"

	TypeMeans[CreateV2GCurveCmdReq] = "V2G计划曲线生成指令.下发"
	TypeMeans[CreateV2GCurveCmdAns] = "V2G计划曲线生成指令响应.上传"
	TypeMeans[V2GCurveResultReq] = "V2G计划曲线确认结果.下发"
	TypeMeans[V2GCurveResultAns] = "V2G计划曲线确认结果响应.上传"

	TypeMeans[ReserveCmdReq] = "预约状态指令.请求"
	TypeMeans[ReserveCmdAns] = "预约状态指令.响应"

	TypeMeans[CallAllDataReq] = "平台下发总召.请求"
	TypeMeans[VinPlateNotiftReq] = "设备侧上报VIN/车牌号.通知"

	TypeMeans[GetChargeLimitReq] = "集控请求新的余额信息"
	TypeMeans[GetChargeLimitAns] = "集控请求新的余额信息.响应"
	TypeMeans[GetChargeFeeReq] = "充电过程中计费数据.召唤"
	TypeMeans[GetChargeFeeAns] = "充电过程中计费数据.上传"

	//////////////////////////智能运维///////////////
	TypeMeans[MonitorTelemetryReq] = "智能运维遥测数据召唤.请求"
	TypeMeans[MonitorTelemetryAns] = "智能运维遥测数据召唤.响应"
	TypeMeans[MonitorSignalReq] = "智能运维遥信数据召唤.请求"
	TypeMeans[MonitorSignalAns] = "智能运维遥信数据召唤.响应"
	TypeMeans[MonitorStateReq] = "智能运维状态数据召唤.请求"
	TypeMeans[MonitorStateAns] = "智能运维状态数据召唤.响应"

	TypeMeans[MonitorTelemetryNotifyReq] = "智能运维遥测数据通知.请求"
	TypeMeans[MonitorTelemetryNotifyAns] = "智能运维遥测数据通知.响应"
	TypeMeans[MonitorSignalNotifyReq] = "智能运维遥信数据通知.请求"
	TypeMeans[MonitorSignalNotifyAns] = "智能运维遥信数据通知.响应"
	TypeMeans[MonitorStateNotifyReq] = "智能运维状态数据通知.请求"
	TypeMeans[MonitorStateNotifyAns] = "智能运维状态数据通知.响应"

	TypeMeans[MonitorDynamicStrReq] = "动态字符数据召唤.请求"
	TypeMeans[MonitorDynamicStrNotifyReq] = "动态字符数据通知.请求"

	TypeMeans[AlarmReq] = "智能运维告警数据通知.请求"
	TypeMeans[AlarmAns] = "智能运维告警数据通知.响应"

	TypeMeans[SetGunGroupInfoReq] = "平台下发多枪分组信息.下发"
	TypeMeans[SetGunGroupInfoAns] = "平台下发多枪分组信息.响应"
	TypeMeans[GunGroupInfoReq] = "集控上传主副枪信息通知.请求"
	TypeMeans[GunGroupInfoAns] = "集控上传主副枪信息通知.响应"

	TypeMeans[ModelConfCmdReq] = "模型配置召唤命令.请求"
	TypeMeans[ModelConfCmdAns] = "召唤模型配置命令.响应"
	TypeMeans[DevConfInfoReq] = "设备模型配置信息召唤.请求"
	TypeMeans[DevConfInfoAns] = "设备模型配置信息召唤.响应"
	TypeMeans[MeasureConfInfoReq] = "量测模型配置信息召唤.请求"
	TypeMeans[MeasureConfInfoAns] = "量测模型配置信息召唤.响应"

	TypeMeans[HLHTMessageReq] = "互联互通信息上传请求"

	TypeMeans[SetOfflineLoadCtrlStrategyReq] = "离网负荷约束下发.请求"
	TypeMeans[SetOfflineLoadCtrlStrategyAns] = "离网负荷约束下发.响应"

	// 虚负荷检定
	TypeMeans[VirtualLoadCheckReq] = "虚负荷检定.指令下发"
	TypeMeans[VirtualLoadCheckAns] = "虚负荷检定.指令下发确认"
	TypeMeans[VirtualLoadMeterReq] = "虚负荷电量上送"
	TypeMeans[VirtualLoadMeterAns] = "虚负荷电量上送确认"

	///////////////mgc////////////////////
	TypeMeans[CallMgcAllDataReq] = "数据总召唤.请求"
	TypeMeans[CallMgcAllDataAns] = "数据总召唤.响应"
	TypeMeans[MGCTelemetryNotifyReq] = "遥测数据通知.请求"
	TypeMeans[MGCTelemetryNotifyAns] = "遥测数据通知.响应"
	TypeMeans[MGCSignalNotifyReq] = "遥信数据通知.请求"
	TypeMeans[MGCSignalNotifyAns] = "遥信数据通知.响应"
	TypeMeans[TeleSetlReq] = "遥调指令.请求"
	TypeMeans[TeleSetlAns] = "遥调指令.响应"
	TypeMeans[MGCSOENotifyReq] = "SOE数据通知.请求"
	TypeMeans[MGCSOENotifyAns] = "SOE数据通知.响应"
	TypeMeans[TelecontrolPreReq] = "遥控预置指令.请求"
	TypeMeans[TelecontrolPreAns] = "遥控反校结果.响应"
	TypeMeans[TelecontrolExecReq] = "遥控执行/撤销指令.请求"
	TypeMeans[TelecontrolExecAns] = "遥控执行/撤销指令.响应"
	TypeMeans[CallAmmeterReq] = "电表数据召唤.请求"
	TypeMeans[CallAmmeterAns] = "电表数据召唤.响应"
	TypeMeans[ReadFileReq] = "读文件激活.请求"
	TypeMeans[ReadFileAns] = "读文件激活.确认"
	TypeMeans[ReadFileDataNotifyReq] = "读文件数据传输.请求"
	TypeMeans[ReadFileDataFinishAns] = "读文件数据传输完成.确认"
	TypeMeans[WriteFileReq] = "写文件激活.请求"
	TypeMeans[WriteFileAns] = "写文件激活.确认"
	TypeMeans[WriteFileDataNotifyReq] = "写文件数据传输.请求"
	TypeMeans[WriteFileDataFinishAns] = "写文件数据传输完成.确认"

	TypeMeans[BillUploadFailReq] = "订单上传失败通知.请求"
	TypeMeans[BillUploadFailAns] = "订单上传失败通知.响应"

	TypeMeans[LockCmdReq] = "平台下发地锁遥控指令.请求"
	TypeMeans[LockCmdAck] = " 设备回复地锁遥控指令.确认"
	TypeMeans[LockCmdAns] = " 设备上报地锁遥控指令执行结果.请求"
	TypeMeans[LockCmdAnsAck] = " 平台回复地锁遥控指令执行结果.确认"
	TypeMeans[LockStateChangeReq] = " 地锁状态变化通知.请求"
	TypeMeans[LockParaSetReq] = " 平台下发地锁参数设置指令.请求"
	TypeMeans[LockParaSetAck] = " 设备回复地锁参数设置指令.确认"
	TypeMeans[LockParaSetAns] = " 设备回复地锁参数设置结果.请求"
	TypeMeans[LockParaSetAnsAck] = " 平台回复地锁参数设置结果.确认"
	TypeMeans[OutputRelayYkCmqReq] = " 平台下发输出继电器控制指令.请求"
	TypeMeans[OutputRelayYkAck] = " 设备回复继电器控制指令.确认"
	TypeMeans[OutputRelayYkAns] = " 设备上报输出继电器控制执行结果.请求"
	TypeMeans[OutputRelayYkAnsAck] = " 平台回复输出继电器控制执行结果.确认"
	TypeMeans[MeterCallReq] = " 平台下发召唤电表数据.请求"
	TypeMeans[MeterCallAck] = " 设备回复召唤电表数据.确认"
	TypeMeans[MeterNotifyReq] = " 设备上报电表数据通知.请求"
	TypeMeans[MeterCallAns] = " 平台召唤电表数据.响应"
	TypeMeans[MeterParaSetReq] = " 平台下发电表地址参数设置指令.请求"
	TypeMeans[MeterParaSetAns] = "平台下发电表地址参数设置指令.响应"
	TypeMeans[EmergencyConfigReq] = "应急充电配置.请求"
	TypeMeans[EmergencyConfigAns] = "应急充电配置.响应"
	TypeMeans[EmergencyDetailConfigReq] = "应急充电详细配置.请求"
	TypeMeans[EmergencyDetailConfigAns] = "应急充电详细配置.响应"
	TypeMeans[GetDevConfigReq] = "获取集控配置.请求"
	TypeMeans[GetDevConfigAns] = "获取集控配置.响应"
	TypeMeans[RebootDevReq] = "重启设备指令.请求"
	TypeMeans[RebootDevAns] = "重启设备指令.响应"
	TypeMeans[GetDevVersionReq] = "获取设备版本.请求"
	TypeMeans[GetDevVersionAns] = "获取设备版本.响应"
	TypeMeans[SetChargeModReq] = "设置充电模式指令.下发"
	TypeMeans[SetChargeModAns] = "置充电模式指令.响应"
	TypeMeans[SetGroupStrategyReq] = "设置群充策略指令.下发"
	TypeMeans[SetGroupStrategyAns] = "设置群充策略指令.响应"
	TypeMeans[CallAlarmCmd] = "告警召唤命令.下发"
	TypeMeans[SendShutDownLinkWithReasonReq] = "平台主动断开通信链路和原因"
	TypeMeans[ControllerSwitchReq] = "平台要求电站切换引流的消息体"
	TypeMeans[SendShutDownLinkWithReasonAns] = "平台主动断开通信链路和原因.响应"
	TypeMeans[ControllerSwitchAns] = "平台要求电站切换引流的应答消息体"
	TypeMeans[QueryButtonStatusReq] = "平台查询检修按钮状态.下发"
	TypeMeans[QueryButtonStatusAns] = "设备上送检修按钮状态"
	TypeMeans[SendClosePowerFailureReq] = "平台关闭开门断电指令.下发"
	TypeMeans[SendClosePowerFailureAns] = "设备响应平台关闭开门断电指令"
	TypeMeans[SwitchAuxPowerReq] = "平台下发切换辅源指令"
	TypeMeans[SwitchAuxPowerAns] = "设备响应切换辅源指令"
	TypeMeans[DevUpdateCmdReq] = "平台下发设备升级指令"
	TypeMeans[DevUpdateCmdAns] = "平台下发设备升级指令.响应"
	TypeMeans[CtrlEmergencyReq] = "3.0公交站应急指令"

	TypeMeans[EconomicPerformanceCurveReq] = "经济运行曲线.下发"
	TypeMeans[EconomicPerformanceCurveAns] = "经济运行曲线.反校"

	TypeMeans[ActiveProtecParametersAns] = "主动防护参数.下发"
	TypeMeans[CMSDataReq] = "主动防护数据上报"
	TypeMeans[ModelChangeUploadReq] = "模型变化通知上送"
	TypeMeans[ModelChangeUploadAns] = "模型变化通知上送.响应"

	TypeMeans[FileDataTransSchedule] = "读文件数据传输进度.确认"
	TypeMeans[V2GRealDataReq] = "V2G实时数据上传"
	TypeMeans[V2GRecordReq] = "V2G充放电记录上传"
	TypeMeans[V2GRecordAns] = "V2G充放电记录上传.响应"

	TypeMeans[MeterActiveUploadReq] = "设备侧主动上报电表数据.请求"
	TypeMeans[CtrlConfigInfoCall] = "集控配置信息召唤"
	TypeMeans[CtrlConfigInfoAns] = "集控配置信息召唤.响应"

	TypeMeans[DevCodeInfoReq] = "设备终端编号请求"
	TypeMeans[DevCodeInfoAns] = "设备终端编号请求.响应"

	TypeMeans[SetQRCode] = "下发设置二维码"

	TypeMeans[SetDeviceConfigInfoReq] = "设备配置信息下发"
	TypeMeans[SetDeviceConfigInfoAns] = "设备配置信息下发.响应"
	TypeMeans[RestartCmdReq] = "平台下发重启指令"
	TypeMeans[RestartCmdAns] = "平台下发重启指令.响应"
	TypeMeans[GetDeviceFileReq] = "设备文件获取"
	TypeMeans[GetDeviceFileAns] = "设备文件获取.响应"
	TypeMeans[SendFileToDeviceReq] = "下发文件到设备"
	TypeMeans[SendFileToDeviceAns] = "下发文件到设备.响应"

	TypeMeans[SendTask] = "平台下发任务（0x4B）"
	TypeMeans[SendTaskStatus] = "设备上送升级结果（0x4C）"
	TypeMeans[SendTaskAns] = "设备响应远程升级指令(0xCB)"
	TypeMeans[SendTaskStatusAns] = "平台确认升级结果（0xCC）"
	TypeMeans[DevUpdateResultAns] = "平台确认升级结果 0xAE"
	TypeMeans[LoginReqUpdata] = "设备升级登录申请"
}
