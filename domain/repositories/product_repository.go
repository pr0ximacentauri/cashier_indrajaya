package repositories

import "cashier_indrajaya/domain/entities"

type ProductRepository interface {
	GetAll() ([]entities.Product, error)
	GetByID(id int) (*entities.Product, error)
	Create(product *entities.Product) error
	Update(product *entities.Product) error
	Delete(id int) error
}