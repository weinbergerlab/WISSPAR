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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclassdenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineClassDenomQuery is the builder for querying BaselineClassDenom entities.
type BaselineClassDenomQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineClassDenom
	// eager-loading edges.
	withParent                      *BaselineClassQuery
	withBaselineClassDenomCountList *BaselineClassDenomCountQuery
	withFKs                         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineClassDenomQuery builder.
func (bcdq *BaselineClassDenomQuery) Where(ps ...predicate.BaselineClassDenom) *BaselineClassDenomQuery {
	bcdq.predicates = append(bcdq.predicates, ps...)
	return bcdq
}

// Limit adds a limit step to the query.
func (bcdq *BaselineClassDenomQuery) Limit(limit int) *BaselineClassDenomQuery {
	bcdq.limit = &limit
	return bcdq
}

// Offset adds an offset step to the query.
func (bcdq *BaselineClassDenomQuery) Offset(offset int) *BaselineClassDenomQuery {
	bcdq.offset = &offset
	return bcdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcdq *BaselineClassDenomQuery) Unique(unique bool) *BaselineClassDenomQuery {
	bcdq.unique = &unique
	return bcdq
}

// Order adds an order step to the query.
func (bcdq *BaselineClassDenomQuery) Order(o ...OrderFunc) *BaselineClassDenomQuery {
	bcdq.order = append(bcdq.order, o...)
	return bcdq
}

