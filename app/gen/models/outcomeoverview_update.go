// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// OutcomeOverviewUpdate is the builder for updating OutcomeOverview entities.
type OutcomeOverviewUpdate struct {
	config
	hooks    []Hook
	mutation *OutcomeOverviewMutation
}

// Where appends a list predicates to the OutcomeOverviewUpdate builder.
func (oou *OutcomeOverviewUpdate) Where(ps ...predicate.OutcomeOverview) *OutcomeOverviewUpdate {
	oou.mutation.Where(ps...)
	return oou
}

// SetOutcomeOverviewID sets the "outcome_overview_id" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewID(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewID(s)
	return oou
}

// SetOutcomeOverviewTitle sets the "outcome_overview_title" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewTitle(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewTitle(s)
	return oou
}

// SetOutcomeOverviewDescription sets the "outcome_overview_description" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewDescription(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewDescription(s)
	return oou
}

// SetOutcomeOverviewParticipants sets the "outcome_overview_participants" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewParticipants(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewParticipants(s)
	return oou
}

// SetOutcomeOverviewTimeFrame sets the "outcome_overview_time_frame" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewTimeFrame(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewTimeFrame(s)
	return oou
}

// SetOutcomeOverviewSerotype sets the "outcome_overview_serotype" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewSerotype(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewSerotype(s)
	return oou
}

// SetOutcomeOverviewAssay sets the "outcome_overview_assay" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewAssay(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewAssay(s)
	return oou
}

// SetOutcomeOverviewDoseNumber sets the "outcome_overview_dose_number" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewDoseNumber(i int64) *OutcomeOverviewUpdate {
	oou.mutation.ResetOutcomeOverviewDoseNumber()
	oou.mutation.SetOutcomeOverviewDoseNumber(i)
	return oou
}

// SetNillableOutcomeOverviewDoseNumber sets the "outcome_overview_dose_number" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewDoseNumber(i *int64) *OutcomeOverviewUpdate {
	if i != nil {
		oou.SetOutcomeOverviewDoseNumber(*i)
	}
	return oou
}

// AddOutcomeOverviewDoseNumber adds i to the "outcome_overview_dose_number" field.
func (oou *OutcomeOverviewUpdate) AddOutcomeOverviewDoseNumber(i int64) *OutcomeOverviewUpdate {
	oou.mutation.AddOutcomeOverviewDoseNumber(i)
	return oou
}

// SetOutcomeOverviewValue sets the "outcome_overview_value" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewValue(f float64) *OutcomeOverviewUpdate {
	oou.mutation.ResetOutcomeOverviewValue()
	oou.mutation.SetOutcomeOverviewValue(f)
	return oou
}

// AddOutcomeOverviewValue adds f to the "outcome_overview_value" field.
func (oou *OutcomeOverviewUpdate) AddOutcomeOverviewValue(f float64) *OutcomeOverviewUpdate {
	oou.mutation.AddOutcomeOverviewValue(f)
	return oou
}

// SetOutcomeOverviewUpperLimit sets the "outcome_overview_upper_limit" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewUpperLimit(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewUpperLimit(s)
	return oou
}

// SetOutcomeOverviewLowerLimit sets the "outcome_overview_lower_limit" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewLowerLimit(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewLowerLimit(s)
	return oou
}

// SetOutcomeOverviewGroupID sets the "outcome_overview_group_id" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewGroupID(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewGroupID(s)
	return oou
}

// SetNillableOutcomeOverviewGroupID sets the "outcome_overview_group_id" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewGroupID(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewGroupID(*s)
	}
	return oou
}

// SetOutcomeOverviewRatio sets the "outcome_overview_ratio" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewRatio(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewRatio(s)
	return oou
}

// SetNillableOutcomeOverviewRatio sets the "outcome_overview_ratio" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewRatio(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewRatio(*s)
	}
	return oou
}

// SetOutcomeOverviewMeasureTitle sets the "outcome_overview_measure_title" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewMeasureTitle(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewMeasureTitle(s)
	return oou
}

// SetNillableOutcomeOverviewMeasureTitle sets the "outcome_overview_measure_title" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewMeasureTitle(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewMeasureTitle(*s)
	}
	return oou
}

