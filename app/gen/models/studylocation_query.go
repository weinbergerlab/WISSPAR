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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// StudyLocationQuery is the builder for querying StudyLocation entities.
type StudyLocationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StudyLocation
	// eager-loading edges.
	withParent *ClinicalTrialQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StudyLocationQuery builder.
func (slq *StudyLocationQuery) Where(ps ...predicate.StudyLocation) *StudyLocationQuery {
	slq.predicates = append(slq.predicates, ps...)
	return slq
}

// Limit adds a limit step to the query.
func (slq *StudyLocationQuery) Limit(limit int) *StudyLocationQuery {
	slq.limit = &limit
	return slq
}

// Offset adds an offset step to the query.
func (slq *StudyLocationQuery) Offset(offset int) *StudyLocationQuery {
	slq.offset = &offset
	return slq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (slq *StudyLocationQuery) Unique(unique bool) *StudyLocationQuery {
	slq.unique = &unique
	return slq
}

// Order adds an order step to the query.
func (slq *StudyLocationQuery) Order(o ...OrderFunc) *StudyLocationQuery {
	slq.order = append(slq.order, o...)
	return slq
}

// QueryParent chains the current query on the "parent" edge.
func (slq *StudyLocationQuery) QueryParent() *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: slq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := slq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := slq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studylocation.Table, studylocation.FieldID, selector),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, studylocation.ParentTable, studylocation.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(slq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StudyLocation entity from the query.
// Returns a *NotFoundError when no StudyLocation was found.
func (slq *StudyLocationQuery) First(ctx context.Context) (*StudyLocation, error) {
	nodes, err := slq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{studylocation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (slq *StudyLocationQuery) FirstX(ctx context.Context) *StudyLocation {
	node, err := slq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StudyLocation ID from the query.
// Returns a *NotFoundError when no StudyLocation ID was found.
func (slq *StudyLocationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = slq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{studylocation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (slq *StudyLocationQuery) FirstIDX(ctx context.Context) int {
	id, err := slq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StudyLocation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StudyLocation entity is found.
// Returns a *NotFoundError when no StudyLocation entities are found.
func (slq *StudyLocationQuery) Only(ctx context.Context) (*StudyLocation, error) {
	nodes, err := slq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{studylocation.Label}
	default:
		return nil, &NotSingularError{studylocation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (slq *StudyLocationQuery) OnlyX(ctx context.Context) *StudyLocation {
	node, err := slq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StudyLocation ID in the query.
// Returns a *NotSingularError when more than one StudyLocation ID is found.
// Returns a *NotFoundError when no entities are found.
func (slq *StudyLocationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = slq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{studylocation.Label}
	default:
		err = &NotSingularError{studylocation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (slq *StudyLocationQuery) OnlyIDX(ctx context.Context) int {
	id, err := slq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StudyLocations.
func (slq *StudyLocationQuery) All(ctx context.Context) ([]*StudyLocation, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return slq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (slq *StudyLocationQuery) AllX(ctx context.Context) []*StudyLocation {
	nodes, err := slq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StudyLocation IDs.
func (slq *StudyLocationQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := slq.Select(studylocation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (slq *StudyLocationQuery) IDsX(ctx context.Context) []int {
	ids, err := slq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (slq *StudyLocationQuery) Count(ctx context.Context) (int, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return slq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (slq *StudyLocationQuery) CountX(ctx context.Context) int {
	count, err := slq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (slq *StudyLocationQuery) Exist(ctx context.Context) (bool, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return slq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (slq *StudyLocationQuery) ExistX(ctx context.Context) bool {
	exist, err := slq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StudyLocationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (slq *StudyLocationQuery) Clone() *StudyLocationQuery {
	if slq == nil {
		return nil
	}
	return &StudyLocationQuery{
		config:     slq.config,
		limit:      slq.limit,
		offset:     slq.offset,
		order:      append([]OrderFunc{}, slq.order...),
		predicates: append([]predicate.StudyLocation{}, slq.predicates...),
		withParent: slq.withParent.Clone(),
		// clone intermediate query.
		sql:    slq.sql.Clone(),
		path:   slq.path,
		unique: slq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (slq *StudyLocationQuery) WithParent(opts ...func(*ClinicalTrialQuery)) *StudyLocationQuery {
	query := &ClinicalTrialQuery{config: slq.config}
	for _, opt := range opts {
		opt(query)
	}
	slq.withParent = query
	return slq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LocationFacility string `json:"LocationFacility,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StudyLocation.Query().
//		GroupBy(studylocation.FieldLocationFacility).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (slq *StudyLocationQuery) GroupBy(field string, fields ...string) *StudyLocationGroupBy {
	grbuild := &StudyLocationGroupBy{config: slq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := slq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return slq.sqlQuery(ctx), nil
	}
	grbuild.label = studylocation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LocationFacility string `json:"LocationFacility,omitempty"`
//	}
//
//	client.StudyLocation.Query().
//		Select(studylocation.FieldLocationFacility).
//		Scan(ctx, &v)
//
func (slq *StudyLocationQuery) Select(fields ...string) *StudyLocationSelect {
	slq.fields = append(slq.fields, fields...)
	selbuild := &StudyLocationSelect{StudyLocationQuery: slq}
	selbuild.label = studylocation.Label
	selbuild.flds, selbuild.scan = &slq.fields, selbuild.Scan
	return selbuild
}

func (slq *StudyLocationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range slq.fields {
		if !studylocation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if slq.path != nil {
		prev, err := slq.path(ctx)
		if err != nil {
			return err
		}
		slq.sql = prev
	}
	return nil
}

func (slq *StudyLocationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StudyLocation, error) {
	var (
		nodes       = []*StudyLocation{}
		withFKs     = slq.withFKs
		_spec       = slq.querySpec()
		loadedTypes = [1]bool{
			slq.withParent != nil,
		}
	)
	if slq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, studylocation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*StudyLocation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &StudyLocation{config: slq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, slq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := slq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*StudyLocation)
		for i := range nodes {
			if nodes[i].clinical_trial_study_locations == nil {
				continue
			}
			fk := *nodes[i].clinical_trial_study_locations
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
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_study_locations" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (slq *StudyLocationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := slq.querySpec()
	_spec.Node.Columns = slq.fields
	if len(slq.fields) > 0 {
		_spec.Unique = slq.unique != nil && *slq.unique
	}
	return sqlgraph.CountNodes(ctx, slq.driver, _spec)
}

func (slq *StudyLocationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := slq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (slq *StudyLocationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studylocation.Table,
			Columns: studylocation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studylocation.FieldID,
			},
		},
		From:   slq.sql,
		Unique: true,
	}
	if unique := slq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := slq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studylocation.FieldID)
		for i := range fields {
			if fields[i] != studylocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := slq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := slq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := slq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := slq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (slq *StudyLocationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(slq.driver.Dialect())
	t1 := builder.Table(studylocation.Table)
	columns := slq.fields
	if len(columns) == 0 {
		columns = studylocation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if slq.sql != nil {
		selector = slq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if slq.unique != nil && *slq.unique {
		selector.Distinct()
	}
	for _, p := range slq.predicates {
		p(selector)
	}
	for _, p := range slq.order {
		p(selector)
	}
	if offset := slq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := slq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StudyLocationGroupBy is the group-by builder for StudyLocation entities.
type StudyLocationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (slgb *StudyLocationGroupBy) Aggregate(fns ...AggregateFunc) *StudyLocationGroupBy {
	slgb.fns = append(slgb.fns, fns...)
	return slgb
}

// Scan applies the group-by query and scans the result into the given value.
func (slgb *StudyLocationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := slgb.path(ctx)
	if err != nil {
		return err
	}
	slgb.sql = query
	return slgb.sqlScan(ctx, v)
}

func (slgb *StudyLocationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range slgb.fields {
		if !studylocation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := slgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := slgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (slgb *StudyLocationGroupBy) sqlQuery() *sql.Selector {
	selector := slgb.sql.Select()
	aggregation := make([]string, 0, len(slgb.fns))
	for _, fn := range slgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(slgb.fields)+len(slgb.fns))
		for _, f := range slgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(slgb.fields...)...)
}

// StudyLocationSelect is the builder for selecting fields of StudyLocation entities.
type StudyLocationSelect struct {
	*StudyLocationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sls *StudyLocationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sls.prepareQuery(ctx); err != nil {
		return err
	}
	sls.sql = sls.StudyLocationQuery.sqlQuery(ctx)
	return sls.sqlScan(ctx, v)
}

func (sls *StudyLocationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sls.sql.Query()
	if err := sls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
