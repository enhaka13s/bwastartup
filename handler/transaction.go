package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get campaign's transaction: error to bind URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Error to get campaign's transaction: error to call service.GetTransactionsByCampaignID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign's transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Error to get user's transaction: error to call service.GetTransactionsByUserID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get user's transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)

	if err != nil {
		response := helper.APIResponse("Failed to create transaction: Error to call service CreateTransaction()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create transaction", http.StatusCreated, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusCreated, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification: Failed to bind JSON to input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification: Failed to call service -> ProcessPayment()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	c.JSON(http.StatusOK, input)
}
