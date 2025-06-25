package domain

type Asset struct {
	id               string
	vaultAccountId   string
	balanceTotal     string
	balanceAvailable string
}

func NewAsset(id, vaultAccountId, balanceTotal, balanceAvailable string) *Asset {
	return &Asset{
		id:               id,
		vaultAccountId:   vaultAccountId,
		balanceTotal:     balanceTotal,
		balanceAvailable: balanceAvailable,
	}
}

func (a Asset) ID() string {
	return a.id
}

func (a Asset) VaultAccountID() string {
	return a.vaultAccountId
}

func (a Asset) BalanceTotal() string {
	return a.balanceTotal
}

func (a Asset) BalanceAvailable() string {
	return a.balanceAvailable
}
