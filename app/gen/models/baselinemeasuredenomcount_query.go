// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDenomCountQuery is the builder for querying BaselineMeasureDenomCount entities.
type BaselineMeasureDenomCountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineMeasureDenomCount
	// eager-loading edges.
	withParent *BaselineMeasureDenomQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineMeasureDenomCountQuery builder.
func (bmdcq *BaselineMeasureDenomCountQuery) Where(ps ...predicate.BaselineMeasureDenomCount) *BaselineMeasureDenomCountQuery {
	bmdcq.predicates = append(bmdcq.predicates, ps...)
	return bmdcq
}

// Limit adds a limit step to the query.
func (bmdcq *BaselineMeasureDenomCountQuery) Limit(limit int) *BaselineMeasureDenomCountQuery {
	bmdcq.limit = &limit
	return bmdcq
}

// Offset adds an offset step to the query.
func (bmdcq *BaselineMeasureDenomCountQuery) Offset(offset int) *BaselineMeasureDenomCountQuery {
	bmdcq.offset = &offset
	return bmdcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bmdcq *BaselineMeasureDenomCountQuery) Unique(unique bool) *BaselineMeasureDenomCountQuery {
	bmdcq.unique = &unique
	return bmdcq
}

// Order adds an order step to the query.
func (bmdcq *BaselineMeasureDenomCountQuery) Order(o ...OrderFunc) *BaselineMeasureDenomCountQuery {
	bmdcq.order = append(bmdcq.order, o...)
	return bmdcq
}

