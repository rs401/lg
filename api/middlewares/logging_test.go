package middlewares

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func GetTestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {}
	return http.HandlerFunc(fn)
}

func TestLogMiddleware(t *testing.T) {
	buf := &bytes.Buffer{}

	// Redirect STDOUT to a buffer
	stdout := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("Failed to redirect STDOUT")
	}
	os.Stderr = w
	// Try either using stderr ^ or just setoutput bellow to the buffer
	log.SetOutput(w)
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			buf.WriteString(scanner.Text())
		}
	}()

	// Create test HTTP server
	ts := httptest.NewServer(LogMiddleware(GetTestHandler()))
	defer ts.Close()

	// Trigger a request to get output to log
	http.Get(fmt.Sprintf("%s/api/user/", ts.URL))

	// Reset output
	w.Close()
	os.Stderr = stdout
	log.SetOutput(os.Stderr)

	// Test output
	t.Log("LogMiddleware output: ", buf)
	if buf.Len() == 0 {
		t.Error("No information logged to STDOUT")
	}
	if strings.Count(buf.String(), "\n") > 1 {
		t.Error("Expected only a single line of log output")
	}

}
