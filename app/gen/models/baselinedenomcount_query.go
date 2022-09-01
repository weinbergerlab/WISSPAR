// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineDenomCountQuery is the builder for querying BaselineDenomCount entities.
type BaselineDenomCountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineDenomCount
	// eager-loading edges.
	withParent *BaselineDenomQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineDenomCountQuery builder.
func (bdcq *BaselineDenomCountQuery) Where(ps ...predicate.BaselineDenomCount) *BaselineDenomCountQuery {
	bdcq.predicates = append(bdcq.predicates, ps...)
	return bdcq
}

// Limit adds a limit step to the query.
func (bdcq *BaselineDenomCountQuery) Limit(limit int) *BaselineDenomCountQuery {
	bdcq.limit = &limit
	return bdcq
}

// Offset adds an offset step to the query.
func (bdcq *BaselineDenomCountQuery) Offset(offset int) *BaselineDenomCountQuery {
	bdcq.offset = &offset
	return bdcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bdcq *BaselineDenomCountQuery) Unique(unique bool) *BaselineDenomCountQuery {
	bdcq.unique = &unique
	return bdcq
}

// Order adds an order step to the query.
func (bdcq *BaselineDenomCountQuery) Order(o ...OrderFunc) *BaselineDenomCountQuery {
	bdcq.order = append(bdcq.order, o...)
	return bdcq
}

