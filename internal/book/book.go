package book

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
    ErrNotImplement = errors.New("Not implemented")
)

type Book struct {
    Id uuid.UUID
    AccountId uuid.UUID
    Title string
    Author string
    Year int
    UpVotes int
}

type Store interface {
    GetBook(ctx context.Context, id uuid.UUID) (Book, error)
    CreateBook(ctx context.Context, book Book) (Book, error)
    UpdateBook(ctx context.Context, book Book) (Book, error)
    DeleteBook(ctx context.Context, id uuid.UUID) error
    UpVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error
    DownVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error
    GetUpVoteCount(ctx context.Context, id uuid.UUID) (int, error)
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
    fmt.Println("Retrieving Book")
    book, BookErr := s.Store.GetBook(ctx, id)
    if BookErr != nil {
        fmt.Println(BookErr)
        return Book{}, nil
    }
    return book, nil
}

func (s *Service) CreateBook(ctx context.Context, book Book) (Book, error) {
    fmt.Println("Creating Book")
    bk, BookErr := s.Store.CreateBook(ctx, book)
    if BookErr != nil {
        fmt.Println(BookErr)
        return Book{}, nil
    }
    return bk, nil
}

func (s *Service) UpdateBook(ctx context.Context, book Book) (Book, error) {
    fmt.Println("Updating Book")
    bk, acctErr := s.Store.UpdateBook(ctx, book)
    if acctErr != nil {
        fmt.Println(acctErr)
        return Book{}, nil
    }
    return bk, nil
}

func (s *Service) DeleteBook(ctx context.Context, id uuid.UUID) error {
    fmt.Println("Deleting Book")
    BookErr := s.Store.DeleteBook(ctx, id)
    if BookErr != nil {
        fmt.Println(BookErr)
        return nil
    }
    return nil
}

func (s *Service) GetUpVoteCount(ctx context.Context, id uuid.UUID) (int, error) {
    fmt.Println("Upvoting Book")
    count, upVoteErr := s.Store.GetUpVoteCount(ctx, id)
    if upVoteErr != nil {
        fmt.Println(upVoteErr)
        return -1, upVoteErr
    }
    return count, nil
}

func (s *Service) UpVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error {
    fmt.Println("Upvoting Book")
    upVoteErr := s.Store.UpVoteBook(ctx, accountId, bookId)
    if upVoteErr != nil {
        fmt.Println(upVoteErr)
        return nil
    }
    return nil
}

func (s *Service) DownVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error {
    return nil
}
