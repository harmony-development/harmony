// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/guildlistentry"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// GuildListEntryQuery is the builder for querying GuildListEntry entities.
type GuildListEntryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GuildListEntry
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GuildListEntryQuery builder.
func (gleq *GuildListEntryQuery) Where(ps ...predicate.GuildListEntry) *GuildListEntryQuery {
	gleq.predicates = append(gleq.predicates, ps...)
	return gleq
}

// Limit adds a limit step to the query.
func (gleq *GuildListEntryQuery) Limit(limit int) *GuildListEntryQuery {
	gleq.limit = &limit
	return gleq
}

// Offset adds an offset step to the query.
func (gleq *GuildListEntryQuery) Offset(offset int) *GuildListEntryQuery {
	gleq.offset = &offset
	return gleq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gleq *GuildListEntryQuery) Unique(unique bool) *GuildListEntryQuery {
	gleq.unique = &unique
	return gleq
}

// Order adds an order step to the query.
func (gleq *GuildListEntryQuery) Order(o ...OrderFunc) *GuildListEntryQuery {
	gleq.order = append(gleq.order, o...)
	return gleq
}

// QueryUser chains the current query on the "user" edge.
func (gleq *GuildListEntryQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: gleq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gleq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gleq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(guildlistentry.Table, guildlistentry.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, guildlistentry.UserTable, guildlistentry.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gleq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GuildListEntry entity from the query.
// Returns a *NotFoundError when no GuildListEntry was found.
func (gleq *GuildListEntryQuery) First(ctx context.Context) (*GuildListEntry, error) {
	nodes, err := gleq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{guildlistentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gleq *GuildListEntryQuery) FirstX(ctx context.Context) *GuildListEntry {
	node, err := gleq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GuildListEntry ID from the query.
// Returns a *NotFoundError when no GuildListEntry ID was found.
func (gleq *GuildListEntryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gleq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{guildlistentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gleq *GuildListEntryQuery) FirstIDX(ctx context.Context) int {
	id, err := gleq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GuildListEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GuildListEntry entity is not found.
// Returns a *NotFoundError when no GuildListEntry entities are found.
func (gleq *GuildListEntryQuery) Only(ctx context.Context) (*GuildListEntry, error) {
	nodes, err := gleq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{guildlistentry.Label}
	default:
		return nil, &NotSingularError{guildlistentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gleq *GuildListEntryQuery) OnlyX(ctx context.Context) *GuildListEntry {
	node, err := gleq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GuildListEntry ID in the query.
// Returns a *NotSingularError when exactly one GuildListEntry ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gleq *GuildListEntryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gleq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = &NotSingularError{guildlistentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gleq *GuildListEntryQuery) OnlyIDX(ctx context.Context) int {
	id, err := gleq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GuildListEntries.
func (gleq *GuildListEntryQuery) All(ctx context.Context) ([]*GuildListEntry, error) {
	if err := gleq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gleq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gleq *GuildListEntryQuery) AllX(ctx context.Context) []*GuildListEntry {
	nodes, err := gleq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GuildListEntry IDs.
func (gleq *GuildListEntryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gleq.Select(guildlistentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gleq *GuildListEntryQuery) IDsX(ctx context.Context) []int {
	ids, err := gleq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gleq *GuildListEntryQuery) Count(ctx context.Context) (int, error) {
	if err := gleq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gleq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gleq *GuildListEntryQuery) CountX(ctx context.Context) int {
	count, err := gleq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gleq *GuildListEntryQuery) Exist(ctx context.Context) (bool, error) {
	if err := gleq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gleq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gleq *GuildListEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := gleq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GuildListEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gleq *GuildListEntryQuery) Clone() *GuildListEntryQuery {
	if gleq == nil {
		return nil
	}
	return &GuildListEntryQuery{
		config:     gleq.config,
		limit:      gleq.limit,
		offset:     gleq.offset,
		order:      append([]OrderFunc{}, gleq.order...),
		predicates: append([]predicate.GuildListEntry{}, gleq.predicates...),
		withUser:   gleq.withUser.Clone(),
		// clone intermediate query.
		sql:  gleq.sql.Clone(),
		path: gleq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gleq *GuildListEntryQuery) WithUser(opts ...func(*UserQuery)) *GuildListEntryQuery {
	query := &UserQuery{config: gleq.config}
	for _, opt := range opts {
		opt(query)
	}
	gleq.withUser = query
	return gleq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GuildID uint64 `json:"guild_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GuildListEntry.Query().
//		GroupBy(guildlistentry.FieldGuildID).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
//
func (gleq *GuildListEntryQuery) GroupBy(field string, fields ...string) *GuildListEntryGroupBy {
	group := &GuildListEntryGroupBy{config: gleq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gleq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gleq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GuildID uint64 `json:"guild_id,omitempty"`
//	}
//
//	client.GuildListEntry.Query().
//		Select(guildlistentry.FieldGuildID).
//		Scan(ctx, &v)
//
func (gleq *GuildListEntryQuery) Select(field string, fields ...string) *GuildListEntrySelect {
	gleq.fields = append([]string{field}, fields...)
	return &GuildListEntrySelect{GuildListEntryQuery: gleq}
}

func (gleq *GuildListEntryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gleq.fields {
		if !guildlistentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if gleq.path != nil {
		prev, err := gleq.path(ctx)
		if err != nil {
			return err
		}
		gleq.sql = prev
	}
	return nil
}

func (gleq *GuildListEntryQuery) sqlAll(ctx context.Context) ([]*GuildListEntry, error) {
	var (
		nodes       = []*GuildListEntry{}
		withFKs     = gleq.withFKs
		_spec       = gleq.querySpec()
		loadedTypes = [1]bool{
			gleq.withUser != nil,
		}
	)
	if gleq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, guildlistentry.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GuildListEntry{config: gleq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("entgen: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gleq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gleq.withUser; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*GuildListEntry)
		for i := range nodes {
			if nodes[i].user_listentry == nil {
				continue
			}
			fk := *nodes[i].user_listentry
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_listentry" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (gleq *GuildListEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gleq.querySpec()
	return sqlgraph.CountNodes(ctx, gleq.driver, _spec)
}

func (gleq *GuildListEntryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gleq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entgen: check existence: %w", err)
	}
	return n > 0, nil
}

func (gleq *GuildListEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guildlistentry.Table,
			Columns: guildlistentry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: guildlistentry.FieldID,
			},
		},
		From:   gleq.sql,
		Unique: true,
	}
	if unique := gleq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gleq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guildlistentry.FieldID)
		for i := range fields {
			if fields[i] != guildlistentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gleq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gleq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gleq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gleq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gleq *GuildListEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gleq.driver.Dialect())
	t1 := builder.Table(guildlistentry.Table)
	selector := builder.Select(t1.Columns(guildlistentry.Columns...)...).From(t1)
	if gleq.sql != nil {
		selector = gleq.sql
		selector.Select(selector.Columns(guildlistentry.Columns...)...)
	}
	for _, p := range gleq.predicates {
		p(selector)
	}
	for _, p := range gleq.order {
		p(selector)
	}
	if offset := gleq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gleq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GuildListEntryGroupBy is the group-by builder for GuildListEntry entities.
type GuildListEntryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (glegb *GuildListEntryGroupBy) Aggregate(fns ...AggregateFunc) *GuildListEntryGroupBy {
	glegb.fns = append(glegb.fns, fns...)
	return glegb
}

// Scan applies the group-by query and scans the result into the given value.
func (glegb *GuildListEntryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := glegb.path(ctx)
	if err != nil {
		return err
	}
	glegb.sql = query
	return glegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := glegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(glegb.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := glegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) StringsX(ctx context.Context) []string {
	v, err := glegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = glegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) StringX(ctx context.Context) string {
	v, err := glegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(glegb.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := glegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) IntsX(ctx context.Context) []int {
	v, err := glegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = glegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) IntX(ctx context.Context) int {
	v, err := glegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(glegb.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := glegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := glegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = glegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := glegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(glegb.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := glegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := glegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (glegb *GuildListEntryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = glegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (glegb *GuildListEntryGroupBy) BoolX(ctx context.Context) bool {
	v, err := glegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (glegb *GuildListEntryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range glegb.fields {
		if !guildlistentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := glegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := glegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (glegb *GuildListEntryGroupBy) sqlQuery() *sql.Selector {
	selector := glegb.sql
	columns := make([]string, 0, len(glegb.fields)+len(glegb.fns))
	columns = append(columns, glegb.fields...)
	for _, fn := range glegb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(glegb.fields...)
}

// GuildListEntrySelect is the builder for selecting fields of GuildListEntry entities.
type GuildListEntrySelect struct {
	*GuildListEntryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gles *GuildListEntrySelect) Scan(ctx context.Context, v interface{}) error {
	if err := gles.prepareQuery(ctx); err != nil {
		return err
	}
	gles.sql = gles.GuildListEntryQuery.sqlQuery(ctx)
	return gles.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gles *GuildListEntrySelect) ScanX(ctx context.Context, v interface{}) {
	if err := gles.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Strings(ctx context.Context) ([]string, error) {
	if len(gles.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntrySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gles.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gles *GuildListEntrySelect) StringsX(ctx context.Context) []string {
	v, err := gles.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gles.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntrySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gles *GuildListEntrySelect) StringX(ctx context.Context) string {
	v, err := gles.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Ints(ctx context.Context) ([]int, error) {
	if len(gles.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntrySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gles.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gles *GuildListEntrySelect) IntsX(ctx context.Context) []int {
	v, err := gles.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gles.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntrySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gles *GuildListEntrySelect) IntX(ctx context.Context) int {
	v, err := gles.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gles.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntrySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gles.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gles *GuildListEntrySelect) Float64sX(ctx context.Context) []float64 {
	v, err := gles.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gles.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntrySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gles *GuildListEntrySelect) Float64X(ctx context.Context) float64 {
	v, err := gles.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gles.fields) > 1 {
		return nil, errors.New("entgen: GuildListEntrySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gles.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gles *GuildListEntrySelect) BoolsX(ctx context.Context) []bool {
	v, err := gles.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gles *GuildListEntrySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gles.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{guildlistentry.Label}
	default:
		err = fmt.Errorf("entgen: GuildListEntrySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gles *GuildListEntrySelect) BoolX(ctx context.Context) bool {
	v, err := gles.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gles *GuildListEntrySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gles.sqlQuery().Query()
	if err := gles.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gles *GuildListEntrySelect) sqlQuery() sql.Querier {
	selector := gles.sql
	selector.Select(selector.Columns(gles.fields...)...)
	return selector
}
