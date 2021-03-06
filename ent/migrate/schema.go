// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CachesColumns holds the columns for the "caches" table.
	CachesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "walks", Type: field.TypeInt},
	}
	// CachesTable holds the schema information for the "caches" table.
	CachesTable = &schema.Table{
		Name:       "caches",
		Columns:    CachesColumns,
		PrimaryKey: []*schema.Column{CachesColumns[0]},
	}
	// DogsColumns holds the columns for the "dogs" table.
	DogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
	}
	// DogsTable holds the schema information for the "dogs" table.
	DogsTable = &schema.Table{
		Name:       "dogs",
		Columns:    DogsColumns,
		PrimaryKey: []*schema.Column{DogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dogs_users_pets",
				Columns:    []*schema.Column{DogsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeString},
		{Name: "last_digits", Type: field.TypeString},
		{Name: "user_cache", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_caches_cache",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{CachesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CachesTable,
		DogsTable,
		UsersTable,
	}
)

func init() {
	DogsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = CachesTable
}
