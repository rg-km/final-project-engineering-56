package controller

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"final-project-engineering-56/Backend/middleware"
	"final-project-engineering-56/Backend/model"
)

type UserLogin struct {
	db *sql.DB
}

func NewUserLogin(db *sql.DB) *UserLogin {
	return &UserLogin{db: db}
}

func (u *UserLogin) Login(email string, password string) *string {
	var user model.Users

	err := u.db.QueryRow(`
		SELECT email FROM users WHERE email = ? AND password = ?`, email, password).Scan(&user.Email)
	if err != nil {
		return nil
	}

	match := middleware.Verifypassword(password, user.Password)
	if match == false {
		return nil
	}

	return &user.Email
}

func (u *UserLogin) Register(username string, password string, email string) error {

	p := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"

	_, err := u.db.Exec(p, username, password, email)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserLogin) GetALLUser() ([]model.Users, error) {
	rows, err := u.db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserLogin) Update(id int, username string, password string, email string) error {
	_, err := u.db.Exec(`
		UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?`, username, password, email, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserLogin) Delete(id int) error {
	_, err := u.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserLogin) GetbyID(id int) (model.Users, error) {
	var user model.Users

	err := u.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}
