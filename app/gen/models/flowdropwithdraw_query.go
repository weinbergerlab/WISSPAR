// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowdropwithdraw"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowreason"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowDropWithdrawQuery is the builder for querying FlowDropWithdraw entities.
type FlowDropWithdrawQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowDropWithdraw
	// eager-loading edges.
	withParent         *FlowPeriodQuery
	withFlowReasonList *FlowReasonQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowDropWithdrawQuery builder.
func (fdwq *FlowDropWithdrawQuery) Where(ps ...predicate.FlowDropWithdraw) *FlowDropWithdrawQuery {
	fdwq.predicates = append(fdwq.predicates, ps...)
	return fdwq
}

// Limit adds a limit step to the query.
func (fdwq *FlowDropWithdrawQuery) Limit(limit int) *FlowDropWithdrawQuery {
	fdwq.limit = &limit
	return fdwq
}

// Offset adds an offset step to the query.
func (fdwq *FlowDropWithdrawQuery) Offset(offset int) *FlowDropWithdrawQuery {
	fdwq.offset = &offset
	return fdwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fdwq *FlowDropWithdrawQuery) Unique(unique bool) *FlowDropWithdrawQuery {
	fdwq.unique = &unique
	return fdwq
}

// Order adds an order step to the query.
func (fdwq *FlowDropWithdrawQuery) Order(o ...OrderFunc) *FlowDropWithdrawQuery {
	fdwq.order = append(fdwq.order, o...)
	return fdwq
}

