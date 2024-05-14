package e

type ErrorWithCode struct {
	Err  error
	Code int
}

func NewE(e error, code int) ErrorWithCode {
	return ErrorWithCode{
		Err:  e,
		Code: code,
	}
}
