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
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("Failed to redirect STDOUT")
	}
	os.Stdout = w
	log.SetOutput(w)
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			buf.WriteString(scanner.Text())
		}
		log.Printf("Scanner error: %v\n", scanner.Err())
	}()

	// Create test HTTP server
	ts := httptest.NewServer(LogMiddleware(GetTestHandler()))
	defer ts.Close()

	// Trigger a request to get output to log
	http.Get(fmt.Sprintf("%s/api/user/", ts.URL))

	// Reset output
	w.Close()
	os.Stdout = stdout
	log.SetOutput(os.Stdout)

	// Test output
	t.Log("LogMiddleware output: ", buf)
	if buf.Len() == 0 {
		t.Error("No information logged to STDOUT")
	}
	if strings.Count(buf.String(), "\n") > 1 {
		t.Error("Expected only a single line of log output")
	}

	// type args struct {
	// 	next http.Handler
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want http.Handler
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := LogMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
