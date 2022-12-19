package main

import (
	"fmt"

	"github.com/idylicaro/go-bank/internal/domain/customer"
	"github.com/idylicaro/go-bank/internal/services/transaction"
	utils "github.com/idylicaro/go-bank/internal/utils"
	"github.com/shopspring/decimal"
)

func Hello() string {
	return utils.ViperEnvVariable("HELLO_WORLD")
}

func main() {
	ts, err := transaction.NewTransactionService(
		transaction.WithMemoryCustomerRepository(),
	)
	if err != nil {
		panic(err)
	}
	customer, err := customer.NewCustomer("Idyl Icaro", "000.000.000-00")
	if err != nil {
		panic(err)
	}
	err = ts.CreateDepositTransaction(customer.GetID(), decimal.NewFromInt(1))
	if err != nil {
		panic(err)
	}
	fmt.Print(customer.GetBalance())
	fmt.Print("HelloWorld!")
}
