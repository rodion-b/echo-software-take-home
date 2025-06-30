# 🚀 FireGo Wallet Service - How to Run

## Prerequisites

- **Go** 1.20+ (for local development)
- **Docker** & **Docker Compose** (for containerized development)
- **PostgreSQL** (if running locally, not via Docker)
- (Optional) **golangci-lint** for linting

---

## 🐳 Running with Docker Compose (Recommended)

### 1. **Create and configure your .env file in your root use .env.example**
- Use `.env.example` as an example.
  
### 2. **Place your Fireblocks private key:**
- Save your PEM file as `private_key.pem` in the project root.

### 3. **Run the application**
```sh
make run
```

### 4. **Testing using postman**
```
postman_collection_v1.json is provided for testing
```


- The API will be available at [http://localhost:8080](http://localhost:8080)
- PostgreSQL will be available at `localhost:5432` (user: `postgres`, password: `postgres`)

## 🧪 Running Tests

### Run all tests:
```sh
make test
```

### Run specific test:
```sh
go test -v ./internal/app/service -run TestWalletService_CreateWallet
```

### Run tests with coverage:
```sh
make test-coverage
```

---

## 🧹 Linting

```sh
make lint
```

---

## 🔄 Regenerate Mocks

```sh
make regenerate-mocks
```

---

## 📝 API Endpoints

### Create Wallet
```http
POST /wallets
Content-Type: application/json

{
  "name": "My Personal Wallet"
}
```

### Get Wallet Balance
```http
GET /wallets/{walletId}/assets/{assetId}/balance
```

### Get Deposit Address
```http
GET /wallets/{walletId}/assets/{assetId}/address
```

### Initiate Transfer
```http
POST /wallets/{walletId}/transactions
{
  "assetId": "TEST_BTC",
  "amount": "0.001",
  "sourceAddress": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
  "destinationAddress": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
  "note": "Test transfer"
}
```


## 🐛 Work in progress Items

1. **Bug in transaction create needs source and destination as objects**
2. **Add more tests and e2e test** 
  

