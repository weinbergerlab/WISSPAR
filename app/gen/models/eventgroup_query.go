// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// EventGroupQuery is the builder for querying EventGroup entities.
type EventGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EventGroup
	// eager-loading edges.
	withParent *AdverseEventsModuleQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventGroupQuery builder.
func (egq *EventGroupQuery) Where(ps ...predicate.EventGroup) *EventGroupQuery {
	egq.predicates = append(egq.predicates, ps...)
	return egq
}

// Limit adds a limit step to the query.
func (egq *EventGroupQuery) Limit(limit int) *EventGroupQuery {
	egq.limit = &limit
	return egq
}

// Offset adds an offset step to the query.
func (egq *EventGroupQuery) Offset(offset int) *EventGroupQuery {
	egq.offset = &offset
	return egq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (egq *EventGroupQuery) Unique(unique bool) *EventGroupQuery {
	egq.unique = &unique
	return egq
}

// Order adds an order step to the query.
func (egq *EventGroupQuery) Order(o ...OrderFunc) *EventGroupQuery {
	egq.order = append(egq.order, o...)
	return egq
}

// QueryParent chains the current query on the "parent" edge.
func (egq *EventGroupQuery) QueryParent() *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: egq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := egq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := egq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventgroup.Table, eventgroup.FieldID, selector),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eventgroup.ParentTable, eventgroup.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(egq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EventGroup entity from the query.
// Returns a *NotFoundError when no EventGroup was found.
func (egq *EventGroupQuery) First(ctx context.Context) (*EventGroup, error) {
	nodes, err := egq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (egq *EventGroupQuery) FirstX(ctx context.Context) *EventGroup {
	node, err := egq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventGroup ID from the query.
// Returns a *NotFoundError when no EventGroup ID was found.
func (egq *EventGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = egq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (egq *EventGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := egq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EventGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EventGroup entity is found.
// Returns a *NotFoundError when no EventGroup entities are found.
func (egq *EventGroupQuery) Only(ctx context.Context) (*EventGroup, error) {
	nodes, err := egq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventgroup.Label}
	default:
		return nil, &NotSingularError{eventgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (egq *EventGroupQuery) OnlyX(ctx context.Context) *EventGroup {
	node, err := egq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EventGroup ID in the query.
// Returns a *NotSingularError when more than one EventGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (egq *EventGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = egq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventgroup.Label}
	default:
		err = &NotSingularError{eventgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (egq *EventGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := egq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventGroups.
func (egq *EventGroupQuery) All(ctx context.Context) ([]*EventGroup, error) {
	if err := egq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return egq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (egq *EventGroupQuery) AllX(ctx context.Context) []*EventGroup {
	nodes, err := egq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventGroup IDs.
func (egq *EventGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := egq.Select(eventgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (egq *EventGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := egq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (egq *EventGroupQuery) Count(ctx context.Context) (int, error) {
	if err := egq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return egq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (egq *EventGroupQuery) CountX(ctx context.Context) int {
	count, err := egq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (egq *EventGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := egq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return egq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (egq *EventGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := egq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (egq *EventGroupQuery) Clone() *EventGroupQuery {
	if egq == nil {
		return nil
	}
	return &EventGroupQuery{
		config:     egq.config,
		limit:      egq.limit,
		offset:     egq.offset,
		order:      append([]OrderFunc{}, egq.order...),
		predicates: append([]predicate.EventGroup{}, egq.predicates...),
		withParent: egq.withParent.Clone(),
		// clone intermediate query.
		sql:    egq.sql.Clone(),
		path:   egq.path,
		unique: egq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (egq *EventGroupQuery) WithParent(opts ...func(*AdverseEventsModuleQuery)) *EventGroupQuery {
	query := &AdverseEventsModuleQuery{config: egq.config}
	for _, opt := range opts {
		opt(query)
	}
	egq.withParent = query
	return egq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventGroupID string `json:"event_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EventGroup.Query().
//		GroupBy(eventgroup.FieldEventGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (egq *EventGroupQuery) GroupBy(field string, fields ...string) *EventGroupGroupBy {
	grbuild := &EventGroupGroupBy{config: egq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := egq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return egq.sqlQuery(ctx), nil
	}
	grbuild.label = eventgroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventGroupID string `json:"event_group_id,omitempty"`
//	}
//
//	client.EventGroup.Query().
//		Select(eventgroup.FieldEventGroupID).
//		Scan(ctx, &v)
//
func (egq *EventGroupQuery) Select(fields ...string) *EventGroupSelect {
	egq.fields = append(egq.fields, fields...)
	selbuild := &EventGroupSelect{EventGroupQuery: egq}
	selbuild.label = eventgroup.Label
	selbuild.flds, selbuild.scan = &egq.fields, selbuild.Scan
	return selbuild
}

func (egq *EventGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range egq.fields {
		if !eventgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if egq.path != nil {
		prev, err := egq.path(ctx)
		if err != nil {
			return err
		}
		egq.sql = prev
	}
	return nil
}

func (egq *EventGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EventGroup, error) {
	var (
		nodes       = []*EventGroup{}
		withFKs     = egq.withFKs
		_spec       = egq.querySpec()
		loadedTypes = [1]bool{
			egq.withParent != nil,
		}
	)
	if egq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, eventgroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*EventGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &EventGroup{config: egq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, egq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := egq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*EventGroup)
		for i := range nodes {
			if nodes[i].adverse_events_module_event_group_list == nil {
				continue
			}
			fk := *nodes[i].adverse_events_module_event_group_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "adverse_events_module_event_group_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (egq *EventGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := egq.querySpec()
	_spec.Node.Columns = egq.fields
	if len(egq.fields) > 0 {
		_spec.Unique = egq.unique != nil && *egq.unique
	}
	return sqlgraph.CountNodes(ctx, egq.driver, _spec)
}

func (egq *EventGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := egq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (egq *EventGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventgroup.Table,
			Columns: eventgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventgroup.FieldID,
			},
		},
		From:   egq.sql,
		Unique: true,
	}
	if unique := egq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := egq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventgroup.FieldID)
		for i := range fields {
			if fields[i] != eventgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := egq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := egq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := egq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := egq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (egq *EventGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(egq.driver.Dialect())
	t1 := builder.Table(eventgroup.Table)
	columns := egq.fields
	if len(columns) == 0 {
		columns = eventgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if egq.sql != nil {
		selector = egq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if egq.unique != nil && *egq.unique {
		selector.Distinct()
	}
	for _, p := range egq.predicates {
		p(selector)
	}
	for _, p := range egq.order {
		p(selector)
	}
	if offset := egq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := egq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventGroupGroupBy is the group-by builder for EventGroup entities.
type EventGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eggb *EventGroupGroupBy) Aggregate(fns ...AggregateFunc) *EventGroupGroupBy {
	eggb.fns = append(eggb.fns, fns...)
	return eggb
}

// Scan applies the group-by query and scans the result into the given value.
func (eggb *EventGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := eggb.path(ctx)
	if err != nil {
		return err
	}
	eggb.sql = query
	return eggb.sqlScan(ctx, v)
}

func (eggb *EventGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range eggb.fields {
		if !eventgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := eggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eggb *EventGroupGroupBy) sqlQuery() *sql.Selector {
	selector := eggb.sql.Select()
	aggregation := make([]string, 0, len(eggb.fns))
	for _, fn := range eggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(eggb.fields)+len(eggb.fns))
		for _, f := range eggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(eggb.fields...)...)
}

// EventGroupSelect is the builder for selecting fields of EventGroup entities.
type EventGroupSelect struct {
	*EventGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (egs *EventGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := egs.prepareQuery(ctx); err != nil {
		return err
	}
	egs.sql = egs.EventGroupQuery.sqlQuery(ctx)
	return egs.sqlScan(ctx, v)
}

func (egs *EventGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := egs.sql.Query()
	if err := egs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
