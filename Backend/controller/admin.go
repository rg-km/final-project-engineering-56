package controller

import (
	"database/sql"
)

type Admin struct {
	db *sql.DB
}

func NewAdmin(db *sql.DB) *Admin {
	return &Admin{db: db}
}
