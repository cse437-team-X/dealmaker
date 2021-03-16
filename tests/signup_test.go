package tests

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/procedure/item_upload"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/stretchr/testify/require"
	"testing"
)


func InitForTest() {
	log.InitZapSugared(true, false, 1)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker?parseTime=true", nil, "mysql")
	auth_db.InitUserCredModel()
	item_upload.InitItemModel()
	item_upload.InitTagsModel()
	factory.BuildStreamlines()
}

func TestSignUp(t *testing.T) {
	InitForTest()

	dataDomain := handler.UserSignupDomain{
		UserCredModel: auth_db.UserCredModel{
			CredUser: model2.CredUser{
				LoginName:      "admin4",
				HashedPassword: "admin",
				Status:         1,
			},
		},
	}
	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          factory.Factory.Get("/auth/user/signup"),
		Ctx:        nil,
		Logger:     log.GlobalLogger,
		LogInfoGen: func(belt *streamline.ConveyorBelt) string {
			return belt.S.Name
		},
	}
	code, err := c.Run()
	require.Equal(t, 200, code)
	require.Nil(t, err)
}
