package domain

type Wallet struct {
	id      string
	vaultId string
	name    string
}

func NewWallet(id, vaultId, name string) *Wallet {
	return &Wallet{
		id:      id,
		vaultId: vaultId,
		name:    name,
	}
}

func (w *Wallet) ID() string {
	return w.id
}

func (w *Wallet) VaultID() string {
	return w.vaultId
}

func (w *Wallet) Name() string {
	return w.name
}
