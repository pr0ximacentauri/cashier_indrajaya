package repositories

import (
	"cashier_indrajaya/domain/entities"
	"database/sql"
	"time"
)

type SQLiteSaleRepository struct {
	db *sql.DB
}

func NewSQLiteSaleRepository(db *sql.DB) *SQLiteSaleRepository {
	return &SQLiteSaleRepository{db: db}
}

func (r *SQLiteSaleRepository) GetAll() ([]entities.Sale, error) {
	rows, err := r.db.Query(`
		SELECT s.id, s.product_id, p.name, s.quantity, s.price, s.total, s.sale_date 
		FROM sales s 
		JOIN products p ON s.product_id = p.id 
		ORDER BY s.sale_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []entities.Sale
	for rows.Next() {
		var s entities.Sale
		err := rows.Scan(&s.ID, &s.ProductID, &s.ProductName, &s.Quantity, &s.Price, &s.Total, &s.SaleDate)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}
	return sales, nil
}

func (r *SQLiteSaleRepository) Create(sale *entities.Sale) error {
	
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO sales (product_id, quantity, price, total, sale_date) VALUES (?, ?, ?, ?, ?)",
		sale.ProductID, sale.Quantity, sale.Price, sale.Total, sale.SaleDate)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	sale.ID = int(id)

	_, err = tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", sale.Quantity, sale.ProductID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *SQLiteSaleRepository) GetByDateRange(start, end time.Time) ([]entities.Sale, error) {
	rows, err := r.db.Query(`
		SELECT s.id, s.product_id, p.name, s.quantity, s.price, s.total, s.sale_date 
		FROM sales s 
		JOIN products p ON s.product_id = p.id 
		WHERE s.sale_date BETWEEN ? AND ?
		ORDER BY s.sale_date DESC
	`, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []entities.Sale
	for rows.Next() {
		var s entities.Sale
		err := rows.Scan(&s.ID, &s.ProductID, &s.ProductName, &s.Quantity, &s.Price, &s.Total, &s.SaleDate)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}
	return sales, nil
}
