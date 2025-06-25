package httpserver

import (
	"context"
	"echo-software-take-home/internal/app/domain"
)

type WalletService interface {
	CreateWallet(
		ctx context.Context,
		walletName string,
	) (*domain.Wallet, error)

	GetWalletBalance(
		ctx context.Context,
		vaultAccountId string,
		assetId string,
	) (*domain.Asset, error)

	GetDepositAddress(
		ctx context.Context,
		vaultAccountId string,
		assetId string,
	) ([]*domain.Address, error)

	InititateTransfer(
		ctx context.Context,
		vaultAccountId string,
		assetId string,
		amount string,
		sourceAddress string,
		destinationAddress string,
	) (*domain.Transaction, error)
}
