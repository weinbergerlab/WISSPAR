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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// BaselineDenomQuery is the builder for querying BaselineDenom entities.
type BaselineDenomQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineDenom
	// eager-loading edges.
	withParent                 *BaselineCharacteristicsModuleQuery
	withBaselineDenomCountList *BaselineDenomCountQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineDenomQuery builder.
func (bdq *BaselineDenomQuery) Where(ps ...predicate.BaselineDenom) *BaselineDenomQuery {
	bdq.predicates = append(bdq.predicates, ps...)
	return bdq
}

// Limit adds a limit step to the query.
func (bdq *BaselineDenomQuery) Limit(limit int) *BaselineDenomQuery {
	bdq.limit = &limit
	return bdq
}

// Offset adds an offset step to the query.
func (bdq *BaselineDenomQuery) Offset(offset int) *BaselineDenomQuery {
	bdq.offset = &offset
	return bdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bdq *BaselineDenomQuery) Unique(unique bool) *BaselineDenomQuery {
	bdq.unique = &unique
	return bdq
}

// Order adds an order step to the query.
func (bdq *BaselineDenomQuery) Order(o ...OrderFunc) *BaselineDenomQuery {
	bdq.order = append(bdq.order, o...)
	return bdq
}

// QueryParent chains the current query on the "parent" edge.
func (bdq *BaselineDenomQuery) QueryParent() *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: bdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinedenom.Table, baselinedenom.FieldID, selector),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, baselinedenom.ParentTable, baselinedenom.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineDenomCountList chains the current query on the "baseline_denom_count_list" edge.
func (bdq *BaselineDenomQuery) QueryBaselineDenomCountList() *BaselineDenomCountQuery {
	query := &BaselineDenomCountQuery{config: bdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinedenom.Table, baselinedenom.FieldID, selector),
			sqlgraph.To(baselinedenomcount.Table, baselinedenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinedenom.BaselineDenomCountListTable, baselinedenom.BaselineDenomCountListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineDenom entity from the query.
// Returns a *NotFoundError when no BaselineDenom was found.
func (bdq *BaselineDenomQuery) First(ctx context.Context) (*BaselineDenom, error) {
	nodes, err := bdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinedenom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bdq *BaselineDenomQuery) FirstX(ctx context.Context) *BaselineDenom {
	node, err := bdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineDenom ID from the query.
// Returns a *NotFoundError when no BaselineDenom ID was found.
func (bdq *BaselineDenomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinedenom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bdq *BaselineDenomQuery) FirstIDX(ctx context.Context) int {
	id, err := bdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineDenom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineDenom entity is found.
// Returns a *NotFoundError when no BaselineDenom entities are found.
func (bdq *BaselineDenomQuery) Only(ctx context.Context) (*BaselineDenom, error) {
	nodes, err := bdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinedenom.Label}
	default:
		return nil, &NotSingularError{baselinedenom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bdq *BaselineDenomQuery) OnlyX(ctx context.Context) *BaselineDenom {
	node, err := bdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineDenom ID in the query.
// Returns a *NotSingularError when more than one BaselineDenom ID is found.
// Returns a *NotFoundError when no entities are found.
func (bdq *BaselineDenomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinedenom.Label}
	default:
		err = &NotSingularError{baselinedenom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bdq *BaselineDenomQuery) OnlyIDX(ctx context.Context) int {
	id, err := bdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineDenoms.
func (bdq *BaselineDenomQuery) All(ctx context.Context) ([]*BaselineDenom, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bdq *BaselineDenomQuery) AllX(ctx context.Context) []*BaselineDenom {
	nodes, err := bdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineDenom IDs.
func (bdq *BaselineDenomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bdq.Select(baselinedenom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bdq *BaselineDenomQuery) IDsX(ctx context.Context) []int {
	ids, err := bdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bdq *BaselineDenomQuery) Count(ctx context.Context) (int, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bdq *BaselineDenomQuery) CountX(ctx context.Context) int {
	count, err := bdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bdq *BaselineDenomQuery) Exist(ctx context.Context) (bool, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bdq *BaselineDenomQuery) ExistX(ctx context.Context) bool {
	exist, err := bdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineDenomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bdq *BaselineDenomQuery) Clone() *BaselineDenomQuery {
	if bdq == nil {
		return nil
	}
	return &BaselineDenomQuery{
		config:                     bdq.config,
		limit:                      bdq.limit,
		offset:                     bdq.offset,
		order:                      append([]OrderFunc{}, bdq.order...),
		predicates:                 append([]predicate.BaselineDenom{}, bdq.predicates...),
		withParent:                 bdq.withParent.Clone(),
		withBaselineDenomCountList: bdq.withBaselineDenomCountList.Clone(),
		// clone intermediate query.
		sql:    bdq.sql.Clone(),
		path:   bdq.path,
		unique: bdq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bdq *BaselineDenomQuery) WithParent(opts ...func(*BaselineCharacteristicsModuleQuery)) *BaselineDenomQuery {
	query := &BaselineCharacteristicsModuleQuery{config: bdq.config}
	for _, opt := range opts {
		opt(query)
	}
	bdq.withParent = query
	return bdq
}

// WithBaselineDenomCountList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_denom_count_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bdq *BaselineDenomQuery) WithBaselineDenomCountList(opts ...func(*BaselineDenomCountQuery)) *BaselineDenomQuery {
	query := &BaselineDenomCountQuery{config: bdq.config}
	for _, opt := range opts {
		opt(query)
	}
	bdq.withBaselineDenomCountList = query
	return bdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselineDenomUnits string `json:"baseline_denom_units,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineDenom.Query().
//		GroupBy(baselinedenom.FieldBaselineDenomUnits).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bdq *BaselineDenomQuery) GroupBy(field string, fields ...string) *BaselineDenomGroupBy {
	grbuild := &BaselineDenomGroupBy{config: bdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bdq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinedenom.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselineDenomUnits string `json:"baseline_denom_units,omitempty"`
//	}
//
//	client.BaselineDenom.Query().
//		Select(baselinedenom.FieldBaselineDenomUnits).
//		Scan(ctx, &v)
//
func (bdq *BaselineDenomQuery) Select(fields ...string) *BaselineDenomSelect {
	bdq.fields = append(bdq.fields, fields...)
	selbuild := &BaselineDenomSelect{BaselineDenomQuery: bdq}
	selbuild.label = baselinedenom.Label
	selbuild.flds, selbuild.scan = &bdq.fields, selbuild.Scan
	return selbuild
}

func (bdq *BaselineDenomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bdq.fields {
		if !baselinedenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bdq.path != nil {
		prev, err := bdq.path(ctx)
		if err != nil {
			return err
		}
		bdq.sql = prev
	}
	return nil
}

func (bdq *BaselineDenomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineDenom, error) {
	var (
		nodes       = []*BaselineDenom{}
		withFKs     = bdq.withFKs
		_spec       = bdq.querySpec()
		loadedTypes = [2]bool{
			bdq.withParent != nil,
			bdq.withBaselineDenomCountList != nil,
		}
	)
	if bdq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinedenom.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineDenom).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineDenom{config: bdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bdq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineDenom)
		for i := range nodes {
			if nodes[i].baseline_characteristics_module_baseline_denom_list == nil {
				continue
			}
			fk := *nodes[i].baseline_characteristics_module_baseline_denom_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_characteristics_module_baseline_denom_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bdq.withBaselineDenomCountList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineDenom)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineDenomCountList = []*BaselineDenomCount{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineDenomCount(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinedenom.BaselineDenomCountListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_denom_baseline_denom_count_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_denom_baseline_denom_count_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_denom_baseline_denom_count_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineDenomCountList = append(node.Edges.BaselineDenomCountList, n)
		}
	}

	return nodes, nil
}

func (bdq *BaselineDenomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bdq.querySpec()
	_spec.Node.Columns = bdq.fields
	if len(bdq.fields) > 0 {
		_spec.Unique = bdq.unique != nil && *bdq.unique
	}
	return sqlgraph.CountNodes(ctx, bdq.driver, _spec)
}

func (bdq *BaselineDenomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bdq *BaselineDenomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinedenom.Table,
			Columns: baselinedenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinedenom.FieldID,
			},
		},
		From:   bdq.sql,
		Unique: true,
	}
	if unique := bdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinedenom.FieldID)
		for i := range fields {
			if fields[i] != baselinedenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bdq *BaselineDenomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bdq.driver.Dialect())
	t1 := builder.Table(baselinedenom.Table)
	columns := bdq.fields
	if len(columns) == 0 {
		columns = baselinedenom.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bdq.sql != nil {
		selector = bdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bdq.unique != nil && *bdq.unique {
		selector.Distinct()
	}
	for _, p := range bdq.predicates {
		p(selector)
	}
	for _, p := range bdq.order {
		p(selector)
	}
	if offset := bdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineDenomGroupBy is the group-by builder for BaselineDenom entities.
type BaselineDenomGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bdgb *BaselineDenomGroupBy) Aggregate(fns ...AggregateFunc) *BaselineDenomGroupBy {
	bdgb.fns = append(bdgb.fns, fns...)
	return bdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bdgb *BaselineDenomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bdgb.path(ctx)
	if err != nil {
		return err
	}
	bdgb.sql = query
	return bdgb.sqlScan(ctx, v)
}

func (bdgb *BaselineDenomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bdgb.fields {
		if !baselinedenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bdgb *BaselineDenomGroupBy) sqlQuery() *sql.Selector {
	selector := bdgb.sql.Select()
	aggregation := make([]string, 0, len(bdgb.fns))
	for _, fn := range bdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bdgb.fields)+len(bdgb.fns))
		for _, f := range bdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bdgb.fields...)...)
}

// BaselineDenomSelect is the builder for selecting fields of BaselineDenom entities.
type BaselineDenomSelect struct {
	*BaselineDenomQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bds *BaselineDenomSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bds.prepareQuery(ctx); err != nil {
		return err
	}
	bds.sql = bds.BaselineDenomQuery.sqlQuery(ctx)
	return bds.sqlScan(ctx, v)
}

func (bds *BaselineDenomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bds.sql.Query()
	if err := bds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
