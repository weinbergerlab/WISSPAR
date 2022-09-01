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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselineclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasuredenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineMeasureQuery is the builder for querying BaselineMeasure entities.
type BaselineMeasureQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineMeasure
	// eager-loading edges.
	withParent                   *BaselineCharacteristicsModuleQuery
	withBaselineMeasureDenomList *BaselineMeasureDenomQuery
	withBaselineClassList        *BaselineClassQuery
	withFKs                      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineMeasureQuery builder.
func (bmq *BaselineMeasureQuery) Where(ps ...predicate.BaselineMeasure) *BaselineMeasureQuery {
	bmq.predicates = append(bmq.predicates, ps...)
	return bmq
}

// Limit adds a limit step to the query.
func (bmq *BaselineMeasureQuery) Limit(limit int) *BaselineMeasureQuery {
	bmq.limit = &limit
	return bmq
}

// Offset adds an offset step to the query.
func (bmq *BaselineMeasureQuery) Offset(offset int) *BaselineMeasureQuery {
	bmq.offset = &offset
	return bmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bmq *BaselineMeasureQuery) Unique(unique bool) *BaselineMeasureQuery {
	bmq.unique = &unique
	return bmq
}

// Order adds an order step to the query.
func (bmq *BaselineMeasureQuery) Order(o ...OrderFunc) *BaselineMeasureQuery {
	bmq.order = append(bmq.order, o...)
	return bmq
}

