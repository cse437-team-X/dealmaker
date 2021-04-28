package dal

import (
	"github.com/dealmaker/shared/auth/model"
	"gorm.io/gorm"
)

func InsertUser(u *model.CredUser) error {
	res := DB.Create(u)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(terms *model.CredUser) *model.CredUser {
	trimTerm := model.CredUser{LoginName: terms.LoginName}
	dbRes := model.CredUser{}
	res := DB.Where(&trimTerm).First(&dbRes)
	err := res.Error
	if err != nil {
		return nil
	}

	return &dbRes
}

func UpdateUser(u *model.CredUser) error {
	res := DB.Model(&model.CredUser{
		Model: gorm.Model{
			ID: u.ID,
		},
	}).Updates(u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
