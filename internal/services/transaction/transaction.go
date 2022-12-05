// Package services holds all the services that connects repositories into a business flow
package services

import (
	"github.com/idylicaro/go-bank/internal/domain/customer"
)

// TransactionConfiguration is an alias for a function that will take in a pointer to an TransactionService and modify it
type TransactionConfiguration func(ts *TransactionService) error

// TransactionService is a implementation of the TransactionService
type TransactionService struct {
	customers customer.CustomerRepository
}

// NewTransactionService takes a variable amount of TransactionConfiguration functions and returns a new TransactionService
// Each TransactionConfiguration will be called in the order they are passed in
func NewTransactionService(cfgs ...TransactionConfiguration) (*TransactionService, error) {
	// Create the transactionService
	ts := &TransactionService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(ts)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}
