package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andruixxd31/beginner-project/internal/book"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BooksService interface {
    GetBook(ctx context.Context, id uuid.UUID) (book.Book, error)
    CreateBook(ctx context.Context, book book.Book) (book.Book, error)
    UpdateBook(ctx context.Context, book book.Book) (book.Book, error)
    DeleteBook(ctx context.Context, id uuid.UUID) error
    UpVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error
    DownVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error
    GetUpVoteCount(ctx context.Context, id uuid.UUID) (int, error)
}


func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
    reqVars := mux.Vars(r)    
    id := reqVars["id"]

    if id == ""{
        w.WriteHeader(http.StatusBadGateway)
        return
    }
    book, err := h.BooksService.GetBook(r.Context(), uuid.MustParse(id))
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
    }

    if err := json.NewEncoder(w).Encode(book); err != nil {
        panic(err)
    }     
    w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
    var bookRequest book.Book
    fmt.Println("req: ", r.Body)
    if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    fmt.Println(bookRequest)

    book, err := h.BooksService.CreateBook(r.Context(), bookRequest)
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(book); err != nil {
        panic(err)
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
    var bookRequest book.Book
    vars := mux.Vars(r)
    id := vars["id"]
    
    if id == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    errDecode := json.NewDecoder(r.Body).Decode(&bookRequest)
    if errDecode != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    book, err := h.BooksService.UpdateBook(r.Context(), bookRequest)
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
     
    if err := json.NewEncoder(w).Encode(book); err != nil {
        panic(err)
    }
    w.WriteHeader(http.StatusOK)

}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    if id == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    err := h.BooksService.DeleteBook(r.Context(),uuid.MustParse(id)) 
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    errEncode := json.NewEncoder(w).Encode(Response{Message: "Succesfully deleted the book"})
    if errEncode != nil {
        panic(err)
    }
    w.WriteHeader(http.StatusOK)

}

func (h *Handler) UpVoteBook(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DownVoteBook(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetUpVoteCount(w http.ResponseWriter, r *http.Request) {

}
