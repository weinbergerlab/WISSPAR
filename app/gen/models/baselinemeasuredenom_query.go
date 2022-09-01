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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureDenomQuery is the builder for querying BaselineMeasureDenom entities.
type BaselineMeasureDenomQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineMeasureDenom
	// eager-loading edges.
	withParent                        *BaselineMeasureQuery
	withBaselineMeasureDenomCountList *BaselineMeasureDenomCountQuery
	withFKs                           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineMeasureDenomQuery builder.
func (bmdq *BaselineMeasureDenomQuery) Where(ps ...predicate.BaselineMeasureDenom) *BaselineMeasureDenomQuery {
	bmdq.predicates = append(bmdq.predicates, ps...)
	return bmdq
}

// Limit adds a limit step to the query.
func (bmdq *BaselineMeasureDenomQuery) Limit(limit int) *BaselineMeasureDenomQuery {
	bmdq.limit = &limit
	return bmdq
}

// Offset adds an offset step to the query.
func (bmdq *BaselineMeasureDenomQuery) Offset(offset int) *BaselineMeasureDenomQuery {
	bmdq.offset = &offset
	return bmdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bmdq *BaselineMeasureDenomQuery) Unique(unique bool) *BaselineMeasureDenomQuery {
	bmdq.unique = &unique
	return bmdq
}

// Order adds an order step to the query.
func (bmdq *BaselineMeasureDenomQuery) Order(o ...OrderFunc) *BaselineMeasureDenomQuery {
	bmdq.order = append(bmdq.order, o...)
	return bmdq
}

