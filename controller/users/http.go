package users

import (
	"cozy-inn/businesses/users"
	"cozy-inn/controller/users/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(userUC users.UseCase) *UserController {
	return &UserController{
		userUseCase: userUC,
	}
}

func (userCtrl *UserController) Register(c echo.Context) error {
	userInput := request.User{}

	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	if err := userCtrl.userUseCase.Register(userInput.ToDomain()); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "success to register",
	})
}
