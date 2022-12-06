package api

import (
	"fmt"
	"net/http"

	"unitasks.josefjantzen.de/backend/auth"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	claims := &auth.Claims{}
	auth.Auth(w, r, claims)
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Id)))
}
