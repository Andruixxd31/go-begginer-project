// go:build integration
package database

import (
	"context"
	"testing"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/stretchr/testify/assert"
)

func TestAccountTable(t *testing.T) {
    t.Run("test create account", func(t *testing.T) {
        db, err := NewDatabase()
        assert.NoError(t, err)

        account, err := db.CreateAccount(context.Background(), account.Account{
            Name: "John",
        })
        assert.NoError(t, err)

        newAccount, err := db.GetAccount(context.Background(), account.Id)
        assert.NoError(t, err)
        assert.Equal(t, newAccount.Name, account.Name)
    })

    t.Run("test delete account", func(t *testing.T) {
        db, err := NewDatabase()
        assert.NoError(t, err)

        account, err := db.CreateAccount(context.Background(), account.Account{
            Name: "Juan",
        })
        assert.NoError(t, err)

        delErr := db.DeleteAccount(context.Background(), account.Id)
        assert.NoError(t, delErr)

        _, getErr := db.GetAccount(context.Background(), account.Id)
        assert.Error(t, getErr)
    })

    t.Run("test update account", func(t *testing.T) {
        db, err := NewDatabase()
        assert.NoError(t, err)

        createdAccount, err := db.CreateAccount(context.Background(), account.Account{
            Name: "John",
        })
        assert.NoError(t, err)

        updatedAccount, updateErr := db.UpdateAccount(context.Background(), createdAccount.Id, account.Account{
            Name: "Juan",
        })
        assert.NoError(t, updateErr)

        newAccount, err := db.GetAccount(context.Background(), createdAccount.Id)
        assert.NoError(t, err)
        assert.Equal(t, newAccount.Name, updatedAccount.Name)
    })

}
