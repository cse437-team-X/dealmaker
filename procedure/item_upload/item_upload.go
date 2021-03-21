package item_upload

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/model"
	"github.com/dealmaker/shared/auth"
	"net/http"
)

type GetItemInterface interface {
	GetItem() *model.Item
}

func InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetItem()
	jwtData := c.DataDomain.(auth.JwtInterface).GetJwtAuth()
	c.Debugw(
		"desc",data.Description,
		"title", data.Title,
		"urls", data.ImageUrls,
		"tags", data.Tags)

	data.Uploader = jwtData.Uid

	_, err := dal.ItemCollection.InsertOne(c.Ctx, data)
	if err != nil {
		c.Errorw("Insert item", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}