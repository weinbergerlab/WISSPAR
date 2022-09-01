// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassDenomCountQuery is the builder for querying BaselineClassDenomCount entities.
type BaselineClassDenomCountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineClassDenomCount
	// eager-loading edges.
	withParent *BaselineClassDenomQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineClassDenomCountQuery builder.
func (bcdcq *BaselineClassDenomCountQuery) Where(ps ...predicate.BaselineClassDenomCount) *BaselineClassDenomCountQuery {
	bcdcq.predicates = append(bcdcq.predicates, ps...)
	return bcdcq
}

// Limit adds a limit step to the query.
func (bcdcq *BaselineClassDenomCountQuery) Limit(limit int) *BaselineClassDenomCountQuery {
	bcdcq.limit = &limit
	return bcdcq
}

// Offset adds an offset step to the query.
func (bcdcq *BaselineClassDenomCountQuery) Offset(offset int) *BaselineClassDenomCountQuery {
	bcdcq.offset = &offset
	return bcdcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcdcq *BaselineClassDenomCountQuery) Unique(unique bool) *BaselineClassDenomCountQuery {
	bcdcq.unique = &unique
	return bcdcq
}

// Order adds an order step to the query.
func (bcdcq *BaselineClassDenomCountQuery) Order(o ...OrderFunc) *BaselineClassDenomCountQuery {
	bcdcq.order = append(bcdcq.order, o...)
	return bcdcq
}

