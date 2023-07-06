package book

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
    ErrNotImplement = errors.New("Not implemented")
)

type Book struct {
    Id uuid.UUID
    Title string
    ISBN string
    Author string
    Genre string
    UpVotes int
}

type Store interface {
    GetBook(ctx context.Context, id uuid.UUID) (Book, error)
    CreateBook(ctx context.Context, book Book) (Book, error)
    UpdateBook(ctx context.Context, book Book) error
    DeleteBook(ctx context.Context, id uuid.UUID) error
    UpVoteBook(ctx context.Context, id uuid.UUID) error
    GetUpVoteCount(ctx context.Context, id uuid.UUID) int
}

type Service struct {
    Store Store
}

func NewService(store Store) *Service {
    return &Service{
        Store: store,
    }
}

func (s *Service) GetBook(ctx context.Context, id uuid.UUID) (Book, error) {
    return Book{}, nil
}

func (s *Service) CreateBook(ctx context.Context, book Book) (Book, error) {
    return Book{}, nil
}

func (s *Service) UpdateBook(ctx context.Context, book Book) error {
    return nil
}

func (s *Service) DeleteBook(ctx context.Context, id uuid.UUID) error {
    return nil
}

func (s *Service) GetUpVoteCount(ctx context.Context, id uuid.UUID) error {
    return nil
}

func (s *Service) UpVoteBook(ctx context.Context, id uuid.UUID) int {
    return 0
}
