package json

//Response struct is to make a valid response
type Response struct {
	Code    int
	Success bool
	Message string
	Data    interface{}
}

//NewSuccessResponse makes a new Response object with success state
func NewSuccessResponse(statusCode int, data interface{}) *Response {
	return &Response{
		Success: true,
		Code:    statusCode,
		Data:    data,
	}
}

//NewErrorResponse makes a new Response object with fail state
func NewErrorResponse(statusCode int, message string) *Response {
	return &Response{
		Code: statusCode,
		Data: message,
	}
}
