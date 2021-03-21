package main

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init begin
	log.InitZapSugared(true, false, 1)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker?parseTime=true", nil, "mysql")
	dal.InitMongoDB()
	auth_db.InitUserCredModel()
	//item_upload.InitItemModel()
	//item_upload.InitTagsModel()

	factory.BuildStreamlines()
	// Init end

	r := gin.Default()

	//store := memstore.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("user_info", store))

	r.POST("/auth/user/signup", handler.UserSignup)
	r.POST("/auth/user/login", handler.UserLogin)
	r.POST("/item/upload", handler.ItemUpload)
	r.POST("/item/get", handler.ItemGetHandler)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
