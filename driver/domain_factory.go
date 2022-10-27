package driver

import (
	userDomain "cozy-inn/businesses/users"
	userDB "cozy-inn/driver/firestore/users"

	"cloud.google.com/go/firestore"
)

func NewUserRepository(fs *firestore.Client) userDomain.Repository {
	return userDB.NewFirestoreRepository(fs)
}
