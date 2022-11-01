package rooms

import (
	"context"
	"cozy-inn/businesses/rooms"

	"cloud.google.com/go/firestore"
)

type RoomRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewRoomRepository(client *firestore.Client, ctx context.Context) rooms.Repository {
	if client == nil {
		panic("No firestore client")
	}
	return &RoomRepository{client, ctx}
}

func (rr *RoomRepository) roomsCollection() *firestore.CollectionRef {
	return rr.client.Collection("rooms")
}

func (rr *RoomRepository) GetAllRoom() ([]rooms.Domain, error) {
	roomList := []rooms.Domain{}
	roomsDoc, err := rr.roomsCollection().Documents(rr.ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, roomDoc := range roomsDoc {
		roomData := Model{}

		if err := roomDoc.DataTo(&roomData); err != nil {
			return []rooms.Domain{}, err
		}

		roomList = append(roomList, roomData.ToDomain())
	}

	return roomList, nil
}
