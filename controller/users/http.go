package users

import (
	"cozy-inn/app/middleware"
	"cozy-inn/businesses/users"
	"cozy-inn/controller/users/request"
	"cozy-inn/controller/users/response"
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

func (userCtrl *UserController) UserRegister(c echo.Context) error {
	userInput := request.User{}

	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	userInput.Role = "user"
	userInput.Status = true
	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token, err := userCtrl.userUseCase.UserRegister(userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "success to register",
		"token":   token,
	})
}

func (userCtrl *UserController) Login(c echo.Context) error {
	userInput := request.UserLogin{}
	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "required filled form is invalid",
		})
	}

	token, err := userCtrl.userUseCase.Login(userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success to login",
		"token":   token,
	})
}

func (userCtrl *UserController) GetUserProfile(c echo.Context) error {
	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	user, err := userCtrl.userUseCase.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get user profile",
		"user":    response.FromDomain(user),
	})
}

func (userCtrl *UserController) UpdateUserProfile(c echo.Context) error {
	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	userInput := request.UserUpdate{}
	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "required filled form is invalid",
		})
	}

	user, err := userCtrl.userUseCase.UserUpdate(email, userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to update user profile",
		"user":    response.FromDomain(user),
	})
}

func (userCtrl *UserController) AdminGetUserList(c echo.Context) error {
	userList, err := userCtrl.userUseCase.GetUserList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get user list",
		"user":    userList,
	})
}

func (userCtrl *UserController) AdminRegister(c echo.Context) error {
	userInput := request.User{}
	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	userInput.Status = true
	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	err := userCtrl.userUseCase.AdminRegister(userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "success to register",
	})
}

func (userCtrl *UserController) AdminGetUser(c echo.Context) error {
	userEmail := c.Param("user-email")

	user, err := userCtrl.userUseCase.GetUserByEmail(userEmail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get user profile",
		"user":    response.FromDomain(user),
	})
}

func (userCtrl *UserController) AdminUpdate(c echo.Context) error {
	userEmail := c.Param("user-email")

	userInput := request.AdminUpdate{}
	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "required filled form is invalid",
		})
	}

	user, err := userCtrl.userUseCase.AdminUpdate(userEmail, userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to update user profile",
		"user":    response.FromDomain(user),
	})
}

func (userCtrl *UserController) AdminDelete(c echo.Context) error {
	userEmail := c.Param("user-email")

	err := userCtrl.userUseCase.AdminDelete(userEmail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success to delete user",
	})
}
