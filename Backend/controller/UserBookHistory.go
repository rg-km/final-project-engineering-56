package controller

import (
	"database/sql"
	"final-project-engineering-56/Backend/model"
)

type UserBookHistory struct {
	db *sql.DB
}

func NewUserBookHistory(db *sql.DB) *UserBookHistory {
	return &UserBookHistory{db: db}
}

func (h *UserBookHistory) HistoryBuku(id_user string, id_books string, ) error {

	sqlStament := `
	INSERT INTO user_book_history (id_user, id_books, true) VALUES (?, ?, ?)
	`

	_, err := h.db.Exec(sqlStament, id_user, id_books)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserBookHistory) UpdateHistoryBuku(id int, id_user string, id_books string, favorite bool) error {

	_, err := h.db.Exec(`
		UPDATE user_book_history SET id_user= ?, id_books = ?, favorite = ? WHERE id = ?`, id_user, id_books, favorite, id)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserBookHistory) SetFavorit(id int, id_user string, id_books string, favorite bool) error {

	_, err := h.db.Exec(`
		UPDATE user_book_history SET id_user= ?, id_books = ?, favorite = ? WHERE id = ?`, id_user, id_books, favorite, id)
	if err != nil {
		return err
	}

	return nil
}


func (h *UserBookHistory) GetALLFav() ([]model.HistoryBuku, error) {
	rows, err := h.db.Query(`SELECT * FROM user_book_history`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var HistoryBuku []model.HistoryBuku
	for rows.Next() {
		var getbuku model.HistoryBuku
		err := rows.Scan(&getbuku.ID, &getbuku.ID_user, &getbuku.ID_books, &getbuku.Favorite)
		if err != nil {
			return nil, err
		}
		HistoryBuku = append(HistoryBuku, getbuku)
	}

	return HistoryBuku, nil
}