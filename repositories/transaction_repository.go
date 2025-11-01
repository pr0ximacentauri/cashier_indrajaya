package repositories

import (
	"context"
	"time"

	"cashier_indrajaya/config"
	"cashier_indrajaya/models"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetAll() ([]models.Transaction, error)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(tx *models.Transaction) error {
	tx.CreatedAt = time.Now()
	_, _, err := config.Firestore.Collection("transactions").Add(context.Background(), map[string]interface{}{
		"items":     tx.Items,
		"total":     tx.Total,
		"createdAt": tx.CreatedAt,
	})
	return err
}

func (r *transactionRepository) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	iter := config.Firestore.Collection("transactions").Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var tx models.Transaction
		doc.DataTo(&tx)
		tx.ID = doc.Ref.ID
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
