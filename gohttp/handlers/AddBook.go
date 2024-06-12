package handlers

import (
	"encoding/json"
	"gohttp/mocks"
	"gohttp/models"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	log.Println("read something", string(body))

	var book models.Book
	if err := json.Unmarshal(body, &book); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// generate random id (simple one)
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // response code 201
	json.NewEncoder(w).Encode("Created!")
}
