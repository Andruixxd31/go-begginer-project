package main

import (
	"fmt"

	"github.com/andruixxd31/beginner-project/internal/account"
	"github.com/andruixxd31/beginner-project/internal/book"
	"github.com/andruixxd31/beginner-project/internal/database"
	transportHttp "github.com/andruixxd31/beginner-project/internal/transport/http"
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
    bookService := book.NewService(db)

    httpHandler := transportHttp.NewHandler(bookService, accountService)
    if err := httpHandler.Serve(); err != nil {
        return err
    }
    return nil
}

func main(){
    fmt.Println("Go app")
    if err := Run(); err != nil {
        fmt.Println(err)
    }
}
