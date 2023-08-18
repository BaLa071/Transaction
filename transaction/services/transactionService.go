package services

import (
	"context"
	"transaction/interfaces"
	"transaction/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	ProfileCollection *mongo.Collection
	ctx               context.Context
}

func InitTranasctionService(collection *mongo.Collection, ctx context.Context) interfaces.ITransaction {
	return &TransactionService{collection, ctx}
}

func (p *TransactionService) CreateTransaction(transaction *models.Transaction) (string, error) {
	_, err := p.ProfileCollection.InsertOne(p.ctx, &transaction)

	if err != nil {
		return "error", err
	}
	return "successfull", nil
}
func (p *TransactionService) GetTransactionById(user *models.Transaction) (*models.Transaction, error) {
	var newUser *models.Transaction
	return newUser, nil
}
