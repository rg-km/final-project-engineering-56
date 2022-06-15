package main

import (
	"database/sql"
	"final-project-engineering-56/Backend/api"
	"final-project-engineering-56/Backend/controller"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "D:/Final-project-engineering-56/Backend/database/BukuKu.db")
	if err != nil {
		panic(err)
	}

	usersRepo := controller.NewUserLogin(db)
	booksRepo := controller.NewBuku(db)

	mainAPI := api.NewApi(*usersRepo, *booksRepo)
	mainAPI.Start()
}
