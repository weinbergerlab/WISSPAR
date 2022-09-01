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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysis"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeanalysisgroupid"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeAnalysisQuery is the builder for querying OutcomeAnalysis entities.
type OutcomeAnalysisQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeAnalysis
	// eager-loading edges.
	withParent                     *OutcomeMeasureQuery
	withOutcomeAnalysisGroupIDList *OutcomeAnalysisGroupIDQuery
	withFKs                        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeAnalysisQuery builder.
func (oaq *OutcomeAnalysisQuery) Where(ps ...predicate.OutcomeAnalysis) *OutcomeAnalysisQuery {
	oaq.predicates = append(oaq.predicates, ps...)
	return oaq
}

// Limit adds a limit step to the query.
func (oaq *OutcomeAnalysisQuery) Limit(limit int) *OutcomeAnalysisQuery {
	oaq.limit = &limit
	return oaq
}

// Offset adds an offset step to the query.
func (oaq *OutcomeAnalysisQuery) Offset(offset int) *OutcomeAnalysisQuery {
	oaq.offset = &offset
	return oaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oaq *OutcomeAnalysisQuery) Unique(unique bool) *OutcomeAnalysisQuery {
	oaq.unique = &unique
	return oaq
}

// Order adds an order step to the query.
func (oaq *OutcomeAnalysisQuery) Order(o ...OrderFunc) *OutcomeAnalysisQuery {
	oaq.order = append(oaq.order, o...)
	return oaq
}

