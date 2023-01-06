package customer_test

import (
	"testing"

	"github.com/idylicaro/go-bank/internal/domain/customer"
)

func TestCustomer_NewCustomer(t *testing.T) {
	// testcase data struct
	type testCase struct {
		test        string
		name        string
		cpf         string
		expectedErr error
	}

	// Create new test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			cpf:         "000.000.000-00",
			expectedErr: customer.ErrInvalidName,
		},
		{
			test:        "Empty CPF validation",
			name:        "Igor Iure",
			cpf:         "",
			expectedErr: customer.ErrInvalidCPF,
		},
		{
			test:        "Valid Customer",
			name:        "Percy Bolmer",
			cpf:         "000.000.000-00",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := customer.NewCustomer(tc.name, tc.cpf)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
