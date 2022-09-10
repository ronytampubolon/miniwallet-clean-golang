package models

import (
	"time"

	"github.com/google/uuid"
)

type Deposited struct {
	ID          uuid.UUID `json:"id"`
	DepositedBy uuid.UUID `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      float64   `json:"amount"`
	ReferenceID uuid.UUID `json:"reference_id"`
}
