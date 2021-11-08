package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhvu2510/golang-gin/pkg/gredis"
	"github.com/minhvu2510/golang-gin/pkg/util"

	//"github.com/minhvu2510/golang-gin/utils" // notify via telegram

	"github.com/minhvu2510/golang-gin/models"
	"github.com/minhvu2510/golang-gin/pkg/logging"
	"github.com/minhvu2510/golang-gin/pkg/setting"
)

func init() {
	fmt.Println("----int setting server app----")
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}
func main() {
	fmt.Println("----int main server app----")
	//utils.Notify()
	// router test
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// setup router mới

	fmt.Println("+++++", setting.ServerSetting)
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
	// log.Printf("[info] start http server listening %s", endPoint)
	// logging.Info("loi khong xac dinh")
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