// QueryParent chains the current query on the "parent" edge.
func (bcdcq *BaselineClassDenomCountQuery) QueryParent() *BaselineClassDenomQuery {
	query := &BaselineClassDenomQuery{config: bcdcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcdcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclassdenomcount.Table, baselineclassdenomcount.FieldID, selector),
			sqlgraph.To(baselineclassdenom.Table, baselineclassdenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselineclassdenomcount.ParentTable, baselineclassdenomcount.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcdcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineClassDenomCount entity from the query.
// Returns a *NotFoundError when no BaselineClassDenomCount was found.
func (bcdcq *BaselineClassDenomCountQuery) First(ctx context.Context) (*BaselineClassDenomCount, error) {
	nodes, err := bcdcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselineclassdenomcount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) FirstX(ctx context.Context) *BaselineClassDenomCount {
	node, err := bcdcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineClassDenomCount ID from the query.
// Returns a *NotFoundError when no BaselineClassDenomCount ID was found.
func (bcdcq *BaselineClassDenomCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcdcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselineclassdenomcount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) FirstIDX(ctx context.Context) int {
	id, err := bcdcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineClassDenomCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineClassDenomCount entity is found.
// Returns a *NotFoundError when no BaselineClassDenomCount entities are found.
func (bcdcq *BaselineClassDenomCountQuery) Only(ctx context.Context) (*BaselineClassDenomCount, error) {
	nodes, err := bcdcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselineclassdenomcount.Label}
	default:
		return nil, &NotSingularError{baselineclassdenomcount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) OnlyX(ctx context.Context) *BaselineClassDenomCount {
	node, err := bcdcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineClassDenomCount ID in the query.
// Returns a *NotSingularError when more than one BaselineClassDenomCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcdcq *BaselineClassDenomCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcdcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselineclassdenomcount.Label}
	default:
		err = &NotSingularError{baselineclassdenomcount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcdcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineClassDenomCounts.
func (bcdcq *BaselineClassDenomCountQuery) All(ctx context.Context) ([]*BaselineClassDenomCount, error) {
	if err := bcdcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcdcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) AllX(ctx context.Context) []*BaselineClassDenomCount {
	nodes, err := bcdcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineClassDenomCount IDs.
func (bcdcq *BaselineClassDenomCountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bcdcq.Select(baselineclassdenomcount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) IDsX(ctx context.Context) []int {
	ids, err := bcdcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcdcq *BaselineClassDenomCountQuery) Count(ctx context.Context) (int, error) {
	if err := bcdcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcdcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) CountX(ctx context.Context) int {
	count, err := bcdcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcdcq *BaselineClassDenomCountQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcdcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcdcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcdcq *BaselineClassDenomCountQuery) ExistX(ctx context.Context) bool {
	exist, err := bcdcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineClassDenomCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcdcq *BaselineClassDenomCountQuery) Clone() *BaselineClassDenomCountQuery {
	if bcdcq == nil {
		return nil
	}
	return &BaselineClassDenomCountQuery{
		config:     bcdcq.config,
		limit:      bcdcq.limit,
		offset:     bcdcq.offset,
		order:      append([]OrderFunc{}, bcdcq.order...),
		predicates: append([]predicate.BaselineClassDenomCount{}, bcdcq.predicates...),
		withParent: bcdcq.withParent.Clone(),
		// clone intermediate query.
		sql:    bcdcq.sql.Clone(),
		path:   bcdcq.path,
		unique: bcdcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bcdcq *BaselineClassDenomCountQuery) WithParent(opts ...func(*BaselineClassDenomQuery)) *BaselineClassDenomCountQuery {
	query := &BaselineClassDenomQuery{config: bcdcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcdcq.withParent = query
	return bcdcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineClassDenomCountGroupID string `json:"baseline_class_denom_count_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineClassDenomCount.Query().
//		GroupBy(baselineclassdenomcount.FieldBaselineClassDenomCountGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bcdcq *BaselineClassDenomCountQuery) GroupBy(field string, fields ...string) *BaselineClassDenomCountGroupBy {
	grbuild := &BaselineClassDenomCountGroupBy{config: bcdcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcdcq.sqlQuery(ctx), nil
	}
	grbuild.label = baselineclassdenomcount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineClassDenomCountGroupID string `json:"baseline_class_denom_count_group_id,omitempty"`
//	}
//
//	client.BaselineClassDenomCount.Query().
//		Select(baselineclassdenomcount.FieldBaselineClassDenomCountGroupID).
//		Scan(ctx, &v)
//
func (bcdcq *BaselineClassDenomCountQuery) Select(fields ...string) *BaselineClassDenomCountSelect {
	bcdcq.fields = append(bcdcq.fields, fields...)
	selbuild := &BaselineClassDenomCountSelect{BaselineClassDenomCountQuery: bcdcq}
	selbuild.label = baselineclassdenomcount.Label
	selbuild.flds, selbuild.scan = &bcdcq.fields, selbuild.Scan
	return selbuild
}

func (bcdcq *BaselineClassDenomCountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcdcq.fields {
		if !baselineclassdenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bcdcq.path != nil {
		prev, err := bcdcq.path(ctx)
		if err != nil {
			return err
		}
		bcdcq.sql = prev
	}
	return nil
}

func (bcdcq *BaselineClassDenomCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineClassDenomCount, error) {
	var (
		nodes       = []*BaselineClassDenomCount{}
		withFKs     = bcdcq.withFKs
		_spec       = bcdcq.querySpec()
		loadedTypes = [1]bool{
			bcdcq.withParent != nil,
		}
	)
	if bcdcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclassdenomcount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineClassDenomCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineClassDenomCount{config: bcdcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcdcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bcdcq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineClassDenomCount)
		for i := range nodes {
			if nodes[i].baseline_class_denom_baseline_class_denom_count_list == nil {
				continue
			}
			fk := *nodes[i].baseline_class_denom_baseline_class_denom_count_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselineclassdenom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_class_denom_baseline_class_denom_count_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (bcdcq *BaselineClassDenomCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcdcq.querySpec()
	_spec.Node.Columns = bcdcq.fields
	if len(bcdcq.fields) > 0 {
		_spec.Unique = bcdcq.unique != nil && *bcdcq.unique
	}
	return sqlgraph.CountNodes(ctx, bcdcq.driver, _spec)
}

func (bcdcq *BaselineClassDenomCountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bcdcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bcdcq *BaselineClassDenomCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclassdenomcount.Table,
			Columns: baselineclassdenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenomcount.FieldID,
			},
		},
		From:   bcdcq.sql,
		Unique: true,
	}
	if unique := bcdcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcdcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclassdenomcount.FieldID)
		for i := range fields {
			if fields[i] != baselineclassdenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcdcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcdcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcdcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcdcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcdcq *BaselineClassDenomCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcdcq.driver.Dialect())
	t1 := builder.Table(baselineclassdenomcount.Table)
	columns := bcdcq.fields
	if len(columns) == 0 {
		columns = baselineclassdenomcount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcdcq.sql != nil {
		selector = bcdcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcdcq.unique != nil && *bcdcq.unique {
		selector.Distinct()
	}
	for _, p := range bcdcq.predicates {
		p(selector)
	}
	for _, p := range bcdcq.order {
		p(selector)
	}
	if offset := bcdcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcdcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineClassDenomCountGroupBy is the group-by builder for BaselineClassDenomCount entities.
type BaselineClassDenomCountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcdcgb *BaselineClassDenomCountGroupBy) Aggregate(fns ...AggregateFunc) *BaselineClassDenomCountGroupBy {
	bcdcgb.fns = append(bcdcgb.fns, fns...)
	return bcdcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcdcgb *BaselineClassDenomCountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bcdcgb.path(ctx)
	if err != nil {
		return err
	}
	bcdcgb.sql = query
	return bcdcgb.sqlScan(ctx, v)
}

func (bcdcgb *BaselineClassDenomCountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bcdcgb.fields {
		if !baselineclassdenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcdcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcdcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcdcgb *BaselineClassDenomCountGroupBy) sqlQuery() *sql.Selector {
	selector := bcdcgb.sql.Select()
	aggregation := make([]string, 0, len(bcdcgb.fns))
	for _, fn := range bcdcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcdcgb.fields)+len(bcdcgb.fns))
		for _, f := range bcdcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcdcgb.fields...)...)
}

// BaselineClassDenomCountSelect is the builder for selecting fields of BaselineClassDenomCount entities.
type BaselineClassDenomCountSelect struct {
	*BaselineClassDenomCountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcdcs *BaselineClassDenomCountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bcdcs.prepareQuery(ctx); err != nil {
		return err
	}
	bcdcs.sql = bcdcs.BaselineClassDenomCountQuery.sqlQuery(ctx)
	return bcdcs.sqlScan(ctx, v)
}

func (bcdcs *BaselineClassDenomCountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bcdcs.sql.Query()
	if err := bcdcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
