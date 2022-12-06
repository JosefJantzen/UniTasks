package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
