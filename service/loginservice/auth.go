package loginservice

import "github.com/njylll/thirdparty_auxiliary_tool_go/models"

type Login struct {
	Username string
	Password string
}

func (a *Login) Check() (bool, error, string) {
	return models.CheckAuth(a.Username, a.Password)
}
