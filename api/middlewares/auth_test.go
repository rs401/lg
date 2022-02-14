package middlewares

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// This is passing but shouldn't so it's not working correctly yet
func TestAuthMiddleware(t *testing.T) {
	myNext := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
	req := httptest.NewRequest(http.MethodGet, "http://site.com/", nil)
	res := httptest.NewRecorder()
	req.Header.Set("Authorization", "badtokenstring")
	myNext(res, req)
	AuthMiddleware(myNext)
	if res.Result().StatusCode != http.StatusOK {
		t.Errorf("Bad status code: %v\n", res.Result().StatusCode)
	}
	t.Log(res.Result().StatusCode)
	log.Println(res.Result().StatusCode)

}
