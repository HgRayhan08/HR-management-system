package dto

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func SucsessResponse(code int, message string, data any) response[any] {
	return response[any]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func FailedResponse(code int, message string, data any) response[any] {
	return response[any]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
