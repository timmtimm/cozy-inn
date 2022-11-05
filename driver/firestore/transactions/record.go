package transactions

import (
	"cozy-inn/businesses/transactions"
	"time"
)

type Model struct {
	TransactionID string    `firestore:"transactionID"`
	UserEmail     string    `firestore:"userEmail"`
	RoomType      string    `firestore:"roomType"`
	RoomNumber    int       `firestore:"roomNumber"`
	StartDate     time.Time `firestore:"startDate"`
	EndDate       time.Time `firestore:"endDate"`
	CheckIn       time.Time `firestore:"checkIn,omitempty"`
	CheckOut      time.Time `firestore:"checkOut,omitempty"`
	Status        string    `firestore:"status"`
	Bill          int       `firestore:"bill"`
	Payment_URL   string    `firestore:"payment_URL,omitempty"`
	CreatedAt     time.Time `firestore:"createdAt"`
	UpdatedAt     time.Time `firestore:"updatedAt"`
}

func FromDomain(domain *transactions.Domain) *Model {
	return &Model{
		TransactionID: domain.TransactionID,
		RoomType:      domain.RoomType,
		RoomNumber:    domain.RoomNumber,
		StartDate:     domain.StartDate,
		EndDate:       domain.EndDate,
		CheckIn:       domain.CheckIn,
		CheckOut:      domain.CheckOut,
		Status:        domain.Status,
		Bill:          domain.Bill,
		Payment_URL:   domain.Payment_URL,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func (rec *Model) ToDomain() transactions.Domain {
	return transactions.Domain{
		TransactionID: rec.TransactionID,
		UserEmail:     rec.UserEmail,
		RoomType:      rec.RoomType,
		RoomNumber:    rec.RoomNumber,
		StartDate:     rec.StartDate,
		EndDate:       rec.EndDate,
		CheckIn:       rec.CheckIn,
		CheckOut:      rec.CheckOut,
		Status:        rec.Status,
		Bill:          rec.Bill,
		Payment_URL:   rec.Payment_URL,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}
