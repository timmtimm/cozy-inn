package rooms

import (
	"time"
)

type Domain struct {
	RoomType       string
	Room           []RoomCondition
	Description    string
	ImageRoom_URLS []string
	Rules          []string
	Facilities     []string
	Capacity       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RoomCondition struct {
	Number int    `json:"number" validate:"required" firestore:"number"`
	Status string `json:"status" validate:"required" firestore:"status"`
}

type UseCase interface {
	GetAllRoom() ([]Domain, error)
	CreateRoom(roomDomain *Domain) error
	UpdateRoom(roomDomain *Domain) (Domain, error)
	DeleteRoom(roomType string) error
}

type Repository interface {
	GetAllRoom() ([]Domain, error)
	CreateRoom(roomDomain *Domain) error
	UpdateRoom(roomDomain *Domain) (Domain, error)
	DeleteRoom(roomType string) error
}
