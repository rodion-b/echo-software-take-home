package service

import (
	"context"
	"echo-software-take-home/internal/app/domain"
)

func (s *WalletService) GetDepositAddress(
	ctx context.Context,
	vaultAccountId string,
	assetId string,
) ([]*domain.Address, error) {
	// Get the deposit addressed for the specified vault account and asset ID
	resp, err := s.fbclient.GetAssetAddressPaged(ctx, vaultAccountId, assetId)
	if err != nil {
		return nil, err
	}

	assetAddresses := make([]*domain.Address, len(resp.Addresses))

	for i, a := range resp.Addresses {
		// Create a new Address domain object for each address
		assetAddresses[i] = domain.NewAddress(
			a.Address,
			assetId,
			vaultAccountId,
		)
	}

	return assetAddresses, nil
}