// QueryParent chains the current query on the "parent" edge.
func (oaq *OutcomeAnalysisQuery) QueryParent() *OutcomeMeasureQuery {
	query := &OutcomeMeasureQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeanalysis.Table, outcomeanalysis.FieldID, selector),
			sqlgraph.To(outcomemeasure.Table, outcomemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomeanalysis.ParentTable, outcomeanalysis.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeAnalysisGroupIDList chains the current query on the "outcome_analysis_group_id_list" edge.
func (oaq *OutcomeAnalysisQuery) QueryOutcomeAnalysisGroupIDList() *OutcomeAnalysisGroupIDQuery {
	query := &OutcomeAnalysisGroupIDQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomeanalysis.Table, outcomeanalysis.FieldID, selector),
			sqlgraph.To(outcomeanalysisgroupid.Table, outcomeanalysisgroupid.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomeanalysis.OutcomeAnalysisGroupIDListTable, outcomeanalysis.OutcomeAnalysisGroupIDListColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeAnalysis entity from the query.
// Returns a *NotFoundError when no OutcomeAnalysis was found.
func (oaq *OutcomeAnalysisQuery) First(ctx context.Context) (*OutcomeAnalysis, error) {
	nodes, err := oaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomeanalysis.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) FirstX(ctx context.Context) *OutcomeAnalysis {
	node, err := oaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeAnalysis ID from the query.
// Returns a *NotFoundError when no OutcomeAnalysis ID was found.
func (oaq *OutcomeAnalysisQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomeanalysis.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) FirstIDX(ctx context.Context) int {
	id, err := oaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeAnalysis entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeAnalysis entity is found.
// Returns a *NotFoundError when no OutcomeAnalysis entities are found.
func (oaq *OutcomeAnalysisQuery) Only(ctx context.Context) (*OutcomeAnalysis, error) {
	nodes, err := oaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomeanalysis.Label}
	default:
		return nil, &NotSingularError{outcomeanalysis.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) OnlyX(ctx context.Context) *OutcomeAnalysis {
	node, err := oaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeAnalysis ID in the query.
// Returns a *NotSingularError when more than one OutcomeAnalysis ID is found.
// Returns a *NotFoundError when no entities are found.
func (oaq *OutcomeAnalysisQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomeanalysis.Label}
	default:
		err = &NotSingularError{outcomeanalysis.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) OnlyIDX(ctx context.Context) int {
	id, err := oaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeAnalyses.
func (oaq *OutcomeAnalysisQuery) All(ctx context.Context) ([]*OutcomeAnalysis, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) AllX(ctx context.Context) []*OutcomeAnalysis {
	nodes, err := oaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeAnalysis IDs.
func (oaq *OutcomeAnalysisQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oaq.Select(outcomeanalysis.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) IDsX(ctx context.Context) []int {
	ids, err := oaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oaq *OutcomeAnalysisQuery) Count(ctx context.Context) (int, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) CountX(ctx context.Context) int {
	count, err := oaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oaq *OutcomeAnalysisQuery) Exist(ctx context.Context) (bool, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oaq *OutcomeAnalysisQuery) ExistX(ctx context.Context) bool {
	exist, err := oaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeAnalysisQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oaq *OutcomeAnalysisQuery) Clone() *OutcomeAnalysisQuery {
	if oaq == nil {
		return nil
	}
	return &OutcomeAnalysisQuery{
		config:                         oaq.config,
		limit:                          oaq.limit,
		offset:                         oaq.offset,
		order:                          append([]OrderFunc{}, oaq.order...),
		predicates:                     append([]predicate.OutcomeAnalysis{}, oaq.predicates...),
		withParent:                     oaq.withParent.Clone(),
		withOutcomeAnalysisGroupIDList: oaq.withOutcomeAnalysisGroupIDList.Clone(),
		// clone intermediate query.
		sql:    oaq.sql.Clone(),
		path:   oaq.path,
		unique: oaq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OutcomeAnalysisQuery) WithParent(opts ...func(*OutcomeMeasureQuery)) *OutcomeAnalysisQuery {
	query := &OutcomeMeasureQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withParent = query
	return oaq
}

// WithOutcomeAnalysisGroupIDList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_analysis_group_id_list" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OutcomeAnalysisQuery) WithOutcomeAnalysisGroupIDList(opts ...func(*OutcomeAnalysisGroupIDQuery)) *OutcomeAnalysisQuery {
	query := &OutcomeAnalysisGroupIDQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOutcomeAnalysisGroupIDList = query
	return oaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeAnalysisGroupDescription string `json:"outcome_analysis_group_description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeAnalysis.Query().
//		GroupBy(outcomeanalysis.FieldOutcomeAnalysisGroupDescription).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (oaq *OutcomeAnalysisQuery) GroupBy(field string, fields ...string) *OutcomeAnalysisGroupBy {
	grbuild := &OutcomeAnalysisGroupBy{config: oaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oaq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomeanalysis.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeAnalysisGroupDescription string `json:"outcome_analysis_group_description,omitempty"`
//	}
//
//	client.OutcomeAnalysis.Query().
//		Select(outcomeanalysis.FieldOutcomeAnalysisGroupDescription).
//		Scan(ctx, &v)
//
func (oaq *OutcomeAnalysisQuery) Select(fields ...string) *OutcomeAnalysisSelect {
	oaq.fields = append(oaq.fields, fields...)
	selbuild := &OutcomeAnalysisSelect{OutcomeAnalysisQuery: oaq}
	selbuild.label = outcomeanalysis.Label
	selbuild.flds, selbuild.scan = &oaq.fields, selbuild.Scan
	return selbuild
}

func (oaq *OutcomeAnalysisQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oaq.fields {
		if !outcomeanalysis.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if oaq.path != nil {
		prev, err := oaq.path(ctx)
		if err != nil {
			return err
		}
		oaq.sql = prev
	}
	return nil
}

func (oaq *OutcomeAnalysisQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeAnalysis, error) {
	var (
		nodes       = []*OutcomeAnalysis{}
		withFKs     = oaq.withFKs
		_spec       = oaq.querySpec()
		loadedTypes = [2]bool{
			oaq.withParent != nil,
			oaq.withOutcomeAnalysisGroupIDList != nil,
		}
	)
	if oaq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeanalysis.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeAnalysis).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeAnalysis{config: oaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oaq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeAnalysis)
		for i := range nodes {
			if nodes[i].outcome_measure_outcome_analysis_list == nil {
				continue
			}
			fk := *nodes[i].outcome_measure_outcome_analysis_list
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
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_analysis_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := oaq.withOutcomeAnalysisGroupIDList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeAnalysis)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeAnalysisGroupIDList = []*OutcomeAnalysisGroupID{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeAnalysisGroupID(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomeanalysis.OutcomeAnalysisGroupIDListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_analysis_outcome_analysis_group_id_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_analysis_outcome_analysis_group_id_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_analysis_outcome_analysis_group_id_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeAnalysisGroupIDList = append(node.Edges.OutcomeAnalysisGroupIDList, n)
		}
	}

	return nodes, nil
}

func (oaq *OutcomeAnalysisQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oaq.querySpec()
	_spec.Node.Columns = oaq.fields
	if len(oaq.fields) > 0 {
		_spec.Unique = oaq.unique != nil && *oaq.unique
	}
	return sqlgraph.CountNodes(ctx, oaq.driver, _spec)
}

func (oaq *OutcomeAnalysisQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (oaq *OutcomeAnalysisQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeanalysis.Table,
			Columns: outcomeanalysis.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeanalysis.FieldID,
			},
		},
		From:   oaq.sql,
		Unique: true,
	}
	if unique := oaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeanalysis.FieldID)
		for i := range fields {
			if fields[i] != outcomeanalysis.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oaq *OutcomeAnalysisQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oaq.driver.Dialect())
	t1 := builder.Table(outcomeanalysis.Table)
	columns := oaq.fields
	if len(columns) == 0 {
		columns = outcomeanalysis.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oaq.sql != nil {
		selector = oaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oaq.unique != nil && *oaq.unique {
		selector.Distinct()
	}
	for _, p := range oaq.predicates {
		p(selector)
	}
	for _, p := range oaq.order {
		p(selector)
	}
	if offset := oaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeAnalysisGroupBy is the group-by builder for OutcomeAnalysis entities.
type OutcomeAnalysisGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oagb *OutcomeAnalysisGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeAnalysisGroupBy {
	oagb.fns = append(oagb.fns, fns...)
	return oagb
}

// Scan applies the group-by query and scans the result into the given value.
func (oagb *OutcomeAnalysisGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oagb.path(ctx)
	if err != nil {
		return err
	}
	oagb.sql = query
	return oagb.sqlScan(ctx, v)
}

func (oagb *OutcomeAnalysisGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oagb.fields {
		if !outcomeanalysis.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oagb *OutcomeAnalysisGroupBy) sqlQuery() *sql.Selector {
	selector := oagb.sql.Select()
	aggregation := make([]string, 0, len(oagb.fns))
	for _, fn := range oagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oagb.fields)+len(oagb.fns))
		for _, f := range oagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oagb.fields...)...)
}

// OutcomeAnalysisSelect is the builder for selecting fields of OutcomeAnalysis entities.
type OutcomeAnalysisSelect struct {
	*OutcomeAnalysisQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oas *OutcomeAnalysisSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oas.prepareQuery(ctx); err != nil {
		return err
	}
	oas.sql = oas.OutcomeAnalysisQuery.sqlQuery(ctx)
	return oas.sqlScan(ctx, v)
}

func (oas *OutcomeAnalysisSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oas.sql.Query()
	if err := oas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
