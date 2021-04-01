package dal

import "github.com/dealmaker/shared/auth/model"

func InitCredUser() {
	err := DB.AutoMigrate(&model.CredUser{})
	if err != nil {
		panic(err)
	}
}
