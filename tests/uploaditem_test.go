package tests

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/dealmaker/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUploadItem(t *testing.T) {
	InitForTest()

	dataDomain := handler.UploadItemDomain{
		JwtAuth:   model2.JwtAuth{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOiI2IiwiUm9sZSI6InVzZXIiLCJpYXQiOjE2MTU4MjAxOTUsImV4cCI6MTYxNTgyMDc5NX0.4-8O2tu6UOB5LUHXRHXj06bMDZ40E_9kKSiChhmgNUc",
		},
		Item: model.Item{
				Description: "ruaruarua",
				Title:       "Test item1",
				ImageUrls:   []string{"a","b","c"},
				Tags:        []string{"aa","bb","cc"},
			},
		}

	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          factory.Factory.Get("/item/upload"),
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