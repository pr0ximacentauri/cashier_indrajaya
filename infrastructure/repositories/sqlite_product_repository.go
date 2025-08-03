package repositories

import (
	"cashier_indrajaya/domain/entities"
	"database/sql"
)

type SQLiteProductRepository struct {
	db *sql.DB
}

func NewSQLiteProductRepository(db *sql.DB) *SQLiteProductRepository {
	return &SQLiteProductRepository{db: db}
}

func (r *SQLiteProductRepository) GetAll() ([]entities.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, stock, category FROM products ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var p entities.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *SQLiteProductRepository) GetByID(id int) (*entities.Product, error) {
	var p entities.Product
	err := r.db.QueryRow("SELECT id, name, price, stock, category FROM products WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *SQLiteProductRepository) Create(product *entities.Product) error {
	result, err := r.db.Exec("INSERT INTO products (name, price, stock, category) VALUES (?, ?, ?, ?)",
		product.Name, product.Price, product.Stock, product.Category)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	product.ID = int(id)
	return nil
}

func (r *SQLiteProductRepository) Update(product *entities.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = ?, price = ?, stock = ?, category = ? WHERE id = ?",
		product.Name, product.Price, product.Stock, product.Category, product.ID)
	return err
}

func (r *SQLiteProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}