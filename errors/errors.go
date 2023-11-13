package errors

import "fmt"

const (
	BillResponseError = "BillResponseError"
	InternalCodeError = "InternalCodeError"
)

type Error struct {
	srcError  error
	errorCode string
}

func NewError(code string, err error) *Error {
	return &Error{errorCode: code, srcError: err}
}

func (e *Error) Error() string {
	return fmt.Sprintf("ErrorCode: {%s}   SourceError: {%s}", e.errorCode, e.srcError.Error())
}
