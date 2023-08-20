package main

import (
	"context"
	"fmt"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/andruixxd31/beginner-project/internal/database"
	"github.com/google/uuid"
)

func Run() error{
    fmt.Println("Starting up app")
    db, err := database.NewDatabase()
    if err != nil {
        return fmt.Errorf("Failed to connect to db: %w", err)
    }

    if err := db.MigrateDB(); err != nil {
        fmt.Println("Failed to migrate db")
        return err
    }
    accountService := account.NewService(db)
    fmt.Println(accountService.GetAccount(context.Background(), uuid.MustParse("94ba2858-0be6-4c31-b967-9f3fbf20f755")))
    fmt.Println(accountService.UpdateAccount(context.Background(), account.Account{Id: uuid.MustParse("8773488e-f4cc-45bd-859b-8367244a9fe4"), Name: "Andrew"}))
    return nil
}

func main(){
    fmt.Println("Go app")
    if err := Run(); err != nil {
        fmt.Println(err)
    }
}
