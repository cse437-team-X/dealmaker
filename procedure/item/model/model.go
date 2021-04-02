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
