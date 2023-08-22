package database

import (
	"github.com/andruixxd31/beginner-project/internal/upvote"
    "github.com/google/uuid"
)

type UpvoteRow struct {
    AccountId uuid.UUID
    BookId uuid.UUID
}

func convertUpvoteRowToUpvote(upvoteRow UpvoteRow) upvote.Upvote {
    return upvote.Upvote{
        AccountId: upvoteRow.AccountId,
        BookId: upvoteRow.BookId,
    }
}