// QueryParent chains the current query on the "parent" edge.
func (fdwq *FlowDropWithdrawQuery) QueryParent() *FlowPeriodQuery {
	query := &FlowPeriodQuery{config: fdwq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fdwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fdwq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowdropwithdraw.Table, flowdropwithdraw.FieldID, selector),
			sqlgraph.To(flowperiod.Table, flowperiod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowdropwithdraw.ParentTable, flowdropwithdraw.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(fdwq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowReasonList chains the current query on the "flow_reason_list" edge.
func (fdwq *FlowDropWithdrawQuery) QueryFlowReasonList() *FlowReasonQuery {
	query := &FlowReasonQuery{config: fdwq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fdwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fdwq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowdropwithdraw.Table, flowdropwithdraw.FieldID, selector),
			sqlgraph.To(flowreason.Table, flowreason.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowdropwithdraw.FlowReasonListTable, flowdropwithdraw.FlowReasonListColumn),
		)
		fromU = sqlgraph.SetNeighbors(fdwq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowDropWithdraw entity from the query.
// Returns a *NotFoundError when no FlowDropWithdraw was found.
func (fdwq *FlowDropWithdrawQuery) First(ctx context.Context) (*FlowDropWithdraw, error) {
	nodes, err := fdwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowdropwithdraw.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) FirstX(ctx context.Context) *FlowDropWithdraw {
	node, err := fdwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowDropWithdraw ID from the query.
// Returns a *NotFoundError when no FlowDropWithdraw ID was found.
func (fdwq *FlowDropWithdrawQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fdwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowdropwithdraw.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) FirstIDX(ctx context.Context) int {
	id, err := fdwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowDropWithdraw entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowDropWithdraw entity is found.
// Returns a *NotFoundError when no FlowDropWithdraw entities are found.
func (fdwq *FlowDropWithdrawQuery) Only(ctx context.Context) (*FlowDropWithdraw, error) {
	nodes, err := fdwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowdropwithdraw.Label}
	default:
		return nil, &NotSingularError{flowdropwithdraw.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) OnlyX(ctx context.Context) *FlowDropWithdraw {
	node, err := fdwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowDropWithdraw ID in the query.
// Returns a *NotSingularError when more than one FlowDropWithdraw ID is found.
// Returns a *NotFoundError when no entities are found.
func (fdwq *FlowDropWithdrawQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fdwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowdropwithdraw.Label}
	default:
		err = &NotSingularError{flowdropwithdraw.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) OnlyIDX(ctx context.Context) int {
	id, err := fdwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowDropWithdraws.
func (fdwq *FlowDropWithdrawQuery) All(ctx context.Context) ([]*FlowDropWithdraw, error) {
	if err := fdwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fdwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) AllX(ctx context.Context) []*FlowDropWithdraw {
	nodes, err := fdwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowDropWithdraw IDs.
func (fdwq *FlowDropWithdrawQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fdwq.Select(flowdropwithdraw.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) IDsX(ctx context.Context) []int {
	ids, err := fdwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fdwq *FlowDropWithdrawQuery) Count(ctx context.Context) (int, error) {
	if err := fdwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fdwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) CountX(ctx context.Context) int {
	count, err := fdwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fdwq *FlowDropWithdrawQuery) Exist(ctx context.Context) (bool, error) {
	if err := fdwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fdwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fdwq *FlowDropWithdrawQuery) ExistX(ctx context.Context) bool {
	exist, err := fdwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowDropWithdrawQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fdwq *FlowDropWithdrawQuery) Clone() *FlowDropWithdrawQuery {
	if fdwq == nil {
		return nil
	}
	return &FlowDropWithdrawQuery{
		config:             fdwq.config,
		limit:              fdwq.limit,
		offset:             fdwq.offset,
		order:              append([]OrderFunc{}, fdwq.order...),
		predicates:         append([]predicate.FlowDropWithdraw{}, fdwq.predicates...),
		withParent:         fdwq.withParent.Clone(),
		withFlowReasonList: fdwq.withFlowReasonList.Clone(),
		// clone intermediate query.
		sql:    fdwq.sql.Clone(),
		path:   fdwq.path,
		unique: fdwq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (fdwq *FlowDropWithdrawQuery) WithParent(opts ...func(*FlowPeriodQuery)) *FlowDropWithdrawQuery {
	query := &FlowPeriodQuery{config: fdwq.config}
	for _, opt := range opts {
		opt(query)
	}
	fdwq.withParent = query
	return fdwq
}

// WithFlowReasonList tells the query-builder to eager-load the nodes that are connected to
// the "flow_reason_list" edge. The optional arguments are used to configure the query builder of the edge.
func (fdwq *FlowDropWithdrawQuery) WithFlowReasonList(opts ...func(*FlowReasonQuery)) *FlowDropWithdrawQuery {
	query := &FlowReasonQuery{config: fdwq.config}
	for _, opt := range opts {
		opt(query)
	}
	fdwq.withFlowReasonList = query
	return fdwq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowDropWithdrawType string `json:"flow_drop_withdraw_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowDropWithdraw.Query().
//		GroupBy(flowdropwithdraw.FieldFlowDropWithdrawType).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (fdwq *FlowDropWithdrawQuery) GroupBy(field string, fields ...string) *FlowDropWithdrawGroupBy {
	grbuild := &FlowDropWithdrawGroupBy{config: fdwq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fdwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fdwq.sqlQuery(ctx), nil
	}
	grbuild.label = flowdropwithdraw.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowDropWithdrawType string `json:"flow_drop_withdraw_type,omitempty"`
//	}
//
//	client.FlowDropWithdraw.Query().
//		Select(flowdropwithdraw.FieldFlowDropWithdrawType).
//		Scan(ctx, &v)
//
func (fdwq *FlowDropWithdrawQuery) Select(fields ...string) *FlowDropWithdrawSelect {
	fdwq.fields = append(fdwq.fields, fields...)
	selbuild := &FlowDropWithdrawSelect{FlowDropWithdrawQuery: fdwq}
	selbuild.label = flowdropwithdraw.Label
	selbuild.flds, selbuild.scan = &fdwq.fields, selbuild.Scan
	return selbuild
}

func (fdwq *FlowDropWithdrawQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fdwq.fields {
		if !flowdropwithdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if fdwq.path != nil {
		prev, err := fdwq.path(ctx)
		if err != nil {
			return err
		}
		fdwq.sql = prev
	}
	return nil
}

func (fdwq *FlowDropWithdrawQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowDropWithdraw, error) {
	var (
		nodes       = []*FlowDropWithdraw{}
		withFKs     = fdwq.withFKs
		_spec       = fdwq.querySpec()
		loadedTypes = [2]bool{
			fdwq.withParent != nil,
			fdwq.withFlowReasonList != nil,
		}
	)
	if fdwq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flowdropwithdraw.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowDropWithdraw).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowDropWithdraw{config: fdwq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fdwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fdwq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*FlowDropWithdraw)
		for i := range nodes {
			if nodes[i].flow_period_flow_drop_withdraw_list == nil {
				continue
			}
			fk := *nodes[i].flow_period_flow_drop_withdraw_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flowperiod.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_period_flow_drop_withdraw_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := fdwq.withFlowReasonList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*FlowDropWithdraw)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowReasonList = []*FlowReason{}
		}
		query.withFKs = true
		query.Where(predicate.FlowReason(func(s *sql.Selector) {
			s.Where(sql.InValues(flowdropwithdraw.FlowReasonListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flow_drop_withdraw_flow_reason_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flow_drop_withdraw_flow_reason_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_drop_withdraw_flow_reason_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlowReasonList = append(node.Edges.FlowReasonList, n)
		}
	}

	return nodes, nil
}

func (fdwq *FlowDropWithdrawQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fdwq.querySpec()
	_spec.Node.Columns = fdwq.fields
	if len(fdwq.fields) > 0 {
		_spec.Unique = fdwq.unique != nil && *fdwq.unique
	}
	return sqlgraph.CountNodes(ctx, fdwq.driver, _spec)
}

func (fdwq *FlowDropWithdrawQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fdwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (fdwq *FlowDropWithdrawQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowdropwithdraw.Table,
			Columns: flowdropwithdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowdropwithdraw.FieldID,
			},
		},
		From:   fdwq.sql,
		Unique: true,
	}
	if unique := fdwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fdwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowdropwithdraw.FieldID)
		for i := range fields {
			if fields[i] != flowdropwithdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fdwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fdwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fdwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fdwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fdwq *FlowDropWithdrawQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fdwq.driver.Dialect())
	t1 := builder.Table(flowdropwithdraw.Table)
	columns := fdwq.fields
	if len(columns) == 0 {
		columns = flowdropwithdraw.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fdwq.sql != nil {
		selector = fdwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fdwq.unique != nil && *fdwq.unique {
		selector.Distinct()
	}
	for _, p := range fdwq.predicates {
		p(selector)
	}
	for _, p := range fdwq.order {
		p(selector)
	}
	if offset := fdwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fdwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowDropWithdrawGroupBy is the group-by builder for FlowDropWithdraw entities.
type FlowDropWithdrawGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fdwgb *FlowDropWithdrawGroupBy) Aggregate(fns ...AggregateFunc) *FlowDropWithdrawGroupBy {
	fdwgb.fns = append(fdwgb.fns, fns...)
	return fdwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fdwgb *FlowDropWithdrawGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fdwgb.path(ctx)
	if err != nil {
		return err
	}
	fdwgb.sql = query
	return fdwgb.sqlScan(ctx, v)
}

func (fdwgb *FlowDropWithdrawGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fdwgb.fields {
		if !flowdropwithdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fdwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fdwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fdwgb *FlowDropWithdrawGroupBy) sqlQuery() *sql.Selector {
	selector := fdwgb.sql.Select()
	aggregation := make([]string, 0, len(fdwgb.fns))
	for _, fn := range fdwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fdwgb.fields)+len(fdwgb.fns))
		for _, f := range fdwgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fdwgb.fields...)...)
}

// FlowDropWithdrawSelect is the builder for selecting fields of FlowDropWithdraw entities.
type FlowDropWithdrawSelect struct {
	*FlowDropWithdrawQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fdws *FlowDropWithdrawSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fdws.prepareQuery(ctx); err != nil {
		return err
	}
	fdws.sql = fdws.FlowDropWithdrawQuery.sqlQuery(ctx)
	return fdws.sqlScan(ctx, v)
}

func (fdws *FlowDropWithdrawSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fdws.sql.Query()
	if err := fdws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
