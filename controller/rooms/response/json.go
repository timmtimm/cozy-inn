package response

import (
	"cozy-inn/businesses/rooms"
	"time"
)

type Room struct {
	RoomType       string                `json:"roomType" firestore:"roomType"`
	Room           []rooms.RoomCondition `json:"room" firestore:"room"`
	Description    string                `json:"description" firestore:"description"`
	ImageRoom_URLS []string              `json:"imageRoom_URLS" validate:"required,url" firestore:"imageRoom_URLS"`
	Rules          []string              `json:"rules" firestore:"rules"`
	Facilities     []string              `json:"facilities" firestore:"facilities"`
	Capacity       int                   `json:"capacity" firestore:"capacity"`
	CreatedAt      time.Time             `json:"createdAt" firestore:"createdAt,serverTimestamp"`
	UpdatedAt      time.Time             `json:"updatedAt" firestore:"updatedAt,omitempty"`
}

func FromDomain(domain rooms.Domain) Room {
	return Room{
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
