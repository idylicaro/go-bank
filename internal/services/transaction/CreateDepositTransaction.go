package transaction

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CreateDepositTransaction will create deposit transaction to increase the balance
func (t *TransactionService) CreateDepositTransaction(customerID uuid.UUID, depositValue decimal.Decimal) error {
	// Get the customer
	c, err := t.customers.Get(customerID)
	if err != nil {
		return err
	}

	c.SetBalance(c.GetBalance().Add(depositValue)) // balance += depositValue

	err = t.customers.Update((c))
	if err != nil {
		return err
	}
	return nil
}
