package response

import (
	"cozy-inn/businesses/users"
	"time"
)

type User struct {
	Role        string    `json:"role" firestore:"role"`
	Name        string    `json:"name" firestore:"name"`
	Email       string    `json:"email" firestore:"email"`
	Password    string    `json:"password" firestore:"password"`
	ImageID_URL string    `json:"imageID_URL" firestore:"imageID_URL"`
	Status      string    `json:"status" firestore:"status"`
	CreatedAt   time.Time `json:"createdAt" firestore:"createdAt,serverTimestamp"`
	UpdatedAt   time.Time `json:"updatedAt" firestore:"updatedAt,omitempty"`
}

func FromDomain(domain users.Domain) User {
	return User{
		Role:        domain.Role,
		Email:       domain.Email,
		Password:    domain.Password,
		ImageID_URL: domain.ImageID_URL,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
