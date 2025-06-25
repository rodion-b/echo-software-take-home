package service

import (
	"context"
	"echo-software-take-home/internal/app/client/fireblocks"
	"echo-software-take-home/internal/app/domain"
)

func (s *WalletService) InititateTransfer(
	ctx context.Context,
	vaultAccountId string,
	assetId string,
	amount string,
	sourceAddress string,
	destinationAddress string,
) (*domain.Transaction, error) {
	// Creating the paaload for the new transaction request
	createNewTransactionRequest := fireblocks.CreateNewTransactionRequest{
		AssetId:     assetId,
		Amount:      amount,
		Source:      sourceAddress,
		Destination: destinationAddress,
	}

	// Sending the request to create a new transaction
	resp, err := s.fbclient.CreateNewTransaction(ctx, createNewTransactionRequest)
	if err != nil {
		return nil, err
	}

	// Converting to domain.Transaction
	tx := domain.NewTransaction(
		resp.ID,
		vaultAccountId,
		assetId,
		sourceAddress,
		destinationAddress,
		amount,
		resp.Status,
	)

	return tx, nil
}
