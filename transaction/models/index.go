package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Amount          float32 `json:"amount" bson:"amount"`
	TransactionDate string  `json:"transactionDate" bson:"transactionDate"`
	// TransactionType []TransactionType    `json:"transactiontype" bson:"transactionType"`
	CustomerId string    `json:"transactionId" bson:"transactionId"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type DBResponse struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Amount          float32            `json:"amount" bson:"amount"`
	TransactionDate time.Time          `json:"transactionDate" bson:"transactionDate"`
	// TransactionType []TransactionType    `json:"transactiontype" bson:"transactionType"`
	CustomerId string    `json:"transactionId" bson:"transactionId"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}
