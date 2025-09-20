package models

import "time"

type PaymentInitiateRequest struct {
	ID                  string  `json:"id" binding:"required"`
	Amount              float64 `json:"amount" binding:"required"`
	PaymentReason       string  `json:"paymentReason" binding:"required"`
	SuccessRedirectURL  string  `json:"successRedirectUrl" binding:"required"`
	FailureRedirectURL  string  `json:"failureRedirectUrl" binding:"required"`
	NotifyURL           string  `json:"notifyUrl" binding:"required"`
	PhoneNumber         string  `json:"phoneNumber,omitempty"`
	CancelRedirectURL   string  `json:"cancelRedirectUrl,omitempty"`
}

type DirectPaymentRequest struct {
	ID            string  `json:"id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	PaymentReason string  `json:"paymentReason" binding:"required"`
	NotifyURL     string  `json:"notifyUrl" binding:"required"`
	PhoneNumber   string  `json:"phoneNumber" binding:"required"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
}

type PayoutTransferRequest struct {
	ID            string  `json:"id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	PaymentReason string  `json:"paymentReason" binding:"required"`
	PhoneNumber   string  `json:"phoneNumber" binding:"required"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
	NotifyURL     string  `json:"notifyUrl" binding:"required"`
}

type TransactionStatusRequest struct {
	ID string `json:"id" binding:"required"`
}

type PaymentInitiatePayload struct {
	ID                  string  `json:"id"`
	Amount              float64 `json:"amount"`
	Reason              string  `json:"reason"`
	MerchantID          string  `json:"merchantId"`
	SignedToken         string  `json:"signedToken"`
	SuccessRedirectURL  string  `json:"successRedirectUrl"`
	FailureRedirectURL  string  `json:"failureRedirectUrl"`
	NotifyURL           string  `json:"notifyUrl"`
	CancelRedirectURL   string  `json:"cancelRedirectUrl,omitempty"`
	PhoneNumber         string  `json:"phoneNumber,omitempty"`
}

type DirectPaymentPayload struct {
	ID            string  `json:"id"`
	Amount        float64 `json:"amount"`
	Reason        string  `json:"reason"`
	MerchantID    string  `json:"merchantId"`
	SignedToken   string  `json:"signedToken"`
	PhoneNumber   string  `json:"phoneNumber"`
	PaymentMethod string  `json:"paymentMethod"`
	NotifyURL     string  `json:"notifyUrl"`
}

type PayoutTransferPayload struct {
	ID                     string  `json:"id"`
	ClientReference        string  `json:"clientReference"`
	Amount                 float64 `json:"amount"`
	Reason                 string  `json:"reason"`
	MerchantID             string  `json:"merchantId"`
	SignedToken            string  `json:"signedToken"`
	ReceiverAccountNumber  string  `json:"receiverAccountNumber"`
	NotifyURL              string  `json:"notifyUrl"`
	PaymentMethod          string  `json:"paymentMethod"`
}

type TransactionStatusPayload struct {
	ID          string `json:"id"`
	MerchantID  string `json:"merchantId"`
	SignedToken string `json:"signedToken"`
}

type TokenPayload struct {
	Amount        float64 `json:"amount"`
	PaymentReason string  `json:"paymentReason"`
	MerchantID    string  `json:"merchantId"`
	Generated     int64   `json:"generated"`
	PaymentMethod string  `json:"paymentMethod,omitempty"`
	PhoneNumber   string  `json:"phoneNumber,omitempty"`
	ID            string  `json:"id,omitempty"`
	MerID         string  `json:"merId,omitempty"`
}

type PaymentResponse struct {
	URL string `json:"url"`
}

type DirectPaymentResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type TransactionStatusResponse struct {
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	Transaction Transaction `json:"transaction,omitempty"`
}

type Transaction struct {
	ID            string    `json:"id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"paymentMethod"`
	PhoneNumber   string    `json:"phoneNumber"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}