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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinedenom"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinegroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// BaselineCharacteristicsModuleQuery is the builder for querying BaselineCharacteristicsModule entities.
type BaselineCharacteristicsModuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaselineCharacteristicsModule
	// eager-loading edges.
	withParent              *ResultsDefinitionQuery
	withBaselineGroupList   *BaselineGroupQuery
	withBaselineDenomList   *BaselineDenomQuery
	withBaselineMeasureList *BaselineMeasureQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaselineCharacteristicsModuleQuery builder.
func (bcmq *BaselineCharacteristicsModuleQuery) Where(ps ...predicate.BaselineCharacteristicsModule) *BaselineCharacteristicsModuleQuery {
	bcmq.predicates = append(bcmq.predicates, ps...)
	return bcmq
}

// Limit adds a limit step to the query.
func (bcmq *BaselineCharacteristicsModuleQuery) Limit(limit int) *BaselineCharacteristicsModuleQuery {
	bcmq.limit = &limit
	return bcmq
}

// Offset adds an offset step to the query.
func (bcmq *BaselineCharacteristicsModuleQuery) Offset(offset int) *BaselineCharacteristicsModuleQuery {
	bcmq.offset = &offset
	return bcmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcmq *BaselineCharacteristicsModuleQuery) Unique(unique bool) *BaselineCharacteristicsModuleQuery {
	bcmq.unique = &unique
	return bcmq
}

// Order adds an order step to the query.
func (bcmq *BaselineCharacteristicsModuleQuery) Order(o ...OrderFunc) *BaselineCharacteristicsModuleQuery {
	bcmq.order = append(bcmq.order, o...)
	return bcmq
}

