package rooms

import (
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoomUseCase struct {
	roomRepository Repository
}

func NewRoomUsecase(rr Repository) UseCase {
	return &RoomUseCase{
		roomRepository: rr,
	}
}

func (ru *RoomUseCase) GetAllRoom() ([]Domain, error) {
	return ru.roomRepository.GetAllRoom()
}

func (ru *RoomUseCase) CreateRoom(roomInput Domain) error {
	err := ru.roomRepository.Create(roomInput)
	if err != nil {
		return err
	}

	return nil
}

func (ru RoomUseCase) UpdateRoom(roomInput Domain) (Domain, error) {
	room, err := ru.roomRepository.GetRoomByType(roomInput.RoomType)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return Domain{}, errors.New("room type not registered")
		}
	}

	room.Room = roomInput.Room
	room.Description = roomInput.Description
	room.ImageRoom_URLS = roomInput.ImageRoom_URLS
	room.Capacity = roomInput.Capacity
	room.Price = roomInput.Price
	room.Facilities = roomInput.Facilities
	room.Rules = roomInput.Rules
	room.UpdatedAt = time.Now()

	err = ru.roomRepository.Update(room)
	if err != nil {
		return Domain{}, err
	}

	return room, nil
}

func (ru RoomUseCase) DeleteRoom(roomType string) error {
	_, err := ru.roomRepository.GetRoomByType(roomType)
	if err != nil {
		return err
	}

	err = ru.roomRepository.Delete(roomType)
	if err != nil {
		return err
	}

	return nil
}
