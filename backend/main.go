package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cenkalti/backoff"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Server: ")
	initServer()
	fmt.Println("Finished")

	db, err := initStore()
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(fmt.Printf("failed to initialise the store: %s", err))
	}
	defer db.Close()

	fmt.Println("Completly Started")
	select {}
}

func initServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	})

	/*myRouter.Handle("/tasks", isAuthorized(returnAllTasks))

	myRouter.Handle("/tasks/{id}", isAuthorized(returnSingleTask))
	myRouter.Handle("/task", isAuthorized(createNewTask)).Methods("POST")
	myRouter.Handle("/task/{id}", isAuthorized(deleteTask)).Methods("DELETE")
	myRouter.Handle("/task/{id}", isAuthorized(updateTask)).Methods("PUT")*/

	//  Start HTTP
	go func() {
		err_http := http.ListenAndServe(":8080", myRouter)
		fmt.Println(err_http)
		if err_http != nil {
			log.Fatal("Web server (HTTP): ", err_http)
		}
	}()

	//  Start HTTPS
	go func() {
		err_https := http.ListenAndServeTLS(":443", "server.crt", "server.key", myRouter)
		if err_https != nil {
			log.Fatal("Web server (HTTPS): ", err_https)
		}
	}()

}

func initStore() (*sql.DB, error) {
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
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS users (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), e_mail STRING, pwd STRING);"); err != nil {
		return nil, err
	}
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS recurring_tasks (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), name STRING, interval INTERVAL, parentUser UUID REFERENCES users(id));"); err != nil {
		return nil, err
	}
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS tasks (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), name STRING, due TIME, description STRING, parentTask UUID REFERENCES recurring_tasks(id));"); err != nil {
		return nil, err
	}
	fmt.Println("End")
	return db, err
}
