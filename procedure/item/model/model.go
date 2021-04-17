package model

type Item struct {
	ObjId string `bson:"_id"`

	Description string
	Title string
	Images []string
	Thumbnails []string
	Tags []string

	OriginalPrice float32
	NewPrice float32

	Uploader uint
	UpdateTime int64
	IsDeleted int
}

func (i *Item) GetItem() *Item {
	return i
}

type GetItemDomain struct {
	QueryFilter
	Result []Item
}
func (i *GetItemDomain) GetGetItemDomain() *GetItemDomain {
	return i
}

// None nil conditions will be connected with ANDs
type QueryFilter struct {
	ObjId string

	Uploader uint
	Tags []string
	BeginTime int64
	EndTime int64
	FuzzyTitle string

	PriceLow float64
	PriceHigh float64

	// 0: full size, 1: thumbnails, 2: no image
	//ImageType int
}
func (q *QueryFilter) GetQueryFilter() *QueryFilter {
	return q
}


type ItemUpdate struct {
	ObjId string
}
func (i *ItemUpdate) GetItemUpdate() *ItemUpdate {
	return i
}