package item_upload

import (
	"github.com/dealmaker/dal"
	"github.com/dealmaker/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type GetItemInterface interface {
	GetItem() *model.Item
}

func InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetItem()
	jwtData := c.DataDomain.(model2.JwtInterface).GetJwtAuth()
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