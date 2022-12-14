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

type EMail struct {
	Id   uuid.UUID `json:"id"`
	Mail string    `json:"eMail"`
}

func NewApiService(s *database.DBService) *ApiService {
	return &ApiService{DB: s}
}

type AllTasks struct {
	Tasks          []database.Task          `json:"tasks"`
	RecurringTasks []database.RecurringTask `json:"recurringTasks"`
}

func (s *ApiService) Welcome(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Id)))
}

func (s *ApiService) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.SignIn(w, r, s.DB, creds)
}

func (s *ApiService) SignUp(w http.ResponseWriter, r *http.Request) {
	auth.SignUp(w, r, s.DB)
}

func (s *ApiService) UpdateMail(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	var mail EMail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if mail.Id != claims.Id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.DB.UpdateMail(claims.Id, mail.Mail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s *ApiService) UpdatePwd(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	var pwd auth.Password
	err := json.NewDecoder(r.Body).Decode(&pwd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if pwd.Id != claims.Id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.UpdatePwd(w, r, s.DB, pwd)
}

func (s *ApiService) DeleteUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	var creds auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.DeleteUser(w, r, creds, s.DB)
}

func (s *ApiService) GetTaskById(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := s.DB.GetTaskById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("GetTaskById error: ", err)
		return
	}
	if task == nil {
		*task = database.Task{}
	}
	if task.UserId != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(task)
}

func (s *ApiService) GetTasksByUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	tasks, err := s.DB.GetTasksByUser(claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("GetTasksByUser error: ", err)
		return
	}
	if tasks == nil {
		tasks = []database.Task{}
	}

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func (s *ApiService) InsertTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.Task
	json.Unmarshal(reqBody, &task)
	task.UserId = claims.Id
	id, err := s.DB.InsertTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("InsertTask error: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "{\"id\": \""+id.String()+"\" }")
}

func (s *ApiService) UpdateTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateTask error: ", err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.Task
	json.Unmarshal(reqBody, &task)

	task.Id = id

	if task.UserId != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	task.UserId = claims.Id
	err = s.DB.UpdateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateTask error: ", err)
		return
	}
}

func (s *ApiService) UpdateTaskDone(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateTaskDone error: ", err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.Task
	json.Unmarshal(reqBody, &task)

	task.Id = id

	if task.UserId != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	task.UserId = claims.Id
	err = s.DB.UpdateTaskDone(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateTaskDone error: ", err)
		return
	}
}

func (s *ApiService) DeleteTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteTask error: ", err)
		return
	}

	err = s.DB.DeleteTask(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteTask error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	if task.UserId != claims.Id {
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
	task.UserId = claims.Id
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

	if task.UserId != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	task.UserId = claims.Id
	err = s.DB.UpdateRecurringTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s *ApiService) DeleteRecurringTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.DB.DeleteRecurringTask(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ApiService) GetAllTasksByUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	tasks, err := s.DB.GetTasksByUser(claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if tasks == nil {
		tasks = []database.Task{}
	}

	recurringTasks := s.DB.GetRecurringTasksByUser(claims.Id)
	if recurringTasks == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	allTasks := AllTasks{Tasks: tasks, RecurringTasks: recurringTasks}

	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	json.NewEncoder(w).Encode(allTasks)
}
