package domain

type Address struct {
	address        string
	assetId        string
	vaultAccountId string
}

// NewAddress creates a new Address instance
func NewAddress(address, assetId, vaultAccountId string) *Address {
	return &Address{
		address:        address,
		assetId:        assetId,
		vaultAccountId: vaultAccountId,
	}
}

// ID returns the ID of the address
func (a *Address) Address() string {
	return a.address
}

// AssetID returns the asset ID associated with the address
func (a *Address) AssetID() string {
	return a.assetId
}

// VaultAccountID returns the vault account ID associated with the address
func (a *Address) VaultAccountID() string {
	return a.vaultAccountId
}
