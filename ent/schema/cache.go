package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Cache holds the schema definition for the Cache entity.
type Cache struct {
	ent.Schema
}

// Fields of the Cache.
func (Cache) Fields() []ent.Field {
	return []ent.Field{
		field.Int("walks"),
	}
}
