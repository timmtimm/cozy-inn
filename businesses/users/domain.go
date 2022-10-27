package users

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type Domain struct {
	UserID     string
	Role       string
	Email      string
	Password   string
	ImageIdUrl string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Repository interface {
	Register(ctx context.Context, client *firestore.Client)
}

type Usecase interface {
	Register(userDomain *Domain) Domain
}
