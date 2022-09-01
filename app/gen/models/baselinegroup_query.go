// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineGroupQuery is the builder for querying BaselineGroup entities.
type BaselineGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineGroup
	// eager-loading edges.
	withParent *BaselineCharacteristicsModuleQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineGroupQuery builder.
func (bgq *BaselineGroupQuery) Where(ps ...predicate.BaselineGroup) *BaselineGroupQuery {
	bgq.predicates = append(bgq.predicates, ps...)
	return bgq
}

// Limit adds a limit step to the query.
func (bgq *BaselineGroupQuery) Limit(limit int) *BaselineGroupQuery {
	bgq.limit = &limit
	return bgq
}

// Offset adds an offset step to the query.
func (bgq *BaselineGroupQuery) Offset(offset int) *BaselineGroupQuery {
	bgq.offset = &offset
	return bgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bgq *BaselineGroupQuery) Unique(unique bool) *BaselineGroupQuery {
	bgq.unique = &unique
	return bgq
}

// Order adds an order step to the query.
func (bgq *BaselineGroupQuery) Order(o ...OrderFunc) *BaselineGroupQuery {
	bgq.order = append(bgq.order, o...)
	return bgq
}

// QueryParent chains the current query on the "parent" edge.
func (bgq *BaselineGroupQuery) QueryParent() *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: bgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinegroup.Table, baselinegroup.FieldID, selector),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinegroup.ParentTable, baselinegroup.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineGroup entity from the query.
// Returns a *NotFoundError when no BaselineGroup was found.
func (bgq *BaselineGroupQuery) First(ctx context.Context) (*BaselineGroup, error) {
	nodes, err := bgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinegroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bgq *BaselineGroupQuery) FirstX(ctx context.Context) *BaselineGroup {
	node, err := bgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineGroup ID from the query.
// Returns a *NotFoundError when no BaselineGroup ID was found.
func (bgq *BaselineGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinegroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bgq *BaselineGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := bgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineGroup entity is found.
// Returns a *NotFoundError when no BaselineGroup entities are found.
func (bgq *BaselineGroupQuery) Only(ctx context.Context) (*BaselineGroup, error) {
	nodes, err := bgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinegroup.Label}
	default:
		return nil, &NotSingularError{baselinegroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bgq *BaselineGroupQuery) OnlyX(ctx context.Context) *BaselineGroup {
	node, err := bgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineGroup ID in the query.
// Returns a *NotSingularError when more than one BaselineGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (bgq *BaselineGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinegroup.Label}
	default:
		err = &NotSingularError{baselinegroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bgq *BaselineGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := bgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineGroups.
func (bgq *BaselineGroupQuery) All(ctx context.Context) ([]*BaselineGroup, error) {
	if err := bgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bgq *BaselineGroupQuery) AllX(ctx context.Context) []*BaselineGroup {
	nodes, err := bgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineGroup IDs.
func (bgq *BaselineGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bgq.Select(baselinegroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bgq *BaselineGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := bgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bgq *BaselineGroupQuery) Count(ctx context.Context) (int, error) {
	if err := bgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bgq *BaselineGroupQuery) CountX(ctx context.Context) int {
	count, err := bgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bgq *BaselineGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := bgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bgq *BaselineGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := bgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bgq *BaselineGroupQuery) Clone() *BaselineGroupQuery {
	if bgq == nil {
		return nil
	}
	return &BaselineGroupQuery{
		config:     bgq.config,
		limit:      bgq.limit,
		offset:     bgq.offset,
		order:      append([]OrderFunc{}, bgq.order...),
		predicates: append([]predicate.BaselineGroup{}, bgq.predicates...),
		withParent: bgq.withParent.Clone(),
		// clone intermediate query.
		sql:    bgq.sql.Clone(),
		path:   bgq.path,
		unique: bgq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bgq *BaselineGroupQuery) WithParent(opts ...func(*BaselineCharacteristicsModuleQuery)) *BaselineGroupQuery {
	query := &BaselineCharacteristicsModuleQuery{config: bgq.config}
	for _, opt := range opts {
		opt(query)
	}
	bgq.withParent = query
	return bgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineGroupID string `json:"baseline_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineGroup.Query().
//		GroupBy(baselinegroup.FieldBaselineGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bgq *BaselineGroupQuery) GroupBy(field string, fields ...string) *BaselineGroupGroupBy {
	grbuild := &BaselineGroupGroupBy{config: bgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bgq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinegroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineGroupID string `json:"baseline_group_id,omitempty"`
//	}
//
//	client.BaselineGroup.Query().
//		Select(baselinegroup.FieldBaselineGroupID).
//		Scan(ctx, &v)
//
func (bgq *BaselineGroupQuery) Select(fields ...string) *BaselineGroupSelect {
	bgq.fields = append(bgq.fields, fields...)
	selbuild := &BaselineGroupSelect{BaselineGroupQuery: bgq}
	selbuild.label = baselinegroup.Label
	selbuild.flds, selbuild.scan = &bgq.fields, selbuild.Scan
	return selbuild
}

func (bgq *BaselineGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bgq.fields {
		if !baselinegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bgq.path != nil {
		prev, err := bgq.path(ctx)
		if err != nil {
			return err
		}
		bgq.sql = prev
	}
	return nil
}

func (bgq *BaselineGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineGroup, error) {
	var (
		nodes       = []*BaselineGroup{}
		withFKs     = bgq.withFKs
		_spec       = bgq.querySpec()
		loadedTypes = [1]bool{
			bgq.withParent != nil,
		}
	)
	if bgq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinegroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineGroup{config: bgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bgq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineGroup)
		for i := range nodes {
			if nodes[i].baseline_characteristics_module_baseline_group_list == nil {
				continue
			}
			fk := *nodes[i].baseline_characteristics_module_baseline_group_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselinecharacteristicsmodule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_characteristics_module_baseline_group_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (bgq *BaselineGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bgq.querySpec()
	_spec.Node.Columns = bgq.fields
	if len(bgq.fields) > 0 {
		_spec.Unique = bgq.unique != nil && *bgq.unique
	}
	return sqlgraph.CountNodes(ctx, bgq.driver, _spec)
}

func (bgq *BaselineGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bgq *BaselineGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinegroup.Table,
			Columns: baselinegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinegroup.FieldID,
			},
		},
		From:   bgq.sql,
		Unique: true,
	}
	if unique := bgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinegroup.FieldID)
		for i := range fields {
			if fields[i] != baselinegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bgq *BaselineGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bgq.driver.Dialect())
	t1 := builder.Table(baselinegroup.Table)
	columns := bgq.fields
	if len(columns) == 0 {
		columns = baselinegroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bgq.sql != nil {
		selector = bgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bgq.unique != nil && *bgq.unique {
		selector.Distinct()
	}
	for _, p := range bgq.predicates {
		p(selector)
	}
	for _, p := range bgq.order {
		p(selector)
	}
	if offset := bgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineGroupGroupBy is the group-by builder for BaselineGroup entities.
type BaselineGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bggb *BaselineGroupGroupBy) Aggregate(fns ...AggregateFunc) *BaselineGroupGroupBy {
	bggb.fns = append(bggb.fns, fns...)
	return bggb
}

// Scan applies the group-by query and scans the result into the given value.
func (bggb *BaselineGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bggb.path(ctx)
	if err != nil {
		return err
	}
	bggb.sql = query
	return bggb.sqlScan(ctx, v)
}

func (bggb *BaselineGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bggb.fields {
		if !baselinegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bggb *BaselineGroupGroupBy) sqlQuery() *sql.Selector {
	selector := bggb.sql.Select()
	aggregation := make([]string, 0, len(bggb.fns))
	for _, fn := range bggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bggb.fields)+len(bggb.fns))
		for _, f := range bggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bggb.fields...)...)
}

// BaselineGroupSelect is the builder for selecting fields of BaselineGroup entities.
type BaselineGroupSelect struct {
	*BaselineGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bgs *BaselineGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bgs.prepareQuery(ctx); err != nil {
		return err
	}
	bgs.sql = bgs.BaselineGroupQuery.sqlQuery(ctx)
	return bgs.sqlScan(ctx, v)
}

func (bgs *BaselineGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bgs.sql.Query()
	if err := bgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
