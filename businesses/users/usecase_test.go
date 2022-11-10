package users_test

import (
	_middleware "cozy-inn/app/middleware"
	"cozy-inn/businesses/users"
	_userMock "cozy-inn/businesses/users/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository _userMock.Repository
	userUseCase    users.UseCase
	userDomain     users.Domain
)

func TestMain(m *testing.M) {
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       "secret",
		ExpiresDuration: 1,
	}

	userUseCase = users.NewUserUsecase(&userRepository, &configJWT)

	userDomain = users.Domain{
		Email:       "test@gmail.com",
		Password:    "test123",
		Name:        "tester",
		Role:        "user",
		ImageID_URL: "https://www.google.com",
		Status:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.Run()
}

func TestGetUserByEmail(t *testing.T) {
	t.Run("Test Case 1 | Valid Get User By Email", func(t *testing.T) {
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()

		result, err := userUseCase.GetUserByEmail(userDomain.Email)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, result)
	})

	t.Run("Test Case 2 | Invalid Get User By Email", func(t *testing.T) {
		expectedErr := errors.New("user not found")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(users.Domain{}, expectedErr).Once()

		result, err := userUseCase.GetUserByEmail(userDomain.Email)

		assert.Equal(t, expectedErr, err)
		assert.Equal(t, users.Domain{}, result)
	})
}

func TestGetUserList(t *testing.T) {
	t.Run("Test Case 1 | Valid Get User List", func(t *testing.T) {
		userRepository.On("GetUserList").Return([]users.Domain{userDomain}, nil).Once()

		result, err := userUseCase.GetUserList()

		assert.Nil(t, err)
		assert.Equal(t, []users.Domain{userDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get User List", func(t *testing.T) {
		expectedErr := errors.New("")
		userRepository.On("GetUserList").Return([]users.Domain{}, expectedErr).Once()

		result, actualErr := userUseCase.GetUserList()

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []users.Domain{}, result)
	})
}

func TestUserRegister(t *testing.T) {
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		userRepository.On("Register", userDomain).Return(nil).Once()

		actualToken, err := userUseCase.UserRegister(userDomain)
		assert.NotNil(t, actualToken)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Register", func(t *testing.T) {
		expectedErr := errors.New("email already registered")
		userRepository.On("Register", userDomain).Return(expectedErr).Once()

		actualToken, actualErr := userUseCase.UserRegister(userDomain)
		assert.Equal(t, expectedErr, actualErr)
		assert.Empty(t, actualToken)
	})
}

func TestAdminRegister(t *testing.T) {
	t.Run("Test Case 1 | Valid Admin Register", func(t *testing.T) {
		userRepository.On("Register", userDomain).Return(nil).Once()

		actualErr := userUseCase.AdminRegister(userDomain)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Admin Register", func(t *testing.T) {
		expectedErr := errors.New("email already registered")
		userRepository.On("Register", userDomain).Return(expectedErr).Once()

		actualErr := userUseCase.AdminRegister(userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Admin Register", func(t *testing.T) {
		userDomain.Role = "invalid"
		expectedErr := errors.New("invalid role")
		userRepository.On("Register", userDomain).Return(expectedErr).Once()

		actualErr := userUseCase.AdminRegister(userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		userRepository.On("Login", userDomain).Return(nil).Once()
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()

		actualToken, err := userUseCase.Login(userDomain)
		assert.NotNil(t, actualToken)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Login", func(t *testing.T) {
		expectedErr := errors.New("wrong email or password")
		userRepository.On("Login", userDomain).Return(expectedErr).Once()

		actualToken, actualErr := userUseCase.Login(userDomain)
		assert.Equal(t, expectedErr, actualErr)
		assert.Empty(t, actualToken)
	})

	t.Run("Test Case 3 | Invalid Login", func(t *testing.T) {
		expectedErrGetUserByEmail := errors.New("user not found")
		userRepository.On("Login", userDomain).Return(nil).Once()
		userRepository.On("GetUserByEmail", userDomain.Email).Return(users.Domain{}, expectedErrGetUserByEmail).Once()

		actualToken, actualErr := userUseCase.Login(userDomain)
		assert.Equal(t, expectedErrGetUserByEmail, actualErr)
		assert.Empty(t, actualToken)
	})
}

func TestUserUpdate(t *testing.T) {
	t.Run("Test Case 1 | Valid User Update", func(t *testing.T) {
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()
		userRepository.On("Update", userDomain.Email, mock.Anything).Return(nil).Once()

		updatedUser, actualErr := userUseCase.UserUpdate(userDomain.Email, userDomain)
		assert.Nil(t, actualErr)
		assert.NotNil(t, updatedUser)
	})

	t.Run("Test Case 2 | Invalid User Update", func(t *testing.T) {
		expectedErr := errors.New("failed to update")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()
		userRepository.On("Update", userDomain.Email, mock.Anything).Return(expectedErr).Once()

		_, actualErr := userUseCase.UserUpdate(userDomain.Email, userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid User Update", func(t *testing.T) {
		userDomain.Email = "invalid"
		expectedErr := errors.New("user not found")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(users.Domain{}, expectedErr).Once()

		_, actualErr := userUseCase.UserUpdate(userDomain.Email, userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestAdminUpdate(t *testing.T) {
	userDomain.Role = "user"
	t.Run("Test Case 1 | Valid Admin Update", func(t *testing.T) {
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()
		userRepository.On("Update", userDomain.Email, mock.Anything).Return(nil).Once()

		updatedUser, err := userUseCase.AdminUpdate(userDomain.Email, userDomain)
		assert.Nil(t, err)
		assert.NotNil(t, updatedUser)
	})

	t.Run("Test Case 2 | Invalid Admin Update", func(t *testing.T) {
		userDomain.Role = "user"
		expectedErr := errors.New("failed to update")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()
		userRepository.On("Update", userDomain.Email, mock.Anything).Return(expectedErr).Once()

		_, actualErr := userUseCase.AdminUpdate(userDomain.Email, userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Admin Update", func(t *testing.T) {
		userDomain.Email = "invalid"
		expectedErr := errors.New("invalid role")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(users.Domain{}, expectedErr).Once()

		_, actualErr := userUseCase.AdminUpdate(userDomain.Email, userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Admin Update", func(t *testing.T) {
		userDomain.Role = "admin"
		expectedErr := errors.New("invalid role")

		_, actualErr := userUseCase.AdminUpdate(userDomain.Email, userDomain)
		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestAdminDelete(t *testing.T) {
	t.Run("Test Case 1 | Valid Admin Delete", func(t *testing.T) {
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()
		userRepository.On("Delete", userDomain.Email).Return(nil).Once()

		actualErr := userUseCase.AdminDelete(userDomain.Email)
		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Admin Delete", func(t *testing.T) {
		expectedErr := errors.New("failed to delete")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(userDomain, nil).Once()
		userRepository.On("Delete", userDomain.Email).Return(expectedErr).Once()

		actualErr := userUseCase.AdminDelete(userDomain.Email)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Admin Delete", func(t *testing.T) {
		userDomain.Email = "invalid"
		expectedErr := errors.New("user not found")
		userRepository.On("GetUserByEmail", userDomain.Email).Return(users.Domain{}, expectedErr).Once()

		actualErr := userUseCase.AdminDelete(userDomain.Email)
		assert.Equal(t, expectedErr, actualErr)
	})

}
