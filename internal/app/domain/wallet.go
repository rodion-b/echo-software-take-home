package domain

import "time"

type Wallet struct {
	id        string    `gorm:"primaryKey"`
	name      string    `gorm:"type:varchar(255);not null"`
	vaultId   string    `gorm:"type:varchar(255);not null"`
	createdAt time.Time `gorm:"autoCreateTime"`
	updatedAt time.Time `gorm:"autoUpdateTime"`
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

func (w *Wallet) CreatedAt() time.Time {
	return w.createdAt
}

func (w *Wallet) UpdatedAt() time.Time {
	return w.updatedAt
}