// QueryParent chains the current query on the "parent" edge.
func (bdcq *BaselineDenomCountQuery) QueryParent() *BaselineDenomQuery {
	query := &BaselineDenomQuery{config: bdcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bdcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinedenomcount.Table, baselinedenomcount.FieldID, selector),
			sqlgraph.To(baselinedenom.Table, baselinedenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinedenomcount.ParentTable, baselinedenomcount.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bdcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineDenomCount entity from the query.
// Returns a *NotFoundError when no BaselineDenomCount was found.
func (bdcq *BaselineDenomCountQuery) First(ctx context.Context) (*BaselineDenomCount, error) {
	nodes, err := bdcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinedenomcount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) FirstX(ctx context.Context) *BaselineDenomCount {
	node, err := bdcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineDenomCount ID from the query.
// Returns a *NotFoundError when no BaselineDenomCount ID was found.
func (bdcq *BaselineDenomCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bdcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinedenomcount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) FirstIDX(ctx context.Context) int {
	id, err := bdcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineDenomCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineDenomCount entity is found.
// Returns a *NotFoundError when no BaselineDenomCount entities are found.
func (bdcq *BaselineDenomCountQuery) Only(ctx context.Context) (*BaselineDenomCount, error) {
	nodes, err := bdcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinedenomcount.Label}
	default:
		return nil, &NotSingularError{baselinedenomcount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) OnlyX(ctx context.Context) *BaselineDenomCount {
	node, err := bdcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineDenomCount ID in the query.
// Returns a *NotSingularError when more than one BaselineDenomCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (bdcq *BaselineDenomCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bdcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinedenomcount.Label}
	default:
		err = &NotSingularError{baselinedenomcount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := bdcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineDenomCounts.
func (bdcq *BaselineDenomCountQuery) All(ctx context.Context) ([]*BaselineDenomCount, error) {
	if err := bdcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bdcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) AllX(ctx context.Context) []*BaselineDenomCount {
	nodes, err := bdcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineDenomCount IDs.
func (bdcq *BaselineDenomCountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bdcq.Select(baselinedenomcount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) IDsX(ctx context.Context) []int {
	ids, err := bdcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bdcq *BaselineDenomCountQuery) Count(ctx context.Context) (int, error) {
	if err := bdcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bdcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) CountX(ctx context.Context) int {
	count, err := bdcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bdcq *BaselineDenomCountQuery) Exist(ctx context.Context) (bool, error) {
	if err := bdcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bdcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bdcq *BaselineDenomCountQuery) ExistX(ctx context.Context) bool {
	exist, err := bdcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineDenomCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bdcq *BaselineDenomCountQuery) Clone() *BaselineDenomCountQuery {
	if bdcq == nil {
		return nil
	}
	return &BaselineDenomCountQuery{
		config:     bdcq.config,
		limit:      bdcq.limit,
		offset:     bdcq.offset,
		order:      append([]OrderFunc{}, bdcq.order...),
		predicates: append([]predicate.BaselineDenomCount{}, bdcq.predicates...),
		withParent: bdcq.withParent.Clone(),
		// clone intermediate query.
		sql:    bdcq.sql.Clone(),
		path:   bdcq.path,
		unique: bdcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bdcq *BaselineDenomCountQuery) WithParent(opts ...func(*BaselineDenomQuery)) *BaselineDenomCountQuery {
	query := &BaselineDenomQuery{config: bdcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bdcq.withParent = query
	return bdcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineDenomCountGroupID string `json:"baseline_denom_count_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineDenomCount.Query().
//		GroupBy(baselinedenomcount.FieldBaselineDenomCountGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bdcq *BaselineDenomCountQuery) GroupBy(field string, fields ...string) *BaselineDenomCountGroupBy {
	grbuild := &BaselineDenomCountGroupBy{config: bdcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bdcq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinedenomcount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineDenomCountGroupID string `json:"baseline_denom_count_group_id,omitempty"`
//	}
//
//	client.BaselineDenomCount.Query().
//		Select(baselinedenomcount.FieldBaselineDenomCountGroupID).
//		Scan(ctx, &v)
//
func (bdcq *BaselineDenomCountQuery) Select(fields ...string) *BaselineDenomCountSelect {
	bdcq.fields = append(bdcq.fields, fields...)
	selbuild := &BaselineDenomCountSelect{BaselineDenomCountQuery: bdcq}
	selbuild.label = baselinedenomcount.Label
	selbuild.flds, selbuild.scan = &bdcq.fields, selbuild.Scan
	return selbuild
}

func (bdcq *BaselineDenomCountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bdcq.fields {
		if !baselinedenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bdcq.path != nil {
		prev, err := bdcq.path(ctx)
		if err != nil {
			return err
		}
		bdcq.sql = prev
	}
	return nil
}

func (bdcq *BaselineDenomCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineDenomCount, error) {
	var (
		nodes       = []*BaselineDenomCount{}
		withFKs     = bdcq.withFKs
		_spec       = bdcq.querySpec()
		loadedTypes = [1]bool{
			bdcq.withParent != nil,
		}
	)
	if bdcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinedenomcount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineDenomCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineDenomCount{config: bdcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bdcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bdcq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineDenomCount)
		for i := range nodes {
			if nodes[i].baseline_denom_baseline_denom_count_list == nil {
				continue
			}
			fk := *nodes[i].baseline_denom_baseline_denom_count_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselinedenom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_denom_baseline_denom_count_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (bdcq *BaselineDenomCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bdcq.querySpec()
	_spec.Node.Columns = bdcq.fields
	if len(bdcq.fields) > 0 {
		_spec.Unique = bdcq.unique != nil && *bdcq.unique
	}
	return sqlgraph.CountNodes(ctx, bdcq.driver, _spec)
}

func (bdcq *BaselineDenomCountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bdcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bdcq *BaselineDenomCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinedenomcount.Table,
			Columns: baselinedenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenomcount.FieldID,
			},
		},
		From:   bdcq.sql,
		Unique: true,
	}
	if unique := bdcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bdcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinedenomcount.FieldID)
		for i := range fields {
			if fields[i] != baselinedenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bdcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bdcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bdcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bdcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bdcq *BaselineDenomCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bdcq.driver.Dialect())
	t1 := builder.Table(baselinedenomcount.Table)
	columns := bdcq.fields
	if len(columns) == 0 {
		columns = baselinedenomcount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bdcq.sql != nil {
		selector = bdcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bdcq.unique != nil && *bdcq.unique {
		selector.Distinct()
	}
	for _, p := range bdcq.predicates {
		p(selector)
	}
	for _, p := range bdcq.order {
		p(selector)
	}
	if offset := bdcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bdcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineDenomCountGroupBy is the group-by builder for BaselineDenomCount entities.
type BaselineDenomCountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bdcgb *BaselineDenomCountGroupBy) Aggregate(fns ...AggregateFunc) *BaselineDenomCountGroupBy {
	bdcgb.fns = append(bdcgb.fns, fns...)
	return bdcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bdcgb *BaselineDenomCountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bdcgb.path(ctx)
	if err != nil {
		return err
	}
	bdcgb.sql = query
	return bdcgb.sqlScan(ctx, v)
}

func (bdcgb *BaselineDenomCountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bdcgb.fields {
		if !baselinedenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bdcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bdcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bdcgb *BaselineDenomCountGroupBy) sqlQuery() *sql.Selector {
	selector := bdcgb.sql.Select()
	aggregation := make([]string, 0, len(bdcgb.fns))
	for _, fn := range bdcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bdcgb.fields)+len(bdcgb.fns))
		for _, f := range bdcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bdcgb.fields...)...)
}

// BaselineDenomCountSelect is the builder for selecting fields of BaselineDenomCount entities.
type BaselineDenomCountSelect struct {
	*BaselineDenomCountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bdcs *BaselineDenomCountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bdcs.prepareQuery(ctx); err != nil {
		return err
	}
	bdcs.sql = bdcs.BaselineDenomCountQuery.sqlQuery(ctx)
	return bdcs.sqlScan(ctx, v)
}

func (bdcs *BaselineDenomCountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bdcs.sql.Query()
	if err := bdcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
