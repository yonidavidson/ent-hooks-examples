package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	gen "github.com/yonidavidson/ent-side-effect-hooks-example/ent"
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

type syncCtxKey struct{}

// newSyncContext returns a new context marked TX Commit hook enabled.
func newSyncContext(parent context.Context) context.Context {
	return context.WithValue(parent, syncCtxKey{}, "")
}

func cloudSync(ctx context.Context, tx *gen.Tx, cloudID int) {
	tx.OnCommit(func(next gen.Committer) gen.Committer {
		return gen.CommitFunc(func(ctx context.Context, tx *gen.Tx) error {
			fmt.Println("hi")
			// Validate that each TX registers start on the Commit hook only once even if more than one hook is triggered.
			_, ok := ctx.Value(syncCtxKey{}).(string)
			ctx = newSyncContext(ctx)
			if err := next.Commit(ctx, tx); err != nil {
				return err
			}
			if !ok {
				tx.Client().CloudSyncer.Sync(ctx, cloudID)
			}
			return nil
		})
	})
}
