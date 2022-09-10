package models

import (
	"time"

	"github.com/google/uuid"
)

type Withdrawal struct {
	ID          uuid.UUID `json:"id"`
	WithdrawnBy uuid.UUID `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      float64   `json:"amount"`
	ReferenceID uuid.UUID `json:"reference_id"`
}
