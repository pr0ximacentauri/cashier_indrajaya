package usecases

import(
	"cashier_indrajaya/domain/entities"
	"cashier_indrajaya/domain/repositories"
	"time"
	"fmt"
)

type SaleService struct {
	repo        repositories.SaleRepository
	productRepo repositories.ProductRepository
}

func NewSaleService(repo repositories.SaleRepository, productRepo repositories.ProductRepository) *SaleService {
	return &SaleService{repo: repo, productRepo: productRepo}
}

func (s *SaleService) CreateSale(sale *entities.Sale) error {
	product, err := s.productRepo.GetByID(sale.ProductID)
	if err != nil {
		return fmt.Errorf("produk tidak ditemukan")
	}
	
	if product.Stock < sale.Quantity {
		return fmt.Errorf("stok tidak mencukupi. Stok tersedia: %d", product.Stock)
	}

	sale.Price = product.Price
	sale.Total = product.Price * float64(sale.Quantity)
	sale.SaleDate = time.Now()

	return s.repo.Create(sale)
}

func (s *SaleService) GetAllSales() ([]entities.Sale, error) {
	return s.repo.GetAll()
}