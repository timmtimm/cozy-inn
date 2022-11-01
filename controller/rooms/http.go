package rooms

import (
	"cozy-inn/businesses/rooms"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomController struct {
	roomUseCase rooms.UseCase
}

func NewRoomController(roomUC rooms.UseCase) *RoomController {
	return &RoomController{
		roomUseCase: roomUC,
	}
}

func (roomCtrl *RoomController) GetAllRoom(c echo.Context) error {
	rooms, err := roomCtrl.roomUseCase.GetAllRoom()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all room",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all room",
		"data":    rooms,
	})
}
