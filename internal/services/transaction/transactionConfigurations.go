package transaction

import (
	"github.com/idylicaro/go-bank/internal/domain/customer"
	customerRepo "github.com/idylicaro/go-bank/internal/infrastructure/repositories/customer/memory"
)

// WithCustomerRepository applies a given customer repository to the TransactionService
func WithCustomerRepository(cr customer.CustomerRepository) TransactionConfiguration {
	// return a function that matches the TransactionConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(ts *TransactionService) error {
		ts.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the TransactionService
func WithMemoryCustomerRepository() TransactionConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := customerRepo.New()
	return WithCustomerRepository(cr)
}
