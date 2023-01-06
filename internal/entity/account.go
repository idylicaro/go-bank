// Package entity entities holds all the entities that are shared across all subdomains
package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrAccountOriginNotFound = errors.New("account origin not found")

	ErrAccountDestinationNotFound = errors.New("account destination not found")
)

type Account struct {
	ID        uuid.UUID
	Name      string
	CPF       string
	Balance   decimal.Decimal
	CreatedAt time.Time
}
