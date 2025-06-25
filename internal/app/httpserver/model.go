package httpserver

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"` // optional
	Message   string `json:"message"`
}

type GetWalletBalanceResponse struct {
	WalletID         string `json:"walletId"`
	AssetID          string `json:"assetId"`
	BalanceTotal     string `json:"balanceTotal"`
	BalanceAvailable string `json:"balanceAvailable"`
}

type GetDepositAddressResponse struct {
	WalletID  string   `json:"walletId"`
	AssetID   string   `json:"assetId"`
	Addresses []string `json:"addresses"`
}

type CreateWalletRequest struct {
	Name string `json:"name"`
}

type CreateWalletResponse struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	VaultID string `json:"vault_id"`
}

type InitiateTransferRequest struct {
	AssetId            string `json:"assetId"`
	Amount             string `json:"amount"`
	SourceAddress      string `json:"sourceAddress"` // Missing from the task
	DestinationAddress string `json:"destinationAddress"`
	Note               string `json:"note"`
}
// Bug here Source: should be object, destination: should be object
type InitiateTransferResponse struct {
	TransactionID      string `json:"transactionId"`
	Status             string `json:"status"`
	AssetId            string `json:"assetId"`
	Amount             string `json:"amount"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	Note               string `json:"note"`
}
