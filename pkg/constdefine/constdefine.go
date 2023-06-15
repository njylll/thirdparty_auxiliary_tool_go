/*
 * Created by 丁德鑫  on 2017/11/22 9:23:21
 * ==============================================
 * 包说明：常量的定义，集中存放
 * ==============================================
 * 南京德睿能源研究院版权所有
 * ==============================================
 */
package constdefine

const (
	ConstHeader       = 0x63
	HeaderLengthIndex = 1    // 报文头长度下标
	HeaderLength      = 0x20 // 报文头长度
	PayloadDataLength = 4    // payload长度
	PayloadStartIndex = 4    // payload长度的开始下标
	Fixheaderlength   = 32   // 与集控器通信的固定报文头长
	EncodingTypeIndex = 3    // 消息体编码类型下标
	PROTOCOLLEN       = 1    // 版本号长度
	MessageTypeIndex  = 8    // 报文的基本类型的下标
	Reserved1Index    = 10   // 报文结构第一个预留字段下标
	CtrlAddrIndex     = 12   // 集控器地址的下标
	CtrlAddrLen       = 16   // 集控器地址长度

	ChargeBowAddrIndex       = 12 // 充电弓地址下标
	ChargeBowAddrLen         = 23 // 充电弓地址长度
	ChargeBowFixheaderlength = 39 // 与集控器通信的固定报文头长

	// 前置与后台的报文头
	FesToRtpHearderlength   = 64   // 前置与后台的报文头长度
	FesWithRtpCtrlAddrIndex = 16   // 前置与后台的报文头中集控器地址的下标
	FesSendTimeIndex        = 32   // 前置服务器时间下标
	FesSendTimeLen          = 8    // 前置服务器时间长度
	ServerCodeIndex         = 10   // 通信节点编号起始下标 2个字节
	CtrlVendorIndex         = 12   // 充电设备生产厂商标识下标
	DRDCSendTimeIndex       = 40   // 下行消息的时间起始下标
	DRDCSendTimeLen         = 8    // 下行消息的时间长度
	GuidIndex               = 48   // Guid起始下标
	GuidLen                 = 16   // Guid长度
	MaxServerCodeLen        = 2    // 通信节点编号长度
	ProtocolType            = 0x03 // 协议版本号，第一版为0x03
	EncodingType            = 0x01 // 消息体编码类型，Protobuf为0x01
	JsonEncodingType        = 0x02 // 消息体编码类型，Json为0x02
	BinaryEncodingType      = 0x03 // 消息体编码类型，二进制
	CtrlProtoVersion        = "3"

	// 充电弓前置与后台的报文头
	ChargeBowFesToRtpHearderlength   = 64 + 7 // 前置与后台的报文头长度
	ChargeBowFesWithRtpCtrlAddrIndex = 16     // 前置与后台的报文头中集控器地址的下标
	ChargeBowFesSendTimeIndex        = 32 + 7 // 前置服务器时间下标
	ChargeBowFesSendTimeLen          = 8      // 前置服务器时间长度
	ChargeBowServerCodeIndex         = 10     // 通信节点编号起始下标 2个字节
	ChargeBowCtrlVendorIndex         = 12     // 充电设备生产厂商标识下标
	ChargeBowDRDCSendTimeIndex       = 40 + 7 // 下行消息的时间起始下标
	ChargeBowDRDCSendTimeLen         = 8      // 下行消息的时间长度
	ChargeBowGuidIndex               = 48 + 7 // Guid起始下标
	ChargeBowGuidLen                 = 16     // Guid长度
	ChargeBowMaxServerCodeLen        = 2      // 通信节点编号长度
	ChargeBowProtocolType            = 0x03   // 协议版本号，第一版为0x03
	ChargeBowEncodingType            = 0x01   // 消息体编码类型，Protobuf为0x01
	ChargeBowJsonEncodingType        = 0x02   // 消息体编码类型，Json为0x02
	ChargeBowBinaryEncodingType      = 0x03   // 消息体编码类型，二进制
	ChargeBowCtrlProtoVersion        = "3"

	// 升级服务报文头
	UpdateStartSymbol     = 0x64
	ResponseReq           = 0
	UpdateProtocolVersion = 0x01
	UpdateVendorIndex     = 11 // 厂商标识下标
	UpdateCtrlAddrIndex   = 12 // 集控地址下标

	MinFrozenNum = 2 // 最小冻结点个数

	RemoteUpgradeHearderlength = 32 // 远程升级的报文头长度

	// 智能运维报文头
	OperStartSymbol = 0x65
)

const (
	FromDownChannel = 1 // 前置直接下发给集控的消息
	FromCmdChannel  = 2 // 后台下发给集控的消息
)

// DataCenter
const (
	BJC = "BJC"
	BJ  = "BJ"
	SHC = "SHC"
)

// ShutdownLinkReason
const (
	RouteSwitch              = 1 // 路由切换
	DataCenterDifferent      = 2 // 集控在平台里所属数据中心和通信服务所处的数据中心不一致
	KeyDifferent             = 3 // 集控地址和秘钥在平台中不匹配
	RepeatLogin              = 4 // 重复上线，关闭链路
	FlatShutdown             = 5 // 平台主动断
	FlatCtrlProtoVersionNot3 = 6 // 平台内不是3.0协议的集控
)

// 合肥桩企
const (
	TELDCTRL  = 0 // 特来电集控
	TAITAN    = 1 // 泰坦
	KEDA_A    = 2 // 科大
	SHENGHONG = 3 // 盛弘

)

// A+新协议 redis key
const (
	ThirdEmergencyBillCodeMapping = "thirdEmergencyBillCodeMapping" // 第三方应急订单映射HashKey
)

// B模式redis key
const (
	// 第三方应急订单key前缀  type:String expireTime:default(2day) key:SmcBEmergency:$companyCode:$ctrlAddr:$canIndex:$billCode(plat or dev) value:billCode(dev or plat
	ThirdEmergencyBillPrefixKey = "SmcBEmergency"
)

// A+协议常量
const (
	NotAllowDefaultAccount = 0 // 不允许默认用户充电
	AllowDefaultAccount    = 1 // 允许默认用户充电
)

// redis 通用常量集合
const (
	RedisConnDefaultMonitorPeriod = 2 // redis 活性检测默认探测周期 			单位:(s)
	RedisConnDefaultMaxTries      = 3 // redis 初始化连接默认最大的尝试此时	单位:(次)

	RedisNoExpireTime = 0 // 过期时间为0

	RedisDb5  = 5  // redis db5 用来存储静态信息 如集控 canList等
	RedisDb10 = 10 // redis db10 用来存储变化信息
)

// 第三方桩企
const (
	GUOXIN               = 4  // 国信
	SHENRUI_A            = 5  // 长原深瑞
	AONAI                = 6  // 奥耐
	DINGCHONG            = 7  // 鼎充
	ZHENGHUA             = 8  // 杭州正华
	XINGXING             = 9  // 星星充电
	SnailB               = 10 // 百城新能源
	KeHuaB               = 11 // 科华B
	TaiYuanLongTou       = 12 // 太原龙投
	ChengDuWeiYu         = 13 // 成都蔚宇
	ZhuHaiYinLong        = 14 // 珠海银隆
	YunkcB               = 15 // 云快充B
	DISIDAIJIAOLIUZHUANG = 16 // 第四代交流桩
	SinExcelB            = 17 // 盛弘电气B
	LVNENG_B             = 18 // 绿能B
	NEWYUNKCB            = 19 // 新版云快充 B
	FEIHONGB             = 20 // 飞宏B
	ZHONGNENG            = 21 // 中能
	HDZYT                = 22 // 核达中远通
	LYSLX                = 23 // 洛阳硕力信
	QDDQDC               = 24 // 青岛电气部直流 台湾外包
	SZJUDIAN             = 25 // 深圳聚电 B
	SSEB                 = 26 // 追日电气 B
	CHENGDIAN            = 27 // 橙电
	SZYONGLIAN           = 28 // 深圳永联
	HAIHUIDE             = 29 // 海汇德
	SZYIPULESHIB         = 30 // 深圳驿普乐氏 B
	HUANGSHANTIANSHI_A   = 31 // 黄山天时A
	GUOYAO_B             = 32 // 国耀B
	LUOBINSEN            = 33 // 罗宾森
	XUJI                 = 34 // 许继
	GAOSIBAO             = 35 // 高斯宝
	NANRV                = 36 // 南瑞
	CECB                 = 37 // 中电联
	KELU_B               = 38 // 科陆B
	JUNMA_AP             = 39 // 骏马A+
	YL_B                 = 40 // 深圳永联B模式
	AONENG_B             = 41 // 奥能B
	KESHIDA_B            = 42 // 科士达B
	ZHUIRI_AP            = 43 // 追日A+
	YISHITE_AP           = 44 // 易事特A+
	MAOSHUO_AP           = 45 // 江西茂硕A+
	PAINUO_B_V2_4        = 46 // 珠海派诺B v2.4
	WANMA_AP             = 47 // 万马A+
	PAINUO_B_V1_2        = 48 // 珠海派诺B v1.2
	SUNGROW_AP           = 49 // 阳光电源A+
	XJu                  = 50 // 小桔V1.1
	XinYeB               = 51 // 新页
	XingXingB            = 52 // 星星B
	GAOSIBAOLowPower     = 53 // 高斯宝7kw直流
	XINGXINGWAIBAO       = 54 // 外协星星
	Ieslab               = 55 // 积成电子B
	GUANGFA_B            = 56 // 光法B
	ZHONGXING_B          = 57 // 中兴B
	DINGWANG_B           = 58 // 丁旺B
	DONGXU_B             = 59 // 东旭B
	JUNENG_B             = 60 // 聚能B
	HESHENG_B            = 61 // 赫胜B
	CHENGDIAN_B          = 62 // 橙电
	JIASHENG_AP          = 63 // 洛阳嘉盛A+
	KEDA_B_V3_6          = 64 // 科大智能B v3.6
	KEDA_B_V3_8          = 65 // 科大智能B v3.8
	KESHIDA_AP           = 66 // 科士达A+
	SANYOU_AP            = 67 // 三优A+
	XIAMEN_B             = 68 // 厦门
	YINGWT_B             = 69 // 英威腾B
	GUANGTC_B            = 70 // 广天川B
	JINGNENG_B           = 71 // 京能B模式
	ZHIBANG_AP           = 72 // 智邦A+
	YUNSHAN_B            = 73 // 云杉B
	ANHENG_B             = 74 // 安恒B
	YUNKC_B_V1_2         = 75 // 云快充B V1.2
	HeKang_B             = 76 // 合康B
	MINGRUI_B            = 77 // 明瑞B
	DONGYUAN_AP          = 78 // 东源A+
	TAIHONG_B            = 79 // 泰宏B
	TEBD_B               = 80 // 特变电工B
	ZHIDA_AP             = 81 // 挚达A+
	JIEDT_B              = 82 // 捷电通B
	NENGRUI_B            = 83 // 能瑞B
	BOSHIDUN_AP          = 84 // 铂士顿A+
	HEXINRT_B            = 85 // 和信瑞通B
	FZ_B                 = 86 // 方智B
	JIANCHONG_AP         = 87 // 简充A+
	YIDIAN_AP            = 88 // 亦电A+
	ANHUIYICHONG_AP      = 89 // 安徽易充A+
	SINEXCEL_AP          = 90 // 盛弘A+
	SUNGROW_B            = 96 // 阳关电源乐充B
	HEXINRT_NEW_B        = 99 // 新版和信瑞通B
	SINEXCEL_V5_B        = 98 // 盛弘v5交流桩B

	// 有智能运维需求
	TELD_7KW_AC_PILE_AP   = 200 // 特来电7kw交流单桩
	LINGCH_5TH_AC_PILE_AP = 201 // 领充五代交流单桩
)

// 第三方桩企B模式
const (
	SnailCode    = 1000 // 蜗牛之家
	KehuaCode    = 1001 // 科华
	YunkcCode    = 1002 // 云快充
	SinExcelCode = 1003 // 盛弘电气
)

const (
	EM   = "EM"   // 能量管理抄表标记
	PILE = "Pile" // 能量管理终端数据标记
	EQPT = "Eqpt"
)

const (
	TTAN   = "泰坦A" // 泰坦
	SHONG  = "盛弘A"
	DCHONG = "鼎充A" //
	GXIN   = "国信A"
	SRUI   = "长原深瑞A"
	ANAI   = "奥耐A"
	ZHENGH = "杭州正华A"
	XINGX  = "星星充电A"
)

// 不同B模式厂家的区分
const (
	SNAIL               = "snail"
	KEHUA               = "kehua"
	YUNKC               = "yunkc"
	NEWYUNKC            = "newyunkc"
	SINEXCEL            = "sinexcel"
	GAOSIBAO_A          = "gaosibao"
	GAOSIBAO_A_LOWPOWER = "gaosibao_7kw"
	CHENGDIANB          = "chengdian"
	YL                  = "yonglian"
	LVNENG              = "lvneng"
	FEIHONG             = "feihong"
	COSTAR              = "costar"
	SSE                 = "sse"
	ENPLUS              = "enplus"
	GuoYao              = "guoyao"
	KELU                = "kelu"
	AONENG              = "aoneng"
	KESHIDA             = "keshida"
	XINYE               = "xinye"
	PAINUO_V1_2         = "painuoV1.2"
	IESLAB              = "ieslab"
	GUANGFA             = "guangfa"
	ZHONGXING           = "zhongxing"
	DINGWANG            = "dingwang"
	DONGXU              = "dongxu"
	JUNENG              = "juneng"
	HESHENG             = "hesheng"
	KEDA_V3_6           = "keda_v3_6"
	KEDA_V3_8           = "keda_v3_8"
	YINGWT              = "yingwt"
	JINGNENG            = "jingneng"
	TAIHONG             = "taihong"
	YUNSHAN             = "yunshan"
	ANHENG              = "anheng"
	YUNKC_V1_2          = "yunkc_v1_2"
	MINGRUI             = "mingrui"
	TEBD                = "tebd"
	JIEDT               = "jiedt"
	HEXINRT             = "hexinrt"
	NENGRUI             = "nengrui"
	FZ                  = "fangzhi"
	CCTIA               = "cctia" // 充电联盟
	SUNGROW             = "sungrow"
	SHENRUI             = "shenrui" // 长园深瑞
	TAIYONG             = "taiyong"
	HEXINRT_NEW         = "hexinrt_new"
	SINEXCEL_V5         = "sinexcel_v5"
	RUISU               = "ruisu"
	HEXINRT_V3_2        = "hexinrt_v3_2"

	// A+
	TELD = "teld"

	// mqtt
	CEC         = "cec"
	Xju         = "xiaoju"
	PAINUO_V2_4 = "painuo"
	XingXing    = "XingXing"
	GUANGTC     = "guangtc"
	XIAMEN      = "xiamen"
	HeKang      = "hekang"
)

