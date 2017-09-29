package main

import (
	"net/http"
)

func response(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello"))
}

func fail(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Fail na"))
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/fail", fail)
	http.ListenAndServe(":9000", nil)
}
