package api

import (
	"github.com/njylll/thirdparty_auxiliary_tool_go/dto"
	"github.com/njylll/thirdparty_auxiliary_tool_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/app"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/e"
	"github.com/njylll/thirdparty_auxiliary_tool_go/service/loginservice"
)

// @Summary Post Login
// @Produce  json
// @Param loginParam body LoginParam true "loginParam"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /login [post]
func PostLogin(c *gin.Context) {
	appG := app.Gin{C: c}

	//解析参数
	revBody := new(dto.LoginParam)
	if err := c.ShouldBind(revBody); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	//查询是否有此用户
	loginService := loginservice.Login{Username: revBody.UserName, Password: revBody.PassWord}
	isExist, err, usrId := loginService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	//用户不存在返回鉴权失败
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenerateToken(usrId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
