package dal

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/model/obj"
	"gorm.io/gorm"
)

type UserDBModel struct {
	obj.UserInfo
	gorm.Model
}

func MigrateUserTable() {
	err := DB.AutoMigrate(&UserDBModel{})
	if err != nil {
		panic(err)
	}
}

func AddUser(user obj.UserInfo) error {
	dbuser := &UserDBModel{
		UserInfo: user,
	}
	res := DB.Create(dbuser)
	err := res.Error
	if err != nil {
		log.Errorw("Add user", "err", err.Error())
		return err
	}
	return nil
}

func GetUser(email string) obj.UserInfo {
	res := UserDBModel{}
	DB.Take(&res, "email = ?", email)
	return res.UserInfo
}
