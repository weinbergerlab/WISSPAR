// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeOverviewQuery is the builder for querying OutcomeOverview entities.
type OutcomeOverviewQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeOverview
	// eager-loading edges.
	withParent *OutcomeMeasureQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeOverviewQuery builder.
func (ooq *OutcomeOverviewQuery) Where(ps ...predicate.OutcomeOverview) *OutcomeOverviewQuery {
	ooq.predicates = append(ooq.predicates, ps...)
	return ooq
}

// Limit adds a limit step to the query.
func (ooq *OutcomeOverviewQuery) Limit(limit int) *OutcomeOverviewQuery {
	ooq.limit = &limit
	return ooq
}

// Offset adds an offset step to the query.
func (ooq *OutcomeOverviewQuery) Offset(offset int) *OutcomeOverviewQuery {
	ooq.offset = &offset
	return ooq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ooq *OutcomeOverviewQuery) Unique(unique bool) *OutcomeOverviewQuery {
	ooq.unique = &unique
	return ooq
}

// Order adds an order step to the query.
func (ooq *OutcomeOverviewQuery) Order(o ...OrderFunc) *OutcomeOverviewQuery {
	ooq.order = append(ooq.order, o...)
	return ooq
}

// QueryParent chains the current query on the "parent" edge.
func (ooq *OutcomeOverviewQuery) QueryParent() *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: ooq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ooq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ooq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeoverview.Table, outcomeoverview.FieldID, selector),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeoverview.ParentTable, outcomeoverview.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ooq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeOverview entity from the query.
// Returns a *NotFoundError when no OutcomeOverview was found.
func (ooq *OutcomeOverviewQuery) First(ctx context.Context) (*OutcomeOverview, error) {
	nodes, err := ooq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomeoverview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) FirstX(ctx context.Context) *OutcomeOverview {
	node, err := ooq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeOverview ID from the query.
// Returns a *NotFoundError when no OutcomeOverview ID was found.
func (ooq *OutcomeOverviewQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ooq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomeoverview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) FirstIDX(ctx context.Context) int {
	id, err := ooq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeOverview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeOverview entity is found.
// Returns a *NotFoundError when no OutcomeOverview entities are found.
func (ooq *OutcomeOverviewQuery) Only(ctx context.Context) (*OutcomeOverview, error) {
	nodes, err := ooq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomeoverview.Label}
	default:
		return nil, &NotSingularError{outcomeoverview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) OnlyX(ctx context.Context) *OutcomeOverview {
	node, err := ooq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeOverview ID in the query.
// Returns a *NotSingularError when more than one OutcomeOverview ID is found.
// Returns a *NotFoundError when no entities are found.
func (ooq *OutcomeOverviewQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ooq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomeoverview.Label}
	default:
		err = &NotSingularError{outcomeoverview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) OnlyIDX(ctx context.Context) int {
	id, err := ooq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeOverviews.
func (ooq *OutcomeOverviewQuery) All(ctx context.Context) ([]*OutcomeOverview, error) {
	if err := ooq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ooq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) AllX(ctx context.Context) []*OutcomeOverview {
	nodes, err := ooq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeOverview IDs.
func (ooq *OutcomeOverviewQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ooq.Select(outcomeoverview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) IDsX(ctx context.Context) []int {
	ids, err := ooq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ooq *OutcomeOverviewQuery) Count(ctx context.Context) (int, error) {
	if err := ooq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ooq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) CountX(ctx context.Context) int {
	count, err := ooq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ooq *OutcomeOverviewQuery) Exist(ctx context.Context) (bool, error) {
	if err := ooq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ooq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ooq *OutcomeOverviewQuery) ExistX(ctx context.Context) bool {
	exist, err := ooq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeOverviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ooq *OutcomeOverviewQuery) Clone() *OutcomeOverviewQuery {
	if ooq == nil {
		return nil
	}
	return &OutcomeOverviewQuery{
		config:     ooq.config,
		limit:      ooq.limit,
		offset:     ooq.offset,
		order:      append([]OrderFunc{}, ooq.order...),
		predicates: append([]predicate.OutcomeOverview{}, ooq.predicates...),
		withParent: ooq.withParent.Clone(),
		// clone intermediate query.
		sql:    ooq.sql.Clone(),
		path:   ooq.path,
		unique: ooq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ooq *OutcomeOverviewQuery) WithParent(opts ...func(*OutcomeMeasureQuery)) *OutcomeOverviewQuery {
	query := &OutcomeMeasureQuery{config: ooq.config}
	for _, opt := range opts {
		opt(query)
	}
	ooq.withParent = query
	return ooq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeOverviewID string `json:"outcome_overview_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeOverview.Query().
//		GroupBy(outcomeoverview.FieldOutcomeOverviewID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ooq *OutcomeOverviewQuery) GroupBy(field string, fields ...string) *OutcomeOverviewGroupBy {
	grbuild := &OutcomeOverviewGroupBy{config: ooq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ooq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ooq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomeoverview.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeOverviewID string `json:"outcome_overview_id,omitempty"`
//	}
//
//	client.OutcomeOverview.Query().
//		Select(outcomeoverview.FieldOutcomeOverviewID).
//		Scan(ctx, &v)
//
func (ooq *OutcomeOverviewQuery) Select(fields ...string) *OutcomeOverviewSelect {
	ooq.fields = append(ooq.fields, fields...)
	selbuild := &OutcomeOverviewSelect{OutcomeOverviewQuery: ooq}
	selbuild.label = outcomeoverview.Label
	selbuild.flds, selbuild.scan = &ooq.fields, selbuild.Scan
	return selbuild
}

func (ooq *OutcomeOverviewQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ooq.fields {
		if !outcomeoverview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ooq.path != nil {
		prev, err := ooq.path(ctx)
		if err != nil {
			return err
		}
		ooq.sql = prev
	}
	return nil
}

func (ooq *OutcomeOverviewQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeOverview, error) {
	var (
		nodes       = []*OutcomeOverview{}
		withFKs     = ooq.withFKs
		_spec       = ooq.querySpec()
		loadedTypes = [1]bool{
			ooq.withParent != nil,
		}
	)
	if ooq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeoverview.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeOverview).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeOverview{config: ooq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ooq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ooq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeOverview)
		for i := range nodes {
			if nodes[i].outcome_measure_outcome_overview_list == nil {
				continue
			}
			fk := *nodes[i].outcome_measure_outcome_overview_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_overview_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (ooq *OutcomeOverviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ooq.querySpec()
	_spec.Node.Columns = ooq.fields
	if len(ooq.fields) > 0 {
		_spec.Unique = ooq.unique != nil && *ooq.unique
	}
	return sqlgraph.CountNodes(ctx, ooq.driver, _spec)
}

func (ooq *OutcomeOverviewQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ooq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ooq *OutcomeOverviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeoverview.Table,
			Columns: outcomeoverview.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeoverview.FieldID,
			},
		},
		From:   ooq.sql,
		Unique: true,
	}
	if unique := ooq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ooq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeoverview.FieldID)
		for i := range fields {
			if fields[i] != outcomeoverview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ooq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ooq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ooq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ooq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ooq *OutcomeOverviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ooq.driver.Dialect())
	t1 := builder.Table(outcomeoverview.Table)
	columns := ooq.fields
	if len(columns) == 0 {
		columns = outcomeoverview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ooq.sql != nil {
		selector = ooq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ooq.unique != nil && *ooq.unique {
		selector.Distinct()
	}
	for _, p := range ooq.predicates {
		p(selector)
	}
	for _, p := range ooq.order {
		p(selector)
	}
	if offset := ooq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ooq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeOverviewGroupBy is the group-by builder for OutcomeOverview entities.
type OutcomeOverviewGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oogb *OutcomeOverviewGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeOverviewGroupBy {
	oogb.fns = append(oogb.fns, fns...)
	return oogb
}

// Scan applies the group-by query and scans the result into the given value.
func (oogb *OutcomeOverviewGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oogb.path(ctx)
	if err != nil {
		return err
	}
	oogb.sql = query
	return oogb.sqlScan(ctx, v)
}

func (oogb *OutcomeOverviewGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oogb.fields {
		if !outcomeoverview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oogb *OutcomeOverviewGroupBy) sqlQuery() *sql.Selector {
	selector := oogb.sql.Select()
	aggregation := make([]string, 0, len(oogb.fns))
	for _, fn := range oogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oogb.fields)+len(oogb.fns))
		for _, f := range oogb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oogb.fields...)...)
}

// OutcomeOverviewSelect is the builder for selecting fields of OutcomeOverview entities.
type OutcomeOverviewSelect struct {
	*OutcomeOverviewQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oos *OutcomeOverviewSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oos.prepareQuery(ctx); err != nil {
		return err
	}
	oos.sql = oos.OutcomeOverviewQuery.sqlQuery(ctx)
	return oos.sqlScan(ctx, v)
}

func (oos *OutcomeOverviewSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oos.sql.Query()
	if err := oos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