// QueryParent chains the current query on the "parent" edge.
func (bmdcq *BaselineMeasureDenomCountQuery) QueryParent() *BaselineMeasureDenomQuery {
	query := &BaselineMeasureDenomQuery{config: bmdcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmdcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasuredenomcount.Table, baselinemeasuredenomcount.FieldID, selector),
			sqlgraph.To(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasuredenomcount.ParentTable, baselinemeasuredenomcount.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmdcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineMeasureDenomCount entity from the query.
// Returns a *NotFoundError when no BaselineMeasureDenomCount was found.
func (bmdcq *BaselineMeasureDenomCountQuery) First(ctx context.Context) (*BaselineMeasureDenomCount, error) {
	nodes, err := bmdcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinemeasuredenomcount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) FirstX(ctx context.Context) *BaselineMeasureDenomCount {
	node, err := bmdcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineMeasureDenomCount ID from the query.
// Returns a *NotFoundError when no BaselineMeasureDenomCount ID was found.
func (bmdcq *BaselineMeasureDenomCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmdcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinemeasuredenomcount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) FirstIDX(ctx context.Context) int {
	id, err := bmdcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineMeasureDenomCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineMeasureDenomCount entity is found.
// Returns a *NotFoundError when no BaselineMeasureDenomCount entities are found.
func (bmdcq *BaselineMeasureDenomCountQuery) Only(ctx context.Context) (*BaselineMeasureDenomCount, error) {
	nodes, err := bmdcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinemeasuredenomcount.Label}
	default:
		return nil, &NotSingularError{baselinemeasuredenomcount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) OnlyX(ctx context.Context) *BaselineMeasureDenomCount {
	node, err := bmdcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineMeasureDenomCount ID in the query.
// Returns a *NotSingularError when more than one BaselineMeasureDenomCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (bmdcq *BaselineMeasureDenomCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmdcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinemeasuredenomcount.Label}
	default:
		err = &NotSingularError{baselinemeasuredenomcount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := bmdcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineMeasureDenomCounts.
func (bmdcq *BaselineMeasureDenomCountQuery) All(ctx context.Context) ([]*BaselineMeasureDenomCount, error) {
	if err := bmdcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bmdcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) AllX(ctx context.Context) []*BaselineMeasureDenomCount {
	nodes, err := bmdcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineMeasureDenomCount IDs.
func (bmdcq *BaselineMeasureDenomCountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bmdcq.Select(baselinemeasuredenomcount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) IDsX(ctx context.Context) []int {
	ids, err := bmdcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bmdcq *BaselineMeasureDenomCountQuery) Count(ctx context.Context) (int, error) {
	if err := bmdcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bmdcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) CountX(ctx context.Context) int {
	count, err := bmdcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bmdcq *BaselineMeasureDenomCountQuery) Exist(ctx context.Context) (bool, error) {
	if err := bmdcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bmdcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bmdcq *BaselineMeasureDenomCountQuery) ExistX(ctx context.Context) bool {
	exist, err := bmdcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineMeasureDenomCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bmdcq *BaselineMeasureDenomCountQuery) Clone() *BaselineMeasureDenomCountQuery {
	if bmdcq == nil {
		return nil
	}
	return &BaselineMeasureDenomCountQuery{
		config:     bmdcq.config,
		limit:      bmdcq.limit,
		offset:     bmdcq.offset,
		order:      append([]OrderFunc{}, bmdcq.order...),
		predicates: append([]predicate.BaselineMeasureDenomCount{}, bmdcq.predicates...),
		withParent: bmdcq.withParent.Clone(),
		// clone intermediate query.
		sql:    bmdcq.sql.Clone(),
		path:   bmdcq.path,
		unique: bmdcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bmdcq *BaselineMeasureDenomCountQuery) WithParent(opts ...func(*BaselineMeasureDenomQuery)) *BaselineMeasureDenomCountQuery {
	query := &BaselineMeasureDenomQuery{config: bmdcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmdcq.withParent = query
	return bmdcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineMeasureDenomCountGroupID string `json:"baseline_measure_denom_count_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineMeasureDenomCount.Query().
//		GroupBy(baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bmdcq *BaselineMeasureDenomCountQuery) GroupBy(field string, fields ...string) *BaselineMeasureDenomCountGroupBy {
	grbuild := &BaselineMeasureDenomCountGroupBy{config: bmdcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bmdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bmdcq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinemeasuredenomcount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineMeasureDenomCountGroupID string `json:"baseline_measure_denom_count_group_id,omitempty"`
//	}
//
//	client.BaselineMeasureDenomCount.Query().
//		Select(baselinemeasuredenomcount.FieldBaselineMeasureDenomCountGroupID).
//		Scan(ctx, &v)
//
func (bmdcq *BaselineMeasureDenomCountQuery) Select(fields ...string) *BaselineMeasureDenomCountSelect {
	bmdcq.fields = append(bmdcq.fields, fields...)
	selbuild := &BaselineMeasureDenomCountSelect{BaselineMeasureDenomCountQuery: bmdcq}
	selbuild.label = baselinemeasuredenomcount.Label
	selbuild.flds, selbuild.scan = &bmdcq.fields, selbuild.Scan
	return selbuild
}

func (bmdcq *BaselineMeasureDenomCountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bmdcq.fields {
		if !baselinemeasuredenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bmdcq.path != nil {
		prev, err := bmdcq.path(ctx)
		if err != nil {
			return err
		}
		bmdcq.sql = prev
	}
	return nil
}

func (bmdcq *BaselineMeasureDenomCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineMeasureDenomCount, error) {
	var (
		nodes       = []*BaselineMeasureDenomCount{}
		withFKs     = bmdcq.withFKs
		_spec       = bmdcq.querySpec()
		loadedTypes = [1]bool{
			bmdcq.withParent != nil,
		}
	)
	if bmdcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasuredenomcount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineMeasureDenomCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineMeasureDenomCount{config: bmdcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bmdcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bmdcq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineMeasureDenomCount)
		for i := range nodes {
			if nodes[i].baseline_measure_denom_baseline_measure_denom_count_list == nil {
				continue
			}
			fk := *nodes[i].baseline_measure_denom_baseline_measure_denom_count_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselinemeasuredenom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_measure_denom_baseline_measure_denom_count_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (bmdcq *BaselineMeasureDenomCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bmdcq.querySpec()
	_spec.Node.Columns = bmdcq.fields
	if len(bmdcq.fields) > 0 {
		_spec.Unique = bmdcq.unique != nil && *bmdcq.unique
	}
	return sqlgraph.CountNodes(ctx, bmdcq.driver, _spec)
}

func (bmdcq *BaselineMeasureDenomCountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bmdcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bmdcq *BaselineMeasureDenomCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasuredenomcount.Table,
			Columns: baselinemeasuredenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenomcount.FieldID,
			},
		},
		From:   bmdcq.sql,
		Unique: true,
	}
	if unique := bmdcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bmdcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasuredenomcount.FieldID)
		for i := range fields {
			if fields[i] != baselinemeasuredenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bmdcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bmdcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bmdcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bmdcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bmdcq *BaselineMeasureDenomCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bmdcq.driver.Dialect())
	t1 := builder.Table(baselinemeasuredenomcount.Table)
	columns := bmdcq.fields
	if len(columns) == 0 {
		columns = baselinemeasuredenomcount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bmdcq.sql != nil {
		selector = bmdcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bmdcq.unique != nil && *bmdcq.unique {
		selector.Distinct()
	}
	for _, p := range bmdcq.predicates {
		p(selector)
	}
	for _, p := range bmdcq.order {
		p(selector)
	}
	if offset := bmdcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bmdcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineMeasureDenomCountGroupBy is the group-by builder for BaselineMeasureDenomCount entities.
type BaselineMeasureDenomCountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bmdcgb *BaselineMeasureDenomCountGroupBy) Aggregate(fns ...AggregateFunc) *BaselineMeasureDenomCountGroupBy {
	bmdcgb.fns = append(bmdcgb.fns, fns...)
	return bmdcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bmdcgb *BaselineMeasureDenomCountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bmdcgb.path(ctx)
	if err != nil {
		return err
	}
	bmdcgb.sql = query
	return bmdcgb.sqlScan(ctx, v)
}

func (bmdcgb *BaselineMeasureDenomCountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bmdcgb.fields {
		if !baselinemeasuredenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bmdcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bmdcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bmdcgb *BaselineMeasureDenomCountGroupBy) sqlQuery() *sql.Selector {
	selector := bmdcgb.sql.Select()
	aggregation := make([]string, 0, len(bmdcgb.fns))
	for _, fn := range bmdcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bmdcgb.fields)+len(bmdcgb.fns))
		for _, f := range bmdcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bmdcgb.fields...)...)
}

// BaselineMeasureDenomCountSelect is the builder for selecting fields of BaselineMeasureDenomCount entities.
type BaselineMeasureDenomCountSelect struct {
	*BaselineMeasureDenomCountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bmdcs *BaselineMeasureDenomCountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bmdcs.prepareQuery(ctx); err != nil {
		return err
	}
	bmdcs.sql = bmdcs.BaselineMeasureDenomCountQuery.sqlQuery(ctx)
	return bmdcs.sqlScan(ctx, v)
}

func (bmdcs *BaselineMeasureDenomCountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bmdcs.sql.Query()
	if err := bmdcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
