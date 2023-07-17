package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (db *DB) MigrateDB() error {
    fmt.Println("migrating db")

    driver, err := postgres.WithInstance(db.Client.DB, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("Could not create the postgres driver")  
    }

    migration, err := migrate.NewWithDatabaseInstance(
        "file:///migrations",
        "postgres",
        driver,
    )
    if err != nil {
        fmt.Println(err)
        return err
    }
    if err := migration.Up(); err != nil {
        if !errors.Is(err, migrate.ErrNoChange){
            return fmt.Errorf("Could not run up migrations: %w", err)
        }
    }

    fmt.Println("Succesfully migrated db")
    return nil
}
