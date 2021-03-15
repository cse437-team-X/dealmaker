package factory

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogin(t *testing.T) {
	InitForTest()

	dataDomain := struct {
		base.Base
		model.JwtAuth
		auth_db.UserCredModel
	}{}

	dataDomain.LoginName = "admin"
	dataDomain.HashedPassword = "admin"

	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          Factory.Get("/auth/user/login"),
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