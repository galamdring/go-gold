// Code generated by ent, DO NOT EDIT.

package account

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeBudget holds the string denoting the budget edge name in mutations.
	EdgeBudget = "budget"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// BudgetTable is the table that holds the budget relation/edge.
	BudgetTable = "accounts"
	// BudgetInverseTable is the table name for the Budget entity.
	// It exists in this package in order to avoid circular dependency with the "budget" package.
	BudgetInverseTable = "budgets"
	// BudgetColumn is the table column denoting the budget relation/edge.
	BudgetColumn = "budget_accounts"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transactions"
	// TransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionsInverseTable = "transactions"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "account_transactions"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"budget_accounts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeChecking        Type = "checking"
	TypeSavings         Type = "savings"
	TypeCreditCard      Type = "credit_card"
	TypeInstallmentLoan Type = "installment_loan"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeChecking, TypeSavings, TypeCreditCard, TypeInstallmentLoan:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByBudgetField orders the results by budget field.
func ByBudgetField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBudgetStep(), sql.OrderByField(field, opts...))
	}
}

// ByTransactionsCount orders the results by transactions count.
func ByTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionsStep(), opts...)
	}
}

// ByTransactions orders the results by transactions terms.
func ByTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBudgetStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BudgetInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BudgetTable, BudgetColumn),
	)
}
func newTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
	)
}
