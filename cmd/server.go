package main

import (
	"net/http"

	"github.com/kurozakizz/go-rest-api/api"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	api.CreateRouter(mux)
	http.ListenAndServe(port, mux)
}
