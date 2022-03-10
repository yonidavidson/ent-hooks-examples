package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Cloud holds the schema definition for the Cloud entity.
type Cloud struct {
	ent.Schema
}

// Fields of the Cloud.
func (Cloud) Fields() []ent.Field {
	return []ent.Field{
		field.Int("walks"),
	}
}
