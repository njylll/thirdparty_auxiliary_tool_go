package e

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "err",
	INVALID_PARAMS: "request param happen error",
	NOT_LOGIN:      "please login",
	LOGIN_ERR:      "login error",
	SYS_ERR:        "system error",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