// QueryParent chains the current query on the "parent" edge.
func (bcmq *BaselineCharacteristicsModuleQuery) QueryParent() *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: bcmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, selector),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, baselinecharacteristicsmodule.ParentTable, baselinecharacteristicsmodule.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineGroupList chains the current query on the "baseline_group_list" edge.
func (bcmq *BaselineCharacteristicsModuleQuery) QueryBaselineGroupList() *BaselineGroupQuery {
	query := &BaselineGroupQuery{config: bcmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, selector),
			sqlgraph.To(baselinegroup.Table, baselinegroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecharacteristicsmodule.BaselineGroupListTable, baselinecharacteristicsmodule.BaselineGroupListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineDenomList chains the current query on the "baseline_denom_list" edge.
func (bcmq *BaselineCharacteristicsModuleQuery) QueryBaselineDenomList() *BaselineDenomQuery {
	query := &BaselineDenomQuery{config: bcmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, selector),
			sqlgraph.To(baselinedenom.Table, baselinedenom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecharacteristicsmodule.BaselineDenomListTable, baselinecharacteristicsmodule.BaselineDenomListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBaselineMeasureList chains the current query on the "baseline_measure_list" edge.
func (bcmq *BaselineCharacteristicsModuleQuery) QueryBaselineMeasureList() *BaselineMeasureQuery {
	query := &BaselineMeasureQuery{config: bcmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(baselinecharacteristicsmodule.Table, baselinecharacteristicsmodule.FieldID, selector),
			sqlgraph.To(baselinemeasure.Table, baselinemeasure.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, baselinecharacteristicsmodule.BaselineMeasureListTable, baselinecharacteristicsmodule.BaselineMeasureListColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BaselineCharacteristicsModule entity from the query.
// Returns a *NotFoundError when no BaselineCharacteristicsModule was found.
func (bcmq *BaselineCharacteristicsModuleQuery) First(ctx context.Context) (*BaselineCharacteristicsModule, error) {
	nodes, err := bcmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{baselinecharacteristicsmodule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) FirstX(ctx context.Context) *BaselineCharacteristicsModule {
	node, err := bcmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaselineCharacteristicsModule ID from the query.
// Returns a *NotFoundError when no BaselineCharacteristicsModule ID was found.
func (bcmq *BaselineCharacteristicsModuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{baselinecharacteristicsmodule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) FirstIDX(ctx context.Context) int {
	id, err := bcmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaselineCharacteristicsModule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaselineCharacteristicsModule entity is found.
// Returns a *NotFoundError when no BaselineCharacteristicsModule entities are found.
func (bcmq *BaselineCharacteristicsModuleQuery) Only(ctx context.Context) (*BaselineCharacteristicsModule, error) {
	nodes, err := bcmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{baselinecharacteristicsmodule.Label}
	default:
		return nil, &NotSingularError{baselinecharacteristicsmodule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) OnlyX(ctx context.Context) *BaselineCharacteristicsModule {
	node, err := bcmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaselineCharacteristicsModule ID in the query.
// Returns a *NotSingularError when more than one BaselineCharacteristicsModule ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcmq *BaselineCharacteristicsModuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{baselinecharacteristicsmodule.Label}
	default:
		err = &NotSingularError{baselinecharacteristicsmodule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaselineCharacteristicsModules.
func (bcmq *BaselineCharacteristicsModuleQuery) All(ctx context.Context) ([]*BaselineCharacteristicsModule, error) {
	if err := bcmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) AllX(ctx context.Context) []*BaselineCharacteristicsModule {
	nodes, err := bcmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaselineCharacteristicsModule IDs.
func (bcmq *BaselineCharacteristicsModuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bcmq.Select(baselinecharacteristicsmodule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) IDsX(ctx context.Context) []int {
	ids, err := bcmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcmq *BaselineCharacteristicsModuleQuery) Count(ctx context.Context) (int, error) {
	if err := bcmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) CountX(ctx context.Context) int {
	count, err := bcmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcmq *BaselineCharacteristicsModuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcmq *BaselineCharacteristicsModuleQuery) ExistX(ctx context.Context) bool {
	exist, err := bcmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaselineCharacteristicsModuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcmq *BaselineCharacteristicsModuleQuery) Clone() *BaselineCharacteristicsModuleQuery {
	if bcmq == nil {
		return nil
	}
	return &BaselineCharacteristicsModuleQuery{
		config:                  bcmq.config,
		limit:                   bcmq.limit,
		offset:                  bcmq.offset,
		order:                   append([]OrderFunc{}, bcmq.order...),
		predicates:              append([]predicate.BaselineCharacteristicsModule{}, bcmq.predicates...),
		withParent:              bcmq.withParent.Clone(),
		withBaselineGroupList:   bcmq.withBaselineGroupList.Clone(),
		withBaselineDenomList:   bcmq.withBaselineDenomList.Clone(),
		withBaselineMeasureList: bcmq.withBaselineMeasureList.Clone(),
		// clone intermediate query.
		sql:    bcmq.sql.Clone(),
		path:   bcmq.path,
		unique: bcmq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (bcmq *BaselineCharacteristicsModuleQuery) WithParent(opts ...func(*ResultsDefinitionQuery)) *BaselineCharacteristicsModuleQuery {
	query := &ResultsDefinitionQuery{config: bcmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcmq.withParent = query
	return bcmq
}

// WithBaselineGroupList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_group_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcmq *BaselineCharacteristicsModuleQuery) WithBaselineGroupList(opts ...func(*BaselineGroupQuery)) *BaselineCharacteristicsModuleQuery {
	query := &BaselineGroupQuery{config: bcmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcmq.withBaselineGroupList = query
	return bcmq
}

// WithBaselineDenomList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_denom_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcmq *BaselineCharacteristicsModuleQuery) WithBaselineDenomList(opts ...func(*BaselineDenomQuery)) *BaselineCharacteristicsModuleQuery {
	query := &BaselineDenomQuery{config: bcmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcmq.withBaselineDenomList = query
	return bcmq
}

// WithBaselineMeasureList tells the query-builder to eager-load the nodes that are connected to
// the "baseline_measure_list" edge. The optional arguments are used to configure the query builder of the edge.
func (bcmq *BaselineCharacteristicsModuleQuery) WithBaselineMeasureList(opts ...func(*BaselineMeasureQuery)) *BaselineCharacteristicsModuleQuery {
	query := &BaselineMeasureQuery{config: bcmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcmq.withBaselineMeasureList = query
	return bcmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BaselinePopulationDescription string `json:"baseline_population_description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaselineCharacteristicsModule.Query().
//		GroupBy(baselinecharacteristicsmodule.FieldBaselinePopulationDescription).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (bcmq *BaselineCharacteristicsModuleQuery) GroupBy(field string, fields ...string) *BaselineCharacteristicsModuleGroupBy {
	grbuild := &BaselineCharacteristicsModuleGroupBy{config: bcmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcmq.sqlQuery(ctx), nil
	}
	grbuild.label = baselinecharacteristicsmodule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BaselinePopulationDescription string `json:"baseline_population_description,omitempty"`
//	}
//
//	client.BaselineCharacteristicsModule.Query().
//		Select(baselinecharacteristicsmodule.FieldBaselinePopulationDescription).
//		Scan(ctx, &v)
//
func (bcmq *BaselineCharacteristicsModuleQuery) Select(fields ...string) *BaselineCharacteristicsModuleSelect {
	bcmq.fields = append(bcmq.fields, fields...)
	selbuild := &BaselineCharacteristicsModuleSelect{BaselineCharacteristicsModuleQuery: bcmq}
	selbuild.label = baselinecharacteristicsmodule.Label
	selbuild.flds, selbuild.scan = &bcmq.fields, selbuild.Scan
	return selbuild
}

func (bcmq *BaselineCharacteristicsModuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcmq.fields {
		if !baselinecharacteristicsmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if bcmq.path != nil {
		prev, err := bcmq.path(ctx)
		if err != nil {
			return err
		}
		bcmq.sql = prev
	}
	return nil
}

func (bcmq *BaselineCharacteristicsModuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaselineCharacteristicsModule, error) {
	var (
		nodes       = []*BaselineCharacteristicsModule{}
		withFKs     = bcmq.withFKs
		_spec       = bcmq.querySpec()
		loadedTypes = [4]bool{
			bcmq.withParent != nil,
			bcmq.withBaselineGroupList != nil,
			bcmq.withBaselineDenomList != nil,
			bcmq.withBaselineMeasureList != nil,
		}
	)
	if bcmq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, baselinecharacteristicsmodule.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaselineCharacteristicsModule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaselineCharacteristicsModule{config: bcmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bcmq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BaselineCharacteristicsModule)
		for i := range nodes {
			if nodes[i].results_definition_baseline_characteristics_module == nil {
				continue
			}
			fk := *nodes[i].results_definition_baseline_characteristics_module
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(resultsdefinition.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_baseline_characteristics_module" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := bcmq.withBaselineGroupList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineCharacteristicsModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineGroupList = []*BaselineGroup{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineGroup(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinecharacteristicsmodule.BaselineGroupListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_characteristics_module_baseline_group_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_characteristics_module_baseline_group_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_characteristics_module_baseline_group_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineGroupList = append(node.Edges.BaselineGroupList, n)
		}
	}

	if query := bcmq.withBaselineDenomList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineCharacteristicsModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineDenomList = []*BaselineDenom{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineDenom(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinecharacteristicsmodule.BaselineDenomListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_characteristics_module_baseline_denom_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_characteristics_module_baseline_denom_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_characteristics_module_baseline_denom_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineDenomList = append(node.Edges.BaselineDenomList, n)
		}
	}

	if query := bcmq.withBaselineMeasureList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*BaselineCharacteristicsModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BaselineMeasureList = []*BaselineMeasure{}
		}
		query.withFKs = true
		query.Where(predicate.BaselineMeasure(func(s *sql.Selector) {
			s.Where(sql.InValues(baselinecharacteristicsmodule.BaselineMeasureListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.baseline_characteristics_module_baseline_measure_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "baseline_characteristics_module_baseline_measure_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "baseline_characteristics_module_baseline_measure_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BaselineMeasureList = append(node.Edges.BaselineMeasureList, n)
		}
	}

	return nodes, nil
}

func (bcmq *BaselineCharacteristicsModuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcmq.querySpec()
	_spec.Node.Columns = bcmq.fields
	if len(bcmq.fields) > 0 {
		_spec.Unique = bcmq.unique != nil && *bcmq.unique
	}
	return sqlgraph.CountNodes(ctx, bcmq.driver, _spec)
}

func (bcmq *BaselineCharacteristicsModuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bcmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (bcmq *BaselineCharacteristicsModuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   baselinecharacteristicsmodule.Table,
			Columns: baselinecharacteristicsmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: baselinecharacteristicsmodule.FieldID,
			},
		},
		From:   bcmq.sql,
		Unique: true,
	}
	if unique := bcmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, baselinecharacteristicsmodule.FieldID)
		for i := range fields {
			if fields[i] != baselinecharacteristicsmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcmq *BaselineCharacteristicsModuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcmq.driver.Dialect())
	t1 := builder.Table(baselinecharacteristicsmodule.Table)
	columns := bcmq.fields
	if len(columns) == 0 {
		columns = baselinecharacteristicsmodule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcmq.sql != nil {
		selector = bcmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcmq.unique != nil && *bcmq.unique {
		selector.Distinct()
	}
	for _, p := range bcmq.predicates {
		p(selector)
	}
	for _, p := range bcmq.order {
		p(selector)
	}
	if offset := bcmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaselineCharacteristicsModuleGroupBy is the group-by builder for BaselineCharacteristicsModule entities.
type BaselineCharacteristicsModuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcmgb *BaselineCharacteristicsModuleGroupBy) Aggregate(fns ...AggregateFunc) *BaselineCharacteristicsModuleGroupBy {
	bcmgb.fns = append(bcmgb.fns, fns...)
	return bcmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcmgb *BaselineCharacteristicsModuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bcmgb.path(ctx)
	if err != nil {
		return err
	}
	bcmgb.sql = query
	return bcmgb.sqlScan(ctx, v)
}

func (bcmgb *BaselineCharacteristicsModuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bcmgb.fields {
		if !baselinecharacteristicsmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcmgb *BaselineCharacteristicsModuleGroupBy) sqlQuery() *sql.Selector {
	selector := bcmgb.sql.Select()
	aggregation := make([]string, 0, len(bcmgb.fns))
	for _, fn := range bcmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcmgb.fields)+len(bcmgb.fns))
		for _, f := range bcmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcmgb.fields...)...)
}

// BaselineCharacteristicsModuleSelect is the builder for selecting fields of BaselineCharacteristicsModule entities.
type BaselineCharacteristicsModuleSelect struct {
	*BaselineCharacteristicsModuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcms *BaselineCharacteristicsModuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bcms.prepareQuery(ctx); err != nil {
		return err
	}
	bcms.sql = bcms.BaselineCharacteristicsModuleQuery.sqlQuery(ctx)
	return bcms.sqlScan(ctx, v)
}

func (bcms *BaselineCharacteristicsModuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bcms.sql.Query()
	if err := bcms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
