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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ResultsDefinitionQuery is the builder for querying ResultsDefinition entities.
type ResultsDefinitionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ResultsDefinition
	// eager-loading edges.
	withParent                        *ClinicalTrialQuery
	withParticipantFlowModule         *ParticipantFlowModuleQuery
	withBaselineCharacteristicsModule *BaselineCharacteristicsModuleQuery
	withOutcomeMeasuresModule         *OutcomeMeasuresModuleQuery
	withAdverseEventsModule           *AdverseEventsModuleQuery
	withMoreInfoModule                *MoreInfoModuleQuery
	withFKs                           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResultsDefinitionQuery builder.
func (rdq *ResultsDefinitionQuery) Where(ps ...predicate.ResultsDefinition) *ResultsDefinitionQuery {
	rdq.predicates = append(rdq.predicates, ps...)
	return rdq
}

// Limit adds a limit step to the query.
func (rdq *ResultsDefinitionQuery) Limit(limit int) *ResultsDefinitionQuery {
	rdq.limit = &limit
	return rdq
}

// Offset adds an offset step to the query.
func (rdq *ResultsDefinitionQuery) Offset(offset int) *ResultsDefinitionQuery {
	rdq.offset = &offset
	return rdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rdq *ResultsDefinitionQuery) Unique(unique bool) *ResultsDefinitionQuery {
	rdq.unique = &unique
	return rdq
}

// Order adds an order step to the query.
func (rdq *ResultsDefinitionQuery) Order(o ...OrderFunc) *ResultsDefinitionQuery {
	rdq.order = append(rdq.order, o...)
	return rdq
}

