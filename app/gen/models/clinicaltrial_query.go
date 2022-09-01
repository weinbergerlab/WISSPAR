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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// ClinicalTrialQuery is the builder for querying ClinicalTrial entities.
type ClinicalTrialQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ClinicalTrial
	// eager-loading edges.
	withResultsDefinition *ResultsDefinitionQuery
	withStudyLocations    *StudyLocationQuery
	withStudyEligibility  *StudyEligibilityQuery
	withMetadata          *ClinicalTrialMetadataQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClinicalTrialQuery builder.
func (ctq *ClinicalTrialQuery) Where(ps ...predicate.ClinicalTrial) *ClinicalTrialQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit adds a limit step to the query.
func (ctq *ClinicalTrialQuery) Limit(limit int) *ClinicalTrialQuery {
	ctq.limit = &limit
	return ctq
}

// Offset adds an offset step to the query.
func (ctq *ClinicalTrialQuery) Offset(offset int) *ClinicalTrialQuery {
	ctq.offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *ClinicalTrialQuery) Unique(unique bool) *ClinicalTrialQuery {
	ctq.unique = &unique
	return ctq
}

// Order adds an order step to the query.
func (ctq *ClinicalTrialQuery) Order(o ...OrderFunc) *ClinicalTrialQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryResultsDefinition chains the current query on the "results_definition" edge.
func (ctq *ClinicalTrialQuery) QueryResultsDefinition() *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: ctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, selector),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, clinicaltrial.ResultsDefinitionTable, clinicaltrial.ResultsDefinitionColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudyLocations chains the current query on the "study_locations" edge.
func (ctq *ClinicalTrialQuery) QueryStudyLocations() *StudyLocationQuery {
	query := &StudyLocationQuery{config: ctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, selector),
			sqlgraph.To(studylocation.Table, studylocation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinicaltrial.StudyLocationsTable, clinicaltrial.StudyLocationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudyEligibility chains the current query on the "study_eligibility" edge.
func (ctq *ClinicalTrialQuery) QueryStudyEligibility() *StudyEligibilityQuery {
	query := &StudyEligibilityQuery{config: ctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, selector),
			sqlgraph.To(studyeligibility.Table, studyeligibility.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinicaltrial.StudyEligibilityTable, clinicaltrial.StudyEligibilityColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMetadata chains the current query on the "metadata" edge.
func (ctq *ClinicalTrialQuery) QueryMetadata() *ClinicalTrialMetadataQuery {
	query := &ClinicalTrialMetadataQuery{config: ctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrial.Table, clinicaltrial.FieldID, selector),
			sqlgraph.To(clinicaltrialmetadata.Table, clinicaltrialmetadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clinicaltrial.MetadataTable, clinicaltrial.MetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClinicalTrial entity from the query.
// Returns a *NotFoundError when no ClinicalTrial was found.
func (ctq *ClinicalTrialQuery) First(ctx context.Context) (*ClinicalTrial, error) {
	nodes, err := ctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{clinicaltrial.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) FirstX(ctx context.Context) *ClinicalTrial {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClinicalTrial ID from the query.
// Returns a *NotFoundError when no ClinicalTrial ID was found.
func (ctq *ClinicalTrialQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{clinicaltrial.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClinicalTrial entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ClinicalTrial entity is found.
// Returns a *NotFoundError when no ClinicalTrial entities are found.
func (ctq *ClinicalTrialQuery) Only(ctx context.Context) (*ClinicalTrial, error) {
	nodes, err := ctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{clinicaltrial.Label}
	default:
		return nil, &NotSingularError{clinicaltrial.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) OnlyX(ctx context.Context) *ClinicalTrial {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClinicalTrial ID in the query.
// Returns a *NotSingularError when more than one ClinicalTrial ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *ClinicalTrialQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{clinicaltrial.Label}
	default:
		err = &NotSingularError{clinicaltrial.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClinicalTrials.
func (ctq *ClinicalTrialQuery) All(ctx context.Context) ([]*ClinicalTrial, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) AllX(ctx context.Context) []*ClinicalTrial {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClinicalTrial IDs.
func (ctq *ClinicalTrialQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctq.Select(clinicaltrial.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *ClinicalTrialQuery) Count(ctx context.Context) (int, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *ClinicalTrialQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *ClinicalTrialQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClinicalTrialQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *ClinicalTrialQuery) Clone() *ClinicalTrialQuery {
	if ctq == nil {
		return nil
	}
	return &ClinicalTrialQuery{
		config:                ctq.config,
		limit:                 ctq.limit,
		offset:                ctq.offset,
		order:                 append([]OrderFunc{}, ctq.order...),
		predicates:            append([]predicate.ClinicalTrial{}, ctq.predicates...),
		withResultsDefinition: ctq.withResultsDefinition.Clone(),
		withStudyLocations:    ctq.withStudyLocations.Clone(),
		withStudyEligibility:  ctq.withStudyEligibility.Clone(),
		withMetadata:          ctq.withMetadata.Clone(),
		// clone intermediate query.
		sql:    ctq.sql.Clone(),
		path:   ctq.path,
		unique: ctq.unique,
	}
}

// WithResultsDefinition tells the query-builder to eager-load the nodes that are connected to
// the "results_definition" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ClinicalTrialQuery) WithResultsDefinition(opts ...func(*ResultsDefinitionQuery)) *ClinicalTrialQuery {
	query := &ResultsDefinitionQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctq.withResultsDefinition = query
	return ctq
}

// WithStudyLocations tells the query-builder to eager-load the nodes that are connected to
// the "study_locations" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ClinicalTrialQuery) WithStudyLocations(opts ...func(*StudyLocationQuery)) *ClinicalTrialQuery {
	query := &StudyLocationQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctq.withStudyLocations = query
	return ctq
}

// WithStudyEligibility tells the query-builder to eager-load the nodes that are connected to
// the "study_eligibility" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ClinicalTrialQuery) WithStudyEligibility(opts ...func(*StudyEligibilityQuery)) *ClinicalTrialQuery {
	query := &StudyEligibilityQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctq.withStudyEligibility = query
	return ctq
}

// WithMetadata tells the query-builder to eager-load the nodes that are connected to
// the "metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ClinicalTrialQuery) WithMetadata(opts ...func(*ClinicalTrialMetadataQuery)) *ClinicalTrialQuery {
	query := &ClinicalTrialMetadataQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctq.withMetadata = query
	return ctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StudyName string `json:"study_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClinicalTrial.Query().
//		GroupBy(clinicaltrial.FieldStudyName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ctq *ClinicalTrialQuery) GroupBy(field string, fields ...string) *ClinicalTrialGroupBy {
	grbuild := &ClinicalTrialGroupBy{config: ctq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctq.sqlQuery(ctx), nil
	}
	grbuild.label = clinicaltrial.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StudyName string `json:"study_name,omitempty"`
//	}
//
//	client.ClinicalTrial.Query().
//		Select(clinicaltrial.FieldStudyName).
//		Scan(ctx, &v)
//
func (ctq *ClinicalTrialQuery) Select(fields ...string) *ClinicalTrialSelect {
	ctq.fields = append(ctq.fields, fields...)
	selbuild := &ClinicalTrialSelect{ClinicalTrialQuery: ctq}
	selbuild.label = clinicaltrial.Label
	selbuild.flds, selbuild.scan = &ctq.fields, selbuild.Scan
	return selbuild
}

func (ctq *ClinicalTrialQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ctq.fields {
		if !clinicaltrial.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *ClinicalTrialQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ClinicalTrial, error) {
	var (
		nodes       = []*ClinicalTrial{}
		_spec       = ctq.querySpec()
		loadedTypes = [4]bool{
			ctq.withResultsDefinition != nil,
			ctq.withStudyLocations != nil,
			ctq.withStudyEligibility != nil,
			ctq.withMetadata != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ClinicalTrial).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ClinicalTrial{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ctq.withResultsDefinition; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ClinicalTrial)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ResultsDefinition(func(s *sql.Selector) {
			s.Where(sql.InValues(clinicaltrial.ResultsDefinitionColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.clinical_trial_results_definition
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "clinical_trial_results_definition" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_results_definition" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ResultsDefinition = n
		}
	}

	if query := ctq.withStudyLocations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ClinicalTrial)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.StudyLocations = []*StudyLocation{}
		}
		query.withFKs = true
		query.Where(predicate.StudyLocation(func(s *sql.Selector) {
			s.Where(sql.InValues(clinicaltrial.StudyLocationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.clinical_trial_study_locations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "clinical_trial_study_locations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_study_locations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.StudyLocations = append(node.Edges.StudyLocations, n)
		}
	}

	if query := ctq.withStudyEligibility; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ClinicalTrial)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.StudyEligibility = []*StudyEligibility{}
		}
		query.withFKs = true
		query.Where(predicate.StudyEligibility(func(s *sql.Selector) {
			s.Where(sql.InValues(clinicaltrial.StudyEligibilityColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.clinical_trial_study_eligibility
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "clinical_trial_study_eligibility" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_study_eligibility" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.StudyEligibility = append(node.Edges.StudyEligibility, n)
		}
	}

	if query := ctq.withMetadata; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ClinicalTrial)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Metadata = []*ClinicalTrialMetadata{}
		}
		query.withFKs = true
		query.Where(predicate.ClinicalTrialMetadata(func(s *sql.Selector) {
			s.Where(sql.InValues(clinicaltrial.MetadataColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.clinical_trial_metadata
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "clinical_trial_metadata" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_metadata" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Metadata = append(node.Edges.Metadata, n)
		}
	}

	return nodes, nil
}

func (ctq *ClinicalTrialQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	_spec.Node.Columns = ctq.fields
	if len(ctq.fields) > 0 {
		_spec.Unique = ctq.unique != nil && *ctq.unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *ClinicalTrialQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ctq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ctq *ClinicalTrialQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicaltrial.Table,
			Columns: clinicaltrial.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrial.FieldID,
			},
		},
		From:   ctq.sql,
		Unique: true,
	}
	if unique := ctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clinicaltrial.FieldID)
		for i := range fields {
			if fields[i] != clinicaltrial.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *ClinicalTrialQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(clinicaltrial.Table)
	columns := ctq.fields
	if len(columns) == 0 {
		columns = clinicaltrial.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.unique != nil && *ctq.unique {
		selector.Distinct()
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClinicalTrialGroupBy is the group-by builder for ClinicalTrial entities.
type ClinicalTrialGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *ClinicalTrialGroupBy) Aggregate(fns ...AggregateFunc) *ClinicalTrialGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ctgb *ClinicalTrialGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ctgb.path(ctx)
	if err != nil {
		return err
	}
	ctgb.sql = query
	return ctgb.sqlScan(ctx, v)
}

func (ctgb *ClinicalTrialGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ctgb.fields {
		if !clinicaltrial.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctgb *ClinicalTrialGroupBy) sqlQuery() *sql.Selector {
	selector := ctgb.sql.Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ctgb.fields)+len(ctgb.fns))
		for _, f := range ctgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ctgb.fields...)...)
}

// ClinicalTrialSelect is the builder for selecting fields of ClinicalTrial entities.
type ClinicalTrialSelect struct {
	*ClinicalTrialQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cts *ClinicalTrialSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	cts.sql = cts.ClinicalTrialQuery.sqlQuery(ctx)
	return cts.sqlScan(ctx, v)
}

func (cts *ClinicalTrialSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cts.sql.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
