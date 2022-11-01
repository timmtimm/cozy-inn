package users

import (
	"time"
)

type Domain struct {
	Role        string
	Name        string
	Email       string
	Password    string `json:"-"`
	ImageID_URL string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	Register(userDomain *Domain) (string, error)
	SudoRegister(userDomain *Domain) (string, error)
	Login(userDomain *Domain) (string, error)
	GetUserByEmail(email string) (Domain, error)
	UpdateUser(email string, userDomain *Domain) (Domain, error)
	GetUserList() ([]Domain, error)
	AdminUpdateUser(email string, userDomain *Domain) (Domain, error)
	AdminDeleteUser(email string) error
}

type Repository interface {
	GetUserByEmail(email string) (Domain, error)
	Register(userDomain *Domain) error
	Login(userDomain *Domain) error
	Update(email string, userDomain *Domain) (Domain, error)
	GetUserList() ([]Domain, error)
	AdminUpdateUser(email string, userDomain *Domain) (Domain, error)
	Delete(email string) error
}
