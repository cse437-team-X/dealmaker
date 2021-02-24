package main

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/model"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	log.InitZapSugared(true, false)
	factory.BuildStreamlines()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("test", func(context *gin.Context) {
		s := factory.Factory.Get("/auth/user/login")
		s.Run(&model.UserLoginRequest{
			BaseRequest:    model.BaseRequest{
				BaseTime:  time.Now().Unix(),
				BaseLogId: "LOGID2ijdisdi",
			},
			Username:       "test1",
			HashedPassword: "asdq23ewidqcia",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
