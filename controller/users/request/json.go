package request

import (
	"cozy-inn/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name        string `json:"name" validate:"required" firestore:"name"`
	Email       string `json:"email" validate:"required,email" firestore:"email"`
	Password    string `json:"password" validate:"required" firestore:"password"`
	ImageID_URL string `json:"imageID_URL" validate:"required,url" firestore:"imageID_URL"`
	Role        string `json:"role" validate:"required" firestore:"role"`
	Status      bool   `json:"status" firestore:"status"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		ImageID_URL: req.ImageID_URL,
		Role:        req.Role,
		Status:      req.Status,
	}
}

func (req *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" firestore:"email"`
	Password string `json:"password" validate:"required" firestore:"password"`
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}

type UserUpdate struct {
	Name        string `json:"name" validate:"required" firestore:"name"`
	ImageID_URL string `json:"imageID_URL" validate:"required,url" firestore:"imageID_URL"`
}

func (req *UserUpdate) ToDomain() *users.Domain {
	return &users.Domain{
		Name:        req.Name,
		ImageID_URL: req.ImageID_URL,
	}
}

func (req *UserUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}

type AdminUpdate struct {
	Role   string `json:"role" validate:"required" firestore:"role"`
	Status bool   `json:"status" firestore:"status"`
}

func (req *AdminUpdate) ToDomain() *users.Domain {
	return &users.Domain{
		Role:   req.Role,
		Status: req.Status,
	}
}

func (req *AdminUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
