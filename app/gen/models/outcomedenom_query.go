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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeDenomQuery is the builder for querying OutcomeDenom entities.
type OutcomeDenomQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeDenom
	// eager-loading edges.
	withParent                *OutcomeMeasureQuery
	withOutcomeDenomCountList *OutcomeDenomCountQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeDenomQuery builder.
func (odq *OutcomeDenomQuery) Where(ps ...predicate.OutcomeDenom) *OutcomeDenomQuery {
	odq.predicates = append(odq.predicates, ps...)
	return odq
}

// Limit adds a limit step to the query.
func (odq *OutcomeDenomQuery) Limit(limit int) *OutcomeDenomQuery {
	odq.limit = &limit
	return odq
}

// Offset adds an offset step to the query.
func (odq *OutcomeDenomQuery) Offset(offset int) *OutcomeDenomQuery {
	odq.offset = &offset
	return odq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (odq *OutcomeDenomQuery) Unique(unique bool) *OutcomeDenomQuery {
	odq.unique = &unique
	return odq
}

// Order adds an order step to the query.
func (odq *OutcomeDenomQuery) Order(o ...OrderFunc) *OutcomeDenomQuery {
	odq.order = append(odq.order, o...)
	return odq
}

// QueryParent chains the current query on the "parent" edge.
func (odq *OutcomeDenomQuery) QueryParent() *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: odq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomedenom.Table, outcomedenom.FieldID, selector),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomedenom.ParentTable, outcomedenom.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(odq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeDenomCountList chains the current query on the "outcome_denom_count_list" edge.
func (odq *OutcomeDenomQuery) QueryOutcomeDenomCountList() *OutcomeDenomCountQuery {
	query := &OutcomeDenomCountQuery{config: odq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomedenom.Table, outcomedenom.FieldID, selector),
			sqlgraph.To(outcomedenomcount.Table, outcomedenomcount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomedenom.OutcomeDenomCountListTable, outcomedenom.OutcomeDenomCountListColumn),
		)
		fromU = sqlgraph.SetNeighbors(odq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeDenom entity from the query.
// Returns a *NotFoundError when no OutcomeDenom was found.
func (odq *OutcomeDenomQuery) First(ctx context.Context) (*OutcomeDenom, error) {
	nodes, err := odq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomedenom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (odq *OutcomeDenomQuery) FirstX(ctx context.Context) *OutcomeDenom {
	node, err := odq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeDenom ID from the query.
// Returns a *NotFoundError when no OutcomeDenom ID was found.
func (odq *OutcomeDenomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = odq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomedenom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (odq *OutcomeDenomQuery) FirstIDX(ctx context.Context) int {
	id, err := odq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeDenom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeDenom entity is found.
// Returns a *NotFoundError when no OutcomeDenom entities are found.
func (odq *OutcomeDenomQuery) Only(ctx context.Context) (*OutcomeDenom, error) {
	nodes, err := odq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomedenom.Label}
	default:
		return nil, &NotSingularError{outcomedenom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (odq *OutcomeDenomQuery) OnlyX(ctx context.Context) *OutcomeDenom {
	node, err := odq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeDenom ID in the query.
// Returns a *NotSingularError when more than one OutcomeDenom ID is found.
// Returns a *NotFoundError when no entities are found.
func (odq *OutcomeDenomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = odq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomedenom.Label}
	default:
		err = &NotSingularError{outcomedenom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (odq *OutcomeDenomQuery) OnlyIDX(ctx context.Context) int {
	id, err := odq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeDenoms.
func (odq *OutcomeDenomQuery) All(ctx context.Context) ([]*OutcomeDenom, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return odq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (odq *OutcomeDenomQuery) AllX(ctx context.Context) []*OutcomeDenom {
	nodes, err := odq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeDenom IDs.
func (odq *OutcomeDenomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := odq.Select(outcomedenom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (odq *OutcomeDenomQuery) IDsX(ctx context.Context) []int {
	ids, err := odq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (odq *OutcomeDenomQuery) Count(ctx context.Context) (int, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return odq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (odq *OutcomeDenomQuery) CountX(ctx context.Context) int {
	count, err := odq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (odq *OutcomeDenomQuery) Exist(ctx context.Context) (bool, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return odq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (odq *OutcomeDenomQuery) ExistX(ctx context.Context) bool {
	exist, err := odq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeDenomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (odq *OutcomeDenomQuery) Clone() *OutcomeDenomQuery {
	if odq == nil {
		return nil
	}
	return &OutcomeDenomQuery{
		config:                    odq.config,
		limit:                     odq.limit,
		offset:                    odq.offset,
		order:                     append([]OrderFunc{}, odq.order...),
		predicates:                append([]predicate.OutcomeDenom{}, odq.predicates...),
		withParent:                odq.withParent.Clone(),
		withOutcomeDenomCountList: odq.withOutcomeDenomCountList.Clone(),
		// clone intermediate query.
		sql:    odq.sql.Clone(),
		path:   odq.path,
		unique: odq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (odq *OutcomeDenomQuery) WithParent(opts ...func(*OutcomeMeasureQuery)) *OutcomeDenomQuery {
	query := &OutcomeMeasureQuery{config: odq.config}
	for _, opt := range opts {
		opt(query)
	}
	odq.withParent = query
	return odq
}

// WithOutcomeDenomCountList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_denom_count_list" edge. The optional arguments are used to configure the query builder of the edge.
func (odq *OutcomeDenomQuery) WithOutcomeDenomCountList(opts ...func(*OutcomeDenomCountQuery)) *OutcomeDenomQuery {
	query := &OutcomeDenomCountQuery{config: odq.config}
	for _, opt := range opts {
		opt(query)
	}
	odq.withOutcomeDenomCountList = query
	return odq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeDenomUnits string `json:"outcome_denom_units,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeDenom.Query().
//		GroupBy(outcomedenom.FieldOutcomeDenomUnits).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (odq *OutcomeDenomQuery) GroupBy(field string, fields ...string) *OutcomeDenomGroupBy {
	grbuild := &OutcomeDenomGroupBy{config: odq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return odq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomedenom.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeDenomUnits string `json:"outcome_denom_units,omitempty"`
//	}
//
//	client.OutcomeDenom.Query().
//		Select(outcomedenom.FieldOutcomeDenomUnits).
//		Scan(ctx, &v)
//
func (odq *OutcomeDenomQuery) Select(fields ...string) *OutcomeDenomSelect {
	odq.fields = append(odq.fields, fields...)
	selbuild := &OutcomeDenomSelect{OutcomeDenomQuery: odq}
	selbuild.label = outcomedenom.Label
	selbuild.flds, selbuild.scan = &odq.fields, selbuild.Scan
	return selbuild
}

func (odq *OutcomeDenomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range odq.fields {
		if !outcomedenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if odq.path != nil {
		prev, err := odq.path(ctx)
		if err != nil {
			return err
		}
		odq.sql = prev
	}
	return nil
}

func (odq *OutcomeDenomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeDenom, error) {
	var (
		nodes       = []*OutcomeDenom{}
		withFKs     = odq.withFKs
		_spec       = odq.querySpec()
		loadedTypes = [2]bool{
			odq.withParent != nil,
			odq.withOutcomeDenomCountList != nil,
		}
	)
	if odq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomedenom.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeDenom).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeDenom{config: odq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, odq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := odq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeDenom)
		for i := range nodes {
			if nodes[i].outcome_measure_outcome_denom_list == nil {
				continue
			}
			fk := *nodes[i].outcome_measure_outcome_denom_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomemeasure.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_denom_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := odq.withOutcomeDenomCountList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeDenom)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeDenomCountList = []*OutcomeDenomCount{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeDenomCount(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomedenom.OutcomeDenomCountListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_denom_outcome_denom_count_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_denom_outcome_denom_count_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_denom_outcome_denom_count_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeDenomCountList = append(node.Edges.OutcomeDenomCountList, n)
		}
	}

	return nodes, nil
}

func (odq *OutcomeDenomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := odq.querySpec()
	_spec.Node.Columns = odq.fields
	if len(odq.fields) > 0 {
		_spec.Unique = odq.unique != nil && *odq.unique
	}
	return sqlgraph.CountNodes(ctx, odq.driver, _spec)
}

func (odq *OutcomeDenomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := odq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (odq *OutcomeDenomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomedenom.Table,
			Columns: outcomedenom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenom.FieldID,
			},
		},
		From:   odq.sql,
		Unique: true,
	}
	if unique := odq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := odq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomedenom.FieldID)
		for i := range fields {
			if fields[i] != outcomedenom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := odq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := odq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := odq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := odq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (odq *OutcomeDenomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(odq.driver.Dialect())
	t1 := builder.Table(outcomedenom.Table)
	columns := odq.fields
	if len(columns) == 0 {
		columns = outcomedenom.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if odq.sql != nil {
		selector = odq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if odq.unique != nil && *odq.unique {
		selector.Distinct()
	}
	for _, p := range odq.predicates {
		p(selector)
	}
	for _, p := range odq.order {
		p(selector)
	}
	if offset := odq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := odq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeDenomGroupBy is the group-by builder for OutcomeDenom entities.
type OutcomeDenomGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (odgb *OutcomeDenomGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeDenomGroupBy {
	odgb.fns = append(odgb.fns, fns...)
	return odgb
}

// Scan applies the group-by query and scans the result into the given value.
func (odgb *OutcomeDenomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := odgb.path(ctx)
	if err != nil {
		return err
	}
	odgb.sql = query
	return odgb.sqlScan(ctx, v)
}

func (odgb *OutcomeDenomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range odgb.fields {
		if !outcomedenom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := odgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := odgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (odgb *OutcomeDenomGroupBy) sqlQuery() *sql.Selector {
	selector := odgb.sql.Select()
	aggregation := make([]string, 0, len(odgb.fns))
	for _, fn := range odgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(odgb.fields)+len(odgb.fns))
		for _, f := range odgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(odgb.fields...)...)
}

// OutcomeDenomSelect is the builder for selecting fields of OutcomeDenom entities.
type OutcomeDenomSelect struct {
	*OutcomeDenomQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ods *OutcomeDenomSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ods.prepareQuery(ctx); err != nil {
		return err
	}
	ods.sql = ods.OutcomeDenomQuery.sqlQuery(ctx)
	return ods.sqlScan(ctx, v)
}

func (ods *OutcomeDenomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ods.sql.Query()
	if err := ods.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
