package api

import (
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
	auth.SignIn(w, r, s.DB)
}

func (s *ApiService) GetUser(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {

}
