package charge

import "github.com/njylll/thirdparty_auxiliary_tool_go/pkg/feslib"

// 判断集控器地址是否合法
func checkCtrlAddr(ctrlAddr string) bool {
	if ctrlAddr == "" {
		return false
	}
	return feslib.CheckCtrlAddrErr(ctrlAddr)
}
