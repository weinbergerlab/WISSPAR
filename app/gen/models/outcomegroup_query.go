// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeGroupQuery is the builder for querying OutcomeGroup entities.
type OutcomeGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeGroup
	// eager-loading edges.
	withParent *OutcomeMeasureQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeGroupQuery builder.
func (ogq *OutcomeGroupQuery) Where(ps ...predicate.OutcomeGroup) *OutcomeGroupQuery {
	ogq.predicates = append(ogq.predicates, ps...)
	return ogq
}

// Limit adds a limit step to the query.
func (ogq *OutcomeGroupQuery) Limit(limit int) *OutcomeGroupQuery {
	ogq.limit = &limit
	return ogq
}

// Offset adds an offset step to the query.
func (ogq *OutcomeGroupQuery) Offset(offset int) *OutcomeGroupQuery {
	ogq.offset = &offset
	return ogq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ogq *OutcomeGroupQuery) Unique(unique bool) *OutcomeGroupQuery {
	ogq.unique = &unique
	return ogq
}

// Order adds an order step to the query.
func (ogq *OutcomeGroupQuery) Order(o ...OrderFunc) *OutcomeGroupQuery {
	ogq.order = append(ogq.order, o...)
	return ogq
}

// QueryParent chains the current query on the "parent" edge.
func (ogq *OutcomeGroupQuery) QueryParent() *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: ogq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ogq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ogq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomegroup.Table, outcomegroup.FieldID, selector),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomegroup.ParentTable, outcomegroup.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ogq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeGroup entity from the query.
// Returns a *NotFoundError when no OutcomeGroup was found.
func (ogq *OutcomeGroupQuery) First(ctx context.Context) (*OutcomeGroup, error) {
	nodes, err := ogq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomegroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) FirstX(ctx context.Context) *OutcomeGroup {
	node, err := ogq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeGroup ID from the query.
// Returns a *NotFoundError when no OutcomeGroup ID was found.
func (ogq *OutcomeGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ogq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomegroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := ogq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeGroup entity is found.
// Returns a *NotFoundError when no OutcomeGroup entities are found.
func (ogq *OutcomeGroupQuery) Only(ctx context.Context) (*OutcomeGroup, error) {
	nodes, err := ogq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomegroup.Label}
	default:
		return nil, &NotSingularError{outcomegroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) OnlyX(ctx context.Context) *OutcomeGroup {
	node, err := ogq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeGroup ID in the query.
// Returns a *NotSingularError when more than one OutcomeGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (ogq *OutcomeGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ogq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomegroup.Label}
	default:
		err = &NotSingularError{outcomegroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := ogq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeGroups.
func (ogq *OutcomeGroupQuery) All(ctx context.Context) ([]*OutcomeGroup, error) {
	if err := ogq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ogq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) AllX(ctx context.Context) []*OutcomeGroup {
	nodes, err := ogq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeGroup IDs.
func (ogq *OutcomeGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ogq.Select(outcomegroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := ogq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ogq *OutcomeGroupQuery) Count(ctx context.Context) (int, error) {
	if err := ogq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ogq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) CountX(ctx context.Context) int {
	count, err := ogq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ogq *OutcomeGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := ogq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ogq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ogq *OutcomeGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := ogq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ogq *OutcomeGroupQuery) Clone() *OutcomeGroupQuery {
	if ogq == nil {
		return nil
	}
	return &OutcomeGroupQuery{
		config:     ogq.config,
		limit:      ogq.limit,
		offset:     ogq.offset,
		order:      append([]OrderFunc{}, ogq.order...),
		predicates: append([]predicate.OutcomeGroup{}, ogq.predicates...),
		withParent: ogq.withParent.Clone(),
		// clone intermediate query.
		sql:    ogq.sql.Clone(),
		path:   ogq.path,
		unique: ogq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ogq *OutcomeGroupQuery) WithParent(opts ...func(*OutcomeMeasureQuery)) *OutcomeGroupQuery {
	query := &OutcomeMeasureQuery{config: ogq.config}
	for _, opt := range opts {
		opt(query)
	}
	ogq.withParent = query
	return ogq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeGroupID string `json:"outcome_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeGroup.Query().
//		GroupBy(outcomegroup.FieldOutcomeGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ogq *OutcomeGroupQuery) GroupBy(field string, fields ...string) *OutcomeGroupGroupBy {
	grbuild := &OutcomeGroupGroupBy{config: ogq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ogq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ogq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomegroup.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeGroupID string `json:"outcome_group_id,omitempty"`
//	}
//
//	client.OutcomeGroup.Query().
//		Select(outcomegroup.FieldOutcomeGroupID).
//		Scan(ctx, &v)
//
func (ogq *OutcomeGroupQuery) Select(fields ...string) *OutcomeGroupSelect {
	ogq.fields = append(ogq.fields, fields...)
	selbuild := &OutcomeGroupSelect{OutcomeGroupQuery: ogq}
	selbuild.label = outcomegroup.Label
	selbuild.flds, selbuild.scan = &ogq.fields, selbuild.Scan
	return selbuild
}

func (ogq *OutcomeGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ogq.fields {
		if !outcomegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ogq.path != nil {
		prev, err := ogq.path(ctx)
		if err != nil {
			return err
		}
		ogq.sql = prev
	}
	return nil
}

func (ogq *OutcomeGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeGroup, error) {
	var (
		nodes       = []*OutcomeGroup{}
		withFKs     = ogq.withFKs
		_spec       = ogq.querySpec()
		loadedTypes = [1]bool{
			ogq.withParent != nil,
		}
	)
	if ogq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomegroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeGroup{config: ogq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ogq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ogq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeGroup)
		for i := range nodes {
			if nodes[i].outcome_measure_outcome_group_list == nil {
				continue
			}
			fk := *nodes[i].outcome_measure_outcome_group_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_group_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (ogq *OutcomeGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ogq.querySpec()
	_spec.Node.Columns = ogq.fields
	if len(ogq.fields) > 0 {
		_spec.Unique = ogq.unique != nil && *ogq.unique
	}
	return sqlgraph.CountNodes(ctx, ogq.driver, _spec)
}

func (ogq *OutcomeGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ogq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ogq *OutcomeGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomegroup.Table,
			Columns: outcomegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomegroup.FieldID,
			},
		},
		From:   ogq.sql,
		Unique: true,
	}
	if unique := ogq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ogq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomegroup.FieldID)
		for i := range fields {
			if fields[i] != outcomegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ogq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ogq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ogq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ogq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ogq *OutcomeGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ogq.driver.Dialect())
	t1 := builder.Table(outcomegroup.Table)
	columns := ogq.fields
	if len(columns) == 0 {
		columns = outcomegroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ogq.sql != nil {
		selector = ogq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ogq.unique != nil && *ogq.unique {
		selector.Distinct()
	}
	for _, p := range ogq.predicates {
		p(selector)
	}
	for _, p := range ogq.order {
		p(selector)
	}
	if offset := ogq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ogq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeGroupGroupBy is the group-by builder for OutcomeGroup entities.
type OutcomeGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oggb *OutcomeGroupGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeGroupGroupBy {
	oggb.fns = append(oggb.fns, fns...)
	return oggb
}

// Scan applies the group-by query and scans the result into the given value.
func (oggb *OutcomeGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oggb.path(ctx)
	if err != nil {
		return err
	}
	oggb.sql = query
	return oggb.sqlScan(ctx, v)
}

func (oggb *OutcomeGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oggb.fields {
		if !outcomegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oggb *OutcomeGroupGroupBy) sqlQuery() *sql.Selector {
	selector := oggb.sql.Select()
	aggregation := make([]string, 0, len(oggb.fns))
	for _, fn := range oggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oggb.fields)+len(oggb.fns))
		for _, f := range oggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oggb.fields...)...)
}

// OutcomeGroupSelect is the builder for selecting fields of OutcomeGroup entities.
type OutcomeGroupSelect struct {
	*OutcomeGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ogs *OutcomeGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ogs.prepareQuery(ctx); err != nil {
		return err
	}
	ogs.sql = ogs.OutcomeGroupQuery.sqlQuery(ctx)
	return ogs.sqlScan(ctx, v)
}

func (ogs *OutcomeGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ogs.sql.Query()
	if err := ogs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