var companyName [2000]string

func InitCompanyName() {
	companyName[0] = "特来电"
	companyName[TAITAN] = TTAN
	companyName[KEDA_A] = "科大智能A"
	companyName[SHENGHONG] = SHONG
	companyName[DINGCHONG] = DCHONG
	companyName[SHENRUI_A] = SRUI
	companyName[GUOXIN] = GXIN
	companyName[AONAI] = ANAI
	companyName[ZHENGHUA] = ZHENGH
	companyName[XINGXING] = XINGX
	companyName[SnailB] = "百城新能源B"
	companyName[KeHuaB] = "科华B"
	companyName[TaiYuanLongTou] = "太原龙投A"
	companyName[ChengDuWeiYu] = "成都蔚宇A"
	companyName[ZhuHaiYinLong] = "珠海银隆A"
	companyName[YunkcB] = "云快充B"
	companyName[SinExcelB] = "盛弘电气B模式"
	companyName[DISIDAIJIAOLIUZHUANG] = "第四代交流桩"
	companyName[LVNENG_B] = "绿能B"
	companyName[NEWYUNKCB] = "新版云快充B"
	companyName[FEIHONGB] = "飞宏B"
	companyName[ZHONGNENG] = "中能A"
	companyName[HDZYT] = "核达中远通A"
	companyName[LYSLX] = "洛阳硕力信A"
	companyName[QDDQDC] = "青岛电气直流"
	companyName[SZJUDIAN] = "深圳聚电B"
	companyName[SSEB] = "追日电气B"
	companyName[CHENGDIAN] = "橙电A"
	companyName[SZYONGLIAN] = "深圳永联A"
	companyName[HAIHUIDE] = "海汇德A"
	companyName[SZYIPULESHIB] = "深圳驿普乐氏B"
	companyName[HUANGSHANTIANSHI_A] = "黄山天时A"
	companyName[GUOYAO_B] = "国耀B"
	companyName[LUOBINSEN] = "罗宾森A"
	companyName[XUJI] = "许继A"
	companyName[GAOSIBAO] = "高斯宝A"
	companyName[GAOSIBAOLowPower] = "高斯宝西安7kw直流"
	companyName[CECB] = "中电联"
	companyName[NANRV] = "南瑞A"
	companyName[KELU_B] = "科陆B"
	companyName[AONENG_B] = "奥能B"
	companyName[JUNMA_AP] = "骏马A+"
	companyName[XJu] = "小桔v1.1"
	companyName[YL_B] = "深圳永联B模式"
	companyName[KESHIDA_B] = "科士达B"
	companyName[ZHUIRI_AP] = "追日A+"
	companyName[YISHITE_AP] = "易事特A+"
	companyName[MAOSHUO_AP] = "茂硕A+"
	companyName[PAINUO_B_V2_4] = "珠海派诺B v2.4"
	companyName[PAINUO_B_V1_2] = "珠海派诺B v1.2"
	companyName[WANMA_AP] = "万马A+"
	companyName[SUNGROW_AP] = "阳光电源A+"
	companyName[XinYeB] = "新页B"
	companyName[XingXingB] = "星星充电B"
	companyName[XINGXINGWAIBAO] = "外包A+"
	companyName[Ieslab] = "积成电子B"
	companyName[GUANGFA_B] = "光法B"
	companyName[ZHONGXING_B] = "中兴B"
	companyName[DINGWANG_B] = "丁旺B"
	companyName[DONGXU_B] = "东旭B"
	companyName[JUNENG_B] = "聚能B"
	companyName[HESHENG_B] = "赫胜B"
	companyName[CHENGDIAN_B] = "橙电B"
	companyName[JIASHENG_AP] = "洛阳嘉盛A+"
	companyName[KEDA_B_V3_6] = "科大智能B v3.6"
	companyName[KEDA_B_V3_8] = "科大智能B v3.8"
	companyName[KESHIDA_AP] = "科士达A+"
	companyName[SANYOU_AP] = "三优A+"
	companyName[ZHIBANG_AP] = "智邦A+"
	companyName[XIAMEN_B] = "厦门市政"
	companyName[YINGWT_B] = "英威腾B"
	companyName[JINGNENG_B] = "京能B"
	companyName[GUANGTC_B] = "广天川B"
	companyName[TAIHONG_B] = "泰宏B"
	companyName[YUNSHAN_B] = "云杉B"
	companyName[ANHENG_B] = "安恒B"
	companyName[YUNKC_B_V1_2] = "云快充BV1.2"
	companyName[HeKang_B] = "合康B"
	companyName[MINGRUI_B] = "明瑞B"
	companyName[DONGYUAN_AP] = "东源A+"
	companyName[TEBD_B] = "特变电工B"
	companyName[FZ_B] = "方智B"
	companyName[ZHIDA_AP] = "挚达A+"
	companyName[JIEDT_B] = "捷电通B"
	companyName[HEXINRT_B] = "和信瑞通B"
	companyName[NENGRUI_B] = "能瑞B"
	companyName[BOSHIDUN_AP] = "铂士顿A+"
	companyName[JIANCHONG_AP] = "简充A+"
	companyName[YIDIAN_AP] = "亦电A+"
	companyName[ANHUIYICHONG_AP] = "安徽易充A+"
	companyName[SUNGROW_B] = "阳光电源B"
	companyName[SINEXCEL_V5_B] = "盛弘V5交流桩"

	// 有智能运维需求
	companyName[TELD_7KW_AC_PILE_AP] = "特来电7kw交流单桩"
	companyName[LINGCH_5TH_AC_PILE_AP] = "领充五代交流单桩"
}

func GetCompanyName(ctrlVendorIndex byte) string {
	if ctrlVendorIndex < 0 || ctrlVendorIndex > 255 {
		return ""
	}
	return companyName[ctrlVendorIndex]
}

// 链路状态
const (
	INIT      = 0
	CONNECTED = 1
	HEARTOK   = 2
	TIMEOUT   = 3
)

// uuid五部分
const (
	FirstIndex  = 4
	SecondIndex = 6
	ThirdIndex  = 8
	FouthIndex  = 10
)
const (
	ONLINE  = 1 // 集控器上线请求
	OFFLINE = 2 // 集控器离线
)

const (
	CTRLSTATTOPIC = "CtrlStat"
	CTRLSTATGROUP = "CtrlStatGroup"

	OFFLINESTATUS = "2" // 集控器离线
	ONLINESTATUS  = "1" // 集控在线
)

// ErrorCode-错误代码 定义
const (
	Default              = 0x00 // 缺省值
	StartSymbolError     = 0x10 // 起始符错误
	HeaderLengthError    = 0x20 // 报文头长度错误
	EncodingTypeError    = 0x30 //	消息体编码类型错误
	ProtocolVersionError = 0x40 //	协议版本号错误
	MessageTypeError     = 0x50 // 消息类型不支持
	CtrlAddrError        = 0x60 // 集控器地址格式错误
	ParseError           = 0x70 // 解析payload出错，或payload解析正确，但有域数据和消息类型不匹配
	ApplicationError     = 0x80 //	业务应用处理消息时出错
	GenericError         = 0x90 //	其它错误

)

// 集控器上线申请判断结果
const (
	PASS                byte = 1  // 鉴权通过且唯一
	INVALID             byte = 2  // key错误
	NOTEXIST            byte = 3  // 集控器地址不存在
	REPEATED            byte = 5  // 重复的上线申请
	ShutDownLink        byte = 7  // 前置主动解除应急后，前置主动断开链路，这样订单可以上传
	REJECT              byte = 8  // 所属数据中心不一致，拒绝
	KeepingBlock        byte = 9  // 前置一直阻断集控连接
	NotProtocolVersion3 byte = 10 // 平台内不是3.0协议的集控
	ONLINEERROR         byte = 11 // 集控上线时读取redis或者前置程序中的某些错误
	VersionObsolete     byte = 12 // 版本废弃
)

const (
	LenVin      = 17 // bms中vin的长度
	LenBillCode = 18 // 平台订单号长度
)
const ReadRedisNilValueErr string = "nil value"
const (
	// 上行TOPIC 定义
	TOPIC_CDRESPONSE                = "CdResponse"          // 下行指令的响应结果。
	TOPIC_CDREQUEST                 = "CdRequest"           // 集控发起的请求
	TOPIC_CDBILL                    = "CdBill"              // 订单上传
	TOPIC_CDREALTIMEDATA            = "CdRealtime"          // 实时数据
	TOPIC_CDSTATE                   = "CdState"             // 充电机状态变化通知
	TOPIC_CDONLINE                  = "CdOnline"            // 集控器上线离线
	TOPIC_MONITOR_REALTIME          = "MRDPRealTime"        // 运行监控智能运维数据通知 主题
	TOPIC_MONITOR_DEVRELATIONMODEL  = "DevRelationModel"    // 量测模型配置信息
	TOPIC_MONITOR_DEVTELEMETRYDATA  = "DevTelemetryData"    // 遥测数据
	TOPIC_MONITOR_DEVSIGNALDATA     = "DevSignalData"       // 遥信数据
	TOPIC_MONITOR_DEVMEASUREMODEL   = "DevMeasureModel"     // 设备模型配置信息
	TOPIC_MONITOR_DEVSTATEDATA      = "DevStateData"        // 状态数据
	TOPIC_MONITOR_DEVALARMDATA      = "DevAlarmData"        // 告警数据
	TOPIC_METER_ACTIVE_DATA         = "MeterActiveData"     // 电表数据
	TOPIC_CONTROLLERSWITCHACK       = "ControllerSwitchAck" // 电站阻断引流前置通信和后台消息回执主题
	TOPIC_METER_MAC_DATA            = "KeyMeterExchange"    // 发送给MAC的Topic
	TOPIC_MONITOR_DEVDYNAMICSTRDATA = "DevDynamicStrData"   // 动态字符数据

	// 平台下行 MQ TOPIC 定义
	TOPIC_CONTROLLERSWITCH = "ControllerSwitch" // 电站阻断引流前置通信和后台消息主题

	//	TOPIC_MONITOR_BRIDGE   = "MRBridge"            //监控桥接器异常，并发送邮件
	TOPIC_MGC_REALDATA         = "MgcRealData"         // mgc实时数据
	TOPIC_MGC_HISDATA          = "mgc_his_data"        // mgc历史数据
	TOPIC_CHARGE_RTULOG        = "ChargeRtuLog"        // 充电集控器日志
	TOPIC_CTRL_STAT            = "CtrlStat"            // 集控统计
	TOPIC_OPER_RTULOG          = "OperRtuLog"          // 智能运维集控器日志
	TOPIC_OPER_STATE           = "OperState"           // 开门断电和辐源切换上行主题
	TOPIC_MGC_RTULOG           = "MgcRtuLog"           // 微网控制器日志
	TOPIC_CDBRGSTATE           = "BrgCdState"          // 前置发出的充电机状态变化通知
	TOPIC_OldDevMgr            = "HaDeviceQueue"       // 2.0的设备管理的主题
	TOPIC_LOCK                 = "HaCdLockQueue"       // 3.0地锁上行数据主题
	TOPIC_MGCFILEDATA          = "MgcFileData"         // 微网文件数据主题
	TOPIC_GETPILECODE          = "CdGetPileCodeByCtrl" // 设备侧申请CAN地址和终端编号对应关系
	TOPIC_MGCFIXEDVALUE        = "MgcRealData"         // 定值修改数据主题
	TOPIC_CMS_DATA             = "CMSProtectData"      // cms主动防护告警数据
	TOPIC_EM_METER_EXCHANGE    = "emMeterExchange"     // 与充电程序交互抄表数据的主题
	TOPIC_EM_PILE_METER_ESDATA = "EMPileMeterESData"   // 与平台插入电表数据的主题
	TOPIC_EM_BOX_METER_ESDATA  = "EMMeterBoxESData"    // 与平台插入电表数据的主题

	// HEARTTIMEOUT           = 35             //集控器20秒心跳发送一次，超时时间设定为35秒（先参照生产系统）
	MODULECALLCYCLE    = 25  // 终端充电时的召唤周期 此处虽然是20秒，但代码中是10秒判断一次，所以是30秒召唤一次
	DCMODULECANINDEX   = 181 // 直流充电的开始can地址
	TPACMODULECANINDEX = 151 // 三相交流充电的开始can地址

	TimeoutTime          int64 = 10 // 超时时间
	SendHeartReqInterval       = 20 // 前置主动下发心跳的周期

	// MQ定义

	MQQUEUE_HadeviceExchangeQueue = "HadeviceExchangeQueue" // dvMge消费子站管理消息的队列名字
	MQQUEUE_CpCommandQueue        = "CpCommandBrgQueue-"    // 充电进程消费CpCommand的队列名字
	// MQQUEUE_HeFeiCpCommandQueue         = "HeFeiCpCommandBrgQueue-"         //合肥充电进程消费CpCommand的队列名字
	MQQUEUE_BridgeMonitorCpCommandQueue    = "BridgeMonitorCpCommandBrgQueue-" // bridge_monitor进程消费CpCommand的队列名字
	MQQUEUE_MgcCommandQueue                = "CpCommandMgcQueue-"              // mgc进程消费CpCommand的队列名字
	MQQUEUE_MgcDataMqQueue                 = "MgcDataMqQueue-"
	MQQUEUE_MgcResponseMqQueue             = "MgcResponseMqQueue-"
	MQQUEUE_MgcCpCmdMqQueue                = "MgcCpCmdMqQueue-"
	MQQUEUE_KafkaSwitchQueue               = "BrgKafkaSwitchQueue-" // kafka切换队列名
	MQQUEUE_DataSyncProducerCpCommandQueue = "DataSyncProducerCpCommandBrgQueue-"

	// 第三方报文kafka主题
	TOPIC_ThirdChargeRtuLog = "ThirdChargeRtuLog" // 第三方设备发送的报文

	// Kafka下行Topic
	MgcResponse = "MgecResp" // Mgc响应topic

	// TopicMgcReqResponse mgc -> mgec   Mgc响应微网能量云
	TopicMgcReqResponse = "MgecResp"
	// TopicMgecCmd mgec -> mgc 云端向微网控制器发起请求
	TopicMgecCmd = "mgec_cmd" //

	// TopicMgecCmdAns mgec -> mgc 微网能量云回复微网控制器信息
	TopicMgecCmdAns = "mgec_cmd_ans" // 控制器回复

	// TopicMgcReq mgc -> mgec 微网控制器向微网能量云发起请求
	TopicMgcReq = "mgc_req"

	//  TopicMgcBatteryTransmit 大数据梯次电池转发
	TopicMgcBatteryTransmit = "mgc_battery_transmit"
)

