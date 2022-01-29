package main

import (
	"net/http"
)

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":8080", nil)

}
