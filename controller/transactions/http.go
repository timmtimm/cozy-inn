package transactions

import (
	"cozy-inn/app/middleware"
	"cozy-inn/businesses/transactions"
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
