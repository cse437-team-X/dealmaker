package factory

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/model"
	"github.com/dealmaker/procedure/auth_db"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/stretchr/testify/require"
	"testing"
)


func InitForTest() {
	log.InitZapSugared(true, false, 1)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker?parseTime=true", nil, "mysql")
	auth_db.InitModel()
	model.InitItemModel()
	BuildStreamlines()
}

func TestSignUp(t *testing.T) {
	InitForTest()

	type SignUpDomain struct {
		base.Base
		auth_db.UserCredModel
	}

	dataDomain := SignUpDomain{
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
		S:          Factory.Get("/auth/user/signup"),
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
