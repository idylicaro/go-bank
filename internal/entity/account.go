// Package entities holds all the entities that are shared across all subdomains
package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID        uuid.UUID
	Balance   decimal.Decimal
	CreatedAt time.Time
}