// QueryParent chains the current query on the "parent" edge.
func (bmdq *BaselineMeasureDenomQuery) QueryParent() *BaselineMeasureQuery {
	query := &BaselineMeasureQuery{config: bmdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID, selector),
			sqlgraph.To(baselinemeasure.Table, baselinemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasuredenom.ParentTable, baselinemeasuredenom.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineMeasureDenomCountList chains the current query on the "baseline_measure_denom_count_list" edge.
func (bmdq *BaselineMeasureDenomQuery) QueryBaselineMeasureDenomCountList() *BaselineMeasureDenomCountQuery {
	query := &BaselineMeasureDenomCountQuery{config: bmdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID, selector),
			sqlgraph.To(baselinemeasuredenomcount.Table, baselinemeasuredenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinemeasuredenom.BaselineMeasureDenomCountListTable, baselinemeasuredenom.BaselineMeasureDenomCountListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineMeasureDenom entity from the query.
// Returns a *NotFoundError when no BaselineMeasureDenom was found.
func (bmdq *BaselineMeasureDenomQuery) First(ctx context.Context) (*BaselineMeasureDenom, error) {
	nodes, err := bmdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinemeasuredenom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) FirstX(ctx context.Context) *BaselineMeasureDenom {
	node, err := bmdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineMeasureDenom ID from the query.
// Returns a *NotFoundError when no BaselineMeasureDenom ID was found.
func (bmdq *BaselineMeasureDenomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinemeasuredenom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) FirstIDX(ctx context.Context) int {
	id, err := bmdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineMeasureDenom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineMeasureDenom entity is found.
// Returns a *NotFoundError when no BaselineMeasureDenom entities are found.
func (bmdq *BaselineMeasureDenomQuery) Only(ctx context.Context) (*BaselineMeasureDenom, error) {
	nodes, err := bmdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinemeasuredenom.Label}
	default:
		return nil, &NotSingularError{baselinemeasuredenom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) OnlyX(ctx context.Context) *BaselineMeasureDenom {
	node, err := bmdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineMeasureDenom ID in the query.
// Returns a *NotSingularError when more than one BaselineMeasureDenom ID is found.
// Returns a *NotFoundError when no entities are found.
func (bmdq *BaselineMeasureDenomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinemeasuredenom.Label}
	default:
		err = &NotSingularError{baselinemeasuredenom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) OnlyIDX(ctx context.Context) int {
	id, err := bmdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineMeasureDenoms.
func (bmdq *BaselineMeasureDenomQuery) All(ctx context.Context) ([]*BaselineMeasureDenom, error) {
	if err := bmdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bmdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) AllX(ctx context.Context) []*BaselineMeasureDenom {
	nodes, err := bmdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineMeasureDenom IDs.
func (bmdq *BaselineMeasureDenomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bmdq.Select(baselinemeasuredenom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) IDsX(ctx context.Context) []int {
	ids, err := bmdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bmdq *BaselineMeasureDenomQuery) Count(ctx context.Context) (int, error) {
	if err := bmdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bmdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) CountX(ctx context.Context) int {
	count, err := bmdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bmdq *BaselineMeasureDenomQuery) Exist(ctx context.Context) (bool, error) {
	if err := bmdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bmdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bmdq *BaselineMeasureDenomQuery) ExistX(ctx context.Context) bool {
	exist, err := bmdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineMeasureDenomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bmdq *BaselineMeasureDenomQuery) Clone() *BaselineMeasureDenomQuery {
	if bmdq == nil {
		return nil
	}
	return &BaselineMeasureDenomQuery{
		config:                            bmdq.config,
		limit:                             bmdq.limit,
		offset:                            bmdq.offset,
		order:                             append([]OrderFunc{}, bmdq.order...),
		predicates:                        append([]predicate.BaselineMeasureDenom{}, bmdq.predicates...),
		withParent:                        bmdq.withParent.Clone(),
		withBaselineMeasureDenomCountList: bmdq.withBaselineMeasureDenomCountList.Clone(),
		// clone intermediate query.
		sql:    bmdq.sql.Clone(),
		path:   bmdq.path,
		unique: bmdq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bmdq *BaselineMeasureDenomQuery) WithParent(opts ...func(*BaselineMeasureQuery)) *BaselineMeasureDenomQuery {
	query := &BaselineMeasureQuery{config: bmdq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmdq.withParent = query
	return bmdq
}

// WithBaselineMeasureDenomCountList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_measure_denom_count_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bmdq *BaselineMeasureDenomQuery) WithBaselineMeasureDenomCountList(opts ...func(*BaselineMeasureDenomCountQuery)) *BaselineMeasureDenomQuery {
	query := &BaselineMeasureDenomCountQuery{config: bmdq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmdq.withBaselineMeasureDenomCountList = query
	return bmdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineMeasureDenomUnits string `json:"baseline_measure_denom_units,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineMeasureDenom.Query().
//		GroupBy(baselinemeasuredenom.FieldBaselineMeasureDenomUnits).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bmdq *BaselineMeasureDenomQuery) GroupBy(field string, fields ...string) *BaselineMeasureDenomGroupBy {
	grbuild := &BaselineMeasureDenomGroupBy{config: bmdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bmdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bmdq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinemeasuredenom.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineMeasureDenomUnits string `json:"baseline_measure_denom_units,omitempty"`
//	}
//
//	client.BaselineMeasureDenom.Query().
//		Select(baselinemeasuredenom.FieldBaselineMeasureDenomUnits).
//		Scan(ctx, &v)
//
func (bmdq *BaselineMeasureDenomQuery) Select(fields ...string) *BaselineMeasureDenomSelect {
	bmdq.fields = append(bmdq.fields, fields...)
	selbuild := &BaselineMeasureDenomSelect{BaselineMeasureDenomQuery: bmdq}
	selbuild.label = baselinemeasuredenom.Label
	selbuild.flds, selbuild.scan = &bmdq.fields, selbuild.Scan
	return selbuild
}

func (bmdq *BaselineMeasureDenomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bmdq.fields {
		if !baselinemeasuredenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bmdq.path != nil {
		prev, err := bmdq.path(ctx)
		if err != nil {
			return err
		}
		bmdq.sql = prev
	}
	return nil
}

func (bmdq *BaselineMeasureDenomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineMeasureDenom, error) {
	var (
		nodes       = []*BaselineMeasureDenom{}
		withFKs     = bmdq.withFKs
		_spec       = bmdq.querySpec()
		loadedTypes = [2]bool{
			bmdq.withParent != nil,
			bmdq.withBaselineMeasureDenomCountList != nil,
		}
	)
	if bmdq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasuredenom.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineMeasureDenom).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineMeasureDenom{config: bmdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bmdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bmdq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineMeasureDenom)
		for i := range nodes {
			if nodes[i].baseline_measure_baseline_measure_denom_list == nil {
				continue
			}
			fk := *nodes[i].baseline_measure_baseline_measure_denom_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselinemeasure.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_measure_baseline_measure_denom_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bmdq.withBaselineMeasureDenomCountList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineMeasureDenom)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineMeasureDenomCountList = []*BaselineMeasureDenomCount{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineMeasureDenomCount(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinemeasuredenom.BaselineMeasureDenomCountListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_measure_denom_baseline_measure_denom_count_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_measure_denom_baseline_measure_denom_count_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_measure_denom_baseline_measure_denom_count_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineMeasureDenomCountList = append(node.Edges.BaselineMeasureDenomCountList, n)
		}
	}

	return nodes, nil
}

func (bmdq *BaselineMeasureDenomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bmdq.querySpec()
	_spec.Node.Columns = bmdq.fields
	if len(bmdq.fields) > 0 {
		_spec.Unique = bmdq.unique != nil && *bmdq.unique
	}
	return sqlgraph.CountNodes(ctx, bmdq.driver, _spec)
}

func (bmdq *BaselineMeasureDenomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bmdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bmdq *BaselineMeasureDenomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasuredenom.Table,
			Columns: baselinemeasuredenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasuredenom.FieldID,
			},
		},
		From:   bmdq.sql,
		Unique: true,
	}
	if unique := bmdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bmdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasuredenom.FieldID)
		for i := range fields {
			if fields[i] != baselinemeasuredenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bmdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bmdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bmdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bmdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bmdq *BaselineMeasureDenomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bmdq.driver.Dialect())
	t1 := builder.Table(baselinemeasuredenom.Table)
	columns := bmdq.fields
	if len(columns) == 0 {
		columns = baselinemeasuredenom.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bmdq.sql != nil {
		selector = bmdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bmdq.unique != nil && *bmdq.unique {
		selector.Distinct()
	}
	for _, p := range bmdq.predicates {
		p(selector)
	}
	for _, p := range bmdq.order {
		p(selector)
	}
	if offset := bmdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bmdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineMeasureDenomGroupBy is the group-by builder for BaselineMeasureDenom entities.
type BaselineMeasureDenomGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bmdgb *BaselineMeasureDenomGroupBy) Aggregate(fns ...AggregateFunc) *BaselineMeasureDenomGroupBy {
	bmdgb.fns = append(bmdgb.fns, fns...)
	return bmdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bmdgb *BaselineMeasureDenomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bmdgb.path(ctx)
	if err != nil {
		return err
	}
	bmdgb.sql = query
	return bmdgb.sqlScan(ctx, v)
}

func (bmdgb *BaselineMeasureDenomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bmdgb.fields {
		if !baselinemeasuredenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bmdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bmdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bmdgb *BaselineMeasureDenomGroupBy) sqlQuery() *sql.Selector {
	selector := bmdgb.sql.Select()
	aggregation := make([]string, 0, len(bmdgb.fns))
	for _, fn := range bmdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bmdgb.fields)+len(bmdgb.fns))
		for _, f := range bmdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bmdgb.fields...)...)
}

// BaselineMeasureDenomSelect is the builder for selecting fields of BaselineMeasureDenom entities.
type BaselineMeasureDenomSelect struct {
	*BaselineMeasureDenomQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bmds *BaselineMeasureDenomSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bmds.prepareQuery(ctx); err != nil {
		return err
	}
	bmds.sql = bmds.BaselineMeasureDenomQuery.sqlQuery(ctx)
	return bmds.sqlScan(ctx, v)
}

func (bmds *BaselineMeasureDenomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bmds.sql.Query()
	if err := bmds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
