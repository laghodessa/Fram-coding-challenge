package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"personia/domain/hr"
	apphttp "personia/infra/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	cases := []struct {
		name string
		err  error
		code int
		resp string
	}{
		{
			name: "it handles domain error",
			err:  hr.ErrHierarchyHasLoop,
			code: http.StatusUnprocessableEntity,
			resp: `
{
  "code": "hierarchy_has_loop",
  "message": "hierarchy has loop"
}`,
		},
		{
			name: "it respects echo HTTPError",
			err:  echo.NewHTTPError(http.StatusBadRequest, "echo message"),
			code: http.StatusBadRequest,
			resp: `
{
  "message": "echo message"
}`,
		},
		{
			name: "it returns 500 on unknown error",
			err:  errors.New("unknown"),
			code: http.StatusInternalServerError,
			resp: `
{
	"code": "internal",
  "message": "internal server error"
}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			resp := httptest.NewRecorder()

			apphttp.ErrorHandler()(tc.err, e.NewContext(req, resp))
			assert.Equal(t, tc.code, resp.Code)
			assert.JSONEq(t, tc.resp, resp.Body.String())
		})
	}
}
