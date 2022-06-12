package main

import (
	"database/sql"
	"final-project-engineering-56/Backend/database"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	res, err := database.ConnectSQlite()
	if err != nil {
		fmt.Println("error open database", err)
	}
	fmt.Println(res)

	db, err := sql.Open("sqlite3", "./Bukuku.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL);
		
		INSERT INTO users(username, password, email) VALUES
			("nanda", "nanda123","nanda@gmail.com");`)

	if err != nil {
		panic(err)
	}
}
