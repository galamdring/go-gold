// Code generated by ent, DO NOT EDIT.

package account

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/galamdring/go-gold/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldName, v))
}

// CurrentBalance applies equality check predicate on the "current_balance" field. It's identical to CurrentBalanceEQ.
func CurrentBalance(v float64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCurrentBalance, v))
}

// ClearedBalance applies equality check predicate on the "cleared_balance" field. It's identical to ClearedBalanceEQ.
func ClearedBalance(v float64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldClearedBalance, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldType, vs...))
}

// CurrentBalanceEQ applies the EQ predicate on the "current_balance" field.
func CurrentBalanceEQ(v float64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCurrentBalance, v))
}

// CurrentBalanceNEQ applies the NEQ predicate on the "current_balance" field.
func CurrentBalanceNEQ(v float64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCurrentBalance, v))
}

// CurrentBalanceIn applies the In predicate on the "current_balance" field.
func CurrentBalanceIn(vs ...float64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCurrentBalance, vs...))
}

// CurrentBalanceNotIn applies the NotIn predicate on the "current_balance" field.
func CurrentBalanceNotIn(vs ...float64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCurrentBalance, vs...))
}

// CurrentBalanceGT applies the GT predicate on the "current_balance" field.
func CurrentBalanceGT(v float64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCurrentBalance, v))
}

// CurrentBalanceGTE applies the GTE predicate on the "current_balance" field.
func CurrentBalanceGTE(v float64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCurrentBalance, v))
}

// CurrentBalanceLT applies the LT predicate on the "current_balance" field.
func CurrentBalanceLT(v float64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCurrentBalance, v))
}

// CurrentBalanceLTE applies the LTE predicate on the "current_balance" field.
func CurrentBalanceLTE(v float64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCurrentBalance, v))
}

// ClearedBalanceEQ applies the EQ predicate on the "cleared_balance" field.
func ClearedBalanceEQ(v float64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldClearedBalance, v))
}

// ClearedBalanceNEQ applies the NEQ predicate on the "cleared_balance" field.
func ClearedBalanceNEQ(v float64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldClearedBalance, v))
}

// ClearedBalanceIn applies the In predicate on the "cleared_balance" field.
func ClearedBalanceIn(vs ...float64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldClearedBalance, vs...))
}

// ClearedBalanceNotIn applies the NotIn predicate on the "cleared_balance" field.
func ClearedBalanceNotIn(vs ...float64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldClearedBalance, vs...))
}

// ClearedBalanceGT applies the GT predicate on the "cleared_balance" field.
func ClearedBalanceGT(v float64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldClearedBalance, v))
}

// ClearedBalanceGTE applies the GTE predicate on the "cleared_balance" field.
func ClearedBalanceGTE(v float64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldClearedBalance, v))
}

// ClearedBalanceLT applies the LT predicate on the "cleared_balance" field.
func ClearedBalanceLT(v float64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldClearedBalance, v))
}

// ClearedBalanceLTE applies the LTE predicate on the "cleared_balance" field.
func ClearedBalanceLTE(v float64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldClearedBalance, v))
}

// HasBudget applies the HasEdge predicate on the "budget" edge.
func HasBudget() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BudgetTable, BudgetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBudgetWith applies the HasEdge predicate on the "budget" edge with a given conditions (other predicates).
func HasBudgetWith(preds ...predicate.Budget) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newBudgetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransactions applies the HasEdge predicate on the "transactions" edge.
func HasTransactions() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionsWith applies the HasEdge predicate on the "transactions" edge with a given conditions (other predicates).
func HasTransactionsWith(preds ...predicate.Transaction) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newTransactionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(sql.NotPredicates(p))
}
