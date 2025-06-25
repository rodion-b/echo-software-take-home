package service

import (
	"context"
	"echo-software-take-home/internal/app/domain"
)

func (s *WalletService) GetWalletBalance(
	ctx context.Context,
	vaultAccountId string,
	assetId string,
) (*domain.Asset, error) {
	resp, err := s.fbclient.GetAssetBalanceForVaultAccountId(ctx, vaultAccountId, assetId)
	if err != nil {
		return nil, err
	}
	// Validate the response
	if resp.ID == "" {
		return nil, ErrEmptyVaultAccountID
	}
	if resp.Total == "" {
		return nil, ErrEmptyAssetBalanceTotal
	}

	// Create a new Asset domain object
	asset := domain.NewAsset(
		resp.ID,
		vaultAccountId,
		resp.Total,
		resp.Available,
	)

	return asset, nil
}