// SetOutcomeOverviewVaccine sets the "outcome_overview_vaccine" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewVaccine(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewVaccine(s)
	return oou
}

// SetNillableOutcomeOverviewVaccine sets the "outcome_overview_vaccine" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewVaccine(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewVaccine(*s)
	}
	return oou
}

// SetOutcomeOverviewImmunocompromisedPopulation sets the "outcome_overview_immunocompromised_population" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewImmunocompromisedPopulation(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewImmunocompromisedPopulation(s)
	return oou
}

// SetNillableOutcomeOverviewImmunocompromisedPopulation sets the "outcome_overview_immunocompromised_population" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewImmunocompromisedPopulation(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewImmunocompromisedPopulation(*s)
	}
	return oou
}

// SetOutcomeOverviewManufacturer sets the "outcome_overview_manufacturer" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewManufacturer(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewManufacturer(s)
	return oou
}

// SetNillableOutcomeOverviewManufacturer sets the "outcome_overview_manufacturer" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewManufacturer(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewManufacturer(*s)
	}
	return oou
}

// SetOutcomeOverviewConfidenceInterval sets the "outcome_overview_confidence_interval" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewConfidenceInterval(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewConfidenceInterval(s)
	return oou
}

// SetNillableOutcomeOverviewConfidenceInterval sets the "outcome_overview_confidence_interval" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewConfidenceInterval(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewConfidenceInterval(*s)
	}
	return oou
}

// SetOutcomeOverviewPercentResponders sets the "outcome_overview_percent_responders" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewPercentResponders(f float64) *OutcomeOverviewUpdate {
	oou.mutation.ResetOutcomeOverviewPercentResponders()
	oou.mutation.SetOutcomeOverviewPercentResponders(f)
	return oou
}

// SetNillableOutcomeOverviewPercentResponders sets the "outcome_overview_percent_responders" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewPercentResponders(f *float64) *OutcomeOverviewUpdate {
	if f != nil {
		oou.SetOutcomeOverviewPercentResponders(*f)
	}
	return oou
}

// AddOutcomeOverviewPercentResponders adds f to the "outcome_overview_percent_responders" field.
func (oou *OutcomeOverviewUpdate) AddOutcomeOverviewPercentResponders(f float64) *OutcomeOverviewUpdate {
	oou.mutation.AddOutcomeOverviewPercentResponders(f)
	return oou
}

// SetOutcomeOverviewTimeFrameWeeks sets the "outcome_overview_time_frame_weeks" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewTimeFrameWeeks(i int64) *OutcomeOverviewUpdate {
	oou.mutation.ResetOutcomeOverviewTimeFrameWeeks()
	oou.mutation.SetOutcomeOverviewTimeFrameWeeks(i)
	return oou
}

// SetNillableOutcomeOverviewTimeFrameWeeks sets the "outcome_overview_time_frame_weeks" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewTimeFrameWeeks(i *int64) *OutcomeOverviewUpdate {
	if i != nil {
		oou.SetOutcomeOverviewTimeFrameWeeks(*i)
	}
	return oou
}

// AddOutcomeOverviewTimeFrameWeeks adds i to the "outcome_overview_time_frame_weeks" field.
func (oou *OutcomeOverviewUpdate) AddOutcomeOverviewTimeFrameWeeks(i int64) *OutcomeOverviewUpdate {
	oou.mutation.AddOutcomeOverviewTimeFrameWeeks(i)
	return oou
}

// SetOutcomeOverviewDoseDescription sets the "outcome_overview_dose_description" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewDoseDescription(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewDoseDescription(s)
	return oou
}

// SetNillableOutcomeOverviewDoseDescription sets the "outcome_overview_dose_description" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewDoseDescription(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewDoseDescription(*s)
	}
	return oou
}

// SetOutcomeOverviewSchedule sets the "outcome_overview_schedule" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewSchedule(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewSchedule(s)
	return oou
}

// SetNillableOutcomeOverviewSchedule sets the "outcome_overview_schedule" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewSchedule(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewSchedule(*s)
	}
	return oou
}

