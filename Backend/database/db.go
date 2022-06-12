package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQlite() (string, error) {
	db, err := sql.Open("sqlite3", "./Bukuku.db")
	if err != nil {
		fmt.Println("error open database", err)
	}
	defer db.Close()
	return "You are connected to SQLite", nil
}
