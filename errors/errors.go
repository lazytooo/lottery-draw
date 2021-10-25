package errors

import (
	"errors"
	"github.com/lazytooo/lottery-draw/models"
)

var (
	ErrorCodeOK   = errors.New("success")
	InvalidParams = errors.New("invalid params")
)

// GetErrorCode ...
func GetErrorBaseResponse(err error) (result models.BaseResponse) {
	result = models.BaseResponse{}
	if err == nil {
		result.Msg = ErrorCodeOK.Error()
	} else {
		result.Msg = err.Error()
	}

	switch err {
	case InvalidParams:
		result.Code = 10001
		return
	case nil:
		result.Code = 10000
		return
	default:
		result.Code = 99999
		return
	}
}
