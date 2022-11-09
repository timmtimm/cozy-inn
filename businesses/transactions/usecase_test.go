package transactions_test

import (
	"cozy-inn/businesses/rooms"
	_roomMock "cozy-inn/businesses/rooms/mocks"
	"cozy-inn/businesses/transactions"
	_transactionMock "cozy-inn/businesses/transactions/mocks"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	transactionRepository _transactionMock.Repository
	roomRepository        _roomMock.Repository
	transactionUseCase    transactions.UseCase
	roomDomain            rooms.Domain
	transactionDomain     transactions.Domain
)

func TestMain(m *testing.M) {
	transactionUseCase = transactions.NewTransactionUsecase(&transactionRepository, &roomRepository)

	roomDomain = rooms.Domain{
		RoomType: "test",
		Room: []rooms.Room{
			{
				Number: 1,
				Status: "available",
			},
			{
				Number: 2,
				Status: "unavailable",
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

	transactionDomain = transactions.Domain{
		TransactionID: "1",
		UserEmail:     "test@gmail.com",
		RoomType:      "test",
		RoomNumber:    1,
		StartDate:     time.Now().AddDate(0, 0, 1),
		EndDate:       time.Now().AddDate(0, 0, 2),
		CheckIn:       time.Now().AddDate(0, 0, 1),
		CheckOut:      time.Now().AddDate(0, 0, 2),
		Status:        "test",
		Bill:          100000,
		Payment_URL:   "https://www.google.com",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	m.Run()
}

func TestGetAllTransactionUser(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All Transaction User", func(t *testing.T) {
		transactionRepository.On("GetAllTransactionByEmail", transactionDomain.UserEmail).Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, actualErr := transactionUseCase.GetAllTransactionUser(transactionDomain.UserEmail)

		assert.Nil(t, actualErr)
		assert.NotNil(t, []transactions.Domain{transactionDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get All Transaction User", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetAllTransactionByEmail", transactionDomain.UserEmail).Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetAllTransactionUser(transactionDomain.UserEmail)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []transactions.Domain{}, result)
	})
}

func TestGetAllReadyCheckIn(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All Ready Check In", func(t *testing.T) {
		transactionRepository.On("GetAllReadyCheckIn").Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, actualErr := transactionUseCase.GetAllReadyCheckIn()

		assert.Nil(t, actualErr)
		assert.Equal(t, []transactions.Domain{transactionDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get All Ready Check In", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetAllReadyCheckIn").Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetAllReadyCheckIn()

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []transactions.Domain{}, result)
	})
}

func TestGetAllReadyCheckOut(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All Ready Check Out", func(t *testing.T) {
		transactionRepository.On("GetAllReadyCheckOut").Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, actualErr := transactionUseCase.GetAllReadyCheckOut()

		assert.Nil(t, actualErr)
		assert.Equal(t, []transactions.Domain{transactionDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get All Ready Check Out", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetAllReadyCheckOut").Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetAllReadyCheckOut()

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []transactions.Domain{}, result)
	})
}

func TestGetAllPaymentNotVerified(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All Payment Not Verified", func(t *testing.T) {
		transactionRepository.On("GetAllPaymentNotVerified").Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, actualErr := transactionUseCase.GetAllPaymentNotVerified()

		assert.Nil(t, actualErr)
		assert.Equal(t, []transactions.Domain{transactionDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Get All Payment Not Verified", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetAllPaymentNotVerified").Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetAllPaymentNotVerified()

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []transactions.Domain{}, result)
	})
}

func TestAdminGetAllTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Admin Get All Transaction", func(t *testing.T) {
		transactionRepository.On("GetAllTransaction").Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, actualErr := transactionUseCase.AdminGetAllTransaction()

		assert.Nil(t, actualErr)
		assert.Equal(t, []transactions.Domain{transactionDomain}, result)
	})

	t.Run("Test Case 2 | Invalid Admin Get All Transaction", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetAllTransaction").Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.AdminGetAllTransaction()

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []transactions.Domain{}, result)
	})
}

