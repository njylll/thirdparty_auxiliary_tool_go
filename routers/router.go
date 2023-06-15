package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/njylll/thirdparty_auxiliary_tool_go/routers/api/charge"

	_ "github.com/njylll/thirdparty_auxiliary_tool_go/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/njylll/thirdparty_auxiliary_tool_go/middleware/jwt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//静态资源目录
	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//登录接口
	r.POST("/login", api.PostLogin)

	//swagger接口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//除登录接口和swagger接口, 其他接口使用jwt鉴权
	group := r.Group("")
	group.Use(jwt.JWT())
	{
		//充电控制
		group.POST("/charge/ctrl", charge.ChargeCtrl)

	}

	return r
}
