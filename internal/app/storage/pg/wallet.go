package pg

import (
	"time"
)

type Wallet struct {
	ID                string    `gorm:"primaryKey"`
	Name              string    `gorm:"type:varchar(255);not null"`
	FireblocksVaultID string    `gorm:"type:varchar(255);not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
