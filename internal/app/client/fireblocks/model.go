package fireblocks

type GetAccountsResponse struct {
	Accounts []Accounts `json:"accounts"`
}

type Assets struct {
	ID           string `json:"id"`
	Total        string `json:"total"`
	Balance      string `json:"balance"`
	LockedAmount string `json:"lockedAmount"`
	Available    string `json:"available"`
	Pending      string `json:"pending"`
	Frozen       string `json:"frozen"`
	Staked       string `json:"staked"`
	BlockHeight  string `json:"blockHeight"`
}

type Accounts struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	HiddenOnUI    bool     `json:"hiddenOnUI"`
	AutoFuel      bool     `json:"autoFuel"`
	CustomerRefID string   `json:"customerRefId,omitempty"`
	Assets        []Assets `json:"assets"`
}

type CreateNewVaultAccountRequest struct {
	Name string `json:"name"`
}

type CreateNewVaultAccountResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	HiddenOnUI    bool   `json:"hiddenOnUI"`
	CustomerRefID string `json:"customerRefId"`
	AutoFuel      bool   `json:"autoFuel"`
}

type GetAssetBalanceForVaultResponse struct {
	ID           string `json:"id"`
	Total        string `json:"total"`
	Available    string `json:"available"`
	Pending      string `json:"pending"`
	Frozen       string `json:"frozen"`
	LockedAmount string `json:"lockedAmount"`
	BlockHeight  string `json:"blockHeight"`
	BlockHash    string `json:"blockHash"`
}

type GetAssetAddressResponse struct {
	Addresses []struct {
		AssetID           string `json:"assetId"`
		Address           string `json:"address"`
		Description       string `json:"description"`
		Tag               string `json:"tag"`
		Type              string `json:"type"`
		CustomerRefID     string `json:"customerRefId"`
		AddressFormat     string `json:"addressFormat"`
		LegacyAddress     string `json:"legacyAddress"`
		EnterpriseAddress string `json:"enterpriseAddress"`
		Bip44AddressIndex int    `json:"bip44AddressIndex"`
		UserDefined       bool   `json:"userDefined"`
	} `json:"addresses"`
	Paging struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"paging"`
}

type TransferRequest struct {
	Operation       string      `json:"operation"`
	Note            string      `json:"note"`
	ExternalTxID    string      `json:"externalTxId"`
	AssetID         string      `json:"assetId"`
	Source          WalletInfo  `json:"source"`
	Destination     WalletInfo  `json:"destination"`
	Destinations    []Transfer  `json:"destinations"`
	Amount          string      `json:"amount"`
	TreatAsGross    bool        `json:"treatAsGrossAmount"`
	ForceSweep      bool        `json:"forceSweep"`
	FeeLevel        string      `json:"feeLevel"`
	PriorityFee     string      `json:"priorityFee"`
	FailOnLowFee    bool        `json:"failOnLowFee"`
	MaxFee          string      `json:"maxFee"`
	GasLimit        string      `json:"gasLimit"`
	ReplaceTxByHash string      `json:"replaceTxByHash"`
	ExtraParameters ExtraParams `json:"extraParameters"`
	CustomerRefID   string      `json:"customerRefId"`
	UseGasless      bool        `json:"useGasless"`
}

type WalletInfo struct {
	Type           string         `json:"type"`
	SubType        string         `json:"subType"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	WalletID       string         `json:"walletId"`
	OneTimeAddress *AddressDetail `json:"oneTimeAddress,omitempty"`
}

type AddressDetail struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

type Transfer struct {
	Amount      string     `json:"amount"`
	Destination WalletInfo `json:"destination"`
}

type ExtraParams struct {
	NodeControls     map[string]interface{} `json:"nodeControls"`
	RawMessageData   map[string]interface{} `json:"rawMessageData"`
	ContractCallData string                 `json:"contractCallData"`
	ProgramCallData  string                 `json:"programCallData"`
	InputsSelection  map[string]interface{} `json:"inputsSelection"`
}


type CreateNewTransactionRequest struct {
	AssetId     string `json:"assetId"`
	Amount      string `json:"amount"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type CreateNewTransactionResponse struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
}