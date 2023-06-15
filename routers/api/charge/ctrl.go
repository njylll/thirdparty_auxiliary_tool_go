package charge

import (
	"github.com/gin-gonic/gin"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/app"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/e"
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

	appG.Response(http.StatusOK, e.SUCCESS, true)
}
