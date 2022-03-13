package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/yonidavidson/ent-side-effect-hooks-example/ent"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent/hook"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Dog.Type),
		edge.To("cloud", Cloud.Type).Unique(),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(userCloudReSync,
			hook.HasOp(ent.OpUpdateOne),
		),
	}
}

func userCloudReSync(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		fmt.Println("start hook from user")
		cloudID, err := m.Client().User.Query().QueryCloud().OnlyID(ctx)
		if err != nil {
			return next.Mutate(ctx, m)
		}
		v, err := next.Mutate(ctx, m)
		if err == nil {
			m.Client().CloudSyncer.Sync(ctx, cloudID)
		}
		return v, err
	})
}
