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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// ClinicalTrialMetadataQuery is the builder for querying ClinicalTrialMetadata entities.
type ClinicalTrialMetadataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ClinicalTrialMetadata
	// eager-loading edges.
	withParent *ClinicalTrialQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClinicalTrialMetadataQuery builder.
func (ctmq *ClinicalTrialMetadataQuery) Where(ps ...predicate.ClinicalTrialMetadata) *ClinicalTrialMetadataQuery {
	ctmq.predicates = append(ctmq.predicates, ps...)
	return ctmq
}

// Limit adds a limit step to the query.
func (ctmq *ClinicalTrialMetadataQuery) Limit(limit int) *ClinicalTrialMetadataQuery {
	ctmq.limit = &limit
	return ctmq
}

// Offset adds an offset step to the query.
func (ctmq *ClinicalTrialMetadataQuery) Offset(offset int) *ClinicalTrialMetadataQuery {
	ctmq.offset = &offset
	return ctmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctmq *ClinicalTrialMetadataQuery) Unique(unique bool) *ClinicalTrialMetadataQuery {
	ctmq.unique = &unique
	return ctmq
}

// Order adds an order step to the query.
func (ctmq *ClinicalTrialMetadataQuery) Order(o ...OrderFunc) *ClinicalTrialMetadataQuery {
	ctmq.order = append(ctmq.order, o...)
	return ctmq
}

