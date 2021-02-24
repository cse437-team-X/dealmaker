package model

type BaseRequest struct {
	BaseTime int64
	BaseLogId string
}

func (b *BaseRequest) GetBaseTime() int64 {
	return b.BaseTime
}
func (b *BaseRequest) GetBaseLogId() string {
	return b.BaseLogId
}
