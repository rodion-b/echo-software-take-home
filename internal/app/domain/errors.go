package domain

import "errors"

var (
	ErrEmptyVaultAccountID    = errors.New("empty vault account ID")
	ErrEmptyAssetBalanceTotal = errors.New("empty asset balance total")
	ErrEmptyVaultName         = errors.New("empty vault name")
)