const (
	SendOperdata           = "1" // 发智能运维数据
	DevModelLastUpdateTime = ""  // 模型更新时间为0
)

const (
	StartBill  = "StartBill"  // 开始订单
	FinishBill = "FinishBill" // 结束订单
)
const (
	RcvRouteSwitchErr = 1 // 路由切换主题
	RabbitMqtopicsErr = 2 // 充电主题
)

const ( // 下行消息类型

	ACCEPTED = 0 // “启动通知”请求消息注册应答结果  成功
	REJECTED = 1 // “启动通知”请求消息注册应答结果 拒绝
	HEART    = 2 // 心跳报文
	CMD      = 3 // 其他下行命令
)

// 发送的离线通知中离线原因重新梳理
const (
	NETERROR          = 1   // TCP连接断开
	HEARTTIMEOUTERROR = 2   // 心跳超时
	ACTIVEDOWN        = 3   // 云端主动断开，（EM页面手动断开链路、其他应用下发断开链路指令）
	DATACENTERSWITCH  = 4   // 数据中心路由切换
	TCPWriteError     = 5   // TCP写失败而断开
	PowerFailure      = 6   // 断电导致的离网
	RemoteRestartCmd  = 7   // 集控监视上远程重启集控而离网
	RemoteUpgrade     = 8   // 升级后也会自动重启而离网
	MAXEXITREASON     = 255 // 其他
)

const (
	ERROR = "[ERROR]"
	WARN  = "[WARN]"
	INFO  = "[INFO]"
	DEBUG = "[DEBUG]"
)

const (
	HEARTINTERVAL = 20 // 心跳时间间隔
	SENTINELMODE  = 1  // redis哨兵模式
	OK            = 1  // 1的宏定义
)

const (
	ProtocolVersion3 = "3" // 平台内是3.0协议的集控
)
const (
	HFC = "HFC" // HFC数据中心
)
const (
	// add way 20180411
	CtrlCanIndex  = 250 // 集控器CanIndex
	CtrlAlarmCode = 500 // 集控器告警码
)

// mgc//
const (
	MAX_FILESEND_ASDU_LEN = 5000 // 发送文件每帧的最大长度
	YK_EXPIRED_TIME       = 45   // 下行命令的实效性时间
	YK_REPLY_EXPIRED_TIME = 30   // 遥控反校的等待超时时间

	MgcV3 = "v3" // 微网3.0
	MgcV2 = "v2" // 微网2.0
)

// 白名单账户同步类型
const (
	ACCOUNT_SYNC_TYPE_CARD = "1" // 1：卡
	ACCOUNT_SYNC_TYPE_CAR  = "2" // 2：车
)

// 白名单是否删除标志
const (
	ACCOUNT_IS_DELETE_TRUE  = "1" // 1：已删除
	ACCOUNT_IS_DELETE_FALSE = "0" // 0：未删除
)

// 当前充电使用的kafka
const (
	MAIN_KAFKA  = 1 // 主kafka
	SPARE_KAFKA = 2 // 备用kafka
)

// 集控器日志类型
const (
	RTULOG_SPLIT                         = "##" // 分隔符
	RTULOG_TIMELOCALATION                = "**" // 时间位置符
	RTULOG_TYPE_STRING                   = "1"  // 字符串
	RTULOG_TYPE_BYTE                     = "2"  // 报文
	RTULOG_TYPE_INVALID                  = "3"  // 鉴权失败
	RTULOG_TYPE_DATACENTERSWITCH         = "4"  // 所属数据中心切换
	RTULOG_TYPE_FlatCtrlProtoVersionNot3 = "5"  // 平台内不是3.0协议的集控
)

const (
	Ctr        = 0 // 集控器
	RtuLogType = 1 // 集控器日志类型
	Uuid       = 2 // uuid
	CtrlName   = 3 // 集控名称（给日志分析使用）
)

const (
	Third_CtrAddrIndex    = 0 // 集控地址下标
	Third_CtrlVendorIndex = 1 // 桩企下标
	Third_RtuLogTypeIndex = 2 // 日志类型下标
	Third_UuidIndex       = 3 // uuid
)

// setflag值
const (
	SETFLAG_ENTER byte = 255 // 255表示进入应急；
	SETFLAG_EXIT  byte = 1   // 1表示关闭集控的应急充电功能。

	SETFLAG_ENTER_STR = "255"
	SETFLAG_EXIT_STR  = "1"

	BMSIFSEND_STR = "1" // 1表示发送bms 0x26到kafka
)

// 数据传输方向
const (
	DirectionUp   = 1 // 上传
	DirectionDown = 2 // 下载
)

// 集控重复上线超过5次
const RepeatedOnlineNum = 5

// 函数返回值定义
const (
	SUCCESS = 1 // 成功
	FAIL    = 0 // 失败
)

// 应急充电配置信息值定义
const (
	OPEN  = 255 // 成功
	CLOSE = 1   // 失败
)

// 充电模式响应结果
const (
	SETSUCCESS = 255 // 成功
	SETFAIL    = 1   // 失败
)

// 群充策略
const (
	FAILED    = 1   // 失败
	STRATEGYA = 2   // 设备已经处于策略A
	STRATEGYB = 3   // 策略B
	STRATEGYC = 4   // 策略C
	SUCCESSED = 255 // 成功
)

// 设备升级
const (
	UPDATESUCCESSED = 0   // 成功
	UPDATEFAILED    = 255 // 失败

	ConfirmFailure = 0 // 失败
	ConfirmSuccess = 1 // 成功

)

// 订单错误
const (
	BATTERYPOWERDOWNYEAR = 2012
)

const (
	PILETYPENAME = "桩"
	CTRLTYPENAME = "集控"
	DEVTYPENAME  = "设备"

	NORMALLEVALE   = 0 // 普通级别
	SERIOUSLEVEL   = 1 // 严重级别
	SECONDARYLEVEL = 2 // 次要级别

	BUSSINESSTYPECTRL     = 0 // 集控
	BUSSINESSTYPEPLATFORM = 1 // 平台
	BUSSINESSTYPEOTHER    = 2 // 其他的不需要关心的

	AUTOSWITCH   = "自动切换"
	MANUALSWITCH = "手动切换"
)

// 文件类型
const (
	LOG          = 1 // 日志
	CONFIGFILE   = 8 // 配置文件
	PROGRAMEFILE = 3 // 程序文件
)

// 设备索引
const (
	CTRL = 250 // 集控
)

// 地锁自动设置
const (
	ALLOWAUTOSET = 1 // 1允许自动升锁
	NOTAUTOSET   = 2 // 2禁止自动升锁
)

