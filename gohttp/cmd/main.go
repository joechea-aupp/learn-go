package main

import (
	"gohttp/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", handlers.UpdateBook).Methods(http.MethodPatch)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
