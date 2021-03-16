package item_upload

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/model"
	"net/http"
)

type GetItemInterface interface {
	GetItem() *model.Item
}

func InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetItem()
	c.Debugw(
		"desc",data.Description,
		"title", data.Title,
		"urls", data.ImageUrls,
		"tags", data.Tags)
	
	dbItem := ItemModel{
		Description: data.Description,
		Title: data.Title,
	}

	err := dal.DB.Create(&dbItem).Error
	if err != nil {
		return http.StatusInternalServerError
	}

	var tagsModel []TagsModel
	for _, v :=range data.Tags {
		tagsModel = append(tagsModel, TagsModel{
			ItemId: dbItem.ID,
			Tag:    v,
		})
	}
	err = dal.DB.Create(&tagsModel).Error
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}