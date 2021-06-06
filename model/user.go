package model

import (
	"order-food-app-golang/server/request"
	responses "order-food-app-golang/server/response"
)

// UserModel ...
type UserModel struct{}

// GetAll ...
func (model *UserModel) GetAll(offset, limit int) (data []responses.UserModel, count int, err error) {
	query := `SELECT id, username,email,role_id FROM users limit ? offset ?`
	rows, err := GetDB().Query(query, limit, offset)
	if err != nil {
		return data, count, err
	}
	defer rows.Close()
	for rows.Next() {
		d := responses.UserModel{}
		err = rows.Scan(&d.ID, &d.Name, &d.Email, &d.RoleId)
		if err != nil {

			return data, count, err
		}
		data = append(data, d)
	}
	// Query row count
	query = `SELECT COUNT(*) FROM users`
	err = GetDB().QueryRow(query).Scan(&count)

	return data, count, err
}

// FindByUUID ...
func (model *UserModel) FindByID(id float64) (data responses.UserModel, err error) {
	query := `SELECT id, username, email FROM users where id = ?`
	err = GetDB().QueryRow(query, id).Scan(&data.ID, &data.Name, &data.Email)
	if err != nil {

		return data, err
	}
	return data, err
}

// FindByEmail ...
func (model *UserModel) FindByEmail(email string) (data responses.UserModel, err error) {
	query := `SELECT id, username, email, password, role_id FROM users where email = ?`
	err = GetDB().QueryRow(query, email).Scan(&data.ID, &data.Name, &data.Email, &data.Password, &data.RoleId)

	if err != nil {

		return data, err
	}
	return data, err
}

// Create ...
func (model *UserModel) Create(body request.Register) (res int64, err error) {
	tx, _ := GetDB().Begin()
	query, err := tx.Exec(`INSERT INTO users(username, email, password)
	values(?,?,?)`, body.Name, body.Email, body.Password)

	if err != nil {
		tx.Rollback()

		return res, err
	}
	res, err = query.LastInsertId()
	if err != nil {
		tx.Rollback()

		return res, err
	}
	tx.Commit()
	return res, err
}
