package transactions

type TransactionUseCase struct {
	transactionRepository Repository
}

func NewTransactionUsecase(rr Repository) UseCase {
	return &TransactionUseCase{
		transactionRepository: rr,
	}
}

func (tu *TransactionUseCase) GetAllTransaction(email string) ([]Domain, error) {
	transactions, err := tu.transactionRepository.GetAllTransaction(email)

	if err != nil {
		return []Domain{}, err
	}

	return transactions, nil
}
