package schema_test

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"github.com/yonidavidson/ent-hooks-examples/ent/enttest"
)

func TestUserConnectionStringHook(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite,
		"file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1",
	)
	u := c.User.Create().SetName("Yoni").SetConnectionString("mysql://root:pass@localhost:3306)").SaveX(ctx)
	require.Equal(t, "mysql://root:****@localhost:3306)", u.ConnectionString)
	require.Equal(t, "pass", u.Password)
	require.Equal(t, "mysql://root:pass@localhost:3306)", u.FullConnectionString())
}
