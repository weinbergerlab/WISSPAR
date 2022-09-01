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
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/eventgroup"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/otherevent"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/seriousevent"
)

// AdverseEventsModuleQuery is the builder for querying AdverseEventsModule entities.
type AdverseEventsModuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AdverseEventsModule
	// eager-loading edges.
	withParent           *ResultsDefinitionQuery
	withEventGroupList   *EventGroupQuery
	withSeriousEventList *SeriousEventQuery
	withOtherEventList   *OtherEventQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdverseEventsModuleQuery builder.
func (aemq *AdverseEventsModuleQuery) Where(ps ...predicate.AdverseEventsModule) *AdverseEventsModuleQuery {
	aemq.predicates = append(aemq.predicates, ps...)
	return aemq
}

// Limit adds a limit step to the query.
func (aemq *AdverseEventsModuleQuery) Limit(limit int) *AdverseEventsModuleQuery {
	aemq.limit = &limit
	return aemq
}

// Offset adds an offset step to the query.
func (aemq *AdverseEventsModuleQuery) Offset(offset int) *AdverseEventsModuleQuery {
	aemq.offset = &offset
	return aemq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aemq *AdverseEventsModuleQuery) Unique(unique bool) *AdverseEventsModuleQuery {
	aemq.unique = &unique
	return aemq
}

// Order adds an order step to the query.
func (aemq *AdverseEventsModuleQuery) Order(o ...OrderFunc) *AdverseEventsModuleQuery {
	aemq.order = append(aemq.order, o...)
	return aemq
}

