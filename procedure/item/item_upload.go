package item

import (
	"github.com/dealmaker/procedure/item/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
	"time"
)

type InsertItemInterface interface {
	GetItem() *model.Item
	GetJwtAuth() *model2.JwtAuth
}

func (w *WorkerInstance) InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(InsertItemInterface).GetItem()
	jwtData := c.DataDomain.(InsertItemInterface).GetJwtAuth()
	c.Debugw(
		"desc",data.Description,
		"title", data.Title,
		"urls", data.ImageUrls,
		"tags", data.Tags)

	data.Uploader = jwtData.Uid
	data.UpdateTime = time.Now().UnixNano() / 1000

	objid, err := w.FuncInsertItem(c.Ctx, data)
	if err != nil {
		c.Errorw("Insert item", err)
		return http.StatusForbidden
	}
	data.ObjId = objid
	return http.StatusOK
}