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
	Number int
	Status string
}

type UseCase interface {
	GetAllRoom() ([]Domain, error)
}

type Repository interface {
	GetAllRoom() ([]Domain, error)
}
