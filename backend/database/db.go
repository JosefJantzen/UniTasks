package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cenkalti/backoff"
	"unitasks.josefjantzen.de/backend/config"
)

type DBService struct {
	db *sql.DB
}

func NewDBService(d *sql.DB) *DBService {
	return &DBService{db: d}
}

func InitDB(config *config.Config) *DBService {
	fmt.Print("Start DB init: ")
	var (
		db  *sql.DB
		err error
	)
	openDB := func() error {
		db, err = sql.Open("postgres", config.DB.GetDBConnectString())
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
