package users

import (
	"time"
)

type Model struct {
	UserID     string    `firestore:"userID"`
	Role       string    `firestore:"role"`
	Email      string    `firestore:"email"`
	Password   string    `firestore:"password"`
	ImageIdUrl string    `firestore:"image_id_url"`
	CreatedAt  time.Time `firestore:"createdAt,serverTimestamp"`
	UpdatedAt  time.Time `firestore:"updatedAt,omitempty"`
}
