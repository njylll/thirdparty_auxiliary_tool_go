package jwt

import (
	"github.com/njylll/thirdparty_auxiliary_tool_go/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
			code = e.ERROR_AUTH
		}

		//校验token
		if token == "" {
			code = e.ERROR_AUTH
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		//校验token不通过,返回错误
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
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
