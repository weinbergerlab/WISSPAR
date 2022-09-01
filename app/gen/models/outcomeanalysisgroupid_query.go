// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeAnalysisGroupIDQuery is the builder for querying OutcomeAnalysisGroupID entities.
type OutcomeAnalysisGroupIDQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeAnalysisGroupID
	// eager-loading edges.
	withParent *OutcomeAnalysisQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeAnalysisGroupIDQuery builder.
func (oagiq *OutcomeAnalysisGroupIDQuery) Where(ps ...predicate.OutcomeAnalysisGroupID) *OutcomeAnalysisGroupIDQuery {
	oagiq.predicates = append(oagiq.predicates, ps...)
	return oagiq
}

// Limit adds a limit step to the query.
func (oagiq *OutcomeAnalysisGroupIDQuery) Limit(limit int) *OutcomeAnalysisGroupIDQuery {
	oagiq.limit = &limit
	return oagiq
}

// Offset adds an offset step to the query.
func (oagiq *OutcomeAnalysisGroupIDQuery) Offset(offset int) *OutcomeAnalysisGroupIDQuery {
	oagiq.offset = &offset
	return oagiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oagiq *OutcomeAnalysisGroupIDQuery) Unique(unique bool) *OutcomeAnalysisGroupIDQuery {
	oagiq.unique = &unique
	return oagiq
}

// Order adds an order step to the query.
func (oagiq *OutcomeAnalysisGroupIDQuery) Order(o ...OrderFunc) *OutcomeAnalysisGroupIDQuery {
	oagiq.order = append(oagiq.order, o...)
	return oagiq
}

