package rooms_test

import (
	"cozy-inn/businesses/rooms"
	_roomMock "cozy-inn/businesses/rooms/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	roomRepository _roomMock.Repository
	roomUseCase    rooms.UseCase
	roomDomain     rooms.Domain
)

func TestMain(m *testing.M) {
	roomUseCase = rooms.NewRoomUsecase(&roomRepository)

	roomDomain = rooms.Domain{
		RoomType: "test",
		Room: []rooms.Room{
			{
				Number: 1,
				Status: "available",
			},
		},
		Description:    "test",
		ImageRoom_URLS: []string{"https://www.google.com"},
		Capacity:       2,
		Price:          100000,
		Facilities:     []string{"test"},
		Rules:          []string{"test"},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	m.Run()
}

func TestGetAllRoom(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All Room", func(t *testing.T) {
		roomRepository.On("GetAllRoom").Return([]rooms.Domain{roomDomain}, nil).Once()

		result, actualErr := roomUseCase.GetAllRoom()

		assert.Nil(t, actualErr)
		assert.Equal(t, []rooms.Domain{roomDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get All Room", func(t *testing.T) {
		expectedErr := errors.New("")
		roomRepository.On("GetAllRoom").Return([]rooms.Domain{}, expectedErr).Once()

		result, actualErr := roomUseCase.GetAllRoom()

		assert.Equal(t, actualErr, expectedErr)
		assert.Equal(t, []rooms.Domain{}, result)
	})
}

func TestGetRoom(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Room", func(t *testing.T) {
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(roomDomain, nil).Once()

		result, actualErr := roomUseCase.GetRoom(roomDomain.RoomType)

		assert.Nil(t, actualErr)
		assert.Equal(t, roomDomain, result)
	})

	t.Run("Test Case 2 | Invalid Get Room", func(t *testing.T) {
		roomDomain.RoomType = "not exist"
		expectedErr := errors.New("room type not registered")
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(rooms.Domain{}, expectedErr).Once()

		result, actualErr := roomUseCase.GetRoom(roomDomain.RoomType)

		assert.Equal(t, actualErr, expectedErr)
		assert.Equal(t, rooms.Domain{}, result)
	})
}

func TestCreateRoom(t *testing.T) {
	t.Run("Test Case 1 | Valid Create Room", func(t *testing.T) {
		roomRepository.On("Create", roomDomain).Return(nil).Once()

		actualErr := roomUseCase.CreateRoom(roomDomain)

		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Create Room", func(t *testing.T) {
		roomDomain.RoomType = "exist"
		expectedErr := errors.New("room type already added")
		roomRepository.On("Create", roomDomain).Return(expectedErr).Once()

		actualErr := roomUseCase.CreateRoom(roomDomain)

		assert.Equal(t, actualErr, expectedErr)
	})

	t.Run("Test Case 3 | Invalid Create Room", func(t *testing.T) {
		expectedErr := errors.New("failed to add room")
		roomRepository.On("Create", roomDomain).Return(expectedErr).Once()

		actualErr := roomUseCase.CreateRoom(roomDomain)

		assert.Equal(t, actualErr, expectedErr)
	})

	t.Run("Test Case 4 | Invalid Create Room", func(t *testing.T) {
		roomDomain.Room[0].Status = "invalid"
		expectedErr := errors.New("invalid status")

		actualErr := roomUseCase.CreateRoom(roomDomain)

		assert.Equal(t, actualErr, expectedErr)
	})
}

func TestUpdateRoom(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Room", func(t *testing.T) {
		roomDomain.Room[0].Status = "available"
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(roomDomain, nil).Once()
		roomRepository.On("Update", mock.Anything).Return(nil).Once()

		updatedRoom, actualErr := roomUseCase.UpdateRoom(roomDomain)

		assert.Nil(t, actualErr)
		assert.NotNil(t, updatedRoom)
	})

	t.Run("Test Case 2 | Invalid Update Room", func(t *testing.T) {
		roomDomain.Room[0].Status = "available"
		expectedErr := errors.New("failed to update room")
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(roomDomain, nil).Once()
		roomRepository.On("Update", mock.Anything).Return(expectedErr).Once()

		updatedRoom, actualErr := roomUseCase.UpdateRoom(roomDomain)

		assert.Equal(t, actualErr, expectedErr)
		assert.NotNil(t, updatedRoom)
	})

	t.Run("Test Case 3 | Invalid Update Room", func(t *testing.T) {
		roomDomain.Room[0].Status = "invalid"
		expectedErr := errors.New("invalid status")

		result, actualErr := roomUseCase.UpdateRoom(roomDomain)

		assert.Equal(t, actualErr, expectedErr)
		assert.Equal(t, rooms.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Update Room", func(t *testing.T) {
		roomDomain.RoomType = "not exist"
		roomDomain.Room[0].Status = "available"
		expectedErr := errors.New("room type not registered")
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(rooms.Domain{}, expectedErr).Once()

		result, actualErr := roomUseCase.UpdateRoom(roomDomain)

		assert.Equal(t, actualErr, expectedErr)
		assert.Equal(t, rooms.Domain{}, result)
	})
}

func TestDeleteRoom(t *testing.T) {
	t.Run("Test Case 1 | Valid Delete Room", func(t *testing.T) {
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(roomDomain, nil).Once()
		roomRepository.On("Delete", roomDomain.RoomType).Return(nil).Once()

		actualErr := roomUseCase.DeleteRoom(roomDomain.RoomType)

		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Delete Room", func(t *testing.T) {
		roomDomain.RoomType = "not exist"
		expectedErr := errors.New("room type not registered")
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(rooms.Domain{}, expectedErr).Once()

		actualErr := roomUseCase.DeleteRoom(roomDomain.RoomType)

		assert.Equal(t, actualErr, expectedErr)
	})

	t.Run("Test Case 3 | Invalid Delete Room", func(t *testing.T) {
		expectedErr := errors.New("failed to delete room")
		roomRepository.On("GetRoomByType", roomDomain.RoomType).Return(roomDomain, nil).Once()
		roomRepository.On("Delete", roomDomain.RoomType).Return(expectedErr).Once()

		actualErr := roomUseCase.DeleteRoom(roomDomain.RoomType)

		assert.Equal(t, actualErr, expectedErr)
	})
}
