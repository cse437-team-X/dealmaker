package model

type SearchDomain struct {
	BaseDomain
	TargetTags []string

	RespItemList []Item
}
