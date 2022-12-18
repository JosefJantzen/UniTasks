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

func (s *ApiService) GetRecurringTaskById(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := s.DB.GetRecurringTaskById(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("GetRecurringTaskById error: ", err)
		return
	}
	if task == nil {
		*task = database.RecurringTask{}
	}
	if task.UserId != claims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(task)
}

func (s *ApiService) GetRecurringTasksByUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	tasks, err := s.DB.GetRecurringTasksByUser(claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("GetRecurringTasksByUser error: ", err)
		return
	}
	if tasks == nil {
		tasks = []database.RecurringTask{}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func (s *ApiService) InsertRecurringTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTask
	json.Unmarshal(reqBody, &task)

	task.UserId = claims.Id

	id, err := s.DB.InsertRecurringTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("InsertRecurringTask error: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "{\"id\": \""+id.String()+"\" }")
}

func (s *ApiService) UpdateRecurringTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateRecurringTask error: ", err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTask
	json.Unmarshal(reqBody, &task)

	task.Id = id
	task.UserId = claims.Id

	err = s.DB.UpdateRecurringTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateRecurringTask error: ", err)
		return
	}
}

func (s *ApiService) DeleteRecurringTask(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteRecurringTask error: ", err)
		return
	}

	err = s.DB.DeleteRecurringTask(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteRecurringTask error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
