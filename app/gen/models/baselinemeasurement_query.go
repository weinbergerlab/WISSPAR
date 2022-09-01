// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasurementQuery is the builder for querying BaselineMeasurement entities.
type BaselineMeasurementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineMeasurement
	// eager-loading edges.
	withParent *BaselineCategoryQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineMeasurementQuery builder.
func (bmq *BaselineMeasurementQuery) Where(ps ...predicate.BaselineMeasurement) *BaselineMeasurementQuery {
	bmq.predicates = append(bmq.predicates, ps...)
	return bmq
}

// Limit adds a limit step to the query.
func (bmq *BaselineMeasurementQuery) Limit(limit int) *BaselineMeasurementQuery {
	bmq.limit = &limit
	return bmq
}

// Offset adds an offset step to the query.
func (bmq *BaselineMeasurementQuery) Offset(offset int) *BaselineMeasurementQuery {
	bmq.offset = &offset
	return bmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bmq *BaselineMeasurementQuery) Unique(unique bool) *BaselineMeasurementQuery {
	bmq.unique = &unique
	return bmq
}

// Order adds an order step to the query.
func (bmq *BaselineMeasurementQuery) Order(o ...OrderFunc) *BaselineMeasurementQuery {
	bmq.order = append(bmq.order, o...)
	return bmq
}

