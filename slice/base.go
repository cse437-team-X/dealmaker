package slice

type BaseReqInterface interface {
	GetBaseTime() int64

	SetBaseTime(v int64)
}

type BaseRespInterface interface {
	GetBaseLogId() string

	SetBaseLogId(v string)
	SetBaseCode(v int)
}