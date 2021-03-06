package api

import (
	"final-project-engineering-56/Backend/controller"

	"fmt"
	"net/http"
)

type API struct {
	userRepo    controller.UserLogin
	bukuRepo    controller.Buku
	adminRepo   controller.Admin
	historyRepo controller.UserBookHistory
	mux         *http.ServeMux
}

func NewApi(userRepo controller.UserLogin, bukuRepo controller.Buku, adminRepo controller.Admin, historyRepo controller.UserBookHistory) API {
	mux := http.NewServeMux()
	api := API{
		userRepo, bukuRepo, adminRepo, historyRepo, mux,
	}
	//user
	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))       // POST
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))     // POST
	mux.Handle("/api/user/register", api.POST(http.HandlerFunc(api.Register))) // POST
	mux.Handle("/api/user/update", api.POST(http.HandlerFunc(api.UpdateUser))) // POST
	mux.Handle("/api/user/delete", api.POST(http.HandlerFunc(api.DeleteUser))) // POST
	mux.Handle("/api/user/getall", api.GET(http.HandlerFunc(api.GetallUser)))  // GET
	// mux.Handle("/api/user/getiduser", api.GET(http.HandlerFunc(api.GetById)))  // GET}")

	//admin
	mux.Handle("/api/admin/login", api.POST(http.HandlerFunc(api.AdminLogin)))         // POST
	mux.Handle("/api/admin/logout", api.POST(http.HandlerFunc(api.AdminLogout)))       // POST
	mux.Handle("/api/admin/register", api.POST(http.HandlerFunc(api.AdminInputAdmin))) //POST
	mux.Handle("/api/admin/update", api.POST(http.HandlerFunc(api.AdminUpdateAdmin)))  //POST
	mux.Handle("/api/admin/delete", api.POST(http.HandlerFunc(api.AdminDeleteAdmin)))  //POST
	mux.Handle("/api/admin/getall", api.GET(http.HandlerFunc(api.AdminGetAllAdmin)))   //GET
	// mux.Handle("/api/admin/getidadmin", api.GET(http.HandlerFunc(api.AdminGetById))) //GET

	//buku

	mux.Handle("/api/buku/addbuku", api.POST(http.HandlerFunc(api.InputBuku)))     // POST
	mux.Handle("/api/buku/delete", api.POST(http.HandlerFunc(api.DeleteBukuByID))) // POST
	mux.Handle("/api/buku/update", api.POST(http.HandlerFunc(api.UpdateBukuByID))) // POST
	mux.Handle("/api/buku/setfavorit", api.POST(http.HandlerFunc(api.setFavorit))) // POST
	mux.Handle("/api/buku/getidbuku", api.GET(http.HandlerFunc(api.GetById)))      // GET
	mux.Handle("/api/buku/getall", api.GET(http.HandlerFunc(api.Getallbuku)))      // GET
	mux.Handle("/api/buku/search", api.GET(http.HandlerFunc(api.searchBooks)))     // GET
	mux.Handle("/api/buku/fetch", api.GET(http.HandlerFunc(api.fetchBooks)))       // GET
	mux.Handle("/api/buku/getfav", api.GET(http.HandlerFunc(api.GetFav)))       // GET

	return api
}

func (api *API) Handler() http.Handler {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