func TestGetTransactionOnVerification(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Transaction On Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verification-pending"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.GetTransactionOnVerification(transactionInput.TransactionID)

		assert.Nil(t, actualErr)
		assert.Equal(t, transactionInput, result)
	})

	t.Run("Test Case 2 | Invalid Get Transaction On Verification", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetTransactionOnVerification(transactionDomain.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Get Transaction On Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.TransactionID = "not exist"
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetTransactionOnVerification(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Get Transaction On Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "not verification-pending"
		expectedErr := errors.New("transaction is not verification-pending")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.GetTransactionOnVerification(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestGetCheckInTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Check In Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.GetCheckInTransaction(transactionInput.TransactionID)

		assert.Nil(t, actualErr)
		assert.Equal(t, transactionInput, result)
	})

	t.Run("Test Case 2 | Invalid Get Check In Transaction", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetCheckInTransaction(transactionDomain.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Get Check In Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.TransactionID = "not exist"
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetCheckInTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Get Check In Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "not verified"
		expectedErr := errors.New("transaction is not verified")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.GetCheckInTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestGetCheckOutTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Check Out Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.GetCheckOutTransaction(transactionInput.TransactionID)

		assert.Nil(t, actualErr)
		assert.Equal(t, transactionInput, result)
	})

	t.Run("Test Case 2 | Invalid Get Check Out Transaction", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetCheckOutTransaction(transactionDomain.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Get Check Out Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.TransactionID = "not exist"
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetCheckOutTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Get Check Out Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "not checked-in"
		expectedErr := errors.New("transaction is not checked-in")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.GetCheckOutTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestGetTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Transaction", func(t *testing.T) {
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactionDomain, nil).Once()

		result, actualErr := transactionUseCase.GetTransaction(transactionDomain.TransactionID)

		assert.Nil(t, actualErr)
		assert.Equal(t, transactionDomain, result)
	})

	t.Run("Test Case 2 | Invalid Get Transaction", func(t *testing.T) {
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetTransaction(transactionDomain.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Get Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.TransactionID = "not exist"
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.GetTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestCheckAvailabilityAllRoom(t *testing.T) {
	t.Run("Test Case 1 | Valid Check Availability All Room", func(t *testing.T) {
		expectedResult := []string{}
		for _, room := range roomDomain.Room {
			if room.Status == "available" {
				expectedResult = append(expectedResult, fmt.Sprintf("%s-%d", roomDomain.RoomType, room.Number))
			}
		}

		roomRepository.On("GetAllRoom").Return([]rooms.Domain{roomDomain}, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, nil).Once()

		actualResult, actualErr := transactionUseCase.CheckAvailabilityAllRoom(transactionDomain.StartDate, transactionDomain.CheckOut)

		assert.Nil(t, actualErr)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Test Case 2 | Invalid Check Availability All Room", func(t *testing.T) {
		expectedErr := errors.New("")
		roomRepository.On("GetAllRoom").Return([]rooms.Domain{roomDomain}, expectedErr).Once()

		result, actualErr := transactionUseCase.CheckAvailabilityAllRoom(transactionDomain.StartDate, transactionDomain.EndDate)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []string{}, result)
	})

	t.Run("Test Case 3 | Invalid Check Availability All Room", func(t *testing.T) {
		expectedErr := errors.New("")
		roomRepository.On("GetAllRoom").Return([]rooms.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.CheckAvailabilityAllRoom(transactionDomain.StartDate, transactionDomain.EndDate)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []string{}, result)
	})

	t.Run("Test Case 4 | Invalid Check Availability All Room", func(t *testing.T) {
		expectedErr := errors.New("")
		roomRepository.On("GetAllRoom").Return([]rooms.Domain{roomDomain}, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.CheckAvailabilityAllRoom(transactionDomain.StartDate, transactionDomain.EndDate)

		assert.Empty(t, actualErr)
		assert.Equal(t, []string{}, result)
	})

	t.Run("Test Case 5 | Invalid Check Availability All Room", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.StartDate = time.Now().Add(time.Duration(-24) * time.Hour)
		expectedErr := errors.New("invalid date")

		result, actualErr := transactionUseCase.CheckAvailabilityAllRoom(transactionInput.StartDate, transactionInput.EndDate)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, []string{}, result)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Create Transaction", func(t *testing.T) {
		roomInput := roomDomain
		roomInput.RoomType = transactionDomain.RoomType
		roomInput.Room[0].Number = transactionDomain.RoomNumber
		roomInput.Room[0].Status = "available"
		roomRepository.On("GetRoomByType", transactionDomain.RoomType).Return(roomInput, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, nil).Once()
		transactionRepository.On("Create", transactionDomain.UserEmail, mock.Anything).Return(transactionDomain, nil).Once()

		result, actualErr := transactionUseCase.CreateTransaction(transactionDomain.UserEmail, transactionDomain)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Create Transaction", func(t *testing.T) {
		expectedErr := errors.New("")
		roomRepository.On("GetRoomByType", transactionDomain.RoomType).Return(roomDomain, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, nil).Once()
		transactionRepository.On("Create", transactionDomain.UserEmail, mock.Anything).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.CreateTransaction(transactionDomain.UserEmail, transactionDomain)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Create Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.RoomType = "not exist"

		expectedErr := errors.New("room type not registered")
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(rooms.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.CreateTransaction(transactionInput.UserEmail, transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Create Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.StartDate = time.Now().Add(time.Duration(-24) * time.Hour)

		expectedErr := errors.New("invalid date")
		result, actualErr := transactionUseCase.CreateTransaction(transactionInput.UserEmail, transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Create Transaction", func(t *testing.T) {
		roomInput := roomDomain
		roomInput.Room[0].Status = "not available"

		expectedErr := errors.New("room is not available")
		roomRepository.On("GetRoomByType", transactionDomain.RoomType).Return(roomInput, nil).Once()

		result, actualErr := transactionUseCase.CreateTransaction(transactionDomain.UserEmail, transactionDomain)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 6 | Invalid Create Transaction", func(t *testing.T) {
		roomDomain.Room[0].Status = "available"
		expectedErr := errors.New("room is not available")
		roomRepository.On("GetRoomByType", transactionDomain.RoomType).Return(roomDomain, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{transactionDomain}, expectedErr).Once()

		result, actualErr := transactionUseCase.CreateTransaction(transactionDomain.UserEmail, transactionDomain)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 7 | Invalid Create Transaction", func(t *testing.T) {
		expectedErr := errors.New("")
		roomRepository.On("GetRoomByType", transactionDomain.RoomType).Return(roomDomain, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.CreateTransaction(transactionDomain.UserEmail, transactionDomain)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestReceptionistCreateTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Receptionist Create Transaction", func(t *testing.T) {
		roomInput := roomDomain
		roomInput.RoomType = transactionDomain.RoomType
		roomInput.Room[0].Number = transactionDomain.RoomNumber
		roomInput.Room[0].Status = "available"

		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		transactionInput.RoomType = "test"
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(roomInput, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionInput.RoomType, mock.Anything, transactionInput.RoomNumber).Return([]transactions.Domain{}, nil).Once()
		transactionRepository.On("Create", transactionInput.UserEmail, mock.Anything).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transactionInput)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Receptionist Create Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.StartDate = time.Now().Add(time.Duration(-24) * time.Hour)
		expectedErr := errors.New("invalid date")

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Receptionist Create Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.RoomType = "not exist"
		transactionInput.Status = "verified"

		expectedErr := errors.New("room type not registered")
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(rooms.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Receptionist Create Transaction", func(t *testing.T) {
		roomInput := roomDomain
		roomInput.Room[0].Status = "not available"

		transactionInput := transactionDomain
		transactionInput.Status = "verified"

		expectedErr := errors.New("room is not available")
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(roomInput, nil).Once()

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Receptionist Create Transaction", func(t *testing.T) {
		transctionInput := transactionDomain
		transctionInput.Status = "not valid"
		expectedErr := errors.New("status must be verified or checked-in")

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transctionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 6 | Invalid Receptionist Create Transaction", func(t *testing.T) {
		transcationInput := transactionDomain
		transcationInput.Status = "verified"
		roomInput := roomDomain
		roomInput.Room[0].Status = "available"
		expectedErr := errors.New("")
		roomRepository.On("GetRoomByType", transcationInput.RoomType).Return(roomInput, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transcationInput.RoomType, mock.Anything, transcationInput.RoomNumber).Return([]transactions.Domain{}, nil).Once()
		transactionRepository.On("Create", transcationInput.UserEmail, mock.Anything).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transcationInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 7 | Invalid Receptionist Create Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		roomInput := roomDomain
		roomInput.Room[0].Status = "available"
		expectedErr := errors.New("")
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(roomInput, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionInput.RoomType, mock.Anything, transactionInput.RoomNumber).Return([]transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.ReceptionistCreateTransaction(transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestCancelTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Cancel Transaction", func(t *testing.T) {
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactionDomain, nil).Once()
		transactionRepository.On("Update", transactionDomain.TransactionID, mock.Anything).Return(nil).Once()

		actualErr := transactionUseCase.CancelTransaction(transactionDomain.TransactionID, transactionDomain.UserEmail)

		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Cancel Transaction", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		actualErr := transactionUseCase.CancelTransaction(transactionDomain.TransactionID, transactionDomain.UserEmail)

		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Cancel Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		expectedErr := errors.New("transaction is verified")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		actualErr := transactionUseCase.CancelTransaction(transactionInput.TransactionID, transactionInput.UserEmail)

		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 4 | Invalid Cancel Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		transactionInput.UserEmail = "invalid@gmail.com"
		actualErr := transactionUseCase.CancelTransaction(transactionInput.TransactionID, transactionInput.UserEmail)

		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 5 | Invalid Cancel Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(expectedErr).Once()

		actualErr := transactionUseCase.CancelTransaction(transactionInput.TransactionID, transactionInput.UserEmail)

		assert.Equal(t, expectedErr, actualErr)
	})
}

func TestUpdatePayment(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Payemnt", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(nil).Once()

		result, actualErr := transactionUseCase.UpdatePayment(transactionInput.TransactionID, transactionInput.UserEmail, transactionInput.Payment_URL)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Update Payemnt", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.UpdatePayment(transactionDomain.TransactionID, transactionDomain.UserEmail, transactionDomain.Payment_URL)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Update Payemnt", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		expectedErr := errors.New("transaction is verified")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdatePayment(transactionInput.TransactionID, transactionInput.UserEmail, transactionInput.Payment_URL)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Update Payemnt", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		transactionInput.UserEmail = "not same email"
		result, actualErr := transactionUseCase.UpdatePayment(transactionInput.TransactionID, transactionInput.UserEmail, transactionInput.Payment_URL)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Update Payemnt", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(expectedErr).Once()

		result, actualErr := transactionUseCase.UpdatePayment(transactionInput.TransactionID, transactionInput.UserEmail, transactionInput.Payment_URL)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestUpdateVerification(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verification-pending"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionInput.Status = "verified"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(nil).Once()

		result, actualErr := transactionUseCase.UpdateVerification(transactionInput.TransactionID, transactionInput.Status)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Update Verification", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.UpdateVerification(transactionDomain.TransactionID, transactionDomain.Status)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Update Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("transaction is unpaid")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdateVerification(transactionInput.TransactionID, transactionInput.Status)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Update Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verification-pending"
		expectedErr := errors.New("invalid status")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdateVerification(transactionInput.TransactionID, transactionInput.Status)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Update Verification", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verification-pending"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionInput.Status = "verified"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(expectedErr).Once()

		result, actualErr := transactionUseCase.UpdateVerification(transactionInput.TransactionID, transactionInput.Status)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestUpdateCheckIn(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Check In", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionInput.Status = "checked-in"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(nil).Once()

		result, actualErr := transactionUseCase.UpdateCheckIn(transactionInput.TransactionID)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Update Check In", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.UpdateCheckIn(transactionDomain.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Update Check In", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("transaction is unpaid")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdateCheckIn(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Update Check In", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		expectedErr := errors.New("transaction is checked-in")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdateCheckIn(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Update Check In", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "verified"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionInput.Status = "checked-in"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(expectedErr).Once()

		result, actualErr := transactionUseCase.UpdateCheckIn(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestUpdateCheckOut(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Check Out", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionInput.Status = "checked-out"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(nil).Once()

		result, actualErr := transactionUseCase.UpdateCheckOut(transactionInput.TransactionID)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Update Check Out", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionRepository.On("GetTransactionByID", transactionDomain.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.UpdateCheckOut(transactionDomain.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Update Check Out", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "unpaid"
		expectedErr := errors.New("transaction is unpaid")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdateCheckOut(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Update Check Out", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-out"
		expectedErr := errors.New("transaction is checked-out")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()

		result, actualErr := transactionUseCase.UpdateCheckOut(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Update Check Out", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionInput.Status = "checked-out"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(expectedErr).Once()

		result, actualErr := transactionUseCase.UpdateCheckOut(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestAdminUpdateTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Admin Update Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		transactionInput.RoomType = "test"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(roomDomain, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, nil).Once()
		transactionInput.Status = "done"
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(nil).Once()

		result, actualErr := transactionUseCase.AdminUpdateTransaction(transactionInput.TransactionID, transactionInput)

		assert.Nil(t, actualErr)
		assert.NotNil(t, result)
	})

	t.Run("Test Case 2 | Invalid Admin Update Transaction", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.AdminUpdateTransaction(transactionInput.TransactionID, transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 3 | Invalid Admin Update Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		expectedErr := errors.New("room not found")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(rooms.Domain{}, expectedErr).Once()

		result, actualErr := transactionUseCase.AdminUpdateTransaction(transactionInput.TransactionID, transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 4 | Invalid Admin Update Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		expectedErr := errors.New("room is not available")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(roomDomain, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{transactionDomain}, nil).Once()

		result, actualErr := transactionUseCase.AdminUpdateTransaction(transactionInput.TransactionID, transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})

	t.Run("Test Case 5 | Invalid Admin Update Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		roomRepository.On("GetRoomByType", transactionInput.RoomType).Return(roomDomain, nil).Once()
		transactionRepository.On("GetTransactionByRoomAndEndDate", transactionDomain.RoomType, mock.Anything, transactionDomain.RoomNumber).Return([]transactions.Domain{}, nil).Once()
		transactionRepository.On("Update", transactionInput.TransactionID, mock.Anything).Return(expectedErr).Once()

		result, actualErr := transactionUseCase.AdminUpdateTransaction(transactionInput.TransactionID, transactionInput)

		assert.Equal(t, expectedErr, actualErr)
		assert.Equal(t, transactions.Domain{}, result)
	})
}

func TestAdminDeleteTransaction(t *testing.T) {
	t.Run("Test Case 1 | Valid Admin Delete Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionRepository.On("Delete", transactionInput.TransactionID).Return(nil).Once()

		actualErr := transactionUseCase.AdminDeleteTransaction(transactionInput.TransactionID)

		assert.Nil(t, actualErr)
	})

	t.Run("Test Case 2 | Invalid Admin Delete Transaction", func(t *testing.T) {
		expectedErr := errors.New("transaction not found")
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactions.Domain{}, expectedErr).Once()

		actualErr := transactionUseCase.AdminDeleteTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("Test Case 3 | Invalid Admin Delete Transaction", func(t *testing.T) {
		transactionInput := transactionDomain
		transactionInput.Status = "checked-in"
		expectedErr := errors.New("")
		transactionRepository.On("GetTransactionByID", transactionInput.TransactionID).Return(transactionInput, nil).Once()
		transactionRepository.On("Delete", transactionInput.TransactionID).Return(expectedErr).Once()

		actualErr := transactionUseCase.AdminDeleteTransaction(transactionInput.TransactionID)

		assert.Equal(t, expectedErr, actualErr)
	})
}
