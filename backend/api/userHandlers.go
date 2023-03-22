package api

import (
	"encoding/json"
	"net/http"

	"unitasks.josefjantzen.de/backend/auth"
	"unitasks.josefjantzen.de/backend/config"
)

func (s *ApiService) SignIn(config *config.Config, w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var creds auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.SignIn(w, r, s.DB, creds, config)
}

func (s *ApiService) SignUp(config *config.Config, w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	auth.SignUp(w, r, s.DB, config)
}

func (s *ApiService) UpdateMail(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	auth.HandleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var mail EMail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mail.Id = claims.Id

	err = s.DB.UpdateMail(claims.Id, mail.Mail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s *ApiService) UpdatePwd(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	auth.HandleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
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
	auth.HandleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var creds auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth.DeleteUser(w, r, creds, s.DB)
}
