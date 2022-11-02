package transactions

import "time"

type Domain struct {
	TransactionID string
	Room          string
	StartDate     time.Time
	EndDate       time.Time
	CheckIn       time.Time
	CheckOut      time.Time
	Status        string
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetAllTransaction(email string) ([]Domain, error)
}

type Repository interface {
	GetAllTransaction(email string) ([]Domain, error)
}
