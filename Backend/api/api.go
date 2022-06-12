package api

import (
	"final-project-engineering-56/Backend/controller"
	"fmt"
	"net/http"
)

type API struct {
	userRepo controller.UserLogin
	mux      *http.ServeMux
}

func NewApi(userRepo controller.UserLogin) API {
	mux := http.NewServeMux()
	api := API{
		userRepo, mux,
	}

	mux.Handle("/api/user/login", (http.HandlerFunc(api.login))) // POST
	return api
}

func (api *API) Handler() http.Handler {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
