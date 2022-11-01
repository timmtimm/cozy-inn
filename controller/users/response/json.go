package response

import (
	"cozy-inn/businesses/users"
	"time"
)

type User struct {
	Role        string    `json:"role" firestore:"role"`
	Name        string    `json:"name" firestore:"name"`
	Email       string    `json:"email" firestore:"email"`
	ImageID_URL string    `json:"imageID_URL" firestore:"imageID_URL"`
	Status      bool      `json:"status" firestore:"status"`
	CreatedAt   time.Time `json:"createdAt" firestore:"createdAt,serverTimestamp"`
	UpdatedAt   time.Time `json:"updatedAt" firestore:"updatedAt,omitempty"`
}

func FromDomain(domain users.Domain) User {
	return User{
		Name:        domain.Name,
		Role:        domain.Role,
		Email:       domain.Email,
		ImageID_URL: domain.ImageID_URL,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
