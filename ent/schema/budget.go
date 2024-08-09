package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Budget holds the schema definition for the Budget entity.
type Budget struct {
	ent.Schema
}

// Fields of the Budget.
func (Budget) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
		field.Float("amount"),
	}
}

// Edges of the Budget.
func (Budget) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("budgets").Unique(),
		edge.To("accounts", Account.Type),
		edge.To("transactions", Transaction.Type),
	}
}
