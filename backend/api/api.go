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

	myRouter.HandleFunc("/v1/api/signin", apiService.SignIn).Methods("POST")
	myRouter.HandleFunc("/v1/api/signup", apiService.SignUp).Methods("POST")
	myRouter.HandleFunc("/v1/api/logout", auth.Logout).Methods("POST")
	myRouter.HandleFunc("/v1/api/refresh", auth.Refresh).Methods("GET")
	myRouter.Handle("/", auth.Auth(apiService.Welcome)).Methods("GET")

	myRouter.Handle("/v1/api/tasks/{id}", auth.Auth(apiService.GetTaskById)).Methods("POST")
	myRouter.Handle("/v1/api/tasks", auth.Auth(apiService.GetTasksByUser)).Methods("POST")
	//myRouter.Handle("/", auth.Auth(Welcome))

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
