/*
Package models provides data structures and types for managing financial accounts, budgets, transactions, and categories in a budgeting API.

This package includes:
  - Account types and their properties
  - Budget management
  - Transaction handling
  - Category definitions

*/
package models

import "github.com/shopspring/decimal"

// AccountType is an enum for the type of Account. Unknown, Checking, Savings, CreditCard, Loan
type AccountType int

const (
	// Unknown is the default account type, and is likely a sign that its not an actual account
	Unknown AccountType = iota
	// Checking represents a checking account
	Checking
	// Savings represents a savings account
	Savings
	// CreditCard represents a credit card account
	CreditCard
	// Loan represents a loan account
	Loan
)

// Account is a representation of a single account
type Account struct {
	// ID is the internal identifier
	ID       int     `json:"id"`
	// BudgetId is the ID if the budget that this belongs to
	BudgetId int     `json:"budget_id"`
	// Name is the display name of the account
	Name     string  `json:"name"`
	// Type is the AccountType of the account
	Type     AccountType  `json:"type"`
	// CurrentBalance is the current balance of the account with all transactions processed
	CurrentBalance decimal.Decimal  `json:"current_balance"`
	// ClearedBalance is the balance of the account based on all cleared transactions
	ClearedBalance decimal.Decimal  `json:"cleared_balance"`
}