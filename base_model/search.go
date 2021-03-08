package base_model

import "github.com/dealmaker/base_model/obj"

type SearchDomain struct {
	BaseDomain
	TargetTags []string

	RespItemList []obj.Item
}
