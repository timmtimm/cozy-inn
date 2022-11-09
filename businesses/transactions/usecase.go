package transactions

import (
	"cozy-inn/businesses/rooms"
	"errors"
	"fmt"
	"time"
)

type TransactionUseCase struct {
	transactionRepository Repository
	RoomRepository        rooms.Repository
}

func NewTransactionUsecase(rr Repository, RoomRepository rooms.Repository) UseCase {
	return &TransactionUseCase{
		transactionRepository: rr,
		RoomRepository:        RoomRepository,
	}
}

func (tu *TransactionUseCase) GetAllTransactionUser(email string) ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllTransactionByEmail(email)
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) GetAllReadyCheckIn() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllReadyCheckIn()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) GetAllReadyCheckOut() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllReadyCheckOut()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) GetAllPaymentNotVerified() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllPaymentNotVerified()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) AdminGetAllTransaction() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllTransaction()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) GetTransactionOnVerification(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "verification-pending" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	return transaction, nil
}

func (tu *TransactionUseCase) GetCheckInTransaction(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "verified" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	return transaction, nil
}

func (tu *TransactionUseCase) GetCheckOutTransaction(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "checked-in" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	return transaction, nil
}

func (tu *TransactionUseCase) GetTransaction(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) CheckAvailabilityAllRoom(startDate time.Time, endDate time.Time) ([]string, error) {
	checkDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 23, 59, 59, 59, time.UTC)
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)

	if startDate.Before(checkDate) || endDate.Before(checkDate) || startDate.After(endDate) {
		return []string{}, errors.New("invalid date")
	}

	roomTypes, err := tu.RoomRepository.GetAllRoom()
	if err != nil {
		return []string{}, err
	}

	availableRooms := []string{}
	for _, roomType := range roomTypes {
		for _, room := range roomType.Room {
			if room.Status == "available" {
				transactions, err := tu.transactionRepository.GetTransactionByRoomAndEndDate(
					roomType.RoomType,
					startDate,
					room.Number)
				if err != nil {
					return []string{}, err
				}

				if len(transactions) == 0 {
					availableRooms = append(availableRooms, fmt.Sprintf("%s-%d", roomType.RoomType, room.Number))
				} else {
					available := true
					for _, transaction := range transactions {
						if (transaction.StartDate == startDate && transaction.EndDate == endDate) ||
							(transaction.StartDate.Before(endDate) && transaction.StartDate.After(startDate)) ||
							(transaction.EndDate.Before(endDate) && transaction.EndDate.After(startDate)) {
							available = false
						}
					}

					if available {
						availableRooms = append(availableRooms, fmt.Sprintf("%s-%d", roomType.RoomType, room.Number))
					}
				}
			}
		}
	}

	return availableRooms, nil
}

