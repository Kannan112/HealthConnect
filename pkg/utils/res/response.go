package res

type Response struct {
	StatusCode int         `json:"stastus_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}

func ErrorResponse(statusCode int, message string, err interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Errors:     err,
	}
}

func SuccessResponse(statsCode int, message string, data interface{}) Response {
	return Response{
		StatusCode: statsCode,
		Message:    message,
		Data:       data,
	}
}
