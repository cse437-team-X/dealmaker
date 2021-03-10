package slice

import (
	"context"
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/base_model"
	"github.com/dealmaker/base_model/obj"
	"github.com/dealmaker/dal"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func InitForTest() {
	log.InitZapSugared(true, false, 2)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
}

func TestAuth(t *testing.T) {
	InitForTest()

	dataDomain := base_model.JwtAuth{}
	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          nil,
		Ctx:        context.Background(),
		Logger:     log.GlobalLogger,
	}

	r := SignToken(&c)
	require.NotNil(t, dataDomain.Token)
	require.Equal(t, r,http.StatusOK)

	r = Validate(&c)
	require.Equal(t, r,http.StatusOK)

	r = Logout(&c)
	require.Equal(t, r,http.StatusOK)

	r = Validate(&c)
	require.Equal(t, r,http.StatusForbidden)
}

func TestValidateUsernamePassword(t *testing.T) {
	log.InitZapSugared(true, false, 2)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
	dataDomain := base_model.UserInfoDomain{}
	dataDomain.Email = "jiayi.zhang@wustl.edu"
	dataDomain.HashedPassword = "fakepw"
	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          nil,
		Ctx:        context.Background(),
		Logger:     log.GlobalLogger,
	}
	r := ValidateUsernamePassword(&c)
	require.Equal(t, r,http.StatusOK)
}

func TestFullAuth(t *testing.T) {
	InitForTest()

	type LoginDomain struct {
		base_model.JwtAuth
		obj.UserInfo
	}

	dataDomain := LoginDomain{
		UserInfo: obj.UserInfo{
			Username:       "admin",
			Email:          "jiayi.zhang@wustl.edu",
			HashedPassword: "fakepw",
			Status:         1,
		},
	}
	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          nil,
		Ctx:        context.Background(),
		Logger:     log.GlobalLogger,
	}

	r := ValidateUsernamePassword(&c)
	require.Equal(t, r, http.StatusForbidden)
	r = SignUp(&c)
	require.Equal(t, r, http.StatusOK)
	r = ValidateUsernamePassword(&c)
	require.Equal(t, r, http.StatusOK)
	r = SignToken(&c)
	require.NotNil(t, dataDomain.Token)
	require.Equal(t, r,http.StatusOK)

	r = Validate(&c)
	require.Equal(t, r,http.StatusOK)

	r = Logout(&c)
	require.Equal(t, r,http.StatusOK)

	r = Validate(&c)
	require.Equal(t, r,http.StatusForbidden)
}