package transactions

import (
	"cozy-inn/businesses/rooms"
	"errors"
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
			return Domain{}, errors.New("room is not available")
		}
	}

	transactionList, err := tu.transactionRepository.GetFinishedTransactionByRoom(transactionDomain.RoomType, transactionDomain.StartDate)
	if err != nil {
		return Domain{}, err
	}

	for _, transaction := range transactionList {
		// start date between input end date and input start date
		if transaction.StartDate.Before(transactionDomain.EndDate) && transaction.StartDate.After(transactionDomain.StartDate) {
			return Domain{}, errors.New("room is not available")
		}

		// end date between input end date and input start date
		if transaction.EndDate.Before(transactionDomain.EndDate) && transaction.EndDate.After(transactionDomain.StartDate) {
			return Domain{}, errors.New("room is not available")
		}
	}

	transaction, err := tu.transactionRepository.CreateTransaction(email, transactionDomain, RoomData)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}