// SetOutcomeOverviewUseCaseCode sets the "outcome_overview_use_case_code" field.
func (oou *OutcomeOverviewUpdate) SetOutcomeOverviewUseCaseCode(s string) *OutcomeOverviewUpdate {
	oou.mutation.SetOutcomeOverviewUseCaseCode(s)
	return oou
}

// SetNillableOutcomeOverviewUseCaseCode sets the "outcome_overview_use_case_code" field if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableOutcomeOverviewUseCaseCode(s *string) *OutcomeOverviewUpdate {
	if s != nil {
		oou.SetOutcomeOverviewUseCaseCode(*s)
	}
	return oou
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oou *OutcomeOverviewUpdate) SetParentID(id int) *OutcomeOverviewUpdate {
	oou.mutation.SetParentID(id)
	return oou
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oou *OutcomeOverviewUpdate) SetNillableParentID(id *int) *OutcomeOverviewUpdate {
	if id != nil {
		oou = oou.SetParentID(*id)
	}
	return oou
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oou *OutcomeOverviewUpdate) SetParent(o *OutcomeMeasure) *OutcomeOverviewUpdate {
	return oou.SetParentID(o.ID)
}

// Mutation returns the OutcomeOverviewMutation object of the builder.
func (oou *OutcomeOverviewUpdate) Mutation() *OutcomeOverviewMutation {
	return oou.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (oou *OutcomeOverviewUpdate) ClearParent() *OutcomeOverviewUpdate {
	oou.mutation.ClearParent()
	return oou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oou *OutcomeOverviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oou.hooks) == 0 {
		affected, err = oou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeOverviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oou.mutation = mutation
			affected, err = oou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oou.hooks) - 1; i >= 0; i-- {
			if oou.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oou *OutcomeOverviewUpdate) SaveX(ctx context.Context) int {
	affected, err := oou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oou *OutcomeOverviewUpdate) Exec(ctx context.Context) error {
	_, err := oou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oou *OutcomeOverviewUpdate) ExecX(ctx context.Context) {
	if err := oou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oou *OutcomeOverviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeoverview.Table,
			Columns: outcomeoverview.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeoverview.FieldID,
			},
		},
	}
	if ps := oou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oou.mutation.OutcomeOverviewID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewID,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTitle,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDescription,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewParticipants(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewParticipants,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewTimeFrame(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrame,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewSerotype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewSerotype,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewAssay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewAssay,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewDoseNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseNumber,
		})
	}
	if value, ok := oou.mutation.AddedOutcomeOverviewDoseNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseNumber,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewValue,
		})
	}
	if value, ok := oou.mutation.AddedOutcomeOverviewValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewValue,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewUpperLimit,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewLowerLimit,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewGroupID,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewRatio,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewMeasureTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewMeasureTitle,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewVaccine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewVaccine,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewImmunocompromisedPopulation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewImmunocompromisedPopulation,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewManufacturer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewManufacturer,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewConfidenceInterval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewConfidenceInterval,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewPercentResponders(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewPercentResponders,
		})
	}
	if value, ok := oou.mutation.AddedOutcomeOverviewPercentResponders(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewPercentResponders,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewTimeFrameWeeks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks,
		})
	}
	if value, ok := oou.mutation.AddedOutcomeOverviewTimeFrameWeeks(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewDoseDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseDescription,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewSchedule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewSchedule,
		})
	}
	if value, ok := oou.mutation.OutcomeOverviewUseCaseCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewUseCaseCode,
		})
	}
	if oou.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeoverview.ParentTable,
			Columns: []string{outcomeoverview.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oou.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeoverview.ParentTable,
			Columns: []string{outcomeoverview.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeoverview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OutcomeOverviewUpdateOne is the builder for updating a single OutcomeOverview entity.
type OutcomeOverviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutcomeOverviewMutation
}

// SetOutcomeOverviewID sets the "outcome_overview_id" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewID(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewID(s)
	return oouo
}

// SetOutcomeOverviewTitle sets the "outcome_overview_title" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewTitle(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewTitle(s)
	return oouo
}

