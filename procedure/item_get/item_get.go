package item_get

import (
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/model"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type ItemGetResult struct {
	Result []model.Item
}

// None nil conditions will be connected with ANDs
type ItemFilter struct {
	Uploader uint
	Tags []string
	//BeginTime time.Time
	//EndTime time.Time
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
	mongoFilter := bson.M{}
	if filter.Uploader != 0 {
		mongoFilter["uploader"] = filter.Uploader
	}
	if filter.Tags != nil {
		mongoFilter["tags"] = bson.M{"$in":filter.Tags}
	}
	c.Infow("filter", mongoFilter)
	cursor, err := dal.ItemCollection.Find(c.Ctx, mongoFilter)
	if err != nil {
		c.Errorw("Read Item Collection", err)
		return http.StatusInternalServerError
	}

	var dbRes []model.Item
	if err = cursor.All(c.Ctx, &dbRes); err != nil {
		c.Errorw("Read Item Collection", err)
		return http.StatusInternalServerError
	}
	//
	//c.Debugw("res", dbRes)

	data.Result = dbRes
	c.Debugw("vals", data.Result)
	return 200
}