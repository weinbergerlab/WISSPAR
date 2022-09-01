// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenomcount"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeDenomCountQuery is the builder for querying OutcomeDenomCount entities.
type OutcomeDenomCountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeDenomCount
	// eager-loading edges.
	withParent *OutcomeDenomQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeDenomCountQuery builder.
func (odcq *OutcomeDenomCountQuery) Where(ps ...predicate.OutcomeDenomCount) *OutcomeDenomCountQuery {
	odcq.predicates = append(odcq.predicates, ps...)
	return odcq
}

// Limit adds a limit step to the query.
func (odcq *OutcomeDenomCountQuery) Limit(limit int) *OutcomeDenomCountQuery {
	odcq.limit = &limit
	return odcq
}

// Offset adds an offset step to the query.
func (odcq *OutcomeDenomCountQuery) Offset(offset int) *OutcomeDenomCountQuery {
	odcq.offset = &offset
	return odcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (odcq *OutcomeDenomCountQuery) Unique(unique bool) *OutcomeDenomCountQuery {
	odcq.unique = &unique
	return odcq
}

// Order adds an order step to the query.
func (odcq *OutcomeDenomCountQuery) Order(o ...OrderFunc) *OutcomeDenomCountQuery {
	odcq.order = append(odcq.order, o...)
	return odcq
}

// QueryParent chains the current query on the "parent" edge.
func (odcq *OutcomeDenomCountQuery) QueryParent() *OutcomeDenomQuery {
	query := &OutcomeDenomQuery{config: odcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomedenomcount.Table, outcomedenomcount.FieldID, selector),
			sqlgraph.To(outcomedenom.Table, outcomedenom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomedenomcount.ParentTable, outcomedenomcount.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(odcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeDenomCount entity from the query.
// Returns a *NotFoundError when no OutcomeDenomCount was found.
func (odcq *OutcomeDenomCountQuery) First(ctx context.Context) (*OutcomeDenomCount, error) {
	nodes, err := odcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomedenomcount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) FirstX(ctx context.Context) *OutcomeDenomCount {
	node, err := odcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeDenomCount ID from the query.
// Returns a *NotFoundError when no OutcomeDenomCount ID was found.
func (odcq *OutcomeDenomCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = odcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomedenomcount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) FirstIDX(ctx context.Context) int {
	id, err := odcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeDenomCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeDenomCount entity is found.
// Returns a *NotFoundError when no OutcomeDenomCount entities are found.
func (odcq *OutcomeDenomCountQuery) Only(ctx context.Context) (*OutcomeDenomCount, error) {
	nodes, err := odcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomedenomcount.Label}
	default:
		return nil, &NotSingularError{outcomedenomcount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) OnlyX(ctx context.Context) *OutcomeDenomCount {
	node, err := odcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeDenomCount ID in the query.
// Returns a *NotSingularError when more than one OutcomeDenomCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (odcq *OutcomeDenomCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = odcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomedenomcount.Label}
	default:
		err = &NotSingularError{outcomedenomcount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := odcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeDenomCounts.
func (odcq *OutcomeDenomCountQuery) All(ctx context.Context) ([]*OutcomeDenomCount, error) {
	if err := odcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return odcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) AllX(ctx context.Context) []*OutcomeDenomCount {
	nodes, err := odcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeDenomCount IDs.
func (odcq *OutcomeDenomCountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := odcq.Select(outcomedenomcount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) IDsX(ctx context.Context) []int {
	ids, err := odcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (odcq *OutcomeDenomCountQuery) Count(ctx context.Context) (int, error) {
	if err := odcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return odcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) CountX(ctx context.Context) int {
	count, err := odcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (odcq *OutcomeDenomCountQuery) Exist(ctx context.Context) (bool, error) {
	if err := odcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return odcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (odcq *OutcomeDenomCountQuery) ExistX(ctx context.Context) bool {
	exist, err := odcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeDenomCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (odcq *OutcomeDenomCountQuery) Clone() *OutcomeDenomCountQuery {
	if odcq == nil {
		return nil
	}
	return &OutcomeDenomCountQuery{
		config:     odcq.config,
		limit:      odcq.limit,
		offset:     odcq.offset,
		order:      append([]OrderFunc{}, odcq.order...),
		predicates: append([]predicate.OutcomeDenomCount{}, odcq.predicates...),
		withParent: odcq.withParent.Clone(),
		// clone intermediate query.
		sql:    odcq.sql.Clone(),
		path:   odcq.path,
		unique: odcq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (odcq *OutcomeDenomCountQuery) WithParent(opts ...func(*OutcomeDenomQuery)) *OutcomeDenomCountQuery {
	query := &OutcomeDenomQuery{config: odcq.config}
	for _, opt := range opts {
		opt(query)
	}
	odcq.withParent = query
	return odcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeDenomCountGroupID string `json:"outcome_denom_count_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeDenomCount.Query().
//		GroupBy(outcomedenomcount.FieldOutcomeDenomCountGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (odcq *OutcomeDenomCountQuery) GroupBy(field string, fields ...string) *OutcomeDenomCountGroupBy {
	grbuild := &OutcomeDenomCountGroupBy{config: odcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := odcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return odcq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomedenomcount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeDenomCountGroupID string `json:"outcome_denom_count_group_id,omitempty"`
//	}
//
//	client.OutcomeDenomCount.Query().
//		Select(outcomedenomcount.FieldOutcomeDenomCountGroupID).
//		Scan(ctx, &v)
//
func (odcq *OutcomeDenomCountQuery) Select(fields ...string) *OutcomeDenomCountSelect {
	odcq.fields = append(odcq.fields, fields...)
	selbuild := &OutcomeDenomCountSelect{OutcomeDenomCountQuery: odcq}
	selbuild.label = outcomedenomcount.Label
	selbuild.flds, selbuild.scan = &odcq.fields, selbuild.Scan
	return selbuild
}

func (odcq *OutcomeDenomCountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range odcq.fields {
		if !outcomedenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if odcq.path != nil {
		prev, err := odcq.path(ctx)
		if err != nil {
			return err
		}
		odcq.sql = prev
	}
	return nil
}

func (odcq *OutcomeDenomCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeDenomCount, error) {
	var (
		nodes       = []*OutcomeDenomCount{}
		withFKs     = odcq.withFKs
		_spec       = odcq.querySpec()
		loadedTypes = [1]bool{
			odcq.withParent != nil,
		}
	)
	if odcq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomedenomcount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeDenomCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeDenomCount{config: odcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, odcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := odcq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeDenomCount)
		for i := range nodes {
			if nodes[i].outcome_denom_outcome_denom_count_list == nil {
				continue
			}
			fk := *nodes[i].outcome_denom_outcome_denom_count_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomedenom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_denom_outcome_denom_count_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (odcq *OutcomeDenomCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := odcq.querySpec()
	_spec.Node.Columns = odcq.fields
	if len(odcq.fields) > 0 {
		_spec.Unique = odcq.unique != nil && *odcq.unique
	}
	return sqlgraph.CountNodes(ctx, odcq.driver, _spec)
}

func (odcq *OutcomeDenomCountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := odcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (odcq *OutcomeDenomCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomedenomcount.Table,
			Columns: outcomedenomcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomedenomcount.FieldID,
			},
		},
		From:   odcq.sql,
		Unique: true,
	}
	if unique := odcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := odcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomedenomcount.FieldID)
		for i := range fields {
			if fields[i] != outcomedenomcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := odcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := odcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := odcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := odcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (odcq *OutcomeDenomCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(odcq.driver.Dialect())
	t1 := builder.Table(outcomedenomcount.Table)
	columns := odcq.fields
	if len(columns) == 0 {
		columns = outcomedenomcount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if odcq.sql != nil {
		selector = odcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if odcq.unique != nil && *odcq.unique {
		selector.Distinct()
	}
	for _, p := range odcq.predicates {
		p(selector)
	}
	for _, p := range odcq.order {
		p(selector)
	}
	if offset := odcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := odcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeDenomCountGroupBy is the group-by builder for OutcomeDenomCount entities.
type OutcomeDenomCountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (odcgb *OutcomeDenomCountGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeDenomCountGroupBy {
	odcgb.fns = append(odcgb.fns, fns...)
	return odcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (odcgb *OutcomeDenomCountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := odcgb.path(ctx)
	if err != nil {
		return err
	}
	odcgb.sql = query
	return odcgb.sqlScan(ctx, v)
}

func (odcgb *OutcomeDenomCountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range odcgb.fields {
		if !outcomedenomcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := odcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := odcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (odcgb *OutcomeDenomCountGroupBy) sqlQuery() *sql.Selector {
	selector := odcgb.sql.Select()
	aggregation := make([]string, 0, len(odcgb.fns))
	for _, fn := range odcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(odcgb.fields)+len(odcgb.fns))
		for _, f := range odcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(odcgb.fields...)...)
}

// OutcomeDenomCountSelect is the builder for selecting fields of OutcomeDenomCount entities.
type OutcomeDenomCountSelect struct {
	*OutcomeDenomCountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (odcs *OutcomeDenomCountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := odcs.prepareQuery(ctx); err != nil {
		return err
	}
	odcs.sql = odcs.OutcomeDenomCountQuery.sqlQuery(ctx)
	return odcs.sqlScan(ctx, v)
}

func (odcs *OutcomeDenomCountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := odcs.sql.Query()
	if err := odcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