// SetOutcomeOverviewDescription sets the "outcome_overview_description" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewDescription(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewDescription(s)
	return oouo
}

// SetOutcomeOverviewParticipants sets the "outcome_overview_participants" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewParticipants(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewParticipants(s)
	return oouo
}

// SetOutcomeOverviewTimeFrame sets the "outcome_overview_time_frame" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewTimeFrame(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewTimeFrame(s)
	return oouo
}

// SetOutcomeOverviewSerotype sets the "outcome_overview_serotype" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewSerotype(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewSerotype(s)
	return oouo
}

// SetOutcomeOverviewAssay sets the "outcome_overview_assay" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewAssay(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewAssay(s)
	return oouo
}

// SetOutcomeOverviewDoseNumber sets the "outcome_overview_dose_number" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewDoseNumber(i int64) *OutcomeOverviewUpdateOne {
	oouo.mutation.ResetOutcomeOverviewDoseNumber()
	oouo.mutation.SetOutcomeOverviewDoseNumber(i)
	return oouo
}

// SetNillableOutcomeOverviewDoseNumber sets the "outcome_overview_dose_number" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewDoseNumber(i *int64) *OutcomeOverviewUpdateOne {
	if i != nil {
		oouo.SetOutcomeOverviewDoseNumber(*i)
	}
	return oouo
}

// AddOutcomeOverviewDoseNumber adds i to the "outcome_overview_dose_number" field.
func (oouo *OutcomeOverviewUpdateOne) AddOutcomeOverviewDoseNumber(i int64) *OutcomeOverviewUpdateOne {
	oouo.mutation.AddOutcomeOverviewDoseNumber(i)
	return oouo
}

// SetOutcomeOverviewValue sets the "outcome_overview_value" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewValue(f float64) *OutcomeOverviewUpdateOne {
	oouo.mutation.ResetOutcomeOverviewValue()
	oouo.mutation.SetOutcomeOverviewValue(f)
	return oouo
}

// AddOutcomeOverviewValue adds f to the "outcome_overview_value" field.
func (oouo *OutcomeOverviewUpdateOne) AddOutcomeOverviewValue(f float64) *OutcomeOverviewUpdateOne {
	oouo.mutation.AddOutcomeOverviewValue(f)
	return oouo
}

// SetOutcomeOverviewUpperLimit sets the "outcome_overview_upper_limit" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewUpperLimit(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewUpperLimit(s)
	return oouo
}

// SetOutcomeOverviewLowerLimit sets the "outcome_overview_lower_limit" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewLowerLimit(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewLowerLimit(s)
	return oouo
}

// SetOutcomeOverviewGroupID sets the "outcome_overview_group_id" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewGroupID(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewGroupID(s)
	return oouo
}

// SetNillableOutcomeOverviewGroupID sets the "outcome_overview_group_id" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewGroupID(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewGroupID(*s)
	}
	return oouo
}

// SetOutcomeOverviewRatio sets the "outcome_overview_ratio" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewRatio(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewRatio(s)
	return oouo
}

// SetNillableOutcomeOverviewRatio sets the "outcome_overview_ratio" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewRatio(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewRatio(*s)
	}
	return oouo
}

// SetOutcomeOverviewMeasureTitle sets the "outcome_overview_measure_title" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewMeasureTitle(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewMeasureTitle(s)
	return oouo
}

// SetNillableOutcomeOverviewMeasureTitle sets the "outcome_overview_measure_title" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewMeasureTitle(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewMeasureTitle(*s)
	}
	return oouo
}

// SetOutcomeOverviewVaccine sets the "outcome_overview_vaccine" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewVaccine(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewVaccine(s)
	return oouo
}

// SetNillableOutcomeOverviewVaccine sets the "outcome_overview_vaccine" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewVaccine(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewVaccine(*s)
	}
	return oouo
}

// SetOutcomeOverviewImmunocompromisedPopulation sets the "outcome_overview_immunocompromised_population" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewImmunocompromisedPopulation(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewImmunocompromisedPopulation(s)
	return oouo
}

