package service

type WalletService struct {
	fbclient  FireblocksClient
	pgStorage PgStorage
}

func NewWalletService(fbclient FireblocksClient, pgStorage PgStorage) *WalletService {
	return &WalletService{
		fbclient:  fbclient,
		pgStorage: pgStorage,
	}
}
