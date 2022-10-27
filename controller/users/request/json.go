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
	Role        string `json:"role" firestore:"role"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		ImageID_URL: req.ImageID_URL,
		Role:        req.Role,
	}
}

func (req *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
