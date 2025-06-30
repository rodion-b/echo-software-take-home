package pg

import (
	"echo-software-take-home/internal/app/config"
	"echo-software-take-home/internal/app/domain"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgStorage struct {
	db *gorm.DB
}

func NewPgStorage(config config.Config) (*PgStorage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
	)

	log.Info().Msg(fmt.Sprintf("Connecting to database with DSN: %s", dsn))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&domain.Wallet{}); err != nil {
		log.Error().Msgf("Failed to migrate database: %v", err)
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &PgStorage{
		db: db,
	}, nil
}

func (s *PgStorage) SaveWallet(wallet *domain.Wallet) error {
	tx := s.db.Create(wallet)
	if tx.Error != nil {
		return fmt.Errorf("failed to save new wallet: %w", tx.Error)
	}

	return nil
}
