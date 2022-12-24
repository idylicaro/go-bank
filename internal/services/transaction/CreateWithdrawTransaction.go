package transaction

import (
	"github.com/google/uuid"
	"github.com/idylicaro/go-bank/internal/domain/customer"
	"github.com/shopspring/decimal"
)

// CreateOrder will chaintogether all repositories to create a order for a customer
func (t *TransactionService) CreateWithdrawTransaction(customerID uuid.UUID, withdrawValue decimal.Decimal) error {
	// Get the customer
	c, err := t.customers.Get(customerID)
	if err != nil {
		return err
	}
	// LessThan is '<'
	if c.GetBalance().LessThan(withdrawValue) {
		return customer.ErrInsufficientBalance
	}
	c.SetBalance(c.GetBalance().Sub(withdrawValue)) // balance -= withdrawValue

	err = t.customers.Update((c))
	if err != nil {
		return err
	}
	return nil
}
