package errors

type ErrCode int

const (
	ERR_VALIDATOR_NOT_FOUND ErrCode = 10001
)

type Error struct {
	Code int
	Msg  interface{}
}
