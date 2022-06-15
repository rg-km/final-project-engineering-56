package api

import (
	"final-project-engineering-56/Backend/controller"

	"fmt"
	"net/http"
)

type API struct {
	userRepo controller.UserLogin
	bukuRepo controller.Buku
	mux      *http.ServeMux
}

func NewApi(userRepo controller.UserLogin, bukuRepo controller.Buku) API {
	mux := http.NewServeMux()
	api := API{
		userRepo, bukuRepo, mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))       // POST
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))     // POST
	mux.Handle("/api/user/register", api.POST(http.HandlerFunc(api.Register))) // POST
	mux.Handle("/api/user/all", api.GET(http.HandlerFunc(api.GetallUser)))     // GET

	//buku
	mux.Handle("/api/buku/getall", api.GET(http.HandlerFunc(api.Getallbuku))) // GET
	mux.Handle("/api/buku/addbuku", api.POST(http.HandlerFunc(api.InputBuku)))

	return api
}

func (api *API) Handler() http.Handler {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
