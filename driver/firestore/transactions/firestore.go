package transactions

import (
	"context"
	"cozy-inn/businesses/rooms"
	"cozy-inn/businesses/transactions"
	"errors"
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

func (tr *TransactionRepository) GetAllTransaction() ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection()

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

func (tr *TransactionRepository) GetAllTransactionByEmail(email string) ([]transactions.Domain, error) {
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
		if err := doc.DataTo(&transaction); err != nil {
			return []transactions.Domain{}, err
		}

		transactionList = append(transactionList, transaction)
	}

	return transactionList, nil
}

func (tr *TransactionRepository) GetAllReadyCheckIn() ([]transactions.Domain, error) {
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

func (tr *TransactionRepository) GetAllReadyCheckOut() ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("status", "==", "checked-in")

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

func (tr *TransactionRepository) GetAllPaymentNotVerified() ([]transactions.Domain, error) {
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

func (tr *TransactionRepository) GetTransactionByRoomAndEndDate(roomType string, startDate time.Time, roomNumber int) ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("roomType", "==", roomType).Where("roomNumber", "==", roomNumber).Where(
		"endDate", ">=", startDate).Where("status", "in", []string{"unpaid", "verification-pending", "checked-in", "verified"})

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

func (tr *TransactionRepository) GetTransactionByID(transactionID string) (transactions.Domain, error) {
	transactionDoc, err := tr.transactionsCollection().Doc(transactionID).Get(tr.ctx)
	if err != nil {
		return transactions.Domain{}, errors.New("transaction not found")
	}

	transaction := transactions.Domain{}
	if err := transactionDoc.DataTo(&transaction); err != nil {
		return transactions.Domain{}, err
	}

	return transaction, nil
}

func (tr *TransactionRepository) GetTransactionOngoing() ([]transactions.Domain, error) {
	transactionList := []transactions.Domain{}
	transactionDoc := tr.transactionsCollection().Where("status", "not-in", []string{"done, rejected, cancelled"})

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

func (tr *TransactionRepository) Create(email string, transactionInput transactions.Domain, RoomData rooms.Domain) (transactions.Domain, error) {
	rec := FromDomain(transactionInput)

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

func (tr *TransactionRepository) Update(transcationID string, transactionInput transactions.Domain) error {
	rec := FromDomain(transactionInput)

	_, err := tr.transactionsCollection().Doc(transcationID).Set(tr.ctx, rec)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TransactionRepository) Delete(transactionID string) error {
	_, err := tr.transactionsCollection().Doc(transactionID).Delete(tr.ctx)
	if err != nil {
		return err
	}

	return nil
}
