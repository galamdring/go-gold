// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/galamdring/go-gold/ent/account"
	"github.com/galamdring/go-gold/ent/budget"
	"github.com/galamdring/go-gold/ent/schema"
	"github.com/galamdring/go-gold/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount schema.Decimal `json:"amount,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Cleared holds the value of the "cleared" field.
	Cleared bool `json:"cleared,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                TransactionEdges `json:"edges"`
	account_transactions *int
	budget_transactions  *int
	selectValues         sql.SelectValues
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// Budget holds the value of the budget edge.
	Budget *Budget `json:"budget,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) AccountOrErr() (*Account, error) {
	if e.Account != nil {
		return e.Account, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "account"}
}

// BudgetOrErr returns the Budget value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) BudgetOrErr() (*Budget, error) {
	if e.Budget != nil {
		return e.Budget, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: budget.Label}
	}
	return nil, &NotLoadedError{edge: "budget"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldAmount:
			values[i] = new(schema.Decimal)
		case transaction.FieldCleared:
			values[i] = new(sql.NullBool)
		case transaction.FieldID:
			values[i] = new(sql.NullInt64)
		case transaction.FieldNote:
			values[i] = new(sql.NullString)
		case transaction.ForeignKeys[0]: // account_transactions
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[1]: // budget_transactions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transaction.FieldAmount:
			if value, ok := values[i].(*schema.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				t.Amount = *value
			}
		case transaction.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				t.Note = value.String
			}
		case transaction.FieldCleared:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cleared", values[i])
			} else if value.Valid {
				t.Cleared = value.Bool
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_transactions", value)
			} else if value.Valid {
				t.account_transactions = new(int)
				*t.account_transactions = int(value.Int64)
			}
		case transaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field budget_transactions", value)
			} else if value.Valid {
				t.budget_transactions = new(int)
				*t.budget_transactions = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transaction.
// This includes values selected through modifiers, order, etc.
func (t *Transaction) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the Transaction entity.
func (t *Transaction) QueryAccount() *AccountQuery {
	return NewTransactionClient(t.config).QueryAccount(t)
}

// QueryBudget queries the "budget" edge of the Transaction entity.
func (t *Transaction) QueryBudget() *BudgetQuery {
	return NewTransactionClient(t.config).QueryBudget(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return NewTransactionClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(t.Note)
	builder.WriteString(", ")
	builder.WriteString("cleared=")
	builder.WriteString(fmt.Sprintf("%v", t.Cleared))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction
