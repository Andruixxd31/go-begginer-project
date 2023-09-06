package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AccountsService interface {
    GetAccount(ctx context.Context, id uuid.UUID) (account.Account, error)
    CreateAccount(ctx context.Context, account account.Account) (account.Account, error)
    UpdateAccount(ctx context.Context, id uuid.UUID, account account.Account) (account.Account, error)
    DeleteAccount(ctx context.Context, id uuid.UUID) error
}


func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        w.WriteHeader(http.StatusBadGateway)
        return
    }

    account, err := h.AccountsService.GetAccount(r.Context(), uuid.MustParse(id))
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(account); err != nil {
        panic(err)
    }     
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    var account account.Account
    if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
        return
    }

    account, err := h.AccountsService.CreateAccount(r.Context(), account)
    if err != nil {
        panic(err)
    }

    if err := json.NewEncoder(w).Encode(account); err != nil {
        panic(err)
    }     
}

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r) 
    var account account.Account

    id := vars["id"]
    if id == "" {
        w.WriteHeader(http.StatusBadGateway)
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
        return
    }

    account, err := h.AccountsService.UpdateAccount(r.Context(), uuid.MustParse(id), account)
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(account); err != nil {
        panic(err)
    }     
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r) 
    id := vars["id"]
    if id == "" {
        w.WriteHeader(http.StatusBadGateway)
        return
    }

    err := h.AccountsService.DeleteAccount(r.Context(), uuid.MustParse(id))
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(Response{Message: "Succesfully Deleted Account"}); err != nil {
        panic(err)
    }

}
