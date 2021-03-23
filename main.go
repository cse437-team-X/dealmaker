package main

import (
	"fmt"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/procedure/email"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/mentally-friendly-infra/config"
	"github.com/itzmeerkat/mentally-friendly-infra/log"
)

type MySqlConfig struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
	Address string `yaml:"Address"`
}
func (m *MySqlConfig) GetMasterDSN() string {
	return fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", m.Username, m.Password, m.Address, m.Database)
}

type EnvConfig struct {
	DEALMAKER_ENV string
}
func (e *EnvConfig) IsProd() bool {
	return e.DEALMAKER_ENV == "prod"
}

type SendGridConfig struct {
	SENDGRID_API_KEY string
}

type AllConfig struct {
	MySqlConfig
	SendGridConfig
	EnvConfig
}

var conf AllConfig
func LoadConfigs() {
	config.LoadEnvVar(&conf.EnvConfig)
	if conf.EnvConfig.DEALMAKER_ENV == "" {
		panic("Set environment before launch")
	}
	confPath := "./conf/" + conf.EnvConfig.DEALMAKER_ENV +".yml"
	config.LoadConfigFile(confPath, &conf.MySqlConfig)
	config.LoadEnvVar(&conf.SendGridConfig)
	fmt.Println(conf)
}

func main() {
	LoadConfigs()
	// Init begin
	log.InitZapSugared(true, conf.EnvConfig.IsProd(), 1)
	dal.InitDatabaseClient(conf.MySqlConfig.GetMasterDSN(), nil, "mysql")
	dal.InitMongoDB()
	email.InitEmailClient()
	auth_db.InitUserCredModel()
	email.InitEmailClient()
	//item_upload.InitItemModel()
	//item_upload.InitTagsModel()

	factory.BuildStreamlines()
	// Init end

	r := gin.Default()

	//store := memstore.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("user_info", store))
	r.Use(cors.Default())

	r.POST("/auth/user/signup", handler.UserSignup)
	r.POST("/auth/user/login", handler.UserLogin)
	r.GET("/auth/user/recover", handler.UserRecover)
	r.POST("/item/upload", handler.ItemUpload)
	r.POST("/item/get", handler.ItemGetHandler)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
