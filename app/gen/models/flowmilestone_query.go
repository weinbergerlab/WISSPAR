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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowperiod"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowMilestoneQuery is the builder for querying FlowMilestone entities.
type FlowMilestoneQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowMilestone
	// eager-loading edges.
	withParent              *FlowPeriodQuery
	withFlowAchievementList *FlowAchievementQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowMilestoneQuery builder.
func (fmq *FlowMilestoneQuery) Where(ps ...predicate.FlowMilestone) *FlowMilestoneQuery {
	fmq.predicates = append(fmq.predicates, ps...)
	return fmq
}

// Limit adds a limit step to the query.
func (fmq *FlowMilestoneQuery) Limit(limit int) *FlowMilestoneQuery {
	fmq.limit = &limit
	return fmq
}

// Offset adds an offset step to the query.
func (fmq *FlowMilestoneQuery) Offset(offset int) *FlowMilestoneQuery {
	fmq.offset = &offset
	return fmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fmq *FlowMilestoneQuery) Unique(unique bool) *FlowMilestoneQuery {
	fmq.unique = &unique
	return fmq
}

// Order adds an order step to the query.
func (fmq *FlowMilestoneQuery) Order(o ...OrderFunc) *FlowMilestoneQuery {
	fmq.order = append(fmq.order, o...)
	return fmq
}