// QueryParent chains the current query on the "parent" edge.
func (bcdq *BaselineClassDenomQuery) QueryParent() *BaselineClassQuery {
	query := &BaselineClassQuery{config: bcdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclassdenom.Table, baselineclassdenom.FieldID, selector),
			sqlgraph.To(baselineclass.Table, baselineclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselineclassdenom.ParentTable, baselineclassdenom.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineClassDenomCountList chains the current query on the "baseline_class_denom_count_list" edge.
func (bcdq *BaselineClassDenomQuery) QueryBaselineClassDenomCountList() *BaselineClassDenomCountQuery {
	query := &BaselineClassDenomCountQuery{config: bcdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselineclassdenom.Table, baselineclassdenom.FieldID, selector),
			sqlgraph.To(baselineclassdenomcount.Table, baselineclassdenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselineclassdenom.BaselineClassDenomCountListTable, baselineclassdenom.BaselineClassDenomCountListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineClassDenom entity from the query.
// Returns a *NotFoundError when no BaselineClassDenom was found.
func (bcdq *BaselineClassDenomQuery) First(ctx context.Context) (*BaselineClassDenom, error) {
	nodes, err := bcdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselineclassdenom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) FirstX(ctx context.Context) *BaselineClassDenom {
	node, err := bcdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineClassDenom ID from the query.
// Returns a *NotFoundError when no BaselineClassDenom ID was found.
func (bcdq *BaselineClassDenomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselineclassdenom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) FirstIDX(ctx context.Context) int {
	id, err := bcdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineClassDenom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineClassDenom entity is found.
// Returns a *NotFoundError when no BaselineClassDenom entities are found.
func (bcdq *BaselineClassDenomQuery) Only(ctx context.Context) (*BaselineClassDenom, error) {
	nodes, err := bcdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselineclassdenom.Label}
	default:
		return nil, &NotSingularError{baselineclassdenom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) OnlyX(ctx context.Context) *BaselineClassDenom {
	node, err := bcdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineClassDenom ID in the query.
// Returns a *NotSingularError when more than one BaselineClassDenom ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcdq *BaselineClassDenomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselineclassdenom.Label}
	default:
		err = &NotSingularError{baselineclassdenom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineClassDenoms.
func (bcdq *BaselineClassDenomQuery) All(ctx context.Context) ([]*BaselineClassDenom, error) {
	if err := bcdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) AllX(ctx context.Context) []*BaselineClassDenom {
	nodes, err := bcdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineClassDenom IDs.
func (bcdq *BaselineClassDenomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bcdq.Select(baselineclassdenom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) IDsX(ctx context.Context) []int {
	ids, err := bcdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcdq *BaselineClassDenomQuery) Count(ctx context.Context) (int, error) {
	if err := bcdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) CountX(ctx context.Context) int {
	count, err := bcdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcdq *BaselineClassDenomQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcdq *BaselineClassDenomQuery) ExistX(ctx context.Context) bool {
	exist, err := bcdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineClassDenomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcdq *BaselineClassDenomQuery) Clone() *BaselineClassDenomQuery {
	if bcdq == nil {
		return nil
	}
	return &BaselineClassDenomQuery{
		config:                          bcdq.config,
		limit:                           bcdq.limit,
		offset:                          bcdq.offset,
		order:                           append([]OrderFunc{}, bcdq.order...),
		predicates:                      append([]predicate.BaselineClassDenom{}, bcdq.predicates...),
		withParent:                      bcdq.withParent.Clone(),
		withBaselineClassDenomCountList: bcdq.withBaselineClassDenomCountList.Clone(),
		// clone intermediate query.
		sql:    bcdq.sql.Clone(),
		path:   bcdq.path,
		unique: bcdq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bcdq *BaselineClassDenomQuery) WithParent(opts ...func(*BaselineClassQuery)) *BaselineClassDenomQuery {
	query := &BaselineClassQuery{config: bcdq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcdq.withParent = query
	return bcdq
}

// WithBaselineClassDenomCountList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_class_denom_count_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcdq *BaselineClassDenomQuery) WithBaselineClassDenomCountList(opts ...func(*BaselineClassDenomCountQuery)) *BaselineClassDenomQuery {
	query := &BaselineClassDenomCountQuery{config: bcdq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcdq.withBaselineClassDenomCountList = query
	return bcdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineClassDenomUnits string `json:"baseline_class_denom_units,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineClassDenom.Query().
//		GroupBy(baselineclassdenom.FieldBaselineClassDenomUnits).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bcdq *BaselineClassDenomQuery) GroupBy(field string, fields ...string) *BaselineClassDenomGroupBy {
	grbuild := &BaselineClassDenomGroupBy{config: bcdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcdq.sqlQuery(ctx), nil
	}
	grbuild.label = baselineclassdenom.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineClassDenomUnits string `json:"baseline_class_denom_units,omitempty"`
//	}
//
//	client.BaselineClassDenom.Query().
//		Select(baselineclassdenom.FieldBaselineClassDenomUnits).
//		Scan(ctx, &v)
//
func (bcdq *BaselineClassDenomQuery) Select(fields ...string) *BaselineClassDenomSelect {
	bcdq.fields = append(bcdq.fields, fields...)
	selbuild := &BaselineClassDenomSelect{BaselineClassDenomQuery: bcdq}
	selbuild.label = baselineclassdenom.Label
	selbuild.flds, selbuild.scan = &bcdq.fields, selbuild.Scan
	return selbuild
}

func (bcdq *BaselineClassDenomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcdq.fields {
		if !baselineclassdenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bcdq.path != nil {
		prev, err := bcdq.path(ctx)
		if err != nil {
			return err
		}
		bcdq.sql = prev
	}
	return nil
}

func (bcdq *BaselineClassDenomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineClassDenom, error) {
	var (
		nodes       = []*BaselineClassDenom{}
		withFKs     = bcdq.withFKs
		_spec       = bcdq.querySpec()
		loadedTypes = [2]bool{
			bcdq.withParent != nil,
			bcdq.withBaselineClassDenomCountList != nil,
		}
	)
	if bcdq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclassdenom.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineClassDenom).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineClassDenom{config: bcdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bcdq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineClassDenom)
		for i := range nodes {
			if nodes[i].baseline_class_baseline_class_denom_list == nil {
				continue
			}
			fk := *nodes[i].baseline_class_baseline_class_denom_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(baselineclass.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_class_baseline_class_denom_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bcdq.withBaselineClassDenomCountList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineClassDenom)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineClassDenomCountList = []*BaselineClassDenomCount{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineClassDenomCount(func(s *sql.Selector) {
			s.Where(sql.InValues(baselineclassdenom.BaselineClassDenomCountListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_class_denom_baseline_class_denom_count_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_class_denom_baseline_class_denom_count_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_class_denom_baseline_class_denom_count_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineClassDenomCountList = append(node.Edges.BaselineClassDenomCountList, n)
		}
	}

	return nodes, nil
}

func (bcdq *BaselineClassDenomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcdq.querySpec()
	_spec.Node.Columns = bcdq.fields
	if len(bcdq.fields) > 0 {
		_spec.Unique = bcdq.unique != nil && *bcdq.unique
	}
	return sqlgraph.CountNodes(ctx, bcdq.driver, _spec)
}

func (bcdq *BaselineClassDenomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bcdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bcdq *BaselineClassDenomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselineclassdenom.Table,
			Columns: baselineclassdenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselineclassdenom.FieldID,
			},
		},
		From:   bcdq.sql,
		Unique: true,
	}
	if unique := bcdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselineclassdenom.FieldID)
		for i := range fields {
			if fields[i] != baselineclassdenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcdq *BaselineClassDenomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcdq.driver.Dialect())
	t1 := builder.Table(baselineclassdenom.Table)
	columns := bcdq.fields
	if len(columns) == 0 {
		columns = baselineclassdenom.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcdq.sql != nil {
		selector = bcdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcdq.unique != nil && *bcdq.unique {
		selector.Distinct()
	}
	for _, p := range bcdq.predicates {
		p(selector)
	}
	for _, p := range bcdq.order {
		p(selector)
	}
	if offset := bcdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineClassDenomGroupBy is the group-by builder for BaselineClassDenom entities.
type BaselineClassDenomGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcdgb *BaselineClassDenomGroupBy) Aggregate(fns ...AggregateFunc) *BaselineClassDenomGroupBy {
	bcdgb.fns = append(bcdgb.fns, fns...)
	return bcdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcdgb *BaselineClassDenomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bcdgb.path(ctx)
	if err != nil {
		return err
	}
	bcdgb.sql = query
	return bcdgb.sqlScan(ctx, v)
}

func (bcdgb *BaselineClassDenomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bcdgb.fields {
		if !baselineclassdenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcdgb *BaselineClassDenomGroupBy) sqlQuery() *sql.Selector {
	selector := bcdgb.sql.Select()
	aggregation := make([]string, 0, len(bcdgb.fns))
	for _, fn := range bcdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcdgb.fields)+len(bcdgb.fns))
		for _, f := range bcdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcdgb.fields...)...)
}

// BaselineClassDenomSelect is the builder for selecting fields of BaselineClassDenom entities.
type BaselineClassDenomSelect struct {
	*BaselineClassDenomQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcds *BaselineClassDenomSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bcds.prepareQuery(ctx); err != nil {
		return err
	}
	bcds.sql = bcds.BaselineClassDenomQuery.sqlQuery(ctx)
	return bcds.sqlScan(ctx, v)
}

func (bcds *BaselineClassDenomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bcds.sql.Query()
	if err := bcds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
