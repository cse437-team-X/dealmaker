package dal

import (
	"context"
	"errors"
	"fmt"
	"github.com/dealmaker/procedure/item/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetItem(ctx context.Context, filter model.QueryFilter) ([]model.Item, error) {
	mongoFilter := buildMongoFilter(filter)
	if mongoFilter == nil {
		return nil, errors.New("invalid obj_id")
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

func InsertItem(ctx context.Context, data *model.Item) (string, error) {
	res, err := ItemCollection.InsertOne(ctx, bson.D{
		{"description", data.Description},
		{"title", data.Title},
		{"tags", data.Tags},
		{"uploader", data.Uploader},
		{"images", data.Images},
		{"thumbnails", data.Thumbnails},
		{"isdeleted", data.IsDeleted},
		{"newprice", data.NewPrice},
	})
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func buildMongoFilter(filter model.QueryFilter) bson.M {
	mongoFilter := bson.M{}
	mongoFilter["isdeleted"] = 0

	if filter.ObjId != "" {

		id, err:=primitive.ObjectIDFromHex(filter.ObjId)
		if err != nil {
			return nil
		}
		mongoFilter["_id"] = id
		return mongoFilter
	}

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

	priceRange := bson.M{}
	if filter.PriceLow != 0 {
		priceRange["$gte"] = filter.PriceLow
	}
	if filter.PriceHigh != 0 {
		priceRange["$lt"] = filter.PriceHigh
	}
	if len(priceRange) > 0 {
		//fmt.Println(mongoFilter)
		mongoFilter["newprice"] = priceRange
	}
	//fmt.Println(filter.PriceLow)
	//fmt.Println(len(priceRange))
	fmt.Println(mongoFilter)

	return mongoFilter
}

func DeleteItem(ctx context.Context, objId string) error {
	update := bson.M{
		"$set":bson.M{
			"isdeleted":1,
		},
	}

	id, err:=primitive.ObjectIDFromHex(objId)
	if err != nil {
		return err
	}
	_, err = ItemCollection.UpdateByID(ctx, id, update)
	return err
}