package dal

import (
	"github.com/dealmaker/conf"
	"github.com/dealmaker/shared/auth/model"
)

func init() {
	InitDatabaseClient(conf.Conf.MySqlConfig.GetMasterDSN(), nil, "mysql")
	InitMongoDB()
	err := DB.AutoMigrate(&model.CredUser{})
	if err != nil {
		panic(err)
	}
}
