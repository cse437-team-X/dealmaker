package model

import (
	"github.com/dealmaker/model/obj"
	"gorm.io/gorm"
)

type ItemModel struct {
	obj.Item
	gorm.Model
}
