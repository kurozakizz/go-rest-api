package api

import (
	"net/http"
)

func CreateRouter(mux *http.ServeMux) {
	mux.HandleFunc("/clients", GetClientsHandler)
	mux.HandleFunc("/client", CreateClientHandler)
}