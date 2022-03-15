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
		field.String("phone_number").
			NotEmpty(),
		field.String("last_digits").
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
		hook.If(maskPhoneNumber,
			hook.And(hook.Or(hook.HasOp(ent.OpUpdateOne), hook.HasOp(ent.OpCreate)), hook.HasFields(user.FieldPhoneNumber)),
		),
	}
}

func maskPhoneNumber(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		cs, ok := m.PhoneNumber()
		if !ok {
			return next.Mutate(ctx, m)
		}
		sp := strings.Split(cs, "-")
		if len(sp) != 3 {
			return next.Mutate(ctx, m)
		}
		m.SetLastDigits(sp[2])
		sp[2] = "****"
		m.SetPhoneNumber(strings.Join(sp[0:3], "-"))
		return next.Mutate(ctx, m)
	})
}
