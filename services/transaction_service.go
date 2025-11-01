package services

import (
	"cashier_indrajaya/models"
	"cashier_indrajaya/repositories"
)

type TransactionService interface {
	Create(transaction *models.Transaction) error
	GetAll() ([]models.Transaction, error)
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

func (s *transactionService) Create(tx *models.Transaction) error {
	return s.repo.Create(tx)
}

func (s *transactionService) GetAll() ([]models.Transaction, error) {
	return s.repo.GetAll()
}
