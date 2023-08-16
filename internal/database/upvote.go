package database

import "github.com/google/uuid"

type UpvoteRow struct {
    AccountId uuid.UUID
    BookId uuid.UUID
}
