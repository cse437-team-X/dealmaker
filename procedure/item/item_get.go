package item

import (
	"github.com/dealmaker/procedure/item/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type GetItemInterface interface {
	GetGetItemDomain() *model.GetItemDomain
}

func (w *WorkerInstance) GetItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetGetItemDomain()
	filter := data.QueryFilter
	dbRes,err := w.FuncGetItem(c.Ctx, filter)

	if c.S.Action == "get" {
		for i:=0;i<len(dbRes);i++ {
			dbRes[i].Images = nil
		}
	}
	if c.S.Action == "detail" {
		for i:=0;i<len(dbRes);i++ {
			dbRes[i].Thumbnails = nil
		}
	}

	if err != nil {
		c.Errorw("get item", err)
		return http.StatusInternalServerError
	}
	data.Result = dbRes
	c.Debugw("vals", data.Result)
	return http.StatusOK
}