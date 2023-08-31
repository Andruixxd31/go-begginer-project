package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/google/uuid"
)

type AccountsService interface {
    GetAccount(ctx context.Context, id uuid.UUID) (account.Account, error)
    CreateAccount(ctx context.Context, account account.Account) (account.Account, error)
    UpdateAccount(ctx context.Context, account account.Account) (account.Account, error)
    DeleteAccount(ctx context.Context, id uuid.UUID) error
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
     
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "CreateAccount")
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
     
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
     
}
