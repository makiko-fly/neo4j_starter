package std

// short cut function to create specific error
func NewNeo4jQueryErr(msg string) *Err {
	return &Err{
		Code: Neo4jQueryErrCode,
		Msg:  msg,
	}
}

type Err struct {
	Code int64
	Msg  string
}

func (self *Err) Error() string {
	return self.Msg
}

var (
	SuccessCode       int64 = 20000
	DefaultErrCode    int64 = 50000
	Neo4jQueryErrCode int64 = 50001
)
