// error_handler
package model

import (
	"encoding/json"
)

type ErrorType string

const (
	ErrSuccess    ErrorType = "success"
	ErrNoMsgFound           = "can not found message info in request"
)

var (
	errCode = map[ErrorType]int{
		ErrNoMsgFound: 10001,
	}
)

type Response struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  ErrorType   `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func ParseResult(errMsg ErrorType, result interface{}) []byte {
	code := errCode[errMsg]
	if code == 0 && errMsg != ErrSuccess {
		code = -1
	}
	r := Response{
		ErrCode: code,
		ErrMsg:  errMsg,
		Data:    result,
	}

	b, err := json.Marshal(r)
	if err != nil {

	}
	return b
}
