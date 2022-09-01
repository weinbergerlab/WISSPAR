// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/adverseeventsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/baselinecharacteristicsmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasuresmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/participantflowmodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// ResultsDefinitionCreate is the builder for creating a ResultsDefinition entity.
type ResultsDefinitionCreate struct {
	config
	mutation *ResultsDefinitionMutation
	hooks    []Hook
}

// SetParentID sets the "parent" edge to the ClinicalTrial entity by ID.
func (rdc *ResultsDefinitionCreate) SetParentID(id int) *ResultsDefinitionCreate {
	rdc.mutation.SetParentID(id)
	return rdc
}

// SetNillableParentID sets the "parent" edge to the ClinicalTrial entity by ID if the given value is not nil.
func (rdc *ResultsDefinitionCreate) SetNillableParentID(id *int) *ResultsDefinitionCreate {
	if id != nil {
		rdc = rdc.SetParentID(*id)
	}
	return rdc
}

// SetParent sets the "parent" edge to the ClinicalTrial entity.
func (rdc *ResultsDefinitionCreate) SetParent(c *ClinicalTrial) *ResultsDefinitionCreate {
	return rdc.SetParentID(c.ID)
}

// SetParticipantFlowModuleID sets the "participant_flow_module" edge to the ParticipantFlowModule entity by ID.
func (rdc *ResultsDefinitionCreate) SetParticipantFlowModuleID(id int) *ResultsDefinitionCreate {
	rdc.mutation.SetParticipantFlowModuleID(id)
	return rdc
}

// SetNillableParticipantFlowModuleID sets the "participant_flow_module" edge to the ParticipantFlowModule entity by ID if the given value is not nil.
func (rdc *ResultsDefinitionCreate) SetNillableParticipantFlowModuleID(id *int) *ResultsDefinitionCreate {
	if id != nil {
		rdc = rdc.SetParticipantFlowModuleID(*id)
	}
	return rdc
}

// SetParticipantFlowModule sets the "participant_flow_module" edge to the ParticipantFlowModule entity.
func (rdc *ResultsDefinitionCreate) SetParticipantFlowModule(p *ParticipantFlowModule) *ResultsDefinitionCreate {
	return rdc.SetParticipantFlowModuleID(p.ID)
}

// SetBaselineCharacteristicsModuleID sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity by ID.
func (rdc *ResultsDefinitionCreate) SetBaselineCharacteristicsModuleID(id int) *ResultsDefinitionCreate {
	rdc.mutation.SetBaselineCharacteristicsModuleID(id)
	return rdc
}

// SetNillableBaselineCharacteristicsModuleID sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity by ID if the given value is not nil.
func (rdc *ResultsDefinitionCreate) SetNillableBaselineCharacteristicsModuleID(id *int) *ResultsDefinitionCreate {
	if id != nil {
		rdc = rdc.SetBaselineCharacteristicsModuleID(*id)
	}
	return rdc
}

// SetBaselineCharacteristicsModule sets the "baseline_characteristics_module" edge to the BaselineCharacteristicsModule entity.
func (rdc *ResultsDefinitionCreate) SetBaselineCharacteristicsModule(b *BaselineCharacteristicsModule) *ResultsDefinitionCreate {
	return rdc.SetBaselineCharacteristicsModuleID(b.ID)
}

// SetOutcomeMeasuresModuleID sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity by ID.
func (rdc *ResultsDefinitionCreate) SetOutcomeMeasuresModuleID(id int) *ResultsDefinitionCreate {
	rdc.mutation.SetOutcomeMeasuresModuleID(id)
	return rdc
}

// SetNillableOutcomeMeasuresModuleID sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity by ID if the given value is not nil.
func (rdc *ResultsDefinitionCreate) SetNillableOutcomeMeasuresModuleID(id *int) *ResultsDefinitionCreate {
	if id != nil {
		rdc = rdc.SetOutcomeMeasuresModuleID(*id)
	}
	return rdc
}

// SetOutcomeMeasuresModule sets the "outcome_measures_module" edge to the OutcomeMeasuresModule entity.
func (rdc *ResultsDefinitionCreate) SetOutcomeMeasuresModule(o *OutcomeMeasuresModule) *ResultsDefinitionCreate {
	return rdc.SetOutcomeMeasuresModuleID(o.ID)
}

// SetAdverseEventsModuleID sets the "adverse_events_module" edge to the AdverseEventsModule entity by ID.
func (rdc *ResultsDefinitionCreate) SetAdverseEventsModuleID(id int) *ResultsDefinitionCreate {
	rdc.mutation.SetAdverseEventsModuleID(id)
	return rdc
}

