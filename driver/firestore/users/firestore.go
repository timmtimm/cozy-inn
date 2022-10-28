package users

import (
	"context"
	"cozy-inn/businesses/users"
	"errors"
	"time"

	"log"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (ur *UserRepository) GetuserByEmail(email string) users.Domain {
	doc, err := ur.usersCollection().Doc(email).Get(ur.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return users.Domain{}
		}
	}

	userData := Model{}
	if err := doc.DataTo(&userData); err != nil {
		return users.Domain{}
	}

	return userData.ToDomain()
}

func (ur *UserRepository) Register(userDomain *users.Domain) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	doc, err := ur.usersCollection().Doc(rec.Email).Get(ur.ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err = ur.usersCollection().Doc(rec.Email).Set(ur.ctx, Model{
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
	}

	if doc != nil {
		return errors.New("email already registered")
	}

	return errors.New("failed to register")
}

func (ur *UserRepository) Login(userDomain *users.Domain) error {
	rec := FromDomain(userDomain)

	doc, err := ur.usersCollection().Doc(rec.Email).Get(ur.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return errors.New("email not registered")
		}
	}

	userData := Model{}
	if err := doc.DataTo(&userData); err != nil {
		return errors.New("failed get data")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(rec.Password))
	if err != nil {
		return errors.New("wrong email or password")
	}

	return nil
}
