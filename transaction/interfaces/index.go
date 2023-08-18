package interfaces

import "transaction/models"

type ITransaction interface {
	CreateTransaction(transaction *models.Transaction) (string, error)
	GetTransactionById(transaction *models.Transaction) (*models.Transaction, error)
}
