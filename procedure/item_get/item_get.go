package item_get

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"time"
)

// None nil conditions will be connected with ANDs
type ItemFilter struct {
	Uploader uint
	Tags []string
	BeginTime time.Time
	EndTime time.Time
	//FuzzyTitle string
}
func (i *ItemFilter) GetItemFilter() *ItemFilter {
	return i
}

type ItemFilterInterface interface {
	GetItemFilter() *ItemFilter
}

func QueryItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(ItemFilterInterface).GetItemFilter()
	query := dal.DB.Table(dal.TableItem).Select("description, title, tag, item_models.id").Joins("JOIN "+dal.TableTags + " a ON a.item_id = "+dal.TableItem+".id")
	if data.Uploader > 0 {
		query = query.Where("uploader = ?", data.Uploader)
	}
	if len(data.Tags) > 0 {
		query = query.Where("a.tag IN (?)", data.Tags)
	}
	if data.BeginTime.Unix() > 0 {
		query = query.Where("updated_at >= ?", data.BeginTime)
	}
	if data.EndTime.Unix() > 0 {
		query = query.Where("updated_at <= ?", data.EndTime)
	}

	type ItemTagModel struct {
		ID uint
		Description string
		Title string
		Tag string
	}

	var res []ItemTagModel
	query.Scan(&res)

	//for rows.Next() {
	//	cols, _:=rows.Columns()
	//	res := ItemTagModel{}
	//	dal.DB.ScanRows(rows, &res)
	//	//rows.Scan(&res)
	//	c.Debugw("cols", cols, "vals", res)
	//}
	c.Debugw("vals", res)
	return 200
}