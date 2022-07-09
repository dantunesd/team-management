package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-chi/chi/v5"
)

const payloadExample = `{
    "name": "example",
    "type": "employee",
    "typeData": {
        "role": "software eng"
    },
    "tags": [
        "backend",
        "frontend",
        "ops"
    ]
}`

func getHttpRequestAndResponse(path string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(body))
	res := httptest.NewRecorder()
	return req, res
}

func setIdParameter(req *http.Request) *http.Request {
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
}
