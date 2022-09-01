// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/dosedescription"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// DoseDescriptionQuery is the builder for querying DoseDescription entities.
type DoseDescriptionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DoseDescription
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DoseDescriptionQuery builder.
func (ddq *DoseDescriptionQuery) Where(ps ...predicate.DoseDescription) *DoseDescriptionQuery {
	ddq.predicates = append(ddq.predicates, ps...)
	return ddq
}

// Limit adds a limit step to the query.
func (ddq *DoseDescriptionQuery) Limit(limit int) *DoseDescriptionQuery {
	ddq.limit = &limit
	return ddq
}

// Offset adds an offset step to the query.
func (ddq *DoseDescriptionQuery) Offset(offset int) *DoseDescriptionQuery {
	ddq.offset = &offset
	return ddq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ddq *DoseDescriptionQuery) Unique(unique bool) *DoseDescriptionQuery {
	ddq.unique = &unique
	return ddq
}

// Order adds an order step to the query.
func (ddq *DoseDescriptionQuery) Order(o ...OrderFunc) *DoseDescriptionQuery {
	ddq.order = append(ddq.order, o...)
	return ddq
}

// First returns the first DoseDescription entity from the query.
// Returns a *NotFoundError when no DoseDescription was found.
func (ddq *DoseDescriptionQuery) First(ctx context.Context) (*DoseDescription, error) {
	nodes, err := ddq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dosedescription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) FirstX(ctx context.Context) *DoseDescription {
	node, err := ddq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DoseDescription ID from the query.
// Returns a *NotFoundError when no DoseDescription ID was found.
func (ddq *DoseDescriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ddq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dosedescription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) FirstIDX(ctx context.Context) int {
	id, err := ddq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DoseDescription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DoseDescription entity is found.
// Returns a *NotFoundError when no DoseDescription entities are found.
func (ddq *DoseDescriptionQuery) Only(ctx context.Context) (*DoseDescription, error) {
	nodes, err := ddq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dosedescription.Label}
	default:
		return nil, &NotSingularError{dosedescription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) OnlyX(ctx context.Context) *DoseDescription {
	node, err := ddq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DoseDescription ID in the query.
// Returns a *NotSingularError when more than one DoseDescription ID is found.
// Returns a *NotFoundError when no entities are found.
func (ddq *DoseDescriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ddq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dosedescription.Label}
	default:
		err = &NotSingularError{dosedescription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) OnlyIDX(ctx context.Context) int {
	id, err := ddq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DoseDescriptions.
func (ddq *DoseDescriptionQuery) All(ctx context.Context) ([]*DoseDescription, error) {
	if err := ddq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ddq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) AllX(ctx context.Context) []*DoseDescription {
	nodes, err := ddq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DoseDescription IDs.
func (ddq *DoseDescriptionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ddq.Select(dosedescription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := ddq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ddq *DoseDescriptionQuery) Count(ctx context.Context) (int, error) {
	if err := ddq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ddq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) CountX(ctx context.Context) int {
	count, err := ddq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ddq *DoseDescriptionQuery) Exist(ctx context.Context) (bool, error) {
	if err := ddq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ddq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ddq *DoseDescriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := ddq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DoseDescriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ddq *DoseDescriptionQuery) Clone() *DoseDescriptionQuery {
	if ddq == nil {
		return nil
	}
	return &DoseDescriptionQuery{
		config:     ddq.config,
		limit:      ddq.limit,
		offset:     ddq.offset,
		order:      append([]OrderFunc{}, ddq.order...),
		predicates: append([]predicate.DoseDescription{}, ddq.predicates...),
		// clone intermediate query.
		sql:    ddq.sql.Clone(),
		path:   ddq.path,
		unique: ddq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DoseDescription.Query().
//		GroupBy(dosedescription.FieldName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ddq *DoseDescriptionQuery) GroupBy(field string, fields ...string) *DoseDescriptionGroupBy {
	grbuild := &DoseDescriptionGroupBy{config: ddq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ddq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ddq.sqlQuery(ctx), nil
	}
	grbuild.label = dosedescription.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.DoseDescription.Query().
//		Select(dosedescription.FieldName).
//		Scan(ctx, &v)
//
func (ddq *DoseDescriptionQuery) Select(fields ...string) *DoseDescriptionSelect {
	ddq.fields = append(ddq.fields, fields...)
	selbuild := &DoseDescriptionSelect{DoseDescriptionQuery: ddq}
	selbuild.label = dosedescription.Label
	selbuild.flds, selbuild.scan = &ddq.fields, selbuild.Scan
	return selbuild
}

func (ddq *DoseDescriptionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ddq.fields {
		if !dosedescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ddq.path != nil {
		prev, err := ddq.path(ctx)
		if err != nil {
			return err
		}
		ddq.sql = prev
	}
	return nil
}

func (ddq *DoseDescriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DoseDescription, error) {
	var (
		nodes = []*DoseDescription{}
		_spec = ddq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DoseDescription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DoseDescription{config: ddq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ddq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ddq *DoseDescriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ddq.querySpec()
	_spec.Node.Columns = ddq.fields
	if len(ddq.fields) > 0 {
		_spec.Unique = ddq.unique != nil && *ddq.unique
	}
	return sqlgraph.CountNodes(ctx, ddq.driver, _spec)
}

func (ddq *DoseDescriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ddq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ddq *DoseDescriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dosedescription.Table,
			Columns: dosedescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dosedescription.FieldID,
			},
		},
		From:   ddq.sql,
		Unique: true,
	}
	if unique := ddq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ddq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dosedescription.FieldID)
		for i := range fields {
			if fields[i] != dosedescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ddq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ddq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ddq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ddq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ddq *DoseDescriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ddq.driver.Dialect())
	t1 := builder.Table(dosedescription.Table)
	columns := ddq.fields
	if len(columns) == 0 {
		columns = dosedescription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ddq.sql != nil {
		selector = ddq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ddq.unique != nil && *ddq.unique {
		selector.Distinct()
	}
	for _, p := range ddq.predicates {
		p(selector)
	}
	for _, p := range ddq.order {
		p(selector)
	}
	if offset := ddq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ddq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DoseDescriptionGroupBy is the group-by builder for DoseDescription entities.
type DoseDescriptionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ddgb *DoseDescriptionGroupBy) Aggregate(fns ...AggregateFunc) *DoseDescriptionGroupBy {
	ddgb.fns = append(ddgb.fns, fns...)
	return ddgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ddgb *DoseDescriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ddgb.path(ctx)
	if err != nil {
		return err
	}
	ddgb.sql = query
	return ddgb.sqlScan(ctx, v)
}

func (ddgb *DoseDescriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ddgb.fields {
		if !dosedescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ddgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ddgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ddgb *DoseDescriptionGroupBy) sqlQuery() *sql.Selector {
	selector := ddgb.sql.Select()
	aggregation := make([]string, 0, len(ddgb.fns))
	for _, fn := range ddgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ddgb.fields)+len(ddgb.fns))
		for _, f := range ddgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ddgb.fields...)...)
}

// DoseDescriptionSelect is the builder for selecting fields of DoseDescription entities.
type DoseDescriptionSelect struct {
	*DoseDescriptionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dds *DoseDescriptionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dds.prepareQuery(ctx); err != nil {
		return err
	}
	dds.sql = dds.DoseDescriptionQuery.sqlQuery(ctx)
	return dds.sqlScan(ctx, v)
}

func (dds *DoseDescriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dds.sql.Query()
	if err := dds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
