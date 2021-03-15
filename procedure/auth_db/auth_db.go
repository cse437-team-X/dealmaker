package auth_db

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"net/http"
)

func InitModel()  {
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
	res := dal.DB.Take(dbRes, "login_name = ?", dbRes.LoginName)
	err := res.Error
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}