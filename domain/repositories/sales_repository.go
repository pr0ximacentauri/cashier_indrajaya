package repositories

import(
	"cashier_indrajaya/domain/entities"
	"time"
)

type SaleRepository interface {
	GetAll() ([]entities.Sale, error)
	Create(sale *entities.Sale) error
	GetByDateRange(start, end time.Time) ([]entities.Sale, error)
}