package dal

import (
	"github.com/dealmaker/base_model/obj"
	"gorm.io/gorm"
)

type UserDBModel struct {
	obj.User
	gorm.Model
}

func MigrateUserTable() {
	err := DB.AutoMigrate(&UserDBModel{})
	if err != nil {
		panic(err)
	}
}

func AddUser(user obj.User) {
	dbuser := &UserDBModel{
		User:  user,
	}
	DB.Create(dbuser)
}

func GetUser(email string) obj.User {
	res := UserDBModel{}
	DB.Take(&res, "email = ?", email)
	return res.User
}