// redis标记
const (
	DEVKEY                                 = "DevKey:"
	DEVSTATUS                              = "DevStatus:"
	BOWDEVSTATUS                           = "BowDevStatus"
	DEVCANSTATUS                           = "DevCanStatus:"
	DevProtocolVersion                     = "DevProtocolVersion:"
	SendDingDingLastTime                   = "SendDingDingLastTime"
	SendDingDingProtocolVersionErrLastTime = "SendDingDingProtocolVersionErrLastTime"
	SendDingDingRepeateOnlineLastTime      = "SendDingDingRepeateOnlineLastTime"
	//	DEVEMERGECSTATUS = "DevEmergStatus:"
	DEVMODELLASTTIME         = "DevModelLastTime:"
	CHARGEBOWMODELLASTTIME   = "ChargeBowModeLastTime:"
	OPERCHARGEBOWCTRLADDRESS = "CtrlAddress"
	DEVCTRLIST               = "DEVCtrlAddrList"

	CHARGEBOWBINDHASH = "ChargeBowBindHash"

	CHARGEBOWDEVCONF     = "ChargeBowDevConf:"
	CHARGEBOWMEASURECONF = "ChargeBowMeasureConf:"

	ALLCTRLEMERQENCYSTATUS = "AllCtrlEmergencyStatus"

	OPERTEL                    = "OperTEL"
	UPDATETIME                 = "time"
	OPERUPDATETIME             = "OperTime"
	OPERBOWCTRLADDRESS         = "CtrlAddress"
	HOSTNAME                   = "hostname"
	OPERHOSTNAME               = "OperHostname"
	STATUS                     = "status"
	OPERSTATUS                 = "OperStatus"
	EPOCH                      = "epoch"
	LastCallAllMonitorDataTime = "LastCallAllMonitorDataTime"
	OPEREPOCH                  = "OperEpoch"
	CompanyCode                = "CompanyCode"
	BillStrategy               = "BillStrategy"
	StrategyUpdateTime         = "StrategyUpdateTime"
	CTRLSWVER                  = "ctrlswversion"
	STAID                      = "StaID"
	CTRLNAME                   = "Name"
	STANAME                    = "StaName"
	FORCEEMERGENCY             = "forceEmergency"
	BMSIFSEND                  = "BmsIfSend"
	DATACENTER                 = "DataCenter"
	CTRLKEY                    = "Key"
	CODE                       = "Code"
	HASHIDHEAD                 = "Controller:"
	HASHIDTAIL                 = ":StaticInfo"
	CONTROLLERCODE             = "ControllerCode"
	CONtROLLERSTATIONID        = "ControllerStationID"
	HASHKEY                    = "HashKey"
	IFSENDOPERDATA             = "IsSupportOps"
	IFKEEPBLOCK                = "ifKeepBlock"         // 一直保持阻断
	DefaultCustomerInfo        = "defaultCustomerInfo" // 默认用户信息字段

	IfAlarmRestrain       = "IfAlarmRestrain" // 是否屏蔽钉钉告警   1表示是，0或者没有表示否
	CanList               = "CanList"
	StaIndustryType       = "StaIndustryType"       // 场站类型
	ProtocolVersion       = "ProtocolVersion"       // 平台内集控协议类型
	BUSEMERGENCY          = "BusEmergencyStatus"    // 公交一键应急
	EMERGENCYSTATUS       = "EmergencyStatus"       // 应急状态 1：强制应急已下发；2：不是强制应急
	HEARTTIME             = "HeartTime"             // 集控心跳时间
	QrCodeRule            = "QrCodeRule"            // 二维码规则
	UPTransferRate        = "UpTransferRate"        // 上行报文速率
	DOWNTransferRate      = "DownTransferRate"      // 下行报文速率
	ControllerStationCode = "ControllerStationCode" // redis db5，集控所属场站编号
	ControllerStationName = "ControllerStationName" // 获取平台2.0协议的集控的场站名称
	PILEKEY               = "Pile:"
	PILECODE              = "PileCode"     // 终端编码
	CallMgcCycle          = "callMgcCycle" // 不同集控对应不同的总召周期
	RecvMqInfo            = "RecvMqInfo:"  // 接收mq消息存入redis db5中
	CtrlAddress           = "CtrlAddress"  // 集控器地址
	CanIndex              = "CanSN"        // 枪地址

	VERTIFYREDISCONNECTIONFLAG_KEY   = "SystemParameterNoDelete" // 程序验证redis是否可用的标记
	VERTIFYREDISCONNECTIONFLAG_VALUE = "available"               // 程序验证redis是否可用的标记

	ONLINECTRLNUMS = "OnlineCtrlNums" // 各个数据中心在线集控数量
	ONLINENUMSTR   = "OnlineNum"
	ONLINESTR      = "Online"
	OFFLINESTR     = "Offline"

	// 集控的配置项
	ChargingEquipment                                                     = "ChargingEquipment"                                                     // 充电设备
	ChargingService                                                       = "ChargingService"                                                       // 充电服务
	GetNetworkAddress                                                     = "GetNetworkAddress"                                                     // 获取网络地址
	ChargingCommunication                                                 = "ChargingCommunication"                                                 // 充电通信
	IntelligentOperationAndMaintenance                                    = "IntelligentOperationAndMaintenance"                                    // 智能运维
	EquipmentManagement                                                   = "EquipmentManagement"                                                   // 设备管理
	GPIO                                                                  = "GPIO"                                                                  // GPIO
	Upgrade                                                               = "Upgrade"                                                               // 升级
	ElectricMeter                                                         = "ElectricMeter"                                                         // 电表
	LCDDisplay                                                            = "LCDDisplay"                                                            // LCD显示屏
	SerialDisplay                                                         = "SerialDisplay"                                                         // 串口显示屏
	ActiveProtection                                                      = "ActiveProtection"                                                      // 主动防护
	ChargingBow                                                           = "ChargingBow"                                                           // 充电弓
	BoxChangeSecond                                                       = "BoxChangeSecond"                                                       // 箱变2.0
	CardReader                                                            = "CardReader"                                                            // 读卡器
	LoadConstraint                                                        = "LoadConstraint"                                                        // 负荷约束
	ParkingLock                                                           = "ParkingLock"                                                           // 车位锁
	Scanner                                                               = "Scanner"                                                               // 扫码器
	Thermometer                                                           = "Thermometer"                                                           // 温湿度计
	TicketMachine                                                         = "TicketMachine"                                                         // 小票机
	CtrlStationName                                                       = "CtrlStationName"                                                       // 站名称
	BasePlateType                                                         = "BasePlateType"                                                         // 底板类型
	VersionNumber                                                         = "VersionNumber"                                                         // 版本号
	DNS                                                                   = "DNS"                                                                   // DNS
	ScreenPassword                                                        = "ScreenPassword"                                                        // 屏密码
	NumberOfDCPiles                                                       = "NumberOfDCPiles"                                                       // 直流桩数量
	NumberOfExchangePiles                                                 = "NumberOfExchangePiles"                                                 // 交流桩数量
	NumberOfThreePhaseACPiles                                             = "NumberOfThreePhaseACPiles"                                             // 三相交流桩数量
	TimeServerAddress                                                     = "TimeServerAddress"                                                     // 对时服务器地址
	VPNServerAddress                                                      = "VPNServerAddress"                                                      // VPN服务器地址
	TimeZoneFileSettings                                                  = "TimeZoneFileSettings"                                                  // 时区文件设置
	DoubleGunChargingEnable                                               = "DoubleGunChargingEnable"                                               // 版本号
	VINStartsChargingAfterDisconnection                                   = "VINStartsChargingAfterDisconnection"                                   // 断网后VIN启动充电
	LocalEndButton                                                        = "LocalEndButton"                                                        // 本地结束按钮
	CardType                                                              = "CardType"                                                              // 卡类型
	VINAutomaticCharging                                                  = "VINAutomaticCharging"                                                  // VIN自动充电
	CreditCardAutomaticCharging                                           = "CreditCardAutomaticCharging"                                           // 刷卡自动充电
	VehicleInformationIdentification                                      = "VehicleInformationIdentification"                                      // 车辆信息识别VIN/车牌号
	MeasurementType                                                       = "MeasurementType"                                                       // 计量类型（电表类型）
	SmallTicketMachineEnabled                                             = "SmallTicketMachineEnabled"                                             // 小票机使能
	LanguageSelection                                                     = "LanguageSelection"                                                     // 语言选择
	CalculateLineLossSwitch                                               = "CalculateLineLossSwitch"                                               // 计算线损开关
	EmergencyEnable                                                       = "EmergencyEnable"                                                       // 使能
	VINAuthenticationEnable                                               = "VINAuthenticationEnable"                                               // VIN鉴权使能
	CardNumberAuthenticationEnable                                        = "CardNumberAuthenticationEnable"                                        // 卡号鉴权使能
	LicensePlateNumberAuthenticationEnable                                = "LicensePlateNumberAuthenticationEnable"                                // 车牌号鉴权使能
	DefaultUserAuthenticationIsEnabled                                    = "DefaultUserAuthenticationIsEnabled"                                    // 默认用户鉴权使能
	DefaultUserType                                                       = "DefaultUserType"                                                       // 默认用户类型
	DefaultUserID                                                         = "DefaultUserID"                                                         // 默认用户ID
	TriggerEmergencyStateDisconnectionTime                                = "TriggerEmergencyStateDisconnectionTime"                                // 触发应急状态断网时间
	EmergencyChargingMaximumDuration                                      = "EmergencyChargingMaximumDuration"                                      // 应急充电最长持续时间
	MaximumNumberOfEmergencyChargingOrders                                = "MaximumNumberOfEmergencyChargingOrders"                                // 应急充电最大订单数
	Lan0Ip                                                                = "Lan0Ip"                                                                // ip
	Lan0SubnetMask                                                        = "Lan0SubnetMask"                                                        // 子网掩码
	Lan0Gateway                                                           = "Lan0Gateway"                                                           // 网关
	Lan1Pp                                                                = "Lan1Pp"                                                                // ip
	Lan1SubnetMask                                                        = "Lan1SubnetMask"                                                        // 子网掩码
	Lan2Gateway                                                           = "Lan2Gateway"                                                           // 网关
	Server0IP1                                                            = "Server0IP1"                                                            // 服务器IP1
	Server0Port1                                                          = "Server0Port1"                                                          // 服务器端口1
	Server0IP2                                                            = "Server0IP2"                                                            // 服务器IP2
	Server0Port2                                                          = "Server0Port2"                                                          // 服务器端口2
	Server0IP3                                                            = "Server0IP3"                                                            // 服务器IP3
	Server0Port3                                                          = "Server0Port3"                                                          // 服务器端口3
	Server0StationAddress                                                 = "Server0StationAddress"                                                 // 站地址
	Server0Key                                                            = "Server0Key"                                                            // 密钥
	Server0WhetherToEncrypt                                               = "Server0WhetherToEncrypt"                                               // 是否加密
	Server1IP                                                             = "Server1IP"                                                             // 服务器IP
	Server1Port                                                           = "Server1Port"                                                           // 服务器端口
	Server1StationAddress                                                 = "Server1StationAddress"                                                 // 站地址
	Server1Key                                                            = "Server1Key"                                                            // 密钥
	Server1WhetherToEncrypt                                               = "Server1WhetherToEncrypt"                                               // 是否加密
	UpgradeServerAddress                                                  = "UpgradeServerAddress"                                                  // 服务器地址
	UpgradeServerPortNumber                                               = "UpgradeServerPortNumber"                                               // 服务器端口号
	CAN0Address                                                           = "CAN0Address"                                                           // CAN地址
	CAN0BaudRate                                                          = "CAN0BaudRate"                                                          // 波特率
	CAN1Address                                                           = "CAN1Address"                                                           // CAN地址
	CAN1BaudRate                                                          = "CAN1BaudRate"                                                          // 波特率
	DI1                                                                   = "DI1"                                                                   // DI1
	DI2                                                                   = "DI2"                                                                   // DI2
	DI3                                                                   = "DI3"                                                                   // DI3
	DI4                                                                   = "DI4"                                                                   // DI4
	DI5                                                                   = "DI5"                                                                   // DI5
	DI6                                                                   = "DI6"                                                                   // DI6
	DI7                                                                   = "DI7"                                                                   // DI7
	DI8                                                                   = "DI8"                                                                   // DI8
	DI1EN                                                                 = "DI1EN"                                                                 // DI1EN
	DI2EN                                                                 = "DI2EN"                                                                 // DI2EN
	DI3EN                                                                 = "DI3EN"                                                                 // DI3EN
	DI4EN                                                                 = "DI4EN"                                                                 // DI4EN
	DI5EN                                                                 = "DI5EN"                                                                 // DI5EN
	DI6EN                                                                 = "DI6EN"                                                                 // DI6EN
	DI7EN                                                                 = "DI7EN"                                                                 // DI7EN
	DI8EN                                                                 = "DI8EN"                                                                 // DI8EN
	DO1                                                                   = "DO1"                                                                   // DO1
	DO2                                                                   = "DO2"                                                                   // DO2
	DOD1                                                                  = "DOD1"                                                                  // DOD1
	DOD2                                                                  = "DOD2"                                                                  // DOD2
	DOD3                                                                  = "DOD3"                                                                  // DOD3
	DOD4                                                                  = "DOD4"                                                                  // DOD4
	LightEnable                                                           = "LightEnable"                                                           // 使能
	LightStartHour                                                        = "LightStartHour"                                                        // 开始小时
	LightStartMinute                                                      = "LightStartMinute"                                                      // 开始分钟
	LightEndHour                                                          = "LightEndHour"                                                          // 结束小时
	LightEndMinute                                                        = "LightEndMinute"                                                        // 结束分钟
	MeterAddress                                                          = "MeterAddress"                                                          // 电表设置
	MeterVoltageRatio                                                     = "MeterVoltageRatio"                                                     // 电压变比
	MeterCurrentRatio                                                     = "MeterCurrentRatio"                                                     // 电流变比
	MeterEnable                                                           = "MeterEnable"                                                           // 使能
	MeterEquipmentType                                                    = "MeterEquipmentType"                                                    // 设备类型
	MeterFunctionType                                                     = "MeterFunctionType"                                                     // 功能类型
	OpenTheDoorAndPowerOffEnable                                          = "OpenTheDoorAndPowerOffEnable"                                          // 使能
	CANATo485Adress                                                       = "CANATo485Adress"                                                       // CAN地址
	CANATo485PGN                                                          = "CANATo485PGN"                                                          // PGN
	WuhanNonStandard300KwDoubleGun                                        = "WuhanNonStandard300KwDoubleGun"                                        // 武汉非标300Kw双枪
	DirectConnectionUsingLocalServerAddress                               = "DirectConnectionUsingLocalServerAddress"                               // 使用本地服务器地址直连（不使用网络地址功能）
	IgnoreVINUnderEmergencyChargingAndAutomaticallyApplyForChargingSwitch = "IgnoreVINUnderEmergencyChargingAndAutomaticallyApplyForChargingSwitch" // 应急充电下忽略VIN自动申请充电开关
	OperationSettingWithoutPowerCurve                                     = "OperationSettingWithoutPowerCurve"                                     // 无功率曲线执行操作设置
	FreezePowerPolicySettings                                             = "FreezePowerPolicySettings"                                             // 补冻结电量策略设置
	CANDataSource                                                         = "CANDataSource"                                                         // CAN数据源
	CANTemperatureEnable                                                  = "CANTemperatureEnable"                                                  // 使能
	ThresholdTemperature1                                                 = "ThresholdTemperature1"                                                 // 阈值温度1
	Duration1                                                             = "Duration1"                                                             // 持续时间1
	AdjustCurrentRatio1                                                   = "AdjustCurrentRatio1"                                                   // 调整电流比1
	ThresholdTemperature2                                                 = "ThresholdTemperature2"                                                 // 阈值温度2
	Duration2                                                             = "Duration2"                                                             // 持续时间2
	AdjustCurrentRatio2                                                   = "AdjustCurrentRatio2"                                                   // 调整电流比2
	ThresholdTemperature3                                                 = "ThresholdTemperature3"                                                 // 阈值温度3
	Duration3                                                             = "Duration3"                                                             // 持续时间3
	AdjustCurrentRatio3                                                   = "AdjustCurrentRatio3"                                                   // 调整电流比3
	ThresholdTemperature4                                                 = "ThresholdTemperature4"                                                 // 阈值温度4
	Duration4                                                             = "Duration4"                                                             // 持续时间4
	AdjustCurrentRatio4                                                   = "AdjustCurrentRatio4"                                                   // 调整电流比4
	RecoveryTemperature                                                   = "RecoveryTemperature"                                                   // 恢复温度
	RecoveryDuration                                                      = "RecoveryDuration"                                                      // 恢复持续时间
	ThreePhaseSetting                                                     = "ThreePhaseSetting"                                                     // 三相识别配置
	LoadConstraintEnable                                                  = "LoadConstraintEnable"                                                  // 使能
	LoadConstraintNumberOfCCU                                             = "LoadConstraintNumberOfCCU"                                             // CCU数量
	TotalStationPowerLimit                                                = "TotalStationPowerLimit"                                                // 场站总限制功率
	StationSafeChargingPower                                              = "StationSafeChargingPower"                                              // 场站安全充电功率
	ManuallySetChargingLimitPower                                         = "ManuallySetChargingLimitPower"                                         // 手动设置充电限制功率
	TheServerIssuesChargingLimitPower                                     = "TheServerIssuesChargingLimitPower"                                     // 服务器下发充电限制功率
	PowerMeterDynamicCalculationSettingLimitPowerEnable                   = "PowerMeterDynamicCalculationSettingLimitPowerEnable"                   // 电表动态计算设置限制功率使能
	ScreenInputSettingsLimitPowerEnable                                   = "ScreenInputSettingsLimitPowerEnable"                                   // 屏幕输入设置限制功率使能
	TheServerDeliversSettingsToLimitPowerEnable                           = "TheServerDeliversSettingsToLimitPowerEnable"                           // 服务器下发设置限制功率使能
	CrossPeakChargingLoadConstraintEnable                                 = "CrossPeakChargingLoadConstraintEnable"                                 // 负荷约束使能
	VehiclePrioritySchedulingEnable                                       = "VehiclePrioritySchedulingEnable"                                       // 使能
	VehiclePrioritySchedulingNumberOfCCU                                  = "VehiclePrioritySchedulingNumberOfCCU"                                  // CCU数量
	NumberOfPDUsUnderTheFirstCCU                                          = "NumberOfPDUsUnderTheFirstCCU"                                          // 第一个CCU下带的PDU数量
	NumberOfModulesUnderTheFirstCCU                                       = "NumberOfModulesUnderTheFirstCCU"                                       // 第一个CCU下带的模块数量
	PeakAndValleyTimeSetting                                              = "PeakAndValleyTimeSetting"                                              // 尖峰平谷
	PlugAndChargeEnable                                                   = "PlugAndChargeEnable"                                                   // 即插即充使能

	BillRealData                  = "ToolBillRealData"              // 订单实时信息
	MgcUpRateStatisticPrefixKey   = "DCB:MGC:MgcUpRateStatistic:"   // 微网单个集控的上行报文速率key前缀 key:$统计区间 type:hash fileName:mgcCtrlAddr value: 单位byte
	MgcDownRateStatisticPrefixKey = "DCB:MGC:MgcDownRateStatistic:" // 微网单个集控的下行报文速率key前缀 key:$统计区间 type:hash fileName:mgcCtrlAddr value: 单位byte
	KeHuaDeviceNumberMappingKey   = "SmcBKeHuaMapping"              // B模式科华映射key type:hash fieldName:deviceNumber

	PreposeInstanceStatusPrefixKey      = "DCB:CHARGE:STATUS:" // 充电前置程序存活性标识 key:PreposeInstanceStatusPrefixKey:${processName}:${hostname} value:hostname TTL
	MgcPreposeInstanceStatusPrefixKey   = "DCB:MGC:STATUS:"    // 微网前置程序存活性标识 key:MgcPreposeInstanceStatusPrefixKey:${processName}:${hostname} value:hostname TTL
	ThirdPreposeInstanceStatusPrefixKey = "DCB:THIRD:STATUS:"  // 第三方充电前置程序存活性标识
)

