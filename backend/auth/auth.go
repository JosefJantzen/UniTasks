package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"unitasks.josefjantzen.de/backend/database"
)

var expireMin = 5

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

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
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return ""
	}
	return tokenString
}

func Auth(endpoint func(w http.ResponseWriter, r *http.Request, c *Claims)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			return jwtKey, nil
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

func SignIn(w http.ResponseWriter, r *http.Request, s *database.DBService, creds Credentials) {
	user, err := s.GetUserByMail(creds.EMail)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(creds.Pwd, user.Pwd) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Duration(expireMin) * time.Minute)

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

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func SignUp(w http.ResponseWriter, r *http.Request, s *database.DBService) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if s.CheckMailUsed(creds.EMail) {
		SignIn(w, r, s, creds)
		return
	}

	pwd, err := hashPassword(creds.Pwd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := s.InsertUser(creds.EMail, pwd)

	expirationTime := time.Now().Add(time.Duration(expireMin) * time.Minute)

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

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
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
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}

func Refresh(w http.ResponseWriter, r *http.Request) {
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
		return jwtKey, nil
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

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
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
