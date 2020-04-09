package errors_helper

import (
	"errors"
	"fmt"
)

type ApplicationError struct {
	 origin    error
	 errorCode ErrorCode
}

func NewApplicationError(errorCode ErrorCode, args ...interface{}) error {
	var insertionString = getInsertionString(errorCode)
	return &ApplicationError{
		origin: errors.New(fmt.Sprintf(insertionString, args)),
		errorCode: errorCode,
	}
}

func (err ApplicationError) Error() string {
	return err.origin.Error()
}

func (err ApplicationError) Code() ErrorCode {
	return err.errorCode
}

func getInsertionString(errorCode ErrorCode) string {

	for _, info := range insertionStrings {
		if info.ErrorCode == errorCode {
			return info.msgString
		}
	}

	return ""
}