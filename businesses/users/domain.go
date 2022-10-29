package users

import (
	"time"
)

type Domain struct {
	Role        string
	Name        string
	Email       string
	Password    string
	ImageID_URL string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	Register(userDomain *Domain) (string, error)
	Login(userDomain *Domain) (string, error)
	GetUserByEmail(email string) Domain
}

type Repository interface {
	GetUserByEmail(email string) Domain
	Register(userDomain *Domain) error
	Login(userDomain *Domain) error
}