const DEFAULT_END_REASON = 79 // 默认充电结束原因

// 充电终止原因
var EndReasonOfCharging [300]string
var REASONSTR [MAXEXITREASON]string

func InitEndReasonOfCharging() {
	EndReasonOfCharging[1] = "BMS中止"
	EndReasonOfCharging[2] = "设备侧集中控制器终止"
	EndReasonOfCharging[3] = "人工设定条件中止"
	EndReasonOfCharging[4] = "手动点击直流机界面中止"
	EndReasonOfCharging[5] = "充电机故障中止"
	EndReasonOfCharging[6] = "连接器故障中止"
	EndReasonOfCharging[7] = "连接器拔出中止"
	EndReasonOfCharging[8] = "强制开关中止"
	EndReasonOfCharging[9] = "主接触器异常中止"
	EndReasonOfCharging[10] = "辅助接触器异常中止"
	EndReasonOfCharging[11] = "电子锁异常中止"
	EndReasonOfCharging[12] = "充电电压异常中止"
	EndReasonOfCharging[13] = "充电电流异常中止"
	EndReasonOfCharging[14] = "充电电流过流中止"
	EndReasonOfCharging[15] = "充电机最小电压大于BMS需求电压"
	EndReasonOfCharging[16] = "控制器通信故障中止"
	EndReasonOfCharging[17] = "电能表通信故障中止"
	EndReasonOfCharging[18] = "后台通讯中止"
	EndReasonOfCharging[19] = "SOC满中止"
	EndReasonOfCharging[20] = "系统模式类型转换中止"
	EndReasonOfCharging[21] = "断电中止"
	EndReasonOfCharging[22] = "主动防护电池过温中止"
	EndReasonOfCharging[23] = "主动防护电池低温中止"
	EndReasonOfCharging[24] = "主动防护电池温升异常终止"
	EndReasonOfCharging[25] = "主动防护电池过充中止"
	EndReasonOfCharging[26] = "BMS辅助电源异常中止"
	EndReasonOfCharging[27] = "BMS接触器开路故障中止"
	EndReasonOfCharging[28] = "主动防护BMS数据超范围中止"
	EndReasonOfCharging[29] = "PDU过温中止"
	EndReasonOfCharging[30] = "主动防护电池端口电压异常中止"
	EndReasonOfCharging[31] = "主动防护BMS通讯数据不刷新"
	EndReasonOfCharging[32] = "导引电路电压跳变中止"
	EndReasonOfCharging[33] = "拔枪中止"
	EndReasonOfCharging[34] = "设备过温中止"
	EndReasonOfCharging[35] = "过流中止"
	EndReasonOfCharging[36] = "过压中止"
	EndReasonOfCharging[37] = "欠压中止"
	EndReasonOfCharging[38] = "CP线低压故障中止"
	EndReasonOfCharging[39] = "漏电故障中止"
	EndReasonOfCharging[40] = "充电枪座过温中止"
	EndReasonOfCharging[41] = "继电器异常中止"
	EndReasonOfCharging[42] = "CAN地址冲突中止"
	EndReasonOfCharging[43] = "枪头过温中止"
	EndReasonOfCharging[44] = "模块开启失败中止"
	EndReasonOfCharging[49] = "未满足启动条件终止"
	EndReasonOfCharging[51] = "主动防护电池单体过压中止"
	EndReasonOfCharging[52] = "主动防护电池整包过压中止"
	EndReasonOfCharging[53] = "电表校验错误中止"
	EndReasonOfCharging[54] = "BMS通讯超时中止"
	EndReasonOfCharging[55] = "电池电压超过充电机输出最大电压"
	EndReasonOfCharging[56] = "风扇故障中止"
	EndReasonOfCharging[57] = "急停中止"
	EndReasonOfCharging[58] = "充电机绝缘异常中止"
	EndReasonOfCharging[59] = "DC模块不响应（改名为“充电机无可用模块”）"
	EndReasonOfCharging[60] = "车辆BMS粘连中止"
	EndReasonOfCharging[61] = "未获取车辆VIN码或车牌号中止"
	EndReasonOfCharging[62] = "熔断器故障中止"
	EndReasonOfCharging[63] = "模块匹配异常中止"
	EndReasonOfCharging[64] = "配置中禁止充电"
	EndReasonOfCharging[65] = "设备主动防护终止"
	EndReasonOfCharging[66] = "车辆电池达到目标SOC中止"
	EndReasonOfCharging[67] = "车辆达到总电压目标值中止"
	EndReasonOfCharging[68] = "车辆达到单体电压目标值中止"
	EndReasonOfCharging[69] = "车辆bms绝缘故障中止"
	EndReasonOfCharging[70] = "车辆输出连接器过温故障中止"
	EndReasonOfCharging[71] = "车辆BMS元件、输出连接器过温中止"
	EndReasonOfCharging[72] = "车辆充电连接器故障中止"
	EndReasonOfCharging[73] = "车辆电池组温度过高故障中止   "
	EndReasonOfCharging[74] = "车辆高压继电器故障中止"
	EndReasonOfCharging[75] = "车辆监测点CC2电压检测故障中止"
	EndReasonOfCharging[76] = "车辆其他故障"
	EndReasonOfCharging[77] = "车辆电流过大中止"
	EndReasonOfCharging[78] = "车辆电压异常中止"
	EndReasonOfCharging[DEFAULT_END_REASON] = "BMS中止"
	EndReasonOfCharging[83] = "温差过大中止"
	EndReasonOfCharging[84] = "SOC数据异常保护中止"
	EndReasonOfCharging[91] = "云平台指令中止"
	EndReasonOfCharging[92] = "其他平台指令中止"
	EndReasonOfCharging[93] = "集控器刷卡中止"
	EndReasonOfCharging[94] = "设备刷卡中止"
	EndReasonOfCharging[95] = "集控器界面按钮中止"
	EndReasonOfCharging[96] = "限制金额充电中止"
	EndReasonOfCharging[97] = "限制电量充电中止"
	EndReasonOfCharging[98] = "限制时间充电中止"
	EndReasonOfCharging[99] = "错峰充电中止"
	EndReasonOfCharging[100] = "异常电量中止"
	EndReasonOfCharging[101] = "门磁触发中止"
	EndReasonOfCharging[102] = "集控主动防护中止"
	EndReasonOfCharging[103] = "用户拔枪终止"
	EndReasonOfCharging[104] = "充电机主动终止"
	EndReasonOfCharging[105] = "功率曲线终止"
	EndReasonOfCharging[106] = "自启动设备终止"
	EndReasonOfCharging[107] = "枪头过温终止"
	EndReasonOfCharging[108] = "终端响应失败"
	EndReasonOfCharging[109] = "柱挂式直流终端按钮结束充电"
	EndReasonOfCharging[110] = "冻结点超过限制值终止"
	EndReasonOfCharging[128] = "单体电池欠压保护中止"
	EndReasonOfCharging[129] = "放电过流保护中止"
	EndReasonOfCharging[130] = "过放保护中止"
	EndReasonOfCharging[131] = "SOC过低保护中止"
	EndReasonOfCharging[132] = "电流反向保护中止"
	EndReasonOfCharging[134] = "短路故障中止"
	EndReasonOfCharging[135] = "车辆高低压类型识别失败中止"
	EndReasonOfCharging[136] = "枪头反接中止"
	EndReasonOfCharging[133] = "电池组欠压保护中止"
	EndReasonOfCharging[181] = "车载屏中止"
	EndReasonOfCharging[182] = "弓急停中止"
	EndReasonOfCharging[183] = "弓故障中止"
	EndReasonOfCharging[184] = "PLC离线中止"
	EndReasonOfCharging[185] = "WIFI连接异常"
	EndReasonOfCharging[256] = "充电功率异常终止"
}

// 错误码
var ErrorCode [136]string

func InitErrorCode() {
	ErrorCode[1] = "用户余额不足，请充值后再充电"
	ErrorCode[2] = "用户卡未绑定"
	ErrorCode[3] = "您的用户已冻结，无法开启充电"
	ErrorCode[4] = "未找到用户卡"
	ErrorCode[5] = "未找到用户"
	ErrorCode[6] = "未找到终端"
	ErrorCode[7] = "通信故障"
	ErrorCode[8] = "已经充电，不能重复充电"
	ErrorCode[9] = "该电站尚未运营"
	ErrorCode[10] = "终端正在服务中,请选择其他终端"
	ErrorCode[11] = "充电枪未与车连接"
	ErrorCode[12] = "通信超时"
	ErrorCode[13] = "未找到用户账户"
	ErrorCode[14] = "未找到计费策略"
	ErrorCode[15] = "未找到充电业务信息"
	ErrorCode[17] = "开始充电异常"
	ErrorCode[18] = "时间段不能超过24小时，请重设"
	ErrorCode[19] = "终端离网，请选择其他终端"
	ErrorCode[20] = "您最多只能同时进行三个充电业务"
	ErrorCode[21] = "操作过于频繁，请稍后再试"
	ErrorCode[22] = "电站已运营不允许充电"
	ErrorCode[23] = "轮充设备不支持刷卡充电"
	ErrorCode[24] = "终端故障不允许充电"
	ErrorCode[25] = "此终端为调度充电，不允许app充电"
	ErrorCode[26] = "运营电站，不允许测试充电"
	ErrorCode[27] = "该电站不支持直接充"
	ErrorCode[28] = "该电站没有初始电量"
	ErrorCode[29] = "获取正在消费的代金卷异常"
	ErrorCode[30] = "获取代金卷余额异常"
	ErrorCode[31] = "此代金卷不抵扣电费或者服务费"
	ErrorCode[32] = "此代金卷费用限制不正确"
	ErrorCode[33] = "获取正在消费的信用账户异常"
	ErrorCode[34] = "获取信用账户余额帐异常"
	ErrorCode[35] = "获取正在消费的现金账户异常"
	ErrorCode[36] = "获取现金余额账异常"
	ErrorCode[37] = "创建订单异常"
	ErrorCode[38] = "结束充电失败"
	ErrorCode[39] = "初始电量异常"
	ErrorCode[40] = "APP只支持收费电站的轮充充电"
	ErrorCode[41] = "该电站收费不允许直接充电"
	ErrorCode[42] = "该终端初始电量大于最大计量值"
	ErrorCode[43] = "终端未响应"
	ErrorCode[44] = "有未支付的充电订单,不允许开启充电"
	ErrorCode[45] = "北京公车改革刷卡充电中国移动身份确认失败"
	ErrorCode[46] = "获取电站电价高峰时段信息异常"
	ErrorCode[47] = "强制暂停充电失败"
	ErrorCode[48] = "强制开启充电失败"
	ErrorCode[50] = "终端不存在地锁"
	ErrorCode[51] = "下发降锁命令失败"
	ErrorCode[52] = "下发升锁命令失败"
	ErrorCode[53] = "强制创建订单异常"
	ErrorCode[54] = "充电机故障,请选择其他终端"
	ErrorCode[55] = "BMS故障,请选择其他终端"
	ErrorCode[56] = "连接故障,请选择其他终端"
	ErrorCode[57] = "APP定时充电不支持轮充设备充电"
	ErrorCode[58] = "易通卡刷卡申请创建电子钱包记录异常"
	ErrorCode[59] = "电子钱包支付未完成"
	ErrorCode[60] = "电子钱包补扣未支付失败"
	ErrorCode[61] = "未获取到车架号信息"
	ErrorCode[62] = "车牌号不匹配"
	ErrorCode[63] = "未获取到电子钱包记录"
	ErrorCode[64] = "未获取到运营商"
	ErrorCode[65] = "冻结优惠卷失败"
	ErrorCode[66] = "检查终端是否含有正在充电订单异常"
	ErrorCode[67] = "更新充电明细异常"
	ErrorCode[68] = "结束充电通知执行充电业务结算异常"
	ErrorCode[69] = "结束充电通知执行逻辑异常"
	ErrorCode[70] = "您最多只能同时进行一个充电业务"
	ErrorCode[71] = "根据车牌号未获取到车辆信息"
	ErrorCode[72] = "检查更新未完全支付状态异常"
	ErrorCode[73] = "账号未绑定，请重新登录"
	ErrorCode[75] = "获取平台账户失败"
	ErrorCode[76] = "不能停止他人启动的订单"
	ErrorCode[77] = "检查客户未支付订单异常"
	ErrorCode[78] = "企业信用账户有未还款的账单"
	ErrorCode[79] = "检查企业信用账户还款情况异常"
	ErrorCode[80] = "该电站不支持限制充电"
	ErrorCode[81] = "限制金额不能小于零元"
	ErrorCode[82] = "限制电量不能小于零度"
	ErrorCode[83] = "此终端已被预约，请选择其他终端进行充电"
	ErrorCode[84] = "获取不到折扣策略"
	ErrorCode[87] = "创建客户信息失败"
	ErrorCode[88] = "创建客户信息异常"
	ErrorCode[89] = "创建客户卡信息失败"
	ErrorCode[90] = "创建客户卡信息异常"
	ErrorCode[91] = "获取用户卡信息异常"
	ErrorCode[92] = "枪未插好"
	ErrorCode[93] = "电池已充满"
	ErrorCode[94] = "车辆电池温度过高"
	ErrorCode[95] = "车辆电池温度过低"
	ErrorCode[96] = "车辆电池温度变化异常"
	ErrorCode[97] = "车辆BMS接触器故障"
	ErrorCode[98] = "车辆电池端口电压异常"
	ErrorCode[99] = "车辆控制器故障"
	ErrorCode[100] = "车辆电池过压"
	ErrorCode[101] = "充电机无可用模块"
	ErrorCode[102] = "车辆故障"
	ErrorCode[103] = "车辆BMS故障"
	ErrorCode[104] = "副枪不允许开启充电"
	ErrorCode[105] = "副枪不允许结束充电"
	ErrorCode[106] = "终端状态未知"
	ErrorCode[107] = "计费策略不完整"
	ErrorCode[108] = "应急模式时不能从云端开启充电"
	ErrorCode[109] = "该电站不支持经济充电"
	ErrorCode[110] = "该用户未交押金"
	ErrorCode[111] = "检查用户是否交押金异常"
	ErrorCode[112] = "经济充电不支持轮充"
	ErrorCode[113] = "您还未预约，无法开启充电"
	ErrorCode[114] = "客户不满足场站对于充电客户的限制"
	ErrorCode[116] = "该终端非营业时间不允许充电"
	ErrorCode[117] = "充电机切换中不允许充电"
	ErrorCode[120] = "该终端非营业时间不允许充电"
	ErrorCode[121] = "充电机切换中不允许充电"
	ErrorCode[122] = "车辆已充满"
	ErrorCode[123] = "此车位已被使用，请选择其他终端进行充电"
	ErrorCode[124] = "充电账户和预约账户不一致"
	ErrorCode[125] = "充电账户和使用地锁的账户不一致"
	ErrorCode[126] = "您的车已经被限制了充电范围,请到可以允许的电站充电"
	ErrorCode[127] = "终端不支持您的充电启动方式"
	ErrorCode[128] = "终端不支持定时充电"
	ErrorCode[129] = "支付账户限制的VIN不存在"
	ErrorCode[130] = "支付账户与VIN绑定的支付账户不一致"
	ErrorCode[131] = "第三方终端只允许通过APP扫码或调度充电"
	ErrorCode[132] = "第三方终端设备校验异常"
	ErrorCode[133] = "地锁处于升起状态不允许充电"
	ErrorCode[134] = "电站不允许使用芝麻信用"
	ErrorCode[135] = "获取终端信息异常"
}

