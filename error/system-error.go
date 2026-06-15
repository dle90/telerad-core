package errors

import (
	"errors"
)

type SystemError struct {
	errorCode int
	err       error
}

func (e *SystemError) Error() error {
	return e.err
}

func (e *SystemError) ErrorMessage() string {
	if e == nil || e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *SystemError) ErrorCode() int {
	return e.errorCode
}

func NewChecksumInvalid() *SystemError {
	return &SystemError{
		errorCode: CHECKSUM_INVALID,
		err:       errors.New(""),
	}
}

func New(err error) *SystemError {
	return &SystemError{
		errorCode: 0,
		err:       err,
	}
}

func NewErrorByString(err string) *SystemError {
	return &SystemError{
		errorCode: 0,
		err:       errors.New(err),
	}
}
