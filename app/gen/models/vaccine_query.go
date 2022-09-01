// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/vaccine"
)

// VaccineQuery is the builder for querying Vaccine entities.
type VaccineQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Vaccine
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VaccineQuery builder.
func (vq *VaccineQuery) Where(ps ...predicate.Vaccine) *VaccineQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VaccineQuery) Limit(limit int) *VaccineQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VaccineQuery) Offset(offset int) *VaccineQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VaccineQuery) Unique(unique bool) *VaccineQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *VaccineQuery) Order(o ...OrderFunc) *VaccineQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// First returns the first Vaccine entity from the query.
// Returns a *NotFoundError when no Vaccine was found.
func (vq *VaccineQuery) First(ctx context.Context) (*Vaccine, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vaccine.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VaccineQuery) FirstX(ctx context.Context) *Vaccine {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vaccine ID from the query.
// Returns a *NotFoundError when no Vaccine ID was found.
func (vq *VaccineQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vaccine.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VaccineQuery) FirstIDX(ctx context.Context) int {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vaccine entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vaccine entity is found.
// Returns a *NotFoundError when no Vaccine entities are found.
func (vq *VaccineQuery) Only(ctx context.Context) (*Vaccine, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vaccine.Label}
	default:
		return nil, &NotSingularError{vaccine.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VaccineQuery) OnlyX(ctx context.Context) *Vaccine {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vaccine ID in the query.
// Returns a *NotSingularError when more than one Vaccine ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VaccineQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vaccine.Label}
	default:
		err = &NotSingularError{vaccine.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VaccineQuery) OnlyIDX(ctx context.Context) int {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vaccines.
func (vq *VaccineQuery) All(ctx context.Context) ([]*Vaccine, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VaccineQuery) AllX(ctx context.Context) []*Vaccine {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vaccine IDs.
func (vq *VaccineQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := vq.Select(vaccine.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VaccineQuery) IDsX(ctx context.Context) []int {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VaccineQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VaccineQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VaccineQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VaccineQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VaccineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VaccineQuery) Clone() *VaccineQuery {
	if vq == nil {
		return nil
	}
	return &VaccineQuery{
		config:     vq.config,
		limit:      vq.limit,
		offset:     vq.offset,
		order:      append([]OrderFunc{}, vq.order...),
		predicates: append([]predicate.Vaccine{}, vq.predicates...),
		// clone intermediate query.
		sql:    vq.sql.Clone(),
		path:   vq.path,
		unique: vq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VaccineName string `json:"vaccine_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vaccine.Query().
//		GroupBy(vaccine.FieldVaccineName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (vq *VaccineQuery) GroupBy(field string, fields ...string) *VaccineGroupBy {
	grbuild := &VaccineGroupBy{config: vq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	grbuild.label = vaccine.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VaccineName string `json:"vaccine_name,omitempty"`
//	}
//
//	client.Vaccine.Query().
//		Select(vaccine.FieldVaccineName).
//		Scan(ctx, &v)
//
func (vq *VaccineQuery) Select(fields ...string) *VaccineSelect {
	vq.fields = append(vq.fields, fields...)
	selbuild := &VaccineSelect{VaccineQuery: vq}
	selbuild.label = vaccine.Label
	selbuild.flds, selbuild.scan = &vq.fields, selbuild.Scan
	return selbuild
}

func (vq *VaccineQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !vaccine.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VaccineQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vaccine, error) {
	var (
		nodes = []*Vaccine{}
		_spec = vq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Vaccine).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Vaccine{config: vq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vq *VaccineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.fields
	if len(vq.fields) > 0 {
		_spec.Unique = vq.unique != nil && *vq.unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VaccineQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (vq *VaccineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vaccine.Table,
			Columns: vaccine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vaccine.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vaccine.FieldID)
		for i := range fields {
			if fields[i] != vaccine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VaccineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vaccine.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = vaccine.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.unique != nil && *vq.unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VaccineGroupBy is the group-by builder for Vaccine entities.
type VaccineGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VaccineGroupBy) Aggregate(fns ...AggregateFunc) *VaccineGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VaccineGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

func (vgb *VaccineGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vgb.fields {
		if !vaccine.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VaccineGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}

// VaccineSelect is the builder for selecting fields of Vaccine entities.
type VaccineSelect struct {
	*VaccineQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VaccineSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VaccineQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

func (vs *VaccineSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
