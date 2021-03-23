package auth_db

import (
	"github.com/dealmaker/dal"
	"github.com/itzmeerkat/streamline"
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
	data := c.DataDomain.(AuthDBInterface)
	res := dal.DB.Create(data.GetUserCredModel())
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