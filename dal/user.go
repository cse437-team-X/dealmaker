package dal

import (
	"github.com/dealmaker/shared/auth/model"
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
	dbRes := model.CredUser{}
	res := DB.Where(terms).First(&dbRes)
	err := res.Error
	if err != nil {
		return nil
	}

	return &dbRes
}

func UpdateUser(u *model.CredUser) error {
	res := DB.Save(u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
