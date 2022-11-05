package transactions

import (
	"cozy-inn/businesses/rooms"
	"time"
)

type Domain struct {
	TransactionID string
	UserEmail     string
	RoomType      string
	RoomNumber    int
	StartDate     time.Time
	EndDate       time.Time
	CheckIn       time.Time
	CheckOut      time.Time
	Status        string
	Bill          int
	Payment_URL   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetAllTransaction(email string) ([]Domain, error)
	CreateTransaction(email string, transactionDomain *Domain) (Domain, error)
	UpdatePayment(transactionID string, payment_URL string) (Domain, error)
	GetPaymentNotVerified() ([]Domain, error)
	GetTransactionOnVerification(transactionID string) (Domain, error)
	UpdateVerification(transactionID string, status string) (Domain, error)
	GetAllCheckIn() ([]Domain, error)
	GetCheckInTransaction(transactionID string) (Domain, error)
	UpdateCheckIn(transactionID string) (Domain, error)
	GetAllCheckOut() ([]Domain, error)
	GetCheckOutTransaction(transactionID string) (Domain, error)
	UpdateCheckOut(transactionID string) (Domain, error)
	DeleteTransaction(transactionID string) error
}

type Repository interface {
	GetAllTransaction(email string) ([]Domain, error)
	CreateTransaction(email string, transactionDomain *Domain, RoomData rooms.Domain) (Domain, error)
	GetTransactionByRoomAndDate(roomType string, startDate time.Time, roomNumber int) ([]Domain, error)
	GetTransactionByID(transactionID string) (Domain, error)
	UpdatePayment(transactionID string, payment_URL string) (Domain, error)
	GetPaymentNotVerified() ([]Domain, error)
	GetTransactionOnVerification(transactionID string) (Domain, error)
	UpdateVerification(transactionID string, status string) (Domain, error)
	GetAllCheckIn() ([]Domain, error)
	GetCheckInTransaction(transactionID string) (Domain, error)
	UpdateCheckIn(transactionID string) (Domain, error)
	GetAllCheckOut() ([]Domain, error)
	Update(transcationID string, transactionDomain Domain) error
	Delete(transactionID string) error
}
