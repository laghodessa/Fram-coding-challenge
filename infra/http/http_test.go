package http_test

import (
	"bytes"
	"database/sql"
	stdhttp "net/http"
	"net/http/httptest"
	"personia/infra/http"
	"personia/infra/sqlite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEndpoints(t *testing.T) {
	cases := []struct {
		name      string
		apiSecret string // override the default secret header
		method    string
		url       string
		code      int
		req       string
		resp      string
	}{
		{
			name:   "when request without Authorization header",
			method: stdhttp.MethodGet,
			url:    "/api/hierarchy",
			code:   400,
			resp:   `{"message":"missing key in request header"}`,
		},
		{
			name:      "when request with invalid Authorization header",
			apiSecret: "wrongsecret",
			method:    stdhttp.MethodGet,
			url:       "/api/hierarchy",
			code:      401,
			resp:      `{"message":"Unauthorized"}`,
		},
		{
			name:      "when query empty hierarhcy, it succeeds",
			apiSecret: "secret",
			method:    stdhttp.MethodGet,
			url:       "/api/hierarchy",
			code:      200,
			resp:      "{}",
		},
	}

	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	sqlite.MigrateDB(db)

	s := http.NewServer(http.ServerOpts{
		DB:        db,
		APISecret: "secret",
	})

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, bytes.NewBufferString(tc.req))
		if tc.apiSecret != "" {
			req.Header.Set("Authorization", "Bearer "+tc.apiSecret)
		}
		resp := httptest.NewRecorder()
		s.ServeHTTP(resp, req)

		// assertions
		assert.Equal(t, tc.code, resp.Code)
		assert.JSONEq(t, tc.resp, resp.Body.String())
	}
}
