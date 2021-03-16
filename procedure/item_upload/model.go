package item_upload

import (
	"github.com/dealmaker/dal"
	"gorm.io/gorm"
)

type TagsModel struct {
	gorm.Model
	ItemId uint `gorm:"index"`
	Tag string
}

func InitTagsModel() {
	err := dal.DB.AutoMigrate(&TagsModel{})
	if err != nil {
		panic(err)
	}
}

type ItemModel struct {
	Description string
	Title string
	gorm.Model
}

func InitItemModel() {
	dal.DB.AutoMigrate(&ItemModel{})
}
