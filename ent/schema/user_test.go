package schema_test

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"github.com/yonidavidson/ent-hooks-examples/ent/enttest"
)

func TestUserPhoneNumberHook(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite, "file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1")
	u := c.User.Create().SetName("Yoni").SetPhoneNumber("315-194-6020").SaveX(ctx)
	require.Equal(t, "315-194-****", u.PhoneNumber)
	require.Equal(t, "6020", u.LastDigits)
	require.Equal(t, "315-194-6020", u.FullPhoneNumber())
}
