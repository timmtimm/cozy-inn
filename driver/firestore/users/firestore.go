package users

import (
	"context"
	"cozy-inn/businesses/users"
	"errors"
	"time"

	"log"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

type UserRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewUserRepository(client *firestore.Client, ctx context.Context) users.Repository {
	if client == nil {
		panic("No firestore client")
	}
	return &UserRepository{client, ctx}
}

func (ur *UserRepository) usersCollection() *firestore.CollectionRef {
	return ur.client.Collection("users")
}

func (ur *UserRepository) Register(userDomain *users.Domain) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	iter := ur.usersCollection().Where("email", "==", rec.Email).Documents(ur.ctx)
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		return errors.New("email already exists")
	}

	_, err := ur.usersCollection().Doc(rec.Email).Set(ur.ctx, Model{
		Role:        rec.Role,
		Name:        rec.Name,
		Email:       rec.Email,
		Password:    rec.Password,
		ImageID_URL: rec.ImageID_URL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}

	return nil
}
