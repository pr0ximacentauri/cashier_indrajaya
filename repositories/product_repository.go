package repositories

import (
	"cashier_indrajaya/config"
	"cashier_indrajaya/models"
	"context"
	"time"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) error
	Update(id string, product *models.Product) error
	Delete(id string) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	iter := config.Firestore.Collection("products").Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var product models.Product
		doc.DataTo(&product)
		product.ID = doc.Ref.ID
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) Create(product *models.Product) error {
	product.CreatedAt = time.Now()
	_, _, err := config.Firestore.Collection("products").Add(context.Background(), map[string]interface{}{
		"name":      product.Name,
		"category":  product.Category,
		"price":     product.Price,
		"stock":     product.Stock,
		"createdAt": product.CreatedAt,
	})
	return err
}

func (r *productRepository) Update(id string, product *models.Product) error {
	_, err := config.Firestore.Collection("products").Doc(id).Set(context.Background(), map[string]interface{}{
		"name":     product.Name,
		"category": product.Category,
		"price":    product.Price,
		"stock":    product.Stock,
	}, nil)
	return err
}

func (r *productRepository) Delete(id string) error {
	_, err := config.Firestore.Collection("products").Doc(id).Delete(context.Background())
	return err
}