// QueryParent chains the current query on the "parent" edge.
func (aemq *AdverseEventsModuleQuery) QueryParent() *ResultsDefinitionQuery {
	query := &ResultsDefinitionQuery{config: aemq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aemq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aemq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, selector),
			sqlgraph.To(resultsdefinition.Table, resultsdefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, adverseeventsmodule.ParentTable, adverseeventsmodule.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(aemq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEventGroupList chains the current query on the "event_group_list" edge.
func (aemq *AdverseEventsModuleQuery) QueryEventGroupList() *EventGroupQuery {
	query := &EventGroupQuery{config: aemq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aemq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aemq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, selector),
			sqlgraph.To(eventgroup.Table, eventgroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adverseeventsmodule.EventGroupListTable, adverseeventsmodule.EventGroupListColumn),
		)
		fromU = sqlgraph.SetNeighbors(aemq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySeriousEventList chains the current query on the "serious_event_list" edge.
func (aemq *AdverseEventsModuleQuery) QuerySeriousEventList() *SeriousEventQuery {
	query := &SeriousEventQuery{config: aemq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aemq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aemq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, selector),
			sqlgraph.To(seriousevent.Table, seriousevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adverseeventsmodule.SeriousEventListTable, adverseeventsmodule.SeriousEventListColumn),
		)
		fromU = sqlgraph.SetNeighbors(aemq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOtherEventList chains the current query on the "other_event_list" edge.
func (aemq *AdverseEventsModuleQuery) QueryOtherEventList() *OtherEventQuery {
	query := &OtherEventQuery{config: aemq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aemq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aemq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adverseeventsmodule.Table, adverseeventsmodule.FieldID, selector),
			sqlgraph.To(otherevent.Table, otherevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adverseeventsmodule.OtherEventListTable, adverseeventsmodule.OtherEventListColumn),
		)
		fromU = sqlgraph.SetNeighbors(aemq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdverseEventsModule entity from the query.
// Returns a *NotFoundError when no AdverseEventsModule was found.
func (aemq *AdverseEventsModuleQuery) First(ctx context.Context) (*AdverseEventsModule, error) {
	nodes, err := aemq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adverseeventsmodule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) FirstX(ctx context.Context) *AdverseEventsModule {
	node, err := aemq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdverseEventsModule ID from the query.
// Returns a *NotFoundError when no AdverseEventsModule ID was found.
func (aemq *AdverseEventsModuleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aemq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adverseeventsmodule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) FirstIDX(ctx context.Context) int {
	id, err := aemq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdverseEventsModule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdverseEventsModule entity is found.
// Returns a *NotFoundError when no AdverseEventsModule entities are found.
func (aemq *AdverseEventsModuleQuery) Only(ctx context.Context) (*AdverseEventsModule, error) {
	nodes, err := aemq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adverseeventsmodule.Label}
	default:
		return nil, &NotSingularError{adverseeventsmodule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) OnlyX(ctx context.Context) *AdverseEventsModule {
	node, err := aemq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdverseEventsModule ID in the query.
// Returns a *NotSingularError when more than one AdverseEventsModule ID is found.
// Returns a *NotFoundError when no entities are found.
func (aemq *AdverseEventsModuleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aemq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adverseeventsmodule.Label}
	default:
		err = &NotSingularError{adverseeventsmodule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) OnlyIDX(ctx context.Context) int {
	id, err := aemq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdverseEventsModules.
func (aemq *AdverseEventsModuleQuery) All(ctx context.Context) ([]*AdverseEventsModule, error) {
	if err := aemq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aemq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) AllX(ctx context.Context) []*AdverseEventsModule {
	nodes, err := aemq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdverseEventsModule IDs.
func (aemq *AdverseEventsModuleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aemq.Select(adverseeventsmodule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) IDsX(ctx context.Context) []int {
	ids, err := aemq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aemq *AdverseEventsModuleQuery) Count(ctx context.Context) (int, error) {
	if err := aemq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aemq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) CountX(ctx context.Context) int {
	count, err := aemq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aemq *AdverseEventsModuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := aemq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aemq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aemq *AdverseEventsModuleQuery) ExistX(ctx context.Context) bool {
	exist, err := aemq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdverseEventsModuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aemq *AdverseEventsModuleQuery) Clone() *AdverseEventsModuleQuery {
	if aemq == nil {
		return nil
	}
	return &AdverseEventsModuleQuery{
		config:               aemq.config,
		limit:                aemq.limit,
		offset:               aemq.offset,
		order:                append([]OrderFunc{}, aemq.order...),
		predicates:           append([]predicate.AdverseEventsModule{}, aemq.predicates...),
		withParent:           aemq.withParent.Clone(),
		withEventGroupList:   aemq.withEventGroupList.Clone(),
		withSeriousEventList: aemq.withSeriousEventList.Clone(),
		withOtherEventList:   aemq.withOtherEventList.Clone(),
		// clone intermediate query.
		sql:    aemq.sql.Clone(),
		path:   aemq.path,
		unique: aemq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (aemq *AdverseEventsModuleQuery) WithParent(opts ...func(*ResultsDefinitionQuery)) *AdverseEventsModuleQuery {
	query := &ResultsDefinitionQuery{config: aemq.config}
	for _, opt := range opts {
		opt(query)
	}
	aemq.withParent = query
	return aemq
}

// WithEventGroupList tells the query-builder to eager-load the nodes that are connected to
// the "event_group_list" edge. The optional arguments are used to configure the query builder of the edge.
func (aemq *AdverseEventsModuleQuery) WithEventGroupList(opts ...func(*EventGroupQuery)) *AdverseEventsModuleQuery {
	query := &EventGroupQuery{config: aemq.config}
	for _, opt := range opts {
		opt(query)
	}
	aemq.withEventGroupList = query
	return aemq
}

// WithSeriousEventList tells the query-builder to eager-load the nodes that are connected to
// the "serious_event_list" edge. The optional arguments are used to configure the query builder of the edge.
func (aemq *AdverseEventsModuleQuery) WithSeriousEventList(opts ...func(*SeriousEventQuery)) *AdverseEventsModuleQuery {
	query := &SeriousEventQuery{config: aemq.config}
	for _, opt := range opts {
		opt(query)
	}
	aemq.withSeriousEventList = query
	return aemq
}

// WithOtherEventList tells the query-builder to eager-load the nodes that are connected to
// the "other_event_list" edge. The optional arguments are used to configure the query builder of the edge.
func (aemq *AdverseEventsModuleQuery) WithOtherEventList(opts ...func(*OtherEventQuery)) *AdverseEventsModuleQuery {
	query := &OtherEventQuery{config: aemq.config}
	for _, opt := range opts {
		opt(query)
	}
	aemq.withOtherEventList = query
	return aemq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventsFrequencyThreshold string `json:"events_frequency_threshold,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdverseEventsModule.Query().
//		GroupBy(adverseeventsmodule.FieldEventsFrequencyThreshold).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (aemq *AdverseEventsModuleQuery) GroupBy(field string, fields ...string) *AdverseEventsModuleGroupBy {
	grbuild := &AdverseEventsModuleGroupBy{config: aemq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aemq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aemq.sqlQuery(ctx), nil
	}
	grbuild.label = adverseeventsmodule.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventsFrequencyThreshold string `json:"events_frequency_threshold,omitempty"`
//	}
//
//	client.AdverseEventsModule.Query().
//		Select(adverseeventsmodule.FieldEventsFrequencyThreshold).
//		Scan(ctx, &v)
//
func (aemq *AdverseEventsModuleQuery) Select(fields ...string) *AdverseEventsModuleSelect {
	aemq.fields = append(aemq.fields, fields...)
	selbuild := &AdverseEventsModuleSelect{AdverseEventsModuleQuery: aemq}
	selbuild.label = adverseeventsmodule.Label
	selbuild.flds, selbuild.scan = &aemq.fields, selbuild.Scan
	return selbuild
}

func (aemq *AdverseEventsModuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aemq.fields {
		if !adverseeventsmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if aemq.path != nil {
		prev, err := aemq.path(ctx)
		if err != nil {
			return err
		}
		aemq.sql = prev
	}
	return nil
}

func (aemq *AdverseEventsModuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdverseEventsModule, error) {
	var (
		nodes       = []*AdverseEventsModule{}
		withFKs     = aemq.withFKs
		_spec       = aemq.querySpec()
		loadedTypes = [4]bool{
			aemq.withParent != nil,
			aemq.withEventGroupList != nil,
			aemq.withSeriousEventList != nil,
			aemq.withOtherEventList != nil,
		}
	)
	if aemq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, adverseeventsmodule.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AdverseEventsModule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AdverseEventsModule{config: aemq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aemq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aemq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*AdverseEventsModule)
		for i := range nodes {
			if nodes[i].results_definition_adverse_events_module == nil {
				continue
			}
			fk := *nodes[i].results_definition_adverse_events_module
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
				return nil, fmt.Errorf(`unexpected foreign-key "results_definition_adverse_events_module" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := aemq.withEventGroupList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*AdverseEventsModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.EventGroupList = []*EventGroup{}
		}
		query.withFKs = true
		query.Where(predicate.EventGroup(func(s *sql.Selector) {
			s.Where(sql.InValues(adverseeventsmodule.EventGroupListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.adverse_events_module_event_group_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "adverse_events_module_event_group_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "adverse_events_module_event_group_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.EventGroupList = append(node.Edges.EventGroupList, n)
		}
	}

	if query := aemq.withSeriousEventList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*AdverseEventsModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.SeriousEventList = []*SeriousEvent{}
		}
		query.withFKs = true
		query.Where(predicate.SeriousEvent(func(s *sql.Selector) {
			s.Where(sql.InValues(adverseeventsmodule.SeriousEventListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.adverse_events_module_serious_event_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "adverse_events_module_serious_event_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "adverse_events_module_serious_event_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.SeriousEventList = append(node.Edges.SeriousEventList, n)
		}
	}

	if query := aemq.withOtherEventList; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*AdverseEventsModule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OtherEventList = []*OtherEvent{}
		}
		query.withFKs = true
		query.Where(predicate.OtherEvent(func(s *sql.Selector) {
			s.Where(sql.InValues(adverseeventsmodule.OtherEventListColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.adverse_events_module_other_event_list
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "adverse_events_module_other_event_list" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "adverse_events_module_other_event_list" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OtherEventList = append(node.Edges.OtherEventList, n)
		}
	}

	return nodes, nil
}

func (aemq *AdverseEventsModuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aemq.querySpec()
	_spec.Node.Columns = aemq.fields
	if len(aemq.fields) > 0 {
		_spec.Unique = aemq.unique != nil && *aemq.unique
	}
	return sqlgraph.CountNodes(ctx, aemq.driver, _spec)
}

func (aemq *AdverseEventsModuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aemq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %w", err)
	}
	return n > 0, nil
}

func (aemq *AdverseEventsModuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adverseeventsmodule.Table,
			Columns: adverseeventsmodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adverseeventsmodule.FieldID,
			},
		},
		From:   aemq.sql,
		Unique: true,
	}
	if unique := aemq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aemq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adverseeventsmodule.FieldID)
		for i := range fields {
			if fields[i] != adverseeventsmodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aemq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aemq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aemq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aemq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aemq *AdverseEventsModuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aemq.driver.Dialect())
	t1 := builder.Table(adverseeventsmodule.Table)
	columns := aemq.fields
	if len(columns) == 0 {
		columns = adverseeventsmodule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aemq.sql != nil {
		selector = aemq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aemq.unique != nil && *aemq.unique {
		selector.Distinct()
	}
	for _, p := range aemq.predicates {
		p(selector)
	}
	for _, p := range aemq.order {
		p(selector)
	}
	if offset := aemq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aemq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdverseEventsModuleGroupBy is the group-by builder for AdverseEventsModule entities.
type AdverseEventsModuleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aemgb *AdverseEventsModuleGroupBy) Aggregate(fns ...AggregateFunc) *AdverseEventsModuleGroupBy {
	aemgb.fns = append(aemgb.fns, fns...)
	return aemgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aemgb *AdverseEventsModuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aemgb.path(ctx)
	if err != nil {
		return err
	}
	aemgb.sql = query
	return aemgb.sqlScan(ctx, v)
}

func (aemgb *AdverseEventsModuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aemgb.fields {
		if !adverseeventsmodule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aemgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aemgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aemgb *AdverseEventsModuleGroupBy) sqlQuery() *sql.Selector {
	selector := aemgb.sql.Select()
	aggregation := make([]string, 0, len(aemgb.fns))
	for _, fn := range aemgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aemgb.fields)+len(aemgb.fns))
		for _, f := range aemgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aemgb.fields...)...)
}

// AdverseEventsModuleSelect is the builder for selecting fields of AdverseEventsModule entities.
type AdverseEventsModuleSelect struct {
	*AdverseEventsModuleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aems *AdverseEventsModuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aems.prepareQuery(ctx); err != nil {
		return err
	}
	aems.sql = aems.AdverseEventsModuleQuery.sqlQuery(ctx)
	return aems.sqlScan(ctx, v)
}

func (aems *AdverseEventsModuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aems.sql.Query()
	if err := aems.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