const (
	MeterAddrLen = 12
	ERRTIMESTR   = "200000000000"

	INT32_MAX = int32(^uint32(0) >> 1)

	EXCEPTIONVALUE = -1.0 // 表示异常值
	NULLMETERVALUE = -2.0 // 表示没有电表值
	ZEROVALUE      = 0.0  // 表示值为0

	PARAM_NULLMETERVALUE = -2 // 表示没有电表值
	PARAM_ZEROVALUE      = 0  // 表示值为0
)

// 0x09 公交站应急指令
const (
	ENTER_EMERGENCY = 1
	OUTER_EMERGENCY = 2
)

// 日志级别
const (
	EXCLEVEL_WARNING = "Warning"
	EXCLEVEL_ERROR   = "Error"
)

func InitOffLineReasonDesc() {
	REASONSTR[NETERROR] = "网络中断"
	REASONSTR[HEARTTIMEOUTERROR] = "心跳超时"
	REASONSTR[ACTIVEDOWN] = "云端主动断开"
	REASONSTR[DATACENTERSWITCH] = "所属数据中心切换"
}

const CHARGEPROCESSNAME = "prot_tcp_protobuf"
const HLHTTELDID = "395815801" // 互联互通teld标识

// ///////  第三方桩的报文头的格式定义///////////
const (
	// 蜗牛之家
	SNAILHEAD1        = 0xAA // 起始域
	SNAILHEAD2        = 0xF7 // 起始域
	SNAILHEADLEN      = 8    // 报文头长
	SNAILLENINDEX     = 2    // 长度域起始下标
	SNAILHEADFRAMELEN = 2    // 整个报文长度字节数
	SNAILVERSION      = 0x10
	SNAILSendSeqIndex = 5 // 序列号
	SNAILMSGINDEX     = 6 // 命令字起始下标

	// 追日
	SSEHEAD1        = 0xAA // 起始域
	SSEHEAD2        = 0xF5 // 起始域
	SSEHEADLEN      = 8    // 报文头长
	SSELENINDEX     = 2    // 长度域起始下标
	SSEHEADFRAMELEN = 2    // 整个报文长度字节数
	SSEVERSION      = 0x23
	SSESendSeqIndex = 5 // 序列号
	SSEMSGINDEX     = 7 // 命令字起始下标

	// 科华
	KEHUAHHEAD1        = 0x4B // 起始域
	KEHUAHHEAD2        = 0x48 // 起始域
	KEHUAHEADLEN       = 34   // 报文头长
	KEHUAHLENINDEX     = 2    // 长度域起始下标
	KEHUAHHEADFRAMELEN = 2    // 整个报文长度字节数
	DevAddrIndex       = 9    // 设备串号下标
	DevAddrLen         = 20   // 设备串号长度

	// 厦门市政
	XIAMENHEAD1        = 0x47 // 起始域
	XIAMENHEAD2        = 0x47 // 起始域
	XIAMENHEADLEN      = 14   // 报文头长
	XIAMENLENINDEX     = 2    // 长度域起始下标
	XIAMENHEADFRAMELEN = 2    // 整个报文长度字节数
	SourceIndex        = 5    //
	SourceLen          = 4    //

	// 合康
	HeKangCtrlHeader1 = 0x7F
	HeKangCtrlHeader2 = 0x79

	HeKangHeader1          = 0x7E // 起始域
	HeKangHeader2          = 0x68 // 起始域
	HeKangHeaderLength     = 14   // 报文头长
	HeKangFrameLengthIndex = 4    // 长度域起始下标
	HeKangFrameLength      = 4    // 整个报文长度字节数
	HeKangAppId            = 0x04 // 合康应用ID
	HeKangVersionId        = 0x01 // 合康协议版本
	HeKangEndCode          = 0xCC // 合康终止码

	// 京能
	JingNengHead1        = 0x75
	JingNengHead2        = 0x72
	JingNengHeadLen      = 11
	JingNengLenIdx       = 2
	JingNengHeadFrameLen = 2
	JingNengEnd          = 0x68

	// 云快充
	YUNKCHEAD              = 0x68 // 起始域
	YUNKCSTARTFIXEDHEADLEN = 13   // 协议启动帧固定长度
	YUNKCHEARTBEATLEN      = 6    // 心跳帧是固定6个字节
	YUNKCMINLEN            = 13   // 报文最短是13个字节(报文头长)
	YUNKCPROTOTYPE         = 1    // 104协议
	YUNKCHEARTBEAT         = 4    // 心跳U帧
	YUNKCHEARTBEATThird    = 0x43 // 心跳帧第三个字节
	YUNKCHEARTBEATFour     = 0x00 // 心跳帧第四个字节
	YUNKCHEARTBEATFouth    = 0x00 // 心跳帧第五个字节
	YUNKCHEARTBEATSix      = 0x00 // 心跳帧第六个字节
	YUNKCDEVTYPE           = 2    // 桩类型

	// 聚电
	COSTARHEAD              = 0x68 // 起始域
	COSTARSTARTFIXEDHEADLEN = 21   // 协议启动帧固定长度
	COSTARHEARTBEATLEN      = 6    // 心跳帧是固定6个字节
	COSTARMINLEN            = 13   // 报文最短是13个字节(报文头长)
	COSTARPROTOTYPE         = 1    // 104协议
	COSTARHEARTBEAT         = 4    // 心跳U帧
	COSTARHEARTBEATThird    = 0x43 // 心跳帧第三个字节
	COSTARHEARTBEATFour     = 0x00 // 心跳帧第四个字节
	COSTARHEARTBEATFouth    = 0x00 // 心跳帧第五个字节
	COSTARHEARTBEATSix      = 0x00 // 心跳帧第六个字节
	COSTARDEVTYPE           = 2    // 桩类型

	// 新版云快充协议
	NEWYUNKCHEAD         = 0x68 // 起始域
	NewYUNKCHEADLENGTH   = 6
	NEWYUNKCLENINDEX     = 1 // 长度域起始下标
	NEWYUNKCHEADFRAMELEN = 1 // 整个报文长度字节数
	NEWYUNKCMSGINDEX     = 5 // 命令字起始下标

	// 德恒
	DEHENGHEAD         = 0x68 // 起始域
	DEHENGHEADLENGTH   = 6
	DEHENGLENINDEX     = 1 // 长度域起始下标
	DEHENGHEADFRAMELEN = 1 // 整个报文长度字节数
	DEHENGMSGINDEX     = 5 // 命令字起始下标

	// 盛弘
	SINEXCELHEAD1        = 0xAA // 起始域
	SINEXCELHEAD2        = 0xF5 // 起始域
	SINEXCELHEADLEN      = 8    // 报文头长
	SINEXCELLENINDEX     = 2    // 长度域起始下标
	SINEXCELHEADFRAMELEN = 2    // 整个报文长度字节数
	SINEXCELVERSION      = 0x10
	SINEXCELSendSeqIndex = 5 // 序列号
	SINEXCELMSGINDEX     = 6 // 命令字起始下标

	// 万泊
	ENPLUSHEAD1        = 0xAA // 起始域
	ENPLUSHEAD2        = 0xF5 // 起始域
	ENPLUSHEADLEN      = 8    // 报文头长
	ENPLUSLENINDEX     = 2    // 长度域起始下标
	ENPLUSHEADFRAMELEN = 2    // 整个报文长度字节数
	ENPLUSVERSION      = 0x10
	ENPLUSSendSeqIndex = 5 // 序列号
	ENPLUSMSGINDEX     = 6 // 命令字起始下标

	// 绿能
	LVNENGHEAD1        = 0xAA // 起始域
	LVNENGHEAD2        = 0xF5 // 起始域
	LVNENGHEADLEN      = 8    // 报文头长
	LVNENGENINDEX      = 2    // 长度域起始下标
	LVNENGHEADFRAMELEN = 2    // 整个报文长度字节数
	LVNENGInfo         = 0x02
	LVNENGSendSeqIndex = 5 // 序列号
	LVNENGMSGINDEX     = 6 // 命令字起始下标

	// 国耀
	GUOYAOHEAD1         = 0x68 // 起始域
	GUOYAOHEAD2         = 0x68 // 起始域
	GUOYAOHEADLEN       = 15   // 报文头长
	GUOYAOLENINDEX      = 13   // 长度域起始下标
	GUOYAOSendSeqIndex  = 5    // 序列号
	GUOYAOENDCODE       = 0x16 // 帧尾
	GUOYAOTERIDLEN      = 4    // 设备串号长度
	GUOYAOLHEADFRAMELEN = 2    // 整个报文长度字节数
	GUOYAOENDLENGTH     = 2

	// 科陆（车电网）
	KELUHEAD       = 0x68    // 起始字节
	KELUTAIL       = 0x16    // 结束字节
	KELULENINDEX   = 11      // 长度域起始下标
	KELUDATAINDEX  = 13      // 数据域起始下标
	KELUHEADLEN    = 13      // 报文头字节数
	KELUTAILLEN    = 2       // 报文尾字节数
	KELUMAXDATALEN = 2 << 10 // 报文数据域最大长度

	// 奥能
	AONENG_HEAD      = 0x68 // 起始域
	AONENG_TAIL      = 0x16 // 结束域
	AONENG_LEN_IDX   = 1    // 长度域起始下标
	AONENG_LEN_BYTES = 2    // 长度域字节数
	AONENG_TAIL_LEN  = 2    // 报文尾长度

	// 特来电
	TELDHEAD1        = 0xAA // 起始域
	TELDHEAD2        = 0xF5 // 起始域
	TELDHEADLEN      = 26   // 报文头长
	TELDLENINDEX     = 4    // 长度域起始下标
	TELDHEADFRAMELEN = 4    // 整个报文长度字节数
	TELDVERSION      = 0x10
	TELDMSGINDEX     = 8 // 命令字起始下标

	// 特来电
	XingXingHEAD1    = 0x32 // 起始域
	XingXingHEADLEN  = 12   // 报文头长
	XingXingLENINDEX = 11   // 长度域起始下标
	XingXingFRAMELEN = 4    // 整个报文长度字节数
	// TELDVERSION      = 0x10
	XingXingMSGINDEX = 2 // 命令字起始下标

	// 小桔
	XjuStartIdx     = 0 // 起始域下标
	XjuStartLen     = 2 // 起始域长度
	XjuLengthIdx    = 2
	XjuLengthLen    = 2
	XjuVersionIdx   = 4
	XjuVersionLen   = 4
	XjuSerialNumIdx = 8
	XjuSerialNumLen = 4
	XjuCodeIdx      = 12
	XjuCodeLen      = 2
	XjuPayLoadIdx   = 14
	XjuCheckSumLen  = 1

	// 永联
	YLHEAD1        = 0xAA // 起始域
	YLHEAD2        = 0xF5 // 起始域
	YLHEADLEN      = 8    // 报文头长
	YLLENINDEX     = 2    // 长度域起始下标
	YLHEADFRAMELEN = 2    // 整个报文长度字节数
	YLVERSION      = 0x10
	YLSendSeqIndex = 5 // 序列号
	YLMSGINDEX     = 6 // 命令字起始下标

	// 安恒
	ANHENGHEAD1        = 0xAA // 起始域
	ANHENGHEAD2        = 0xF5 // 起始域
	ANHENGHEADLEN      = 8    // 报文头长
	ANHENGLENINDEX     = 2    // 长度域起始下标
	ANHENGHEADFRAMELEN = 2    // 整个报文长度字节数
	ANHENGVERSION      = 0x10
	ANHENGSendSeqIndex = 5 // 序列号
	ANHENGMSGINDEX     = 6 // 命令字起始下标

	// 科士达
	KSD_HEAD_1    = 0x75 // 起始域字节1
	KSD_HEAD_2    = 0x72 // 起始域字节2
	KSD_TAIL      = 0x68 // 结束域
	KSD_LEN_IDX   = 2    // 长度域起始下标
	KSD_LEN_BYTES = 2    // 长度域字节数

	// 京能
	JINGNENG_HEAD_1    = 0x75 // 起始域字节1
	JINGNENG_HEAD_2    = 0x72 // 起始域字节2
	JINGNENG_TAIL      = 0x68 // 结束域
	JINGNENG_LEN_IDX   = 2    // 长度域起始下标
	JINGNENG_LEN_BYTES = 2    // 长度域字节数

	// 新页
	XinYeHEAD1        = 0xAA // 起始域
	XinYeHEAD2        = 0xF5 // 起始域
	XinYeHEADLEN      = 8    // 报文头长
	XinYeLENINDEX     = 2    // 长度域起始下标
	XinYeHEADFRAMELEN = 2    // 整个报文长度字节数
	XinYeVERSION      = 0x10
	XinYeSendSeqIndex = 5 // 序列号
	XinYeMSGINDEX     = 6 // 命令字起始下标

	// 派诺V1.2
	PAINUO_V1_2_HEAD_BYTE       = 0x68
	PAINUO_V1_2_HEAD_2_IDX      = 9
	PAINUO_V1_2_LEN_START       = 11                  // 长度域起始索引
	PAINUO_V1_2_LEN_END         = 13                  // 长度域结束索引
	PAINUO_V1_2_TAIL_BYTE       = 0x16                // 帧尾字节
	PAINUO_V1_2_LEN_BEFORE_DATA = PAINUO_V1_2_LEN_END // 数据域前的字节数
	PAINUO_V1_2_LEN_AFTER_DATA  = 2                   // 数据域后的字节数
	PAINUO_V1_2_MAX_DATA_LEN    = 10 << 10            // 限制一帧的数据域最大长度为10k

	// 光法
	GUANGFA_HEAD_1    = 0xAA // 起始域
	GUANGFA_HEAD_2    = 0xF5 // 起始域
	GUANGFA_HEAD_LEN  = 8    // 报文头长
	GUANGFA_LEN_INDEX = 2    // 长度域起始下标
	GUANGFA_LEN_BYTES = 2    // 长度域字节数
	GUANGFA_MSG_INDEX = 6    // 命令字起始下标

	// 中兴
	ZHONGXING_HEAD_1     = 0xFA // 起始域
	ZHONGXING_HEAD_2     = 0xF5 // 起始域
	ZHONGXING_HEAD_LEN   = 16   // 报文头长
	ZHONGXING_LEN_OFFSET = 14   // 长度域的值加上该值等于整个报文的字节长度
	ZHONGXING_LEN_INDEX  = 2    // 长度域起始下标
	ZHONGXING_LEN_BYTES  = 2    // 长度域字节数
	ZHONGXING_MSG_INDEX  = 15   // 命令字起始下标

	// 丁旺
	DINGWANG_HEAD_1    = 0xAA // 起始域
	DINGWANG_HEAD_2    = 0xF5 // 起始域
	DINGWANG_HEAD_LEN  = 8    // 报文头长
	DINGWANG_LEN_INDEX = 2    // 长度域起始下标
	DINGWANG_LEN_BYTES = 2    // 长度域字节数
	DINGWANG_MSG_INDEX = 6    // 命令字起始下标

	// 东旭
	DONGXUHEAD         = 0x68 // 起始域
	DONGXUHEADLENGTH   = 6
	DONGXULENINDEX     = 1 // 长度域起始下标
	DONGXUHEADFRAMELEN = 1 // 整个报文长度字节数
	DONGXUMSGINDEX     = 5 // 命令字起始下标

	// 方智
	FZHEAD         = 0x68 // 起始域
	FZHEADLENGTH   = 6
	FZLENINDEX     = 1 // 长度域起始下标
	FZHEADFRAMELEN = 1 // 整个报文长度字节数
	FZMSGINDEX     = 5 // 命令字起始下标

	// 东旭
	TAIHONGHEAD         = 0x68 // 起始域
	TAIHONGHEADLENGTH   = 6
	TAIHONGLENINDEX     = 1 // 长度域起始下标
	TAIHONGHEADFRAMELEN = 1 // 整个报文长度字节数
	TAIHONGMSGINDEX     = 5 // 命令字起始下标

	// 电网104
	IEC104_START         = 0x68 // 起始域值
	IEC104_ASDU_META_LEN = 6    // ASDU描述数据长度

	// 聚能（104）
	JUNENG_ID_FRAME_LEN   = 21 // 协议标识帧长度
	JUNENG_PROTOTYPE_104  = 1  // 104协议类型标识
	JUNENG_CONNTYPE_INDEX = 2  // 连接类型字段下标
	JUNENG_CONNTYPE_PILE  = 2  // 连接类型-离散充电桩
	JUNENG_APCI_CTRL_3    = 4  // APCI中控制域第三个字节的下标

	// 捷电通（104）
	JIEDT_ID_FRAME_LEN        = 25   // 协议标识帧长度
	JIEDT_ID_FRAME_TYPE       = 0xFF // 协议标识帧中的启动标识
	JIEDT_ID_FRAME_TYPE_INDEX = 3    // 启动标识下标
	JIEDT_LEN_INDEX           = 1    // 长度域下标
	JIEDT_LEN_BYTES           = 2    // 长度域字节数
	JIEDT_APCI_CTRL_3_INDEX   = 5    // APCI中控制域第三个字节的下标

	// 能瑞（104）
	NENGRUI_ID_FRAME_LEN        = 25   // 协议标识帧长度
	NENGRUI_ID_FRAME_TYPE       = 0xFF // 协议标识帧中的启动标识
	NENGRUI_ID_FRAME_TYPE_INDEX = 3    // 启动标识下标
	NENGRUI_LEN_INDEX           = 1    // 长度域下标
	NENGRUI_LEN_BYTES           = 2    // 长度域字节数
	NENGRUI_APCI_CTRL_3_INDEX   = 5    // APCI中控制域第三个字节的下标

	// 赫胜
	HESHENG_HEAD_1     = 0x55 // 起始域
	HESHENG_HEAD_2     = 0xAA // 起始域
	HESHENG_HEAD_LEN   = 9    // 报文头长
	HESHENG_LEN_OFFSET = 7    // 长度域的值加上该值等于整个报文的字节长度
	HESHENG_LEN_INDEX  = 2    // 长度域起始下标
	HESHENG_LEN_BYTES  = 2    // 长度域字节数
	HESHENG_SEQ_INDEX  = 6    // 帧序号域下标

	// 橙电
	CHENGDIANHEAD1        = 0xAA // 起始域
	CHENGDIANHEAD2        = 0xF5 // 起始域
	CHENGDIANHEADLEN      = 8    // 报文头长
	CHENGDIANLENINDEX     = 2    // 长度域起始下标
	CHENGDIANHEADFRAMELEN = 2    // 整个报文长度字节数
	CHENGDIANVERSION      = 0x10
	CHENGDIANSendSeqIndex = 5 // 序列号
	CHENGDIANMSGINDEX     = 6 // 命令字起始下标

	// 科大v3.6
	KEDA_V3_6_HEAD       = 0x68 // 起始域
	KEDA_V3_6_HEAD_LEN   = 17   // 报文头长
	KEDA_V3_6_LEN_OFFSET = 7    // 长度域的值加上该值等于整个报文的字节长度
	KEDA_V3_6_LEN_INDEX  = 2    // 长度域起始下标
	KEDA_V3_6_LEN_BYTES  = 2    // 长度域字节数
	KEDA_V3_6_CTRL_INDEX = 5    // 控制域下标

	// 科大v3.8
	KEDA_V3_8_HEAD       = 0x68 // 起始域
	KEDA_V3_8_HEAD_LEN   = 17   // 报文头长
	KEDA_V3_8_LEN_OFFSET = 7    // 长度域的值加上该值等于整个报文的字节长度
	KEDA_V3_8_LEN_INDEX  = 2    // 长度域起始下标
	KEDA_V3_8_LEN_BYTES  = 2    // 长度域字节数
	KEDA_V3_8_CTRL_INDEX = 5    // 控制域下标

	// 英威腾
	YINGWT_HEAD_1    = 0x23 // 起始域1
	YINGWT_HEAD_2    = 0x23 // 起始域2
	YINGWT_HEAD_LEN  = 9    // 报文头长
	YINGWT_LEN_INDEX = 3    // 报文体长度域起始下标
	YINGWT_LEN_BYTES = 2    // 报文体长度域字节数
	YINGWT_CS_START  = 2    // 校验和计算起始下标

	// 云杉
	YUNSHAN_HEAD_1     = 0xAA // 起始域
	YUNSHAN_HEAD_2     = 0xF8 // 起始域
	YUNSHAN_HEAD_2_OLD = 0xF5 // 起始域
	YUNSHAN_HEAD_LEN   = 8    // 报文头长
	YUNSHAN_LEN_INDEX  = 2    // 长度域起始下标
	YUNSHAN_LEN_BYTES  = 2    // 长度域字节数
	YUNSHAN_CS_START   = 6    // 校验和计算起始下标

	// 明瑞
	MINGRUI_HEAD      = 0x68 // 起始域
	MINGRUI_HEAD_LEN  = 6
	MINGRUI_LEN_INDEX = 1 // 长度域起始下标
	MINGRUI_MSG_INDEX = 5 // 命令字起始下标
	MINGRUI_META_LEN  = 4 // 长度域的值 - 数据域长度 || 报文长度 - 长度域的值
	MINGRUI_CS_BEGIN  = 2 // 校验和计算起始下标

	// 特变电工
	TEBD_HEAD      = 0x68 // 起始域
	TEBD_HEAD_LEN  = 4
	TEBD_LEN_INDEX = 1    // 长度域起始下标
	TEBD_LEN_BYTES = 2    // 长度域字节数
	TEBD_CS_BEGIN  = 3    // 校验和计算起始下标
	TEBD_TAIL      = 0x16 // 结束域

	// 和信瑞通
	HEXIN_HEAD         = 0x68 //起始域
	HEXIN_HEAD_LEN     = 23   //报文头长度
	HEXIN_CS_BEGIN     = 3    //校验和计算起始下标
	HEXIN_LEN_INDEX    = 1    //长度域起始下标
	HEXIN_LEN_BYTES    = 2    //长度域字节数
	HEXIN_SERIAL_INDEX = 3    //序列号域下标

	// 充电联盟
	CCTIA_HEAD_1    = 0x06 // 起始域
	CCTIA_HEAD_2    = 0x01 // 起始域
	CCTIA_HEAD_LEN  = 24   // 报文头长度
	CCTIA_LEN_INDEX = 2    // 长度域起始下标
	CCTIA_LEN_BYTES = 2    // 长度域字节数

	// 长园深瑞
	SHENRUI_HEAD         = 0x68 // 起始域
	SHENRUI_HEAD_LEN     = 29   // 报文头长度
	SHENRUI_LEN_INDEX    = 1    // 长度域下标
	SHENRUI_CS_BYTES     = 2    // 校验和域长度
	SHENRUI_CS_CAL_BEGIN = 2    // 校验和计算起始下标
	SHENRUI_LEN_OFFSET   = SHENRUI_LEN_INDEX + 1 + SHENRUI_CS_BYTES
)

