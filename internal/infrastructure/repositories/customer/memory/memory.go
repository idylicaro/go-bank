package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/idylicaro/go-bank/internal/domain/customer"
)

// MemoryCustomerRepository fulfills the CustomerRepository interface
type MemoryCustomerRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (mr *MemoryCustomerRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryCustomerRepository) Update(c customer.Customer) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
