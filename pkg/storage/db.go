package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Nicolas-ggd/go-notification/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func NewDB() (*sql.DB, error) {
	dbName, err := removeAndCreateNew("notification")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	err = runMigration(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func removeAndCreateNew(name string) (string, error) {
	err := os.Remove(name)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	log.Printf("Creating new sqlite-%s.db", name)

	file, err := os.Create(fmt.Sprintf("%s.db", name))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = file.Close()
	if err != nil {
		return "", err
	}

	log.Printf("sqlite-%s.db created", name)

	return fmt.Sprintf("%s.db", name), nil
}

func runMigration(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{
		NoTxWrap:     false,
		DatabaseName: config.DatabaseName,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(config.MigrationURL, config.DatabaseName, driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Migrations applied successfully!")

	return nil
}
