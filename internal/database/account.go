package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/google/uuid"
)

type AccountRow struct {
    Id uuid.UUID
    Name string
    CreatedAt sql.NullTime 
    UpdatedAt sql.NullTime
    DeletedAt sql.NullTime
}

func convertAccountRowToAccount(accountRow AccountRow) account.Account {
    return account.Account{
        Id: accountRow.Id,
        Name: accountRow.Name,
    }
}

func (db *DB) GetAccount(ctx context.Context, id uuid.UUID) (account.Account, error) {
    var accountRow AccountRow
    row := db.Client.QueryRowContext(
        ctx,
        `SELECT id, name
        FROM account
        WHERE id = $1
        `,
        id,
    )
    err := row.Scan(&accountRow.Id, &accountRow.Name)
    if err != nil {
        return account.Account{}, fmt.Errorf("error fetching user by id: %w", err)
    }
    
    return convertAccountRowToAccount(accountRow), nil
}

func (db *DB) CreateAccount(ctx context.Context, dbAccount account.Account) (account.Account, error) {
    dbAccount.Id = uuid.New()
    postRow := AccountRow{
        Id: dbAccount.Id,
        Name: dbAccount.Name,
    }
    row, err := db.Client.NamedQueryContext(
        ctx,
        `INSERT INTO account(id, name)
        VALUES(:id, :name)
        `,
        postRow,
    )
    if err != nil {
        return account.Account{}, fmt.Errorf("error creating user by given values %w", err)
    }
    if err := row.Close(); err != nil {
        return account.Account{}, fmt.Errorf("error closing rows: %w", err)
    }
    return dbAccount, nil
}

func (db *DB) DeleteAccount(ctx context.Context, id uuid.UUID) error {
    sqlResult, err := db.Client.ExecContext(
        ctx,
        `DELETE FROM account
        WHERE id = $1
        `,
        id,
    )
    if err != nil {
        return fmt.Errorf("Error deleting account: %w", err)
    }

    rowAffected, err := sqlResult.RowsAffected()
    fmt.Println("affected: ", rowAffected)
    if err != nil {
        return fmt.Errorf("%w", err)
    }
    if rowAffected == 0 {
        return fmt.Errorf("id does not correspond to any account: %w", err)
    }
    return nil
}

func (db *DB) UpdateAccount(ctx context.Context, id uuid.UUID, dbAccount account.Account) (account.Account, error) { 
    updateRow := AccountRow{
        Id: id,
        Name: dbAccount.Name,
    }
    row, err := db.Client.NamedQueryContext(
        ctx,
        `UPDATE account SET
        name = :name
        WHERE id = :id
        `,
        updateRow,
    )
    if err != nil {
        return account.Account{}, fmt.Errorf("error creating user by given values %w", err)
    }
    if err := row.Close(); err != nil {
        return account.Account{}, fmt.Errorf("error closing rows: %w", err)
    }
    return convertAccountRowToAccount(updateRow), nil
}
