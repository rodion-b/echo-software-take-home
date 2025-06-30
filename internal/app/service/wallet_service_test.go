package service

import (
	"context"
	"testing"

	"echo-software-take-home/internal/app/client/fireblocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWalletService_CreateWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock Fireblocks client and PgStorage
	mockFireblocksClient := NewMockFireblocksClient(ctrl)
	mockPgStorage := NewMockPgStorage(ctrl)

	// Create wallet service with mock clients
	walletService := NewWalletService(mockFireblocksClient, mockPgStorage)

	// Test data
	walletName := "Test Wallet"
	expectedVaultID := "vault-123"
	expectedResponse := &fireblocks.CreateNewVaultAccountResponse{
		ID:   expectedVaultID,
		Name: walletName,
	}

	// Set up mock expectations
	mockFireblocksClient.EXPECT().
		CreateNewVaultAccount(
			gomock.Any(), // context
			fireblocks.CreateNewVaultAccountRequest{Name: walletName},
			gomock.Any(), // idempotency key
		).
		Return(expectedResponse, nil).
		Times(1)

	// Set up PgStorage mock expectations
	mockPgStorage.EXPECT().
		SaveWallet(
			gomock.Any(), // wallet object
		).
		Return(nil).
		Times(1)

	// Execute the method
	ctx := context.Background()
	wallet, err := walletService.CreateWallet(ctx, walletName)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, wallet)
	assert.Equal(t, expectedVaultID, wallet.VaultID())
	assert.Equal(t, walletName, wallet.Name())
}

func TestWalletService_GetWalletBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock Fireblocks client and PgStorage
	mockFireblocksClient := NewMockFireblocksClient(ctrl)
	mockPgStorage := NewMockPgStorage(ctrl)

	// Create wallet service with mock clients
	walletService := NewWalletService(mockFireblocksClient, mockPgStorage)

	// Test data
	vaultAccountID := "vault-123"
	assetID := "TEST_BTC"
	expectedResponse := &fireblocks.GetAssetBalanceForVaultResponse{
		ID:        assetID,
		Total:     "1.5",
		Available: "1.0",
		Pending:   "0.5",
	}

	// Set up mock expectations
	mockFireblocksClient.EXPECT().
		GetAssetBalanceForVaultAccountId(
			gomock.Any(), // context
			vaultAccountID,
			assetID,
		).
		Return(expectedResponse, nil).
		Times(1)

	// Execute the method
	ctx := context.Background()
	asset, err := walletService.GetWalletBalance(ctx, vaultAccountID, assetID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.Equal(t, assetID, asset.ID())
	assert.Equal(t, "1.5", asset.BalanceTotal())
	assert.Equal(t, "1.0", asset.BalanceAvailable())
}

func TestWalletService_GetDepositAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock Fireblocks client and PgStorage
	mockFireblocksClient := NewMockFireblocksClient(ctrl)
	mockPgStorage := NewMockPgStorage(ctrl)

	// Create wallet service with mock clients
	walletService := NewWalletService(mockFireblocksClient, mockPgStorage)

	// Test data
	vaultAccountID := "vault-123"
	assetID := "TEST_BTC"
	expectedAddress := "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
	expectedResponse := &fireblocks.GetAssetAddressResponse{
		Addresses: []struct {
			AssetID           string `json:"assetId"`
			Address           string `json:"address"`
			Description       string `json:"description"`
			Tag               string `json:"tag"`
			Type              string `json:"type"`
			CustomerRefID     string `json:"customerRefId"`
			AddressFormat     string `json:"addressFormat"`
			LegacyAddress     string `json:"legacyAddress"`
			EnterpriseAddress string `json:"enterpriseAddress"`
			Bip44AddressIndex int    `json:"bip44AddressIndex"`
			UserDefined       bool   `json:"userDefined"`
		}{
			{
				AssetID: assetID,
				Address: expectedAddress,
			},
		},
	}

	// Set up mock expectations
	mockFireblocksClient.EXPECT().
		GetAssetAddressPaged(
			gomock.Any(), // context
			vaultAccountID,
			assetID,
		).
		Return(expectedResponse, nil).
		Times(1)

	// Execute the method
	ctx := context.Background()
	addresses, err := walletService.GetDepositAddress(ctx, vaultAccountID, assetID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, addresses)
	assert.Len(t, addresses, 1)
	assert.Equal(t, expectedAddress, addresses[0].Address())
}

func TestWalletService_InitiateTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock Fireblocks client and PgStorage
	mockFireblocksClient := NewMockFireblocksClient(ctrl)
	mockPgStorage := NewMockPgStorage(ctrl)

	// Create wallet service with mock clients
	walletService := NewWalletService(mockFireblocksClient, mockPgStorage)

	// Test data
	vaultAccountID := "vault-123"
	assetID := "TEST_BTC"
	amount := "0.001"
	sourceAddress := "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
	destinationAddress := "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
	expectedTxID := "tx-123"
	expectedResponse := &fireblocks.CreateNewTransactionResponse{
		ID:     expectedTxID,
		Status: "SUBMITTED",
	}

	// Set up mock expectations
	mockFireblocksClient.EXPECT().
		CreateNewTransaction(
			gomock.Any(), // context
			fireblocks.CreateNewTransactionRequest{
				AssetId:     assetID,
				Amount:      amount,
				Source:      sourceAddress,
				Destination: destinationAddress,
			},
		).
		Return(expectedResponse, nil).
		Times(1)

	// Execute the method
	ctx := context.Background()
	transaction, err := walletService.InititateTransfer(ctx, vaultAccountID, assetID, amount, sourceAddress, destinationAddress)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, expectedTxID, transaction.ID())
	assert.Equal(t, "SUBMITTED", transaction.Status())
	assert.Equal(t, assetID, transaction.AssetID())
	assert.Equal(t, sourceAddress, transaction.Source())
	assert.Equal(t, destinationAddress, transaction.Destination())
	assert.Equal(t, amount, transaction.Amount())
}
