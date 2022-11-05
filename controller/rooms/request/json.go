package request

import (
	"cozy-inn/businesses/rooms"

	"github.com/go-playground/validator/v10"
)

type Room struct {
	RoomType       string       `json:"roomType" validate:"required" firestore:"roomType"`
	Room           []rooms.Room `json:"room" validate:"required,dive,required" firestore:"room"`
	Description    string       `json:"description" validate:"required" firestore:"description"`
	ImageRoom_URLS []string     `json:"imageRoom_URLS" validate:"required,dive,required,url" firestore:"imageRoom_URLS"`
	Capacity       int          `json:"capacity" validate:"required" firestore:"capacity"`
	Price          int          `json:"price" validate:"required" firestore:"price"`
	Facilities     []string     `json:"facilities" validate:"required,dive,required" firestore:"facilities"`
	Rules          []string     `json:"rules" validate:"required,dive,required" firestore:"rules"`
}

func (req *Room) ToDomain() rooms.Domain {
	return rooms.Domain{
		RoomType:       req.RoomType,
		Room:           req.Room,
		Description:    req.Description,
		ImageRoom_URLS: req.ImageRoom_URLS,
		Capacity:       req.Capacity,
		Price:          req.Price,
		Facilities:     req.Facilities,
		Rules:          req.Rules,
	}
}

func (req *Room) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
