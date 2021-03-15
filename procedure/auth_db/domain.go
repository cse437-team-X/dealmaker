package auth_db

import (
	"github.com/dealmaker/shared/auth/model"
	"gorm.io/gorm"
)

type UserCredModel struct {
	model.CredUser
	gorm.Model
}
func (u *UserCredModel) GetUserCredModel() *UserCredModel {
	return u
}

//func (u *UserCredModel) GetUid() string {
//	return strconv.FormatUint(uint64(u.Model.ID),10)
//}

