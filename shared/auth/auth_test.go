package auth

import (
	"context"
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	model2 "github.com/dealmaker/shared/auth/model"
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

	dataDomain := model2.JwtAuth{}
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

//func TestValidateUsernamePassword(t *testing.T) {
//	log.InitZapSugared(true, false, 2)
//	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
//	dataDomain := model.UserInfoDomain{}
//	dataDomain.Email = "jiayi.zhang@wustl.edu"
//	dataDomain.HashedPassword = "fakepw"
//	c := streamline.ConveyorBelt{
//		DataDomain: &dataDomain,
//		S:          nil,
//		Ctx:        context.Background(),
//		Logger:     log.GlobalLogger,
//	}
//	r := ValidateUsernamePassword(&c)
//	require.Equal(t, r,http.StatusOK)
//}

func TestFullAuth(t *testing.T) {
	InitForTest()

	type LoginDomain struct {
		model2.JwtAuth
		model2.CredUser
	}

	dataDomain := LoginDomain{
		CredUser: model2.CredUser{
			LoginName: "jiayi.zhang2@wustl.edu",
			HashedPassword: "fakepw!",
			Status:         1,
		},
	}
	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          &streamline.Streamline{Name: "TestStreamline"},
		Ctx:        context.Background(),
		Logger:     log.GlobalLogger,
		LogInfoGen: func(belt *streamline.ConveyorBelt) string {
			return belt.S.Name
		},
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