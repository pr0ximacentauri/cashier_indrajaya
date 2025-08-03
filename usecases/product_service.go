package usecases

import (
	"fmt"
	"cashier_indrajaya/domain/entities"
	"cashier_indrajaya/domain/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]entities.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductByID(id int) (*entities.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) CreateProduct(product *entities.Product) error {
	if product.Name == "" {
		return fmt.Errorf("nama produk tidak boleh kosong")
	}
	if product.Price <= 0 {
		return fmt.Errorf("harga harus lebih dari 0")
	}
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *entities.Product) error {
	if product.Name == "" {
		return fmt.Errorf("nama produk tidak boleh kosong")
	}
	if product.Price <= 0 {
		return fmt.Errorf("harga harus lebih dari 0")
	}
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}