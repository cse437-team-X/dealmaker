package model

import (
	"github.com/dealmaker/model/obj"
	"github.com/dealmaker/shared/base"
)

type SearchDomain struct {
	base.Base
	TargetTags []string

	RespItemList []obj.Item
}
