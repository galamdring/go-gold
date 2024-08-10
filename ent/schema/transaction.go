package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Other("amount", Decimal{}).SchemaType(map[string]string{
			"postgres": "numeric", // Specify the database type for PostgreSQL
			"sqlite3":   "TEXT",    // Specify the database type for SQLite
			// Add other database types as needed
		}),
		field.String("note"),
		field.Bool("cleared"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("transactions").Unique(),
		edge.From("budget", Budget.Type).Ref("transactions").Unique(),
	}
}
