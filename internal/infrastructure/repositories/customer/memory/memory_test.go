package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/idylicaro/go-bank/internal/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		id          uuid.UUID
		name        string
		cpf         string
		expectedErr error
	}

	// Create a fake customer to add to repository
	cust, err := customer.NewCustomer("Percy", "000.000.000-00")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "No Customer By	ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		custName    string
		custCpf     string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			custName:    "Percy",
			custCpf:     "000.000.000-00",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryCustomerRepository{
				customers: map[uuid.UUID]customer.Customer{},
			}

			cust, err := customer.NewCustomer(tc.custName, tc.custCpf)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}
