package controller

import (
	"database/sql"
	"final-project-engineering-56/Backend/model"
)

type Admin struct {
	db *sql.DB
}

func NewAdmin(db *sql.DB) *Admin {
	return &Admin{db: db}
}

func (a *Admin) Login(username string, password string) *string {
	var admin model.Admin

	err := a.db.QueryRow(`
		SELECT * FROM admin WHERE username = ? AND password = ?`, username, password).Scan(&admin.ID, &admin.Username, &admin.Password, &admin.Role)
	if err != nil {
		return nil
	}

	return &admin.Username
}

func (a *Admin) Register(username string, password string, role string) error {

	p := "INSERT INTO admin (username, password, role) VALUES (?, ?, ?)"

	_, err := a.db.Exec(p, username, password, role)
	if err != nil {
		return err
	}

	return nil
}

func (a *Admin) GetALLAdmin() ([]model.Admin, error) {
	rows, err := a.db.Query(`SELECT * FROM admin`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var admins []model.Admin
	for rows.Next() {
		var admin model.Admin
		err := rows.Scan(&admin.ID, &admin.Username, &admin.Password, &admin.Role)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil

}

func (a *Admin) GetAdminByID(id int) (model.Admin, error) {
	var admin model.Admin

	err := a.db.QueryRow(`
		SELECT * FROM admin WHERE id = ?`, id).Scan(&admin.ID, &admin.Username, &admin.Password, &admin.Role)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (a *Admin) Update(id int, username string, password string, role string) error {
	_, err := a.db.Exec(`
		UPDATE admin SET username = ?, password = ?, role = ? WHERE id = ?`, username, password, role, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *Admin) Delete(id int) error {
	_, err := a.db.Exec(`DELETE FROM admin WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
