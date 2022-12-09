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
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *entity.Person

	// a customer have an account
	account *entity.Account

	// a customer can perform many transactions
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a new person and generate ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// Create a new account and generate ID
	account := &entity.Account{
		ID:        uuid.New(),
		Balance:   decimal.NewFromInt(0),
		CreatedAt: time.Now(),
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:       person,
		account:      account,
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// GetName return the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}

// GetAccountID return the ID of account
func (c *Customer) GetAccountID() uuid.UUID {
	return c.account.ID
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