const (
	MAXCTRLLEN = 13
	MINCTRLLEN = 11
	STACODELEN = 10 // 电站编号长度
)

const (
	CTRLDEVINDEX   = 250
	PDUMinDevIndex = 181
	PDUMaxDevIndex = 212
	CCUMinDevIndex = 231
	CCUMaxDevIndex = 240
)

// 配置项码
var SettingsCode map[int32]string

func init() {
	SettingsCode = make(map[int32]string)
	SettingsCode[101001] = ChargingEquipment
	SettingsCode[101002] = ChargingService
	SettingsCode[101003] = GetNetworkAddress
	SettingsCode[101004] = ChargingCommunication
	SettingsCode[101005] = IntelligentOperationAndMaintenance
	SettingsCode[101006] = EquipmentManagement
	SettingsCode[101007] = GPIO
	SettingsCode[101008] = Upgrade
	SettingsCode[101009] = ElectricMeter
	SettingsCode[101010] = LCDDisplay
	SettingsCode[101011] = SerialDisplay
	SettingsCode[101012] = ActiveProtection
	SettingsCode[101013] = ChargingBow
	SettingsCode[101014] = BoxChangeSecond
	SettingsCode[101015] = CardReader
	SettingsCode[101016] = LoadConstraint
	SettingsCode[101017] = ParkingLock
	SettingsCode[101018] = Scanner
	SettingsCode[101019] = Thermometer
	SettingsCode[101020] = TicketMachine
	SettingsCode[102001] = CtrlStationName
	SettingsCode[102002] = BasePlateType
	SettingsCode[102003] = VersionNumber
	SettingsCode[102004] = DNS
	SettingsCode[102005] = ScreenPassword
	SettingsCode[102006] = NumberOfDCPiles
	SettingsCode[102007] = NumberOfExchangePiles
	SettingsCode[102008] = NumberOfThreePhaseACPiles
	SettingsCode[102009] = TimeServerAddress
	SettingsCode[102010] = VPNServerAddress
	SettingsCode[102011] = TimeZoneFileSettings
	SettingsCode[201001] = DoubleGunChargingEnable
	SettingsCode[201002] = VINStartsChargingAfterDisconnection
	SettingsCode[201003] = LocalEndButton
	SettingsCode[201004] = CardType
	SettingsCode[201005] = VINAutomaticCharging
	SettingsCode[201006] = CreditCardAutomaticCharging
	SettingsCode[201007] = VehicleInformationIdentification
	SettingsCode[201008] = MeasurementType
	SettingsCode[201009] = SmallTicketMachineEnabled
	SettingsCode[201010] = LanguageSelection
	SettingsCode[201011] = CalculateLineLossSwitch
	SettingsCode[202001] = EmergencyEnable
	SettingsCode[202002] = VINAuthenticationEnable
	SettingsCode[202003] = CardNumberAuthenticationEnable
	SettingsCode[202004] = LicensePlateNumberAuthenticationEnable
	SettingsCode[202005] = DefaultUserAuthenticationIsEnabled
	SettingsCode[202006] = DefaultUserType
	SettingsCode[202007] = DefaultUserID
	SettingsCode[202008] = TriggerEmergencyStateDisconnectionTime
	SettingsCode[202009] = EmergencyChargingMaximumDuration
	SettingsCode[202010] = MaximumNumberOfEmergencyChargingOrders
	SettingsCode[301001] = Lan0Ip
	SettingsCode[301002] = Lan0SubnetMask
	SettingsCode[301003] = Lan0Gateway
	SettingsCode[302001] = Lan1Pp
	SettingsCode[302002] = Lan1SubnetMask
	SettingsCode[302003] = Lan2Gateway
	SettingsCode[303001] = Server0IP1
	SettingsCode[303002] = Server0Port1
	SettingsCode[303003] = Server0IP2
	SettingsCode[303004] = Server0Port2
	SettingsCode[303005] = Server0IP3
	SettingsCode[303006] = Server0Port3
	SettingsCode[303007] = Server0StationAddress
	SettingsCode[303008] = Server0Key
	SettingsCode[303009] = Server0WhetherToEncrypt
	SettingsCode[304001] = Server1IP
	SettingsCode[304002] = Server1Port
	SettingsCode[304003] = Server1StationAddress
	SettingsCode[304004] = Server1Key
	SettingsCode[304005] = Server1WhetherToEncrypt
	SettingsCode[305001] = UpgradeServerAddress
	SettingsCode[305002] = UpgradeServerPortNumber
	SettingsCode[306001] = CAN0Address
	SettingsCode[306002] = CAN0BaudRate
	SettingsCode[307001] = CAN1Address
	SettingsCode[307002] = CAN1BaudRate
	SettingsCode[401001] = DI1
	SettingsCode[401002] = DI2
	SettingsCode[401003] = DI3
	SettingsCode[401004] = DI4
	SettingsCode[401005] = DI5
	SettingsCode[401006] = DI6
	SettingsCode[401007] = DI7
	SettingsCode[401008] = DI8
	SettingsCode[402001] = DI1EN
	SettingsCode[402002] = DI2EN
	SettingsCode[402003] = DI3EN
	SettingsCode[402004] = DI4EN
	SettingsCode[402005] = DI5EN
	SettingsCode[402006] = DI6EN
	SettingsCode[402007] = DI7EN
	SettingsCode[402008] = DI8EN
	SettingsCode[403001] = DO1
	SettingsCode[403002] = DO2
	SettingsCode[403003] = DOD1
	SettingsCode[403004] = DOD2
	SettingsCode[403005] = DOD3
	SettingsCode[403006] = DOD4
	SettingsCode[404001] = LightEnable
	SettingsCode[404002] = LightStartHour
	SettingsCode[404003] = LightStartMinute
	SettingsCode[404004] = LightEndHour
	SettingsCode[404005] = LightEndMinute
	SettingsCode[406001] = OpenTheDoorAndPowerOffEnable
	SettingsCode[407001] = CANATo485Adress
	SettingsCode[407002] = CANATo485PGN
	SettingsCode[501001] = WuhanNonStandard300KwDoubleGun
	SettingsCode[501002] = DirectConnectionUsingLocalServerAddress
	SettingsCode[501003] = IgnoreVINUnderEmergencyChargingAndAutomaticallyApplyForChargingSwitch
	SettingsCode[501004] = OperationSettingWithoutPowerCurve
	SettingsCode[501005] = FreezePowerPolicySettings
	SettingsCode[501006] = CANDataSource
	SettingsCode[502001] = CANTemperatureEnable
	SettingsCode[502002] = ThresholdTemperature1
	SettingsCode[502003] = Duration1
	SettingsCode[502004] = AdjustCurrentRatio1
	SettingsCode[502005] = ThresholdTemperature2
	SettingsCode[502006] = Duration2
	SettingsCode[502007] = AdjustCurrentRatio2
	SettingsCode[502008] = ThresholdTemperature3
	SettingsCode[502009] = Duration3
	SettingsCode[502010] = AdjustCurrentRatio3
	SettingsCode[502011] = ThresholdTemperature4
	SettingsCode[502012] = Duration4
	SettingsCode[502013] = AdjustCurrentRatio4
	SettingsCode[502014] = RecoveryTemperature
	SettingsCode[502015] = RecoveryDuration
	SettingsCode[503000] = ThreePhaseSetting
	SettingsCode[504001] = LoadConstraintEnable
	SettingsCode[504002] = LoadConstraintNumberOfCCU
	SettingsCode[504003] = TotalStationPowerLimit
	SettingsCode[504004] = StationSafeChargingPower
	SettingsCode[504005] = ManuallySetChargingLimitPower
	SettingsCode[504006] = TheServerIssuesChargingLimitPower
	SettingsCode[504007] = PowerMeterDynamicCalculationSettingLimitPowerEnable
	SettingsCode[504008] = ScreenInputSettingsLimitPowerEnable
	SettingsCode[504009] = TheServerDeliversSettingsToLimitPowerEnable
	SettingsCode[505001] = CrossPeakChargingLoadConstraintEnable
	SettingsCode[506001] = VehiclePrioritySchedulingEnable
	SettingsCode[506002] = VehiclePrioritySchedulingNumberOfCCU
	SettingsCode[506003] = NumberOfPDUsUnderTheFirstCCU
	SettingsCode[506004] = NumberOfModulesUnderTheFirstCCU
	SettingsCode[507000] = PeakAndValleyTimeSetting
	SettingsCode[611001] = PlugAndChargeEnable
}

