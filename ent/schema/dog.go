package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/yonidavidson/ent-side-effect-hooks-example/ent"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent/hook"
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

// Hooks of the Dog.
func (Dog) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(syncCache,
			hook.HasOp(ent.OpUpdateOne),
		),
	}
}

func syncCache(next ent.Mutator) ent.Mutator {
	return hook.DogFunc(func(ctx context.Context, m *gen.DogMutation) (ent.Value, error) {
		cloudID, err := m.Client().Dog.Query().QueryOwner().QueryCache().OnlyID(ctx)
		if err != nil {
			return next.Mutate(ctx, m)
		}
		v, err := next.Mutate(ctx, m)
		if err == nil {
			m.Client().CacheSyncer.Sync(ctx, cloudID)
		}
		return v, err
	})
}
