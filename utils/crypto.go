package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignES256(payload map[string]interface{}, privateKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("failed to parse private key: %v", err)
		}
		var ok bool
		privateKey, ok = parsedKey.(*ecdsa.PrivateKey)
		if !ok {
			return "", errors.New("private key is not ECDSA")
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(payload))
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}

func GenerateTokenPayload(amount float64, paymentReason, merchantID string) map[string]interface{} {
	return map[string]interface{}{
		"amount":        amount,
		"paymentReason": paymentReason,
		"merchantId":    merchantID,
		"generated":     time.Now().Unix(),
	}
}

func GenerateDirectPaymentTokenPayload(amount float64, paymentReason, merchantID, paymentMethod, phoneNumber string) map[string]interface{} {
	return map[string]interface{}{
		"amount":        amount,
		"paymentReason": paymentReason,
		"paymentMethod": paymentMethod,
		"phoneNumber":   phoneNumber,
		"merchantId":    merchantID,
		"generated":     time.Now().Unix(),
	}
}

func GenerateTransactionTokenPayload(id, merchantID string) map[string]interface{} {
	return map[string]interface{}{
		"id":        id,
		"merId":     merchantID,
		"generated": time.Now().Unix(),
	}
}