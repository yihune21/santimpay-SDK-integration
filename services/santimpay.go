package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"santimpay-api/models"
	"santimpay-api/utils"
)

const (
	ProductionBaseURL = "https://services.santimpay.com/api/v1/gateway"
	TestBaseURL       = "https://testnet.santimpay.com/api/v1/gateway"
)

type SantimPayService struct {
	MerchantID string
	PrivateKey string
	BaseURL    string
	TestMode   bool
}

func NewSantimPayService(merchantID, privateKey string, testMode bool) *SantimPayService {
	baseURL := ProductionBaseURL
	if testMode {
		baseURL = TestBaseURL
	}

	return &SantimPayService{
		MerchantID: merchantID,
		PrivateKey: privateKey,
		BaseURL:    baseURL,
		TestMode:   testMode,
	}
}

func (s *SantimPayService) GeneratePaymentURL(req models.PaymentInitiateRequest) (string, error) {
	tokenPayload := utils.GenerateTokenPayload(req.Amount, req.PaymentReason, s.MerchantID)
	token, err := utils.SignES256(tokenPayload, s.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	payload := models.PaymentInitiatePayload{
		ID:                  req.ID,
		Amount:              req.Amount,
		Reason:              req.PaymentReason,
		MerchantID:          s.MerchantID,
		SignedToken:         token,
		SuccessRedirectURL:  req.SuccessRedirectURL,
		FailureRedirectURL:  req.FailureRedirectURL,
		NotifyURL:           req.NotifyURL,
		CancelRedirectURL:   req.CancelRedirectURL,
	}

	if req.PhoneNumber != "" {
		payload.PhoneNumber = req.PhoneNumber
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := http.Post(s.BaseURL+"/initiate-payment", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to initiate payment: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResp models.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return "", fmt.Errorf("payment initiation failed with status: %d", resp.StatusCode)
		}
		return "", fmt.Errorf("payment initiation failed: %s", errorResp.Message)
	}

	var paymentResp models.PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return paymentResp.URL, nil
}

func (s *SantimPayService) DirectPayment(req models.DirectPaymentRequest) (*models.DirectPaymentResponse, error) {
	tokenPayload := utils.GenerateDirectPaymentTokenPayload(
		req.Amount, 
		req.PaymentReason, 
		s.MerchantID, 
		req.PaymentMethod, 
		req.PhoneNumber,
	)
	token, err := utils.SignES256(tokenPayload, s.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	payload := models.DirectPaymentPayload{
		ID:            req.ID,
		Amount:        req.Amount,
		Reason:        req.PaymentReason,
		MerchantID:    s.MerchantID,
		SignedToken:   token,
		PhoneNumber:   req.PhoneNumber,
		PaymentMethod: req.PaymentMethod,
		NotifyURL:     req.NotifyURL,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := http.Post(s.BaseURL+"/direct-payment", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to execute direct payment: %v", err)
	}
	defer resp.Body.Close()

	var directPaymentResp models.DirectPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&directPaymentResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &directPaymentResp, fmt.Errorf("direct payment failed: %s", directPaymentResp.Message)
	}

	return &directPaymentResp, nil
}

func (s *SantimPayService) SendToCustomer(req models.PayoutTransferRequest) (*models.DirectPaymentResponse, error) {
	tokenPayload := utils.GenerateDirectPaymentTokenPayload(
		req.Amount, 
		req.PaymentReason, 
		s.MerchantID, 
		req.PaymentMethod, 
		req.PhoneNumber,
	)
	token, err := utils.SignES256(tokenPayload, s.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	payload := models.PayoutTransferPayload{
		ID:                    req.ID,
		ClientReference:       req.ID,
		Amount:                req.Amount,
		Reason:                req.PaymentReason,
		MerchantID:            s.MerchantID,
		SignedToken:           token,
		ReceiverAccountNumber: req.PhoneNumber,
		NotifyURL:             req.NotifyURL,
		PaymentMethod:         req.PaymentMethod,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := http.Post(s.BaseURL+"/payout-transfer", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to execute payout transfer: %v", err)
	}
	defer resp.Body.Close()

	var payoutResp models.DirectPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&payoutResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &payoutResp, fmt.Errorf("payout transfer failed: %s", payoutResp.Message)
	}

	return &payoutResp, nil
}

func (s *SantimPayService) CheckTransactionStatus(id string) (*models.TransactionStatusResponse, error) {
	tokenPayload := utils.GenerateTransactionTokenPayload(id, s.MerchantID)
	token, err := utils.SignES256(tokenPayload, s.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	payload := models.TransactionStatusPayload{
		ID:          id,
		MerchantID:  s.MerchantID,
		SignedToken: token,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := http.Post(s.BaseURL+"/fetch-transaction-status", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction status: %v", err)
	}
	defer resp.Body.Close()

	var statusResp models.TransactionStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &statusResp, fmt.Errorf("failed to fetch transaction status: %s", statusResp.Message)
	}

	return &statusResp, nil
}