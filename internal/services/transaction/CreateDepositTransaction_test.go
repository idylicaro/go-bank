package services

import (
	"testing"

	"github.com/idylicaro/go-bank/internal/aggregate"
	"github.com/shopspring/decimal"
)

func TestTransaction_NewDepositTransactionService(t *testing.T) {
	ts, err := NewTransactionService(
		WithMemoryCustomerRepository(),
	)
	if err != nil {
		t.Error(err)
	}

	// Add customer
	customer, err := aggregate.NewCustomer("Vick")
	if err != nil {
		t.Error(err)
	}

	err = ts.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	err = ts.CreateDepositTransaction(customer.GetID(), decimal.NewFromInt(50))
	if err != nil {
		t.Error(err)
	}

	customer, err = ts.customers.Get(customer.GetID())
	if err != nil {
		t.Error(err)
	}
	if !customer.GetBalance().Equal(decimal.NewFromInt(50)) {
		t.Errorf("Expected balance %v, got %v", decimal.NewFromInt(50), customer.GetBalance())
	}
}
