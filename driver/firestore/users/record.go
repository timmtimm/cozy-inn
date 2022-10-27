package users

import (
	"cozy-inn/businesses/users"
	"time"
)

type Model struct {
	UserID      string    `firestore:"userID"`
	Role        string    `firestore:"role"`
	Name        string    `firestore:"name"`
	Email       string    `firestore:"email"`
	Password    string    `firestore:"password"`
	ImageID_URL string    `firestore:"imageID_URL"`
	CreatedAt   time.Time `firestore:"createdAt,serverTimestamp"`
	UpdatedAt   time.Time `firestore:"updatedAt,omitempty"`
}

func FromDomain(domain *users.Domain) *Model {
	return &Model{
		UserID:      domain.UserID,
		Role:        domain.Role,
		Name:        domain.Name,
		Email:       domain.Email,
		Password:    domain.Password,
		ImageID_URL: domain.ImageID_URL,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (rec *Model) ToDomain() users.Domain {
	return users.Domain{
		UserID:      rec.UserID,
		Role:        rec.Role,
		Name:        rec.Name,
		Email:       rec.Email,
		Password:    rec.Password,
		ImageID_URL: rec.ImageID_URL,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
