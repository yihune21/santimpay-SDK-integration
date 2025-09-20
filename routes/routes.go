package routes

import (
	"santimpay-api/handlers"
	"santimpay-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(paymentHandler *handlers.PaymentHandler) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.CORSMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Santim Pay API is running",
		})
	})

	api := router.Group("/api/v1")
	{
		payment := api.Group("/payment")
		{
			payment.POST("/initiate", paymentHandler.InitiatePayment)
			payment.POST("/direct", paymentHandler.DirectPayment)
			payment.POST("/payout", paymentHandler.PayoutTransfer)
			payment.POST("/status", paymentHandler.CheckTransactionStatus)
			payment.POST("/webhook", paymentHandler.WebhookHandler)
		}
	}

	return router
}