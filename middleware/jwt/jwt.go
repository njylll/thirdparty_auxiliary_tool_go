package jwt

import (
	"github.com/njylll/thirdparty_auxiliary_tool_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/e"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		//取出token
		token, err := c.Cookie("token")
		if err != nil {
			code = e.LOGIN_ERR
		}

		//校验token
		if token == "" {
			code = e.LOGIN_ERR
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				code = e.LOGIN_ERR
			}
		}

		//校验token不通过,返回错误
		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		//校验token通过
		c.Next()
	}
}
