package dal

import (
	"github.com/dealmaker/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"strconv"
)

func MigrateUserTable() {
	err := DB.AutoMigrate(&model.UserCredModel{})
	if err != nil {
		panic(err)
	}
}

func AddCredUser(user model2.CredUser) error {
	dbuser := &model.UserCredModel{
		CredUser: user,
	}
	res := DB.Create(dbuser)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}

func GetCredUser(loginName string) *model2.CredUser {
	dbRes := model.UserCredModel{}
	DB.Take(&dbRes, "login_name = ?", loginName)
	dbRes.SetUid(strconv.FormatUint(uint64(dbRes.ID),10))
	return &dbRes.CredUser
}
