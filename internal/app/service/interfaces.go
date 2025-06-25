package service

import (
	"context"
	"echo-software-take-home/internal/app/client/fireblocks"
)

type PgStorage interface {
	SaveWallet(id string, name string, fireblocksVaultId string) error
}

type FireblocksClient interface {
	CreateNewVaultAccount(
		ctx context.Context,
		payload fireblocks.CreateNewVaultAccountRequest,
		idempotencyKey string,
	) (*fireblocks.CreateNewVaultAccountResponse, error)

	GetAssetBalanceForVaultAccountId(
		ctx context.Context,
		vaultAccountId string,
		assetId string,
	) (*fireblocks.GetAssetBalanceForVaultResponse, error)

	GetAssetAddressPaged(
		ctx context.Context,
		vaultAccountId string,
		assetId string,
	) (*fireblocks.GetAssetAddressResponse, error)

	CreateNewTransaction(
		ctx context.Context,
		payload fireblocks.CreateNewTransactionRequest,
	) (*fireblocks.CreateNewTransactionResponse, error)
}
