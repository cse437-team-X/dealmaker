package model

type BaseRequest struct {
	BaseTime int64
}

func (b *BaseRequest) GetBaseTime() int64 {
	return b.BaseTime
}
func (b *BaseRequest) SetBaseTime(v int64) {
	b.BaseTime = v
}


type BaseResponse struct {
	BaseCode int
	BaseLogId string
}

func (b *BaseResponse) GetBaseLogId() string {
	return b.BaseLogId
}
func (b *BaseResponse) SetBaseLogId(v string) {
	b.BaseLogId = v
}
func (b *BaseResponse) SetBaseCode(v int) {
	b.BaseCode = v
}