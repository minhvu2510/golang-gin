package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhvu2510/golang-gin/routers/api"
	"github.com/minhvu2510/golang-gin/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	fmt.Println("--------init router--------")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/auth", api.GetAuth)
	r.GET("/tags", v1.GetTags)
	// apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	// {
	// 	//get all tags
	// 	apiv1.GET("/tags", v1.GetTags)

	// }
	return r
}
