package account


import (
	"context"
	"errors"

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
    UpdateAccount(ctx context.Context, account Account) error
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
    return Account{}, nil
}

func (s *Service) CreateAccount(ctx context.Context, account Account) (Account, error) {
    return Account{}, nil
}

func (s *Service) UpdateAccount(ctx context.Context, account Account) error {
    return nil
}

func (s *Service) DeleteAccount(ctx context.Context, id uuid.UUID) error {
    return nil
}
