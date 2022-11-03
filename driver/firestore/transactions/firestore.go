package transactions

import (
	"context"
	"cozy-inn/businesses/rooms"
	"cozy-inn/businesses/transactions"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type TransactionRepository struct {
	client         *firestore.Client
	ctx            context.Context
	RoomRepository rooms.Repository
}

func NewTransactionRepository(client *firestore.Client, ctx context.Context, RoomRepository rooms.Repository) transactions.Repository {
	if client == nil {
		panic("No firestore client")
	}
	return &TransactionRepository{client, ctx, RoomRepository}
}

func (tr *TransactionRepository) transactionsCollection() *firestore.CollectionRef {
	return tr.client.Collection("transactions")
}

func (tr *TransactionRepository) GetAllTransaction(email string) ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("userEmail", "==", email)

	iter := transactionDoc.Documents(tr.ctx)
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

func (tr *TransactionRepository) GetFinishedTransactionByRoom(roomType string, startDate time.Time) ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("roomType", "==", roomType).Where("endDate", "<=", startDate).Where("status", "==", "finished")

	iter := transactionDoc.Documents(tr.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []transactions.Domain{}, err
		}

		transaction := transactions.Domain{}
		if err := doc.DataTo(&transaction); err != nil {
			return []transactions.Domain{}, err
		}

		transactionList = append(transactionList, transaction)
	}

	return transactionList, nil
}

func (tr *TransactionRepository) CreateTransaction(email string, transactionDomain *transactions.Domain, RoomData rooms.Domain) (transactions.Domain, error) {
	rec := FromDomain(transactionDomain)

	rec.UserEmail = email
	rec.CreatedAt = time.Now()
	rec.UpdatedAt = rec.CreatedAt
	rec.Bill = RoomData.Price * int(rec.EndDate.Sub(rec.StartDate).Hours()/24)
	rec.Status = "unpaid"

	timeToString, _ := rec.CreatedAt.UTC().MarshalText()
	rec.TransactionID = fmt.Sprintf("%s_%s", timeToString, email)

	_, err := tr.transactionsCollection().Doc(rec.TransactionID).Set(tr.ctx, Model{
		TransactionID: rec.TransactionID,
		UserEmail:     rec.UserEmail,
		RoomType:      rec.RoomType,
		RoomNumber:    rec.RoomNumber,
		StartDate:     rec.StartDate,
		EndDate:       rec.EndDate,
		Status:        rec.Status,
		Bill:          rec.Bill,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	})
	if err != nil {
		return transactions.Domain{}, err
	}

	return rec.ToDomain(), nil
}
