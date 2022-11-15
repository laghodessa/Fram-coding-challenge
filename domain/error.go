package domain

import "fmt"

type Error struct {
	Code    string
	Message string
	Meta    map[string]interface{}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewError(code, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
