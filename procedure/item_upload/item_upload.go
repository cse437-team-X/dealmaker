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
	
	dbItem := ItemModel{
		Description: data.Description,
		Title: data.Title,
		Uploader: jwtData.TokenClaim.Uid,
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