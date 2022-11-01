package rooms

import (
	"cozy-inn/businesses/rooms"
	"cozy-inn/controller/rooms/request"
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

func (roomCtrl *RoomController) CreateRoom(c echo.Context) error {
	RoomInput := request.Room{}

	if err := c.Bind(&RoomInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if RoomInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	err := roomCtrl.roomUseCase.CreateRoom(RoomInput.ToDomain())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to create room",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create room",
	})
}