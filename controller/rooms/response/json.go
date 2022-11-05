package response

import (
	"cozy-inn/businesses/rooms"
	"time"
)

type Room struct {
	RoomType       string       `json:"roomType" firestore:"roomType"`
	Room           []rooms.Room `json:"room" firestore:"room"`
	Description    string       `json:"description" firestore:"description"`
	ImageRoom_URLS []string     `json:"imageRoom_URLS" validate:"required,url" firestore:"imageRoom_URLS"`
	Capacity       int          `json:"capacity" firestore:"capacity"`
	Price          int          `json:"price" validate:"required" firestore:"price"`
	Facilities     []string     `json:"facilities" firestore:"facilities"`
	Rules          []string     `json:"rules" firestore:"rules"`
	CreatedAt      time.Time    `json:"createdAt" firestore:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt" firestore:"updatedAt"`
}

func FromDomain(domain rooms.Domain) Room {
	return Room{
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
