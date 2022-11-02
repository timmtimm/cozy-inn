package driver

import (
	"context"

	userDomain "cozy-inn/businesses/users"
	userDB "cozy-inn/driver/firestore/users"

	roomDomain "cozy-inn/businesses/rooms"
	roomDB "cozy-inn/driver/firestore/rooms"

	transactionDomain "cozy-inn/businesses/transactions"
	transactionDB "cozy-inn/driver/firestore/transactions"

	"cloud.google.com/go/firestore"
)

func NewUserRepository(fs *firestore.Client, ctx context.Context) userDomain.Repository {
	return userDB.NewUserRepository(fs, ctx)
}

func NewRoomRepository(fs *firestore.Client, ctx context.Context) roomDomain.Repository {
	return roomDB.NewRoomRepository(fs, ctx)
}

func NewTransactionRepository(fs *firestore.Client, ctx context.Context) transactionDomain.Repository {
	return transactionDB.NewTransactionRepository(fs, ctx)
}
