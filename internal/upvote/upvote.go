package upvote

import "github.com/google/uuid"

type Upvote struct {
    AccountId uuid.UUID
    BookId uuid.UUID
}
