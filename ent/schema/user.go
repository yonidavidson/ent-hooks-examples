package schema

import (
	"context"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/yonidavidson/ent-hooks-examples/ent"
	"github.com/yonidavidson/ent-hooks-examples/ent/hook"
	"github.com/yonidavidson/ent-hooks-examples/ent/user"
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
		field.String("connection_string").
			NotEmpty(),
		field.String("password").
			Sensitive().
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Dog.Type),
		edge.To("cache", Cache.Type).Unique(),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(clearConnectionString,
			hook.And(hook.Or(hook.HasOp(ent.OpUpdateOne), hook.HasOp(ent.OpCreate)), hook.HasFields(user.FieldConnectionString)),
		),
	}
}

func clearConnectionString(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		cs, ok := m.ConnectionString()
		if !ok {
			return next.Mutate(ctx, m)
		}
		sp := strings.Split(cs, "@")
		if len(sp) != 2 {
			return next.Mutate(ctx, m)
		}
		sp = strings.Split(sp[0], ":")
		if len(sp) != 3 {
			return next.Mutate(ctx, m)
		}
		pass := sp[2]
		m.SetPassword(pass)
		m.SetConnectionString(strings.ReplaceAll(cs, pass, "****"))
		return next.Mutate(ctx, m)
	})
}
