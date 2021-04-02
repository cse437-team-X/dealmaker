package dal

import (
	"context"
	"github.com/dealmaker/procedure/item"
	"github.com/dealmaker/procedure/item/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetItem(ctx context.Context, filter item.QueryFilter) ([]model.Item, error) {
	mongoFilter := bson.M{}
	if filter.Uploader != 0 {
		mongoFilter["uploader"] = filter.Uploader
	}
	if filter.Tags != nil {
		mongoFilter["tags"] = bson.M{"$in":filter.Tags}
	}

	cursor, err := ItemCollection.Find(ctx, mongoFilter)
	if err != nil {
		return nil, err
	}

	var dbRes []model.Item
	if err = cursor.All(ctx, &dbRes); err != nil {
		return nil, err
	}
	return dbRes, nil
}

func InsertItem(ctx context.Context, data *model.Item) error {
	_, err := ItemCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}