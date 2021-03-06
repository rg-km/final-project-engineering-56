package api

import (
	"encoding/json"
	"final-project-engineering-56/Backend/middleware"
	"final-project-engineering-56/Backend/model"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type LoginSuccess struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type RegisSuccess struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
	var user model.Users
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user = api.userRepo.Login(user.Email, user.Password)
	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to login"})
		return
	}

	claims := claims{
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

 	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 60),
	})

	json.NewEncoder(w).Encode(LoginSuccess{Username: user.Username, Email: user.Email, Token: tokenString})

}

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var Regis model.Users
	err := json.NewDecoder(r.Body).Decode(&Regis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.userRepo.Register(Regis.Username, middleware.Hashpassword(Regis.Password), Regis.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to register"})
		return
	}

	json.NewEncoder(w).Encode(RegisSuccess{Username: Regis.Username, Email: Regis.Email, Password: Regis.Password})
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

func (api *API) GetallUser(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var userList []model.Users
	userList, err := api.userRepo.GetALLUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to get all user"})
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(userList)
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var user model.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.userRepo.Update(user.ID, user.Username, user.Password, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to update user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update user success"))
}

func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var user model.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.userRepo.Delete(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to delete user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete user success"))
}

func (api *API) GetById(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user model.Users
	user, err = api.userRepo.GetbyID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to get user"})
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(user)

}
