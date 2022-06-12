package main

import (
	"database/sql"
	"final-project-engineering-56/Backend/api"
	"final-project-engineering-56/Backend/controller"
)

func main() {
	db, err := sql.Open("sqlite3", "final-project-engineering-56/Backend/database/migration/BukuKu.db")
	if err != nil {
		panic(err)
	}

	usersRepo := controller.NewUserLogin(db)

	mainAPI := api.NewApi(*usersRepo)
	mainAPI.Start()
}
