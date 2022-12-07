package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	if tasks == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func (s *ApiService) InsertTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.Task
	json.Unmarshal(reqBody, &task)
	task.ParentUser = claims.Id
	id := s.DB.InsertTask(task)
	if id == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, id)
}

func (s *ApiService) UpdateTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.Task
	json.Unmarshal(reqBody, &task)

	task.Id = id
	task.ParentUser = claims.Id

	if task.ParentUser != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = s.DB.UpdateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s *ApiService) GetRecurringTaskById(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := s.DB.GetRecurringTaskById(id)
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

func (s *ApiService) GetRecurringTasksByUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	tasks := s.DB.GetRecurringTasksByUser(claims.Id)
	if tasks == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func (s *ApiService) InsertRecurringTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTask
	json.Unmarshal(reqBody, &task)
	task.ParentUser = claims.Id
	id := s.DB.InsertRecurringTask(task)
	if id == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, id)
}

func (s *ApiService) UpdateRecurringTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTask
	json.Unmarshal(reqBody, &task)

	task.Id = id
	task.ParentUser = claims.Id

	if task.ParentUser != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = s.DB.UpdateRecurringTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
