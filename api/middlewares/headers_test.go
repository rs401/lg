package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadersMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	hm := HeadersMiddleware(http.HandlerFunc(handler))
	hm.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}
