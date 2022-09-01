// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OtherEventStatsQuery is the builder for querying OtherEventStats entities.
type OtherEventStatsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OtherEventStats
	// eager-loading edges.
	withParent *OtherEventQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OtherEventStatsQuery builder.
func (oesq *OtherEventStatsQuery) Where(ps ...predicate.OtherEventStats) *OtherEventStatsQuery {
	oesq.predicates = append(oesq.predicates, ps...)
	return oesq
}

// Limit adds a limit step to the query.
func (oesq *OtherEventStatsQuery) Limit(limit int) *OtherEventStatsQuery {
	oesq.limit = &limit
	return oesq
}

// Offset adds an offset step to the query.
func (oesq *OtherEventStatsQuery) Offset(offset int) *OtherEventStatsQuery {
	oesq.offset = &offset
	return oesq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oesq *OtherEventStatsQuery) Unique(unique bool) *OtherEventStatsQuery {
	oesq.unique = &unique
	return oesq
}

// Order adds an order step to the query.
func (oesq *OtherEventStatsQuery) Order(o ...OrderFunc) *OtherEventStatsQuery {
	oesq.order = append(oesq.order, o...)
	return oesq
}

// QueryParent chains the current query on the "parent" edge.
func (oesq *OtherEventStatsQuery) QueryParent() *OtherEventQuery {
	query := &OtherEventQuery{config: oesq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oesq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oesq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(othereventstats.Table, othereventstats.FieldID, selector),
			sqlgraph.To(otherevent.Table, otherevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, othereventstats.ParentTable, othereventstats.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(oesq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OtherEventStats entity from the query.
// Returns a *NotFoundError when no OtherEventStats was found.
func (oesq *OtherEventStatsQuery) First(ctx context.Context) (*OtherEventStats, error) {
	nodes, err := oesq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{othereventstats.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) FirstX(ctx context.Context) *OtherEventStats {
	node, err := oesq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OtherEventStats ID from the query.
// Returns a *NotFoundError when no OtherEventStats ID was found.
func (oesq *OtherEventStatsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oesq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{othereventstats.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) FirstIDX(ctx context.Context) int {
	id, err := oesq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OtherEventStats entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OtherEventStats entity is found.
// Returns a *NotFoundError when no OtherEventStats entities are found.
func (oesq *OtherEventStatsQuery) Only(ctx context.Context) (*OtherEventStats, error) {
	nodes, err := oesq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{othereventstats.Label}
	default:
		return nil, &NotSingularError{othereventstats.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) OnlyX(ctx context.Context) *OtherEventStats {
	node, err := oesq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OtherEventStats ID in the query.
// Returns a *NotSingularError when more than one OtherEventStats ID is found.
// Returns a *NotFoundError when no entities are found.
func (oesq *OtherEventStatsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oesq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{othereventstats.Label}
	default:
		err = &NotSingularError{othereventstats.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) OnlyIDX(ctx context.Context) int {
	id, err := oesq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OtherEventStatsSlice.
func (oesq *OtherEventStatsQuery) All(ctx context.Context) ([]*OtherEventStats, error) {
	if err := oesq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oesq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) AllX(ctx context.Context) []*OtherEventStats {
	nodes, err := oesq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OtherEventStats IDs.
func (oesq *OtherEventStatsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oesq.Select(othereventstats.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) IDsX(ctx context.Context) []int {
	ids, err := oesq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oesq *OtherEventStatsQuery) Count(ctx context.Context) (int, error) {
	if err := oesq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oesq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) CountX(ctx context.Context) int {
	count, err := oesq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oesq *OtherEventStatsQuery) Exist(ctx context.Context) (bool, error) {
	if err := oesq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oesq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oesq *OtherEventStatsQuery) ExistX(ctx context.Context) bool {
	exist, err := oesq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OtherEventStatsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oesq *OtherEventStatsQuery) Clone() *OtherEventStatsQuery {
	if oesq == nil {
		return nil
	}
	return &OtherEventStatsQuery{
		config:     oesq.config,
		limit:      oesq.limit,
		offset:     oesq.offset,
		order:      append([]OrderFunc{}, oesq.order...),
		predicates: append([]predicate.OtherEventStats{}, oesq.predicates...),
		withParent: oesq.withParent.Clone(),
		// clone intermediate query.
		sql:    oesq.sql.Clone(),
		path:   oesq.path,
		unique: oesq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (oesq *OtherEventStatsQuery) WithParent(opts ...func(*OtherEventQuery)) *OtherEventStatsQuery {
	query := &OtherEventQuery{config: oesq.config}
	for _, opt := range opts {
		opt(query)
	}
	oesq.withParent = query
	return oesq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OtherEventStatsGroupID string `json:"other_event_stats_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OtherEventStats.Query().
//		GroupBy(othereventstats.FieldOtherEventStatsGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (oesq *OtherEventStatsQuery) GroupBy(field string, fields ...string) *OtherEventStatsGroupBy {
	grbuild := &OtherEventStatsGroupBy{config: oesq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oesq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oesq.sqlQuery(ctx), nil
	}
	grbuild.label = othereventstats.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OtherEventStatsGroupID string `json:"other_event_stats_group_id,omitempty"`
//	}
//
//	client.OtherEventStats.Query().
//		Select(othereventstats.FieldOtherEventStatsGroupID).
//		Scan(ctx, &v)
//
func (oesq *OtherEventStatsQuery) Select(fields ...string) *OtherEventStatsSelect {
	oesq.fields = append(oesq.fields, fields...)
	selbuild := &OtherEventStatsSelect{OtherEventStatsQuery: oesq}
	selbuild.label = othereventstats.Label
	selbuild.flds, selbuild.scan = &oesq.fields, selbuild.Scan
	return selbuild
}

func (oesq *OtherEventStatsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oesq.fields {
		if !othereventstats.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if oesq.path != nil {
		prev, err := oesq.path(ctx)
		if err != nil {
			return err
		}
		oesq.sql = prev
	}
	return nil
}

func (oesq *OtherEventStatsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OtherEventStats, error) {
	var (
		nodes       = []*OtherEventStats{}
		withFKs     = oesq.withFKs
		_spec       = oesq.querySpec()
		loadedTypes = [1]bool{
			oesq.withParent != nil,
		}
	)
	if oesq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, othereventstats.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OtherEventStats).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OtherEventStats{config: oesq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oesq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oesq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OtherEventStats)
		for i := range nodes {
			if nodes[i].other_event_other_event_stats_list == nil {
				continue
			}
			fk := *nodes[i].other_event_other_event_stats_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(otherevent.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "other_event_other_event_stats_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (oesq *OtherEventStatsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oesq.querySpec()
	_spec.Node.Columns = oesq.fields
	if len(oesq.fields) > 0 {
		_spec.Unique = oesq.unique != nil && *oesq.unique
	}
	return sqlgraph.CountNodes(ctx, oesq.driver, _spec)
}

func (oesq *OtherEventStatsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oesq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (oesq *OtherEventStatsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   othereventstats.Table,
			Columns: othereventstats.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: othereventstats.FieldID,
			},
		},
		From:   oesq.sql,
		Unique: true,
	}
	if unique := oesq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oesq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, othereventstats.FieldID)
		for i := range fields {
			if fields[i] != othereventstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oesq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oesq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oesq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oesq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oesq *OtherEventStatsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oesq.driver.Dialect())
	t1 := builder.Table(othereventstats.Table)
	columns := oesq.fields
	if len(columns) == 0 {
		columns = othereventstats.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oesq.sql != nil {
		selector = oesq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oesq.unique != nil && *oesq.unique {
		selector.Distinct()
	}
	for _, p := range oesq.predicates {
		p(selector)
	}
	for _, p := range oesq.order {
		p(selector)
	}
	if offset := oesq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oesq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OtherEventStatsGroupBy is the group-by builder for OtherEventStats entities.
type OtherEventStatsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oesgb *OtherEventStatsGroupBy) Aggregate(fns ...AggregateFunc) *OtherEventStatsGroupBy {
	oesgb.fns = append(oesgb.fns, fns...)
	return oesgb
}

// Scan applies the group-by query and scans the result into the given value.
func (oesgb *OtherEventStatsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oesgb.path(ctx)
	if err != nil {
		return err
	}
	oesgb.sql = query
	return oesgb.sqlScan(ctx, v)
}

func (oesgb *OtherEventStatsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oesgb.fields {
		if !othereventstats.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oesgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oesgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oesgb *OtherEventStatsGroupBy) sqlQuery() *sql.Selector {
	selector := oesgb.sql.Select()
	aggregation := make([]string, 0, len(oesgb.fns))
	for _, fn := range oesgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oesgb.fields)+len(oesgb.fns))
		for _, f := range oesgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oesgb.fields...)...)
}

// OtherEventStatsSelect is the builder for selecting fields of OtherEventStats entities.
type OtherEventStatsSelect struct {
	*OtherEventStatsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oess *OtherEventStatsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oess.prepareQuery(ctx); err != nil {
		return err
	}
	oess.sql = oess.OtherEventStatsQuery.sqlQuery(ctx)
	return oess.sqlScan(ctx, v)
}

func (oess *OtherEventStatsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oess.sql.Query()
	if err := oess.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