// SetNillableAdverseEventsModuleID sets the "adverse_events_module" edge to the AdverseEventsModule entity by ID if the given value is not nil.
func (rdc *ResultsDefinitionCreate) SetNillableAdverseEventsModuleID(id *int) *ResultsDefinitionCreate {
	if id != nil {
		rdc = rdc.SetAdverseEventsModuleID(*id)
	}
	return rdc
}

// SetAdverseEventsModule sets the "adverse_events_module" edge to the AdverseEventsModule entity.
func (rdc *ResultsDefinitionCreate) SetAdverseEventsModule(a *AdverseEventsModule) *ResultsDefinitionCreate {
	return rdc.SetAdverseEventsModuleID(a.ID)
}

// SetMoreInfoModuleID sets the "more_info_module" edge to the MoreInfoModule entity by ID.
func (rdc *ResultsDefinitionCreate) SetMoreInfoModuleID(id int) *ResultsDefinitionCreate {
	rdc.mutation.SetMoreInfoModuleID(id)
	return rdc
}

// SetNillableMoreInfoModuleID sets the "more_info_module" edge to the MoreInfoModule entity by ID if the given value is not nil.
func (rdc *ResultsDefinitionCreate) SetNillableMoreInfoModuleID(id *int) *ResultsDefinitionCreate {
	if id != nil {
		rdc = rdc.SetMoreInfoModuleID(*id)
	}
	return rdc
}

// SetMoreInfoModule sets the "more_info_module" edge to the MoreInfoModule entity.
func (rdc *ResultsDefinitionCreate) SetMoreInfoModule(m *MoreInfoModule) *ResultsDefinitionCreate {
	return rdc.SetMoreInfoModuleID(m.ID)
}

// Mutation returns the ResultsDefinitionMutation object of the builder.
func (rdc *ResultsDefinitionCreate) Mutation() *ResultsDefinitionMutation {
	return rdc.mutation
}

// Save creates the ResultsDefinition in the database.
func (rdc *ResultsDefinitionCreate) Save(ctx context.Context) (*ResultsDefinition, error) {
	var (
		err  error
		node *ResultsDefinition
	)
	if len(rdc.hooks) == 0 {
		if err = rdc.check(); err != nil {
			return nil, err
		}
		node, err = rdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultsDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rdc.check(); err != nil {
				return nil, err
			}
			rdc.mutation = mutation
			if node, err = rdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rdc.hooks) - 1; i >= 0; i-- {
			if rdc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = rdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rdc *ResultsDefinitionCreate) SaveX(ctx context.Context) *ResultsDefinition {
	v, err := rdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdc *ResultsDefinitionCreate) Exec(ctx context.Context) error {
	_, err := rdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdc *ResultsDefinitionCreate) ExecX(ctx context.Context) {
	if err := rdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdc *ResultsDefinitionCreate) check() error {
	return nil
}

func (rdc *ResultsDefinitionCreate) sqlSave(ctx context.Context) (*ResultsDefinition, error) {
	_node, _spec := rdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rdc *ResultsDefinitionCreate) createSpec() (*ResultsDefinition, *sqlgraph.CreateSpec) {
	var (
		_node = &ResultsDefinition{config: rdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resultsdefinition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultsdefinition.FieldID,
			},
		}
	)
	if nodes := rdc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.clinical_trial_results_definition = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rdc.mutation.ParticipantFlowModuleIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rdc.mutation.BaselineCharacteristicsModuleIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rdc.mutation.OutcomeMeasuresModuleIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rdc.mutation.AdverseEventsModuleIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rdc.mutation.MoreInfoModuleIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResultsDefinitionCreateBulk is the builder for creating many ResultsDefinition entities in bulk.
type ResultsDefinitionCreateBulk struct {
	config
	builders []*ResultsDefinitionCreate
}

// Save creates the ResultsDefinition entities in the database.
func (rdcb *ResultsDefinitionCreateBulk) Save(ctx context.Context) ([]*ResultsDefinition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdcb.builders))
	nodes := make([]*ResultsDefinition, len(rdcb.builders))
	mutators := make([]Mutator, len(rdcb.builders))
	for i := range rdcb.builders {
		func(i int, root context.Context) {
			builder := rdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResultsDefinitionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdcb *ResultsDefinitionCreateBulk) SaveX(ctx context.Context) []*ResultsDefinition {
	v, err := rdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdcb *ResultsDefinitionCreateBulk) Exec(ctx context.Context) error {
	_, err := rdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdcb *ResultsDefinitionCreateBulk) ExecX(ctx context.Context) {
	if err := rdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
