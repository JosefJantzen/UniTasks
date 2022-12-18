package database

import (
	"database/sql"
	"fmt"
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
			fmt.Println("InitDB failed to open db connection. Error: ", err)
		}
		return err
	}
	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		fmt.Println("InitDB failed backoff retry for opening db connection. Error: ", err)
		return nil
	}
	body, err := os.ReadFile(config.DB.Initial)
	if err != nil {
		fmt.Println("InitDB error: unable to read file: ", config.DB.Initial)
	}
	if _, err := db.Exec(string(body)); err != nil {
		fmt.Println("InitDB initial error: ", err)
		return nil
	}
	if err != nil {
		fmt.Print("InitDB error: ")
		fmt.Println(fmt.Printf("failed to initialise the store: %s", err))
	}

	if config.Debug {
		body, err := os.ReadFile(config.DB.TestData)
		if err != nil {
			fmt.Println("InitDB error: unable to read file: ", config.DB.TestData)
		}
		if _, err := db.Exec(string(body)); err != nil {
			fmt.Println("InitDB test data error: ", err)
			return nil
		}
		if err != nil {
			fmt.Print("InitDB error: ")
			fmt.Println(fmt.Printf("failed to initialise the store with test data: %s", err))
		}
	}

	service := NewDBService(db)
	fmt.Println("Finished DB init")
	return service
}
