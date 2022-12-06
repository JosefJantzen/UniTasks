package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"unitasks.josefjantzen.de/backend/auth"
	"unitasks.josefjantzen.de/backend/database"
)

type ApiService struct {
	DB *database.DBService
}

func NewApiService(s *database.DBService) *ApiService {
	return &ApiService{DB: s}
}

func (s *ApiService) Welcome(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Id)))
}

func (s *ApiService) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println("2", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.SignIn(w, r, s.DB, creds)
}

func (s *ApiService) SignUp(w http.ResponseWriter, r *http.Request) {
	auth.SignUp(w, r, s.DB)
}

func (s *ApiService) GetTaskById(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := s.DB.GetTaskById(id)
	if task == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if task.ParentUser != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(task)
}

func (s *ApiService) GetTasksByUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	tasks := s.DB.GetTasksByUser(claims.Id)
	fmt.Println("t2 ", tasks)
	if tasks == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}
