package main

import (
	"log"
	"net/http"
	"project/CloudApp/internal/handlers"
	"project/CloudApp/internal/repositories/dataBases/inMem"

	"github.com/gorilla/mux"
)

func main() {
	store := inMem.New()
	handlers := handlers.New(store)
	RunHttp(*handlers)
}

// curl -X PUT -d 'Hello, key-value store!' -v http://localhost:8080/v1/key-a
// curl -v http://localhost:8080/v1/key-a
// curl -X DELETE -v http://localhost:8080/v1/key-a

func RunHttp(handlers handlers.Handlers) {
	log.Println("HTTP server on :8080")

	r := mux.NewRouter()

	r.HandleFunc("/v1/{key}", handlers.PutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", handlers.GetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", handlers.DeleteHandler).Methods("DELETE")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
