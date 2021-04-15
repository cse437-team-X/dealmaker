package main

import (
	"github.com/dealmaker/conf"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/mentally-friendly-infra/log"
)

func init() {
	log.InitZapSugared(true, conf.Conf.EnvConfig.IsProd(), 1)
}
func main() {
	factory.BuildStreamlines()

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/auth/user/signup", handler.UserSignup)
	r.POST("/auth/user/login", handler.UserLogin)
	r.GET("/auth/user/recover", handler.UserRecover)
	r.GET("/auth/user/activate", handler.ActivateUser)
	r.POST("/auth/user/update", handler.UserUpdate)
	r.POST("/item/upload", handler.ItemUpload)
	r.POST("/item/get", handler.ItemGetHandler)
	r.POST("/item/user/contact", handler.UserContact)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
