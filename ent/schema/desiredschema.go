package schema

import "entgo.io/ent"

// DesiredSchema holds the schema definition for the DesiredSchema entity.
type DesiredSchema struct {
	ent.Schema
}

// Fields of the DesiredSchema.
func (DesiredSchema) Fields() []ent.Field {
	return nil
}

// Edges of the DesiredSchema.
func (DesiredSchema) Edges() []ent.Edge {
	return nil
}
