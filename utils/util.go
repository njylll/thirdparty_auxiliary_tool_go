package utils

import (
	"github.com/njylll/thirdparty_auxiliary_tool_go/setting"
)

// Setup Initialize the utils
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
