package main

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init begin
	log.InitZapSugared(true, false, 2)
	factory.BuildStreamlines()
	// Init end

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("user_info", store))

	r.POST("/auth/user/login", handler.UserLogin)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
