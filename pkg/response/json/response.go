package json

//Response struct
type Response struct {
	Code    int
	Success bool
	Message string
	Data    interface{}
}

func NewSuccessResponse(statusCode int, data interface{}) *Response {
	return &Response{
		Success: true,
		Code:    statusCode,
		Data:    data,
	}
}

func NewErrorResponse(statusCode int, message string) *Response {
	return &Response{
		Code: statusCode,
		Data: message,
	}
}
