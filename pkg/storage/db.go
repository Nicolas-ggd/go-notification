package storage

import (
	"database/sql"
	"errors"
	"github.com/Nicolas-ggd/go-notification/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.DBName)
	if err != nil {
		return nil, err
	}

	if err = checkAndRunMigration(db); err != nil {
		return nil, err
	}

	return db, nil
}

func checkAndRunMigration(db *sql.DB) error {
	// check if tables exist
	_, err := db.Exec("SELECT 1 FROM notifications LIMIT 1;")
	if err != nil {
		// if the table doesn't exist, run the migrations
		log.Println("Running migrations as tables don't exist")
		err = runMigration(db)
		if err != nil {
			return err
		}
	} else {
		log.Println("Database and tables already exist")
	}

	return nil
}

func runMigration(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{
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
