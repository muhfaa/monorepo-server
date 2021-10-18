package common

type SuccessResponse string
type errorResponse string

const (
	Success                  SuccessResponse = "success"
	errMissingRequiredHeader errorResponse   = "missing_required_header"
	errInternalServerError                   = "internal_server_error"
)

type DefaultResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		string(Success),
		string(Success),
		data,
	}
}

//NewMissingHeaderResponse bad request caused by header missing
func NewMissingHeaderResponse(key string) DefaultResponse {
	return DefaultResponse{
		string(errMissingRequiredHeader),
		"Missing header " + key,
		map[string]interface{}{},
	}
}

//NewInternalServerError default internal server error with custom message
func NewInternalServerError(message string) DefaultResponse {
	return DefaultResponse{
		errInternalServerError,
		message,
		map[string]interface{}{},
	}
}
