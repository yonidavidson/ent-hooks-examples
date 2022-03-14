// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/yonidavidson/ent-hooks-examples/ent/cache"
)

// Cache is the model entity for the Cache schema.
type Cache struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Walks holds the value of the "walks" field.
	Walks int `json:"walks,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cache) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cache.FieldID, cache.FieldWalks:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cache", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cache fields.
func (c *Cache) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cache.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cache.FieldWalks:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field walks", values[i])
			} else if value.Valid {
				c.Walks = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Cache.
// Note that you need to call Cache.Unwrap() before calling this method if this Cache
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cache) Update() *CacheUpdateOne {
	return (&CacheClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cache entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cache) Unwrap() *Cache {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cache is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cache) String() string {
	var builder strings.Builder
	builder.WriteString("Cache(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", walks=")
	builder.WriteString(fmt.Sprintf("%v", c.Walks))
	builder.WriteByte(')')
	return builder.String()
}

// Caches is a parsable slice of Cache.
type Caches []*Cache

func (c Caches) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}