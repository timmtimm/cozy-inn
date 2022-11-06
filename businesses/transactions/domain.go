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
	GetAllTransactionUser(email string) ([]Domain, error)
	GetAllReadyCheckIn() ([]Domain, error)
	GetAllReadyCheckOut() ([]Domain, error)
	GetAllPaymentNotVerified() ([]Domain, error)
	AdminGetAllTransaction() ([]Domain, error)
	GetTransactionOnVerification(transactionID string) (Domain, error)
	GetCheckInTransaction(transactionID string) (Domain, error)
	GetCheckOutTransaction(transactionID string) (Domain, error)
	GetTransaction(transactionID string) (Domain, error)
	CheckAvailabilityAllRoom(startDate time.Time, endDate time.Time) ([]string, error)
	CreateTransaction(email string, transactionInput Domain) (Domain, error)
	UpdatePayment(transactionID string, email string, payment_URL string) (Domain, error)
	UpdateVerification(transactionID string, status string) (Domain, error)
	UpdateCheckIn(transactionID string) (Domain, error)
	UpdateCheckOut(transactionID string) (Domain, error)
	AdminUpdateTransaction(transactionID string, userInput Domain) (Domain, error)
	AdminDeleteTransaction(transactionID string) error
}

type Repository interface {
	GetAllTransaction() ([]Domain, error)
	GetAllTransactionByEmail(email string) ([]Domain, error)
	GetAllReadyCheckIn() ([]Domain, error)
	GetAllReadyCheckOut() ([]Domain, error)
	GetAllPaymentNotVerified() ([]Domain, error)
	GetTransactionByRoomAndEndDate(roomType string, startDate time.Time, roomNumber int) ([]Domain, error)
	GetTransactionByID(transactionID string) (Domain, error)
	GetTransactionOngoing() ([]Domain, error)
	Create(email string, transactionInput Domain, RoomData rooms.Domain) (Domain, error)
	Update(transcationID string, transactionDomain Domain) error
	Delete(transactionID string) error
}
