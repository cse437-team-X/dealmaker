package factory

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/model"
	"github.com/dealmaker/model/obj"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUploadItem(t *testing.T) {
	InitForTest()

	type UploadItemDomain struct {
		base.Base
		model2.JwtAuth
		model.ItemModel
	}

	dataDomain := UploadItemDomain{
		JwtAuth:   model2.JwtAuth{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOiI2IiwiUm9sZSI6InVzZXIiLCJpYXQiOjE2MTU4MjAxOTUsImV4cCI6MTYxNTgyMDc5NX0.4-8O2tu6UOB5LUHXRHXj06bMDZ40E_9kKSiChhmgNUc",
		},
		ItemModel: model.ItemModel{
			Item: obj.Item{
				Description: "ruaruarua",
				Title:       "Test item1",
				ImageUrls:   "a,b,c",
				Tags:        "aa,bb,cc",
			},
		},
	}

	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          Factory.Get("/item/upload"),
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