// SetNillableOutcomeOverviewImmunocompromisedPopulation sets the "outcome_overview_immunocompromised_population" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewImmunocompromisedPopulation(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewImmunocompromisedPopulation(*s)
	}
	return oouo
}

// SetOutcomeOverviewManufacturer sets the "outcome_overview_manufacturer" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewManufacturer(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewManufacturer(s)
	return oouo
}

// SetNillableOutcomeOverviewManufacturer sets the "outcome_overview_manufacturer" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewManufacturer(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewManufacturer(*s)
	}
	return oouo
}

// SetOutcomeOverviewConfidenceInterval sets the "outcome_overview_confidence_interval" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewConfidenceInterval(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewConfidenceInterval(s)
	return oouo
}

// SetNillableOutcomeOverviewConfidenceInterval sets the "outcome_overview_confidence_interval" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewConfidenceInterval(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewConfidenceInterval(*s)
	}
	return oouo
}

// SetOutcomeOverviewPercentResponders sets the "outcome_overview_percent_responders" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewPercentResponders(f float64) *OutcomeOverviewUpdateOne {
	oouo.mutation.ResetOutcomeOverviewPercentResponders()
	oouo.mutation.SetOutcomeOverviewPercentResponders(f)
	return oouo
}

// SetNillableOutcomeOverviewPercentResponders sets the "outcome_overview_percent_responders" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewPercentResponders(f *float64) *OutcomeOverviewUpdateOne {
	if f != nil {
		oouo.SetOutcomeOverviewPercentResponders(*f)
	}
	return oouo
}

// AddOutcomeOverviewPercentResponders adds f to the "outcome_overview_percent_responders" field.
func (oouo *OutcomeOverviewUpdateOne) AddOutcomeOverviewPercentResponders(f float64) *OutcomeOverviewUpdateOne {
	oouo.mutation.AddOutcomeOverviewPercentResponders(f)
	return oouo
}

// SetOutcomeOverviewTimeFrameWeeks sets the "outcome_overview_time_frame_weeks" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewTimeFrameWeeks(i int64) *OutcomeOverviewUpdateOne {
	oouo.mutation.ResetOutcomeOverviewTimeFrameWeeks()
	oouo.mutation.SetOutcomeOverviewTimeFrameWeeks(i)
	return oouo
}

// SetNillableOutcomeOverviewTimeFrameWeeks sets the "outcome_overview_time_frame_weeks" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewTimeFrameWeeks(i *int64) *OutcomeOverviewUpdateOne {
	if i != nil {
		oouo.SetOutcomeOverviewTimeFrameWeeks(*i)
	}
	return oouo
}

// AddOutcomeOverviewTimeFrameWeeks adds i to the "outcome_overview_time_frame_weeks" field.
func (oouo *OutcomeOverviewUpdateOne) AddOutcomeOverviewTimeFrameWeeks(i int64) *OutcomeOverviewUpdateOne {
	oouo.mutation.AddOutcomeOverviewTimeFrameWeeks(i)
	return oouo
}

// SetOutcomeOverviewDoseDescription sets the "outcome_overview_dose_description" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewDoseDescription(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewDoseDescription(s)
	return oouo
}

// SetNillableOutcomeOverviewDoseDescription sets the "outcome_overview_dose_description" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewDoseDescription(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewDoseDescription(*s)
	}
	return oouo
}

// SetOutcomeOverviewSchedule sets the "outcome_overview_schedule" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewSchedule(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewSchedule(s)
	return oouo
}

// SetNillableOutcomeOverviewSchedule sets the "outcome_overview_schedule" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewSchedule(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewSchedule(*s)
	}
	return oouo
}

// SetOutcomeOverviewUseCaseCode sets the "outcome_overview_use_case_code" field.
func (oouo *OutcomeOverviewUpdateOne) SetOutcomeOverviewUseCaseCode(s string) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetOutcomeOverviewUseCaseCode(s)
	return oouo
}

