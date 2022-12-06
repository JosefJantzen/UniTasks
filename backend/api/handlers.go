package api

import (
	"fmt"
	"net/http"

	"unitasks.josefjantzen.de/backend/auth"
)

func Welcome(w http.ResponseWriter, r *http.Request, claims *auth.Claims) {
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Id)))
}
