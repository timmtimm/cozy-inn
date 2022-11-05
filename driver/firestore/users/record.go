package users

import (
	"cozy-inn/businesses/users"
	"time"
)

type Model struct {
	Email       string    `firestore:"email"`
	Password    string    `firestore:"password"`
	Name        string    `firestore:"name"`
	Role        string    `firestore:"role"`
	ImageID_URL string    `firestore:"imageID_URL"`
	Status      bool      `firestore:"status"`
	CreatedAt   time.Time `firestore:"createdAt"`
	UpdatedAt   time.Time `firestore:"updatedAt"`
}

func FromDomain(domain users.Domain) Model {
	return Model{
		Email:       domain.Email,
		Password:    domain.Password,
		Name:        domain.Name,
		Role:        domain.Role,
		ImageID_URL: domain.ImageID_URL,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (rec *Model) ToDomain() users.Domain {
	return users.Domain{
		Email:       rec.Email,
		Password:    rec.Password,
		Name:        rec.Name,
		Role:        rec.Role,
		ImageID_URL: rec.ImageID_URL,
		Status:      rec.Status,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
