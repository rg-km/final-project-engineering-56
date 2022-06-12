package controller

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"final-project-engineering-56/Backend/model"
)

type UserLogin struct {
	db *sql.DB
}

func NewUserLogin(db *sql.DB) *UserLogin {
	return &UserLogin{db: db}
}

func (u *UserLogin) Login(username string, password string) *string {
	var user model.Users

	err := u.db.QueryRow(`
		SELECT * FROM users WHERE email = ? AND password = ?`, username, password).Scan(&user.Username)
	if err != nil {
		return nil
	}

	return &user.Username
}

func (u *UserLogin) Register(username string, password string, email string) (bool, error) {
	var user model.Users

	err := u.db.QueryRow(`
		INSERT INTO users(username, password, email) VALUES
			(?, ?, ?)`, username, password, email).Scan(&user.Username, &user.Password, &user.Email)
	if err != nil {
		return false, err
	}

	return true, nil
	// err := u.db.Exec(`
	// 	INSERT INTO users (username, password, email) VALUES (?, ?, ?)`, username, password, email)
	// if err != nil {
	// 	return err
	// }

	// return nil
}

// func (u *UserLogin) FetchuserEmail(email string) (*string, error) {
// 	var user model.Users

// 	err := u.db.QueryRow(`
// 		SELECT * FROM users WHERE email = ?`, email).Scan(&user.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user.Email, nil
// }
