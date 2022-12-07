package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"unitasks.josefjantzen.de/backend/auth"
	"unitasks.josefjantzen.de/backend/database"
)

func InitServer(dbService *database.DBService) {
	fmt.Print("Starting Server: ")

	apiService := NewApiService(dbService)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/", auth.Auth(apiService.Welcome)).Methods("GET")

	myRouter.HandleFunc("/v1/api/signin", apiService.SignIn).Methods("POST")
	myRouter.HandleFunc("/v1/api/signup", apiService.SignUp).Methods("POST")
	myRouter.HandleFunc("/v1/api/logout", auth.Logout).Methods("GET")
	myRouter.HandleFunc("/v1/api/refresh", auth.Refresh).Methods("GET")

	myRouter.Handle("/v1/api/tasks/{id}", auth.Auth(apiService.GetTaskById)).Methods("GET")
	myRouter.Handle("/v1/api/tasks", auth.Auth(apiService.GetTasksByUser)).Methods("GET")
	myRouter.Handle("/v1/api/tasks", auth.Auth(apiService.InsertTask)).Methods("POST")
	myRouter.Handle("/v1/api/tasks/{id}", auth.Auth(apiService.UpdateTask)).Methods("POST")
	myRouter.Handle("/v1/api/tasks/{id}", auth.Auth(apiService.DeleteTask)).Methods("DELETE")

	myRouter.Handle("/v1/api/recurring-tasks/{id}", auth.Auth(apiService.GetRecurringTaskById)).Methods("GET")
	myRouter.Handle("/v1/api/recurring-tasks", auth.Auth(apiService.GetRecurringTasksByUser)).Methods("GET")
	myRouter.Handle("/v1/api/recurring-tasks", auth.Auth(apiService.InsertRecurringTask)).Methods("POST")
	myRouter.Handle("/v1/api/recurring-tasks/{id}", auth.Auth(apiService.UpdateRecurringTask)).Methods("POST")
	myRouter.Handle("/v1/api/recurring-tasks/{id}", auth.Auth(apiService.DeleteRecurringTask)).Methods("DELETE")

	/**
	* TODO: DELETE: Beides
	* TODO: get both recurring and normal tasks for one user
	 */

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
		// generate SSL certificate: openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout server.key -out server.crt
		err_https := http.ListenAndServeTLS(":443", "server.crt", "server.key", myRouter)
		if err_https != nil {
			log.Fatal("Web server (HTTPS): ", err_https)
		}
	}()
	fmt.Println("Finished starting Server")
}
