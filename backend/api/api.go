package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"unitasks.josefjantzen.de/backend/auth"
)

func InitServer() {
	fmt.Println("Starting Server: ")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/", auth.Auth(Welcome))
	myRouter.HandleFunc("/signin", auth.Signin)
	myRouter.HandleFunc("/logout", auth.Logout)
	myRouter.HandleFunc("/refresh", auth.Refresh)

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