// QueryParent chains the current query on the "parent" edge.
func (oagiq *OutcomeAnalysisGroupIDQuery) QueryParent() *OutcomeAnalysisQuery {
	query := &OutcomeAnalysisQuery{config: oagiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oagiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oagiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeanalysisgroupid.Table, outcomeanalysisgroupid.FieldID, selector),
			sqlgraph.To(outcomeanalysis.Table, outcomeanalysis.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeanalysisgroupid.ParentTable, outcomeanalysisgroupid.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(oagiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeAnalysisGroupID entity from the query.
// Returns a *NotFoundError when no OutcomeAnalysisGroupID was found.
func (oagiq *OutcomeAnalysisGroupIDQuery) First(ctx context.Context) (*OutcomeAnalysisGroupID, error) {
	nodes, err := oagiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomeanalysisgroupid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) FirstX(ctx context.Context) *OutcomeAnalysisGroupID {
	node, err := oagiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeAnalysisGroupID ID from the query.
// Returns a *NotFoundError when no OutcomeAnalysisGroupID ID was found.
func (oagiq *OutcomeAnalysisGroupIDQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oagiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomeanalysisgroupid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) FirstIDX(ctx context.Context) int {
	id, err := oagiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeAnalysisGroupID entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeAnalysisGroupID entity is found.
// Returns a *NotFoundError when no OutcomeAnalysisGroupID entities are found.
func (oagiq *OutcomeAnalysisGroupIDQuery) Only(ctx context.Context) (*OutcomeAnalysisGroupID, error) {
	nodes, err := oagiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomeanalysisgroupid.Label}
	default:
		return nil, &NotSingularError{outcomeanalysisgroupid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) OnlyX(ctx context.Context) *OutcomeAnalysisGroupID {
	node, err := oagiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeAnalysisGroupID ID in the query.
// Returns a *NotSingularError when more than one OutcomeAnalysisGroupID ID is found.
// Returns a *NotFoundError when no entities are found.
func (oagiq *OutcomeAnalysisGroupIDQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oagiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomeanalysisgroupid.Label}
	default:
		err = &NotSingularError{outcomeanalysisgroupid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) OnlyIDX(ctx context.Context) int {
	id, err := oagiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeAnalysisGroupIDs.
func (oagiq *OutcomeAnalysisGroupIDQuery) All(ctx context.Context) ([]*OutcomeAnalysisGroupID, error) {
	if err := oagiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oagiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) AllX(ctx context.Context) []*OutcomeAnalysisGroupID {
	nodes, err := oagiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeAnalysisGroupID IDs.
func (oagiq *OutcomeAnalysisGroupIDQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oagiq.Select(outcomeanalysisgroupid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) IDsX(ctx context.Context) []int {
	ids, err := oagiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oagiq *OutcomeAnalysisGroupIDQuery) Count(ctx context.Context) (int, error) {
	if err := oagiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oagiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) CountX(ctx context.Context) int {
	count, err := oagiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oagiq *OutcomeAnalysisGroupIDQuery) Exist(ctx context.Context) (bool, error) {
	if err := oagiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oagiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oagiq *OutcomeAnalysisGroupIDQuery) ExistX(ctx context.Context) bool {
	exist, err := oagiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeAnalysisGroupIDQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oagiq *OutcomeAnalysisGroupIDQuery) Clone() *OutcomeAnalysisGroupIDQuery {
	if oagiq == nil {
		return nil
	}
	return &OutcomeAnalysisGroupIDQuery{
		config:     oagiq.config,
		limit:      oagiq.limit,
		offset:     oagiq.offset,
		order:      append([]OrderFunc{}, oagiq.order...),
		predicates: append([]predicate.OutcomeAnalysisGroupID{}, oagiq.predicates...),
		withParent: oagiq.withParent.Clone(),
		// clone intermediate query.
		sql:    oagiq.sql.Clone(),
		path:   oagiq.path,
		unique: oagiq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (oagiq *OutcomeAnalysisGroupIDQuery) WithParent(opts ...func(*OutcomeAnalysisQuery)) *OutcomeAnalysisGroupIDQuery {
	query := &OutcomeAnalysisQuery{config: oagiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oagiq.withParent = query
	return oagiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeAnalysisGroupID string `json:"outcome_analysis_group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeAnalysisGroupID.Query().
//		GroupBy(outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (oagiq *OutcomeAnalysisGroupIDQuery) GroupBy(field string, fields ...string) *OutcomeAnalysisGroupIDGroupBy {
	grbuild := &OutcomeAnalysisGroupIDGroupBy{config: oagiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oagiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oagiq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomeanalysisgroupid.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeAnalysisGroupID string `json:"outcome_analysis_group_id,omitempty"`
//	}
//
//	client.OutcomeAnalysisGroupID.Query().
//		Select(outcomeanalysisgroupid.FieldOutcomeAnalysisGroupID).
//		Scan(ctx, &v)
//
func (oagiq *OutcomeAnalysisGroupIDQuery) Select(fields ...string) *OutcomeAnalysisGroupIDSelect {
	oagiq.fields = append(oagiq.fields, fields...)
	selbuild := &OutcomeAnalysisGroupIDSelect{OutcomeAnalysisGroupIDQuery: oagiq}
	selbuild.label = outcomeanalysisgroupid.Label
	selbuild.flds, selbuild.scan = &oagiq.fields, selbuild.Scan
	return selbuild
}

func (oagiq *OutcomeAnalysisGroupIDQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oagiq.fields {
		if !outcomeanalysisgroupid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if oagiq.path != nil {
		prev, err := oagiq.path(ctx)
		if err != nil {
			return err
		}
		oagiq.sql = prev
	}
	return nil
}

func (oagiq *OutcomeAnalysisGroupIDQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeAnalysisGroupID, error) {
	var (
		nodes       = []*OutcomeAnalysisGroupID{}
		withFKs     = oagiq.withFKs
		_spec       = oagiq.querySpec()
		loadedTypes = [1]bool{
			oagiq.withParent != nil,
		}
	)
	if oagiq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeanalysisgroupid.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeAnalysisGroupID).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeAnalysisGroupID{config: oagiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oagiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oagiq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeAnalysisGroupID)
		for i := range nodes {
			if nodes[i].outcome_analysis_outcome_analysis_group_id_list == nil {
				continue
			}
			fk := *nodes[i].outcome_analysis_outcome_analysis_group_id_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomeanalysis.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_analysis_outcome_analysis_group_id_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (oagiq *OutcomeAnalysisGroupIDQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oagiq.querySpec()
	_spec.Node.Columns = oagiq.fields
	if len(oagiq.fields) > 0 {
		_spec.Unique = oagiq.unique != nil && *oagiq.unique
	}
	return sqlgraph.CountNodes(ctx, oagiq.driver, _spec)
}

func (oagiq *OutcomeAnalysisGroupIDQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oagiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (oagiq *OutcomeAnalysisGroupIDQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeanalysisgroupid.Table,
			Columns: outcomeanalysisgroupid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysisgroupid.FieldID,
			},
		},
		From:   oagiq.sql,
		Unique: true,
	}
	if unique := oagiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oagiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeanalysisgroupid.FieldID)
		for i := range fields {
			if fields[i] != outcomeanalysisgroupid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oagiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oagiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oagiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oagiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oagiq *OutcomeAnalysisGroupIDQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oagiq.driver.Dialect())
	t1 := builder.Table(outcomeanalysisgroupid.Table)
	columns := oagiq.fields
	if len(columns) == 0 {
		columns = outcomeanalysisgroupid.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oagiq.sql != nil {
		selector = oagiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oagiq.unique != nil && *oagiq.unique {
		selector.Distinct()
	}
	for _, p := range oagiq.predicates {
		p(selector)
	}
	for _, p := range oagiq.order {
		p(selector)
	}
	if offset := oagiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oagiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeAnalysisGroupIDGroupBy is the group-by builder for OutcomeAnalysisGroupID entities.
type OutcomeAnalysisGroupIDGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oagigb *OutcomeAnalysisGroupIDGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeAnalysisGroupIDGroupBy {
	oagigb.fns = append(oagigb.fns, fns...)
	return oagigb
}

// Scan applies the group-by query and scans the result into the given value.
func (oagigb *OutcomeAnalysisGroupIDGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oagigb.path(ctx)
	if err != nil {
		return err
	}
	oagigb.sql = query
	return oagigb.sqlScan(ctx, v)
}

func (oagigb *OutcomeAnalysisGroupIDGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oagigb.fields {
		if !outcomeanalysisgroupid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oagigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oagigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oagigb *OutcomeAnalysisGroupIDGroupBy) sqlQuery() *sql.Selector {
	selector := oagigb.sql.Select()
	aggregation := make([]string, 0, len(oagigb.fns))
	for _, fn := range oagigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oagigb.fields)+len(oagigb.fns))
		for _, f := range oagigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oagigb.fields...)...)
}

// OutcomeAnalysisGroupIDSelect is the builder for selecting fields of OutcomeAnalysisGroupID entities.
type OutcomeAnalysisGroupIDSelect struct {
	*OutcomeAnalysisGroupIDQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oagis *OutcomeAnalysisGroupIDSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oagis.prepareQuery(ctx); err != nil {
		return err
	}
	oagis.sql = oagis.OutcomeAnalysisGroupIDQuery.sqlQuery(ctx)
	return oagis.sqlScan(ctx, v)
}

func (oagis *OutcomeAnalysisGroupIDSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oagis.sql.Query()
	if err := oagis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
