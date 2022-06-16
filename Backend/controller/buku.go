package controller

import (
	"database/sql"
	"final-project-engineering-56/Backend/model"
)

type Buku struct {
	db *sql.DB
}

func NewBuku(db *sql.DB) *Buku {
	return &Buku{db: db}
}

func (b *Buku) GetAllBuku() ([]model.Books, error) {
	var bukus []model.Books

	rows, err := b.db.Query(`SELECT * FROM books`)
	if err != nil {
		return bukus, err
	}

	for rows.Next() {
		var buku model.Books
		err := rows.Scan(&buku.IDbuku, &buku.ISBN, &buku.Judul, &buku.Penerbit, &buku.Pengarang, &buku.Tahun, &buku.Gambar, &buku.Deskripsi)
		if err != nil {
			return bukus, err
		}
		bukus = append(bukus, buku)
	}

	return bukus, nil
}

func (b *Buku) GetBukuByID(id int64) (model.Books, error) {
	var buku model.Books

	err := b.db.QueryRow(`SELECT * FROM books WHERE idbuku = ?`, id).Scan(&buku.IDbuku, &buku.ISBN, &buku.Judul, &buku.Penerbit, &buku.Pengarang, &buku.Tahun, &buku.Gambar, &buku.Deskripsi)
	if err != nil {
		return buku, err
	}

	return buku, nil
}

func (b *Buku) FetchBukuByJudul(Judul string) (model.Books, error) {
	var buku model.Books

	err := b.db.QueryRow(`SELECT * FROM books WHERE = ? `, Judul).Scan(&buku.IDbuku, &buku.ISBN, &buku.Judul, &buku.Penerbit, &buku.Pengarang, &buku.Tahun, &buku.Gambar, &buku.Deskripsi)
	if err != nil {
		return buku, err
	}

	return buku, nil
}

func (b *Buku) INSERTBuku(ISBN string, Judul string, Penerbit string, Pengarang string, Tahun string, Gambar string, Deskripsi string) error {
	_, err := b.db.Exec(`INSERT INTO books(ISBN, Judul, Penerbit, Pengarang, Tahun, Gambar, Deskripsi)
		VALUES
		(?, ?, ?, ?, ?, ?, ?)`, ISBN, Judul, Penerbit, Pengarang, Tahun, Gambar, Deskripsi)
	if err != nil {
		return err
	}

	return nil
}

func (b *Buku) UPDATEBuku(id int64, ISBN string, Judul string, Penerbit string, Pengarang string, Tahun string, Gambar string, Deskripsi string) error {
	_, err := b.db.Exec(`UPDATE books SET ISBN = ?, Judul = ?, Penerbit = ?, Pengarang = ?, Tahun = ?, Gambar = ?, Deskripsi = ? WHERE idbuku = ?`, ISBN, Judul, Penerbit, Pengarang, Tahun, Gambar, Deskripsi, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *Buku) DELETEBuku(id int64) error {
	_, err := b.db.Exec(`DELETE FROM books WHERE idbuku = ?`, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *Buku) GetBukuByPengarang(Pengarang string) ([]model.Books, error) {
	var bukus []model.Books

	rows, err := b.db.Query(`SELECT * FROM books WHERE Pengarang = ?`, Pengarang)
	if err != nil {
		return bukus, err
	}

	for rows.Next() {
		var buku model.Books
		err := rows.Scan(&buku.IDbuku, &buku.ISBN, &buku.Judul, &buku.Penerbit, &buku.Pengarang, &buku.Tahun, &buku.Gambar, &buku.Deskripsi)
		if err != nil {
			return bukus, err
		}
		bukus = append(bukus, buku)
	}

	return bukus, nil
}
