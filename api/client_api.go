package api

import (
	"log"
	"net/http"

	"github.com/kurozakizz/go-rest-api/app"
)

func GetClientsHandler(rw http.ResponseWriter, r *http.Request) {
	clients, err := app.GetClients()
	if err != nil {
		panic(err)
	}
	log.Print(clients)
	rw.Write([]byte("Get all clients"))
}

func CreateClientHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Create client"))
}
