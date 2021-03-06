// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/yonidavidson/ent-hooks-examples/ent/migrate"

	"github.com/yonidavidson/ent-hooks-examples/ent/cache"
	"github.com/yonidavidson/ent-hooks-examples/ent/dog"
	"github.com/yonidavidson/ent-hooks-examples/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cache is the client for interacting with the Cache builders.
	Cache *CacheClient
	// Dog is the client for interacting with the Dog builders.
	Dog *DogClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Cache = NewCacheClient(c.config)
	c.Dog = NewDogClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Cache:  NewCacheClient(cfg),
		Dog:    NewDogClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Cache:  NewCacheClient(cfg),
		Dog:    NewDogClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cache.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Cache.Use(hooks...)
	c.Dog.Use(hooks...)
	c.User.Use(hooks...)
}

// CacheClient is a client for the Cache schema.
type CacheClient struct {
	config
}

// NewCacheClient returns a client for the Cache from the given config.
func NewCacheClient(c config) *CacheClient {
	return &CacheClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cache.Hooks(f(g(h())))`.
func (c *CacheClient) Use(hooks ...Hook) {
	c.hooks.Cache = append(c.hooks.Cache, hooks...)
}

// Create returns a create builder for Cache.
func (c *CacheClient) Create() *CacheCreate {
	mutation := newCacheMutation(c.config, OpCreate)
	return &CacheCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cache entities.
func (c *CacheClient) CreateBulk(builders ...*CacheCreate) *CacheCreateBulk {
	return &CacheCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cache.
func (c *CacheClient) Update() *CacheUpdate {
	mutation := newCacheMutation(c.config, OpUpdate)
	return &CacheUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CacheClient) UpdateOne(ca *Cache) *CacheUpdateOne {
	mutation := newCacheMutation(c.config, OpUpdateOne, withCache(ca))
	return &CacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CacheClient) UpdateOneID(id int) *CacheUpdateOne {
	mutation := newCacheMutation(c.config, OpUpdateOne, withCacheID(id))
	return &CacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cache.
func (c *CacheClient) Delete() *CacheDelete {
	mutation := newCacheMutation(c.config, OpDelete)
	return &CacheDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CacheClient) DeleteOne(ca *Cache) *CacheDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CacheClient) DeleteOneID(id int) *CacheDeleteOne {
	builder := c.Delete().Where(cache.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CacheDeleteOne{builder}
}

// Query returns a query builder for Cache.
func (c *CacheClient) Query() *CacheQuery {
	return &CacheQuery{
		config: c.config,
	}
}

// Get returns a Cache entity by its id.
func (c *CacheClient) Get(ctx context.Context, id int) (*Cache, error) {
	return c.Query().Where(cache.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CacheClient) GetX(ctx context.Context, id int) *Cache {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CacheClient) Hooks() []Hook {
	return c.hooks.Cache
}

// DogClient is a client for the Dog schema.
type DogClient struct {
	config
}

// NewDogClient returns a client for the Dog from the given config.
func NewDogClient(c config) *DogClient {
	return &DogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dog.Hooks(f(g(h())))`.
func (c *DogClient) Use(hooks ...Hook) {
	c.hooks.Dog = append(c.hooks.Dog, hooks...)
}

// Create returns a create builder for Dog.
func (c *DogClient) Create() *DogCreate {
	mutation := newDogMutation(c.config, OpCreate)
	return &DogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dog entities.
func (c *DogClient) CreateBulk(builders ...*DogCreate) *DogCreateBulk {
	return &DogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dog.
func (c *DogClient) Update() *DogUpdate {
	mutation := newDogMutation(c.config, OpUpdate)
	return &DogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DogClient) UpdateOne(d *Dog) *DogUpdateOne {
	mutation := newDogMutation(c.config, OpUpdateOne, withDog(d))
	return &DogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DogClient) UpdateOneID(id int) *DogUpdateOne {
	mutation := newDogMutation(c.config, OpUpdateOne, withDogID(id))
	return &DogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dog.
func (c *DogClient) Delete() *DogDelete {
	mutation := newDogMutation(c.config, OpDelete)
	return &DogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DogClient) DeleteOne(d *Dog) *DogDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DogClient) DeleteOneID(id int) *DogDeleteOne {
	builder := c.Delete().Where(dog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DogDeleteOne{builder}
}

// Query returns a query builder for Dog.
func (c *DogClient) Query() *DogQuery {
	return &DogQuery{
		config: c.config,
	}
}

// Get returns a Dog entity by its id.
func (c *DogClient) Get(ctx context.Context, id int) (*Dog, error) {
	return c.Query().Where(dog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DogClient) GetX(ctx context.Context, id int) *Dog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Dog.
func (c *DogClient) QueryOwner(d *Dog) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dog.Table, dog.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dog.OwnerTable, dog.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DogClient) Hooks() []Hook {
	hooks := c.hooks.Dog
	return append(hooks[:len(hooks):len(hooks)], dog.Hooks[:]...)
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPets queries the pets edge of a User.
func (c *UserClient) QueryPets(u *User) *DogQuery {
	query := &DogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(dog.Table, dog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PetsTable, user.PetsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCache queries the cache edge of a User.
func (c *UserClient) QueryCache(u *User) *CacheQuery {
	query := &CacheQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(cache.Table, cache.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.CacheTable, user.CacheColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}
