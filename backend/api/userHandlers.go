package api

import (
	"encoding/json"
	"net/http"

	"unitasks.josefjantzen.de/backend/auth"
)

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
