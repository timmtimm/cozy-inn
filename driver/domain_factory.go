package driver

import (
	"context"
	userDomain "cozy-inn/businesses/users"
	userDB "cozy-inn/driver/firestore/users"

	"cloud.google.com/go/firestore"
)

func NewUserRepository(fs *firestore.Client, ctx context.Context) userDomain.Repository {
	return userDB.NewUserRepository(fs, ctx)
}
