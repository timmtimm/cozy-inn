package transactions

import (
	"context"
	"cozy-inn/businesses/transactions"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type TransactionRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewTransactionRepository(client *firestore.Client, ctx context.Context) transactions.Repository {
	if client == nil {
		panic("No firestore client")
	}
	return &TransactionRepository{client, ctx}
}

func (rr *TransactionRepository) transctionsCollection() *firestore.CollectionRef {
	return rr.client.Collection("transctions")
}

func (rr *TransactionRepository) GetAllTransaction(email string) ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := rr.transctionsCollection().Where("email", "==", email)

	iter := transactionDoc.Documents(rr.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []transactions.Domain{}, err
		}

		transaction := transactions.Domain{}
		doc.DataTo(&transaction)
		transactionList = append(transactionList, transaction)
	}

	return transactionList, nil
}
