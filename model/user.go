package model

import (
	"github.com/dealmaker/shared/auth/model"
	"gorm.io/gorm"
)

type UserCredModel struct {
	model.CredUser
	gorm.Model
}

//type UserInfo struct {
//	Email string
//	Username string
//	ContactInfo string
//}