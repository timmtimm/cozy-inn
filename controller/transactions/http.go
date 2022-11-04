package transactions

import (
	"cozy-inn/app/middleware"
	"cozy-inn/businesses/transactions"
	"cozy-inn/controller/transactions/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUseCase transactions.UseCase
}

func NewTransactionController(transactionUC transactions.UseCase) *TransactionController {
	return &TransactionController{
		transactionUseCase: transactionUC,
	}
}

func (transactionCtrl *TransactionController) GetAllTransaction(c echo.Context) error {
	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	transactions, err := transactionCtrl.transactionUseCase.GetAllTransaction(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all transaction",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all transaction",
		"data":    transactions,
	})
}

func (transactionCtrl *TransactionController) CreateTransaction(c echo.Context) error {
	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	userInput := request.Transaction{}
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if userInput.Validate() != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	transaction, err := transactionCtrl.transactionUseCase.CreateTransaction(email, userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create transaction",
		"data":    transaction,
	})
}