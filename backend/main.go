package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("asd")
	initServer()
}

func initServer() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//myRouter.Handle("/", isAuthorized(homePage))

	/*myRouter.Handle("/tasks", isAuthorized(returnAllTasks))

	myRouter.Handle("/tasks/{id}", isAuthorized(returnSingleTask))
	myRouter.Handle("/task", isAuthorized(createNewTask)).Methods("POST")
	myRouter.Handle("/task/{id}", isAuthorized(deleteTask)).Methods("DELETE")
	myRouter.Handle("/task/{id}", isAuthorized(updateTask)).Methods("PUT")*/

	//  Start HTTP
	//go func() {
	err_http := http.ListenAndServe(":8080", myRouter)
	if err_http != nil {
		log.Fatal("Web server (HTTP): ", err_http)
	}
	//}()

	//  Start HTTPS
	//err_https := http.ListenAndServeTLS(":8081", "server.crt", "server.key", myRouter)
	/*if err_https != nil {
		log.Fatal("Web server (HTTPS): ", err_https)
	}*/
}
