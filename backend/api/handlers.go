package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"unitasks.josefjantzen.de/backend/auth"
	"unitasks.josefjantzen.de/backend/database"
)

type ApiService struct {
	DB *database.DBService
}

func NewApiService(s *database.DBService) *ApiService {
	return &ApiService{DB: s}
}

type EMail struct {
	Id   uuid.UUID `json:"id"`
	Mail string    `json:"eMail"`
}

type AllTasks struct {
	Tasks          []database.Task          `json:"tasks"`
	RecurringTasks []database.RecurringTask `json:"recurringTasks"`
}

func (s *ApiService) Welcome(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Id)))
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

	recurringTasks, err := s.DB.GetRecurringTasksByUser(claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if recurringTasks == nil {
		recurringTasks = []database.RecurringTask{}
	}

	allTasks := AllTasks{Tasks: tasks, RecurringTasks: recurringTasks}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(allTasks)
}