// QueryParent chains the current query on the "parent" edge.
func (bmq *BaselineMeasureQuery) QueryParent() *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: bmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasure.Table, baselinemeasure.FieldID, selector),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinemeasure.ParentTable, baselinemeasure.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineMeasureDenomList chains the current query on the "baseline_measure_denom_list" edge.
func (bmq *BaselineMeasureQuery) QueryBaselineMeasureDenomList() *BaselineMeasureDenomQuery {
	query := &BaselineMeasureDenomQuery{config: bmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasure.Table, baselinemeasure.FieldID, selector),
			sqlgraph.To(baselinemeasuredenom.Table, baselinemeasuredenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinemeasure.BaselineMeasureDenomListTable, baselinemeasure.BaselineMeasureDenomListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineClassList chains the current query on the "baseline_class_list" edge.
func (bmq *BaselineMeasureQuery) QueryBaselineClassList() *BaselineClassQuery {
	query := &BaselineClassQuery{config: bmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinemeasure.Table, baselinemeasure.FieldID, selector),
			sqlgraph.To(baselineclass.Table, baselineclass.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinemeasure.BaselineClassListTable, baselinemeasure.BaselineClassListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineMeasure entity from the query.
// Returns a *NotFoundError when no BaselineMeasure was found.
func (bmq *BaselineMeasureQuery) First(ctx context.Context) (*BaselineMeasure, error) {
	nodes, err := bmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinemeasure.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) FirstX(ctx context.Context) *BaselineMeasure {
	node, err := bmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineMeasure ID from the query.
// Returns a *NotFoundError when no BaselineMeasure ID was found.
func (bmq *BaselineMeasureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinemeasure.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) FirstIDX(ctx context.Context) int {
	id, err := bmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineMeasure entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineMeasure entity is found.
// Returns a *NotFoundError when no BaselineMeasure entities are found.
func (bmq *BaselineMeasureQuery) Only(ctx context.Context) (*BaselineMeasure, error) {
	nodes, err := bmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinemeasure.Label}
	default:
		return nil, &NotSingularError{baselinemeasure.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) OnlyX(ctx context.Context) *BaselineMeasure {
	node, err := bmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineMeasure ID in the query.
// Returns a *NotSingularError when more than one BaselineMeasure ID is found.
// Returns a *NotFoundError when no entities are found.
func (bmq *BaselineMeasureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinemeasure.Label}
	default:
		err = &NotSingularError{baselinemeasure.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) OnlyIDX(ctx context.Context) int {
	id, err := bmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineMeasures.
func (bmq *BaselineMeasureQuery) All(ctx context.Context) ([]*BaselineMeasure, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) AllX(ctx context.Context) []*BaselineMeasure {
	nodes, err := bmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineMeasure IDs.
func (bmq *BaselineMeasureQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bmq.Select(baselinemeasure.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) IDsX(ctx context.Context) []int {
	ids, err := bmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bmq *BaselineMeasureQuery) Count(ctx context.Context) (int, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) CountX(ctx context.Context) int {
	count, err := bmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bmq *BaselineMeasureQuery) Exist(ctx context.Context) (bool, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bmq *BaselineMeasureQuery) ExistX(ctx context.Context) bool {
	exist, err := bmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineMeasureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bmq *BaselineMeasureQuery) Clone() *BaselineMeasureQuery {
	if bmq == nil {
		return nil
	}
	return &BaselineMeasureQuery{
		config:                       bmq.config,
		limit:                        bmq.limit,
		offset:                       bmq.offset,
		order:                        append([]OrderFunc{}, bmq.order...),
		predicates:                   append([]predicate.BaselineMeasure{}, bmq.predicates...),
		withParent:                   bmq.withParent.Clone(),
		withBaselineMeasureDenomList: bmq.withBaselineMeasureDenomList.Clone(),
		withBaselineClassList:        bmq.withBaselineClassList.Clone(),
		// clone intermediate query.
		sql:    bmq.sql.Clone(),
		path:   bmq.path,
		unique: bmq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bmq *BaselineMeasureQuery) WithParent(opts ...func(*BaselineCharacteristicsModuleQuery)) *BaselineMeasureQuery {
	query := &BaselineCharacteristicsModuleQuery{config: bmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmq.withParent = query
	return bmq
}

// WithBaselineMeasureDenomList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_measure_denom_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bmq *BaselineMeasureQuery) WithBaselineMeasureDenomList(opts ...func(*BaselineMeasureDenomQuery)) *BaselineMeasureQuery {
	query := &BaselineMeasureDenomQuery{config: bmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmq.withBaselineMeasureDenomList = query
	return bmq
}

// WithBaselineClassList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_class_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bmq *BaselineMeasureQuery) WithBaselineClassList(opts ...func(*BaselineClassQuery)) *BaselineMeasureQuery {
	query := &BaselineClassQuery{config: bmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmq.withBaselineClassList = query
	return bmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineMeasureTitle string `json:"baseline_measure_title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineMeasure.Query().
//		GroupBy(baselinemeasure.FieldBaselineMeasureTitle).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bmq *BaselineMeasureQuery) GroupBy(field string, fields ...string) *BaselineMeasureGroupBy {
	grbuild := &BaselineMeasureGroupBy{config: bmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bmq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinemeasure.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineMeasureTitle string `json:"baseline_measure_title,omitempty"`
//	}
//
//	client.BaselineMeasure.Query().
//		Select(baselinemeasure.FieldBaselineMeasureTitle).
//		Scan(ctx, &v)
//
func (bmq *BaselineMeasureQuery) Select(fields ...string) *BaselineMeasureSelect {
	bmq.fields = append(bmq.fields, fields...)
	selbuild := &BaselineMeasureSelect{BaselineMeasureQuery: bmq}
	selbuild.label = baselinemeasure.Label
	selbuild.flds, selbuild.scan = &bmq.fields, selbuild.Scan
	return selbuild
}

func (bmq *BaselineMeasureQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bmq.fields {
		if !baselinemeasure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bmq.path != nil {
		prev, err := bmq.path(ctx)
		if err != nil {
			return err
		}
		bmq.sql = prev
	}
	return nil
}

func (bmq *BaselineMeasureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineMeasure, error) {
	var (
		nodes       = []*BaselineMeasure{}
		withFKs     = bmq.withFKs
		_spec       = bmq.querySpec()
		loadedTypes = [3]bool{
			bmq.withParent != nil,
			bmq.withBaselineMeasureDenomList != nil,
			bmq.withBaselineClassList != nil,
		}
	)
	if bmq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasure.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineMeasure).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineMeasure{config: bmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bmq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineMeasure)
		for i := range nodes {
			if nodes[i].baseline_characteristics_module_baseline_measure_list == nil {
				continue
			}
			fk := *nodes[i].baseline_characteristics_module_baseline_measure_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_characteristics_module_baseline_measure_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bmq.withBaselineMeasureDenomList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineMeasureDenomList = []*BaselineMeasureDenom{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineMeasureDenom(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinemeasure.BaselineMeasureDenomListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_measure_baseline_measure_denom_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_measure_baseline_measure_denom_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_measure_baseline_measure_denom_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineMeasureDenomList = append(node.Edges.BaselineMeasureDenomList, n)
		}
	}

	if query := bmq.withBaselineClassList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineClassList = []*BaselineClass{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineClass(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinemeasure.BaselineClassListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_measure_baseline_class_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_measure_baseline_class_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_measure_baseline_class_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineClassList = append(node.Edges.BaselineClassList, n)
		}
	}

	return nodes, nil
}

func (bmq *BaselineMeasureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bmq.querySpec()
	_spec.Node.Columns = bmq.fields
	if len(bmq.fields) > 0 {
		_spec.Unique = bmq.unique != nil && *bmq.unique
	}
	return sqlgraph.CountNodes(ctx, bmq.driver, _spec)
}

func (bmq *BaselineMeasureQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bmq *BaselineMeasureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinemeasure.Table,
			Columns: baselinemeasure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinemeasure.FieldID,
			},
		},
		From:   bmq.sql,
		Unique: true,
	}
	if unique := bmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinemeasure.FieldID)
		for i := range fields {
			if fields[i] != baselinemeasure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bmq *BaselineMeasureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bmq.driver.Dialect())
	t1 := builder.Table(baselinemeasure.Table)
	columns := bmq.fields
	if len(columns) == 0 {
		columns = baselinemeasure.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bmq.sql != nil {
		selector = bmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bmq.unique != nil && *bmq.unique {
		selector.Distinct()
	}
	for _, p := range bmq.predicates {
		p(selector)
	}
	for _, p := range bmq.order {
		p(selector)
	}
	if offset := bmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineMeasureGroupBy is the group-by builder for BaselineMeasure entities.
type BaselineMeasureGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bmgb *BaselineMeasureGroupBy) Aggregate(fns ...AggregateFunc) *BaselineMeasureGroupBy {
	bmgb.fns = append(bmgb.fns, fns...)
	return bmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bmgb *BaselineMeasureGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bmgb.path(ctx)
	if err != nil {
		return err
	}
	bmgb.sql = query
	return bmgb.sqlScan(ctx, v)
}

func (bmgb *BaselineMeasureGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bmgb.fields {
		if !baselinemeasure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bmgb *BaselineMeasureGroupBy) sqlQuery() *sql.Selector {
	selector := bmgb.sql.Select()
	aggregation := make([]string, 0, len(bmgb.fns))
	for _, fn := range bmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bmgb.fields)+len(bmgb.fns))
		for _, f := range bmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bmgb.fields...)...)
}

// BaselineMeasureSelect is the builder for selecting fields of BaselineMeasure entities.
type BaselineMeasureSelect struct {
	*BaselineMeasureQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bms *BaselineMeasureSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bms.prepareQuery(ctx); err != nil {
		return err
	}
	bms.sql = bms.BaselineMeasureQuery.sqlQuery(ctx)
	return bms.sqlScan(ctx, v)
}

func (bms *BaselineMeasureSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bms.sql.Query()
	if err := bms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
