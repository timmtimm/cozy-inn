package users

import (
	"context"
	"cozy-inn/businesses/users"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
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

func (ur *UserRepository) GetUserByEmail(email string) (users.Domain, error) {
	doc, err := ur.usersCollection().Doc(email).Get(ur.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return users.Domain{}, errors.New("user not found")
		}
	}

	userData := Model{}
	if err := doc.DataTo(&userData); err != nil {
		return users.Domain{}, err
	}

	return userData.ToDomain(), nil
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
				Status:      rec.Status,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})
			if err != nil {
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

	if !userData.Status {
		return errors.New("user is inactive")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(rec.Password))
	if err != nil {
		return errors.New("wrong email or password")
	}

	return nil
}

func (ur *UserRepository) Update(email string, userDomain *users.Domain) (users.Domain, error) {
	doc, err := ur.usersCollection().Doc(email).Get(ur.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return users.Domain{}, errors.New("email not registered")
		}
	}

	userData := Model{}
	if err := doc.DataTo(&userData); err != nil {
		return users.Domain{}, err
	}

	rec := FromDomain(userDomain)
	userData.Name = rec.Name
	userData.ImageID_URL = rec.ImageID_URL
	userData.UpdatedAt = time.Now()

	_, err = ur.usersCollection().Doc(email).Set(ur.ctx, userData)
	if err != nil {
		return users.Domain{}, errors.New("failed to update")
	}

	return userData.ToDomain(), nil
}

func (ur *UserRepository) AdminUpdateUser(email string, userDomain *users.Domain) (users.Domain, error) {
	doc, err := ur.usersCollection().Doc(email).Get(ur.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return users.Domain{}, errors.New("email not registered")
		}
	}

	userData := Model{}
	if err := doc.DataTo(&userData); err != nil {
		return users.Domain{}, err
	}

	rec := FromDomain(userDomain)

	userData.Status = rec.Status
	userData.Role = rec.Role
	userData.UpdatedAt = time.Now()

	_, err = ur.usersCollection().Doc(email).Set(ur.ctx, userData)
	if err != nil {
		return users.Domain{}, errors.New("failed to update")
	}

	return userData.ToDomain(), nil
}

func (ur *UserRepository) GetUserList() ([]users.Domain, error) {
	iter := ur.usersCollection().Documents(ur.ctx)
	userData := []users.Domain{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		user := Model{}
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}

		userData = append(userData, user.ToDomain())
	}

	return userData, nil
}

func (ur *UserRepository) Delete(email string) error {
	_, err := ur.usersCollection().Doc(email).Delete(ur.ctx)
	if err != nil {
		return err
	}

	return nil
}
