// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

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

// ResultsDefinitionUpdate is the builder for updating ResultsDefinition entities.
type ResultsDefinitionUpdate struct {
	config
	hooks    []Hook
	mutation *ResultsDefinitionMutation
}

// Where appends a list predicates to the ResultsDefinitionUpdate builder.
func (rdu *ResultsDefinitionUpdate) Where(ps ...predicate.ResultsDefinition) *ResultsDefinitionUpdate {
	rdu.mutation.Where(ps...)
	return rdu
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (rdu *ResultsDefinitionUpdate) SetParentID(id int) *ResultsDefinitionUpdate {
	rdu.mutation.SetParentID(id)
	return rdu
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (rdu *ResultsDefinitionUpdate) SetNillableParentID(id *int) *ResultsDefinitionUpdate {
	if id != nil {
		rdu = rdu.SetParentID(*id)
	}
	return rdu
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (rdu *ResultsDefinitionUpdate) SetParent(c *ClinicalTrial) *ResultsDefinitionUpdate {
	return rdu.SetParentID(c.ID)
}

// SetParticipantFlowModuleID sets the "participant_flow_module" edge to the ParticipantFlowModule entity by ID.
func (rdu *ResultsDefinitionUpdate) SetParticipantFlowModuleID(id int) *ResultsDefinitionUpdate {
	rdu.mutation.SetParticipantFlowModuleID(id)
	return rdu
}

// SetNillableParticipantFlowModuleID sets the "participant_flow_module" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (rdu *ResultsDefinitionUpdate) SetNillableParticipantFlowModuleID(id *int) *ResultsDefinitionUpdate {
	if id != nil {
		rdu = rdu.SetParticipantFlowModuleID(*id)
	}
	return rdu
}

// SetParticipantFlowModule sets the "participant_flow_module" edge to the ParticipantFlowModule entity.
func (rdu *ResultsDefinitionUpdate) SetParticipantFlowModule(p *ParticipantFlowModule) *ResultsDefinitionUpdate {
	return rdu.SetParticipantFlowModuleID(p.ID)
}

// SetBaselineCharacteristicsModuleID sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity by ID.
func (rdu *ResultsDefinitionUpdate) SetBaselineCharacteristicsModuleID(id int) *ResultsDefinitionUpdate {
	rdu.mutation.SetBaselineCharacteristicsModuleID(id)
	return rdu
}

// SetNillableBaselineCharacteristicsModuleID sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (rdu *ResultsDefinitionUpdate) SetNillableBaselineCharacteristicsModuleID(id *int) *ResultsDefinitionUpdate {
	if id != nil {
		rdu = rdu.SetBaselineCharacteristicsModuleID(*id)
	}
	return rdu
}

// SetBaselineCharacteristicsModule sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity.
func (rdu *ResultsDefinitionUpdate) SetBaselineCharacteristicsModule(b *BaselineCharacteristicsModule) *ResultsDefinitionUpdate {
	return rdu.SetBaselineCharacteristicsModuleID(b.ID)
}

// SetOutcomeMeasuresModuleID sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity by ID.
func (rdu *ResultsDefinitionUpdate) SetOutcomeMeasuresModuleID(id int) *ResultsDefinitionUpdate {
	rdu.mutation.SetOutcomeMeasuresModuleID(id)
	return rdu
}

// SetNillableOutcomeMeasuresModuleID sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity by ID if the given value is not nil.
func (rdu *ResultsDefinitionUpdate) SetNillableOutcomeMeasuresModuleID(id *int) *ResultsDefinitionUpdate {
	if id != nil {
		rdu = rdu.SetOutcomeMeasuresModuleID(*id)
	}
	return rdu
}

// SetOutcomeMeasuresModule sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity.
func (rdu *ResultsDefinitionUpdate) SetOutcomeMeasuresModule(o *OutcomeMeasuresModule) *ResultsDefinitionUpdate {
	return rdu.SetOutcomeMeasuresModuleID(o.ID)
}

// SetAdverseEventsModuleID sets the "adverse_events_module" edge to the AdverseEventsModule entity by ID.
func (rdu *ResultsDefinitionUpdate) SetAdverseEventsModuleID(id int) *ResultsDefinitionUpdate {
	rdu.mutation.SetAdverseEventsModuleID(id)
	return rdu
}

// SetNillableAdverseEventsModuleID sets the "adverse_events_module" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (rdu *ResultsDefinitionUpdate) SetNillableAdverseEventsModuleID(id *int) *ResultsDefinitionUpdate {
	if id != nil {
		rdu = rdu.SetAdverseEventsModuleID(*id)
	}
	return rdu
}

