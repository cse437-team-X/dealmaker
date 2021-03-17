package item_get

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"time"
)

type itemTagModel struct {
	ID uint
	Description string
	Title string
	Tag string
}

type ItemGetResult struct {
	Result []itemTagModel
}

// None nil conditions will be connected with ANDs
type ItemFilter struct {
	Uploader uint
	Tags []string
	BeginTime time.Time
	EndTime time.Time
	//FuzzyTitle string
}

type ItemGet struct {
	ItemFilter
	ItemGetResult
}
func (i *ItemGet) GetItemGet() *ItemGet {
	return i
}

type ItemGetInterface interface {
	GetItemGet() *ItemGet
}

func QueryItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(ItemGetInterface).GetItemGet()
	filter := data.ItemFilter
	query := dal.DB.Table(dal.TableItem).Select("description, title, tag, item_models.id").Joins("JOIN "+dal.TableTags + " a ON a.item_id = "+dal.TableItem+".id")
	if filter.Uploader > 0 {
		query = query.Where("uploader = ?", filter.Uploader)
	}
	if len(filter.Tags) > 0 {
		query = query.Where("a.tag IN (?)", filter.Tags)
	}
	if filter.BeginTime.Unix() > 0 {
		query = query.Where("updated_at >= ?", filter.BeginTime)
	}
	if filter.EndTime.Unix() > 0 {
		query = query.Where("updated_at <= ?", filter.EndTime)
	}

	query.Scan(&data.Result)

	c.Debugw("vals", data.Result)
	return 200
}