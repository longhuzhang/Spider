package error_operate

type NoticeError struct {
	Msg string
}

func (error NoticeError)Error() string{
	return error.Msg
}

func (error NoticeError)GetCode() ErrorCode{

	return CommonError
}