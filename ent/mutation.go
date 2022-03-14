// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/yonidavidson/ent-side-effect-hooks-example/ent/cache"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent/dog"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent/predicate"
	"github.com/yonidavidson/ent-side-effect-hooks-example/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCache = "Cache"
	TypeDog   = "Dog"
	TypeUser  = "User"
)

// CacheMutation represents an operation that mutates the Cache nodes in the graph.
type CacheMutation struct {
	config
	op            Op
	typ           string
	id            *int
	walks         *int
	addwalks      *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Cache, error)
	predicates    []predicate.Cache
}

var _ ent.Mutation = (*CacheMutation)(nil)

// cacheOption allows management of the mutation configuration using functional options.
type cacheOption func(*CacheMutation)

// newCacheMutation creates new mutation for the Cache entity.
func newCacheMutation(c config, op Op, opts ...cacheOption) *CacheMutation {
	m := &CacheMutation{
		config:        c,
		op:            op,
		typ:           TypeCache,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCacheID sets the ID field of the mutation.
func withCacheID(id int) cacheOption {
	return func(m *CacheMutation) {
		var (
			err   error
			once  sync.Once
			value *Cache
		)
		m.oldValue = func(ctx context.Context) (*Cache, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Cache.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCache sets the old Cache of the mutation.
func withCache(node *Cache) cacheOption {
	return func(m *CacheMutation) {
		m.oldValue = func(context.Context) (*Cache, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CacheMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CacheMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CacheMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CacheMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Cache.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetWalks sets the "walks" field.
func (m *CacheMutation) SetWalks(i int) {
	m.walks = &i
	m.addwalks = nil
}

// Walks returns the value of the "walks" field in the mutation.
func (m *CacheMutation) Walks() (r int, exists bool) {
	v := m.walks
	if v == nil {
		return
	}
	return *v, true
}

// OldWalks returns the old "walks" field's value of the Cache entity.
// If the Cache object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CacheMutation) OldWalks(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWalks is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWalks requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWalks: %w", err)
	}
	return oldValue.Walks, nil
}

// AddWalks adds i to the "walks" field.
func (m *CacheMutation) AddWalks(i int) {
	if m.addwalks != nil {
		*m.addwalks += i
	} else {
		m.addwalks = &i
	}
}

// AddedWalks returns the value that was added to the "walks" field in this mutation.
func (m *CacheMutation) AddedWalks() (r int, exists bool) {
	v := m.addwalks
	if v == nil {
		return
	}
	return *v, true
}

// ResetWalks resets all changes to the "walks" field.
func (m *CacheMutation) ResetWalks() {
	m.walks = nil
	m.addwalks = nil
}

// Where appends a list predicates to the CacheMutation builder.
func (m *CacheMutation) Where(ps ...predicate.Cache) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CacheMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Cache).
func (m *CacheMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CacheMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.walks != nil {
		fields = append(fields, cache.FieldWalks)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CacheMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case cache.FieldWalks:
		return m.Walks()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CacheMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case cache.FieldWalks:
		return m.OldWalks(ctx)
	}
	return nil, fmt.Errorf("unknown Cache field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CacheMutation) SetField(name string, value ent.Value) error {
	switch name {
	case cache.FieldWalks:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWalks(v)
		return nil
	}
	return fmt.Errorf("unknown Cache field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CacheMutation) AddedFields() []string {
	var fields []string
	if m.addwalks != nil {
		fields = append(fields, cache.FieldWalks)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CacheMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case cache.FieldWalks:
		return m.AddedWalks()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CacheMutation) AddField(name string, value ent.Value) error {
	switch name {
	case cache.FieldWalks:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddWalks(v)
		return nil
	}
	return fmt.Errorf("unknown Cache numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CacheMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CacheMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CacheMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Cache nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CacheMutation) ResetField(name string) error {
	switch name {
	case cache.FieldWalks:
		m.ResetWalks()
		return nil
	}
	return fmt.Errorf("unknown Cache field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CacheMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CacheMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CacheMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CacheMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CacheMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CacheMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CacheMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Cache unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CacheMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Cache edge %s", name)
}

// DogMutation represents an operation that mutates the Dog nodes in the graph.
type DogMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Dog, error)
	predicates    []predicate.Dog
}

var _ ent.Mutation = (*DogMutation)(nil)

// dogOption allows management of the mutation configuration using functional options.
type dogOption func(*DogMutation)

// newDogMutation creates new mutation for the Dog entity.
func newDogMutation(c config, op Op, opts ...dogOption) *DogMutation {
	m := &DogMutation{
		config:        c,
		op:            op,
		typ:           TypeDog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDogID sets the ID field of the mutation.
func withDogID(id int) dogOption {
	return func(m *DogMutation) {
		var (
			err   error
			once  sync.Once
			value *Dog
		)
		m.oldValue = func(ctx context.Context) (*Dog, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Dog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDog sets the old Dog of the mutation.
func withDog(node *Dog) dogOption {
	return func(m *DogMutation) {
		m.oldValue = func(context.Context) (*Dog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DogMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DogMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Dog.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *DogMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *DogMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Dog entity.
// If the Dog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DogMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *DogMutation) ResetName() {
	m.name = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *DogMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *DogMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *DogMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *DogMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *DogMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *DogMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the DogMutation builder.
func (m *DogMutation) Where(ps ...predicate.Dog) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DogMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Dog).
func (m *DogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DogMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, dog.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case dog.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case dog.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Dog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case dog.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Dog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DogMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DogMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DogMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Dog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DogMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DogMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Dog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DogMutation) ResetField(name string) error {
	switch name {
	case dog.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Dog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DogMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, dog.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DogMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case dog.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DogMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, dog.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DogMutation) EdgeCleared(name string) bool {
	switch name {
	case dog.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DogMutation) ClearEdge(name string) error {
	switch name {
	case dog.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Dog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DogMutation) ResetEdge(name string) error {
	switch name {
	case dog.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Dog edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	pets          map[int]struct{}
	removedpets   map[int]struct{}
	clearedpets   bool
	cache         *int
	clearedcache  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddPetIDs adds the "pets" edge to the Dog entity by ids.
func (m *UserMutation) AddPetIDs(ids ...int) {
	if m.pets == nil {
		m.pets = make(map[int]struct{})
	}
	for i := range ids {
		m.pets[ids[i]] = struct{}{}
	}
}

// ClearPets clears the "pets" edge to the Dog entity.
func (m *UserMutation) ClearPets() {
	m.clearedpets = true
}

// PetsCleared reports if the "pets" edge to the Dog entity was cleared.
func (m *UserMutation) PetsCleared() bool {
	return m.clearedpets
}

// RemovePetIDs removes the "pets" edge to the Dog entity by IDs.
func (m *UserMutation) RemovePetIDs(ids ...int) {
	if m.removedpets == nil {
		m.removedpets = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.pets, ids[i])
		m.removedpets[ids[i]] = struct{}{}
	}
}

// RemovedPets returns the removed IDs of the "pets" edge to the Dog entity.
func (m *UserMutation) RemovedPetsIDs() (ids []int) {
	for id := range m.removedpets {
		ids = append(ids, id)
	}
	return
}

// PetsIDs returns the "pets" edge IDs in the mutation.
func (m *UserMutation) PetsIDs() (ids []int) {
	for id := range m.pets {
		ids = append(ids, id)
	}
	return
}

// ResetPets resets all changes to the "pets" edge.
func (m *UserMutation) ResetPets() {
	m.pets = nil
	m.clearedpets = false
	m.removedpets = nil
}

// SetCacheID sets the "cache" edge to the Cache entity by id.
func (m *UserMutation) SetCacheID(id int) {
	m.cache = &id
}

// ClearCache clears the "cache" edge to the Cache entity.
func (m *UserMutation) ClearCache() {
	m.clearedcache = true
}

// CacheCleared reports if the "cache" edge to the Cache entity was cleared.
func (m *UserMutation) CacheCleared() bool {
	return m.clearedcache
}

// CacheID returns the "cache" edge ID in the mutation.
func (m *UserMutation) CacheID() (id int, exists bool) {
	if m.cache != nil {
		return *m.cache, true
	}
	return
}

// CacheIDs returns the "cache" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// CacheID instead. It exists only for internal usage by the builders.
func (m *UserMutation) CacheIDs() (ids []int) {
	if id := m.cache; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCache resets all changes to the "cache" edge.
func (m *UserMutation) ResetCache() {
	m.cache = nil
	m.clearedcache = false
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.pets != nil {
		edges = append(edges, user.EdgePets)
	}
	if m.cache != nil {
		edges = append(edges, user.EdgeCache)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePets:
		ids := make([]ent.Value, 0, len(m.pets))
		for id := range m.pets {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeCache:
		if id := m.cache; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedpets != nil {
		edges = append(edges, user.EdgePets)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePets:
		ids := make([]ent.Value, 0, len(m.removedpets))
		for id := range m.removedpets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedpets {
		edges = append(edges, user.EdgePets)
	}
	if m.clearedcache {
		edges = append(edges, user.EdgeCache)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgePets:
		return m.clearedpets
	case user.EdgeCache:
		return m.clearedcache
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeCache:
		m.ClearCache()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgePets:
		m.ResetPets()
		return nil
	case user.EdgeCache:
		m.ResetCache()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
