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

func (s *ApiService) GetRecurringTasksHistory(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := s.DB.GetRecurringTasksHistory(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("GetRecurringTasksHistory error: ", err)
		return
	}
	if len(tasks) == 0 {
		tasks = []database.RecurringTaskHistory{}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func (s *ApiService) GetRecurringTasksHistoryByUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	tasks, err := s.DB.GetRecurringTaskHistoriesByUser(claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("GetRecurringTasksHistoryByUser error: ", err)
		return
	}
	if tasks == nil {
		tasks = []database.RecurringTaskHistory{}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(tasks)
}

func (s *ApiService) InsertRecurringTaskHistory(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTaskHistory
	json.Unmarshal(reqBody, &task)

	task.UserId = claims.Id
	if !s.DB.CheckRecurringTaskExists(task.RecurringTaskId, task.UserId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.DB.InsertRecurringTaskHistory(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("InsertRecurringTaskHistory error: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "{\"id\": \""+id.String()+"\" }")
}

func (s *ApiService) UpdateRecurringTaskHistory(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateRecurringTaskHistory error: ", err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTaskHistory
	json.Unmarshal(reqBody, &task)

	task.Id = id
	task.UserId = claims.Id
	if !s.DB.CheckRecurringTaskExists(task.RecurringTaskId, task.UserId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.DB.UpdateRecurringTaskHistory(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateRecurringTaskHistory error: ", err)
		return
	}
}

func (s *ApiService) UpdateRecurringTaskHistoryDone(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateRecurringTaskHistoryDone error: ", err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var task database.RecurringTaskHistory
	json.Unmarshal(reqBody, &task)

	task.Id = id
	task.UserId = claims.Id
	if !s.DB.CheckRecurringTaskExists(task.RecurringTaskId, task.UserId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.DB.UpdateRecurringTaskHistoryDone(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("UpdateRecurringTaskHistoryDone error: ", err)
		return
	}
}

func (s *ApiService) DeleteRecurringTaskHistory(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteRecurringTaskHistory error: ", err)
		return
	}

	err = s.DB.DeleteRecurringTaskHistory(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteRecurringTaskHistory error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ApiService) DeleteCompleteRecurringTaskHistory(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteCompleteRecurringTaskHistory error: ", err)
		return
	}

	err = s.DB.DeleteCompleteRecurringTaskHistory(id, claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DeleteCompleteRecurringTaskHistory error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
