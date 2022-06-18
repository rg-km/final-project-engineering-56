package api

import (
	"encoding/json"
	"final-project-engineering-56/Backend/middleware"
	"final-project-engineering-56/Backend/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ResponAdmin struct {
	Error string `json:"error"`
}

type Admin struct {
	Admin model.Admin `json:"admin"`
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var keyy = []byte("admin")

func (api *API) AdminLogin(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var admin model.Admin
	err := json.NewDecoder(req.Body).Decode(&admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims := Claims{
		Username: admin.Username,
		Password: admin.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(keyy)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Minute * 60),
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login success"))

}

func (api *API) AdminLogout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	token, err := req.Cookie("token")
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

func (api *API) AdminGetAllAdmin(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var adminList []model.Admin
	adminList, err := api.adminRepo.GetALLAdmin()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(ResponAdmin{Error: "Failed to get all admin"})
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(adminList)
}

func (api *API) AdminInputAdmin(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var admin model.Admin
	err := json.NewDecoder(req.Body).Decode(&admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.adminRepo.Register(admin.Username, middleware.Hashpassword(admin.Password), admin.Role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(ResponAdmin{Error: "Failed to input admin"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Input admin success"))
}

func (api *API) AdminUpdateAdmin(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var admin model.Admin
	err := json.NewDecoder(req.Body).Decode(&admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.adminRepo.Update(admin.ID, admin.Username, admin.Password, admin.Role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(ResponAdmin{Error: "Failed to update admin"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update admin success"))
}

func (api *API) AdminDeleteAdmin(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var admin model.Admin
	err := json.NewDecoder(req.Body).Decode(&admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.adminRepo.Delete(admin.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(ResponAdmin{Error: "Failed to delete admin"})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete admin success"))
}
