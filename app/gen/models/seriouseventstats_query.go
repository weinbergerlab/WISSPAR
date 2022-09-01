// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventStatsQuery is the builder for querying SeriousEventStats entities.
type SeriousEventStatsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SeriousEventStats
	// eager-loading edges.
	withParent *SeriousEventQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SeriousEventStatsQuery builder.
func (sesq *SeriousEventStatsQuery) Where(ps ...predicate.SeriousEventStats) *SeriousEventStatsQuery {
	sesq.predicates = append(sesq.predicates, ps...)
	return sesq
}

// Limit adds a limit step to the query.
func (sesq *SeriousEventStatsQuery) Limit(limit int) *SeriousEventStatsQuery {
	sesq.limit = &limit
	return sesq
}

// Offset adds an offset step to the query.
func (sesq *SeriousEventStatsQuery) Offset(offset int) *SeriousEventStatsQuery {
	sesq.offset = &offset
	return sesq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sesq *SeriousEventStatsQuery) Unique(unique bool) *SeriousEventStatsQuery {
	sesq.unique = &unique
	return sesq
}

// Order adds an order step to the query.
func (sesq *SeriousEventStatsQuery) Order(o ...OrderFunc) *SeriousEventStatsQuery {
	sesq.order = append(sesq.order, o...)
	return sesq
}

