package base

type Base struct {
	BaseTime int64
	BaseLogId string
	BaseMessage string
}

func (b *Base) GetBase() *Base {
	return b
}