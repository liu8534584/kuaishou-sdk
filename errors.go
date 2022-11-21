package kuaishou

import (
	"time"
)

type ApiError struct {
	Code       int
	ErrMessage string
	Time       time.Time
}

func NewApiError(code int, errMessage string) *ApiError {
	return &ApiError{
		Code:       code,
		ErrMessage: errMessage,
		Time:       time.Now(),
	}
}

func (err *ApiError) Error() string {
	return err.ErrMessage
}

func Ok() *ApiError {
	return &ApiError{
		Code:       0,
		ErrMessage: "ok",
		Time:       time.Now(),
	}
}
