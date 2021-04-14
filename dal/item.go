package dal

import (
	"context"
	"fmt"
	"github.com/dealmaker/procedure/item/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetItem(ctx context.Context, filter model.QueryFilter) ([]model.Item, error) {
	mongoFilter := bson.M{}
	if filter.Uploader != 0 {
		mongoFilter["uploader"] = filter.Uploader
	}
	if filter.Tags != nil {
		mongoFilter["tags"] = bson.M{"$in":filter.Tags}
	}
	if filter.FuzzyTitle != "" {
		mongoFilter["title"] = bson.M{"$regex":filter.FuzzyTitle, "$options":"i"}
	}
    timeRange := bson.M{}
    if filter.BeginTime != 0 {
    	timeRange["$gte"] = filter.BeginTime
	}
	if filter.EndTime != 0 {
		timeRange["$lt"] = filter.EndTime
	}
	if len(timeRange) > 0 {
		mongoFilter["updatetime"] = timeRange
	}
	fmt.Println(mongoFilter)
	fmt.Println(filter)
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

func InsertItem(ctx context.Context, data *model.Item) (string, error) {
	res, err := ItemCollection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).String(), nil
}