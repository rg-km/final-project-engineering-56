package model

type Users struct {
	ID       int    `db:"id"; json:"id"`
	Username string `db:"username"; json:"username"`
	Password string `db:"password";json:"password"`
	Email    string `db:"email";json:"email"`
}

type Books struct {
	IDbuku    int    `db:"idbuku";json:"idbuku"`
	ISBN      string `db:"ISBN";json:"ISBN"`
	Judul     string `db:"judul";json:"judul"`
	Penerbit  string `db:"penerbit";json:"penerbit"`
	Pengarang string `db:"pengarang";json:"pengarang"`
	Tahun     string `db:"tahun";json:"tahun"`
	Gambar    string `db:"gambar";json:"gambar"`
	Deskripsi string `db:"deskripsi";json:"deskripsi"`
}

type Admin struct {
	ID       int    `db:"idadmin";json:"idadmin"`
	Username string `db:"username";json:"username"`
	Password string `db:"password";json:"password"`
	Role     string `db:"role";json:"role"`
}
