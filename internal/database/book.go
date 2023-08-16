package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andruixxd31/beginner-project/internal/book"
	"github.com/google/uuid"
)

type BookRow struct {
    Id uuid.UUID
    AccountId uuid.UUID
    Title string
    Author string
    Year sql.NullInt32
    Upvotes sql.NullInt32
    CreatedAt sql.NullTime 
    UpdatedAt sql.NullTime
    DeletedAt sql.NullTime
}

func convertBookRowToBook(bookRow BookRow) book.Book {
    return book.Book{
        Id: bookRow.Id,
        AccountId: bookRow.Id,
        Title: bookRow.Title,
        Author: bookRow.Author,
        Year: int(bookRow.Year.Int32),
        UpVotes: int(bookRow.Upvotes.Int32),
    }
}

func (db *DB) GetBook(ctx context.Context, uuid uuid.UUID) (book.Book, error) {
    var bookRow BookRow
    row := db.Client.QueryRowContext(
        ctx,
        `SELECT id, account_id, title, author, year, likes
        FROM book 
        WHERE id = $1`,
        uuid,
    )
    err := row.Scan(&bookRow.Id, &bookRow.AccountId, &bookRow.Title, &bookRow.Author, &bookRow.Year, &bookRow.Upvotes)
    if err != nil {
        return book.Book{}, fmt.Errorf("error fetching book by uuid: %w", err)
    }

    return convertBookRowToBook(bookRow), nil
}
