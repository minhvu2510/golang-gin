package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhvu2510/golang-gin/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	fmt.Println("--------init router--------")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/auth", api.GetAuth)
	return r
}
