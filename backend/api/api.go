package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"unitasks.josefjantzen.de/backend/auth"
	"unitasks.josefjantzen.de/backend/config"
	"unitasks.josefjantzen.de/backend/database"
)

func InitServer(dbService *database.DBService, config *config.Config) {
	fmt.Print("Starting Server: ")

	apiService := NewApiService(dbService)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/", auth.Auth(apiService.Welcome)).Methods("GET")

	myRouter.HandleFunc("/api/v1/signin", func(w http.ResponseWriter, r *http.Request) { apiService.SignIn(config, w, r) }).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/api/v1/signup", func(w http.ResponseWriter, r *http.Request) { apiService.SignUp(config, w, r) }).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/api/v1/logout", auth.Logout).Methods("GET")
	myRouter.HandleFunc("/api/v1/refresh", auth.Refresh).Methods("GET")
	myRouter.Handle("/api/v1/deleteUser", auth.Auth(apiService.DeleteUser)).Methods("DELETE", "OPTIONS")
	myRouter.Handle("/api/v1/updateMail", auth.Auth(apiService.UpdateMail)).Methods("PUT", "OPTIONS")
	myRouter.Handle("/api/v1/updatePwd", auth.Auth(apiService.UpdatePwd)).Methods("PUT", "OPTIONS")

	myRouter.Handle("/api/v1/tasks/{id}", auth.Auth(apiService.GetTaskById)).Methods("GET")
	myRouter.Handle("/api/v1/tasks", auth.Auth(apiService.GetTasksByUser)).Methods("GET")
	myRouter.Handle("/api/v1/tasks", auth.Auth(apiService.InsertTask)).Methods("POST", "OPTIONS")
	myRouter.Handle("/api/v1/tasks/{id}", auth.Auth(apiService.UpdateTask)).Methods("PUT", "OPTIONS")
	myRouter.Handle("/api/v1/tasks/{id}/done", auth.Auth(apiService.UpdateTaskDone)).Methods("PUT", "OPTIONS")
	myRouter.Handle("/api/v1/tasks/{id}", auth.Auth(apiService.DeleteTask)).Methods("DELETE", "OPTIONS")

	myRouter.Handle("/api/v1/recurring-tasks/{id}", auth.Auth(apiService.GetRecurringTaskById)).Methods("GET")
	myRouter.Handle("/api/v1/recurring-tasks", auth.Auth(apiService.GetRecurringTasksByUser)).Methods("GET")
	myRouter.Handle("/api/v1/recurring-tasks", auth.Auth(apiService.InsertRecurringTask)).Methods("POST", "OPTIONS")
	myRouter.Handle("/api/v1/recurring-tasks/{id}", auth.Auth(apiService.UpdateRecurringTask)).Methods("PUT", "OPTIONS")
	myRouter.Handle("/api/v1/recurring-tasks/{id}", auth.Auth(apiService.DeleteRecurringTask)).Methods("DELETE", "OPTIONS")

	myRouter.Handle("/api/v1/recurring-tasks/{id}/history", auth.Auth(apiService.GetRecurringTasksHistory)).Methods("GET")
	myRouter.Handle("/api/v1/recurring-tasks-history", auth.Auth(apiService.GetRecurringTasksHistoryByUser)).Methods("GET")
	myRouter.Handle("/api/v1/recurring-tasks-history", auth.Auth(apiService.InsertRecurringTaskHistory)).Methods("POST", "OPTIONS")
	myRouter.Handle("/api/v1/recurring-tasks-history/{id}", auth.Auth(apiService.UpdateRecurringTaskHistory)).Methods("PUT", "OPTIONS")
	myRouter.Handle("/api/v1/recurring-tasks-history/{id}/done", auth.Auth(apiService.UpdateRecurringTaskHistoryDone)).Methods("PUT", "OPTIONS")
	myRouter.Handle("/api/v1/recurring-tasks-history/{id}", auth.Auth(apiService.DeleteRecurringTaskHistory)).Methods("DELETE", "OPTIONS")
	myRouter.Handle("/api/v1/recurring-tasks/{id}/history", auth.Auth(apiService.DeleteCompleteRecurringTaskHistory)).Methods("DELETE", "OPTIONS")

	myRouter.Handle("/api/v1/all", auth.Auth(apiService.GetAllTasksByUser)).Methods("GET")

	//  Start HTTP
	go func() {
		err_http := http.ListenAndServe(":"+config.Port, myRouter)
		fmt.Println(err_http)
		if err_http != nil {
			log.Fatal("Web server (HTTP): ", err_http)
		}
	}()
	fmt.Println("Finished starting Server")
}
