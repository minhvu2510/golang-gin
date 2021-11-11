package jwt

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/minhvu2510/golang-gin/pkg/e"
	"github.com/minhvu2510/golang-gin/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		// token := c.Query("token")
		token_header := c.Request.Header.Get("x-access-token")
		fmt.Println(token_header)
		if token_header == "" {
			code = e.INVALID_PARAMS
		} else {
			// _, err := util.ParseToken(token)
			tkParse, err2 := util.ParseToken(token_header)
			fmt.Println(tkParse)
			fmt.Println(err2)
			if err2 != nil {
				switch err2.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			fmt.Println("----Token err----")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
