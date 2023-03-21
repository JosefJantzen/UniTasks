package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"unitasks.josefjantzen.de/backend/config"
	"unitasks.josefjantzen.de/backend/database"
)

type Credentials struct {
	EMail string `json:"eMail"`
	Pwd   string `json:"pwd"`
}

type Claims struct {
	Id uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type Password struct {
	Id  uuid.UUID `json:"id"`
	Pwd string    `json:"pwd"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func genJWT(claims *Claims, w http.ResponseWriter) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKeyBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return ""
	}
	return tokenString
}

func HandleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", config.FrontendUrl)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")
}

func createCookie(value string, expires time.Time, config *config.Config) *http.Cookie {
	var s = strings.Replace(config.FrontendUrl, "https://", "", 1)
	s = strings.Replace(s, "http://", "", 1)

	return &http.Cookie{
		Name:     "token",
		Value:    value,
		Expires:  expires,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Domain:   s,
	}
}

func Auth(endpoint func(w http.ResponseWriter, r *http.Request, c *Claims)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HandleCors(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := cookie.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKeyBytes, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		endpoint(w, r, claims)
	})
}

func SignIn(w http.ResponseWriter, r *http.Request, s *database.DBService, creds Credentials, config *config.Config) {
	user, err := s.GetUserByMail(creds.EMail)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(creds.Pwd, user.Pwd) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Duration(config.JwtExpireMin) * time.Minute)

	claims := &Claims{
		Id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	tokenString := genJWT(claims, w)
	if tokenString == "" {
		return
	}

	http.SetCookie(w, createCookie(tokenString, expirationTime, config))
}

func SignUp(w http.ResponseWriter, r *http.Request, s *database.DBService, config *config.Config) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if creds.EMail == "" || creds.Pwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if s.CheckMailUsed(creds.EMail) {
		SignIn(w, r, s, creds, config)
		return
	}

	pwd, err := hashPassword(creds.Pwd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("SignUp error: ", err)
		return
	}

	id, err := s.InsertUser(creds.EMail, pwd)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("SignUp error: ", err)
		return
	}

	expirationTime := time.Now().Add(time.Duration(config.JwtExpireMin) * time.Minute)

	claims := &Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	tokenString := genJWT(claims, w)
	if tokenString == "" {
		return
	}

	http.SetCookie(w, createCookie(tokenString, expirationTime, config))
}

func UpdatePwd(w http.ResponseWriter, r *http.Request, s *database.DBService, p Password) {
	pwd, err := hashPassword(p.Pwd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.UpdatePwd(p.Id, pwd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	HandleCors(w)
	http.SetCookie(w, createCookie("", time.Now(), &config.Config{}))

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	HandleCors(w)
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKeyBytes, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Until(claims.ExpiresAt.Time) > 2*time.Minute+30*time.Second {
		w.WriteHeader(http.StatusOK)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKeyBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Domain:   "unitasks.josefjantzen.de",
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request, creds Credentials, s *database.DBService) {
	user, err := s.GetUserByMail(creds.EMail)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(creds.Pwd, user.Pwd) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = s.DeleteUser(user.Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Logout(w, r)
}
