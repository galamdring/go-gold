// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/galamdring/go-gold/ent/account"
	"github.com/galamdring/go-gold/ent/budget"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type account.Type `json:"type,omitempty"`
	// CurrentBalance holds the value of the "current_balance" field.
	CurrentBalance float64 `json:"current_balance,omitempty"`
	// ClearedBalance holds the value of the "cleared_balance" field.
	ClearedBalance float64 `json:"cleared_balance,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges           AccountEdges `json:"edges"`
	budget_accounts *int
	selectValues    sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Budget holds the value of the budget edge.
	Budget *Budget `json:"budget,omitempty"`
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BudgetOrErr returns the Budget value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) BudgetOrErr() (*Budget, error) {
	if e.Budget != nil {
		return e.Budget, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: budget.Label}
	}
	return nil, &NotLoadedError{edge: "budget"}
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[1] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldCurrentBalance, account.FieldClearedBalance:
			values[i] = new(sql.NullFloat64)
		case account.FieldID:
			values[i] = new(sql.NullInt64)
		case account.FieldName, account.FieldType:
			values[i] = new(sql.NullString)
		case account.ForeignKeys[0]: // budget_accounts
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = account.Type(value.String)
			}
		case account.FieldCurrentBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field current_balance", values[i])
			} else if value.Valid {
				a.CurrentBalance = value.Float64
			}
		case account.FieldClearedBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cleared_balance", values[i])
			} else if value.Valid {
				a.ClearedBalance = value.Float64
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field budget_accounts", value)
			} else if value.Valid {
				a.budget_accounts = new(int)
				*a.budget_accounts = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryBudget queries the "budget" edge of the Account entity.
func (a *Account) QueryBudget() *BudgetQuery {
	return NewAccountClient(a.config).QueryBudget(a)
}

// QueryTransactions queries the "transactions" edge of the Account entity.
func (a *Account) QueryTransactions() *TransactionQuery {
	return NewAccountClient(a.config).QueryTransactions(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("current_balance=")
	builder.WriteString(fmt.Sprintf("%v", a.CurrentBalance))
	builder.WriteString(", ")
	builder.WriteString("cleared_balance=")
	builder.WriteString(fmt.Sprintf("%v", a.ClearedBalance))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
