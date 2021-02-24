package main

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/model"
	"github.com/gin-gonic/gin"
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
		resp := model.UserLoginResponse{}
		s.Run(&model.UserLoginRequest{
			Username:       "admin",
			HashedPassword: "admin",
		}, &resp)
		context.JSON(resp.BaseCode, resp)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
