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
		return Domain{}, fmt.Errorf("transaction is on %s", transaction.Status)
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

func (tu *TransactionUseCase) CreateTransaction(email string, transactionInput Domain) (Domain, error) {
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
		// start date between input end date and input start date
		if transaction.StartDate.Before(transactionInput.EndDate) && transaction.StartDate.After(transactionInput.StartDate) {
			return Domain{}, errors.New("room is not available")
		}

		// end date between input end date and input start date
		if transaction.EndDate.Before(transactionInput.EndDate) && transaction.EndDate.After(transactionInput.StartDate) {
			return Domain{}, errors.New("room is not available")
		}
	}

	transactionInput.UserEmail = email
	transactionInput.StartDate = time.Date(transactionInput.StartDate.Year(), transactionInput.StartDate.Month(), transactionInput.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	transactionInput.EndDate = time.Date(transactionInput.EndDate.Year(), transactionInput.EndDate.Month(), transactionInput.EndDate.Day(), 0, 0, 0, 0, time.UTC)
	transactionInput.Status = "unpaid"
	transactionInput.Bill = RoomData.Price * int(transactionInput.EndDate.Sub(transactionInput.StartDate).Hours()/24)
	transactionInput.CreatedAt = time.Now()
	transactionInput.UpdatedAt = transactionInput.CreatedAt

	timeToString, _ := transactionInput.CreatedAt.UTC().MarshalText()
	transactionInput.TransactionID = fmt.Sprintf("%s_%s", timeToString, email)

	transaction, err := tu.transactionRepository.Create(email, transactionInput, RoomData)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
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
		return Domain{}, fmt.Errorf("transaction is on %s", transaction.Status)
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
