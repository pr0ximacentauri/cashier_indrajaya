package services

import (
	"cashier_indrajaya/models"
	"cashier_indrajaya/repositories"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) error
	Update(id string, product *models.Product) error
	Delete(id string) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) Update(id string, product *models.Product) error {
	return s.repo.Update(id, product)
}

func (s *productService) Delete(id string) error {
	return s.repo.Delete(id)
}
