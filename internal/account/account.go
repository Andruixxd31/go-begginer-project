package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)


var (
    ErrNotImplement = errors.New("Not implemented")
)

type Account struct {
    Id uuid.UUID
    Name string
}

type Store interface {
    GetAccount(ctx context.Context, id uuid.UUID) (Account, error)
    CreateAccount(ctx context.Context, account Account) (Account, error)
    UpdateAccount(ctx context.Context, id uuid.UUID, account Account) (Account, error)
    DeleteAccount(ctx context.Context, id uuid.UUID) error
}

type Service struct {
    Store Store
}

func NewService(store Store) *Service {
    return &Service{
        Store: store,
    }
}

func (s *Service) GetAccount(ctx context.Context, id uuid.UUID) (Account, error) {
    fmt.Println("Retrieving Account")
    accnt, acctErr := s.Store.GetAccount(ctx, id)
    if acctErr != nil {
        return Account{}, acctErr
    }
    return accnt, nil
}

func (s *Service) CreateAccount(ctx context.Context, account Account) (Account, error) {
    fmt.Println("Creating Account")
    accnt, acctErr := s.Store.CreateAccount(ctx, account)
    if acctErr != nil {
        return Account{}, acctErr

    }
    return accnt, nil
}

func (s *Service) UpdateAccount(ctx context.Context, id uuid.UUID, account Account) (Account, error) {
    fmt.Println("Updating Account")
    accnt, acctErr := s.Store.UpdateAccount(ctx, id, account)
    if acctErr != nil {
        return Account{}, acctErr
    }
    return accnt, nil
}

func (s *Service) DeleteAccount(ctx context.Context, id uuid.UUID) error {
    fmt.Println("Deleting Account")
    acctErr := s.Store.DeleteAccount(ctx, id)
    if acctErr != nil {
        return acctErr
    }
    return nil
}
