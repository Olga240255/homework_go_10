package main

import (
	"log"
	"net/http"

	"github.com/rinatkh/homework_backend_1/internal/httpapi"
	"github.com/rinatkh/homework_backend_1/internal/note"
)

func main() {
	router := httpapi.NewRouter(note.NewStore())
	const address = ":8080"
	log.Printf("Notes API запущен на http://localhost%s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
