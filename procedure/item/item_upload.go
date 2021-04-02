package item

import (
	"github.com/dealmaker/procedure/item/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type InsertItemInterface interface {
	GetItem() *model.Item
}

func (w *WorkerInstance) InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(InsertItemInterface).GetItem()
	jwtData := c.DataDomain.(model2.JwtInterface).GetJwtAuth()
	c.Debugw(
		"desc",data.Description,
		"title", data.Title,
		"urls", data.ImageUrls,
		"tags", data.Tags)

	data.Uploader = jwtData.Uid

	err := w.FuncInsertItem(c.Ctx, data)
	if err != nil {
		c.Errorw("Insert item", err)
		return http.StatusForbidden
	}
	return http.StatusOK
}