// QueryParent chains the current query on the "parent" edge.
func (ctmq *ClinicalTrialMetadataQuery) QueryParent() *ClinicalTrialQuery {
	query := &ClinicalTrialQuery{config: ctmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clinicaltrialmetadata.Table, clinicaltrialmetadata.FieldID, selector),
			sqlgraph.To(clinicaltrial.Table, clinicaltrial.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, clinicaltrialmetadata.ParentTable, clinicaltrialmetadata.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClinicalTrialMetadata entity from the query.
// Returns a *NotFoundError when no ClinicalTrialMetadata was found.
func (ctmq *ClinicalTrialMetadataQuery) First(ctx context.Context) (*ClinicalTrialMetadata, error) {
	nodes, err := ctmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{clinicaltrialmetadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) FirstX(ctx context.Context) *ClinicalTrialMetadata {
	node, err := ctmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClinicalTrialMetadata ID from the query.
// Returns a *NotFoundError when no ClinicalTrialMetadata ID was found.
func (ctmq *ClinicalTrialMetadataQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{clinicaltrialmetadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) FirstIDX(ctx context.Context) int {
	id, err := ctmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClinicalTrialMetadata entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ClinicalTrialMetadata entity is found.
// Returns a *NotFoundError when no ClinicalTrialMetadata entities are found.
func (ctmq *ClinicalTrialMetadataQuery) Only(ctx context.Context) (*ClinicalTrialMetadata, error) {
	nodes, err := ctmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{clinicaltrialmetadata.Label}
	default:
		return nil, &NotSingularError{clinicaltrialmetadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) OnlyX(ctx context.Context) *ClinicalTrialMetadata {
	node, err := ctmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClinicalTrialMetadata ID in the query.
// Returns a *NotSingularError when more than one ClinicalTrialMetadata ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctmq *ClinicalTrialMetadataQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{clinicaltrialmetadata.Label}
	default:
		err = &NotSingularError{clinicaltrialmetadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClinicalTrialMetadataSlice.
func (ctmq *ClinicalTrialMetadataQuery) All(ctx context.Context) ([]*ClinicalTrialMetadata, error) {
	if err := ctmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) AllX(ctx context.Context) []*ClinicalTrialMetadata {
	nodes, err := ctmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClinicalTrialMetadata IDs.
func (ctmq *ClinicalTrialMetadataQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctmq.Select(clinicaltrialmetadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) IDsX(ctx context.Context) []int {
	ids, err := ctmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctmq *ClinicalTrialMetadataQuery) Count(ctx context.Context) (int, error) {
	if err := ctmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) CountX(ctx context.Context) int {
	count, err := ctmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctmq *ClinicalTrialMetadataQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctmq *ClinicalTrialMetadataQuery) ExistX(ctx context.Context) bool {
	exist, err := ctmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClinicalTrialMetadataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctmq *ClinicalTrialMetadataQuery) Clone() *ClinicalTrialMetadataQuery {
	if ctmq == nil {
		return nil
	}
	return &ClinicalTrialMetadataQuery{
		config:     ctmq.config,
		limit:      ctmq.limit,
		offset:     ctmq.offset,
		order:      append([]OrderFunc{}, ctmq.order...),
		predicates: append([]predicate.ClinicalTrialMetadata{}, ctmq.predicates...),
		withParent: ctmq.withParent.Clone(),
		// clone intermediate query.
		sql:    ctmq.sql.Clone(),
		path:   ctmq.path,
		unique: ctmq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ctmq *ClinicalTrialMetadataQuery) WithParent(opts ...func(*ClinicalTrialQuery)) *ClinicalTrialMetadataQuery {
	query := &ClinicalTrialQuery{config: ctmq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctmq.withParent = query
	return ctmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TagName string `json:"tag_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClinicalTrialMetadata.Query().
//		GroupBy(clinicaltrialmetadata.FieldTagName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (ctmq *ClinicalTrialMetadataQuery) GroupBy(field string, fields ...string) *ClinicalTrialMetadataGroupBy {
	grbuild := &ClinicalTrialMetadataGroupBy{config: ctmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctmq.sqlQuery(ctx), nil
	}
	grbuild.label = clinicaltrialmetadata.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TagName string `json:"tag_name,omitempty"`
//	}
//
//	client.ClinicalTrialMetadata.Query().
//		Select(clinicaltrialmetadata.FieldTagName).
//		Scan(ctx, &v)
//
func (ctmq *ClinicalTrialMetadataQuery) Select(fields ...string) *ClinicalTrialMetadataSelect {
	ctmq.fields = append(ctmq.fields, fields...)
	selbuild := &ClinicalTrialMetadataSelect{ClinicalTrialMetadataQuery: ctmq}
	selbuild.label = clinicaltrialmetadata.Label
	selbuild.flds, selbuild.scan = &ctmq.fields, selbuild.Scan
	return selbuild
}

func (ctmq *ClinicalTrialMetadataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ctmq.fields {
		if !clinicaltrialmetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if ctmq.path != nil {
		prev, err := ctmq.path(ctx)
		if err != nil {
			return err
		}
		ctmq.sql = prev
	}
	return nil
}

func (ctmq *ClinicalTrialMetadataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ClinicalTrialMetadata, error) {
	var (
		nodes       = []*ClinicalTrialMetadata{}
		withFKs     = ctmq.withFKs
		_spec       = ctmq.querySpec()
		loadedTypes = [1]bool{
			ctmq.withParent != nil,
		}
	)
	if ctmq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, clinicaltrialmetadata.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ClinicalTrialMetadata).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ClinicalTrialMetadata{config: ctmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ctmq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ClinicalTrialMetadata)
		for i := range nodes {
			if nodes[i].clinical_trial_metadata == nil {
				continue
			}
			fk := *nodes[i].clinical_trial_metadata
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
				return nil, fmt.Errorf(`unexpected foreign-key "clinical_trial_metadata" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (ctmq *ClinicalTrialMetadataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctmq.querySpec()
	_spec.Node.Columns = ctmq.fields
	if len(ctmq.fields) > 0 {
		_spec.Unique = ctmq.unique != nil && *ctmq.unique
	}
	return sqlgraph.CountNodes(ctx, ctmq.driver, _spec)
}

func (ctmq *ClinicalTrialMetadataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ctmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (ctmq *ClinicalTrialMetadataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicaltrialmetadata.Table,
			Columns: clinicaltrialmetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrialmetadata.FieldID,
			},
		},
		From:   ctmq.sql,
		Unique: true,
	}
	if unique := ctmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clinicaltrialmetadata.FieldID)
		for i := range fields {
			if fields[i] != clinicaltrialmetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctmq *ClinicalTrialMetadataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctmq.driver.Dialect())
	t1 := builder.Table(clinicaltrialmetadata.Table)
	columns := ctmq.fields
	if len(columns) == 0 {
		columns = clinicaltrialmetadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctmq.sql != nil {
		selector = ctmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctmq.unique != nil && *ctmq.unique {
		selector.Distinct()
	}
	for _, p := range ctmq.predicates {
		p(selector)
	}
	for _, p := range ctmq.order {
		p(selector)
	}
	if offset := ctmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClinicalTrialMetadataGroupBy is the group-by builder for ClinicalTrialMetadata entities.
type ClinicalTrialMetadataGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctmgb *ClinicalTrialMetadataGroupBy) Aggregate(fns ...AggregateFunc) *ClinicalTrialMetadataGroupBy {
	ctmgb.fns = append(ctmgb.fns, fns...)
	return ctmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ctmgb *ClinicalTrialMetadataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ctmgb.path(ctx)
	if err != nil {
		return err
	}
	ctmgb.sql = query
	return ctmgb.sqlScan(ctx, v)
}

func (ctmgb *ClinicalTrialMetadataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ctmgb.fields {
		if !clinicaltrialmetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctmgb *ClinicalTrialMetadataGroupBy) sqlQuery() *sql.Selector {
	selector := ctmgb.sql.Select()
	aggregation := make([]string, 0, len(ctmgb.fns))
	for _, fn := range ctmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ctmgb.fields)+len(ctmgb.fns))
		for _, f := range ctmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ctmgb.fields...)...)
}

// ClinicalTrialMetadataSelect is the builder for selecting fields of ClinicalTrialMetadata entities.
type ClinicalTrialMetadataSelect struct {
	*ClinicalTrialMetadataQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ctms *ClinicalTrialMetadataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ctms.prepareQuery(ctx); err != nil {
		return err
	}
	ctms.sql = ctms.ClinicalTrialMetadataQuery.sqlQuery(ctx)
	return ctms.sqlScan(ctx, v)
}

func (ctms *ClinicalTrialMetadataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ctms.sql.Query()
	if err := ctms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
