package handlers

import (
	"encoding/json"
	"gohttp/mocks"
	"gohttp/models"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	var UpdateBook models.Book
	if err := json.Unmarshal(body, &UpdateBook); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	for index, book := range mocks.Books {
		if book.Id == id {
			book.Title = UpdateBook.Title
			book.Author = UpdateBook.Author
			book.Desc = UpdateBook.Desc

			mocks.Books[index] = book

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode("Updated!")
			break
		}
	}
}
