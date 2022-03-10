// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/yonidavidson/entsideeffecthooksexample/ent/remote"
)

// Remote is the model entity for the Remote schema.
type Remote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Walks holds the value of the "walks" field.
	Walks int `json:"walks,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Remote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case remote.FieldID, remote.FieldWalks:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Remote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Remote fields.
func (r *Remote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case remote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case remote.FieldWalks:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field walks", values[i])
			} else if value.Valid {
				r.Walks = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Remote.
// Note that you need to call Remote.Unwrap() before calling this method if this Remote
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Remote) Update() *RemoteUpdateOne {
	return (&RemoteClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Remote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Remote) Unwrap() *Remote {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Remote is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Remote) String() string {
	var builder strings.Builder
	builder.WriteString("Remote(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", walks=")
	builder.WriteString(fmt.Sprintf("%v", r.Walks))
	builder.WriteByte(')')
	return builder.String()
}

// Remotes is a parsable slice of Remote.
type Remotes []*Remote

func (r Remotes) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