// QueryParent chains the current query on the "parent" edge.
func (bmq *BaselineMeasurementQuery) QueryParent() *BaselineCategoryQuery {
	query := &BaselineCategoryQuery{config: bmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasurement.Table, baselinemeasurement.FieldID, selector),
			sqlgraph.To(baselinecategory.Table, baselinecategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasurement.ParentTable, baselinemeasurement.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineMeasurement entity from the query.
// Returns a *NotFoundError when no BaselineMeasurement was found.
func (bmq *BaselineMeasurementQuery) First(ctx context.Context) (*BaselineMeasurement, error) {
	nodes, err := bmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinemeasurement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) FirstX(ctx context.Context) *BaselineMeasurement {
	node, err := bmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineMeasurement ID from the query.
// Returns a *NotFoundError when no BaselineMeasurement ID was found.
func (bmq *BaselineMeasurementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinemeasurement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) FirstIDX(ctx context.Context) int {
	id, err := bmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineMeasurement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineMeasurement entity is found.
// Returns a *NotFoundError when no BaselineMeasurement entities are found.
func (bmq *BaselineMeasurementQuery) Only(ctx context.Context) (*BaselineMeasurement, error) {
	nodes, err := bmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinemeasurement.Label}
	default:
		return nil, &NotSingularError{baselinemeasurement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) OnlyX(ctx context.Context) *BaselineMeasurement {
	node, err := bmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineMeasurement ID in the query.
// Returns a *NotSingularError when more than one BaselineMeasurement ID is found.
// Returns a *NotFoundError when no entities are found.
func (bmq *BaselineMeasurementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinemeasurement.Label}
	default:
		err = &NotSingularError{baselinemeasurement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) OnlyIDX(ctx context.Context) int {
	id, err := bmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineMeasurements.
func (bmq *BaselineMeasurementQuery) All(ctx context.Context) ([]*BaselineMeasurement, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) AllX(ctx context.Context) []*BaselineMeasurement {
	nodes, err := bmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineMeasurement IDs.
func (bmq *BaselineMeasurementQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bmq.Select(baselinemeasurement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) IDsX(ctx context.Context) []int {
	ids, err := bmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bmq *BaselineMeasurementQuery) Count(ctx context.Context) (int, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) CountX(ctx context.Context) int {
	count, err := bmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bmq *BaselineMeasurementQuery) Exist(ctx context.Context) (bool, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bmq *BaselineMeasurementQuery) ExistX(ctx context.Context) bool {
	exist, err := bmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineMeasurementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bmq *BaselineMeasurementQuery) Clone() *BaselineMeasurementQuery {
	if bmq == nil {
		return nil
	}
	return &BaselineMeasurementQuery{
		config:     bmq.config,
		limit:      bmq.limit,
		offset:     bmq.offset,
		order:      append([]OrderFunc{}, bmq.order...),
		predicates: append([]predicate.BaselineMeasurement{}, bmq.predicates...),
		withParent: bmq.withParent.Clone(),
		// clone intermediate query.
		sql:    bmq.sql.Clone(),
		path:   bmq.path,
		unique: bmq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bmq *BaselineMeasurementQuery) WithParent(opts ...func(*BaselineCategoryQuery)) *BaselineMeasurementQuery {
	query := &BaselineCategoryQuery{config: bmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmq.withParent = query
	return bmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineMeasurementGroupID string `json:"baseline_measurement_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineMeasurement.Query().
//		GroupBy(baselinemeasurement.FieldBaselineMeasurementGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bmq *BaselineMeasurementQuery) GroupBy(field string, fields ...string) *BaselineMeasurementGroupBy {
	grbuild := &BaselineMeasurementGroupBy{config: bmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bmq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinemeasurement.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineMeasurementGroupID string `json:"baseline_measurement_group_id,omitempty"`
//	}
//
//	client.BaselineMeasurement.Query().
//		Select(baselinemeasurement.FieldBaselineMeasurementGroupID).
//		Scan(ctx, &v)
//
func (bmq *BaselineMeasurementQuery) Select(fields ...string) *BaselineMeasurementSelect {
	bmq.fields = append(bmq.fields, fields...)
	selbuild := &BaselineMeasurementSelect{BaselineMeasurementQuery: bmq}
	selbuild.label = baselinemeasurement.Label
	selbuild.flds, selbuild.scan = &bmq.fields, selbuild.Scan
	return selbuild
}

func (bmq *BaselineMeasurementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bmq.fields {
		if !baselinemeasurement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bmq.path != nil {
		prev, err := bmq.path(ctx)
		if err != nil {
			return err
		}
		bmq.sql = prev
	}
	return nil
}

func (bmq *BaselineMeasurementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineMeasurement, error) {
	var (
		nodes       = []*BaselineMeasurement{}
		withFKs     = bmq.withFKs
		_spec       = bmq.querySpec()
		loadedTypes = [1]bool{
			bmq.withParent != nil,
		}
	)
	if bmq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasurement.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineMeasurement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineMeasurement{config: bmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bmq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineMeasurement)
		for i := range nodes {
			if nodes[i].baseline_category_baseline_measurement_list == nil {
				continue
			}
			fk := *nodes[i].baseline_category_baseline_measurement_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselinecategory.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_category_baseline_measurement_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (bmq *BaselineMeasurementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bmq.querySpec()
	_spec.Node.Columns = bmq.fields
	if len(bmq.fields) > 0 {
		_spec.Unique = bmq.unique != nil && *bmq.unique
	}
	return sqlgraph.CountNodes(ctx, bmq.driver, _spec)
}

func (bmq *BaselineMeasurementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bmq *BaselineMeasurementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasurement.Table,
			Columns: baselinemeasurement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasurement.FieldID,
			},
		},
		From:   bmq.sql,
		Unique: true,
	}
	if unique := bmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasurement.FieldID)
		for i := range fields {
			if fields[i] != baselinemeasurement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bmq *BaselineMeasurementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bmq.driver.Dialect())
	t1 := builder.Table(baselinemeasurement.Table)
	columns := bmq.fields
	if len(columns) == 0 {
		columns = baselinemeasurement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bmq.sql != nil {
		selector = bmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bmq.unique != nil && *bmq.unique {
		selector.Distinct()
	}
	for _, p := range bmq.predicates {
		p(selector)
	}
	for _, p := range bmq.order {
		p(selector)
	}
	if offset := bmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineMeasurementGroupBy is the group-by builder for BaselineMeasurement entities.
type BaselineMeasurementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bmgb *BaselineMeasurementGroupBy) Aggregate(fns ...AggregateFunc) *BaselineMeasurementGroupBy {
	bmgb.fns = append(bmgb.fns, fns...)
	return bmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bmgb *BaselineMeasurementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bmgb.path(ctx)
	if err != nil {
		return err
	}
	bmgb.sql = query
	return bmgb.sqlScan(ctx, v)
}

func (bmgb *BaselineMeasurementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bmgb.fields {
		if !baselinemeasurement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bmgb *BaselineMeasurementGroupBy) sqlQuery() *sql.Selector {
	selector := bmgb.sql.Select()
	aggregation := make([]string, 0, len(bmgb.fns))
	for _, fn := range bmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bmgb.fields)+len(bmgb.fns))
		for _, f := range bmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bmgb.fields...)...)
}

// BaselineMeasurementSelect is the builder for selecting fields of BaselineMeasurement entities.
type BaselineMeasurementSelect struct {
	*BaselineMeasurementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bms *BaselineMeasurementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bms.prepareQuery(ctx); err != nil {
		return err
	}
	bms.sql = bms.BaselineMeasurementQuery.sqlQuery(ctx)
	return bms.sqlScan(ctx, v)
}

func (bms *BaselineMeasurementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bms.sql.Query()
	if err := bms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
