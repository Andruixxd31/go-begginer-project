package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type BooksService interface {

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

    h.Router.HandleFunc("/api/v1/account/{id}", h.GetAccount).Methods("GET")
    h.Router.HandleFunc("/api/v1/account", h.CreateAccount).Methods("POST")
    h.Router.HandleFunc("/api/v1/account", h.UpdateAccount).Methods("PUT")
    h.Router.HandleFunc("/api/v1/account/{id}", h.DeleteAccount).Methods("DELETE")
}

func (h *Handler) Serve() error {
    go func() {
        if err := h.Server.ListenAndServe(); err != nil {
            log.Println(err.Error())
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <- c

    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    h.Server.Shutdown(ctx)

    log.Println("shut down gracefully")
    return nil
}
