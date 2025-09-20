package handlers

import (
	"net/http"
	"santimpay-api/models"
	"santimpay-api/services"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	santimPayService *services.SantimPayService
}

func NewPaymentHandler(santimPayService *services.SantimPayService) *PaymentHandler {
	return &PaymentHandler{
		santimPayService: santimPayService,
	}
}

func (h *PaymentHandler) InitiatePayment(c *gin.Context) {
	var req models.PaymentInitiateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	paymentURL, err := h.santimPayService.GeneratePaymentURL(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to initiate payment",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Payment initiated successfully",
		"data": gin.H{
			"paymentUrl": paymentURL,
		},
	})
}

func (h *PaymentHandler) DirectPayment(c *gin.Context) {
	var req models.DirectPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.santimPayService.DirectPayment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to process direct payment",
			Error:   err.Error(),
		})
		return
	}

	status := http.StatusOK
	if !response.Success {
		status = http.StatusBadRequest
	}

	c.JSON(status, response)
}

func (h *PaymentHandler) PayoutTransfer(c *gin.Context) {
	var req models.PayoutTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.santimPayService.SendToCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to process payout transfer",
			Error:   err.Error(),
		})
		return
	}

	status := http.StatusOK
	if !response.Success {
		status = http.StatusBadRequest
	}

	c.JSON(status, response)
}

func (h *PaymentHandler) CheckTransactionStatus(c *gin.Context) {
	var req models.TransactionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.santimPayService.CheckTransactionStatus(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to fetch transaction status",
			Error:   err.Error(),
		})
		return
	}

	status := http.StatusOK
	if !response.Success {
		status = http.StatusBadRequest
	}

	c.JSON(status, response)
}

func (h *PaymentHandler) WebhookHandler(c *gin.Context) {
	var webhook map[string]interface{}
	if err := c.ShouldBindJSON(&webhook); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid webhook data",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Webhook received successfully",
	})
}