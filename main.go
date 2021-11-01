package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/minhvu2510/golang-gin/utils"
	"github.com/minhvu2510/golang-gin/pkg/setting"
)

func init() {
	fmt.Println("----int setting server app----")
	setting.Setup()
	//models.Setup()
	//logging.Setup()
	//gredis.Setup()
	//util.Setup()
}
func main() {
	fmt.Println("----int main server app----")
	//utils.Notify()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	fmt.Println(port)
	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//http.ListenAndServe(":8080", router)

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}