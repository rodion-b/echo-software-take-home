package domain

type Transaction struct {
	id             string
	vaultAccountId string
	assetId        string
	source         string
	destination    string
	amount         string
	status         string
}

// NewTransaction creates a new Transaction instance
func NewTransaction(id, vaultAccountId, assetId, source, destination, amount, status string) *Transaction {
	return &Transaction{
		id:             id,
		vaultAccountId: vaultAccountId,
		assetId:        assetId,
		source:         source,
		destination:    destination,
		amount:         amount,
		status:         status,
	}
}

// ID returns the transaction ID
func (t *Transaction) ID() string {
	return t.id
}

// VaultAccountID returns the vault account ID
func (t *Transaction) VaultAccountID() string {
	return t.vaultAccountId
}

// AssetID returns the asset ID
func (t *Transaction) AssetID() string {
	return t.assetId
}

// Source returns the source
func (t *Transaction) Source() string {
	return t.source
}

// Destination returns the destination
func (t *Transaction) Destination() string {
	return t.destination
}

// Amount returns the amount
func (t *Transaction) Amount() string {
	return t.amount
}

func (t *Transaction) Status() string {
	return t.status
}
