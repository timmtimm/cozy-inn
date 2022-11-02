package transactions

import (
	"cozy-inn/businesses/transactions"
	"time"
)

type Model struct {
	TransactionID string    `firestore:"transactionID"`
	Room          string    `firestore:"room"`
	StartDate     time.Time `firestore:"startDate"`
	EndDate       time.Time `firestore:"endDate"`
	CheckIn       time.Time `firestore:"checkIn"`
	CheckOut      time.Time `firestore:"checkOut"`
	Status        string    `firestore:"status"`
	Price         int       `firestore:"price"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
}

func FromDomain(domain transactions.Domain) *Model {
	return &Model{
		TransactionID: domain.TransactionID,
		Room:          domain.Room,
		StartDate:     domain.StartDate,
		EndDate:       domain.EndDate,
		CheckIn:       domain.CheckIn,
		CheckOut:      domain.CheckOut,
		Status:        domain.Status,
		Price:         domain.Price,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func (rec *Model) ToDomain() transactions.Domain {
	return transactions.Domain{
		TransactionID: rec.TransactionID,
		Room:          rec.Room,
		StartDate:     rec.StartDate,
		EndDate:       rec.EndDate,
		CheckIn:       rec.CheckIn,
		CheckOut:      rec.CheckOut,
		Status:        rec.Status,
		Price:         rec.Price,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}
