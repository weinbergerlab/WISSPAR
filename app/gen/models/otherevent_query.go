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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/othereventstats"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OtherEventQuery is the builder for querying OtherEvent entities.
type OtherEventQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OtherEvent
	// eager-loading edges.
	withParent              *AdverseEventsModuleQuery
	withOtherEventStatsList *OtherEventStatsQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OtherEventQuery builder.
func (oeq *OtherEventQuery) Where(ps ...predicate.OtherEvent) *OtherEventQuery {
	oeq.predicates = append(oeq.predicates, ps...)
	return oeq
}

// Limit adds a limit step to the query.
func (oeq *OtherEventQuery) Limit(limit int) *OtherEventQuery {
	oeq.limit = &limit
	return oeq
}

// Offset adds an offset step to the query.
func (oeq *OtherEventQuery) Offset(offset int) *OtherEventQuery {
	oeq.offset = &offset
	return oeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oeq *OtherEventQuery) Unique(unique bool) *OtherEventQuery {
	oeq.unique = &unique
	return oeq
}

// Order adds an order step to the query.
func (oeq *OtherEventQuery) Order(o ...OrderFunc) *OtherEventQuery {
	oeq.order = append(oeq.order, o...)
	return oeq
}

// QueryParent chains the current query on the "parent" edge.
func (oeq *OtherEventQuery) QueryParent() *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: oeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(otherevent.Table, otherevent.FieldID, selector),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, otherevent.ParentTable, otherevent.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(oeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOtherEventStatsList chains the current query on the "other_event_stats_list" edge.
func (oeq *OtherEventQuery) QueryOtherEventStatsList() *OtherEventStatsQuery {
	query := &OtherEventStatsQuery{config: oeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(otherevent.Table, otherevent.FieldID, selector),
			sqlgraph.To(othereventstats.Table, othereventstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, otherevent.OtherEventStatsListTable, otherevent.OtherEventStatsListColumn),
		)
		fromU = sqlgraph.SetNeighbors(oeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OtherEvent entity from the query.
// Returns a *NotFoundError when no OtherEvent was found.
func (oeq *OtherEventQuery) First(ctx context.Context) (*OtherEvent, error) {
	nodes, err := oeq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{otherevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oeq *OtherEventQuery) FirstX(ctx context.Context) *OtherEvent {
	node, err := oeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OtherEvent ID from the query.
// Returns a *NotFoundError when no OtherEvent ID was found.
func (oeq *OtherEventQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oeq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{otherevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oeq *OtherEventQuery) FirstIDX(ctx context.Context) int {
	id, err := oeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OtherEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OtherEvent entity is found.
// Returns a *NotFoundError when no OtherEvent entities are found.
func (oeq *OtherEventQuery) Only(ctx context.Context) (*OtherEvent, error) {
	nodes, err := oeq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{otherevent.Label}
	default:
		return nil, &NotSingularError{otherevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oeq *OtherEventQuery) OnlyX(ctx context.Context) *OtherEvent {
	node, err := oeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OtherEvent ID in the query.
// Returns a *NotSingularError when more than one OtherEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (oeq *OtherEventQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oeq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{otherevent.Label}
	default:
		err = &NotSingularError{otherevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oeq *OtherEventQuery) OnlyIDX(ctx context.Context) int {
	id, err := oeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OtherEvents.
func (oeq *OtherEventQuery) All(ctx context.Context) ([]*OtherEvent, error) {
	if err := oeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oeq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oeq *OtherEventQuery) AllX(ctx context.Context) []*OtherEvent {
	nodes, err := oeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OtherEvent IDs.
func (oeq *OtherEventQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oeq.Select(otherevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oeq *OtherEventQuery) IDsX(ctx context.Context) []int {
	ids, err := oeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oeq *OtherEventQuery) Count(ctx context.Context) (int, error) {
	if err := oeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oeq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oeq *OtherEventQuery) CountX(ctx context.Context) int {
	count, err := oeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oeq *OtherEventQuery) Exist(ctx context.Context) (bool, error) {
	if err := oeq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oeq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oeq *OtherEventQuery) ExistX(ctx context.Context) bool {
	exist, err := oeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OtherEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oeq *OtherEventQuery) Clone() *OtherEventQuery {
	if oeq == nil {
		return nil
	}
	return &OtherEventQuery{
		config:                  oeq.config,
		limit:                   oeq.limit,
		offset:                  oeq.offset,
		order:                   append([]OrderFunc{}, oeq.order...),
		predicates:              append([]predicate.OtherEvent{}, oeq.predicates...),
		withParent:              oeq.withParent.Clone(),
		withOtherEventStatsList: oeq.withOtherEventStatsList.Clone(),
		// clone intermediate query.
		sql:    oeq.sql.Clone(),
		path:   oeq.path,
		unique: oeq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (oeq *OtherEventQuery) WithParent(opts ...func(*AdverseEventsModuleQuery)) *OtherEventQuery {
	query := &AdverseEventsModuleQuery{config: oeq.config}
	for _, opt := range opts {
		opt(query)
	}
	oeq.withParent = query
	return oeq
}

// WithOtherEventStatsList tells the query-builder to eager-load the nodes that are connected to
// the "other_event_stats_list" edge. The optional arguments are used to configure the query builder of the edge.
func (oeq *OtherEventQuery) WithOtherEventStatsList(opts ...func(*OtherEventStatsQuery)) *OtherEventQuery {
	query := &OtherEventStatsQuery{config: oeq.config}
	for _, opt := range opts {
		opt(query)
	}
	oeq.withOtherEventStatsList = query
	return oeq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OtherEventTerm string `json:"other_event_term,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OtherEvent.Query().
//		GroupBy(otherevent.FieldOtherEventTerm).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (oeq *OtherEventQuery) GroupBy(field string, fields ...string) *OtherEventGroupBy {
	grbuild := &OtherEventGroupBy{config: oeq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oeq.sqlQuery(ctx), nil
	}
	grbuild.label = otherevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OtherEventTerm string `json:"other_event_term,omitempty"`
//	}
//
//	client.OtherEvent.Query().
//		Select(otherevent.FieldOtherEventTerm).
//		Scan(ctx, &v)
//
func (oeq *OtherEventQuery) Select(fields ...string) *OtherEventSelect {
	oeq.fields = append(oeq.fields, fields...)
	selbuild := &OtherEventSelect{OtherEventQuery: oeq}
	selbuild.label = otherevent.Label
	selbuild.flds, selbuild.scan = &oeq.fields, selbuild.Scan
	return selbuild
}

func (oeq *OtherEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oeq.fields {
		if !otherevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if oeq.path != nil {
		prev, err := oeq.path(ctx)
		if err != nil {
			return err
		}
		oeq.sql = prev
	}
	return nil
}

func (oeq *OtherEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OtherEvent, error) {
	var (
		nodes       = []*OtherEvent{}
		withFKs     = oeq.withFKs
		_spec       = oeq.querySpec()
		loadedTypes = [2]bool{
			oeq.withParent != nil,
			oeq.withOtherEventStatsList != nil,
		}
	)
	if oeq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, otherevent.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OtherEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OtherEvent{config: oeq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oeq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OtherEvent)
		for i := range nodes {
			if nodes[i].adverse_events_module_other_event_list == nil {
				continue
			}
			fk := *nodes[i].adverse_events_module_other_event_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "adverse_events_module_other_event_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := oeq.withOtherEventStatsList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OtherEvent)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OtherEventStatsList = []*OtherEventStats{}
		}
		query.withFKs = true
		query.Where(predicate.OtherEventStats(func(s *sql.Selector) {
			s.Where(sql.InValues(otherevent.OtherEventStatsListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.other_event_other_event_stats_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "other_event_other_event_stats_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "other_event_other_event_stats_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OtherEventStatsList = append(node.Edges.OtherEventStatsList, n)
		}
	}

	return nodes, nil
}

func (oeq *OtherEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oeq.querySpec()
	_spec.Node.Columns = oeq.fields
	if len(oeq.fields) > 0 {
		_spec.Unique = oeq.unique != nil && *oeq.unique
	}
	return sqlgraph.CountNodes(ctx, oeq.driver, _spec)
}

func (oeq *OtherEventQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oeq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (oeq *OtherEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   otherevent.Table,
			Columns: otherevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: otherevent.FieldID,
			},
		},
		From:   oeq.sql,
		Unique: true,
	}
	if unique := oeq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oeq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, otherevent.FieldID)
		for i := range fields {
			if fields[i] != otherevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oeq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oeq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oeq *OtherEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oeq.driver.Dialect())
	t1 := builder.Table(otherevent.Table)
	columns := oeq.fields
	if len(columns) == 0 {
		columns = otherevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oeq.sql != nil {
		selector = oeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oeq.unique != nil && *oeq.unique {
		selector.Distinct()
	}
	for _, p := range oeq.predicates {
		p(selector)
	}
	for _, p := range oeq.order {
		p(selector)
	}
	if offset := oeq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oeq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OtherEventGroupBy is the group-by builder for OtherEvent entities.
type OtherEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oegb *OtherEventGroupBy) Aggregate(fns ...AggregateFunc) *OtherEventGroupBy {
	oegb.fns = append(oegb.fns, fns...)
	return oegb
}

// Scan applies the group-by query and scans the result into the given value.
func (oegb *OtherEventGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oegb.path(ctx)
	if err != nil {
		return err
	}
	oegb.sql = query
	return oegb.sqlScan(ctx, v)
}

func (oegb *OtherEventGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oegb.fields {
		if !otherevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oegb *OtherEventGroupBy) sqlQuery() *sql.Selector {
	selector := oegb.sql.Select()
	aggregation := make([]string, 0, len(oegb.fns))
	for _, fn := range oegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oegb.fields)+len(oegb.fns))
		for _, f := range oegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oegb.fields...)...)
}

// OtherEventSelect is the builder for selecting fields of OtherEvent entities.
type OtherEventSelect struct {
	*OtherEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oes *OtherEventSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oes.prepareQuery(ctx); err != nil {
		return err
	}
	oes.sql = oes.OtherEventQuery.sqlQuery(ctx)
	return oes.sqlScan(ctx, v)
}

func (oes *OtherEventSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oes.sql.Query()
	if err := oes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
