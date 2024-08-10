/*
Package models provides data structures and types for managing
financial accounts, budgets, transactions, and categories in a budgeting API.

This package includes:
  - Account types and their properties
  - Budget management
  - Transaction handling
  - Category definitions
*/
package models

import (
	"github.com/galamdring/go-gold/ent"
	"github.com/galamdring/go-gold/ent/account"
	"github.com/shopspring/decimal"
)

const (
	InvalidAccountIdMessage = "Invalid account id format"
)

// AccountRestModel is a api facing representation of a single account.
type AccountRestModel struct {
	// ID is the internal identifier
	ID int `json:"id"`
	// Budget is the budget that this belongs to
	Budget *BudgetRestModel `json:"budget,omitempty"`
	// Name is the display name of the account
	Name string `json:"name"`
	// Type is the AccountType of the account
	Type account.Type `json:"type"`
	// CurrentBalance is the current balance of the account with all transactions processed
	CurrentBalance *decimal.Decimal `json:"currentBalance,omitempty"`
	// ClearedBalance is the balance of the account based on all cleared transactions
	ClearedBalance *decimal.Decimal `json:"clearedBalance,omitempty"`
}

func BuildAccountRestModel(acc *ent.Account) *AccountRestModel {
	var bud *BudgetRestModel
	if acc.Edges.Budget != nil {
		bud = BuildBudgetRestModel(acc.Edges.Budget)
	}

	return &AccountRestModel{
		ID: acc.ID,
		Budget: bud,
		Name: acc.Name,
		Type: acc.Type,
		CurrentBalance: GetCurrentBalance(acc.Edges.Transactions),
		ClearedBalance: GetClearedBalance(acc.Edges.Transactions),
	}
}

func GetCurrentBalance(trans []*ent.Transaction) *decimal.Decimal {
	if len(trans) == 0 {
		return nil
	}

	balance := decimal.NewFromInt(0)
	for _, tran := range trans {
		balance.Add(decimal.Decimal(tran.Amount))
	}

	return &balance
}

func GetClearedBalance(trans []*ent.Transaction) *decimal.Decimal {
	if len(trans) == 0 {
		return nil
	}

	balance := decimal.NewFromInt(0)
	for _, tran := range trans {
		if tran.Cleared {
			balance.Add(decimal.Decimal(tran.Amount))
		}
	}

	return &balance
}

func BuildAccountRestModels(accs []*ent.Account) (result []*AccountRestModel) {
	for _, acc := range accs {
		result = append(result, BuildAccountRestModel(acc))
	}

	return result
}

// AccountModel handles account related operations
type AccountModel struct {
	BaseModel
	client *ent.Client
}

// NewAccountModel creates an instance of AccountModel with the provided client
func NewAccountModel(client *ent.Client) *AccountModel {
	return &AccountModel{client: client}
}
