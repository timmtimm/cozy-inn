package transactions

import (
	"cozy-inn/businesses/rooms"
	"errors"
	"fmt"
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

func (tu *TransactionUseCase) GetAllTransaction(email string) ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllTransaction(email)

	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) CreateTransaction(email string, transactionDomain *Domain) (Domain, error) {
	RoomData, err := tu.RoomRepository.GetRoomByType(transactionDomain.RoomType)
	if err != nil {
		return Domain{}, err
	}

	for _, room := range RoomData.Room {
		if room.Number == transactionDomain.RoomNumber && room.Status != "available" {
			return Domain{}, errors.New("room is not available1")
		}
	}

	transactionList, err := tu.transactionRepository.GetTransactionByRoomAndDate(
		transactionDomain.RoomType,
		transactionDomain.StartDate,
		transactionDomain.RoomNumber)
	if err != nil {
		return Domain{}, err
	}

	for _, transaction := range transactionList {
		// start date between input end date and input start date
		fmt.Println(transaction)
		if transaction.StartDate.Before(transactionDomain.EndDate) && transaction.StartDate.After(transactionDomain.StartDate) {
			return Domain{}, errors.New("room is not available2")
		}

		// end date between input end date and input start date
		if transaction.EndDate.Before(transactionDomain.EndDate) && transaction.EndDate.After(transactionDomain.StartDate) {
			return Domain{}, errors.New("room is not available3")
		}
	}

	transaction, err := tu.transactionRepository.CreateTransaction(email, transactionDomain, RoomData)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) UpdatePayment(transactionID string, payment_URL string) (Domain, error) {
	check, err := tu.transactionRepository.GetTransactionOnVerification(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if check.Status != "verification-pending" {
		return Domain{}, fmt.Errorf("transaction is %s", check.Status)
	}

	transaction, err := tu.transactionRepository.UpdatePayment(transactionID, payment_URL)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) GetPaymentNotVerified() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetPaymentNotVerified()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) GetTransactionOnVerification(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetTransactionOnVerification(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "verification-pending" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	return transaction, nil
}

func (tu *TransactionUseCase) UpdateVerification(transactionID string, status string) (Domain, error) {
	check, err := tu.transactionRepository.GetTransactionOnVerification(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if check.Status != "verification-pending" {
		return Domain{}, fmt.Errorf("transaction is %s", check.Status)
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

	transaction, err := tu.transactionRepository.UpdateVerification(transactionID, status)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) GetAllCheckIn() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllCheckIn()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}

func (tu *TransactionUseCase) GetCheckInTransaction(transactionID string) (Domain, error) {
	transaction, err := tu.transactionRepository.GetCheckInTransaction(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if transaction.Status != "verified" {
		return Domain{}, fmt.Errorf("transaction is %s", transaction.Status)
	}

	return transaction, nil
}

func (tu *TransactionUseCase) UpdateCheckIn(transactionID string) (Domain, error) {
	check, err := tu.transactionRepository.GetCheckInTransaction(transactionID)
	if err != nil {
		return Domain{}, err
	}

	if check.Status != "verified" {
		return Domain{}, fmt.Errorf("transaction is %s", check.Status)
	}

	transaction, err := tu.transactionRepository.UpdateCheckIn(transactionID)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (tu *TransactionUseCase) GetAllCheckOut() ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllCheckOut()
	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
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
