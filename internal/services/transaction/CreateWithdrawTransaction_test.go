package transaction

import (
	"testing"

	"github.com/idylicaro/go-bank/internal/domain/customer"
	"github.com/shopspring/decimal"
)

func TestTransaction_NewWithdrawTransactionService(t *testing.T) {
	type testCase struct {
		name            string
		customerName    string
		cpf             string
		initialBalance  decimal.Decimal
		withDrawValue   decimal.Decimal
		expectedBalance decimal.Decimal
		expectedErr     error
	}

	testCases := []testCase{
		{
			name:            "Withdraw with 0 of balance",
			customerName:    "Vick",
			cpf:             "000.000.000-00",
			initialBalance:  decimal.NewFromInt(0),
			withDrawValue:   decimal.NewFromInt(50),
			expectedBalance: decimal.NewFromInt(0),
			expectedErr:     customer.ErrInsufficientBalance,
		},
		{
			name:            "Withdraw with 50 of balance (withdraw 25)",
			customerName:    "Vick",
			cpf:             "000.000.000-00",
			initialBalance:  decimal.NewFromInt(50),
			withDrawValue:   decimal.NewFromInt(25),
			expectedBalance: decimal.NewFromInt(25),
			expectedErr:     nil,
		},
		{
			name:            "Withdraw with 50 of balance (withdraw all)",
			customerName:    "Vick",
			cpf:             "000.000.000-00",
			initialBalance:  decimal.NewFromInt(50),
			withDrawValue:   decimal.NewFromInt(50),
			expectedBalance: decimal.NewFromInt(0),
			expectedErr:     nil,
		},
	}

	ts, err := NewTransactionService(
		WithMemoryCustomerRepository(),
	)
	if err != nil {
		t.Error(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customer, err := customer.NewCustomer(tc.customerName, tc.cpf)
			if err != nil {
				t.Error(err)
			}
			customer.SetBalance(tc.initialBalance)
			err = ts.customers.Add(customer)
			if err != nil {
				t.Error(err)
			}
			err = ts.CreateWithdrawTransaction(customer.GetID(), tc.withDrawValue)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			customer, err = ts.customers.Get(customer.GetID())
			if err != nil {
				t.Error(err)
			}

			if !customer.GetBalance().Equal(tc.expectedBalance) {
				t.Errorf("Expected balance %v, got %v", tc.expectedBalance, customer.GetBalance())
			}
		})
	}
}
