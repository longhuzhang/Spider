package error_operate

type ErrorCode int

const (
	NoError        ErrorCode = iota
	CommonError              = 1
	ReloginAccount           = 1001
	ReloginRole              = 1002
)
