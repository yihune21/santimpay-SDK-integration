package main

import (
	"fmt"
	"log"
	"santimpay-api/config"
	"santimpay-api/handlers"
	"santimpay-api/routes"
	"santimpay-api/services"
)

func main() {
	cfg := config.Load()

	santimPayService := services.NewSantimPayService(
		cfg.MerchantID,
		cfg.PrivateKey,
		cfg.TestMode,
	)

	paymentHandler := handlers.NewPaymentHandler(santimPayService)

	router := routes.SetupRouter(paymentHandler)

	mode := "PRODUCTION"
	if cfg.TestMode {
		mode = "TEST"
	}

	log.Printf("🚀 Santim Pay API Server starting on port %s in %s mode", cfg.ServerPort, mode)
	
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}