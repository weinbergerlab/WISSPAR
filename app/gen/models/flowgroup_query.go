// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowGroupQuery is the builder for querying FlowGroup entities.
type FlowGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowGroup
	// eager-loading edges.
	withParent *ParticipantFlowModuleQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowGroupQuery builder.
func (fgq *FlowGroupQuery) Where(ps ...predicate.FlowGroup) *FlowGroupQuery {
	fgq.predicates = append(fgq.predicates, ps...)
	return fgq
}

// Limit adds a limit step to the query.
func (fgq *FlowGroupQuery) Limit(limit int) *FlowGroupQuery {
	fgq.limit = &limit
	return fgq
}

// Offset adds an offset step to the query.
func (fgq *FlowGroupQuery) Offset(offset int) *FlowGroupQuery {
	fgq.offset = &offset
	return fgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fgq *FlowGroupQuery) Unique(unique bool) *FlowGroupQuery {
	fgq.unique = &unique
	return fgq
}

// Order adds an order step to the query.
func (fgq *FlowGroupQuery) Order(o ...OrderFunc) *FlowGroupQuery {
	fgq.order = append(fgq.order, o...)
	return fgq
}

// QueryParent chains the current query on the "parent" edge.
func (fgq *FlowGroupQuery) QueryParent() *ParticipantFlowModuleQuery {
	query := &ParticipantFlowModuleQuery{config: fgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowgroup.Table, flowgroup.FieldID, selector),
			sqlgraph.To(participantflowmodule.Table, participantflowmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowgroup.ParentTable, flowgroup.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(fgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowGroup entity from the query.
// Returns a *NotFoundError when no FlowGroup was found.
func (fgq *FlowGroupQuery) First(ctx context.Context) (*FlowGroup, error) {
	nodes, err := fgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fgq *FlowGroupQuery) FirstX(ctx context.Context) *FlowGroup {
	node, err := fgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowGroup ID from the query.
// Returns a *NotFoundError when no FlowGroup ID was found.
func (fgq *FlowGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fgq *FlowGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := fgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowGroup entity is found.
// Returns a *NotFoundError when no FlowGroup entities are found.
func (fgq *FlowGroupQuery) Only(ctx context.Context) (*FlowGroup, error) {
	nodes, err := fgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowgroup.Label}
	default:
		return nil, &NotSingularError{flowgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fgq *FlowGroupQuery) OnlyX(ctx context.Context) *FlowGroup {
	node, err := fgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowGroup ID in the query.
// Returns a *NotSingularError when more than one FlowGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (fgq *FlowGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowgroup.Label}
	default:
		err = &NotSingularError{flowgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fgq *FlowGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := fgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowGroups.
func (fgq *FlowGroupQuery) All(ctx context.Context) ([]*FlowGroup, error) {
	if err := fgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fgq *FlowGroupQuery) AllX(ctx context.Context) []*FlowGroup {
	nodes, err := fgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowGroup IDs.
func (fgq *FlowGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fgq.Select(flowgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fgq *FlowGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := fgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fgq *FlowGroupQuery) Count(ctx context.Context) (int, error) {
	if err := fgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fgq *FlowGroupQuery) CountX(ctx context.Context) int {
	count, err := fgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fgq *FlowGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := fgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fgq *FlowGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := fgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fgq *FlowGroupQuery) Clone() *FlowGroupQuery {
	if fgq == nil {
		return nil
	}
	return &FlowGroupQuery{
		config:     fgq.config,
		limit:      fgq.limit,
		offset:     fgq.offset,
		order:      append([]OrderFunc{}, fgq.order...),
		predicates: append([]predicate.FlowGroup{}, fgq.predicates...),
		withParent: fgq.withParent.Clone(),
		// clone intermediate query.
		sql:    fgq.sql.Clone(),
		path:   fgq.path,
		unique: fgq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (fgq *FlowGroupQuery) WithParent(opts ...func(*ParticipantFlowModuleQuery)) *FlowGroupQuery {
	query := &ParticipantFlowModuleQuery{config: fgq.config}
	for _, opt := range opts {
		opt(query)
	}
	fgq.withParent = query
	return fgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowGroupID string `json:"flow_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowGroup.Query().
//		GroupBy(flowgroup.FieldFlowGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (fgq *FlowGroupQuery) GroupBy(field string, fields ...string) *FlowGroupGroupBy {
	grbuild := &FlowGroupGroupBy{config: fgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fgq.sqlQuery(ctx), nil
	}
	grbuild.label = flowgroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowGroupID string `json:"flow_group_id,omitempty"`
//	}
//
//	client.FlowGroup.Query().
//		Select(flowgroup.FieldFlowGroupID).
//		Scan(ctx, &v)
//
func (fgq *FlowGroupQuery) Select(fields ...string) *FlowGroupSelect {
	fgq.fields = append(fgq.fields, fields...)
	selbuild := &FlowGroupSelect{FlowGroupQuery: fgq}
	selbuild.label = flowgroup.Label
	selbuild.flds, selbuild.scan = &fgq.fields, selbuild.Scan
	return selbuild
}

func (fgq *FlowGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fgq.fields {
		if !flowgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if fgq.path != nil {
		prev, err := fgq.path(ctx)
		if err != nil {
			return err
		}
		fgq.sql = prev
	}
	return nil
}

func (fgq *FlowGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowGroup, error) {
	var (
		nodes       = []*FlowGroup{}
		withFKs     = fgq.withFKs
		_spec       = fgq.querySpec()
		loadedTypes = [1]bool{
			fgq.withParent != nil,
		}
	)
	if fgq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flowgroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowGroup{config: fgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fgq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*FlowGroup)
		for i := range nodes {
			if nodes[i].participant_flow_module_flow_group_list == nil {
				continue
			}
			fk := *nodes[i].participant_flow_module_flow_group_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(participantflowmodule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "participant_flow_module_flow_group_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (fgq *FlowGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fgq.querySpec()
	_spec.Node.Columns = fgq.fields
	if len(fgq.fields) > 0 {
		_spec.Unique = fgq.unique != nil && *fgq.unique
	}
	return sqlgraph.CountNodes(ctx, fgq.driver, _spec)
}

func (fgq *FlowGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (fgq *FlowGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowgroup.Table,
			Columns: flowgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowgroup.FieldID,
			},
		},
		From:   fgq.sql,
		Unique: true,
	}
	if unique := fgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowgroup.FieldID)
		for i := range fields {
			if fields[i] != flowgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fgq *FlowGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fgq.driver.Dialect())
	t1 := builder.Table(flowgroup.Table)
	columns := fgq.fields
	if len(columns) == 0 {
		columns = flowgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fgq.sql != nil {
		selector = fgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fgq.unique != nil && *fgq.unique {
		selector.Distinct()
	}
	for _, p := range fgq.predicates {
		p(selector)
	}
	for _, p := range fgq.order {
		p(selector)
	}
	if offset := fgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowGroupGroupBy is the group-by builder for FlowGroup entities.
type FlowGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fggb *FlowGroupGroupBy) Aggregate(fns ...AggregateFunc) *FlowGroupGroupBy {
	fggb.fns = append(fggb.fns, fns...)
	return fggb
}

// Scan applies the group-by query and scans the result into the given value.
func (fggb *FlowGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fggb.path(ctx)
	if err != nil {
		return err
	}
	fggb.sql = query
	return fggb.sqlScan(ctx, v)
}

func (fggb *FlowGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fggb.fields {
		if !flowgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fggb *FlowGroupGroupBy) sqlQuery() *sql.Selector {
	selector := fggb.sql.Select()
	aggregation := make([]string, 0, len(fggb.fns))
	for _, fn := range fggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fggb.fields)+len(fggb.fns))
		for _, f := range fggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fggb.fields...)...)
}

// FlowGroupSelect is the builder for selecting fields of FlowGroup entities.
type FlowGroupSelect struct {
	*FlowGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fgs *FlowGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fgs.prepareQuery(ctx); err != nil {
		return err
	}
	fgs.sql = fgs.FlowGroupQuery.sqlQuery(ctx)
	return fgs.sqlScan(ctx, v)
}

func (fgs *FlowGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fgs.sql.Query()
	if err := fgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
