package api

import (
	"encoding/json"
	"final-project-engineering-56/Backend/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginSuccess struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AuthError struct {
	Error string `json:"error"`
}

type claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

var key = []byte("secret")

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthError{Error: "Login failed"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(AuthError{Error: "Failed to generate token"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 60),
	})

	json.NewEncoder(w).Encode(LoginSuccess{Username: user.Username, Token: tokenString})
}

func (api *API) register(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var user model.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Regis := api.userRepo.Register(user.Username, user.Password, user.Email)
	if Regis == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthError{Error: "Register failed"})
		return
	}

	json.NewEncoder(w).Encode(&Regis)
}

func (api *API) logout(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if token.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout success"))
}
