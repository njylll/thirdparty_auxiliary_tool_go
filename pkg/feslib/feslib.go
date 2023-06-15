/*
 * Created by 丁德鑫  on 2017/10/9 10:22:07
 * ==============================================
 * 包说明：string、byte之间转换等公共函数
 * ==============================================
 * 南京德睿能源研究院版权所有
 * ==============================================
 */
package feslib

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/constdefine"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unicode"
	"github.com/bitly/go-simplejson"
)

const (
	LenVin = 17
)

/*
4字节大头byte转为uint32
*/
func ByteToUint(byteBuffer []byte) uint32 {
	return binary.BigEndian.Uint32(byteBuffer)
}

/*
2字节大头byte转为uint16
*/
func ByteToUint16(byteBuffer []byte) uint16 {
	return binary.BigEndian.Uint16(byteBuffer)
}

func ByteToUint16Little(byteBuffer []byte) uint16 {
	return binary.LittleEndian.Uint16(byteBuffer)
}

/*
uint32转为4字节小头byte
*/
func UintToByte(num uint32) []byte {
	x := num
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 整形转换成4字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 整形转换成4字节
func IntToBytesLittle(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

// 整形转换成2字节
func IntTo2Bytes(n int) []byte {
	x := int16(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 整形转换成小端2字节
func IntTo2LillteBytes(n int) []byte {
	x := int16(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

///*
//BCD码格式的站地址解析为字符串
//*/
//func BCDToCtrlStr(byteBuffer []byte) string {
//	var result bytes.Buffer
//	if len(byteBuffer) != 8 {
//		return ""
//	}
//	var str string
//	var i int = 0
//	for ; i < 5; i++ {
//		result.WriteString(strconv.Itoa(int(byteBuffer[i] >> 4 & 0x0F)))
//		result.WriteString(strconv.Itoa(int(byteBuffer[i] & 0x0F)))

//		str += string(byteBuffer[i] & 0x0F)
//	}

//	result.WriteString(string(byteBuffer[5] & 0xFF))
//	if 0x00 != byteBuffer[6] {
//		result.WriteString(strconv.Itoa(int(byteBuffer[6] >> 4 & 0x0F)))
//		result.WriteString(strconv.Itoa(int(byteBuffer[6] & 0x0F)))
//	}
//	if 0x00 != byteBuffer[7] {

//		result.WriteString(strconv.Itoa(int(byteBuffer[6] >> 4 & 0x0F)))
//		result.WriteString(strconv.Itoa(int(byteBuffer[6] & 0x0F)))
//		result.WriteString(strconv.Itoa(int(byteBuffer[7] >> 4 & 0x0F)))
//		result.WriteString(strconv.Itoa(int(byteBuffer[7] & 0x0F)))
//	}

//	//校验
//	match, _ := regexp.MatchString("[0-9]{10,10}[0-9a-zA-Z]{1,1}[0-9]{0,2}", result.String())
//	if match {
//		return result.String()
//	} else {
//		return ""
//	}
//}

/*
byte格式的站地址解析为字符串,并将结尾连续的空白符去掉
*/
func ByteToCtrlStr(byteBuffer []byte) string {

	if len(byteBuffer) != constdefine.CtrlAddrLen {
		return ""
	}
	return strings.TrimSpace(string(byteBuffer))
}

/*
byte格式的充电弓地址解析为字符串,并将结尾连续的空白符去掉
*/
func ByteToChargeBowStr(byteBuffer []byte) string {

	if len(byteBuffer) != constdefine.ChargeBowAddrLen {
		return ""
	}
	return strings.TrimSpace(string(byteBuffer))
}

// 获取小时与分 HHMM
func GetShortTimeStrBysec(timeInt int64) string {
	return time.Unix(timeInt, 0).Format("1504")
}

/*
将1970年以来的秒数时间转成直到分的字符串格式 201810241330
*/
func GetCMTimeStrBySec(timeInt int64) string {
	if 0 == timeInt {
		return constdefine.ERRTIMESTR
	}
	tb := time.Unix(timeInt, 0)
	return fmt.Sprintf("%04d%02d%02d%02d%02d", tb.Year(), tb.Month(), tb.Day(), tb.Hour(), tb.Minute())
}

/*
将1970年以来的秒数时间转成字符串格式
*/
func GetTimeStrBySec(timeInt int64) string {
	return time.Unix(timeInt, 0).Format("2006-01-02 15:04:05")
}

/*
将1970年以来的带毫秒的时间转成字符串格式
*/
func GetTimeStrByMillsec(timeInt int64) string {
	return time.Unix(timeInt/1e3, 0).Format("2006-01-02 15:04:05") + fmt.Sprintf(":%03d ", timeInt%1e3)
}

/*
将1970年以来的带毫秒的时间转成字符串格式
*/
func GetSimpleTimeStrByMillsec(timeInt int64) string {
	return time.Unix(timeInt/1e3, 0).Format("2006-01-02 15:04:05")
}

/*
将1970年以来的带毫秒的时间转成Time类型
*/
func GetTimeByMillsec(timeInt int64) time.Time {
	return time.Unix(timeInt/1e3, (timeInt%1e3)*1e7)
}

// 获取当天的字符串格式
func GetCurrentDayStr() string {
	nowtime := time.Now().UnixNano()
	return time.Unix(nowtime/1e9, 0).Format("20060102")
}

/*
将固定格式的字符串转成对应的Time类型，例如："2018-07-11 15:45:12"
*/
func GetTimeByTimeStr(timeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")                               //重要：获取时区
	tb, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) //使用模板在对应时区转化为time.time类型
	return tb
}

/*
将固定格式的字符串(带毫秒)转成对应的Time类型，例如："2018-07-11 15:45:12.062"
*/
func GetTimeByTimeStrWithMillsec(timeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")                                   //重要：获取时区
	tb, _ := time.ParseInLocation("2006-01-02 15:04:05.000", timeStr, loc) //使用模板在对应时区转化为time.time类型
	return tb
}

/*
将固定格式的字符串转成对应的Time类型，例如："201810231040"
*/
func GetTimeByTimeStrWithMin(timeStr string) (time.Time, error) {
	if len(timeStr) < 12 {
		return time.Time{}, errors.New("timeStr len is not 12")
	}
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	newTimeStr := timeStr[0:4] + "-" + timeStr[4:6] + "-" + timeStr[6:8] + " " + timeStr[8:10] + ":" + timeStr[10:12] + ":00.000"
	tb, _ := time.ParseInLocation("2006-01-02 15:04:05", newTimeStr, loc) //使用模板在对应时区转化为time.time类型
	return tb, nil
}

/*
带时间的打印函数
*/
func Log(v ...interface{}) {
	fmt.Print(" ", GetCurrentTimeStr())
	fmt.Println(v...)
}

/*
将获得当前时间转的字符串格式
*/
func GetCurrentTimeStr() string {
	nowtime := time.Now().UnixNano()
	return time.Unix(nowtime/1e9, 0).Format("2006-01-02 15:04:05") + fmt.Sprintf(":%03d ", nowtime%1e9/1e6)
}

/*
将获得当前时间转的字符串格式,没有空格，因为工作组通知如果有空格，就推不出来
*/
func GetCurrentTimeStrNoSpace() string {
	nowtime := time.Now().UnixNano()
	return time.Unix(nowtime/1e9, 0).Format("2006-01-02-15:04:05") + fmt.Sprintf(":%03d", nowtime%1e9/1e6)
}

// 获得当前目录，带/
func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

/*
判断目录是否存在
返回值：true  存在 ;  false 不存在
*/
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

/*
判断目录是否存在
返回值：true  存在 ;  false 不存在
*/
func Mkdir(path string) bool {
	if !IsDirExists(path) {
		errMkdir := os.MkdirAll(path, os.ModePerm) //生成多级目录
		if nil != errMkdir {
			fmt.Println(errMkdir)
			return false
		} else {
			fmt.Println("创建文件夹" + path + "成功")
		}
	}
	return true
}

// 获取单个文件的大小
func GetFileSize(fileWithPath string) int64 {
	fileInfo, err := os.Stat(fileWithPath)
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 将2006-01-02T15:04:05.*格式的时间 或者 将2006-01-02 15:04:05.*格式的时间 转为 2006-01-02 15:04:05.000
// 原因是因为平台返回的时间格式不固定，有的没有毫秒，有的毫秒位数不是3
func FormatDateTime(timeStr string) string {
	dateTime := ""
	index := strings.LastIndex(timeStr, ".")
	if -1 == index {
		dateTime = timeStr + ".000"
	} else {
		leftStr := timeStr[:index]
		rightStr := timeStr[index+1:]
		lenOfstr := len(rightStr)
		if lenOfstr == 1 {
			dateTime = leftStr + ".00" + rightStr
		} else if lenOfstr == 2 {
			dateTime = leftStr + ".0" + rightStr
		} else if lenOfstr == 3 {
			dateTime = timeStr
		} else {
			dateTime = "1970-01-01T00:00:00.000"
		}
	}
	standFormat := "2006-01-02 15:04:05.000" // 平台杨磊返回的时间没有T
	if strings.Contains(timeStr, "T") {      //这是EM返回的格式,包含T
		standFormat = "2006-01-02T15:04:05.000"
	}
	newTime, err := time.Parse(standFormat, dateTime)
	if err != nil {
		return newTime.Format("1970-01-01 00:00:00.000")
	}
	return newTime.Format("2006-01-02 15:04:05.000")
}

// TwoPrecision 得到精度为2的double值
func TwoPrecision(i int32) float64 {
	if 0 == i {
		return constdefine.NULLMETERVALUE //没有数值
	}
	if constdefine.INT32_MAX == i {
		return constdefine.EXCEPTIONVALUE //表示异常值
	}
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(i)*0.01), 64)
	if err != nil {
		return constdefine.EXCEPTIONVALUE //表示异常值
	}
	return value
}

// FourPrecision 得到精度为4的double值
func FourPrecision(i int32) float64 {
	return float64(i) * 0.0001
}

// GetOffLineDesc 返回退出的原因描述
func GetOffLineDesc(exitReason byte) string {
	if exitReason < 0 || exitReason > constdefine.MAXEXITREASON {
		return ""
	}
	return constdefine.REASONSTR[exitReason]
}

// 计算绝对值
func Int64Abs(n *int64) {
	y := *n >> 63
	*n = (*n ^ y) - y
}

// 对VIN中非法字符进行替换0处理，位数不足的也补0；空VIN除外
func VinValidRepair0(vin string) string {
	if "" == vin {
		return ""
	}
	newVin := ""
	for _, value := range vin {
		//判断每个字符是否是字母或者数字
		if unicode.IsLetter(value) || unicode.IsDigit(value) {
			newVin += string(value)
			continue
		} else {
			newVin += "0"
		}
	}

	for i := len(vin); i < LenVin; i++ {
		newVin += "0"
	}
	return newVin
}

// 需要处理的vin的数据
var messageVinMap map[string]string // key是结构体名称，value是filed的名称

func init() {
	messageVinMap = make(map[string]string)
	// 0x36
	messageVinMap["BmsNotifyReq"] = "Vin"

	// 0x16
	messageVinMap["BillUploadReq"] = "Bill"
	messageVinMap["BillInfo"] = "VIN"

	// 0x10
	messageVinMap["StartChargeReq"] = "Vin"

	//0x28
	messageVinMap["VinPlateNotiftReq"] = "Value"

	//0x26 vin
	messageVinMap["BmsParamNotifyReq"] = "Vin"

}

func ParseVinValue(data interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			Log(err)
		}
	}()
	if data == nil {
		return errors.New("the interface is nil")
	}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data).Elem()

	infName := ""
	if t.Kind() != reflect.Ptr {
		infName = reflect.PtrTo(t).Name()
	} else {
		infName = t.Elem().Name()
	}

	filedName := messageVinMap[infName]
	for filedName != "" {
		// 取出字段的结构和名称
		if v.Kind() == reflect.Ptr {
			t = reflect.TypeOf(v.Elem().FieldByName(filedName).Interface())
			v = v.Elem().FieldByName(filedName)
		} else {
			t = reflect.TypeOf(v.FieldByName(filedName).Interface())
			v = v.FieldByName(filedName)
		}

		// 如果是结构体或者指针继续找
		if t.Kind() == reflect.Struct || t.Kind() == reflect.Ptr {

			if t.Kind() != reflect.Ptr {
				infName = reflect.PtrTo(t).Name()
			} else {
				infName = t.Elem().Name()
			}

			filedName = messageVinMap[infName]
			if filedName == "" {
				return nil
			}

			continue
		}

		// 不是结构体直接进行设置string
		if v.Kind() == reflect.String {
			v.SetString(VinValidRepair0(v.String()))
		}

		return nil
	}

	return nil
}

type CtrlInfo struct {
	CtrlAddress  string
	StaID        string //所属电站编号
	StaName      string //所属电站名称
	OperMaintain string //所属运维公司编号
	YWID         string //运维负责人ID
	YWName       string //所属运维公司名称
	YWTel        string //运维负责人电话
	YW           string //运维负责人名称
}

func GetYwTelByCtrlAddr(url string, ctrList string) ([]CtrlInfo, error) {
	defer func() {
		if err := recover(); nil != err {
			Log("GetYwTelByCtrlAddr recover err", err)
			debug.PrintStack()
		}
	}()
	visiturl := make(map[string]interface{}, 3)
	visiturl["Type"] = "GetEcoChargingSpiData"
	visiturl["SGType"] = "SelectYwInfoList"
	visiturl["CtrlAdds"] = ctrList

	para, _ := json.Marshal(visiturl)
	str := fmt.Sprintf("%v", string(para))
	Url := url + "&condition=" + str
	resp, err := http.Get(Url)
	ctrlInfoSlice := []CtrlInfo{}
	if err != nil {
		Log("GetYwTelByCtrlAddr http get err!", err, ctrList)
	} else if resp != nil {
		defer func() {
			io.Copy(ioutil.Discard, resp.Body) // 必须要有这个
			resp.Body.Close()
		}()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("解析Http正文失败", err)
		} else {
			js, err := simplejson.NewJson(body)
			if err != nil {
				Log("解析Json失败:", err)
			} else if nil == js {
				Log("获取的数据为空,js:", js)
			} else {
				r, err2 := js.Get("return").Map()
				if err2 != nil {
					fmt.Println("map err", err2)
					return []CtrlInfo{}, err2
				}
				fmt.Println(r)
				result := fmt.Sprintf("%s", r["result"])
				if result != "success" {
					fmt.Println("result is not success!")
				} else {
					data, err3 := js.Get("return").Get("rows").Array()
					if err3 != nil {
						fmt.Println("map err3", err3)
						return []CtrlInfo{}, err2
					}

					for _, row := range data {
						if oneMap, ok := row.(map[string]interface{}); ok {
							ctrlInfo := CtrlInfo{}
							if nil != oneMap["CtrlAddress"] {
								ctrlInfo.CtrlAddress = oneMap["CtrlAddress"].(string)
							}
							if nil != oneMap["StaID"] {
								ctrlInfo.StaID = oneMap["StaID"].(string)
							}
							if nil != oneMap["StaName"] {
								ctrlInfo.StaName = oneMap["StaName"].(string)
							}
							if nil != oneMap["OperMaintain"] {
								ctrlInfo.OperMaintain = oneMap["OperMaintain"].(string)
							}
							if nil != oneMap["YWID"] {
								ctrlInfo.YWID = oneMap["YWID"].(string)
							}
							if nil != oneMap["YWName"] {
								ctrlInfo.YWName = oneMap["YWName"].(string)
							}
							if nil != oneMap["YWTel"] {
								ctrlInfo.YWTel = oneMap["YWTel"].(string)
							}
							if nil != oneMap["YW"] {
								ctrlInfo.YW = oneMap["YW"].(string)
							}
							ctrlInfoSlice = append(ctrlInfoSlice, ctrlInfo)
						}
					}

				}
			}
		}
	}
	return ctrlInfoSlice, nil
}

