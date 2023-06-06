package login_service

import "github.com/EDDYCJY/go-gin-example/models"

type Login struct {
	Username string
	Password string
}

func (a *Login) Check() (bool, error, string) {
	return models.CheckAuth(a.Username, a.Password)
}
