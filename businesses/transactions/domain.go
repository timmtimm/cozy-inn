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
	PaymentProof  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetAllTransaction(email string) ([]Domain, error)
	CreateTransaction(email string, transactionDomain *Domain) (Domain, error)
}

type Repository interface {
	GetAllTransaction(email string) ([]Domain, error)
	CreateTransaction(email string, transactionDomain *Domain, RoomData rooms.Domain) (Domain, error)
	GetFinishedTransactionByRoom(roomType string, startDate time.Time, roomNumber int) ([]Domain, error)
}