func (tu *TransactionUseCase) CreateTransaction(email string, transactionInput Domain) (Domain, error) {
	checkDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 23, 59, 59, 59, time.UTC)
	transactionInput.StartDate = time.Date(transactionInput.StartDate.Year(), transactionInput.StartDate.Month(), transactionInput.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	transactionInput.EndDate = time.Date(transactionInput.EndDate.Year(), transactionInput.EndDate.Month(), transactionInput.EndDate.Day(), 0, 0, 0, 0, time.UTC)

	if transactionInput.StartDate.Before(checkDate) || transactionInput.EndDate.Before(checkDate) || transactionInput.StartDate.After(transactionInput.EndDate) {
		return Domain{}, errors.New("invalid date")
	}

	RoomData, err := tu.RoomRepository.GetRoomByType(transactionInput.RoomType)
	if err != nil {
		return Domain{}, err
	}

	for _, room := range RoomData.Room {
		if room.Number == transactionInput.RoomNumber && room.Status != "available" {
			return Domain{}, errors.New("room is not available")
		}
	}

	transactionList, err := tu.transactionRepository.GetTransactionByRoomAndEndDate(
		transactionInput.RoomType,
		transactionInput.StartDate,
		transactionInput.RoomNumber)
	if err != nil {
		return Domain{}, err
	}

	for _, transaction := range transactionList {
		if (transaction.StartDate == transactionInput.StartDate && transaction.EndDate == transactionInput.EndDate) ||
			(transaction.StartDate.Before(transactionInput.EndDate) && transaction.StartDate.After(transactionInput.StartDate)) ||
			(transaction.EndDate.Before(transactionInput.EndDate) && transaction.EndDate.After(transactionInput.StartDate)) {
			return Domain{}, errors.New("room is not available")
		}
	}

	transactionInput.UserEmail = email
	transactionInput.Status = "unpaid"
	transactionInput.Bill = RoomData.Price * int(transactionInput.EndDate.Sub(transactionInput.StartDate).Hours()/24)
	transactionInput.CreatedAt = time.Now()
	transactionInput.UpdatedAt = transactionInput.CreatedAt

	timeToString, _ := transactionInput.CreatedAt.UTC().MarshalText()
	transactionInput.TransactionID = fmt.Sprintf("%s_%s", timeToString, email)

	transaction, err := tu.transactionRepository.Create(email, transactionInput)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) ReceptionistCreateTransaction(transactionInput Domain) (Domain, error) {
	checkDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 23, 59, 59, 59, time.UTC)
	transactionInput.StartDate = time.Date(transactionInput.StartDate.Year(), transactionInput.StartDate.Month(), transactionInput.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	transactionInput.EndDate = time.Date(transactionInput.EndDate.Year(), transactionInput.EndDate.Month(), transactionInput.EndDate.Day(), 0, 0, 0, 0, time.UTC)

	if transactionInput.StartDate.Before(checkDate) || transactionInput.EndDate.Before(checkDate) || transactionInput.StartDate.After(transactionInput.EndDate) {
		return Domain{}, errors.New("invalid date")
	}

	avaliableStatus := []string{"verified", "checked-in"}
	statusFound := false
	for _, status := range avaliableStatus {
		if transactionInput.Status == status {
			statusFound = true
			break
		}
	}

	if !statusFound {
		return Domain{}, errors.New("status must be verified or checked-in")
	}

	RoomData, err := tu.RoomRepository.GetRoomByType(transactionInput.RoomType)
	if err != nil {
		return Domain{}, err
	}

	for _, room := range RoomData.Room {
		if room.Number == transactionInput.RoomNumber && room.Status != "available" {
			fmt.Println("room number", room.Number, "room status", room.Status)
			return Domain{}, errors.New("room is not available")
		}
	}

	transactionList, err := tu.transactionRepository.GetTransactionByRoomAndEndDate(
		transactionInput.RoomType,
		transactionInput.StartDate,
		transactionInput.RoomNumber)
	if err != nil {
		return Domain{}, err
	}

	for _, transaction := range transactionList {
		if (transaction.StartDate == transactionInput.StartDate && transaction.EndDate == transactionInput.EndDate) ||
			(transaction.StartDate.Before(transactionInput.EndDate) && transaction.StartDate.After(transactionInput.StartDate)) ||
			(transaction.EndDate.Before(transactionInput.EndDate) && transaction.EndDate.After(transactionInput.StartDate)) {
			return Domain{}, errors.New("room is not available")
		}
	}

	transactionInput.Bill = RoomData.Price * int(transactionInput.EndDate.Sub(transactionInput.StartDate).Hours()/24)
	transactionInput.CreatedAt = time.Now()
	transactionInput.UpdatedAt = transactionInput.CreatedAt

	timeToString, _ := transactionInput.CreatedAt.UTC().MarshalText()
	transactionInput.TransactionID = fmt.Sprintf("%s_%s", timeToString, transactionInput.UserEmail)

	transaction, err := tu.transactionRepository.Create(transactionInput.UserEmail, transactionInput)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) CancelTransaction(transactionID string, email string) error {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return err
	}

	if transaction.Status == "verified" || transaction.Status == "checked-in" ||
		transaction.Status == "checked-out" || transaction.Status == "canceled" {
		return fmt.Errorf("transaction is %s", transaction.Status)
	}

	if transaction.UserEmail != email {
		return errors.New("transaction not found")
	}

	transaction.Status = "canceled"

	err = tu.transactionRepository.Update(transactionID, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TransactionUseCase) UpdatePayment(transactionID string, email string, payment_URL string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if email != transaction.UserEmail {
		return Domain{}, errors.New("transaction not found")
	}

	if transaction.Status != "unpaid" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	transaction.Status = "verification-pending"
	transaction.Payment_URL = payment_URL
	transaction.UpdatedAt = time.Now()

	err = tu.transactionRepository.Update(transactionID, transaction)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) UpdateVerification(transactionID string, status string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "verification-pending" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	statusList := []string{"verified", "rejected"}
	isValid := false
	for _, s := range statusList {
		if s == status {
			isValid = true
		}
	}

	if !isValid {
		return Domain{}, errors.New("invalid status")
	}

	transaction.Status = status
	transaction.UpdatedAt = time.Now()

	err = tu.transactionRepository.Update(transactionID, transaction)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) UpdateCheckIn(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "verified" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	transaction.Status = "checked-in"
	transaction.CheckIn = time.Now()
	transaction.UpdatedAt = transaction.CheckIn

	err = tu.transactionRepository.Update(transactionID, transaction)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) UpdateCheckOut(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "checked-in" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	transaction.Status = "done"
	transaction.CheckOut = time.Now()
	transaction.UpdatedAt = transaction.CheckOut

	err = tu.transactionRepository.Update(transactionID, transaction)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) AdminUpdateTransaction(transactionID string, userInput Domain) (Domain, error) {
	avaliableStatus := []string{"unpaid", "verification-pending", "canceled", "verified", "checked-in", "checked-out", "done"}
	statusFound := false
	for _, status := range avaliableStatus {
		if status == userInput.Status {
			statusFound = true
			break
		}
	}

	if !statusFound {
		return Domain{}, errors.New("invalid status")
	}

	transaction, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return Domain{}, err
	}

	roomType, err := tu.RoomRepository.GetRoomByType(userInput.RoomType)
	if err != nil {
		return Domain{}, err
	}

	roomFound := false
	for _, room := range roomType.Room {
		if room.Number == userInput.RoomNumber {
			roomFound = true
			if room.Status != "available" {
				return Domain{}, errors.New("room is not available")
			}
			break
		}
	}

	if !roomFound {
		return Domain{}, errors.New("room not exist")
	}

	transactions, err := tu.transactionRepository.GetTransactionByRoomAndEndDate(
		roomType.RoomType,
		userInput.StartDate,
		userInput.RoomNumber)
	if err != nil {
		return Domain{}, err
	}

	for _, transaction := range transactions {
		if (transaction.StartDate == userInput.StartDate && transaction.EndDate == userInput.EndDate) ||
			(transaction.StartDate.Before(userInput.EndDate) && transaction.StartDate.After(userInput.StartDate)) ||
			(transaction.EndDate.Before(userInput.EndDate) && transaction.EndDate.After(userInput.StartDate)) {
			return Domain{}, errors.New("room is not available")
		}
	}

	transaction.RoomType = userInput.RoomType
	transaction.RoomNumber = userInput.RoomNumber
	transaction.StartDate = userInput.StartDate
	transaction.EndDate = userInput.EndDate
	transaction.CheckIn = userInput.CheckIn
	transaction.CheckOut = userInput.CheckOut
	transaction.Status = userInput.Status
	transaction.Bill = roomType.Price * int(userInput.EndDate.Sub(userInput.StartDate).Hours()/24)
	transaction.UpdatedAt = time.Now()

	err = tu.transactionRepository.Update(transactionID, transaction)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) AdminDeleteTransaction(transactionID string) error {
	_, err := tu.transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return err
	}

	err = tu.transactionRepository.Delete(transactionID)
	if err != nil {
		return err
	}

	return nil
}
