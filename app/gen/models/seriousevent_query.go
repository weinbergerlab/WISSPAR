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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriouseventstats"
)

// SeriousEventQuery is the builder for querying SeriousEvent entities.
type SeriousEventQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SeriousEvent
	// eager-loading edges.
	withParent                *AdverseEventsModuleQuery
	withSeriousEventStatsList *SeriousEventStatsQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SeriousEventQuery builder.
func (seq *SeriousEventQuery) Where(ps ...predicate.SeriousEvent) *SeriousEventQuery {
	seq.predicates = append(seq.predicates, ps...)
	return seq
}

// Limit adds a limit step to the query.
func (seq *SeriousEventQuery) Limit(limit int) *SeriousEventQuery {
	seq.limit = &limit
	return seq
}

// Offset adds an offset step to the query.
func (seq *SeriousEventQuery) Offset(offset int) *SeriousEventQuery {
	seq.offset = &offset
	return seq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (seq *SeriousEventQuery) Unique(unique bool) *SeriousEventQuery {
	seq.unique = &unique
	return seq
}

// Order adds an order step to the query.
func (seq *SeriousEventQuery) Order(o ...OrderFunc) *SeriousEventQuery {
	seq.order = append(seq.order, o...)
	return seq
}

// QueryParent chains the current query on the "parent" edge.
func (seq *SeriousEventQuery) QueryParent() *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: seq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := seq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(seriousevent.Table, seriousevent.FieldID, selector),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, seriousevent.ParentTable, seriousevent.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(seq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySeriousEventStatsList chains the current query on the "serious_event_stats_list" edge.
func (seq *SeriousEventQuery) QuerySeriousEventStatsList() *SeriousEventStatsQuery {
	query := &SeriousEventStatsQuery{config: seq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := seq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(seriousevent.Table, seriousevent.FieldID, selector),
			sqlgraph.To(seriouseventstats.Table, seriouseventstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, seriousevent.SeriousEventStatsListTable, seriousevent.SeriousEventStatsListColumn),
		)
		fromU = sqlgraph.SetNeighbors(seq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SeriousEvent entity from the query.
// Returns a *NotFoundError when no SeriousEvent was found.
func (seq *SeriousEventQuery) First(ctx context.Context) (*SeriousEvent, error) {
	nodes, err := seq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{seriousevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (seq *SeriousEventQuery) FirstX(ctx context.Context) *SeriousEvent {
	node, err := seq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SeriousEvent ID from the query.
// Returns a *NotFoundError when no SeriousEvent ID was found.
func (seq *SeriousEventQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{seriousevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (seq *SeriousEventQuery) FirstIDX(ctx context.Context) int {
	id, err := seq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SeriousEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SeriousEvent entity is found.
// Returns a *NotFoundError when no SeriousEvent entities are found.
func (seq *SeriousEventQuery) Only(ctx context.Context) (*SeriousEvent, error) {
	nodes, err := seq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{seriousevent.Label}
	default:
		return nil, &NotSingularError{seriousevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (seq *SeriousEventQuery) OnlyX(ctx context.Context) *SeriousEvent {
	node, err := seq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SeriousEvent ID in the query.
// Returns a *NotSingularError when more than one SeriousEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (seq *SeriousEventQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{seriousevent.Label}
	default:
		err = &NotSingularError{seriousevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (seq *SeriousEventQuery) OnlyIDX(ctx context.Context) int {
	id, err := seq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SeriousEvents.
func (seq *SeriousEventQuery) All(ctx context.Context) ([]*SeriousEvent, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return seq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (seq *SeriousEventQuery) AllX(ctx context.Context) []*SeriousEvent {
	nodes, err := seq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SeriousEvent IDs.
func (seq *SeriousEventQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := seq.Select(seriousevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (seq *SeriousEventQuery) IDsX(ctx context.Context) []int {
	ids, err := seq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (seq *SeriousEventQuery) Count(ctx context.Context) (int, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return seq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (seq *SeriousEventQuery) CountX(ctx context.Context) int {
	count, err := seq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (seq *SeriousEventQuery) Exist(ctx context.Context) (bool, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return seq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (seq *SeriousEventQuery) ExistX(ctx context.Context) bool {
	exist, err := seq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SeriousEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (seq *SeriousEventQuery) Clone() *SeriousEventQuery {
	if seq == nil {
		return nil
	}
	return &SeriousEventQuery{
		config:                    seq.config,
		limit:                     seq.limit,
		offset:                    seq.offset,
		order:                     append([]OrderFunc{}, seq.order...),
		predicates:                append([]predicate.SeriousEvent{}, seq.predicates...),
		withParent:                seq.withParent.Clone(),
		withSeriousEventStatsList: seq.withSeriousEventStatsList.Clone(),
		// clone intermediate query.
		sql:    seq.sql.Clone(),
		path:   seq.path,
		unique: seq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (seq *SeriousEventQuery) WithParent(opts ...func(*AdverseEventsModuleQuery)) *SeriousEventQuery {
	query := &AdverseEventsModuleQuery{config: seq.config}
	for _, opt := range opts {
		opt(query)
	}
	seq.withParent = query
	return seq
}

// WithSeriousEventStatsList tells the query-builder to eager-load the nodes that are connected to
// the "serious_event_stats_list" edge. The optional arguments are used to configure the query builder of the edge.
func (seq *SeriousEventQuery) WithSeriousEventStatsList(opts ...func(*SeriousEventStatsQuery)) *SeriousEventQuery {
	query := &SeriousEventStatsQuery{config: seq.config}
	for _, opt := range opts {
		opt(query)
	}
	seq.withSeriousEventStatsList = query
	return seq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SeriousEventTerm string `json:"serious_event_term,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SeriousEvent.Query().
//		GroupBy(seriousevent.FieldSeriousEventTerm).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (seq *SeriousEventQuery) GroupBy(field string, fields ...string) *SeriousEventGroupBy {
	grbuild := &SeriousEventGroupBy{config: seq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return seq.sqlQuery(ctx), nil
	}
	grbuild.label = seriousevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SeriousEventTerm string `json:"serious_event_term,omitempty"`
//	}
//
//	client.SeriousEvent.Query().
//		Select(seriousevent.FieldSeriousEventTerm).
//		Scan(ctx, &v)
//
func (seq *SeriousEventQuery) Select(fields ...string) *SeriousEventSelect {
	seq.fields = append(seq.fields, fields...)
	selbuild := &SeriousEventSelect{SeriousEventQuery: seq}
	selbuild.label = seriousevent.Label
	selbuild.flds, selbuild.scan = &seq.fields, selbuild.Scan
	return selbuild
}

func (seq *SeriousEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range seq.fields {
		if !seriousevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if seq.path != nil {
		prev, err := seq.path(ctx)
		if err != nil {
			return err
		}
		seq.sql = prev
	}
	return nil
}

func (seq *SeriousEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SeriousEvent, error) {
	var (
		nodes       = []*SeriousEvent{}
		withFKs     = seq.withFKs
		_spec       = seq.querySpec()
		loadedTypes = [2]bool{
			seq.withParent != nil,
			seq.withSeriousEventStatsList != nil,
		}
	)
	if seq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, seriousevent.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SeriousEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SeriousEvent{config: seq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, seq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := seq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*SeriousEvent)
		for i := range nodes {
			if nodes[i].adverse_events_module_serious_event_list == nil {
				continue
			}
			fk := *nodes[i].adverse_events_module_serious_event_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(adverseeventsmodule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "adverse_events_module_serious_event_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := seq.withSeriousEventStatsList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*SeriousEvent)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.SeriousEventStatsList = []*SeriousEventStats{}
		}
		query.withFKs = true
		query.Where(predicate.SeriousEventStats(func(s *sql.Selector) {
			s.Where(sql.InValues(seriousevent.SeriousEventStatsListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.serious_event_serious_event_stats_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "serious_event_serious_event_stats_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "serious_event_serious_event_stats_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.SeriousEventStatsList = append(node.Edges.SeriousEventStatsList, n)
		}
	}

	return nodes, nil
}

func (seq *SeriousEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := seq.querySpec()
	_spec.Node.Columns = seq.fields
	if len(seq.fields) > 0 {
		_spec.Unique = seq.unique != nil && *seq.unique
	}
	return sqlgraph.CountNodes(ctx, seq.driver, _spec)
}

func (seq *SeriousEventQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := seq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (seq *SeriousEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seriousevent.Table,
			Columns: seriousevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: seriousevent.FieldID,
			},
		},
		From:   seq.sql,
		Unique: true,
	}
	if unique := seq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := seq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seriousevent.FieldID)
		for i := range fields {
			if fields[i] != seriousevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := seq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := seq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := seq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := seq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (seq *SeriousEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(seq.driver.Dialect())
	t1 := builder.Table(seriousevent.Table)
	columns := seq.fields
	if len(columns) == 0 {
		columns = seriousevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if seq.sql != nil {
		selector = seq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if seq.unique != nil && *seq.unique {
		selector.Distinct()
	}
	for _, p := range seq.predicates {
		p(selector)
	}
	for _, p := range seq.order {
		p(selector)
	}
	if offset := seq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := seq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SeriousEventGroupBy is the group-by builder for SeriousEvent entities.
type SeriousEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (segb *SeriousEventGroupBy) Aggregate(fns ...AggregateFunc) *SeriousEventGroupBy {
	segb.fns = append(segb.fns, fns...)
	return segb
}

// Scan applies the group-by query and scans the result into the given value.
func (segb *SeriousEventGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := segb.path(ctx)
	if err != nil {
		return err
	}
	segb.sql = query
	return segb.sqlScan(ctx, v)
}

func (segb *SeriousEventGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range segb.fields {
		if !seriousevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := segb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := segb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (segb *SeriousEventGroupBy) sqlQuery() *sql.Selector {
	selector := segb.sql.Select()
	aggregation := make([]string, 0, len(segb.fns))
	for _, fn := range segb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(segb.fields)+len(segb.fns))
		for _, f := range segb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(segb.fields...)...)
}

// SeriousEventSelect is the builder for selecting fields of SeriousEvent entities.
type SeriousEventSelect struct {
	*SeriousEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ses *SeriousEventSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ses.prepareQuery(ctx); err != nil {
		return err
	}
	ses.sql = ses.SeriousEventQuery.sqlQuery(ctx)
	return ses.sqlScan(ctx, v)
}

func (ses *SeriousEventSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ses.sql.Query()
	if err := ses.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
