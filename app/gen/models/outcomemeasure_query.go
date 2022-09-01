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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeclass"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeMeasureQuery is the builder for querying OutcomeMeasure entities.
type OutcomeMeasureQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutcomeMeasure
	// eager-loading edges.
	withParent              *OutcomeMeasuresModuleQuery
	withOutcomeGroupList    *OutcomeGroupQuery
	withOutcomeOverviewList *OutcomeOverviewQuery
	withOutcomeDenomList    *OutcomeDenomQuery
	withOutcomeClassList    *OutcomeClassQuery
	withOutcomeAnalysisList *OutcomeAnalysisQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutcomeMeasureQuery builder.
func (omq *OutcomeMeasureQuery) Where(ps ...predicate.OutcomeMeasure) *OutcomeMeasureQuery {
	omq.predicates = append(omq.predicates, ps...)
	return omq
}

// Limit adds a limit step to the query.
func (omq *OutcomeMeasureQuery) Limit(limit int) *OutcomeMeasureQuery {
	omq.limit = &limit
	return omq
}

// Offset adds an offset step to the query.
func (omq *OutcomeMeasureQuery) Offset(offset int) *OutcomeMeasureQuery {
	omq.offset = &offset
	return omq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (omq *OutcomeMeasureQuery) Unique(unique bool) *OutcomeMeasureQuery {
	omq.unique = &unique
	return omq
}

// Order adds an order step to the query.
func (omq *OutcomeMeasureQuery) Order(o ...OrderFunc) *OutcomeMeasureQuery {
	omq.order = append(omq.order, o...)
	return omq
}

// QueryParent chains the current query on the "parent" edge.
func (omq *OutcomeMeasureQuery) QueryParent() *OutcomeMeasuresModuleQuery {
	query := &OutcomeMeasuresModuleQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, selector),
			sqlgraph.To(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outcomemeasure.ParentTable, outcomemeasure.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeGroupList chains the current query on the "outcome_group_list" edge.
func (omq *OutcomeMeasureQuery) QueryOutcomeGroupList() *OutcomeGroupQuery {
	query := &OutcomeGroupQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, selector),
			sqlgraph.To(outcomegroup.Table, outcomegroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeGroupListTable, outcomemeasure.OutcomeGroupListColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeOverviewList chains the current query on the "outcome_overview_list" edge.
func (omq *OutcomeMeasureQuery) QueryOutcomeOverviewList() *OutcomeOverviewQuery {
	query := &OutcomeOverviewQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, selector),
			sqlgraph.To(outcomeoverview.Table, outcomeoverview.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeOverviewListTable, outcomemeasure.OutcomeOverviewListColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeDenomList chains the current query on the "outcome_denom_list" edge.
func (omq *OutcomeMeasureQuery) QueryOutcomeDenomList() *OutcomeDenomQuery {
	query := &OutcomeDenomQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, selector),
			sqlgraph.To(outcomedenom.Table, outcomedenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeDenomListTable, outcomemeasure.OutcomeDenomListColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeClassList chains the current query on the "outcome_class_list" edge.
func (omq *OutcomeMeasureQuery) QueryOutcomeClassList() *OutcomeClassQuery {
	query := &OutcomeClassQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, selector),
			sqlgraph.To(outcomeclass.Table, outcomeclass.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeClassListTable, outcomemeasure.OutcomeClassListColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeAnalysisList chains the current query on the "outcome_analysis_list" edge.
func (omq *OutcomeMeasureQuery) QueryOutcomeAnalysisList() *OutcomeAnalysisQuery {
	query := &OutcomeAnalysisQuery{config: omq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outcomemeasure.Table, outcomemeasure.FieldID, selector),
			sqlgraph.To(outcomeanalysis.Table, outcomeanalysis.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outcomemeasure.OutcomeAnalysisListTable, outcomemeasure.OutcomeAnalysisListColumn),
		)
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutcomeMeasure entity from the query.
// Returns a *NotFoundError when no OutcomeMeasure was found.
func (omq *OutcomeMeasureQuery) First(ctx context.Context) (*OutcomeMeasure, error) {
	nodes, err := omq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outcomemeasure.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) FirstX(ctx context.Context) *OutcomeMeasure {
	node, err := omq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutcomeMeasure ID from the query.
// Returns a *NotFoundError when no OutcomeMeasure ID was found.
func (omq *OutcomeMeasureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = omq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outcomemeasure.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) FirstIDX(ctx context.Context) int {
	id, err := omq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutcomeMeasure entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutcomeMeasure entity is found.
// Returns a *NotFoundError when no OutcomeMeasure entities are found.
func (omq *OutcomeMeasureQuery) Only(ctx context.Context) (*OutcomeMeasure, error) {
	nodes, err := omq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outcomemeasure.Label}
	default:
		return nil, &NotSingularError{outcomemeasure.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) OnlyX(ctx context.Context) *OutcomeMeasure {
	node, err := omq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutcomeMeasure ID in the query.
// Returns a *NotSingularError when more than one OutcomeMeasure ID is found.
// Returns a *NotFoundError when no entities are found.
func (omq *OutcomeMeasureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = omq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outcomemeasure.Label}
	default:
		err = &NotSingularError{outcomemeasure.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) OnlyIDX(ctx context.Context) int {
	id, err := omq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutcomeMeasures.
func (omq *OutcomeMeasureQuery) All(ctx context.Context) ([]*OutcomeMeasure, error) {
	if err := omq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return omq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) AllX(ctx context.Context) []*OutcomeMeasure {
	nodes, err := omq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutcomeMeasure IDs.
func (omq *OutcomeMeasureQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := omq.Select(outcomemeasure.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) IDsX(ctx context.Context) []int {
	ids, err := omq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (omq *OutcomeMeasureQuery) Count(ctx context.Context) (int, error) {
	if err := omq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return omq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) CountX(ctx context.Context) int {
	count, err := omq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (omq *OutcomeMeasureQuery) Exist(ctx context.Context) (bool, error) {
	if err := omq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return omq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (omq *OutcomeMeasureQuery) ExistX(ctx context.Context) bool {
	exist, err := omq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutcomeMeasureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (omq *OutcomeMeasureQuery) Clone() *OutcomeMeasureQuery {
	if omq == nil {
		return nil
	}
	return &OutcomeMeasureQuery{
		config:                  omq.config,
		limit:                   omq.limit,
		offset:                  omq.offset,
		order:                   append([]OrderFunc{}, omq.order...),
		predicates:              append([]predicate.OutcomeMeasure{}, omq.predicates...),
		withParent:              omq.withParent.Clone(),
		withOutcomeGroupList:    omq.withOutcomeGroupList.Clone(),
		withOutcomeOverviewList: omq.withOutcomeOverviewList.Clone(),
		withOutcomeDenomList:    omq.withOutcomeDenomList.Clone(),
		withOutcomeClassList:    omq.withOutcomeClassList.Clone(),
		withOutcomeAnalysisList: omq.withOutcomeAnalysisList.Clone(),
		// clone intermediate query.
		sql:    omq.sql.Clone(),
		path:   omq.path,
		unique: omq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasureQuery) WithParent(opts ...func(*OutcomeMeasuresModuleQuery)) *OutcomeMeasureQuery {
	query := &OutcomeMeasuresModuleQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withParent = query
	return omq
}

// WithOutcomeGroupList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_group_list" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasureQuery) WithOutcomeGroupList(opts ...func(*OutcomeGroupQuery)) *OutcomeMeasureQuery {
	query := &OutcomeGroupQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withOutcomeGroupList = query
	return omq
}

// WithOutcomeOverviewList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_overview_list" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasureQuery) WithOutcomeOverviewList(opts ...func(*OutcomeOverviewQuery)) *OutcomeMeasureQuery {
	query := &OutcomeOverviewQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withOutcomeOverviewList = query
	return omq
}

// WithOutcomeDenomList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_denom_list" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasureQuery) WithOutcomeDenomList(opts ...func(*OutcomeDenomQuery)) *OutcomeMeasureQuery {
	query := &OutcomeDenomQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withOutcomeDenomList = query
	return omq
}

// WithOutcomeClassList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_class_list" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasureQuery) WithOutcomeClassList(opts ...func(*OutcomeClassQuery)) *OutcomeMeasureQuery {
	query := &OutcomeClassQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withOutcomeClassList = query
	return omq
}

// WithOutcomeAnalysisList tells the query-builder to eager-load the nodes that are connected to
// the "outcome_analysis_list" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OutcomeMeasureQuery) WithOutcomeAnalysisList(opts ...func(*OutcomeAnalysisQuery)) *OutcomeMeasureQuery {
	query := &OutcomeAnalysisQuery{config: omq.config}
	for _, opt := range opts {
		opt(query)
	}
	omq.withOutcomeAnalysisList = query
	return omq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OutcomeMeasureType string `json:"outcome_measure_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutcomeMeasure.Query().
//		GroupBy(outcomemeasure.FieldOutcomeMeasureType).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (omq *OutcomeMeasureQuery) GroupBy(field string, fields ...string) *OutcomeMeasureGroupBy {
	grbuild := &OutcomeMeasureGroupBy{config: omq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return omq.sqlQuery(ctx), nil
	}
	grbuild.label = outcomemeasure.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OutcomeMeasureType string `json:"outcome_measure_type,omitempty"`
//	}
//
//	client.OutcomeMeasure.Query().
//		Select(outcomemeasure.FieldOutcomeMeasureType).
//		Scan(ctx, &v)
//
func (omq *OutcomeMeasureQuery) Select(fields ...string) *OutcomeMeasureSelect {
	omq.fields = append(omq.fields, fields...)
	selbuild := &OutcomeMeasureSelect{OutcomeMeasureQuery: omq}
	selbuild.label = outcomemeasure.Label
	selbuild.flds, selbuild.scan = &omq.fields, selbuild.Scan
	return selbuild
}

func (omq *OutcomeMeasureQuery) prepareQuery(ctx context.Context) error {
	for _, f := range omq.fields {
		if !outcomemeasure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if omq.path != nil {
		prev, err := omq.path(ctx)
		if err != nil {
			return err
		}
		omq.sql = prev
	}
	return nil
}

func (omq *OutcomeMeasureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutcomeMeasure, error) {
	var (
		nodes       = []*OutcomeMeasure{}
		withFKs     = omq.withFKs
		_spec       = omq.querySpec()
		loadedTypes = [6]bool{
			omq.withParent != nil,
			omq.withOutcomeGroupList != nil,
			omq.withOutcomeOverviewList != nil,
			omq.withOutcomeDenomList != nil,
			omq.withOutcomeClassList != nil,
			omq.withOutcomeAnalysisList != nil,
		}
	)
	if omq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasure.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutcomeMeasure).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutcomeMeasure{config: omq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, omq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := omq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutcomeMeasure)
		for i := range nodes {
			if nodes[i].outcome_measures_module_outcome_measure_list == nil {
				continue
			}
			fk := *nodes[i].outcome_measures_module_outcome_measure_list
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(outcomemeasuresmodule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measures_module_outcome_measure_list" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := omq.withOutcomeGroupList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeGroupList = []*OutcomeGroup{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeGroup(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomemeasure.OutcomeGroupListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_measure_outcome_group_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_measure_outcome_group_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_group_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeGroupList = append(node.Edges.OutcomeGroupList, n)
		}
	}

	if query := omq.withOutcomeOverviewList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeOverviewList = []*OutcomeOverview{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeOverview(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomemeasure.OutcomeOverviewListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_measure_outcome_overview_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_measure_outcome_overview_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_overview_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeOverviewList = append(node.Edges.OutcomeOverviewList, n)
		}
	}

	if query := omq.withOutcomeDenomList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeDenomList = []*OutcomeDenom{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeDenom(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomemeasure.OutcomeDenomListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_measure_outcome_denom_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_measure_outcome_denom_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_denom_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeDenomList = append(node.Edges.OutcomeDenomList, n)
		}
	}

	if query := omq.withOutcomeClassList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeClassList = []*OutcomeClass{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeClass(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomemeasure.OutcomeClassListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_measure_outcome_class_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_measure_outcome_class_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_class_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeClassList = append(node.Edges.OutcomeClassList, n)
		}
	}

	if query := omq.withOutcomeAnalysisList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OutcomeMeasure)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OutcomeAnalysisList = []*OutcomeAnalysis{}
		}
		query.withFKs = true
		query.Where(predicate.OutcomeAnalysis(func(s *sql.Selector) {
			s.Where(sql.InValues(outcomemeasure.OutcomeAnalysisListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outcome_measure_outcome_analysis_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outcome_measure_outcome_analysis_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outcome_measure_outcome_analysis_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeAnalysisList = append(node.Edges.OutcomeAnalysisList, n)
		}
	}

	return nodes, nil
}

func (omq *OutcomeMeasureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := omq.querySpec()
	_spec.Node.Columns = omq.fields
	if len(omq.fields) > 0 {
		_spec.Unique = omq.unique != nil && *omq.unique
	}
	return sqlgraph.CountNodes(ctx, omq.driver, _spec)
}

func (omq *OutcomeMeasureQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := omq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (omq *OutcomeMeasureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomemeasure.Table,
			Columns: outcomemeasure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomemeasure.FieldID,
			},
		},
		From:   omq.sql,
		Unique: true,
	}
	if unique := omq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := omq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomemeasure.FieldID)
		for i := range fields {
			if fields[i] != outcomemeasure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := omq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := omq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := omq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := omq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (omq *OutcomeMeasureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(omq.driver.Dialect())
	t1 := builder.Table(outcomemeasure.Table)
	columns := omq.fields
	if len(columns) == 0 {
		columns = outcomemeasure.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if omq.sql != nil {
		selector = omq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if omq.unique != nil && *omq.unique {
		selector.Distinct()
	}
	for _, p := range omq.predicates {
		p(selector)
	}
	for _, p := range omq.order {
		p(selector)
	}
	if offset := omq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := omq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutcomeMeasureGroupBy is the group-by builder for OutcomeMeasure entities.
type OutcomeMeasureGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (omgb *OutcomeMeasureGroupBy) Aggregate(fns ...AggregateFunc) *OutcomeMeasureGroupBy {
	omgb.fns = append(omgb.fns, fns...)
	return omgb
}

// Scan applies the group-by query and scans the result into the given value.
func (omgb *OutcomeMeasureGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := omgb.path(ctx)
	if err != nil {
		return err
	}
	omgb.sql = query
	return omgb.sqlScan(ctx, v)
}

func (omgb *OutcomeMeasureGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range omgb.fields {
		if !outcomemeasure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := omgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := omgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (omgb *OutcomeMeasureGroupBy) sqlQuery() *sql.Selector {
	selector := omgb.sql.Select()
	aggregation := make([]string, 0, len(omgb.fns))
	for _, fn := range omgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(omgb.fields)+len(omgb.fns))
		for _, f := range omgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(omgb.fields...)...)
}

// OutcomeMeasureSelect is the builder for selecting fields of OutcomeMeasure entities.
type OutcomeMeasureSelect struct {
	*OutcomeMeasureQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oms *OutcomeMeasureSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oms.prepareQuery(ctx); err != nil {
		return err
	}
	oms.sql = oms.OutcomeMeasureQuery.sqlQuery(ctx)
	return oms.sqlScan(ctx, v)
}

func (oms *OutcomeMeasureSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oms.sql.Query()
	if err := oms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
