package schema

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/yonidavidson/ent-hooks-examples/ent"
	"github.com/yonidavidson/ent-hooks-examples/ent/dog"
	"github.com/yonidavidson/ent-hooks-examples/ent/hook"
	"github.com/yonidavidson/ent-hooks-examples/ent/user"
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
		field.Int("owner_id").
			Optional(),
	}
}

// Edges of the Dog.
func (Dog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Field("owner_id").
			Unique(),
	}
}

// Hooks of the Dog.
func (Dog) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			syncCache,
			hook.HasOp(ent.OpUpdateOne),
		),
		hook.If(
			validateName,
			hook.HasFields(dog.FieldOwnerID),
		),
	}
}

func syncCache(next ent.Mutator) ent.Mutator {
	return hook.DogFunc(func(ctx context.Context, m *gen.DogMutation) (ent.Value, error) {
		cacheID, err := m.Client().Dog.Query().QueryOwner().QueryCache().OnlyID(ctx)
		if err != nil {
			return next.Mutate(ctx, m)
		}
		v, err := next.Mutate(ctx, m)
		if err == nil {
			m.Client().CacheSyncer.Sync(ctx, cacheID)
		}
		return v, err
	})
}

func validateName(next ent.Mutator) ent.Mutator {
	return hook.DogFunc(func(ctx context.Context, m *gen.DogMutation) (ent.Value, error) {
		owID, ok := m.OwnerID()
		if !ok {
			return next.Mutate(ctx, m)
		}
		owner, err := m.Client().User.Query().Where(user.ID(owID)).Only(ctx)
		if err != nil {
			return next.Mutate(ctx, m)
		}
		dn, ok := m.Name()
		if !ok {
			return next.Mutate(ctx, m)
		}
		if owner.Name[0:2] == dn[0:2] {
			return nil, errors.New("invalid dog name")
		}
		return next.Mutate(ctx, m)
	})
}
