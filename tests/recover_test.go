package tests

import (
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/dealmaker/shared/auth"
	"github.com/itzmeerkat/mentally-friendly-infra/log"
	"github.com/itzmeerkat/streamline"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRecover(t *testing.T) {
	InitForTest()

	dataDomain := handler.UserLoginDomain{}

	dataDomain.LoginName = "admin"
	//dataDomain.HashedPassword = "admin"

	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          factory.Factory.Get("/auth/user/recover"),
		Ctx:        nil,
		Logger:     log.GlobalLogger,
		LogInfoGen: func(belt *streamline.ConveyorBelt) string {
			return belt.S.Name
		},
	}

	code, err := c.Run()
	require.Equal(t, 200, code)
	require.Nil(t, err)
	auth.Validate(&c)
}
