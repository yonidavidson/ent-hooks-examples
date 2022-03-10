// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yonidavidson/entsideeffecthooksexample/ent/databaseconfig"
	"github.com/yonidavidson/entsideeffecthooksexample/ent/predicate"
)

// DatabaseConfigUpdate is the builder for updating DatabaseConfig entities.
type DatabaseConfigUpdate struct {
	config
	hooks    []Hook
	mutation *DatabaseConfigMutation
}

// Where appends a list predicates to the DatabaseConfigUpdate builder.
func (dcu *DatabaseConfigUpdate) Where(ps ...predicate.DatabaseConfig) *DatabaseConfigUpdate {
	dcu.mutation.Where(ps...)
	return dcu
}

// SetURL sets the "URL" field.
func (dcu *DatabaseConfigUpdate) SetURL(s string) *DatabaseConfigUpdate {
	dcu.mutation.SetURL(s)
	return dcu
}

// Mutation returns the DatabaseConfigMutation object of the builder.
func (dcu *DatabaseConfigUpdate) Mutation() *DatabaseConfigMutation {
	return dcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dcu *DatabaseConfigUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dcu.hooks) == 0 {
		affected, err = dcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatabaseConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dcu.mutation = mutation
			affected, err = dcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dcu.hooks) - 1; i >= 0; i-- {
			if dcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dcu *DatabaseConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := dcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dcu *DatabaseConfigUpdate) Exec(ctx context.Context) error {
	_, err := dcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcu *DatabaseConfigUpdate) ExecX(ctx context.Context) {
	if err := dcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dcu *DatabaseConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   databaseconfig.Table,
			Columns: databaseconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: databaseconfig.FieldID,
			},
		},
	}
	if ps := dcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: databaseconfig.FieldURL,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{databaseconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DatabaseConfigUpdateOne is the builder for updating a single DatabaseConfig entity.
type DatabaseConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DatabaseConfigMutation
}

// SetURL sets the "URL" field.
func (dcuo *DatabaseConfigUpdateOne) SetURL(s string) *DatabaseConfigUpdateOne {
	dcuo.mutation.SetURL(s)
	return dcuo
}

// Mutation returns the DatabaseConfigMutation object of the builder.
func (dcuo *DatabaseConfigUpdateOne) Mutation() *DatabaseConfigMutation {
	return dcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dcuo *DatabaseConfigUpdateOne) Select(field string, fields ...string) *DatabaseConfigUpdateOne {
	dcuo.fields = append([]string{field}, fields...)
	return dcuo
}

// Save executes the query and returns the updated DatabaseConfig entity.
func (dcuo *DatabaseConfigUpdateOne) Save(ctx context.Context) (*DatabaseConfig, error) {
	var (
		err  error
		node *DatabaseConfig
	)
	if len(dcuo.hooks) == 0 {
		node, err = dcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatabaseConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dcuo.mutation = mutation
			node, err = dcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dcuo.hooks) - 1; i >= 0; i-- {
			if dcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dcuo *DatabaseConfigUpdateOne) SaveX(ctx context.Context) *DatabaseConfig {
	node, err := dcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dcuo *DatabaseConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := dcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcuo *DatabaseConfigUpdateOne) ExecX(ctx context.Context) {
	if err := dcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dcuo *DatabaseConfigUpdateOne) sqlSave(ctx context.Context) (_node *DatabaseConfig, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   databaseconfig.Table,
			Columns: databaseconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: databaseconfig.FieldID,
			},
		},
	}
	id, ok := dcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DatabaseConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, databaseconfig.FieldID)
		for _, f := range fields {
			if !databaseconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != databaseconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: databaseconfig.FieldURL,
		})
	}
	_node = &DatabaseConfig{config: dcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{databaseconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
