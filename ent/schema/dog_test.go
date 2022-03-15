package schema_test

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"github.com/yonidavidson/ent-hooks-examples/cache"
	"github.com/yonidavidson/ent-hooks-examples/ent"
	"github.com/yonidavidson/ent-hooks-examples/ent/enttest"
)

func TestCacheHook(t *testing.T) {
	ctx := context.Background()
	cs := cache.NewSyncer()
	c := enttest.Open(t, dialect.SQLite,
		"file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1",
		enttest.WithOptions(ent.CacheSyncer(cs)),
	)
	cs.Start(ctx, c)
	cl := c.Cache.Create().SetWalks(-1).SaveX(ctx)
	d := c.Dog.Create().SetName("Karashindo").SaveX(ctx)
	u := c.User.Create().SetName("Yoni").
		SetCache(cl).
		AddPets(d).
		SetPhoneNumber("315-077-2231").
		SaveX(ctx)
	c.Dog.UpdateOne(d).SetName("Fortuna").ExecX(ctx)
	cs.Close()
	cl = u.QueryCache().OnlyX(ctx)
	require.True(t, cl.Walks > 0)
}

func TestDogNameValidationHook(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite,
		"file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1",
	)
	u := c.User.Create().SetName("Yoni").
		SetPhoneNumber("315-077-2231").
		SaveX(ctx)
	_, err := c.Dog.Create().SetName("Yolo").SetOwner(u).Save(ctx)
	require.Error(t, err)
}
