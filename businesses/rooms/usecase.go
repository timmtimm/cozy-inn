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
	rooms, err := ru.roomRepository.GetAllRoom()
	if err != nil {
		return []Domain{}, err
	}

	return rooms, nil
}

func (ru *RoomUseCase) GetRoom(roomType string) (Domain, error) {
	room, err := ru.roomRepository.GetRoomByType(roomType)
	if err != nil {
		return Domain{}, err
	}

	return room, nil
}

func (ru *RoomUseCase) CreateRoom(roomInput Domain) error {
	availableStatus := []string{"available", "unavailable"}
	statusFound := false
	for _, avaliableRole := range availableStatus {
		for _, room := range roomInput.Room {
			if room.Status == avaliableRole {
				statusFound = true
				break
			}
		}
	}

	if !statusFound {
		return errors.New("invalid status")
	}

	err := ru.roomRepository.Create(roomInput)
	if err != nil {
		return err
	}

	return nil
}

func (ru RoomUseCase) UpdateRoom(roomInput Domain) (Domain, error) {
	availableStatus := []string{"available", "unavailable"}
	statusFound := false
	for _, avaliableRole := range availableStatus {
		for _, room := range roomInput.Room {
			if room.Status == avaliableRole {
				statusFound = true
				break
			}
		}
	}

	if !statusFound {
		return Domain{}, errors.New("invalid status")
	}

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