// SetAdverseEventsModule sets the "adverse_events_module" edge to the AdverseEventsModule entity.
func (rdu *ResultsDefinitionUpdate) SetAdverseEventsModule(a *AdverseEventsModule) *ResultsDefinitionUpdate {
	return rdu.SetAdverseEventsModuleID(a.ID)
}

// SetMoreInfoModuleID sets the "more_info_module" edge to the MoreInfoModule entity by ID.
func (rdu *ResultsDefinitionUpdate) SetMoreInfoModuleID(id int) *ResultsDefinitionUpdate {
	rdu.mutation.SetMoreInfoModuleID(id)
	return rdu
}

// SetNillableMoreInfoModuleID sets the "more_info_module" edge to the MoreInfoModule entity by ID if the given value is not nil.
func (rdu *ResultsDefinitionUpdate) SetNillableMoreInfoModuleID(id *int) *ResultsDefinitionUpdate {
	if id != nil {
		rdu = rdu.SetMoreInfoModuleID(*id)
	}
	return rdu
}

// SetMoreInfoModule sets the "more_info_module" edge to the MoreInfoModule entity.
func (rdu *ResultsDefinitionUpdate) SetMoreInfoModule(m *MoreInfoModule) *ResultsDefinitionUpdate {
	return rdu.SetMoreInfoModuleID(m.ID)
}

