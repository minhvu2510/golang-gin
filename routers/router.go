package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhvu2510/golang-gin/middleware/jwt"
	"github.com/minhvu2510/golang-gin/routers/api"
	v1 "github.com/minhvu2510/golang-gin/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	fmt.Println("--------init router--------")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
	}
	// apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	// {
	// 	//get all tags
	// 	apiv1.GET("/tags", v1.GetTags)

	// }
	return r
}
