// Package aggregates holds aggregates that combines many entities into a full object
package customer

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/idylicaro/go-bank/internal/entity"
	"github.com/idylicaro/go-bank/internal/valueobject"
	"github.com/shopspring/decimal"
)

var (
	// ErrInvalidName is returned when the name is not valid in the NewCustome factory
	ErrInvalidName = errors.New("a customer has to have an valid name")
	ErrInvalidCPF  = errors.New("this cpf is invalid or empty")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// a customer have an account
	account *entity.Account

	// a customer can perform many transactions
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string, cpf string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidName
	}

	// TODO: validate CPF
	if cpf == "" {
		return Customer{}, ErrInvalidCPF
	}

	// Create a new account and generate ID
	account := &entity.Account{
		ID:        uuid.New(),
		Name:      name,
		CPF:       cpf,
		Balance:   decimal.NewFromInt(0),
		CreatedAt: time.Now(),
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		account:      account,
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.account.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.account == nil {
		c.account = &entity.Account{}
	}
	c.account.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.account == nil {
		c.account = &entity.Account{}
	}
	c.account.Name = name
}

// GetName return the name of the Customer
func (c *Customer) GetName() string {
	return c.account.Name
}

// SetBalance changes the Balance of the Customer
func (c *Customer) SetBalance(balance decimal.Decimal) {
	if c.account == nil {
		c.account = &entity.Account{}
	}
	c.account.Balance = balance
}

// GetBalance return the Balance of the Customer
func (c *Customer) GetBalance() decimal.Decimal {
	return c.account.Balance
}
