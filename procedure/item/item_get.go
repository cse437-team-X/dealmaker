package item

import (
	"github.com/dealmaker/procedure/item/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

// None nil conditions will be connected with ANDs
type QueryFilter struct {
	Uploader uint
	Tags []string
	//BeginTime time.Time
	//EndTime time.Time
	//FuzzyTitle string
}

type GetItemDomain struct {
	QueryFilter
	Result []model.Item
}
func (i *GetItemDomain) GetGetItemDomain() *GetItemDomain {
	return i
}

type GetItemInterface interface {
	GetGetItemDomain() *GetItemDomain
}

func (w *WorkerInstance) GetItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(GetItemInterface).GetGetItemDomain()
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