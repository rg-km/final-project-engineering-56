package controller

import (
	"database/sql"
)

type UserBookHistory struct {
	db *sql.DB
}

func NewUserBookHistory(db *sql.DB) *UserBookHistory {
	return &UserBookHistory{db: db}
}


func (h *UserBookHistory) HistoryBuku(iduser string , idbooks string , lastpage string )  {


}