// QueryParent chains the current query on the "parent" edge.
func (sesq *SeriousEventStatsQuery) QueryParent() *SeriousEventQuery {
	query := &SeriousEventQuery{config: sesq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sesq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sesq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(seriouseventstats.Table, seriouseventstats.FieldID, selector),
			sqlgraph.To(seriousevent.Table, seriousevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, seriouseventstats.ParentTable, seriouseventstats.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(sesq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SeriousEventStats entity from the query.
// Returns a *NotFoundError when no SeriousEventStats was found.
func (sesq *SeriousEventStatsQuery) First(ctx context.Context) (*SeriousEventStats, error) {
	nodes, err := sesq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{seriouseventstats.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) FirstX(ctx context.Context) *SeriousEventStats {
	node, err := sesq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SeriousEventStats ID from the query.
// Returns a *NotFoundError when no SeriousEventStats ID was found.
func (sesq *SeriousEventStatsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sesq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{seriouseventstats.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) FirstIDX(ctx context.Context) int {
	id, err := sesq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SeriousEventStats entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SeriousEventStats entity is found.
// Returns a *NotFoundError when no SeriousEventStats entities are found.
func (sesq *SeriousEventStatsQuery) Only(ctx context.Context) (*SeriousEventStats, error) {
	nodes, err := sesq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{seriouseventstats.Label}
	default:
		return nil, &NotSingularError{seriouseventstats.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) OnlyX(ctx context.Context) *SeriousEventStats {
	node, err := sesq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SeriousEventStats ID in the query.
// Returns a *NotSingularError when more than one SeriousEventStats ID is found.
// Returns a *NotFoundError when no entities are found.
func (sesq *SeriousEventStatsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sesq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{seriouseventstats.Label}
	default:
		err = &NotSingularError{seriouseventstats.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) OnlyIDX(ctx context.Context) int {
	id, err := sesq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SeriousEventStatsSlice.
func (sesq *SeriousEventStatsQuery) All(ctx context.Context) ([]*SeriousEventStats, error) {
	if err := sesq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sesq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) AllX(ctx context.Context) []*SeriousEventStats {
	nodes, err := sesq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SeriousEventStats IDs.
func (sesq *SeriousEventStatsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sesq.Select(seriouseventstats.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) IDsX(ctx context.Context) []int {
	ids, err := sesq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sesq *SeriousEventStatsQuery) Count(ctx context.Context) (int, error) {
	if err := sesq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sesq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) CountX(ctx context.Context) int {
	count, err := sesq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sesq *SeriousEventStatsQuery) Exist(ctx context.Context) (bool, error) {
	if err := sesq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sesq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sesq *SeriousEventStatsQuery) ExistX(ctx context.Context) bool {
	exist, err := sesq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SeriousEventStatsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sesq *SeriousEventStatsQuery) Clone() *SeriousEventStatsQuery {
	if sesq == nil {
		return nil
	}
	return &SeriousEventStatsQuery{
		config:     sesq.config,
		limit:      sesq.limit,
		offset:     sesq.offset,
		order:      append([]OrderFunc{}, sesq.order...),
		predicates: append([]predicate.SeriousEventStats{}, sesq.predicates...),
		withParent: sesq.withParent.Clone(),
		// clone intermediate query.
		sql:    sesq.sql.Clone(),
		path:   sesq.path,
		unique: sesq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (sesq *SeriousEventStatsQuery) WithParent(opts ...func(*SeriousEventQuery)) *SeriousEventStatsQuery {
	query := &SeriousEventQuery{config: sesq.config}
	for _, opt := range opts {
		opt(query)
	}
	sesq.withParent = query
	return sesq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SeriousEventStatsGroupID string `json:"serious_event_stats_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SeriousEventStats.Query().
//		GroupBy(seriouseventstats.FieldSeriousEventStatsGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (sesq *SeriousEventStatsQuery) GroupBy(field string, fields ...string) *SeriousEventStatsGroupBy {
	grbuild := &SeriousEventStatsGroupBy{config: sesq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sesq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sesq.sqlQuery(ctx), nil
	}
	grbuild.label = seriouseventstats.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SeriousEventStatsGroupID string `json:"serious_event_stats_group_id,omitempty"`
//	}
//
//	client.SeriousEventStats.Query().
//		Select(seriouseventstats.FieldSeriousEventStatsGroupID).
//		Scan(ctx, &v)
//
func (sesq *SeriousEventStatsQuery) Select(fields ...string) *SeriousEventStatsSelect {
	sesq.fields = append(sesq.fields, fields...)
	selbuild := &SeriousEventStatsSelect{SeriousEventStatsQuery: sesq}
	selbuild.label = seriouseventstats.Label
	selbuild.flds, selbuild.scan = &sesq.fields, selbuild.Scan
	return selbuild
}

func (sesq *SeriousEventStatsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sesq.fields {
		if !seriouseventstats.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if sesq.path != nil {
		prev, err := sesq.path(ctx)
		if err != nil {
			return err
		}
		sesq.sql = prev
	}
	return nil
}

func (sesq *SeriousEventStatsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SeriousEventStats, error) {
	var (
		nodes       = []*SeriousEventStats{}
		withFKs     = sesq.withFKs
		_spec       = sesq.querySpec()
		loadedTypes = [1]bool{
			sesq.withParent != nil,
		}
	)
	if sesq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, seriouseventstats.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SeriousEventStats).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SeriousEventStats{config: sesq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sesq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sesq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*SeriousEventStats)
		for i := range nodes {
			if nodes[i].serious_event_serious_event_stats_list == nil {
				continue
			}
			fk := *nodes[i].serious_event_serious_event_stats_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(seriousevent.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "serious_event_serious_event_stats_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (sesq *SeriousEventStatsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sesq.querySpec()
	_spec.Node.Columns = sesq.fields
	if len(sesq.fields) > 0 {
		_spec.Unique = sesq.unique != nil && *sesq.unique
	}
	return sqlgraph.CountNodes(ctx, sesq.driver, _spec)
}

func (sesq *SeriousEventStatsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sesq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (sesq *SeriousEventStatsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seriouseventstats.Table,
			Columns: seriouseventstats.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriouseventstats.FieldID,
			},
		},
		From:   sesq.sql,
		Unique: true,
	}
	if unique := sesq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sesq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seriouseventstats.FieldID)
		for i := range fields {
			if fields[i] != seriouseventstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sesq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sesq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sesq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sesq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sesq *SeriousEventStatsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sesq.driver.Dialect())
	t1 := builder.Table(seriouseventstats.Table)
	columns := sesq.fields
	if len(columns) == 0 {
		columns = seriouseventstats.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sesq.sql != nil {
		selector = sesq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sesq.unique != nil && *sesq.unique {
		selector.Distinct()
	}
	for _, p := range sesq.predicates {
		p(selector)
	}
	for _, p := range sesq.order {
		p(selector)
	}
	if offset := sesq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sesq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SeriousEventStatsGroupBy is the group-by builder for SeriousEventStats entities.
type SeriousEventStatsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sesgb *SeriousEventStatsGroupBy) Aggregate(fns ...AggregateFunc) *SeriousEventStatsGroupBy {
	sesgb.fns = append(sesgb.fns, fns...)
	return sesgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sesgb *SeriousEventStatsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sesgb.path(ctx)
	if err != nil {
		return err
	}
	sesgb.sql = query
	return sesgb.sqlScan(ctx, v)
}

func (sesgb *SeriousEventStatsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sesgb.fields {
		if !seriouseventstats.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sesgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sesgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sesgb *SeriousEventStatsGroupBy) sqlQuery() *sql.Selector {
	selector := sesgb.sql.Select()
	aggregation := make([]string, 0, len(sesgb.fns))
	for _, fn := range sesgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sesgb.fields)+len(sesgb.fns))
		for _, f := range sesgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sesgb.fields...)...)
}

// SeriousEventStatsSelect is the builder for selecting fields of SeriousEventStats entities.
type SeriousEventStatsSelect struct {
	*SeriousEventStatsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sess *SeriousEventStatsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sess.prepareQuery(ctx); err != nil {
		return err
	}
	sess.sql = sess.SeriousEventStatsQuery.sqlQuery(ctx)
	return sess.sqlScan(ctx, v)
}

func (sess *SeriousEventStatsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sess.sql.Query()
	if err := sess.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