func SendWorkGroupNotify(url string, userId string, msg string) {
	visiturl := make(map[string]interface{}, 4)
	visiturl["UserId"] = userId
	visiturl["SGType"] = "PostDDbyUserId"
	visiturl["Type"] = "GetEcoChargingSpiData"
	visiturl["dingMes"] = msg

	para, _ := json.Marshal(visiturl)
	str := fmt.Sprintf("%v", string(para))
	Url := url + "&condition=" + str
	resp, err := http.Get(Url)
	if err != nil {
		Log("SendWorkGroupNotify http get err!", err, msg)
	} else if resp != nil {
		defer func() {
			io.Copy(ioutil.Discard, resp.Body) // 必须要有这个
			resp.Body.Close()
		}()
		_, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("解析Http正文失败", err)
		} else {
		}
	}
}

func CheckCtrladdrValid(ctrladdr string) bool {
	length := len(ctrladdr)
	if length > constdefine.MAXCTRLLEN || length < constdefine.MINCTRLLEN {
		return false
	}
	return true
}

// 判断集控地址是否含有非法字符
func CheckCtrlAddrErr(ctrladdr string) bool {
	for _, value := range ctrladdr {
		//判断每个字符是否是字母或者数字
		if unicode.IsLetter(value) || unicode.IsDigit(value) {
			continue
		} else {
			return false
		}
	}
	return true
}
