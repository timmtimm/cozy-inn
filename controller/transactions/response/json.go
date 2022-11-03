package response

import (
	"cozy-inn/businesses/transactions"
	"time"
)

type Transaction struct {
	TransactionID string    `json:"transactionID" validate:"required" firestore:"transactionID"`
	UserEmail     string    `json:"userEmail" validate:"required" firestore:"userEmail"`
	RoomType      string    `json:"roomType" validate:"required" firestore:"roomType"`
	RoomNumber    int       `json:"roomNumber" validate:"required" firestore:"roomNumber"`
	StartDate     time.Time `json:"startDate" validate:"required" firestore:"startDate"`
	EndDate       time.Time `json:"EndDate" validate:"required" firestore:"EndDate"`
	CheckIn       time.Time `json:"checkIn" firestore:"checkIn,omitempty"`
	CheckOut      time.Time `json:"checkOut" firestore:"checkOut,omitempty"`
	Status        string    `json:"status" validate:"required" firestore:"status"`
	Bill          int       `json:"bill" validate:"required" firestore:"bill"`
	PaymentProof  string    `json:"paymentProof" firestore:"paymentProof,omitempty"`
	CreatedAt     time.Time `json:"createdAt" firestore:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" firestore:"updatedAt"`
}

func ToDomain(domain transactions.Domain) Transaction {
	return Transaction{
		TransactionID: domain.TransactionID,
		UserEmail:     domain.UserEmail,
		RoomType:      domain.RoomType,
		RoomNumber:    domain.RoomNumber,
		StartDate:     domain.StartDate,
		EndDate:       domain.EndDate,
		CheckIn:       domain.CheckIn,
		CheckOut:      domain.CheckOut,
		Status:        domain.Status,
		Bill:          domain.Bill,
		PaymentProof:  domain.PaymentProof,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
