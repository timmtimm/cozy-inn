package transactions

import (
	"context"
	"cozy-inn/businesses/rooms"
	"cozy-inn/businesses/transactions"
	"errors"
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

func (tr *TransactionRepository) GetTransactionByRoomAndDate(roomType string, startDate time.Time, roomNumber int) ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("roomType", "==", roomType).Where("roomNumber", "==", roomNumber).Where("endDate", ">=", startDate).Where("status", "in", []string{"paid", "unpaid"})

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

	rec.StartDate = time.Date(rec.StartDate.Year(), rec.StartDate.Month(), rec.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	rec.EndDate = time.Date(rec.EndDate.Year(), rec.EndDate.Month(), rec.EndDate.Day(), 0, 0, 0, 0, time.UTC)

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

func (tr *TransactionRepository) GetTransactionByID(transactionID string) (transactions.Domain, error) {
	transactionDoc, err := tr.transactionsCollection().Doc(transactionID).Get(tr.ctx)
	if err != nil {
		return transactions.Domain{}, errors.New("transaction not available")
	}

	transaction := transactions.Domain{}

	if err := transactionDoc.DataTo(&transaction); err != nil {
		return transactions.Domain{}, err
	}

	return transaction, nil
}

func (tr *TransactionRepository) UpdatePayment(transactionID string, payment_URL string) (transactions.Domain, error) {
	transactionDoc := tr.transactionsCollection().Doc(transactionID)
	transaction, err := transactionDoc.Get(tr.ctx)
	if err != nil {
		return transactions.Domain{}, err
	}

	transactionData := Model{}
	if err := transaction.DataTo(&transactionData); err != nil {
		return transactions.Domain{}, err
	}

	transactionData.Payment_URL = payment_URL
	transactionData.Status = "verification-pending"
	transactionData.UpdatedAt = time.Now()

	_, err = transactionDoc.Set(tr.ctx, transactionData)
	if err != nil {
		return transactions.Domain{}, err
	}

	return transactionData.ToDomain(), nil
}

func (tr *TransactionRepository) GetPaymentNotVerified() ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("status", "==", "verification-pending")

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

func (tr *TransactionRepository) GetTransactionOnVerification(transactionID string) (transactions.Domain, error) {
	transactionDoc, err := tr.transactionsCollection().Doc(transactionID).Get(tr.ctx)
	if err != nil {
		return transactions.Domain{}, errors.New("transaction not available")
	}

	transaction := transactions.Domain{}
	if err := transactionDoc.DataTo(&transaction); err != nil {
		return transactions.Domain{}, err
	}

	return transaction, nil
}

func (tr *TransactionRepository) UpdateVerification(transactionID string, status string) (transactions.Domain, error) {
	transactionDoc := tr.transactionsCollection().Doc(transactionID)
	transaction, err := transactionDoc.Get(tr.ctx)
	if err != nil {
		return transactions.Domain{}, errors.New("transaction not available")
	}

	transactionData := Model{}
	if err := transaction.DataTo(&transactionData); err != nil {
		return transactions.Domain{}, err
	}

	transactionData.Status = status
	transactionData.UpdatedAt = time.Now()

	_, err = transactionDoc.Set(tr.ctx, transactionData)
	if err != nil {
		return transactions.Domain{}, err
	}

	return transactionData.ToDomain(), nil
}

func (tr *TransactionRepository) GetAllCheckIn() ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("status", "==", "verified")

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
