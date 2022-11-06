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

func (transactionCtrl *TransactionController) CheckAvailabilityAllRoom(c echo.Context) error {
	userInput := request.CheckAvailability{}
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

	availableRooms, err := transactionCtrl.transactionUseCase.CheckAvailabilityAllRoom(userInput.StartDate, userInput.EndDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all available rooms",
		"data":    availableRooms,
	})
}

func (transactionCtrl *TransactionController) GetAllTransaction(c echo.Context) error {
	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	transactions, err := transactionCtrl.transactionUseCase.GetAllTransactionUser(email)
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

func (transactionCtrl *TransactionController) GetTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	transaction, err := transactionCtrl.transactionUseCase.GetTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get transaction",
		})
	}

	if transaction.UserEmail != email {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "transaction not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get transaction",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) UpdatePayment(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

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

	transaction, err := transactionCtrl.transactionUseCase.UpdatePayment(transactionID, email, userInput.Payment_URL)
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

func (transactionCtrl *TransactionController) CancelTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	email, err := middleware.GetEmailByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	err = transactionCtrl.transactionUseCase.CancelTransaction(transactionID, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success cancel transaction",
	})
}

func (transactionCtrl *TransactionController) AdminDelete(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	err := transactionCtrl.transactionUseCase.AdminDeleteTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaction",
	})
}

func (transactionCtrl *TransactionController) GetAllPaymentNotVerified(c echo.Context) error {
	transactions, err := transactionCtrl.transactionUseCase.GetAllPaymentNotVerified()
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

func (transactionCtrl *TransactionController) GetAllReadyCheckIn(c echo.Context) error {
	transactions, err := transactionCtrl.transactionUseCase.GetAllReadyCheckIn()
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

func (transactionCtrl *TransactionController) GetCheckIn(c echo.Context) error {
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

func (transactionCtrl *TransactionController) UpdateCheckIn(c echo.Context) error {
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

func (transactionCtrl *TransactionController) GetAllReadyCheckOut(c echo.Context) error {
	transactions, err := transactionCtrl.transactionUseCase.GetAllReadyCheckOut()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all check out",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all check out",
		"data":    transactions,
	})
}

func (transactionCtrl *TransactionController) GetCheckOut(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	transaction, err := transactionCtrl.transactionUseCase.GetCheckOutTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment on check out list",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) UpdateCheckOut(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	transaction, err := transactionCtrl.transactionUseCase.UpdateCheckOut(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update check out",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) AdminGetAllTransaction(c echo.Context) error {
	transactions, err := transactionCtrl.transactionUseCase.AdminGetAllTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all transaction",
		"data":    transactions,
	})
}

func (transactionCtrl *TransactionController) AdminGetTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	transaction, err := transactionCtrl.transactionUseCase.GetTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get transaction",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) AdminUpdateTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	userInput := request.Update{}
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

	transaction, err := transactionCtrl.transactionUseCase.AdminUpdateTransaction(transactionID, userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update transaction",
		"data":    transaction,
	})
}

func (transactionCtrl *TransactionController) AdminDeleteTransaction(c echo.Context) error {
	transactionID := c.Param("transaction-id")

	err := transactionCtrl.transactionUseCase.AdminDeleteTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaction",
	})
}

func (transactionCtrl *TransactionController) ReceptionistCreateTransaction(c echo.Context) error {
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

	transaction, err := transactionCtrl.transactionUseCase.ReceptionistCreateTransaction(userInput.ToDomain())
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
