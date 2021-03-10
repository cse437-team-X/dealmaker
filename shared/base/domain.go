package base

type Base struct {
	BaseTime int64
	BaseLogId string
	BaseSessionId string
}

func (b *Base) GetTime() int64 {
	return b.BaseTime
}
func (b *Base) SetTime(v int64) {
	b.BaseTime = v
}
func (b *Base) GetLogId() string {
	return b.BaseLogId
}
func (b *Base) SetLogId(v string) {
	b.BaseLogId = v
}
