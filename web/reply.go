package web

import (
	"encoding/json"
	"net/http"

	"github.com/swanwish/go-helper/logs"
)

const (
	JsonKeyErrorNo      = "code"
	JsonKeyErrorMessage = "msg"
	JsonKeyData         = "data"
	JsonKeyList         = "list"

	StatusCodeInvalidParameter = 400
	StatusCodeInternalError    = 500
	StatusCodeOK               = 200

	ErrorMessageInternalError    = "Internal Error"
	ErrorMessageInvalidParameter = "Invalid Parameter"
)

func ReplyError(rw http.ResponseWriter, errorMessage string, errorCode int) {
	logs.Errorf("Error message: %s and error code is: %d", errorMessage, errorCode)
	data := make(map[string]interface{})
	data[JsonKeyErrorMessage] = errorMessage
	data[JsonKeyErrorNo] = errorCode
	ReplyJson(rw, data)
}

func ReplyInvalidParameterError(rw http.ResponseWriter) {
	ReplyError(rw, ErrorMessageInvalidParameter, StatusCodeInvalidParameter)
}

func ReplyInternalError(rw http.ResponseWriter) {
	ReplyError(rw, ErrorMessageInternalError, StatusCodeInternalError)
}

func ReplyJson(rw http.ResponseWriter, data map[string]interface{}) {
	if data[JsonKeyErrorNo] == nil {
		data[JsonKeyErrorNo] = StatusCodeOK
	}

	js, err := json.Marshal(data)
	if err != nil {
		panic("Marshal json failed." + err.Error())
	}

	//	logs.Debug("json data:", string(js))

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.Write(js)
}

func ReplyJsonData(rw http.ResponseWriter, data interface{}) {
	result := make(map[string]interface{}, 0)
	if data != nil {
		result[JsonKeyData] = data
	}
	ReplyJson(rw, result)
}

func ReplyJsonList(rw http.ResponseWriter, list interface{}) {
	result := make(map[string]interface{}, 0)
	result[JsonKeyList] = list
	ReplyJson(rw, result)
}
