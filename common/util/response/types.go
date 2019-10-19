package response

type ReturnCode int

const (
	SuccessCode             ReturnCode = 200000
	FailCode                ReturnCode = 200001
	NotFoundCode            ReturnCode = 404000
	InternalServerErrorCode ReturnCode = 500000
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Err  interface{} `json:"error"`
}

func New(code ReturnCode, msg string, data interface{}, error interface{}) Response {
	return Response{int(code), msg, data, error}
}

func Ok(data interface{}) Response {
	return New(SuccessCode, "Ok", data, nil)
}

func Fail(msg string, error interface{}) Response {
	return New(FailCode, msg, nil, error)
}

func NotFound(err interface{}) Response {
	return New(NotFoundCode, "Not Found", nil, err)

}

func InternalServerError(err interface{}) Response {
	return New(InternalServerErrorCode, "Internal Server Error", nil, err)
}

func Error(code ReturnCode, msg string, error interface{}) Response {
	return New(code, msg, nil, error)
}
