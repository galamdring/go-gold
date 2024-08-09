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
		field.Other("amount", Decimal{}).SchemaType(map[string]string{
			"postgres": "numeric", // Specify the database type for PostgreSQL
			"sqlite3":   "TEXT",    // Specify the database type for SQLite
			// Add other database types as needed
		}),
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
