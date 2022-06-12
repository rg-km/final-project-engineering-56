package model

type Users struct {
	ID       int    `db:"id"; json:"id"`
	Username string `db:"username"; json:"username"`
	Password string `db:"password";json:"password"`
	Email    string `db:"email";json:"email"`
	Token    string `db"token";json:"token"`
}
