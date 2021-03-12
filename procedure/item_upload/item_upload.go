package item_upload

import (
	"gitee.com/fat_marmota/streamline"
	"net/http"
)

type GetItemInterface interface {
	GetItemDescription() string
	GetItemTitle() string
	GetItemImageUrls() []string
	GetItemTags() []string
}

func InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface)
	c.Debugw(
		"desc",data.GetItemDescription(),
		"title", data.GetItemTitle(),
		"urls", data.GetItemImageUrls(),
		"tags", data.GetItemTags())
	return http.StatusOK
}