// QueryParent chains the current query on the "parent" edge.
func (fmq *FlowMilestoneQuery) QueryParent() *FlowPeriodQuery {
	query := &FlowPeriodQuery{config: fmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowmilestone.Table, flowmilestone.FieldID, selector),
			sqlgraph.To(flowperiod.Table, flowperiod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowmilestone.ParentTable, flowmilestone.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(fmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowAchievementList chains the current query on the "flow_achievement_list" edge.
func (fmq *FlowMilestoneQuery) QueryFlowAchievementList() *FlowAchievementQuery {
	query := &FlowAchievementQuery{config: fmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowmilestone.Table, flowmilestone.FieldID, selector),
			sqlgraph.To(flowachievement.Table, flowachievement.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowmilestone.FlowAchievementListTable, flowmilestone.FlowAchievementListColumn),
		)
		fromU = sqlgraph.SetNeighbors(fmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowMilestone entity from the query.
// Returns a *NotFoundError when no FlowMilestone was found.
func (fmq *FlowMilestoneQuery) First(ctx context.Context) (*FlowMilestone, error) {
	nodes, err := fmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowmilestone.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) FirstX(ctx context.Context) *FlowMilestone {
	node, err := fmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowMilestone ID from the query.
// Returns a *NotFoundError when no FlowMilestone ID was found.
func (fmq *FlowMilestoneQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowmilestone.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) FirstIDX(ctx context.Context) int {
	id, err := fmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowMilestone entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowMilestone entity is found.
// Returns a *NotFoundError when no FlowMilestone entities are found.
func (fmq *FlowMilestoneQuery) Only(ctx context.Context) (*FlowMilestone, error) {
	nodes, err := fmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowmilestone.Label}
	default:
		return nil, &NotSingularError{flowmilestone.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) OnlyX(ctx context.Context) *FlowMilestone {
	node, err := fmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowMilestone ID in the query.
// Returns a *NotSingularError when more than one FlowMilestone ID is found.
// Returns a *NotFoundError when no entities are found.
func (fmq *FlowMilestoneQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowmilestone.Label}
	default:
		err = &NotSingularError{flowmilestone.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) OnlyIDX(ctx context.Context) int {
	id, err := fmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowMilestones.
func (fmq *FlowMilestoneQuery) All(ctx context.Context) ([]*FlowMilestone, error) {
	if err := fmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) AllX(ctx context.Context) []*FlowMilestone {
	nodes, err := fmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowMilestone IDs.
func (fmq *FlowMilestoneQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fmq.Select(flowmilestone.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) IDsX(ctx context.Context) []int {
	ids, err := fmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fmq *FlowMilestoneQuery) Count(ctx context.Context) (int, error) {
	if err := fmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) CountX(ctx context.Context) int {
	count, err := fmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fmq *FlowMilestoneQuery) Exist(ctx context.Context) (bool, error) {
	if err := fmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fmq *FlowMilestoneQuery) ExistX(ctx context.Context) bool {
	exist, err := fmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowMilestoneQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fmq *FlowMilestoneQuery) Clone() *FlowMilestoneQuery {
	if fmq == nil {
		return nil
	}
	return &FlowMilestoneQuery{
		config:                  fmq.config,
		limit:                   fmq.limit,
		offset:                  fmq.offset,
		order:                   append([]OrderFunc{}, fmq.order...),
		predicates:              append([]predicate.FlowMilestone{}, fmq.predicates...),
		withParent:              fmq.withParent.Clone(),
		withFlowAchievementList: fmq.withFlowAchievementList.Clone(),
		// clone intermediate query.
		sql:    fmq.sql.Clone(),
		path:   fmq.path,
		unique: fmq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (fmq *FlowMilestoneQuery) WithParent(opts ...func(*FlowPeriodQuery)) *FlowMilestoneQuery {
	query := &FlowPeriodQuery{config: fmq.config}
	for _, opt := range opts {
		opt(query)
	}
	fmq.withParent = query
	return fmq
}

// WithFlowAchievementList tells the query-builder to eager-load the nodes that are connected to
// the "flow_achievement_list" edge. The optional arguments are used to configure the query builder of the edge.
func (fmq *FlowMilestoneQuery) WithFlowAchievementList(opts ...func(*FlowAchievementQuery)) *FlowMilestoneQuery {
	query := &FlowAchievementQuery{config: fmq.config}
	for _, opt := range opts {
		opt(query)
	}
	fmq.withFlowAchievementList = query
	return fmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowMilestoneType string `json:"flow_milestone_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowMilestone.Query().
//		GroupBy(flowmilestone.FieldFlowMilestoneType).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (fmq *FlowMilestoneQuery) GroupBy(field string, fields ...string) *FlowMilestoneGroupBy {
	grbuild := &FlowMilestoneGroupBy{config: fmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fmq.sqlQuery(ctx), nil
	}
	grbuild.label = flowmilestone.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowMilestoneType string `json:"flow_milestone_type,omitempty"`
//	}
//
//	client.FlowMilestone.Query().
//		Select(flowmilestone.FieldFlowMilestoneType).
//		Scan(ctx, &v)
//
func (fmq *FlowMilestoneQuery) Select(fields ...string) *FlowMilestoneSelect {
	fmq.fields = append(fmq.fields, fields...)
	selbuild := &FlowMilestoneSelect{FlowMilestoneQuery: fmq}
	selbuild.label = flowmilestone.Label
	selbuild.flds, selbuild.scan = &fmq.fields, selbuild.Scan
	return selbuild
}

func (fmq *FlowMilestoneQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fmq.fields {
		if !flowmilestone.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if fmq.path != nil {
		prev, err := fmq.path(ctx)
		if err != nil {
			return err
		}
		fmq.sql = prev
	}
	return nil
}

func (fmq *FlowMilestoneQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowMilestone, error) {
	var (
		nodes       = []*FlowMilestone{}
		withFKs     = fmq.withFKs
		_spec       = fmq.querySpec()
		loadedTypes = [2]bool{
			fmq.withParent != nil,
			fmq.withFlowAchievementList != nil,
		}
	)
	if fmq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flowmilestone.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowMilestone).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowMilestone{config: fmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fmq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*FlowMilestone)
		for i := range nodes {
			if nodes[i].flow_period_flow_milestone_list == nil {
				continue
			}
			fk := *nodes[i].flow_period_flow_milestone_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "flow_period_flow_milestone_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := fmq.withFlowAchievementList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*FlowMilestone)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowAchievementList = []*FlowAchievement{}
		}
		query.withFKs = true
		query.Where(predicate.FlowAchievement(func(s *sql.Selector) {
			s.Where(sql.InValues(flowmilestone.FlowAchievementListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.flow_milestone_flow_achievement_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "flow_milestone_flow_achievement_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_milestone_flow_achievement_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FlowAchievementList = append(node.Edges.FlowAchievementList, n)
		}
	}

	return nodes, nil
}

func (fmq *FlowMilestoneQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fmq.querySpec()
	_spec.Node.Columns = fmq.fields
	if len(fmq.fields) > 0 {
		_spec.Unique = fmq.unique != nil && *fmq.unique
	}
	return sqlgraph.CountNodes(ctx, fmq.driver, _spec)
}

func (fmq *FlowMilestoneQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (fmq *FlowMilestoneQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowmilestone.Table,
			Columns: flowmilestone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowmilestone.FieldID,
			},
		},
		From:   fmq.sql,
		Unique: true,
	}
	if unique := fmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowmilestone.FieldID)
		for i := range fields {
			if fields[i] != flowmilestone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fmq *FlowMilestoneQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fmq.driver.Dialect())
	t1 := builder.Table(flowmilestone.Table)
	columns := fmq.fields
	if len(columns) == 0 {
		columns = flowmilestone.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fmq.sql != nil {
		selector = fmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fmq.unique != nil && *fmq.unique {
		selector.Distinct()
	}
	for _, p := range fmq.predicates {
		p(selector)
	}
	for _, p := range fmq.order {
		p(selector)
	}
	if offset := fmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowMilestoneGroupBy is the group-by builder for FlowMilestone entities.
type FlowMilestoneGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fmgb *FlowMilestoneGroupBy) Aggregate(fns ...AggregateFunc) *FlowMilestoneGroupBy {
	fmgb.fns = append(fmgb.fns, fns...)
	return fmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fmgb *FlowMilestoneGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fmgb.path(ctx)
	if err != nil {
		return err
	}
	fmgb.sql = query
	return fmgb.sqlScan(ctx, v)
}

func (fmgb *FlowMilestoneGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fmgb.fields {
		if !flowmilestone.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fmgb *FlowMilestoneGroupBy) sqlQuery() *sql.Selector {
	selector := fmgb.sql.Select()
	aggregation := make([]string, 0, len(fmgb.fns))
	for _, fn := range fmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fmgb.fields)+len(fmgb.fns))
		for _, f := range fmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fmgb.fields...)...)
}

// FlowMilestoneSelect is the builder for selecting fields of FlowMilestone entities.
type FlowMilestoneSelect struct {
	*FlowMilestoneQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fms *FlowMilestoneSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fms.prepareQuery(ctx); err != nil {
		return err
	}
	fms.sql = fms.FlowMilestoneQuery.sqlQuery(ctx)
	return fms.sqlScan(ctx, v)
}

func (fms *FlowMilestoneSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fms.sql.Query()
	if err := fms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
