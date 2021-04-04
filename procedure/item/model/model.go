package model

type Item struct {
	Description string
	Title string
	ImageUrls []string
	Tags []string
	Uploader uint
}

func (i *Item) GetItem() *Item {
	return i
}

// None nil conditions will be connected with ANDs
type QueryFilter struct {
	Uploader uint
	Tags []string
	//BeginTime time.Time
	//EndTime time.Time
	//FuzzyTitle string
}

type GetItemDomain struct {
	QueryFilter
	Result []Item
}
func (i *GetItemDomain) GetItemDomain() *GetItemDomain {
	return i
}