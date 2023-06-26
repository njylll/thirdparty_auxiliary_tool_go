package charge

import (
	"github.com/gin-gonic/gin"
	"github.com/njylll/thirdparty_auxiliary_tool_go/app"
	"github.com/njylll/thirdparty_auxiliary_tool_go/dto"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/e"
	"github.com/njylll/thirdparty_auxiliary_tool_go/service/chargeservice"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary 充电控制
// @Produce  json
// @Param billCode body ChargeCtrlParam true "ChargeCtrlParam"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /charge/ctrl [post]
func ChargeCtrl(c *gin.Context) {
	appG := app.Gin{C: c}

	//解析参数
	revBody := new(dto.ChargeCtrlParam)
	if err := c.ShouldBind(revBody); err != nil {
		logrus.Errorf("解析参数失败: err: %v", err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	// 检验集控地址
	if !checkCtrlAddr(revBody.CtrlAddr) {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	//调用service
	service := chargeservice.ChargeCtrlService{}

	if service.CtrlCharge(revBody) {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}

	// 充电操作失败
	appG.Response(http.StatusOK, e.ERROR, nil)
}
