package users

import (
	"context"
	"cozy-inn/businesses/users"

	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type UserRepository struct {
	*firestore.Client
}

func NewFirestoreRepository(client *firestore.Client) users.Repository {
	if client == nil {
		panic("No firestore client")
	}
	return &UserRepository{client}
}

func (ur *UserRepository) usersCollection() *firestore.CollectionRef {
	return ur.Collection("users")
}

func (ur *UserRepository) Register(ctx context.Context, client *firestore.Client) {
	res, err := ur.usersCollection().Doc("timotiuswirawan@gmail.com").Set(ctx, Model{
		UserID:     "asd123",
		Role:       "admin",
		Email:      "timotiuswirawan38@gmail.com",
		Password:   "asdasd",
		ImageIdUrl: "blabla",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	fmt.Println(res)
}
