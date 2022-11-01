package rooms

import (
	"context"
	"cozy-inn/businesses/rooms"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (rr *RoomRepository) CreateRoom(roomDomain *rooms.Domain) error {
	rec := FromDomain(*roomDomain)

	doc, err := rr.roomsCollection().Doc(rec.RoomType).Get(rr.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err = rr.roomsCollection().Doc(rec.RoomType).Set(rr.ctx, Model{
				RoomType:       rec.RoomType,
				Room:           rec.Room,
				Description:    rec.Description,
				ImageRoom_URLS: rec.ImageRoom_URLS,
				Rules:          rec.Rules,
				Facilities:     rec.Facilities,
				Capacity:       rec.Capacity,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
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

func (rr *RoomRepository) UpdateRoom(roomDomain *rooms.Domain) (rooms.Domain, error) {
	doc, err := rr.roomsCollection().Doc(roomDomain.RoomType).Get(rr.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return rooms.Domain{}, errors.New("room type not registered")
		}
	}

	roomData := Model{}
	if err := doc.DataTo(&roomData); err != nil {
		return rooms.Domain{}, err
	}

	roomData.Capacity = roomDomain.Capacity
	roomData.Description = roomDomain.Description
	roomData.Facilities = roomDomain.Facilities
	roomData.ImageRoom_URLS = roomDomain.ImageRoom_URLS
	roomData.Rules = roomDomain.Rules
	roomData.UpdatedAt = time.Now()

	_, err = rr.roomsCollection().Doc(roomDomain.RoomType).Set(rr.ctx, roomData)
	if err != nil {
		return rooms.Domain{}, errors.New("failed to update")
	}

	return roomData.ToDomain(), nil
}

func (rr *RoomRepository) DeleteRoom(roomType string) error {
	_, err := rr.roomsCollection().Doc(roomType).Get(rr.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return errors.New("room type not registered")
		}
	}

	_, err = rr.roomsCollection().Doc(roomType).Delete(rr.ctx)
	if err != nil {
		return err
	}

	return nil
}
