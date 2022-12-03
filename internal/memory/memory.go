package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/idylicaro/go-bank/internal/aggregate"
)

// MemoryRepository fulfills the CustomerRepository interface
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(uuid.UUID) (aggregate.Customer, error) {
	return aggregate.Customer{}, nil
}

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(aggregate.Customer) error {
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(aggregate.Customer) error {
	return nil
}
