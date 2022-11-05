package users

import (
	"time"
)

type Domain struct {
	Email       string
	Password    string `json:"-"`
	Name        string
	Role        string
	ImageID_URL string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	GetUserByEmail(email string) (Domain, error)
	GetUserList() ([]Domain, error)
	UserRegister(user Domain) (string, error)
	AdminRegister(user Domain) error
	Login(user Domain) (string, error)
	UserUpdate(email string, user Domain) (Domain, error)
	AdminUpdate(email string, user Domain) (Domain, error)
	AdminDelete(email string) error
}

type Repository interface {
	GetUserByEmail(email string) (Domain, error)
	GetUserList() ([]Domain, error)
	Register(user Domain) error
	Login(user Domain) error
	Update(email string, user Domain) error
	Delete(email string) error
}
