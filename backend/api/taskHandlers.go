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

func (s *ApiService) GetTaskById(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := s.DB.GetTaskById(id, claims.Id)
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
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
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

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
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
	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateTaskDone error: ", err)
		return
	}

	task.Id = id
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
	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateTaskDone error: ", err)
		return
	}

	task.Id = id
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
