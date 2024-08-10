// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/galamdring/go-gold/ent/account"
	"github.com/galamdring/go-gold/ent/budget"
	"github.com/galamdring/go-gold/ent/schema"
	"github.com/galamdring/go-gold/ent/transaction"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (tc *TransactionCreate) SetAmount(s schema.Decimal) *TransactionCreate {
	tc.mutation.SetAmount(s)
	return tc
}

// SetNote sets the "note" field.
func (tc *TransactionCreate) SetNote(s string) *TransactionCreate {
	tc.mutation.SetNote(s)
	return tc
}

// SetCleared sets the "cleared" field.
func (tc *TransactionCreate) SetCleared(b bool) *TransactionCreate {
	tc.mutation.SetCleared(b)
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(i int) *TransactionCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (tc *TransactionCreate) SetAccountID(id int) *TransactionCreate {
	tc.mutation.SetAccountID(id)
	return tc
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableAccountID(id *int) *TransactionCreate {
	if id != nil {
		tc = tc.SetAccountID(*id)
	}
	return tc
}

// SetAccount sets the "account" edge to the Account entity.
func (tc *TransactionCreate) SetAccount(a *Account) *TransactionCreate {
	return tc.SetAccountID(a.ID)
}

// SetBudgetID sets the "budget" edge to the Budget entity by ID.
func (tc *TransactionCreate) SetBudgetID(id int) *TransactionCreate {
	tc.mutation.SetBudgetID(id)
	return tc
}

// SetNillableBudgetID sets the "budget" edge to the Budget entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableBudgetID(id *int) *TransactionCreate {
	if id != nil {
		tc = tc.SetBudgetID(*id)
	}
	return tc
}

// SetBudget sets the "budget" edge to the Budget entity.
func (tc *TransactionCreate) SetBudget(b *Budget) *TransactionCreate {
	return tc.SetBudgetID(b.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Transaction.amount"`)}
	}
	if _, ok := tc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "Transaction.note"`)}
	}
	if _, ok := tc.mutation.Cleared(); !ok {
		return &ValidationError{Name: "cleared", err: errors.New(`ent: missing required field "Transaction.cleared"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(transaction.Table, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeOther, value)
		_node.Amount = value
	}
	if value, ok := tc.mutation.Note(); ok {
		_spec.SetField(transaction.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := tc.mutation.Cleared(); ok {
		_spec.SetField(transaction.FieldCleared, field.TypeBool, value)
		_node.Cleared = value
	}
	if nodes := tc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.AccountTable,
			Columns: []string{transaction.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_transactions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.BudgetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.BudgetTable,
			Columns: []string{transaction.BudgetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(budget.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.budget_transactions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	err      error
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
