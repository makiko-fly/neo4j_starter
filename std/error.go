package std

type ErrWithCode struct {
	Code       int64
	Msg        string
	DisplayMsg string
}

func (self *ErrWithCode) Error() string {
	return self.DisplayMsg
}

var (
	SuccessCode    int64 = 20000
	DefaultErrCode int64 = 50000
)
