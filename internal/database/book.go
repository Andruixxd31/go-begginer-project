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

func (db *DB) CreateBook(ctx context.Context, dbBook book.Book) (book.Book, error) {
    dbBook.Id = uuid.New()
    postRow := BookRow{
        Id: dbBook.Id,
        AccountId: dbBook.AccountId,
        Title: dbBook.Title,
        Author: dbBook.Author,
        Year: sql.NullInt32{Int32: int32(dbBook.Year), Valid: true},
        Upvotes: sql.NullInt32{Int32: int32(dbBook.UpVotes), Valid: true},
    }
    row, err := db.Client.NamedQueryContext(
        ctx,
        `INSERT INTO book(id, account_id, title, author, year, likes)
        VALUES(:id, :accountid, :title, :author, :year, :upvotes)
        `,
        postRow,
    )
    if err != nil {
        return book.Book{}, fmt.Errorf("error creating book by given values %w", err)
    }
    if err := row.Close(); err != nil {
        return book.Book{}, fmt.Errorf("error closing rows: %w", err)
    }
    return dbBook, nil
}

func (db *DB) UpdateBook(ctx context.Context, dbBook book.Book) error {
    return nil
}

func (db *DB) DeleteBook(ctx context.Context, id uuid.UUID) error {
    _, err := db.Client.ExecContext(
        ctx,
        `DELETE FROM book
        WHERE id = $1
        `,
        id,
    )
    if err != nil {
        return fmt.Errorf("Error deleting book: %w", err)
    }
    return nil
}

func (db *DB) GetUpVoteCount(ctx context.Context, id uuid.UUID) (int, error) {
    return 0, nil
}

func (db *DB) UpVoteBook(ctx context.Context, id uuid.UUID) (int, error) {
    return 0, nil
}
