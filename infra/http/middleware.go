package http

import (
	"errors"
	"fmt"
	"net/http"
	"personia/domain"

	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	Code    string                 `json:"code,omitempty"`
	Message string                 `json:"message"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
	status  int
	cause   error
}

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		herr := toHTTPError(err)

		// logs internal error
		if herr.status == http.StatusInternalServerError {
			c.Logger().Error(herr.cause)
		}

		// Send response
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(herr.status)
		} else {
			err = c.JSON(herr.status, herr)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func toHTTPError(err error) HTTPError {
	var derr *domain.Error
	if ok := errors.As(err, &derr); ok {
		return HTTPError{
			Code:    derr.Code,
			Message: derr.Message,
			Meta:    derr.Meta,
			status:  codeToHTTPStatus(derr.Code),
			cause:   nil,
		}
	}

	var herr *echo.HTTPError
	if ok := errors.As(err, &herr); ok {
		return HTTPError{
			status:  herr.Code,
			Message: fmt.Sprintf("%v", herr.Message),
		}
	}

	return HTTPError{
		Code:    "internal",
		Message: "internal server error",
		status:  http.StatusInternalServerError,
		cause:   err,
	}
}

func codeToHTTPStatus(code string) int {
	switch code {
	default:
		return http.StatusUnprocessableEntity
	}
}
