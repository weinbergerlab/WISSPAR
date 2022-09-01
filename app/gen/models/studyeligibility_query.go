// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
)

// StudyEligibilityQuery is the builder for querying StudyEligibility entities.
type StudyEligibilityQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StudyEligibility
	// eager-loading edges.
	withParent *ClinicalTrialQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StudyEligibilityQuery builder.
func (seq *StudyEligibilityQuery) Where(ps ...predicate.StudyEligibility) *StudyEligibilityQuery {
	seq.predicates = append(seq.predicates, ps...)
	return seq
}

// Limit adds a limit step to the query.
func (seq *StudyEligibilityQuery) Limit(limit int) *StudyEligibilityQuery {
	seq.limit = &limit
	return seq
}

// Offset adds an offset step to the query.
func (seq *StudyEligibilityQuery) Offset(offset int) *StudyEligibilityQuery {
	seq.offset = &offset
	return seq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (seq *StudyEligibilityQuery) Unique(unique bool) *StudyEligibilityQuery {
	seq.unique = &unique
	return seq
}

// Order adds an order step to the query.
func (seq *StudyEligibilityQuery) Order(o ...OrderFunc) *StudyEligibilityQuery {
	seq.order = append(seq.order, o...)
	return seq
}

// QueryParent chains the current query on the "parent" edge.
func (seq *StudyEligibilityQuery) QueryParent() *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: seq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := seq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studyeligibility.Table, studyeligibility.FieldID, selector),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, studyeligibility.ParentTable, studyeligibility.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(seq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StudyEligibility entity from the query.
// Returns a *NotFoundError when no StudyEligibility was found.
func (seq *StudyEligibilityQuery) First(ctx context.Context) (*StudyEligibility, error) {
	nodes, err := seq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{studyeligibility.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (seq *StudyEligibilityQuery) FirstX(ctx context.Context) *StudyEligibility {
	node, err := seq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StudyEligibility ID from the query.
// Returns a *NotFoundError when no StudyEligibility ID was found.
func (seq *StudyEligibilityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{studyeligibility.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (seq *StudyEligibilityQuery) FirstIDX(ctx context.Context) int {
	id, err := seq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StudyEligibility entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StudyEligibility entity is found.
// Returns a *NotFoundError when no StudyEligibility entities are found.
func (seq *StudyEligibilityQuery) Only(ctx context.Context) (*StudyEligibility, error) {
	nodes, err := seq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{studyeligibility.Label}
	default:
		return nil, &NotSingularError{studyeligibility.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (seq *StudyEligibilityQuery) OnlyX(ctx context.Context) *StudyEligibility {
	node, err := seq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StudyEligibility ID in the query.
// Returns a *NotSingularError when more than one StudyEligibility ID is found.
// Returns a *NotFoundError when no entities are found.
func (seq *StudyEligibilityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{studyeligibility.Label}
	default:
		err = &NotSingularError{studyeligibility.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (seq *StudyEligibilityQuery) OnlyIDX(ctx context.Context) int {
	id, err := seq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StudyEligibilities.
func (seq *StudyEligibilityQuery) All(ctx context.Context) ([]*StudyEligibility, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return seq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (seq *StudyEligibilityQuery) AllX(ctx context.Context) []*StudyEligibility {
	nodes, err := seq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StudyEligibility IDs.
func (seq *StudyEligibilityQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := seq.Select(studyeligibility.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (seq *StudyEligibilityQuery) IDsX(ctx context.Context) []int {
	ids, err := seq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (seq *StudyEligibilityQuery) Count(ctx context.Context) (int, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return seq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (seq *StudyEligibilityQuery) CountX(ctx context.Context) int {
	count, err := seq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (seq *StudyEligibilityQuery) Exist(ctx context.Context) (bool, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return seq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (seq *StudyEligibilityQuery) ExistX(ctx context.Context) bool {
	exist, err := seq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StudyEligibilityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (seq *StudyEligibilityQuery) Clone() *StudyEligibilityQuery {
	if seq == nil {
		return nil
	}
	return &StudyEligibilityQuery{
		config:     seq.config,
		limit:      seq.limit,
		offset:     seq.offset,
		order:      append([]OrderFunc{}, seq.order...),
		predicates: append([]predicate.StudyEligibility{}, seq.predicates...),
		withParent: seq.withParent.Clone(),
		// clone intermediate query.
		sql:    seq.sql.Clone(),
		path:   seq.path,
		unique: seq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (seq *StudyEligibilityQuery) WithParent(opts ...func(*ClinicalTrialQuery)) *StudyEligibilityQuery {
	query := &ClinicalTrialQuery{config: seq.config}
	for _, opt := range opts {
		opt(query)
	}
	seq.withParent = query
	return seq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EligibilityCriteria string `json:"EligibilityCriteria,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StudyEligibility.Query().
//		GroupBy(studyeligibility.FieldEligibilityCriteria).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (seq *StudyEligibilityQuery) GroupBy(field string, fields ...string) *StudyEligibilityGroupBy {
	grbuild := &StudyEligibilityGroupBy{config: seq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return seq.sqlQuery(ctx), nil
	}
	grbuild.label = studyeligibility.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EligibilityCriteria string `json:"EligibilityCriteria,omitempty"`
//	}
//
//	client.StudyEligibility.Query().
//		Select(studyeligibility.FieldEligibilityCriteria).
//		Scan(ctx, &v)
//
func (seq *StudyEligibilityQuery) Select(fields ...string) *StudyEligibilitySelect {
	seq.fields = append(seq.fields, fields...)
	selbuild := &StudyEligibilitySelect{StudyEligibilityQuery: seq}
	selbuild.label = studyeligibility.Label
	selbuild.flds, selbuild.scan = &seq.fields, selbuild.Scan
	return selbuild
}

func (seq *StudyEligibilityQuery) prepareQuery(ctx context.Context) error {
	for _, f := range seq.fields {
		if !studyeligibility.ValidColumn(f) {
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

func (seq *StudyEligibilityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StudyEligibility, error) {
	var (
		nodes       = []*StudyEligibility{}
		withFKs     = seq.withFKs
		_spec       = seq.querySpec()
		loadedTypes = [1]bool{
			seq.withParent != nil,
		}
	)
	if seq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, studyeligibility.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*StudyEligibility).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &StudyEligibility{config: seq.config}
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
		nodeids := make(map[int][]*StudyEligibility)
		for i := range nodes {
			if nodes[i].clinical_trial_study_eligibility == nil {
				continue
			}
			fk := *nodes[i].clinical_trial_study_eligibility
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
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_study_eligibility" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (seq *StudyEligibilityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := seq.querySpec()
	_spec.Node.Columns = seq.fields
	if len(seq.fields) > 0 {
		_spec.Unique = seq.unique != nil && *seq.unique
	}
	return sqlgraph.CountNodes(ctx, seq.driver, _spec)
}

func (seq *StudyEligibilityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := seq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (seq *StudyEligibilityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studyeligibility.Table,
			Columns: studyeligibility.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyeligibility.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, studyeligibility.FieldID)
		for i := range fields {
			if fields[i] != studyeligibility.FieldID {
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

func (seq *StudyEligibilityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(seq.driver.Dialect())
	t1 := builder.Table(studyeligibility.Table)
	columns := seq.fields
	if len(columns) == 0 {
		columns = studyeligibility.Columns
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

// StudyEligibilityGroupBy is the group-by builder for StudyEligibility entities.
type StudyEligibilityGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (segb *StudyEligibilityGroupBy) Aggregate(fns ...AggregateFunc) *StudyEligibilityGroupBy {
	segb.fns = append(segb.fns, fns...)
	return segb
}

// Scan applies the group-by query and scans the result into the given value.
func (segb *StudyEligibilityGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := segb.path(ctx)
	if err != nil {
		return err
	}
	segb.sql = query
	return segb.sqlScan(ctx, v)
}

func (segb *StudyEligibilityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range segb.fields {
		if !studyeligibility.ValidColumn(f) {
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

func (segb *StudyEligibilityGroupBy) sqlQuery() *sql.Selector {
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

// StudyEligibilitySelect is the builder for selecting fields of StudyEligibility entities.
type StudyEligibilitySelect struct {
	*StudyEligibilityQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ses *StudyEligibilitySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ses.prepareQuery(ctx); err != nil {
		return err
	}
	ses.sql = ses.StudyEligibilityQuery.sqlQuery(ctx)
	return ses.sqlScan(ctx, v)
}

func (ses *StudyEligibilitySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ses.sql.Query()
	if err := ses.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
