package model

type Item struct {
	ObjId string

	Description string
	Title string
	ImageUrls []string
	Tags []string
	Uploader uint
	UpdateTime int64
}

func (i *Item) GetItem() *Item {
	return i
}

// None nil conditions will be connected with ANDs
type QueryFilter struct {
	Uploader uint
	Tags []string
	BeginTime int64
	EndTime int64
	FuzzyTitle string
}

type GetItemDomain struct {
	QueryFilter
	Result []Item
}
func (i *GetItemDomain) GetGetItemDomain() *GetItemDomain {
	return i
}