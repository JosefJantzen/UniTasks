package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cenkalti/backoff"
)

func InitDB() (*sql.DB, error) {
	fmt.Print("Start DB init: ")
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	var (
		db  *sql.DB
		err error
	)
	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		if err != nil {
			fmt.Println(err)
		}
		return err
	}
	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		fmt.Println("Err23: ")
		fmt.Println(err)
		return nil, err
	}
	body, err := os.ReadFile("DB-initial.pgsql")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	if _, err := db.Exec(string(body)); err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(fmt.Printf("failed to initialise the store: %s", err))
	}
	defer db.Close()

	return db, err
}