// Mutation returns the ResultsDefinitionMutation object of the builder.
func (rdu *ResultsDefinitionUpdate) Mutation() *ResultsDefinitionMutation {
	return rdu.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (rdu *ResultsDefinitionUpdate) ClearParent() *ResultsDefinitionUpdate {
	rdu.mutation.ClearParent()
	return rdu
}

// ClearParticipantFlowModule clears the "participant_flow_module" edge to the ParticipantFlowModule entity.
func (rdu *ResultsDefinitionUpdate) ClearParticipantFlowModule() *ResultsDefinitionUpdate {
	rdu.mutation.ClearParticipantFlowModule()
	return rdu
}

// ClearBaselineCharacteristicsModule clears the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity.
func (rdu *ResultsDefinitionUpdate) ClearBaselineCharacteristicsModule() *ResultsDefinitionUpdate {
	rdu.mutation.ClearBaselineCharacteristicsModule()
	return rdu
}

// ClearOutcomeMeasuresModule clears the "outcome_measures_module" edge to the OutcomeMeasuresModule entity.
func (rdu *ResultsDefinitionUpdate) ClearOutcomeMeasuresModule() *ResultsDefinitionUpdate {
	rdu.mutation.ClearOutcomeMeasuresModule()
	return rdu
}

// ClearAdverseEventsModule clears the "adverse_events_module" edge to the AdverseEventsModule entity.
func (rdu *ResultsDefinitionUpdate) ClearAdverseEventsModule() *ResultsDefinitionUpdate {
	rdu.mutation.ClearAdverseEventsModule()
	return rdu
}

// ClearMoreInfoModule clears the "more_info_module" edge to the MoreInfoModule entity.
func (rdu *ResultsDefinitionUpdate) ClearMoreInfoModule() *ResultsDefinitionUpdate {
	rdu.mutation.ClearMoreInfoModule()
	return rdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rdu *ResultsDefinitionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rdu.hooks) == 0 {
		affected, err = rdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultsDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rdu.mutation = mutation
			affected, err = rdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rdu.hooks) - 1; i >= 0; i-- {
			if rdu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = rdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rdu *ResultsDefinitionUpdate) SaveX(ctx context.Context) int {
	affected, err := rdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rdu *ResultsDefinitionUpdate) Exec(ctx context.Context) error {
	_, err := rdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdu *ResultsDefinitionUpdate) ExecX(ctx context.Context) {
	if err := rdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rdu *ResultsDefinitionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resultsdefinition.Table,
			Columns: resultsdefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultsdefinition.FieldID,
			},
		},
	}
	if ps := rdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rdu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resultsdefinition.ParentTable,
			Columns: []string{resultsdefinition.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resultsdefinition.ParentTable,
			Columns: []string{resultsdefinition.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdu.mutation.ParticipantFlowModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.ParticipantFlowModuleTable,
			Columns: []string{resultsdefinition.ParticipantFlowModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdu.mutation.ParticipantFlowModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.ParticipantFlowModuleTable,
			Columns: []string{resultsdefinition.ParticipantFlowModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdu.mutation.BaselineCharacteristicsModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.BaselineCharacteristicsModuleTable,
			Columns: []string{resultsdefinition.BaselineCharacteristicsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdu.mutation.BaselineCharacteristicsModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.BaselineCharacteristicsModuleTable,
			Columns: []string{resultsdefinition.BaselineCharacteristicsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdu.mutation.OutcomeMeasuresModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.OutcomeMeasuresModuleTable,
			Columns: []string{resultsdefinition.OutcomeMeasuresModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasuresmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdu.mutation.OutcomeMeasuresModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.OutcomeMeasuresModuleTable,
			Columns: []string{resultsdefinition.OutcomeMeasuresModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasuresmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdu.mutation.AdverseEventsModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.AdverseEventsModuleTable,
			Columns: []string{resultsdefinition.AdverseEventsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adverseeventsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdu.mutation.AdverseEventsModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.AdverseEventsModuleTable,
			Columns: []string{resultsdefinition.AdverseEventsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adverseeventsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdu.mutation.MoreInfoModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.MoreInfoModuleTable,
			Columns: []string{resultsdefinition.MoreInfoModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdu.mutation.MoreInfoModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.MoreInfoModuleTable,
			Columns: []string{resultsdefinition.MoreInfoModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resultsdefinition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ResultsDefinitionUpdateOne is the builder for updating a single ResultsDefinition entity.
type ResultsDefinitionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResultsDefinitionMutation
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (rduo *ResultsDefinitionUpdateOne) SetParentID(id int) *ResultsDefinitionUpdateOne {
	rduo.mutation.SetParentID(id)
	return rduo
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (rduo *ResultsDefinitionUpdateOne) SetNillableParentID(id *int) *ResultsDefinitionUpdateOne {
	if id != nil {
		rduo = rduo.SetParentID(*id)
	}
	return rduo
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (rduo *ResultsDefinitionUpdateOne) SetParent(c *ClinicalTrial) *ResultsDefinitionUpdateOne {
	return rduo.SetParentID(c.ID)
}

// SetParticipantFlowModuleID sets the "participant_flow_module" edge to the ParticipantFlowModule entity by ID.
func (rduo *ResultsDefinitionUpdateOne) SetParticipantFlowModuleID(id int) *ResultsDefinitionUpdateOne {
	rduo.mutation.SetParticipantFlowModuleID(id)
	return rduo
}

// SetNillableParticipantFlowModuleID sets the "participant_flow_module" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (rduo *ResultsDefinitionUpdateOne) SetNillableParticipantFlowModuleID(id *int) *ResultsDefinitionUpdateOne {
	if id != nil {
		rduo = rduo.SetParticipantFlowModuleID(*id)
	}
	return rduo
}

// SetParticipantFlowModule sets the "participant_flow_module" edge to the ParticipantFlowModule entity.
func (rduo *ResultsDefinitionUpdateOne) SetParticipantFlowModule(p *ParticipantFlowModule) *ResultsDefinitionUpdateOne {
	return rduo.SetParticipantFlowModuleID(p.ID)
}

// SetBaselineCharacteristicsModuleID sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity by ID.
func (rduo *ResultsDefinitionUpdateOne) SetBaselineCharacteristicsModuleID(id int) *ResultsDefinitionUpdateOne {
	rduo.mutation.SetBaselineCharacteristicsModuleID(id)
	return rduo
}

// SetNillableBaselineCharacteristicsModuleID sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (rduo *ResultsDefinitionUpdateOne) SetNillableBaselineCharacteristicsModuleID(id *int) *ResultsDefinitionUpdateOne {
	if id != nil {
		rduo = rduo.SetBaselineCharacteristicsModuleID(*id)
	}
	return rduo
}

// SetBaselineCharacteristicsModule sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity.
func (rduo *ResultsDefinitionUpdateOne) SetBaselineCharacteristicsModule(b *BaselineCharacteristicsModule) *ResultsDefinitionUpdateOne {
	return rduo.SetBaselineCharacteristicsModuleID(b.ID)
}

// SetOutcomeMeasuresModuleID sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity by ID.
func (rduo *ResultsDefinitionUpdateOne) SetOutcomeMeasuresModuleID(id int) *ResultsDefinitionUpdateOne {
	rduo.mutation.SetOutcomeMeasuresModuleID(id)
	return rduo
}

// SetNillableOutcomeMeasuresModuleID sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity by ID if the given value is not nil.
func (rduo *ResultsDefinitionUpdateOne) SetNillableOutcomeMeasuresModuleID(id *int) *ResultsDefinitionUpdateOne {
	if id != nil {
		rduo = rduo.SetOutcomeMeasuresModuleID(*id)
	}
	return rduo
}

// SetOutcomeMeasuresModule sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity.
func (rduo *ResultsDefinitionUpdateOne) SetOutcomeMeasuresModule(o *OutcomeMeasuresModule) *ResultsDefinitionUpdateOne {
	return rduo.SetOutcomeMeasuresModuleID(o.ID)
}

// SetAdverseEventsModuleID sets the "adverse_events_module" edge to the AdverseEventsModule entity by ID.
func (rduo *ResultsDefinitionUpdateOne) SetAdverseEventsModuleID(id int) *ResultsDefinitionUpdateOne {
	rduo.mutation.SetAdverseEventsModuleID(id)
	return rduo
}

// SetNillableAdverseEventsModuleID sets the "adverse_events_module" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (rduo *ResultsDefinitionUpdateOne) SetNillableAdverseEventsModuleID(id *int) *ResultsDefinitionUpdateOne {
	if id != nil {
		rduo = rduo.SetAdverseEventsModuleID(*id)
	}
	return rduo
}

// SetAdverseEventsModule sets the "adverse_events_module" edge to the AdverseEventsModule entity.
func (rduo *ResultsDefinitionUpdateOne) SetAdverseEventsModule(a *AdverseEventsModule) *ResultsDefinitionUpdateOne {
	return rduo.SetAdverseEventsModuleID(a.ID)
}

// SetMoreInfoModuleID sets the "more_info_module" edge to the MoreInfoModule entity by ID.
func (rduo *ResultsDefinitionUpdateOne) SetMoreInfoModuleID(id int) *ResultsDefinitionUpdateOne {
	rduo.mutation.SetMoreInfoModuleID(id)
	return rduo
}

// SetNillableMoreInfoModuleID sets the "more_info_module" edge to the MoreInfoModule entity by ID if the given value is not nil.
func (rduo *ResultsDefinitionUpdateOne) SetNillableMoreInfoModuleID(id *int) *ResultsDefinitionUpdateOne {
	if id != nil {
		rduo = rduo.SetMoreInfoModuleID(*id)
	}
	return rduo
}

// SetMoreInfoModule sets the "more_info_module" edge to the MoreInfoModule entity.
func (rduo *ResultsDefinitionUpdateOne) SetMoreInfoModule(m *MoreInfoModule) *ResultsDefinitionUpdateOne {
	return rduo.SetMoreInfoModuleID(m.ID)
}

// Mutation returns the ResultsDefinitionMutation object of the builder.
func (rduo *ResultsDefinitionUpdateOne) Mutation() *ResultsDefinitionMutation {
	return rduo.mutation
}

// ClearParent clears the "parent" edge to the ClinicalTrial entity.
func (rduo *ResultsDefinitionUpdateOne) ClearParent() *ResultsDefinitionUpdateOne {
	rduo.mutation.ClearParent()
	return rduo
}

// ClearParticipantFlowModule clears the "participant_flow_module" edge to the ParticipantFlowModule entity.
func (rduo *ResultsDefinitionUpdateOne) ClearParticipantFlowModule() *ResultsDefinitionUpdateOne {
	rduo.mutation.ClearParticipantFlowModule()
	return rduo
}

// ClearBaselineCharacteristicsModule clears the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity.
func (rduo *ResultsDefinitionUpdateOne) ClearBaselineCharacteristicsModule() *ResultsDefinitionUpdateOne {
	rduo.mutation.ClearBaselineCharacteristicsModule()
	return rduo
}

// ClearOutcomeMeasuresModule clears the "outcome_measures_module" edge to the OutcomeMeasuresModule entity.
func (rduo *ResultsDefinitionUpdateOne) ClearOutcomeMeasuresModule() *ResultsDefinitionUpdateOne {
	rduo.mutation.ClearOutcomeMeasuresModule()
	return rduo
}

// ClearAdverseEventsModule clears the "adverse_events_module" edge to the AdverseEventsModule entity.
func (rduo *ResultsDefinitionUpdateOne) ClearAdverseEventsModule() *ResultsDefinitionUpdateOne {
	rduo.mutation.ClearAdverseEventsModule()
	return rduo
}

// ClearMoreInfoModule clears the "more_info_module" edge to the MoreInfoModule entity.
func (rduo *ResultsDefinitionUpdateOne) ClearMoreInfoModule() *ResultsDefinitionUpdateOne {
	rduo.mutation.ClearMoreInfoModule()
	return rduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rduo *ResultsDefinitionUpdateOne) Select(field string, fields ...string) *ResultsDefinitionUpdateOne {
	rduo.fields = append([]string{field}, fields...)
	return rduo
}

// Save executes the query and returns the updated ResultsDefinition entity.
func (rduo *ResultsDefinitionUpdateOne) Save(ctx context.Context) (*ResultsDefinition, error) {
	var (
		err  error
		node *ResultsDefinition
	)
	if len(rduo.hooks) == 0 {
		node, err = rduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultsDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rduo.mutation = mutation
			node, err = rduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rduo.hooks) - 1; i >= 0; i-- {
			if rduo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = rduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rduo *ResultsDefinitionUpdateOne) SaveX(ctx context.Context) *ResultsDefinition {
	node, err := rduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rduo *ResultsDefinitionUpdateOne) Exec(ctx context.Context) error {
	_, err := rduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rduo *ResultsDefinitionUpdateOne) ExecX(ctx context.Context) {
	if err := rduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rduo *ResultsDefinitionUpdateOne) sqlSave(ctx context.Context) (_node *ResultsDefinition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resultsdefinition.Table,
			Columns: resultsdefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultsdefinition.FieldID,
			},
		},
	}
	id, ok := rduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "ResultsDefinition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resultsdefinition.FieldID)
		for _, f := range fields {
			if !resultsdefinition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != resultsdefinition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rduo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resultsdefinition.ParentTable,
			Columns: []string{resultsdefinition.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rduo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resultsdefinition.ParentTable,
			Columns: []string{resultsdefinition.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rduo.mutation.ParticipantFlowModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.ParticipantFlowModuleTable,
			Columns: []string{resultsdefinition.ParticipantFlowModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rduo.mutation.ParticipantFlowModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.ParticipantFlowModuleTable,
			Columns: []string{resultsdefinition.ParticipantFlowModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: participantflowmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rduo.mutation.BaselineCharacteristicsModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.BaselineCharacteristicsModuleTable,
			Columns: []string{resultsdefinition.BaselineCharacteristicsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rduo.mutation.BaselineCharacteristicsModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.BaselineCharacteristicsModuleTable,
			Columns: []string{resultsdefinition.BaselineCharacteristicsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: baselinecharacteristicsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rduo.mutation.OutcomeMeasuresModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.OutcomeMeasuresModuleTable,
			Columns: []string{resultsdefinition.OutcomeMeasuresModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasuresmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rduo.mutation.OutcomeMeasuresModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.OutcomeMeasuresModuleTable,
			Columns: []string{resultsdefinition.OutcomeMeasuresModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasuresmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rduo.mutation.AdverseEventsModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.AdverseEventsModuleTable,
			Columns: []string{resultsdefinition.AdverseEventsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adverseeventsmodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rduo.mutation.AdverseEventsModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.AdverseEventsModuleTable,
			Columns: []string{resultsdefinition.AdverseEventsModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adverseeventsmodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rduo.mutation.MoreInfoModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.MoreInfoModuleTable,
			Columns: []string{resultsdefinition.MoreInfoModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rduo.mutation.MoreInfoModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resultsdefinition.MoreInfoModuleTable,
			Columns: []string{resultsdefinition.MoreInfoModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: moreinfomodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ResultsDefinition{config: rduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resultsdefinition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
