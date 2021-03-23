package auth_db

import (
	"github.com/dealmaker/dal"
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"gorm.io/gorm"
	"net/http"
)

func InitUserCredModel()  {
	err := dal.DB.AutoMigrate(&UserCredModel{})
	if err != nil {
		panic(err)
	}
}

type AuthDBInterface interface {
	GetUserCredModel() *UserCredModel
	//SetUserCredModel(*UserCredModel)
}

func InsertUser(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(AuthDBInterface).GetUserCredModel()
	data.Status = 0 // Inactive
	data.Role = "user"
	res := dal.DB.Create(data)
	err := res.Error
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func GetUser(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(AuthDBInterface)
	dbRes := data.GetUserCredModel()
	oldpw := dbRes.HashedPassword
	c.Debugw("getuser", dbRes)
	res := dal.DB.Take(dbRes, "login_name = ?", dbRes.LoginName)
	err := res.Error
	if err != nil {
		return http.StatusInternalServerError
	}

	if oldpw != "" && oldpw != dbRes.HashedPassword {
		return http.StatusForbidden
	}
	return http.StatusOK
}

func ActiveUser(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(model.JwtInterface).GetJwtAuth()
	user := UserCredModel{
		Model:gorm.Model{
			ID: data.TokenClaim.Uid,
		},
	}
	res := dal.DB.Model(&user).Update("status", 1)// 1 means active
	if res.Error != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}