package rooms

import (
	"time"
)

type Domain struct {
	RoomType       string
	Room           []Room
	Description    string
	ImageRoom_URLS []string
	Capacity       int
	Price          int
	Facilities     []string
	Rules          []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Room struct {
	Number int    `json:"number" validate:"required" firestore:"number"`
	Status string `json:"status" validate:"required" firestore:"status"`
}

type UseCase interface {
	GetAllRoom() ([]Domain, error)
	CreateRoom(roomDomain Domain) error
	UpdateRoom(roomDomain Domain) (Domain, error)
	DeleteRoom(roomType string) error
}

type Repository interface {
	GetAllRoom() ([]Domain, error)
	GetRoomByType(roomType string) (Domain, error)
	Create(roomDomain Domain) error
	Update(roomDomain Domain) error
	Delete(roomType string) error
}
