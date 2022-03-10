package schema

import "entgo.io/ent"

// DatabaseConfig holds the schema definition for the DatabaseConfig entity.
type DatabaseConfig struct {
	ent.Schema
}

// Fields of the DatabaseConfig.
func (DatabaseConfig) Fields() []ent.Field {
	return nil
}

// Edges of the DatabaseConfig.
func (DatabaseConfig) Edges() []ent.Edge {
	return nil
}
