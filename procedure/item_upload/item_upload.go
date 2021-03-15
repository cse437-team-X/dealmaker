package item_upload

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/model"
	"github.com/dealmaker/model/obj"
	"net/http"
)

type GetItemInterface interface {
	GetItem() *obj.Item
}

func InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetItem()
	c.Debugw(
		"desc",data.Description,
		"title", data.Title,
		"urls", data.ImageUrls,
		"tags", data.Tags)
	
	dbItem := model.ItemModel{
		Item: *data,
	}

	res := dal.DB.Create(&dbItem)
	err := res.Error
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}