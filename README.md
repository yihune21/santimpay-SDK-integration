# Santim Pay Go API

A custom Go API for integrating with Santim Pay payment gateway.

## Features

- Payment initiation with redirect URLs
- Direct payment processing
- Payout transfers (B2C)
- Transaction status checking
- Webhook handling
- CORS support
- Request logging

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file based on `.env.example`:
```bash
cp .env.example .env
```

4. Update the `.env` file with your Santim Pay credentials:
   - `SANTIMPAY_MERCHANT_ID`: Your merchant ID
   - `SANTIMPAY_PRIVATE_KEY`: Your EC private key for signing tokens
   - `SANTIMPAY_TEST_MODE`: Set to `true` for test environment, `false` for production

## Running the API

```bash
go run main.go
```

The server will start on port 8080 by default (configurable via `SERVER_PORT` in `.env`).

## API Endpoints

### Health Check
- **GET** `/health` - Check if the API is running

### Payment Operations

#### Initiate Payment
- **POST** `/api/v1/payment/initiate`
```json
{
  "id": "unique_transaction_id",
  "amount": 100.00,
  "paymentReason": "Product purchase",
  "successRedirectUrl": "https://yoursite.com/success",
  "failureRedirectUrl": "https://yoursite.com/failure",
  "notifyUrl": "https://yoursite.com/webhook",
  "phoneNumber": "+251911234567",
  "cancelRedirectUrl": "https://yoursite.com/cancel"
}
```

#### Direct Payment
- **POST** `/api/v1/payment/direct`
```json
{
  "id": "unique_transaction_id",
  "amount": 100.00,
  "paymentReason": "Product purchase",
  "notifyUrl": "https://yoursite.com/webhook",
  "phoneNumber": "+251911234567",
  "paymentMethod": "telebirr"
}
```

#### Payout Transfer (B2C)
- **POST** `/api/v1/payment/payout`
```json
{
  "id": "unique_transaction_id",
  "amount": 100.00,
  "paymentReason": "Refund",
  "phoneNumber": "+251911234567",
  "paymentMethod": "telebirr",
  "notifyUrl": "https://yoursite.com/webhook"
}
```

#### Check Transaction Status
- **POST** `/api/v1/payment/status`
```json
{
  "id": "transaction_id_to_check"
}
```

#### Webhook Handler
- **POST** `/api/v1/payment/webhook`
Endpoint to receive payment notifications from Santim Pay.

## Payment Methods

Supported payment methods include:
- `telebirr`
- `cbebirr`
- `mpesa`
- `abyssinia`
- `ebirr`
- `amole`
- `santimpay`

## Environment Modes

- **Test Mode**: Uses `https://testnet.santimpay.com/api/v1/gateway`
- **Production Mode**: Uses `https://services.santimpay.com/api/v1/gateway`

## Error Handling

The API returns consistent error responses:
```json
{
  "success": false,
  "message": "Error description",
  "error": "Detailed error message"
}
```

## Security Notes

- Always keep your private key secure
- Never commit `.env` file to version control
- Use HTTPS in production for webhook URLs
- Validate webhook signatures to ensure authenticity

## Dependencies

- `github.com/gin-gonic/gin` - Web framework
- `github.com/golang-jwt/jwt/v5` - JWT token signing
- `github.com/joho/godotenv` - Environment variable management
- `github.com/gin-contrib/cors` - CORS middleware