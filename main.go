package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HenriqueSilverio/go-rest-api/books"
	"github.com/HenriqueSilverio/go-rest-api/health"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Server starting")
	router := chi.NewRouter()
	health.Register(router)
	books.Register(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
