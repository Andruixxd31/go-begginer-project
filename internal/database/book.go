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

func (db *DB) UpdateBook(ctx context.Context, dbBook book.Book) (book.Book, error) {
    updateRow := BookRow{
        Id: dbBook.Id,
        Title: dbBook.Title,
        Author: dbBook.Author,
        Year: sql.NullInt32{Int32: int32(dbBook.Year), Valid: true},
    }
    row, err := db.Client.NamedQueryContext(
        ctx,
        `UPDATE book SET
        Title = :title,
        Author = :author,
        Year = :year
        WHERE id = :id
        `,
        updateRow,
    )
    if err != nil {
        return book.Book{}, fmt.Errorf("error creating user by given values %w", err)
    }
    if err := row.Close(); err != nil {
        return book.Book{}, fmt.Errorf("error closing rows: %w", err)
    }
    return dbBook, nil
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
    var bookRow BookRow
    row := db.Client.QueryRowContext(
        ctx,
        `SELECT likes
        FROM book 
        WHERE id = $1`,
        id,
    )
    err := row.Scan(&bookRow.Upvotes)
    if err != nil {
        return  -1, fmt.Errorf("error fetching upvote count by id: %w", err)
    }

    return convertBookRowToBook(bookRow).UpVotes, nil
}

func (db *DB) UpdateUpvoteBookCount(ctx context.Context, bookId uuid.UUID) error {
    row, err := db.Client.QueryContext(
        ctx,
        `UPDATE book SET
        likes = tbc.count FROM (SELECT COUNT(*) FROM upvote WHERE book_id = $1) as tbc
        WHERE id = $1
        `,
        bookId,
    )
    if err != nil {
        return fmt.Errorf("error updating upvote count with given id:%w", err)
    }
    if err := row.Close(); err != nil {
        return fmt.Errorf("error closing rows: %w", err)
    }
    return nil
}

func (db *DB) UpVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error {
    postRow := UpvoteRow{
        AccountId: accountId,
        BookId: bookId,
    }
    row, err := db.Client.NamedQueryContext(
        ctx,
        `
        INSERT INTO upvote (account_id, book_id)
        VALUES (:accountid, :bookid)
        ON CONFLICT (account_id, book_id)
        DO NOTHING
        `,
        postRow,
    )
    if err != nil {
        return fmt.Errorf("error creating upvote by given ids %w", err)
    }
    if err := row.Close(); err != nil {
        return fmt.Errorf("error closing rows: %w", err)
    }

    if err := db.UpdateUpvoteBookCount(ctx, bookId); err != nil {
        return fmt.Errorf("error updating upvote count after upvote %w", err)
    }
    return nil
}

func (db *DB) DownVoteBook(ctx context.Context, accountId uuid.UUID, bookId uuid.UUID) error {
    _, err := db.Client.ExecContext(
        ctx,
        `DELETE FROM upvote
        WHERE account_id = $1 AND book_id = $2
        `,
        accountId, bookId,
    )
    if err != nil {
        return fmt.Errorf("Error downvoting book: %w", err)
    }

    if err := db.UpdateUpvoteBookCount(ctx, bookId); err != nil {
        return fmt.Errorf("error updating upvote count after upvote %w", err)
    }
    return nil
}
