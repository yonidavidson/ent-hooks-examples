package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Dog holds the schema definition for the Dog entity.
type Dog struct {
	ent.Schema
}

// Fields of the Dog.
func (Dog) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the Dog.
func (Dog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
