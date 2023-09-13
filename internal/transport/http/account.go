package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
    "github.com/go-playground/validator/v10"
)

type AccountsService interface {
    GetAccount(ctx context.Context, id uuid.UUID) (account.Account, error)
    CreateAccount(ctx context.Context, account account.Account) (account.Account, error)
    UpdateAccount(ctx context.Context, id uuid.UUID, account account.Account) (account.Account, error)
    DeleteAccount(ctx context.Context, id uuid.UUID) error
}

type CreateAccountRequest struct {
    Name string `json:"name" validate:"required"`
}

type UpdateAccountRequest struct {
    Name string `json:"name"`
}

type IdAccountRequest struct {
    Id uuid.UUID `json:"id"`
}

func convertCreateAccountRequestToAccount(a CreateAccountRequest) account.Account {
    return account.Account{
        Name: a.Name,
    }
}

func convertUpdateAccountRequestToAccount(a UpdateAccountRequest) account.Account {
    return account.Account{
        Name: a.Name,
    }
}


func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        json.NewEncoder(w).Encode(Response{Message: errors.New("No id provided").Error()})
        w.WriteHeader(http.StatusBadGateway)
        return
    }

    resAccount, err := h.AccountsService.GetAccount(r.Context(), uuid.MustParse(id))
    if err != nil {
        json.NewEncoder(w).Encode(Response{Message: err.Error()})
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(resAccount); err != nil {
        panic(err)
    }     

    if (account.Account{}) == resAccount {
        if err := json.NewEncoder(w).Encode(Response{Message: "Succesfully Deleted Account"}); err != nil {
            panic(err)
        }
    }
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    var account CreateAccountRequest
    if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
        json.NewEncoder(w).Encode(Response{Message: errors.New("Invalid body request").Error()})
        return
    }

    validate := validator.New()
    err := validate.Struct(account)
    if err != nil {
        http.Error(w, "Not a valid account", http.StatusBadGateway)
        return
    }

    convertedAccount := convertCreateAccountRequestToAccount(account)

    postedAccount, err := h.AccountsService.CreateAccount(r.Context(), convertedAccount)
    if err != nil {
        json.NewEncoder(w).Encode(Response{Message: errors.New("No id provided").Error()})
        panic(err)
    }

    if err := json.NewEncoder(w).Encode(postedAccount); err != nil {
        panic(err)
    }     
}

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r) 
    var account UpdateAccountRequest

    id := vars["id"]
    if id == "" {
        json.NewEncoder(w).Encode(Response{Message: errors.New("No id provided").Error()})
        w.WriteHeader(http.StatusBadGateway)
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
        return
    }

    convertedAccount := convertUpdateAccountRequestToAccount(account)

    updatedAccount, err := h.AccountsService.UpdateAccount(r.Context(), uuid.MustParse(id), convertedAccount)
    if err != nil {
        json.NewEncoder(w).Encode(Response{Message: errors.New("No id provided").Error()})
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(updatedAccount); err != nil {
        panic(err)
    }     
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r) 
    id := vars["id"]
    if id == "" {
        json.NewEncoder(w).Encode(Response{Message: errors.New("No id provided").Error()})
        w.WriteHeader(http.StatusBadGateway)
        return
    }

    err := h.AccountsService.DeleteAccount(r.Context(), uuid.MustParse(id))
    if err != nil {
        json.NewEncoder(w).Encode(Response{Message: errors.New("No id provided").Error()})
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(Response{Message: "Succesfully Deleted Account"}); err != nil {
        panic(err)
    }

}
