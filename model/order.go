package model

import (
	"order-food-app-golang/server/request"
	responses "order-food-app-golang/server/response"
	"time"
)

// OrderModel ...
type OrderModel struct{}

// GetAll ...
func (model *OrderModel) GetAll(offset, limit int) (data []responses.GetOrder, count int, err error) {
	query := `SELECT o.id, i.id, i.name, u.id, u.name, o.status, o.created_at, o.updated_at
	FROM orders o
	JOIN items i ON o.item_id = i.id
	JOIN users u ON o.user_id = u.id
	limit ? offset ?`
	rows, err := GetDB().Query(query, limit, offset)
	if err != nil {
		return data, count, err
	}
	defer rows.Close()
	for rows.Next() {
		d := responses.GetOrder{}
		err = rows.Scan(&d.ID, &d.ItemId, &d.ItemName, &d.UserId, &d.UserName, &d.Status, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {

			return data, count, err
		}
		data = append(data, d)
	}
	// Query row count
	query = `SELECT COUNT(*) FROM orders`
	err = GetDB().QueryRow(query).Scan(&count)

	return data, count, err
}

// FindByID ...
func (model *OrderModel) FindByID(id float64) (data responses.GetOrder, err error) {
	query := `SELECT o.id, i.id, i.name, u.id, u.name, o.status, o.created_at, o.updated_at
	FROM orders o
	JOIN items i ON o.item_id = i.id
	JOIN users u ON o.user_id = u.id
	where o.id = ?`
	err = GetDB().QueryRow(query, id).Scan(&data.ID, &data.ItemId, &data.ItemName, &data.UserId,
		&data.UserName, &data.Status, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	return data, err
}

// Create ...
func (model *OrderModel) Create(body request.SetOrder) (err error) {
	tx, _ := GetDB().Begin()
	query, err := tx.Exec(`INSERT INTO orders(item_id, user_id, status)
	values(?,?,?)`, body.Item, body.User, body.Status)

	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = query.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

// Update ...
func (model *OrderModel) Update(id int, body request.SetOrder, changeAt time.Time) (err error) {
	query, err := db.Exec(`Update orders set item_id = ? , user_id = ?, status = ?
		updated_at = ? where id =? `,
		body.Item, body.User, body.Status, changeAt, id)

	if err != nil {
		return err
	}

	_, err = query.RowsAffected()
	if err != nil {
		return err
	}

	return err
}

func (model *OrderModel) UpdateStatus(id int, status string, changeAt time.Time) (err error) {
	query, err := db.Exec(`Update orders set status = ?, updated_at = ? where id =? `,
		status, changeAt, id)

	if err != nil {
		return err
	}

	_, err = query.RowsAffected()
	if err != nil {
		return err
	}

	return err
}
