// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowReasonQuery is the builder for querying FlowReason entities.
type FlowReasonQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowReason
	// eager-loading edges.
	withParent *FlowDropWithdrawQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowReasonQuery builder.
func (frq *FlowReasonQuery) Where(ps ...predicate.FlowReason) *FlowReasonQuery {
	frq.predicates = append(frq.predicates, ps...)
	return frq
}

// Limit adds a limit step to the query.
func (frq *FlowReasonQuery) Limit(limit int) *FlowReasonQuery {
	frq.limit = &limit
	return frq
}

// Offset adds an offset step to the query.
func (frq *FlowReasonQuery) Offset(offset int) *FlowReasonQuery {
	frq.offset = &offset
	return frq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (frq *FlowReasonQuery) Unique(unique bool) *FlowReasonQuery {
	frq.unique = &unique
	return frq
}

// Order adds an order step to the query.
func (frq *FlowReasonQuery) Order(o ...OrderFunc) *FlowReasonQuery {
	frq.order = append(frq.order, o...)
	return frq
}

// QueryParent chains the current query on the "parent" edge.
func (frq *FlowReasonQuery) QueryParent() *FlowDropWithdrawQuery {
	query := &FlowDropWithdrawQuery{config: frq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := frq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := frq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowreason.Table, flowreason.FieldID, selector),
			sqlgraph.To(flowdropwithdraw.Table, flowdropwithdraw.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowreason.ParentTable, flowreason.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(frq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowReason entity from the query.
// Returns a *NotFoundError when no FlowReason was found.
func (frq *FlowReasonQuery) First(ctx context.Context) (*FlowReason, error) {
	nodes, err := frq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowreason.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (frq *FlowReasonQuery) FirstX(ctx context.Context) *FlowReason {
	node, err := frq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowReason ID from the query.
// Returns a *NotFoundError when no FlowReason ID was found.
func (frq *FlowReasonQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = frq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowreason.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (frq *FlowReasonQuery) FirstIDX(ctx context.Context) int {
	id, err := frq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowReason entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowReason entity is found.
// Returns a *NotFoundError when no FlowReason entities are found.
func (frq *FlowReasonQuery) Only(ctx context.Context) (*FlowReason, error) {
	nodes, err := frq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowreason.Label}
	default:
		return nil, &NotSingularError{flowreason.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (frq *FlowReasonQuery) OnlyX(ctx context.Context) *FlowReason {
	node, err := frq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowReason ID in the query.
// Returns a *NotSingularError when more than one FlowReason ID is found.
// Returns a *NotFoundError when no entities are found.
func (frq *FlowReasonQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = frq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowreason.Label}
	default:
		err = &NotSingularError{flowreason.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (frq *FlowReasonQuery) OnlyIDX(ctx context.Context) int {
	id, err := frq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowReasons.
func (frq *FlowReasonQuery) All(ctx context.Context) ([]*FlowReason, error) {
	if err := frq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return frq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (frq *FlowReasonQuery) AllX(ctx context.Context) []*FlowReason {
	nodes, err := frq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowReason IDs.
func (frq *FlowReasonQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := frq.Select(flowreason.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (frq *FlowReasonQuery) IDsX(ctx context.Context) []int {
	ids, err := frq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (frq *FlowReasonQuery) Count(ctx context.Context) (int, error) {
	if err := frq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return frq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (frq *FlowReasonQuery) CountX(ctx context.Context) int {
	count, err := frq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (frq *FlowReasonQuery) Exist(ctx context.Context) (bool, error) {
	if err := frq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return frq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (frq *FlowReasonQuery) ExistX(ctx context.Context) bool {
	exist, err := frq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowReasonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (frq *FlowReasonQuery) Clone() *FlowReasonQuery {
	if frq == nil {
		return nil
	}
	return &FlowReasonQuery{
		config:     frq.config,
		limit:      frq.limit,
		offset:     frq.offset,
		order:      append([]OrderFunc{}, frq.order...),
		predicates: append([]predicate.FlowReason{}, frq.predicates...),
		withParent: frq.withParent.Clone(),
		// clone intermediate query.
		sql:    frq.sql.Clone(),
		path:   frq.path,
		unique: frq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (frq *FlowReasonQuery) WithParent(opts ...func(*FlowDropWithdrawQuery)) *FlowReasonQuery {
	query := &FlowDropWithdrawQuery{config: frq.config}
	for _, opt := range opts {
		opt(query)
	}
	frq.withParent = query
	return frq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowReasonGroupID string `json:"flow_reason_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowReason.Query().
//		GroupBy(flowreason.FieldFlowReasonGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (frq *FlowReasonQuery) GroupBy(field string, fields ...string) *FlowReasonGroupBy {
	grbuild := &FlowReasonGroupBy{config: frq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := frq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return frq.sqlQuery(ctx), nil
	}
	grbuild.label = flowreason.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowReasonGroupID string `json:"flow_reason_group_id,omitempty"`
//	}
//
//	client.FlowReason.Query().
//		Select(flowreason.FieldFlowReasonGroupID).
//		Scan(ctx, &v)
//
func (frq *FlowReasonQuery) Select(fields ...string) *FlowReasonSelect {
	frq.fields = append(frq.fields, fields...)
	selbuild := &FlowReasonSelect{FlowReasonQuery: frq}
	selbuild.label = flowreason.Label
	selbuild.flds, selbuild.scan = &frq.fields, selbuild.Scan
	return selbuild
}

func (frq *FlowReasonQuery) prepareQuery(ctx context.Context) error {
	for _, f := range frq.fields {
		if !flowreason.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if frq.path != nil {
		prev, err := frq.path(ctx)
		if err != nil {
			return err
		}
		frq.sql = prev
	}
	return nil
}

func (frq *FlowReasonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowReason, error) {
	var (
		nodes       = []*FlowReason{}
		withFKs     = frq.withFKs
		_spec       = frq.querySpec()
		loadedTypes = [1]bool{
			frq.withParent != nil,
		}
	)
	if frq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flowreason.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowReason).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowReason{config: frq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, frq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := frq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*FlowReason)
		for i := range nodes {
			if nodes[i].flow_drop_withdraw_flow_reason_list == nil {
				continue
			}
			fk := *nodes[i].flow_drop_withdraw_flow_reason_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flowdropwithdraw.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_drop_withdraw_flow_reason_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (frq *FlowReasonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := frq.querySpec()
	_spec.Node.Columns = frq.fields
	if len(frq.fields) > 0 {
		_spec.Unique = frq.unique != nil && *frq.unique
	}
	return sqlgraph.CountNodes(ctx, frq.driver, _spec)
}

func (frq *FlowReasonQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := frq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (frq *FlowReasonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowreason.Table,
			Columns: flowreason.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowreason.FieldID,
			},
		},
		From:   frq.sql,
		Unique: true,
	}
	if unique := frq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := frq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowreason.FieldID)
		for i := range fields {
			if fields[i] != flowreason.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := frq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := frq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := frq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := frq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (frq *FlowReasonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(frq.driver.Dialect())
	t1 := builder.Table(flowreason.Table)
	columns := frq.fields
	if len(columns) == 0 {
		columns = flowreason.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if frq.sql != nil {
		selector = frq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if frq.unique != nil && *frq.unique {
		selector.Distinct()
	}
	for _, p := range frq.predicates {
		p(selector)
	}
	for _, p := range frq.order {
		p(selector)
	}
	if offset := frq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := frq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowReasonGroupBy is the group-by builder for FlowReason entities.
type FlowReasonGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (frgb *FlowReasonGroupBy) Aggregate(fns ...AggregateFunc) *FlowReasonGroupBy {
	frgb.fns = append(frgb.fns, fns...)
	return frgb
}

// Scan applies the group-by query and scans the result into the given value.
func (frgb *FlowReasonGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := frgb.path(ctx)
	if err != nil {
		return err
	}
	frgb.sql = query
	return frgb.sqlScan(ctx, v)
}

func (frgb *FlowReasonGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range frgb.fields {
		if !flowreason.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := frgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := frgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (frgb *FlowReasonGroupBy) sqlQuery() *sql.Selector {
	selector := frgb.sql.Select()
	aggregation := make([]string, 0, len(frgb.fns))
	for _, fn := range frgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(frgb.fields)+len(frgb.fns))
		for _, f := range frgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(frgb.fields...)...)
}

// FlowReasonSelect is the builder for selecting fields of FlowReason entities.
type FlowReasonSelect struct {
	*FlowReasonQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (frs *FlowReasonSelect) Scan(ctx context.Context, v interface{}) error {
	if err := frs.prepareQuery(ctx); err != nil {
		return err
	}
	frs.sql = frs.FlowReasonQuery.sqlQuery(ctx)
	return frs.sqlScan(ctx, v)
}

func (frs *FlowReasonSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := frs.sql.Query()
	if err := frs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
