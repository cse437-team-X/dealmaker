package model

type Item struct {
	Description string
	Title string
	ImageUrls []string
	Tags []string
}

func (i *Item) GetItem() *Item {
	return i
}
//
//func (i *Item) GetItemDescription() string {
//	return i.Description
//}
//func (i *Item) GetItemTitle() string {
//	return i.Title
//}
//func (i *Item) GetItemImageUrls() []string {
//	return i.ImageUrls
//}
//func (i *Item) GetItemTags() []string {
//	return i.Tags
//}