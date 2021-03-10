package dal

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/model/obj"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func InitForTest() {
	log.InitZapSugared(true, false, 2)
	InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
}

func TestAddUser(t *testing.T) {
	InitForTest()
	uniqueStr := time.Now().String()

	r := AddUser(obj.UserInfo{
		Username:       uniqueStr,
		Email:          uniqueStr,
		HashedPassword: uniqueStr,
		Status:         0,
	})
	require.Nil(t, r)
	r = AddUser(obj.UserInfo{
		Username:       uniqueStr,
		Email:          uniqueStr,
		HashedPassword: uniqueStr,
		Status:         0,
	})
	require.NotNil(t, r)
}
