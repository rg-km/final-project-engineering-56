package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "D:/final-project-engineering-56/Backend/database/Bukuku.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL);

		CREATE TABLE IF NOT EXISTS admin (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(255) NOT NULL);
		
			CREATE TABLE IF NOT EXISTS books (
				idbuku VARCHAR(255) PRIMARY KEY ,
				ISBN VARCHAR(255) NOT NULL,
				judul VARCHAR(255) NOT NULL,
				penerbit VARCHAR(255) NOT NULL,
				pengarang VARCHAR(255) NOT NULL,
				tahun VARCHAR(255) NOT NULL,
				gambar VARCHAR(255) NOT NULL,
				deskripsi VARCHAR(255) NOT NULL);
	

		CREATE TABLE IF NOT EXISTS user_book_history (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				id_user VARCHAR(255) ,
				id_books VARCHAR(255) ,
				last_page VARCHAR(255),
				favorite Boolean  );
					
			`)

	if err != nil {
		panic(err)
	}
}
