// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/usecase"
)

// UseCaseQuery is the builder for querying UseCase entities.
type UseCaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UseCase
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UseCaseQuery builder.
func (ucq *UseCaseQuery) Where(ps ...predicate.UseCase) *UseCaseQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit adds a limit step to the query.
func (ucq *UseCaseQuery) Limit(limit int) *UseCaseQuery {
	ucq.limit = &limit
	return ucq
}

// Offset adds an offset step to the query.
func (ucq *UseCaseQuery) Offset(offset int) *UseCaseQuery {
	ucq.offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UseCaseQuery) Unique(unique bool) *UseCaseQuery {
	ucq.unique = &unique
	return ucq
}

// Order adds an order step to the query.
func (ucq *UseCaseQuery) Order(o ...OrderFunc) *UseCaseQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// First returns the first UseCase entity from the query.
// Returns a *NotFoundError when no UseCase was found.
func (ucq *UseCaseQuery) First(ctx context.Context) (*UseCase, error) {
	nodes, err := ucq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usecase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UseCaseQuery) FirstX(ctx context.Context) *UseCase {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UseCase ID from the query.
// Returns a *NotFoundError when no UseCase ID was found.
func (ucq *UseCaseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usecase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UseCaseQuery) FirstIDX(ctx context.Context) int {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UseCase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UseCase entity is found.
// Returns a *NotFoundError when no UseCase entities are found.
func (ucq *UseCaseQuery) Only(ctx context.Context) (*UseCase, error) {
	nodes, err := ucq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usecase.Label}
	default:
		return nil, &NotSingularError{usecase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UseCaseQuery) OnlyX(ctx context.Context) *UseCase {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UseCase ID in the query.
// Returns a *NotSingularError when more than one UseCase ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UseCaseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usecase.Label}
	default:
		err = &NotSingularError{usecase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UseCaseQuery) OnlyIDX(ctx context.Context) int {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UseCases.
func (ucq *UseCaseQuery) All(ctx context.Context) ([]*UseCase, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ucq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UseCaseQuery) AllX(ctx context.Context) []*UseCase {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UseCase IDs.
func (ucq *UseCaseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ucq.Select(usecase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UseCaseQuery) IDsX(ctx context.Context) []int {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UseCaseQuery) Count(ctx context.Context) (int, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ucq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UseCaseQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UseCaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ucq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UseCaseQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UseCaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UseCaseQuery) Clone() *UseCaseQuery {
	if ucq == nil {
		return nil
	}
	return &UseCaseQuery{
		config:     ucq.config,
		limit:      ucq.limit,
		offset:     ucq.offset,
		order:      append([]OrderFunc{}, ucq.order...),
		predicates: append([]predicate.UseCase{}, ucq.predicates...),
		// clone intermediate query.
		sql:    ucq.sql.Clone(),
		path:   ucq.path,
		unique: ucq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UseCaseName string `json:"use_case_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UseCase.Query().
//		GroupBy(usecase.FieldUseCaseName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ucq *UseCaseQuery) GroupBy(field string, fields ...string) *UseCaseGroupBy {
	grbuild := &UseCaseGroupBy{config: ucq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ucq.sqlQuery(ctx), nil
	}
	grbuild.label = usecase.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UseCaseName string `json:"use_case_name,omitempty"`
//	}
//
//	client.UseCase.Query().
//		Select(usecase.FieldUseCaseName).
//		Scan(ctx, &v)
//
func (ucq *UseCaseQuery) Select(fields ...string) *UseCaseSelect {
	ucq.fields = append(ucq.fields, fields...)
	selbuild := &UseCaseSelect{UseCaseQuery: ucq}
	selbuild.label = usecase.Label
	selbuild.flds, selbuild.scan = &ucq.fields, selbuild.Scan
	return selbuild
}

func (ucq *UseCaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ucq.fields {
		if !usecase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UseCaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UseCase, error) {
	var (
		nodes = []*UseCase{}
		_spec = ucq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UseCase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UseCase{config: ucq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ucq *UseCaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	_spec.Node.Columns = ucq.fields
	if len(ucq.fields) > 0 {
		_spec.Unique = ucq.unique != nil && *ucq.unique
	}
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UseCaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ucq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ucq *UseCaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usecase.Table,
			Columns: usecase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usecase.FieldID,
			},
		},
		From:   ucq.sql,
		Unique: true,
	}
	if unique := ucq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ucq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usecase.FieldID)
		for i := range fields {
			if fields[i] != usecase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UseCaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(usecase.Table)
	columns := ucq.fields
	if len(columns) == 0 {
		columns = usecase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucq.unique != nil && *ucq.unique {
		selector.Distinct()
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UseCaseGroupBy is the group-by builder for UseCase entities.
type UseCaseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UseCaseGroupBy) Aggregate(fns ...AggregateFunc) *UseCaseGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ucgb *UseCaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ucgb.path(ctx)
	if err != nil {
		return err
	}
	ucgb.sql = query
	return ucgb.sqlScan(ctx, v)
}

func (ucgb *UseCaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ucgb.fields {
		if !usecase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ucgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ucgb *UseCaseGroupBy) sqlQuery() *sql.Selector {
	selector := ucgb.sql.Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ucgb.fields)+len(ucgb.fns))
		for _, f := range ucgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ucgb.fields...)...)
}

// UseCaseSelect is the builder for selecting fields of UseCase entities.
type UseCaseSelect struct {
	*UseCaseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UseCaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	ucs.sql = ucs.UseCaseQuery.sqlQuery(ctx)
	return ucs.sqlScan(ctx, v)
}

func (ucs *UseCaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ucs.sql.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
