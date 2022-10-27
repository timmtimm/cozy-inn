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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	Register(userDomain *Domain) error
}

type Repository interface {
	Register(userDomain *Domain) error
}
