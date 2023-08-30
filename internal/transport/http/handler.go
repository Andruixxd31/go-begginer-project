package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BooksService interface {

}

type AccountsService interface {

}

type Handler struct {
    Router *mux.Router
    BooksService BooksService
    AccountsService AccountsService
    Server *http.Server
}


func NewHandler(booksService BooksService, accountsService AccountsService) *Handler {
    h := &Handler{
        BooksService: booksService,
        AccountsService: accountsService,
    }
    h.Router = mux.NewRouter()

    h.mapRoutes()
    h.Server = &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: h.Router,
    }
    return h
}


func (h *Handler) mapRoutes(){
    h.Router.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world")
    })
}

func (h *Handler) Serve() error {
    if err := h.Server.ListenAndServe(); err != nil {
        return err
    }

    return nil
}
