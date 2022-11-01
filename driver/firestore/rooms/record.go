package rooms

import (
	"cozy-inn/businesses/rooms"
	"time"
)

type Model struct {
	RoomType       string                `firestore:"roomType"`
	Room           []rooms.RoomCondition `firestore:"room"`
	Description    string                `firestore:"description"`
	ImageRoom_URLS []string              `firestore:"imageRoom_URLS"`
	Rules          []string              `firestore:"rules"`
	Facilities     []string              `firestore:"facilities"`
	Capacity       int                   `firestore:"capacity"`
	CreatedAt      time.Time             `firestore:"createdAt"`
	UpdatedAt      time.Time             `firestore:"updatedAt,omitempty"`
}

func FromDomain(domain rooms.Domain) *Model {
	return &Model{
		RoomType:       domain.RoomType,
		Room:           domain.Room,
		Description:    domain.Description,
		ImageRoom_URLS: domain.ImageRoom_URLS,
		Rules:          domain.Rules,
		Facilities:     domain.Facilities,
		Capacity:       domain.Capacity,
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
		Rules:          rec.Rules,
		Facilities:     rec.Facilities,
		Capacity:       rec.Capacity,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}
