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
	log.InitZapSugared(true, false, 1)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
}

func TestFullAuth(t *testing.T) {
	InitForTest()

	type LoginDomain struct {
		model2.JwtAuth
		model2.CredUser
	}

	dataDomain := LoginDomain{
		CredUser: model2.CredUser{
			LoginName: "jiayi.zhsdadsaadsnssg2@wustl.edu",
			HashedPassword: "fakssaepdw!",
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

	r := ValidateCredUser(&c)
	require.Equal(t, r, http.StatusForbidden)
	r = SignUp(&c)
	require.Equal(t, r, http.StatusOK)

	r = SignUp(&c)
	require.Equal(t, r, http.StatusForbidden)

	r = ValidateCredUser(&c)
	require.Equal(t, r, http.StatusOK)
	r = SignToken(&c)
	require.NotNil(t, dataDomain.Token)
	require.Equal(t, r,http.StatusOK)

	r = Validate(&c)
	require.Equal(t, r,http.StatusOK)

	require.Equal(t, dataDomain.Role, "user")

	r = Logout(&c)
	require.Equal(t, r,http.StatusOK)

	r = Validate(&c)
	require.Equal(t, r,http.StatusForbidden)
}