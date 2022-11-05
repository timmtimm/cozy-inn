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

func (transactionCtrl *TransactionController) UpdatePayment(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	userInput := request.Payment{}
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

	transaction, err := transactionCtrl.transactionUseCase.UpdatePayment(transactionID, userInput.Payment_URL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update payment",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) GetPaymentNotVerified(c echo.Context) error {
	transactions, err := transactionCtrl.transactionUseCase.GetPaymentNotVerified()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment not verified",
		"data":    transactions,
	})
}

func (transactionCtrl *TransactionController) GetTransactionOnVerification(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	transaction, err := transactionCtrl.transactionUseCase.GetTransactionOnVerification(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment on verification list",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) UpdateVerification(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	userInput := request.Verification{}
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

	transaction, err := transactionCtrl.transactionUseCase.UpdateVerification(transactionID, userInput.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update verification",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) GetAllCheckIn(c echo.Context) error {
	transactions, err := transactionCtrl.transactionUseCase.GetAllCheckIn()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all check in",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all check in",
		"data":    transactions,
	})
}

func (transactionCtrl *TransactionController) GetCheckInTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	transaction, err := transactionCtrl.transactionUseCase.GetCheckInTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment on check in list",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) CheckInTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	transaction, err := transactionCtrl.transactionUseCase.UpdateCheckIn(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update check in",
		"data":    transaction,
	})
}
