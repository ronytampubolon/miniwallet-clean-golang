package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID         uuid.UUID `json:"id"`
	OwnedBy    uuid.UUID `json:"owned_by"`
	Status     string    `json:"status"`
	EnabledAt  time.Time `json:"enabled_at"`
	DisabledAt time.Time `json:"disabled_at"`
	Balance    float64   `json:"balance"`
}

type DisabledWalletResponse struct {
	ID         uuid.UUID `json:"id"`
	OwnedBy    uuid.UUID `json:"owned_by"`
	Status     string    `json:"status"`
	DisabledAt time.Time `json:"disabled_at"`
	Balance    float64   `json:"balance"`
}

type EnableddWalletResponse struct {
	ID        uuid.UUID `json:"id"`
	OwnedBy   uuid.UUID `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   float64   `json:"balance"`
}

func TransformEnableResponse(account Wallet) EnableddWalletResponse {
	return EnableddWalletResponse{
		ID:        account.ID,
		OwnedBy:   account.OwnedBy,
		Status:    account.Status,
		EnabledAt: account.EnabledAt,
		Balance:   account.Balance,
	}
}
func TransformDisabledResponse(account Wallet) DisabledWalletResponse {
	return DisabledWalletResponse{
		ID:         account.ID,
		OwnedBy:    account.OwnedBy,
		Status:     account.Status,
		DisabledAt: account.EnabledAt,
		Balance:    account.Balance,
	}
}
