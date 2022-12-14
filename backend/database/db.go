package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cenkalti/backoff"
)

type DBService struct {
	db *sql.DB
}

func NewDBService(d *sql.DB) *DBService {
	return &DBService{db: d}
}

func InitDB() *DBService {
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
		return nil
	}
	body, err := os.ReadFile("DB-initial.pgsql")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	if _, err := db.Exec(string(body)); err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(fmt.Printf("failed to initialise the store: %s", err))
	}

	service := NewDBService(db)
	fmt.Println("Finished DB init")
	return service
}
