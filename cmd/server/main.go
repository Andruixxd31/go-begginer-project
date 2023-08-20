package main

import (
	"context"
	"fmt"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/andruixxd31/beginner-project/internal/book"
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

    bookService := book.NewService(db)

    fmt.Println(bookService.UpdateBook(context.Background(), book.Book{
        AccountId: uuid.MustParse("94ba2858-0be6-4c31-b967-9f3fbf20f755"),
        Title: "Kinstugi",
        Author: "Maria Jose Navia",
        Year: 2018,
    }))

    fmt.Println(bookService.GetBook(context.Background(), uuid.MustParse("17e8300c-ff1c-4bd0-b24d-ecba842fd122")))
    return nil
}

func main(){
    fmt.Println("Go app")
    if err := Run(); err != nil {
        fmt.Println(err)
    }
}
