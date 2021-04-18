package item

import (
	"github.com/dealmaker/procedure/item/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ItemDeleteInterface interface {
	GetItemUpdate() *model.ItemUpdate
	GetJwtAuth() *model2.JwtAuth
}

func (w *WorkerInstance) ItemDelete(c *streamline.ConveyorBelt) int {
	objid := c.DataDomain.(ItemDeleteInterface).GetItemUpdate().ObjId
	jwt := c.DataDomain.(ItemDeleteInterface).GetJwtAuth()
	
	targetItem,err := w.FuncGetItem(c.Ctx, model.QueryFilter{
		ObjId: objid,
	})
	if err != nil {
		c.Errorw("pre-delete failed", err)
		return http.StatusInternalServerError
	}

	c.Debugw("titems", targetItem, "uid", jwt.Uid, "objid", objid)

	if jwt.Uid != targetItem[0].Uploader {
		c.Infow("not enough permission deleting", objid)
		return http.StatusForbidden
	}

	err = w.FuncDeleteItem(c.Ctx, objid)
	if err != nil {
		c.Errorw("delete failed", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}