// SetNillableOutcomeOverviewUseCaseCode sets the "outcome_overview_use_case_code" field if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableOutcomeOverviewUseCaseCode(s *string) *OutcomeOverviewUpdateOne {
	if s != nil {
		oouo.SetOutcomeOverviewUseCaseCode(*s)
	}
	return oouo
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (oouo *OutcomeOverviewUpdateOne) SetParentID(id int) *OutcomeOverviewUpdateOne {
	oouo.mutation.SetParentID(id)
	return oouo
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (oouo *OutcomeOverviewUpdateOne) SetNillableParentID(id *int) *OutcomeOverviewUpdateOne {
	if id != nil {
		oouo = oouo.SetParentID(*id)
	}
	return oouo
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (oouo *OutcomeOverviewUpdateOne) SetParent(o *OutcomeMeasure) *OutcomeOverviewUpdateOne {
	return oouo.SetParentID(o.ID)
}

// Mutation returns the OutcomeOverviewMutation object of the builder.
func (oouo *OutcomeOverviewUpdateOne) Mutation() *OutcomeOverviewMutation {
	return oouo.mutation
}

// ClearParent clears the "parent" edge to the OutcomeMeasure entity.
func (oouo *OutcomeOverviewUpdateOne) ClearParent() *OutcomeOverviewUpdateOne {
	oouo.mutation.ClearParent()
	return oouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oouo *OutcomeOverviewUpdateOne) Select(field string, fields ...string) *OutcomeOverviewUpdateOne {
	oouo.fields = append([]string{field}, fields...)
	return oouo
}

// Save executes the query and returns the updated OutcomeOverview entity.
func (oouo *OutcomeOverviewUpdateOne) Save(ctx context.Context) (*OutcomeOverview, error) {
	var (
		err  error
		node *OutcomeOverview
	)
	if len(oouo.hooks) == 0 {
		node, err = oouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeOverviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oouo.mutation = mutation
			node, err = oouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oouo.hooks) - 1; i >= 0; i-- {
			if oouo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = oouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oouo *OutcomeOverviewUpdateOne) SaveX(ctx context.Context) *OutcomeOverview {
	node, err := oouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oouo *OutcomeOverviewUpdateOne) Exec(ctx context.Context) error {
	_, err := oouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oouo *OutcomeOverviewUpdateOne) ExecX(ctx context.Context) {
	if err := oouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oouo *OutcomeOverviewUpdateOne) sqlSave(ctx context.Context) (_node *OutcomeOverview, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outcomeoverview.Table,
			Columns: outcomeoverview.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeoverview.FieldID,
			},
		},
	}
	id, ok := oouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "OutcomeOverview.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outcomeoverview.FieldID)
		for _, f := range fields {
			if !outcomeoverview.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != outcomeoverview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oouo.mutation.OutcomeOverviewID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewID,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTitle,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDescription,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewParticipants(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewParticipants,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewTimeFrame(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrame,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewSerotype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewSerotype,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewAssay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewAssay,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewDoseNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseNumber,
		})
	}
	if value, ok := oouo.mutation.AddedOutcomeOverviewDoseNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseNumber,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewValue,
		})
	}
	if value, ok := oouo.mutation.AddedOutcomeOverviewValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewValue,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewUpperLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewUpperLimit,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewLowerLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewLowerLimit,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewGroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewGroupID,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewRatio,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewMeasureTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewMeasureTitle,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewVaccine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewVaccine,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewImmunocompromisedPopulation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewImmunocompromisedPopulation,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewManufacturer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewManufacturer,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewConfidenceInterval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewConfidenceInterval,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewPercentResponders(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewPercentResponders,
		})
	}
	if value, ok := oouo.mutation.AddedOutcomeOverviewPercentResponders(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewPercentResponders,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewTimeFrameWeeks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks,
		})
	}
	if value, ok := oouo.mutation.AddedOutcomeOverviewTimeFrameWeeks(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewDoseDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseDescription,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewSchedule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewSchedule,
		})
	}
	if value, ok := oouo.mutation.OutcomeOverviewUseCaseCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewUseCaseCode,
		})
	}
	if oouo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeoverview.ParentTable,
			Columns: []string{outcomeoverview.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oouo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outcomeoverview.ParentTable,
			Columns: []string{outcomeoverview.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outcomemeasure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutcomeOverview{config: oouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outcomeoverview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
