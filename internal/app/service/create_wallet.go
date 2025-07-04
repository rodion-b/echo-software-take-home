package service

import (
	"context"
	"echo-software-take-home/internal/app/client/fireblocks"
	"echo-software-take-home/internal/app/domain"

	"github.com/google/uuid"
)

func (s *WalletService) CreateWallet(
	ctx context.Context,
	walletName string,
) (*domain.Wallet, error) {
	// Sending Request to Fireblocks to create a new vault account
	createNewVaultRequest := fireblocks.CreateNewVaultAccountRequest{
		Name: walletName,
	}
	idempotencyKey := uuid.New().String()
	response, err := s.fbclient.CreateNewVaultAccount(ctx, createNewVaultRequest, idempotencyKey)
	if err != nil {
		return nil, err
	}

	// Validate the response
	if response.ID == "" {
		return nil, domain.ErrEmptyVaultAccountID
	}

	if response.Name == "" {
		return nil, domain.ErrEmptyVaultAccountID
	}

	// Can be autogenerated by DB's primary key instead
	localId := uuid.New().String()

	// Create a new Wallet domain object
	wallet := domain.NewWallet(localId, response.ID, walletName)

	// Persist the wallet information in the database
	err = s.pgStorage.SaveWallet(wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
