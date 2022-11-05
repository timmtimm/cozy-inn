package rooms

import (
	"cozy-inn/businesses/rooms"
	"time"
)

type Model struct {
	RoomType       string       `firestore:"roomType"`
	Room           []rooms.Room `firestore:"room"`
	Description    string       `firestore:"description"`
	ImageRoom_URLS []string     `firestore:"imageRoom_URLS"`
	Capacity       int          `firestore:"capacity"`
	Price          int          `firestore:"price"`
	Facilities     []string     `firestore:"facilities"`
	Rules          []string     `firestore:"rules"`
	CreatedAt      time.Time    `firestore:"createdAt"`
	UpdatedAt      time.Time    `firestore:"updatedAt"`
}

func FromDomain(domain rooms.Domain) Model {
	return Model{
		RoomType:       domain.RoomType,
		Room:           domain.Room,
		Description:    domain.Description,
		ImageRoom_URLS: domain.ImageRoom_URLS,
		Capacity:       domain.Capacity,
		Price:          domain.Price,
		Facilities:     domain.Facilities,
		Rules:          domain.Rules,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func (rec *Model) ToDomain() rooms.Domain {
	return rooms.Domain{
		RoomType:       rec.RoomType,
		Room:           rec.Room,
		Description:    rec.Description,
		ImageRoom_URLS: rec.ImageRoom_URLS,
		Capacity:       rec.Capacity,
		Price:          rec.Price,
		Facilities:     rec.Facilities,
		Rules:          rec.Rules,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}
