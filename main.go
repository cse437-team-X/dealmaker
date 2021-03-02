package main

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init begin
	log.InitZapSugared(true, true, 2)
	factory.BuildStreamlines()
	// Init end

	r := gin.Default()
	r.POST("/auth/user/login", handler.UserLogin)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
