package item

import (
	"github.com/dealmaker/procedure/item/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ItemDeleteInterface interface {
	GetItemUpdate() *model.ItemUpdate
}

func (w *WorkerInstance) ItemDelete(c *streamline.ConveyorBelt) int {
	objid := c.DataDomain.(ItemDeleteInterface).GetItemUpdate().ObjId

	err := w.FuncDeleteItem(c.Ctx, objid)
	if err != nil {
		c.Errorw("delete failed", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}