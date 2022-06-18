package code

import "time"

const (
	Successful          = 200
	Sync                = 202
	BadRequest          = 400
	JWTRejected         = 401
	PermissionDenied    = 403
	DoesNotExist        = 404
	FormatError         = 415
	InternalServerError = 500
	ServerDown          = 503
)

var (
	message = map[int]string{
		200: "Successful http requests.",
		202: "Sync function call success.",
		400: "Bad Request",
		401: "JWT rejected.",
		403: "Permission denied.",
		404: "Item does not exist.",
		415: "Data format error.",
		500: "Unexpected server error.",
		503: "Server down.",
	}
)

type codeTime struct {
	//回傳代碼
	Code int `json:"code"`
	//錯誤時間
	Timestamp string `json:"timestamp" example:"2021-07-29T07:23:47Z"`
}

type SuccessfulMessage struct {
	codeTime
	//正確回傳內容
	Body interface{} `json:"body"`
}

type ErrorMessage struct {
	codeTime
	//錯誤回傳訊息
	Message string `json:"message"`
	//詳細錯誤內容
	Detailed interface{} `json:"detailed"`
}

func GetCodeMessage(args ...interface{}) interface{} {
	var codeMessage interface{}

	switch args[0] {
	case 200, 202:
		codeMessage = setSuccessfulMessage(args[0].(int), args[1])
	default:
		codeMessage = setErrorMessage(args[0].(int), args[1])
	}

	return codeMessage
}

func setSuccessfulMessage(code int, body interface{}) *SuccessfulMessage {
	return &SuccessfulMessage{
		codeTime{
			code,
			time.Now().Format(time.RFC3339),
		},
		body,
	}
}

func setErrorMessage(code int, detailed interface{}) *ErrorMessage {
	return &ErrorMessage{
		codeTime{
			code,
			time.Now().Format(time.RFC3339),
		},
		message[code],
		detailed,
	}
}