// QueryParent chains the current query on the "parent" edge.
func (rdq *ResultsDefinitionQuery) QueryParent() *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: rdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, selector),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, resultsdefinition.ParentTable, resultsdefinition.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParticipantFlowModule chains the current query on the "participant_flow_module" edge.
func (rdq *ResultsDefinitionQuery) QueryParticipantFlowModule() *ParticipantFlowModuleQuery {
	query := &ParticipantFlowModuleQuery{config: rdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, selector),
			sqlgraph.To(participantflowmodule.Table, participantflowmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.ParticipantFlowModuleTable, resultsdefinition.ParticipantFlowModuleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineCharacteristicsModule chains the current query on the "baseline_characteristics_module" edge.
func (rdq *ResultsDefinitionQuery) QueryBaselineCharacteristicsModule() *BaselineCharacteristicsModuleQuery {
	query := &BaselineCharacteristicsModuleQuery{config: rdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, selector),
			sqlgraph.To(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.BaselineCharacteristicsModuleTable, resultsdefinition.BaselineCharacteristicsModuleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutcomeMeasuresModule chains the current query on the "outcome_measures_module" edge.
func (rdq *ResultsDefinitionQuery) QueryOutcomeMeasuresModule() *OutcomeMeasuresModuleQuery {
	query := &OutcomeMeasuresModuleQuery{config: rdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, selector),
			sqlgraph.To(outcomemeasuresmodule.Table, outcomemeasuresmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.OutcomeMeasuresModuleTable, resultsdefinition.OutcomeMeasuresModuleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAdverseEventsModule chains the current query on the "adverse_events_module" edge.
func (rdq *ResultsDefinitionQuery) QueryAdverseEventsModule() *AdverseEventsModuleQuery {
	query := &AdverseEventsModuleQuery{config: rdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, selector),
			sqlgraph.To(adverseeventsmodule.Table, adverseeventsmodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.AdverseEventsModuleTable, resultsdefinition.AdverseEventsModuleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMoreInfoModule chains the current query on the "more_info_module" edge.
func (rdq *ResultsDefinitionQuery) QueryMoreInfoModule() *MoreInfoModuleQuery {
	query := &MoreInfoModuleQuery{config: rdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultsdefinition.Table, resultsdefinition.FieldID, selector),
			sqlgraph.To(moreinfomodule.Table, moreinfomodule.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resultsdefinition.MoreInfoModuleTable, resultsdefinition.MoreInfoModuleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ResultsDefinition entity from the query.
// Returns a *NotFoundError when no ResultsDefinition was found.
func (rdq *ResultsDefinitionQuery) First(ctx context.Context) (*ResultsDefinition, error) {
	nodes, err := rdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resultsdefinition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) FirstX(ctx context.Context) *ResultsDefinition {
	node, err := rdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ResultsDefinition ID from the query.
// Returns a *NotFoundError when no ResultsDefinition ID was found.
func (rdq *ResultsDefinitionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resultsdefinition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) FirstIDX(ctx context.Context) int {
	id, err := rdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ResultsDefinition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ResultsDefinition entity is found.
// Returns a *NotFoundError when no ResultsDefinition entities are found.
func (rdq *ResultsDefinitionQuery) Only(ctx context.Context) (*ResultsDefinition, error) {
	nodes, err := rdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resultsdefinition.Label}
	default:
		return nil, &NotSingularError{resultsdefinition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) OnlyX(ctx context.Context) *ResultsDefinition {
	node, err := rdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ResultsDefinition ID in the query.
// Returns a *NotSingularError when more than one ResultsDefinition ID is found.
// Returns a *NotFoundError when no entities are found.
func (rdq *ResultsDefinitionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resultsdefinition.Label}
	default:
		err = &NotSingularError{resultsdefinition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) OnlyIDX(ctx context.Context) int {
	id, err := rdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ResultsDefinitions.
func (rdq *ResultsDefinitionQuery) All(ctx context.Context) ([]*ResultsDefinition, error) {
	if err := rdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) AllX(ctx context.Context) []*ResultsDefinition {
	nodes, err := rdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ResultsDefinition IDs.
func (rdq *ResultsDefinitionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rdq.Select(resultsdefinition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) IDsX(ctx context.Context) []int {
	ids, err := rdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rdq *ResultsDefinitionQuery) Count(ctx context.Context) (int, error) {
	if err := rdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) CountX(ctx context.Context) int {
	count, err := rdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rdq *ResultsDefinitionQuery) Exist(ctx context.Context) (bool, error) {
	if err := rdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rdq *ResultsDefinitionQuery) ExistX(ctx context.Context) bool {
	exist, err := rdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResultsDefinitionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rdq *ResultsDefinitionQuery) Clone() *ResultsDefinitionQuery {
	if rdq == nil {
		return nil
	}
	return &ResultsDefinitionQuery{
		config:                            rdq.config,
		limit:                             rdq.limit,
		offset:                            rdq.offset,
		order:                             append([]OrderFunc{}, rdq.order...),
		predicates:                        append([]predicate.ResultsDefinition{}, rdq.predicates...),
		withParent:                        rdq.withParent.Clone(),
		withParticipantFlowModule:         rdq.withParticipantFlowModule.Clone(),
		withBaselineCharacteristicsModule: rdq.withBaselineCharacteristicsModule.Clone(),
		withOutcomeMeasuresModule:         rdq.withOutcomeMeasuresModule.Clone(),
		withAdverseEventsModule:           rdq.withAdverseEventsModule.Clone(),
		withMoreInfoModule:                rdq.withMoreInfoModule.Clone(),
		// clone intermediate query.
		sql:    rdq.sql.Clone(),
		path:   rdq.path,
		unique: rdq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResultsDefinitionQuery) WithParent(opts ...func(*ClinicalTrialQuery)) *ResultsDefinitionQuery {
	query := &ClinicalTrialQuery{config: rdq.config}
	for _, opt := range opts {
		opt(query)
	}
	rdq.withParent = query
	return rdq
}

// WithParticipantFlowModule tells the query-builder to eager-load the nodes that are connected to
// the "participant_flow_module" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResultsDefinitionQuery) WithParticipantFlowModule(opts ...func(*ParticipantFlowModuleQuery)) *ResultsDefinitionQuery {
	query := &ParticipantFlowModuleQuery{config: rdq.config}
	for _, opt := range opts {
		opt(query)
	}
	rdq.withParticipantFlowModule = query
	return rdq
}

// WithBaselineCharacteristicsModule tells the query-builder to eager-load the nodes that are connected to
// the "baseline_characteristics_module" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResultsDefinitionQuery) WithBaselineCharacteristicsModule(opts ...func(*BaselineCharacteristicsModuleQuery)) *ResultsDefinitionQuery {
	query := &BaselineCharacteristicsModuleQuery{config: rdq.config}
	for _, opt := range opts {
		opt(query)
	}
	rdq.withBaselineCharacteristicsModule = query
	return rdq
}

// WithOutcomeMeasuresModule tells the query-builder to eager-load the nodes that are connected to
// the "outcome_measures_module" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResultsDefinitionQuery) WithOutcomeMeasuresModule(opts ...func(*OutcomeMeasuresModuleQuery)) *ResultsDefinitionQuery {
	query := &OutcomeMeasuresModuleQuery{config: rdq.config}
	for _, opt := range opts {
		opt(query)
	}
	rdq.withOutcomeMeasuresModule = query
	return rdq
}

// WithAdverseEventsModule tells the query-builder to eager-load the nodes that are connected to
// the "adverse_events_module" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResultsDefinitionQuery) WithAdverseEventsModule(opts ...func(*AdverseEventsModuleQuery)) *ResultsDefinitionQuery {
	query := &AdverseEventsModuleQuery{config: rdq.config}
	for _, opt := range opts {
		opt(query)
	}
	rdq.withAdverseEventsModule = query
	return rdq
}

// WithMoreInfoModule tells the query-builder to eager-load the nodes that are connected to
// the "more_info_module" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResultsDefinitionQuery) WithMoreInfoModule(opts ...func(*MoreInfoModuleQuery)) *ResultsDefinitionQuery {
	query := &MoreInfoModuleQuery{config: rdq.config}
	for _, opt := range opts {
		opt(query)
	}
	rdq.withMoreInfoModule = query
	return rdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (rdq *ResultsDefinitionQuery) GroupBy(field string, fields ...string) *ResultsDefinitionGroupBy {
	grbuild := &ResultsDefinitionGroupBy{config: rdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rdq.sqlQuery(ctx), nil
	}
	grbuild.label = resultsdefinition.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (rdq *ResultsDefinitionQuery) Select(fields ...string) *ResultsDefinitionSelect {
	rdq.fields = append(rdq.fields, fields...)
	selbuild := &ResultsDefinitionSelect{ResultsDefinitionQuery: rdq}
	selbuild.label = resultsdefinition.Label
	selbuild.flds, selbuild.scan = &rdq.fields, selbuild.Scan
	return selbuild
}

func (rdq *ResultsDefinitionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rdq.fields {
		if !resultsdefinition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if rdq.path != nil {
		prev, err := rdq.path(ctx)
		if err != nil {
			return err
		}
		rdq.sql = prev
	}
	return nil
}

func (rdq *ResultsDefinitionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ResultsDefinition, error) {
	var (
		nodes       = []*ResultsDefinition{}
		withFKs     = rdq.withFKs
		_spec       = rdq.querySpec()
		loadedTypes = [6]bool{
			rdq.withParent != nil,
			rdq.withParticipantFlowModule != nil,
			rdq.withBaselineCharacteristicsModule != nil,
			rdq.withOutcomeMeasuresModule != nil,
			rdq.withAdverseEventsModule != nil,
			rdq.withMoreInfoModule != nil,
		}
	)
	if rdq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, resultsdefinition.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ResultsDefinition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ResultsDefinition{config: rdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rdq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ResultsDefinition)
		for i := range nodes {
			if nodes[i].clinical_trial_results_definition == nil {
				continue
			}
			fk := *nodes[i].clinical_trial_results_definition
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(clinicaltrial.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_results_definition" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := rdq.withParticipantFlowModule; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ResultsDefinition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ParticipantFlowModule(func(s *sql.Selector) {
			s.Where(sql.InValues(resultsdefinition.ParticipantFlowModuleColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.results_definition_participant_flow_module
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "results_definition_participant_flow_module" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_participant_flow_module" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ParticipantFlowModule = n
		}
	}

	if query := rdq.withBaselineCharacteristicsModule; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ResultsDefinition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.BaselineCharacteristicsModule(func(s *sql.Selector) {
			s.Where(sql.InValues(resultsdefinition.BaselineCharacteristicsModuleColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.results_definition_baseline_characteristics_module
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "results_definition_baseline_characteristics_module" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_baseline_characteristics_module" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineCharacteristicsModule = n
		}
	}

	if query := rdq.withOutcomeMeasuresModule; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ResultsDefinition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.OutcomeMeasuresModule(func(s *sql.Selector) {
			s.Where(sql.InValues(resultsdefinition.OutcomeMeasuresModuleColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.results_definition_outcome_measures_module
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "results_definition_outcome_measures_module" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_outcome_measures_module" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OutcomeMeasuresModule = n
		}
	}

	if query := rdq.withAdverseEventsModule; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ResultsDefinition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.AdverseEventsModule(func(s *sql.Selector) {
			s.Where(sql.InValues(resultsdefinition.AdverseEventsModuleColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.results_definition_adverse_events_module
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "results_definition_adverse_events_module" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_adverse_events_module" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.AdverseEventsModule = n
		}
	}

	if query := rdq.withMoreInfoModule; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ResultsDefinition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.MoreInfoModule(func(s *sql.Selector) {
			s.Where(sql.InValues(resultsdefinition.MoreInfoModuleColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.results_definition_more_info_module
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "results_definition_more_info_module" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_more_info_module" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MoreInfoModule = n
		}
	}

	return nodes, nil
}

func (rdq *ResultsDefinitionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rdq.querySpec()
	_spec.Node.Columns = rdq.fields
	if len(rdq.fields) > 0 {
		_spec.Unique = rdq.unique != nil && *rdq.unique
	}
	return sqlgraph.CountNodes(ctx, rdq.driver, _spec)
}

func (rdq *ResultsDefinitionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (rdq *ResultsDefinitionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resultsdefinition.Table,
			Columns: resultsdefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultsdefinition.FieldID,
			},
		},
		From:   rdq.sql,
		Unique: true,
	}
	if unique := rdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resultsdefinition.FieldID)
		for i := range fields {
			if fields[i] != resultsdefinition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rdq *ResultsDefinitionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rdq.driver.Dialect())
	t1 := builder.Table(resultsdefinition.Table)
	columns := rdq.fields
	if len(columns) == 0 {
		columns = resultsdefinition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rdq.sql != nil {
		selector = rdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rdq.unique != nil && *rdq.unique {
		selector.Distinct()
	}
	for _, p := range rdq.predicates {
		p(selector)
	}
	for _, p := range rdq.order {
		p(selector)
	}
	if offset := rdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResultsDefinitionGroupBy is the group-by builder for ResultsDefinition entities.
type ResultsDefinitionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rdgb *ResultsDefinitionGroupBy) Aggregate(fns ...AggregateFunc) *ResultsDefinitionGroupBy {
	rdgb.fns = append(rdgb.fns, fns...)
	return rdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rdgb *ResultsDefinitionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rdgb.path(ctx)
	if err != nil {
		return err
	}
	rdgb.sql = query
	return rdgb.sqlScan(ctx, v)
}

func (rdgb *ResultsDefinitionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rdgb.fields {
		if !resultsdefinition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rdgb *ResultsDefinitionGroupBy) sqlQuery() *sql.Selector {
	selector := rdgb.sql.Select()
	aggregation := make([]string, 0, len(rdgb.fns))
	for _, fn := range rdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rdgb.fields)+len(rdgb.fns))
		for _, f := range rdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rdgb.fields...)...)
}

// ResultsDefinitionSelect is the builder for selecting fields of ResultsDefinition entities.
type ResultsDefinitionSelect struct {
	*ResultsDefinitionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rds *ResultsDefinitionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rds.prepareQuery(ctx); err != nil {
		return err
	}
	rds.sql = rds.ResultsDefinitionQuery.sqlQuery(ctx)
	return rds.sqlScan(ctx, v)
}

func (rds *ResultsDefinitionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rds.sql.Query()
	if err := rds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