// 启动充电方式 协议附录G
const (
	StartChargeWay_CloudStart          = 1  // 扫码充
	StartChargeWay_CardStart           = 2  // 刷卡充电
	StartChargeWay_VINStart            = 3  // VIN充电
	StartChargeWay_SchedulingStart     = 4  // 调度充电
	StartChargeWay_SchedulingCardStart = 5  // 调度刷卡充电
	StartChargeWay_ReversedCloudStart  = 17 // 反向扫码
)

// mqtt_bridge_server constant
const (
	MqttUrlSplit                     = "," // url分隔符
	MqttClientConnectionTimeout      = 5   // 单位S
	MqttClientPingTimeout            = 5   // 单位S
	MqttSubScribeTimeout             = 20  // 单位S
	MqttTopicSplit                   = "/" // mqtt topic 分隔符
	MqttMaxProtChannelLen            = 50  // 每个集控器的规约上行channel大小
	MqttMaxCommandChannelLen         = 100 // 接收mq下行命令通道大小
	MqttCmdChanBlockRestartThreshold = 50  // 阻塞临界值

	HeartTimeoutOfflineMode = 1 // 心跳检测离线
	LWTOfflineMode          = 2 // 遗嘱检测离线

	NotNeed = 0
	Need    = 1
	Want    = 2

	ElecPriceType = 4 // 电价类型  尖峰平谷
	MobileLen     = 11

	CtrlTypeAC = 1 // 交流
	CtrlTypeDc = 2 // 直流
)

const (
	AuthFailed          = 1 // 鉴权失败
	RepeatedCtrlAddress = 2 // 重复上线
	InvalidVersion      = 3 // 非法版本(非3.0)
	InvalidDataCenter   = 4 // 数据中心非法不一致
	ObsoleteVersion     = 5 // 废弃版本
)

const (
	NoneTLS = 0 // 不使用TLS
	NeedTLS = 1 // 使用TLS

	Prod    = 1
	NotProd = 0

	VinPlateNotiftReq_DataTypeVIN   = 1 // VIN号
	VinPlateNotiftReq_DataTypePlate = 2 // 车牌号
	DefaultOpen                     = 1
	DefaultClose                    = 0
)

// 平台SG调用相关常量
const (
	GetCarInfoByVINFieldName   = "CarNo"      // VIN
	GetCarInfoByCarIDFieldName = "VehicleID"  // 车辆ID
	GetCarInfoByPlateFieldName = "CarLicense" // 车牌号

	GetCarInfoStateFailure = "0" // 失败
	GetCarInfoStateSuccess = "1" // 成功

	GetCarInfoValidFlag   = "1" // 有效
	GetCarInfoInValidFlag = "0" // 无效
)

// V2G
const (
	V2GSpecialFront = 1 // 特殊的V2G前置

	V2GResultDefault = 0 // 默认
	V2GResultTrue    = 1 // 成功
	V2GResultFalse   = 2 // 失败

	V2GGeneralFailureReason_UnrecognizedGun    = -1
	V2GGeneralFailureReason_MissingCarInitInfo = -2
	V2GGeneralFiilureReason_UnabelMeeetDemand  = -3

	V2GConfirmFiilureReason_UnrecognizedGun = -1
	V2GConfirmFiilureReason_MissingV2GCurve = -2
	V2GConfirmFiilureReason_MGCFault        = -3

	V2GRequestExpandHeaderKey = "Referer"
	V2GRequestExpandHeaderVal = "deri"
)

// 轮充
const (
	AlternateControlMode = 6
	StateTypeCharging    = 1
	StateTypeStarting    = 3
	StateTypePause       = 4
)

const (
	CanListBaseTTL   = 300 // 单位S canList 本地缓存基础的ttl
	CanListExpandTTL = 120 // 单位S canList 本地缓存可扩展的ttl
)

type OnlineCode int64

const (
	OnlineSuccess        OnlineCode = 1 // 上线成功
	OnlineMachineOffline OnlineCode = 2 // 上线的机器不存活
	OnlineConflict       OnlineCode = 3 // 其他机器为上线状态
	OnlineLocalConflict  OnlineCode = 4 // 本机重复上线
)

type OfflineCode int64

const (
	OfflineSuccess OfflineCode = 1 // 离线成功
	OfflineFailure OfflineCode = 2 // 离线失败
)

// CustomerLoadByCtrlAddr Redis 加载默认用户key格式
const (
	DefaultCustomerLoadByCtrlAddr    = "DevStatus:%s"       // %s 集控地址
	DefaultCustomerLoadByStationCode = "StationAddition:%s" // %s 站地址
)

type RedisHashFieldName string

const (
	DefaultCustomerFieldName = "defaultCustomer" // 默认用户字段名称
)

// 充电协议遥测数据类型
const (
	TELE_TYPE_U_A   = 0x01 // A相电压
	TELE_TYPE_U_B   = 0x02 // B相电压
	TELE_TYPE_U_C   = 0x03 // C相电压
	TELE_TYPE_U_D   = 0x04 // 直流电压
	TELE_TYPE_I_A   = 0x11 // A相电流
	TELE_TYPE_I_B   = 0x12 // B相电流
	TELE_TYPE_I_C   = 0x13 // C相电流
	TELE_TYPE_I_D   = 0x14 // 直流电流
	TELE_TYPE_POWER = 0x21 // 当前充电功率
	TELE_TYPE_METER = 0x31 // 当前电表读数
)

// mq topics
const (
	// 充电mq
	MQ_TOPIC_CHARGE_DOWN = "CpCommand" // 充电下行mq

	// 升级mq
	MQ_TOPIC_UPGRADE_DOWN     = "UpgradeCommand"   // 升级下行mq
	MQ_TOPIC_UPGRADE_UP       = "UpgradeCommandUp" // 升级上行mq --> 远程升级微服务
	MQ_TOPIC_UPGRADE_UP_THIRD = "ThirdUpgradeUp"   // 升级上行mq --> 远程升级前置

	// 智能运维mq
	MQ_TOPIC_OPER_DOWN = "CtrlMonitorCommandDown" // 智能运维下行mq
	MQ_TOPIC_OPER_UP   = "CtrlMonitorCommandUp"   // 智能运维上行mq --> 智能运维微服务
)

// 时间格式
const (
	DATETIME_FMT       = "2006-01-02 15:04:05"
	DATETIME_MILLI_FMT = "2006-01-02 15:04:05.000"
)

const (
	ONLINE_MANAGER_KEY_PREFIX = "ONLINE_MANAGER:"
	MEASURE_MODEL_KEY_PREFIX  = "MeasureModel:"
	CONFIG_MODEL_KEY_PREFIX   = "ConfigModel:"
)

// 远程配置下发的配置项数据类型
const (
	PARAM_TYPE_INT32  = 1 // int32
	PARAM_TYPE_STRING = 2 // string
	PARAM_TYPE_DOUBLE = 3 // float64
	PARAM_TYPE_BOOL   = 4 // bool
)

const (
	METRICS_PATH = "/metrics"
)

// 登录报文中的集控状态
const (
	CTRL_STATUS_COMM     = 1 // 正常登录（没有历史订单）
	CTRL_STATUS_HAS_BILL = 2 // 有历史订单
)
