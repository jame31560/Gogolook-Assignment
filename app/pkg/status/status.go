package status

import (
	"fmt"
	"net/http"
)

type Status interface {
	Error() string
	WithMsg(string) Status
	WithHttpCode(int) Status

	HttpCode() int
	ErrMsg() string
}

type status struct {
	httpCode int
	title    string
	message  string
}

func (s status) Error() string {
	return s.message
}

func (s status) WithMsg(message string) Status {
	sCopy := s
	sCopy.message = message
	return sCopy
}

func (s status) WithHttpCode(code int) Status {
	sCopy := s
	sCopy.httpCode = code
	return sCopy
}

func (s status) HttpCode() int {
	return s.httpCode
}

func (s status) ErrMsg() string {
	return fmt.Sprintf("%s: %s", s.title, s.message)
}

func NewStatus(title string) Status {
	return &status{
		title:    title,
		httpCode: http.StatusInternalServerError,
	}
}

var (
	SuccessStatus       = NewStatus("Success")
	ErrorStatus         = NewStatus("Error")
	CreateError         = NewStatus("Create Error")
	InternalServerError = NewStatus("Unknown Error").WithHttpCode(http.StatusInternalServerError)
)
