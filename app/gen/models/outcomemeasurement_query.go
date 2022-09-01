// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomecategory"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasurement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasurementQuery is the builder for querying OutcomeMeasurement entities.
type OutcomeMeasurementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeMeasurement
	// eager-loading edges.
	withParent *OutcomeCategoryQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeMeasurementQuery builder.
func (omq *OutcomeMeasurementQuery) Where(ps ...predicate.OutcomeMeasurement) *OutcomeMeasurementQuery {
	omq.predicates = append(omq.predicates, ps...)
	return omq
}

// Limit adds a limit step to the query.
func (omq *OutcomeMeasurementQuery) Limit(limit int) *OutcomeMeasurementQuery {
	omq.limit = &limit
	return omq
}

// Offset adds an offset step to the query.
func (omq *OutcomeMeasurementQuery) Offset(offset int) *OutcomeMeasurementQuery {
	omq.offset = &offset
	return omq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (omq *OutcomeMeasurementQuery) Unique(unique bool) *OutcomeMeasurementQuery {
	omq.unique = &unique
	return omq
}

// Order adds an order step to the query.
func (omq *OutcomeMeasurementQuery) Order(o ...OrderFunc) *OutcomeMeasurementQuery {
	omq.order = append(omq.order, o...)
	return omq
}

// QueryParent chains the current query on the "parent" edge.
func (omq *OutcomeMeasurementQuery) QueryParent() *OutcomeCategoryQuery {
	query := &OutcomeCategoryQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasurement.Table, outcomemeasurement.FieldID, selector),
			sqlgraph.To(outcomecategory.Table, outcomecategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomemeasurement.ParentTable, outcomemeasurement.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeMeasurement entity from the query.
// Returns a *NotFoundError when no OutcomeMeasurement was found.
func (omq *OutcomeMeasurementQuery) First(ctx context.Context) (*OutcomeMeasurement, error) {
	nodes, err := omq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomemeasurement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) FirstX(ctx context.Context) *OutcomeMeasurement {
	node, err := omq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeMeasurement ID from the query.
// Returns a *NotFoundError when no OutcomeMeasurement ID was found.
func (omq *OutcomeMeasurementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = omq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomemeasurement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) FirstIDX(ctx context.Context) int {
	id, err := omq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeMeasurement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeMeasurement entity is found.
// Returns a *NotFoundError when no OutcomeMeasurement entities are found.
func (omq *OutcomeMeasurementQuery) Only(ctx context.Context) (*OutcomeMeasurement, error) {
	nodes, err := omq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomemeasurement.Label}
	default:
		return nil, &NotSingularError{outcomemeasurement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) OnlyX(ctx context.Context) *OutcomeMeasurement {
	node, err := omq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeMeasurement ID in the query.
// Returns a *NotSingularError when more than one OutcomeMeasurement ID is found.
// Returns a *NotFoundError when no entities are found.
func (omq *OutcomeMeasurementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = omq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomemeasurement.Label}
	default:
		err = &NotSingularError{outcomemeasurement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) OnlyIDX(ctx context.Context) int {
	id, err := omq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeMeasurements.
func (omq *OutcomeMeasurementQuery) All(ctx context.Context) ([]*OutcomeMeasurement, error) {
	if err := omq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return omq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) AllX(ctx context.Context) []*OutcomeMeasurement {
	nodes, err := omq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeMeasurement IDs.
func (omq *OutcomeMeasurementQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := omq.Select(outcomemeasurement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) IDsX(ctx context.Context) []int {
	ids, err := omq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (omq *OutcomeMeasurementQuery) Count(ctx context.Context) (int, error) {
	if err := omq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return omq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) CountX(ctx context.Context) int {
	count, err := omq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (omq *OutcomeMeasurementQuery) Exist(ctx context.Context) (bool, error) {
	if err := omq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return omq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (omq *OutcomeMeasurementQuery) ExistX(ctx context.Context) bool {
	exist, err := omq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeMeasurementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (omq *OutcomeMeasurementQuery) Clone() *OutcomeMeasurementQuery {
	if omq == nil {
		return nil
	}
	return &OutcomeMeasurementQuery{
		config:     omq.config,
		limit:      omq.limit,
		offset:     omq.offset,
		order:      append([]OrderFunc{}, omq.order...),
		predicates: append([]predicate.OutcomeMeasurement{}, omq.predicates...),
		withParent: omq.withParent.Clone(),
		// clone intermediate query.
		sql:    omq.sql.Clone(),
		path:   omq.path,
		unique: omq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasurementQuery) WithParent(opts ...func(*OutcomeCategoryQuery)) *OutcomeMeasurementQuery {
	query := &OutcomeCategoryQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withParent = query
	return omq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeMeasurementGroupID string `json:"outcome_measurement_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeMeasurement.Query().
//		GroupBy(outcomemeasurement.FieldOutcomeMeasurementGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (omq *OutcomeMeasurementQuery) GroupBy(field string, fields ...string) *OutcomeMeasurementGroupBy {
	grbuild := &OutcomeMeasurementGroupBy{config: omq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return omq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomemeasurement.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeMeasurementGroupID string `json:"outcome_measurement_group_id,omitempty"`
//	}
//
//	client.OutcomeMeasurement.Query().
//		Select(outcomemeasurement.FieldOutcomeMeasurementGroupID).
//		Scan(ctx, &v)
//
func (omq *OutcomeMeasurementQuery) Select(fields ...string) *OutcomeMeasurementSelect {
	omq.fields = append(omq.fields, fields...)
	selbuild := &OutcomeMeasurementSelect{OutcomeMeasurementQuery: omq}
	selbuild.label = outcomemeasurement.Label
	selbuild.flds, selbuild.scan = &omq.fields, selbuild.Scan
	return selbuild
}

func (omq *OutcomeMeasurementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range omq.fields {
		if !outcomemeasurement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if omq.path != nil {
		prev, err := omq.path(ctx)
		if err != nil {
			return err
		}
		omq.sql = prev
	}
	return nil
}

func (omq *OutcomeMeasurementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeMeasurement, error) {
	var (
		nodes       = []*OutcomeMeasurement{}
		withFKs     = omq.withFKs
		_spec       = omq.querySpec()
		loadedTypes = [1]bool{
			omq.withParent != nil,
		}
	)
	if omq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasurement.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeMeasurement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeMeasurement{config: omq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, omq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := omq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeMeasurement)
		for i := range nodes {
			if nodes[i].outcome_category_outcome_measurement_list == nil {
				continue
			}
			fk := *nodes[i].outcome_category_outcome_measurement_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomecategory.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_category_outcome_measurement_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (omq *OutcomeMeasurementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := omq.querySpec()
	_spec.Node.Columns = omq.fields
	if len(omq.fields) > 0 {
		_spec.Unique = omq.unique != nil && *omq.unique
	}
	return sqlgraph.CountNodes(ctx, omq.driver, _spec)
}

func (omq *OutcomeMeasurementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := omq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (omq *OutcomeMeasurementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasurement.Table,
			Columns: outcomemeasurement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasurement.FieldID,
			},
		},
		From:   omq.sql,
		Unique: true,
	}
	if unique := omq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := omq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasurement.FieldID)
		for i := range fields {
			if fields[i] != outcomemeasurement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := omq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := omq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := omq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := omq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (omq *OutcomeMeasurementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(omq.driver.Dialect())
	t1 := builder.Table(outcomemeasurement.Table)
	columns := omq.fields
	if len(columns) == 0 {
		columns = outcomemeasurement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if omq.sql != nil {
		selector = omq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if omq.unique != nil && *omq.unique {
		selector.Distinct()
	}
	for _, p := range omq.predicates {
		p(selector)
	}
	for _, p := range omq.order {
		p(selector)
	}
	if offset := omq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := omq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeMeasurementGroupBy is the group-by builder for OutcomeMeasurement entities.
type OutcomeMeasurementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (omgb *OutcomeMeasurementGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeMeasurementGroupBy {
	omgb.fns = append(omgb.fns, fns...)
	return omgb
}

// Scan applies the group-by query and scans the result into the given value.
func (omgb *OutcomeMeasurementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := omgb.path(ctx)
	if err != nil {
		return err
	}
	omgb.sql = query
	return omgb.sqlScan(ctx, v)
}

func (omgb *OutcomeMeasurementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range omgb.fields {
		if !outcomemeasurement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := omgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := omgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (omgb *OutcomeMeasurementGroupBy) sqlQuery() *sql.Selector {
	selector := omgb.sql.Select()
	aggregation := make([]string, 0, len(omgb.fns))
	for _, fn := range omgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(omgb.fields)+len(omgb.fns))
		for _, f := range omgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(omgb.fields...)...)
}

// OutcomeMeasurementSelect is the builder for selecting fields of OutcomeMeasurement entities.
type OutcomeMeasurementSelect struct {
	*OutcomeMeasurementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oms *OutcomeMeasurementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oms.prepareQuery(ctx); err != nil {
		return err
	}
	oms.sql = oms.OutcomeMeasurementQuery.sqlQuery(ctx)
	return oms.sqlScan(ctx, v)
}

func (oms *OutcomeMeasurementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oms.sql.Query()
	if err := oms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
