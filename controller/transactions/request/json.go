package request

import (
	"cozy-inn/businesses/transactions"
	"time"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	TransactionID string    `json:"transactionID" firestore:"transactionID"`
	UserEmail     string    `json:"userEmail" firestore:"userEmail"`
	RoomType      string    `json:"roomType" validate:"required" firestore:"roomType"`
	RoomNumber    int       `json:"roomNumber" validate:"required" firestore:"roomNumber"`
	StartDate     time.Time `json:"startDate" validate:"required" firestore:"startDate"`
	EndDate       time.Time `json:"EndDate" validate:"required" firestore:"EndDate"`
	CheckIn       time.Time `json:"checkIn" firestore:"checkIn,omitempty"`
	CheckOut      time.Time `json:"checkOut" firestore:"checkOut,omitempty"`
	Status        string    `json:"status" firestore:"status"`
	Bill          int       `json:"bill" firestore:"bill"`
	PaymentProof  string    `json:"paymentProof" firestore:"paymentProof,omitempty"`
}

func (req *Transaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		TransactionID: req.TransactionID,
		UserEmail:     req.UserEmail,
		RoomType:      req.RoomType,
		RoomNumber:    req.RoomNumber,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		CheckIn:       req.CheckIn,
		Status:        req.Status,
		Bill:          req.Bill,
		PaymentProof:  req.PaymentProof,
	}
}

func (req *Transaction) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
