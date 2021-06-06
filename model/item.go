package model

import (
	"order-food-app-golang/server/request"
	responses "order-food-app-golang/server/response"
	"time"
)

// ItemModel ...
type ItemModel struct{}

// GetAll ...
func (model *ItemModel) GetAll(offset, limit int) (data []responses.GetItem, count int, err error) {
	query := `SELECT f.id, c.name, f.name, f.dsc, f.price, f.stock, f.status
	FROM items f
	JOIN categories c ON f.category_id = c.id
	limit ? offset ?`
	rows, err := GetDB().Query(query, limit, offset)
	if err != nil {
		return data, count, err
	}
	defer rows.Close()
	for rows.Next() {
		d := responses.GetItem{}
		err = rows.Scan(&d.ID, &d.Category, &d.Name, &d.Desc, &d.Price, &d.Stock, &d.Status)
		if err != nil {

			return data, count, err
		}
		data = append(data, d)
	}
	// Query row count
	query = `SELECT COUNT(*) FROM items`
	err = GetDB().QueryRow(query).Scan(&count)

	return data, count, err
}

// FindByID ...
func (model *ItemModel) FindByID(id float64) (data responses.GetItem, err error) {
	query := `SELECT f.id, c.name, f.name, f.dsc, f.price, f.stock, f.status
	FROM items f
	JOIN categories c ON f.category_id = c.id
	where f.id = ?`
	err = GetDB().QueryRow(query, id).Scan(&data.ID, &data.Category, &data.Name, &data.Desc,
		&data.Price, &data.Stock, &data.Status)
	if err != nil {
		return data, err
	}
	return data, err
}

// Create ...
func (model *ItemModel) Create(body request.SetItem) (err error) {
	tx, _ := GetDB().Begin()
	query, err := tx.Exec(`INSERT INTO items(category_id, name, dsc, price, stock, status)
	values(?,?,?,?,?,?)`, body.Category, body.Name, body.Desc, body.Price, body.Stock, body.Status)

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
func (model *ItemModel) Update(id int, body request.SetItem, changeAt time.Time) (err error) {
	query, err := db.Exec(`Update items set category_id = ? , name = ?, dsc = ?, price = ?,
		stock = ? status = ?, updated_at = ? where id =? `,
		body.Category, body.Name, body.Desc, body.Price, body.Stock, body.Status, changeAt, id)

	if err != nil {
		return err
	}

	_, err = query.RowsAffected()
	if err != nil {
		return err
	}

	return err
}

func (model *ItemModel) UpdateStatus(id int, status string, changeAt time.Time) (err error) {
	query, err := db.Exec(`Update items set status = ?, updated_at = ? where id =? `,
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
