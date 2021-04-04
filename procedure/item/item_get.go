package item

import (
	"github.com/dealmaker/procedure/item/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type GetItemInterface interface {
	GetItemDomain() *model.GetItemDomain
}

func (w *WorkerInstance) GetItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetItemDomain()
	filter := data.QueryFilter
	dbRes,err := w.FuncGetItem(c.Ctx, filter)
	if err != nil {
		c.Errorw("get item", err)
		return http.StatusInternalServerError
	}
	data.Result = dbRes
	c.Debugw("vals", data.Result)
	return http.StatusOK
}