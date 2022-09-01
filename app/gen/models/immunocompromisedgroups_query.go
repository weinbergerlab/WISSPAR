// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/immunocompromisedgroups"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ImmunocompromisedGroupsQuery is the builder for querying ImmunocompromisedGroups entities.
type ImmunocompromisedGroupsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ImmunocompromisedGroups
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImmunocompromisedGroupsQuery builder.
func (igq *ImmunocompromisedGroupsQuery) Where(ps ...predicate.ImmunocompromisedGroups) *ImmunocompromisedGroupsQuery {
	igq.predicates = append(igq.predicates, ps...)
	return igq
}

// Limit adds a limit step to the query.
func (igq *ImmunocompromisedGroupsQuery) Limit(limit int) *ImmunocompromisedGroupsQuery {
	igq.limit = &limit
	return igq
}

// Offset adds an offset step to the query.
func (igq *ImmunocompromisedGroupsQuery) Offset(offset int) *ImmunocompromisedGroupsQuery {
	igq.offset = &offset
	return igq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (igq *ImmunocompromisedGroupsQuery) Unique(unique bool) *ImmunocompromisedGroupsQuery {
	igq.unique = &unique
	return igq
}

// Order adds an order step to the query.
func (igq *ImmunocompromisedGroupsQuery) Order(o ...OrderFunc) *ImmunocompromisedGroupsQuery {
	igq.order = append(igq.order, o...)
	return igq
}

// First returns the first ImmunocompromisedGroups entity from the query.
// Returns a *NotFoundError when no ImmunocompromisedGroups was found.
func (igq *ImmunocompromisedGroupsQuery) First(ctx context.Context) (*ImmunocompromisedGroups, error) {
	nodes, err := igq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{immunocompromisedgroups.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) FirstX(ctx context.Context) *ImmunocompromisedGroups {
	node, err := igq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImmunocompromisedGroups ID from the query.
// Returns a *NotFoundError when no ImmunocompromisedGroups ID was found.
func (igq *ImmunocompromisedGroupsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = igq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{immunocompromisedgroups.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) FirstIDX(ctx context.Context) int {
	id, err := igq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImmunocompromisedGroups entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ImmunocompromisedGroups entity is found.
// Returns a *NotFoundError when no ImmunocompromisedGroups entities are found.
func (igq *ImmunocompromisedGroupsQuery) Only(ctx context.Context) (*ImmunocompromisedGroups, error) {
	nodes, err := igq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{immunocompromisedgroups.Label}
	default:
		return nil, &NotSingularError{immunocompromisedgroups.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) OnlyX(ctx context.Context) *ImmunocompromisedGroups {
	node, err := igq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImmunocompromisedGroups ID in the query.
// Returns a *NotSingularError when more than one ImmunocompromisedGroups ID is found.
// Returns a *NotFoundError when no entities are found.
func (igq *ImmunocompromisedGroupsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = igq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{immunocompromisedgroups.Label}
	default:
		err = &NotSingularError{immunocompromisedgroups.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) OnlyIDX(ctx context.Context) int {
	id, err := igq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImmunocompromisedGroupsSlice.
func (igq *ImmunocompromisedGroupsQuery) All(ctx context.Context) ([]*ImmunocompromisedGroups, error) {
	if err := igq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return igq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) AllX(ctx context.Context) []*ImmunocompromisedGroups {
	nodes, err := igq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImmunocompromisedGroups IDs.
func (igq *ImmunocompromisedGroupsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := igq.Select(immunocompromisedgroups.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) IDsX(ctx context.Context) []int {
	ids, err := igq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (igq *ImmunocompromisedGroupsQuery) Count(ctx context.Context) (int, error) {
	if err := igq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return igq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) CountX(ctx context.Context) int {
	count, err := igq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (igq *ImmunocompromisedGroupsQuery) Exist(ctx context.Context) (bool, error) {
	if err := igq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return igq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (igq *ImmunocompromisedGroupsQuery) ExistX(ctx context.Context) bool {
	exist, err := igq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImmunocompromisedGroupsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (igq *ImmunocompromisedGroupsQuery) Clone() *ImmunocompromisedGroupsQuery {
	if igq == nil {
		return nil
	}
	return &ImmunocompromisedGroupsQuery{
		config:     igq.config,
		limit:      igq.limit,
		offset:     igq.offset,
		order:      append([]OrderFunc{}, igq.order...),
		predicates: append([]predicate.ImmunocompromisedGroups{}, igq.predicates...),
		// clone intermediate query.
		sql:    igq.sql.Clone(),
		path:   igq.path,
		unique: igq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GroupName string `json:"group_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ImmunocompromisedGroups.Query().
//		GroupBy(immunocompromisedgroups.FieldGroupName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (igq *ImmunocompromisedGroupsQuery) GroupBy(field string, fields ...string) *ImmunocompromisedGroupsGroupBy {
	grbuild := &ImmunocompromisedGroupsGroupBy{config: igq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := igq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return igq.sqlQuery(ctx), nil
	}
	grbuild.label = immunocompromisedgroups.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GroupName string `json:"group_name,omitempty"`
//	}
//
//	client.ImmunocompromisedGroups.Query().
//		Select(immunocompromisedgroups.FieldGroupName).
//		Scan(ctx, &v)
//
func (igq *ImmunocompromisedGroupsQuery) Select(fields ...string) *ImmunocompromisedGroupsSelect {
	igq.fields = append(igq.fields, fields...)
	selbuild := &ImmunocompromisedGroupsSelect{ImmunocompromisedGroupsQuery: igq}
	selbuild.label = immunocompromisedgroups.Label
	selbuild.flds, selbuild.scan = &igq.fields, selbuild.Scan
	return selbuild
}

func (igq *ImmunocompromisedGroupsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range igq.fields {
		if !immunocompromisedgroups.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if igq.path != nil {
		prev, err := igq.path(ctx)
		if err != nil {
			return err
		}
		igq.sql = prev
	}
	return nil
}

func (igq *ImmunocompromisedGroupsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ImmunocompromisedGroups, error) {
	var (
		nodes = []*ImmunocompromisedGroups{}
		_spec = igq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ImmunocompromisedGroups).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ImmunocompromisedGroups{config: igq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, igq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (igq *ImmunocompromisedGroupsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := igq.querySpec()
	_spec.Node.Columns = igq.fields
	if len(igq.fields) > 0 {
		_spec.Unique = igq.unique != nil && *igq.unique
	}
	return sqlgraph.CountNodes(ctx, igq.driver, _spec)
}

func (igq *ImmunocompromisedGroupsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := igq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (igq *ImmunocompromisedGroupsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   immunocompromisedgroups.Table,
			Columns: immunocompromisedgroups.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: immunocompromisedgroups.FieldID,
			},
		},
		From:   igq.sql,
		Unique: true,
	}
	if unique := igq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := igq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, immunocompromisedgroups.FieldID)
		for i := range fields {
			if fields[i] != immunocompromisedgroups.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := igq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := igq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := igq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := igq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (igq *ImmunocompromisedGroupsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(igq.driver.Dialect())
	t1 := builder.Table(immunocompromisedgroups.Table)
	columns := igq.fields
	if len(columns) == 0 {
		columns = immunocompromisedgroups.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if igq.sql != nil {
		selector = igq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if igq.unique != nil && *igq.unique {
		selector.Distinct()
	}
	for _, p := range igq.predicates {
		p(selector)
	}
	for _, p := range igq.order {
		p(selector)
	}
	if offset := igq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := igq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImmunocompromisedGroupsGroupBy is the group-by builder for ImmunocompromisedGroups entities.
type ImmunocompromisedGroupsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iggb *ImmunocompromisedGroupsGroupBy) Aggregate(fns ...AggregateFunc) *ImmunocompromisedGroupsGroupBy {
	iggb.fns = append(iggb.fns, fns...)
	return iggb
}

// Scan applies the group-by query and scans the result into the given value.
func (iggb *ImmunocompromisedGroupsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := iggb.path(ctx)
	if err != nil {
		return err
	}
	iggb.sql = query
	return iggb.sqlScan(ctx, v)
}

func (iggb *ImmunocompromisedGroupsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range iggb.fields {
		if !immunocompromisedgroups.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := iggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (iggb *ImmunocompromisedGroupsGroupBy) sqlQuery() *sql.Selector {
	selector := iggb.sql.Select()
	aggregation := make([]string, 0, len(iggb.fns))
	for _, fn := range iggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(iggb.fields)+len(iggb.fns))
		for _, f := range iggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(iggb.fields...)...)
}

// ImmunocompromisedGroupsSelect is the builder for selecting fields of ImmunocompromisedGroups entities.
type ImmunocompromisedGroupsSelect struct {
	*ImmunocompromisedGroupsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (igs *ImmunocompromisedGroupsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := igs.prepareQuery(ctx); err != nil {
		return err
	}
	igs.sql = igs.ImmunocompromisedGroupsQuery.sqlQuery(ctx)
	return igs.sqlScan(ctx, v)
}

func (igs *ImmunocompromisedGroupsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := igs.sql.Query()
	if err := igs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
