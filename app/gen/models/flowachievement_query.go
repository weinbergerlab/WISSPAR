// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowachievement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/flowmilestone"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// FlowAchievementQuery is the builder for querying FlowAchievement entities.
type FlowAchievementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowAchievement
	// eager-loading edges.
	withParent *FlowMilestoneQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowAchievementQuery builder.
func (faq *FlowAchievementQuery) Where(ps ...predicate.FlowAchievement) *FlowAchievementQuery {
	faq.predicates = append(faq.predicates, ps...)
	return faq
}

// Limit adds a limit step to the query.
func (faq *FlowAchievementQuery) Limit(limit int) *FlowAchievementQuery {
	faq.limit = &limit
	return faq
}

// Offset adds an offset step to the query.
func (faq *FlowAchievementQuery) Offset(offset int) *FlowAchievementQuery {
	faq.offset = &offset
	return faq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (faq *FlowAchievementQuery) Unique(unique bool) *FlowAchievementQuery {
	faq.unique = &unique
	return faq
}

// Order adds an order step to the query.
func (faq *FlowAchievementQuery) Order(o ...OrderFunc) *FlowAchievementQuery {
	faq.order = append(faq.order, o...)
	return faq
}

// QueryParent chains the current query on the "parent" edge.
func (faq *FlowAchievementQuery) QueryParent() *FlowMilestoneQuery {
	query := &FlowMilestoneQuery{config: faq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := faq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := faq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowachievement.Table, flowachievement.FieldID, selector),
			sqlgraph.To(flowmilestone.Table, flowmilestone.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowachievement.ParentTable, flowachievement.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(faq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowAchievement entity from the query.
// Returns a *NotFoundError when no FlowAchievement was found.
func (faq *FlowAchievementQuery) First(ctx context.Context) (*FlowAchievement, error) {
	nodes, err := faq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowachievement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (faq *FlowAchievementQuery) FirstX(ctx context.Context) *FlowAchievement {
	node, err := faq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowAchievement ID from the query.
// Returns a *NotFoundError when no FlowAchievement ID was found.
func (faq *FlowAchievementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = faq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowachievement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (faq *FlowAchievementQuery) FirstIDX(ctx context.Context) int {
	id, err := faq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowAchievement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowAchievement entity is found.
// Returns a *NotFoundError when no FlowAchievement entities are found.
func (faq *FlowAchievementQuery) Only(ctx context.Context) (*FlowAchievement, error) {
	nodes, err := faq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowachievement.Label}
	default:
		return nil, &NotSingularError{flowachievement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (faq *FlowAchievementQuery) OnlyX(ctx context.Context) *FlowAchievement {
	node, err := faq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowAchievement ID in the query.
// Returns a *NotSingularError when more than one FlowAchievement ID is found.
// Returns a *NotFoundError when no entities are found.
func (faq *FlowAchievementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = faq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowachievement.Label}
	default:
		err = &NotSingularError{flowachievement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (faq *FlowAchievementQuery) OnlyIDX(ctx context.Context) int {
	id, err := faq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowAchievements.
func (faq *FlowAchievementQuery) All(ctx context.Context) ([]*FlowAchievement, error) {
	if err := faq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return faq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (faq *FlowAchievementQuery) AllX(ctx context.Context) []*FlowAchievement {
	nodes, err := faq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowAchievement IDs.
func (faq *FlowAchievementQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := faq.Select(flowachievement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (faq *FlowAchievementQuery) IDsX(ctx context.Context) []int {
	ids, err := faq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (faq *FlowAchievementQuery) Count(ctx context.Context) (int, error) {
	if err := faq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return faq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (faq *FlowAchievementQuery) CountX(ctx context.Context) int {
	count, err := faq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (faq *FlowAchievementQuery) Exist(ctx context.Context) (bool, error) {
	if err := faq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return faq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (faq *FlowAchievementQuery) ExistX(ctx context.Context) bool {
	exist, err := faq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowAchievementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (faq *FlowAchievementQuery) Clone() *FlowAchievementQuery {
	if faq == nil {
		return nil
	}
	return &FlowAchievementQuery{
		config:     faq.config,
		limit:      faq.limit,
		offset:     faq.offset,
		order:      append([]OrderFunc{}, faq.order...),
		predicates: append([]predicate.FlowAchievement{}, faq.predicates...),
		withParent: faq.withParent.Clone(),
		// clone intermediate query.
		sql:    faq.sql.Clone(),
		path:   faq.path,
		unique: faq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (faq *FlowAchievementQuery) WithParent(opts ...func(*FlowMilestoneQuery)) *FlowAchievementQuery {
	query := &FlowMilestoneQuery{config: faq.config}
	for _, opt := range opts {
		opt(query)
	}
	faq.withParent = query
	return faq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlowAchievementGroupID string `json:"flow_achievement_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowAchievement.Query().
//		GroupBy(flowachievement.FieldFlowAchievementGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (faq *FlowAchievementQuery) GroupBy(field string, fields ...string) *FlowAchievementGroupBy {
	grbuild := &FlowAchievementGroupBy{config: faq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := faq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return faq.sqlQuery(ctx), nil
	}
	grbuild.label = flowachievement.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlowAchievementGroupID string `json:"flow_achievement_group_id,omitempty"`
//	}
//
//	client.FlowAchievement.Query().
//		Select(flowachievement.FieldFlowAchievementGroupID).
//		Scan(ctx, &v)
//
func (faq *FlowAchievementQuery) Select(fields ...string) *FlowAchievementSelect {
	faq.fields = append(faq.fields, fields...)
	selbuild := &FlowAchievementSelect{FlowAchievementQuery: faq}
	selbuild.label = flowachievement.Label
	selbuild.flds, selbuild.scan = &faq.fields, selbuild.Scan
	return selbuild
}

func (faq *FlowAchievementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range faq.fields {
		if !flowachievement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if faq.path != nil {
		prev, err := faq.path(ctx)
		if err != nil {
			return err
		}
		faq.sql = prev
	}
	return nil
}

func (faq *FlowAchievementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowAchievement, error) {
	var (
		nodes       = []*FlowAchievement{}
		withFKs     = faq.withFKs
		_spec       = faq.querySpec()
		loadedTypes = [1]bool{
			faq.withParent != nil,
		}
	)
	if faq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, flowachievement.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowAchievement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowAchievement{config: faq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, faq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := faq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*FlowAchievement)
		for i := range nodes {
			if nodes[i].flow_milestone_flow_achievement_list == nil {
				continue
			}
			fk := *nodes[i].flow_milestone_flow_achievement_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flowmilestone.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_milestone_flow_achievement_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (faq *FlowAchievementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := faq.querySpec()
	_spec.Node.Columns = faq.fields
	if len(faq.fields) > 0 {
		_spec.Unique = faq.unique != nil && *faq.unique
	}
	return sqlgraph.CountNodes(ctx, faq.driver, _spec)
}

func (faq *FlowAchievementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := faq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (faq *FlowAchievementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowachievement.Table,
			Columns: flowachievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: flowachievement.FieldID,
			},
		},
		From:   faq.sql,
		Unique: true,
	}
	if unique := faq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := faq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowachievement.FieldID)
		for i := range fields {
			if fields[i] != flowachievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := faq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := faq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := faq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := faq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (faq *FlowAchievementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(faq.driver.Dialect())
	t1 := builder.Table(flowachievement.Table)
	columns := faq.fields
	if len(columns) == 0 {
		columns = flowachievement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if faq.sql != nil {
		selector = faq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if faq.unique != nil && *faq.unique {
		selector.Distinct()
	}
	for _, p := range faq.predicates {
		p(selector)
	}
	for _, p := range faq.order {
		p(selector)
	}
	if offset := faq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := faq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowAchievementGroupBy is the group-by builder for FlowAchievement entities.
type FlowAchievementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fagb *FlowAchievementGroupBy) Aggregate(fns ...AggregateFunc) *FlowAchievementGroupBy {
	fagb.fns = append(fagb.fns, fns...)
	return fagb
}

// Scan applies the group-by query and scans the result into the given value.
func (fagb *FlowAchievementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fagb.path(ctx)
	if err != nil {
		return err
	}
	fagb.sql = query
	return fagb.sqlScan(ctx, v)
}

func (fagb *FlowAchievementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fagb.fields {
		if !flowachievement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fagb *FlowAchievementGroupBy) sqlQuery() *sql.Selector {
	selector := fagb.sql.Select()
	aggregation := make([]string, 0, len(fagb.fns))
	for _, fn := range fagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fagb.fields)+len(fagb.fns))
		for _, f := range fagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fagb.fields...)...)
}

// FlowAchievementSelect is the builder for selecting fields of FlowAchievement entities.
type FlowAchievementSelect struct {
	*FlowAchievementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fas *FlowAchievementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fas.prepareQuery(ctx); err != nil {
		return err
	}
	fas.sql = fas.FlowAchievementQuery.sqlQuery(ctx)
	return fas.sqlScan(ctx, v)
}

func (fas *FlowAchievementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fas.sql.Query()
	if err := fas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
