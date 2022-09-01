// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomemeasure"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/outcomeoverview"
)

// OutcomeOverviewCreate is the builder for creating a OutcomeOverview entity.
type OutcomeOverviewCreate struct {
	config
	mutation *OutcomeOverviewMutation
	hooks    []Hook
}

// SetOutcomeOverviewID sets the "outcome_overview_id" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewID(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewID(s)
	return ooc
}

// SetOutcomeOverviewTitle sets the "outcome_overview_title" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewTitle(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewTitle(s)
	return ooc
}

// SetOutcomeOverviewDescription sets the "outcome_overview_description" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewDescription(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewDescription(s)
	return ooc
}

// SetOutcomeOverviewParticipants sets the "outcome_overview_participants" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewParticipants(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewParticipants(s)
	return ooc
}

// SetOutcomeOverviewTimeFrame sets the "outcome_overview_time_frame" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewTimeFrame(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewTimeFrame(s)
	return ooc
}

// SetOutcomeOverviewSerotype sets the "outcome_overview_serotype" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewSerotype(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewSerotype(s)
	return ooc
}

// SetOutcomeOverviewAssay sets the "outcome_overview_assay" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewAssay(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewAssay(s)
	return ooc
}

// SetOutcomeOverviewDoseNumber sets the "outcome_overview_dose_number" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewDoseNumber(i int64) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewDoseNumber(i)
	return ooc
}

// SetNillableOutcomeOverviewDoseNumber sets the "outcome_overview_dose_number" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewDoseNumber(i *int64) *OutcomeOverviewCreate {
	if i != nil {
		ooc.SetOutcomeOverviewDoseNumber(*i)
	}
	return ooc
}

// SetOutcomeOverviewValue sets the "outcome_overview_value" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewValue(f float64) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewValue(f)
	return ooc
}

// SetOutcomeOverviewUpperLimit sets the "outcome_overview_upper_limit" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewUpperLimit(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewUpperLimit(s)
	return ooc
}

// SetOutcomeOverviewLowerLimit sets the "outcome_overview_lower_limit" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewLowerLimit(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewLowerLimit(s)
	return ooc
}

// SetOutcomeOverviewGroupID sets the "outcome_overview_group_id" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewGroupID(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewGroupID(s)
	return ooc
}

// SetNillableOutcomeOverviewGroupID sets the "outcome_overview_group_id" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewGroupID(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewGroupID(*s)
	}
	return ooc
}

// SetOutcomeOverviewRatio sets the "outcome_overview_ratio" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewRatio(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewRatio(s)
	return ooc
}

// SetNillableOutcomeOverviewRatio sets the "outcome_overview_ratio" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewRatio(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewRatio(*s)
	}
	return ooc
}

// SetOutcomeOverviewMeasureTitle sets the "outcome_overview_measure_title" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewMeasureTitle(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewMeasureTitle(s)
	return ooc
}

// SetNillableOutcomeOverviewMeasureTitle sets the "outcome_overview_measure_title" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewMeasureTitle(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewMeasureTitle(*s)
	}
	return ooc
}

// SetOutcomeOverviewVaccine sets the "outcome_overview_vaccine" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewVaccine(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewVaccine(s)
	return ooc
}

// SetNillableOutcomeOverviewVaccine sets the "outcome_overview_vaccine" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewVaccine(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewVaccine(*s)
	}
	return ooc
}

// SetOutcomeOverviewImmunocompromisedPopulation sets the "outcome_overview_immunocompromised_population" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewImmunocompromisedPopulation(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewImmunocompromisedPopulation(s)
	return ooc
}

// SetNillableOutcomeOverviewImmunocompromisedPopulation sets the "outcome_overview_immunocompromised_population" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewImmunocompromisedPopulation(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewImmunocompromisedPopulation(*s)
	}
	return ooc
}

// SetOutcomeOverviewManufacturer sets the "outcome_overview_manufacturer" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewManufacturer(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewManufacturer(s)
	return ooc
}

// SetNillableOutcomeOverviewManufacturer sets the "outcome_overview_manufacturer" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewManufacturer(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewManufacturer(*s)
	}
	return ooc
}

// SetOutcomeOverviewConfidenceInterval sets the "outcome_overview_confidence_interval" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewConfidenceInterval(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewConfidenceInterval(s)
	return ooc
}

// SetNillableOutcomeOverviewConfidenceInterval sets the "outcome_overview_confidence_interval" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewConfidenceInterval(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewConfidenceInterval(*s)
	}
	return ooc
}

// SetOutcomeOverviewPercentResponders sets the "outcome_overview_percent_responders" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewPercentResponders(f float64) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewPercentResponders(f)
	return ooc
}

// SetNillableOutcomeOverviewPercentResponders sets the "outcome_overview_percent_responders" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewPercentResponders(f *float64) *OutcomeOverviewCreate {
	if f != nil {
		ooc.SetOutcomeOverviewPercentResponders(*f)
	}
	return ooc
}

// SetOutcomeOverviewTimeFrameWeeks sets the "outcome_overview_time_frame_weeks" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewTimeFrameWeeks(i int64) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewTimeFrameWeeks(i)
	return ooc
}

// SetNillableOutcomeOverviewTimeFrameWeeks sets the "outcome_overview_time_frame_weeks" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewTimeFrameWeeks(i *int64) *OutcomeOverviewCreate {
	if i != nil {
		ooc.SetOutcomeOverviewTimeFrameWeeks(*i)
	}
	return ooc
}

// SetOutcomeOverviewDoseDescription sets the "outcome_overview_dose_description" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewDoseDescription(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewDoseDescription(s)
	return ooc
}

// SetNillableOutcomeOverviewDoseDescription sets the "outcome_overview_dose_description" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewDoseDescription(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewDoseDescription(*s)
	}
	return ooc
}

// SetOutcomeOverviewSchedule sets the "outcome_overview_schedule" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewSchedule(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewSchedule(s)
	return ooc
}

// SetNillableOutcomeOverviewSchedule sets the "outcome_overview_schedule" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewSchedule(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewSchedule(*s)
	}
	return ooc
}

// SetOutcomeOverviewUseCaseCode sets the "outcome_overview_use_case_code" field.
func (ooc *OutcomeOverviewCreate) SetOutcomeOverviewUseCaseCode(s string) *OutcomeOverviewCreate {
	ooc.mutation.SetOutcomeOverviewUseCaseCode(s)
	return ooc
}

// SetNillableOutcomeOverviewUseCaseCode sets the "outcome_overview_use_case_code" field if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableOutcomeOverviewUseCaseCode(s *string) *OutcomeOverviewCreate {
	if s != nil {
		ooc.SetOutcomeOverviewUseCaseCode(*s)
	}
	return ooc
}

// SetParentID sets the "parent" edge to the OutcomeMeasure entity by ID.
func (ooc *OutcomeOverviewCreate) SetParentID(id int) *OutcomeOverviewCreate {
	ooc.mutation.SetParentID(id)
	return ooc
}

// SetNillableParentID sets the "parent" edge to the OutcomeMeasure entity by ID if the given value is not nil.
func (ooc *OutcomeOverviewCreate) SetNillableParentID(id *int) *OutcomeOverviewCreate {
	if id != nil {
		ooc = ooc.SetParentID(*id)
	}
	return ooc
}

// SetParent sets the "parent" edge to the OutcomeMeasure entity.
func (ooc *OutcomeOverviewCreate) SetParent(o *OutcomeMeasure) *OutcomeOverviewCreate {
	return ooc.SetParentID(o.ID)
}

// Mutation returns the OutcomeOverviewMutation object of the builder.
func (ooc *OutcomeOverviewCreate) Mutation() *OutcomeOverviewMutation {
	return ooc.mutation
}

// Save creates the OutcomeOverview in the database.
func (ooc *OutcomeOverviewCreate) Save(ctx context.Context) (*OutcomeOverview, error) {
	var (
		err  error
		node *OutcomeOverview
	)
	ooc.defaults()
	if len(ooc.hooks) == 0 {
		if err = ooc.check(); err != nil {
			return nil, err
		}
		node, err = ooc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutcomeOverviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ooc.check(); err != nil {
				return nil, err
			}
			ooc.mutation = mutation
			if node, err = ooc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ooc.hooks) - 1; i >= 0; i-- {
			if ooc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ooc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ooc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ooc *OutcomeOverviewCreate) SaveX(ctx context.Context) *OutcomeOverview {
	v, err := ooc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ooc *OutcomeOverviewCreate) Exec(ctx context.Context) error {
	_, err := ooc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ooc *OutcomeOverviewCreate) ExecX(ctx context.Context) {
	if err := ooc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ooc *OutcomeOverviewCreate) defaults() {
	if _, ok := ooc.mutation.OutcomeOverviewDoseNumber(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewDoseNumber
		ooc.mutation.SetOutcomeOverviewDoseNumber(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewGroupID(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewGroupID
		ooc.mutation.SetOutcomeOverviewGroupID(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewRatio(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewRatio
		ooc.mutation.SetOutcomeOverviewRatio(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewMeasureTitle(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewMeasureTitle
		ooc.mutation.SetOutcomeOverviewMeasureTitle(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewVaccine(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewVaccine
		ooc.mutation.SetOutcomeOverviewVaccine(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewImmunocompromisedPopulation(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewImmunocompromisedPopulation
		ooc.mutation.SetOutcomeOverviewImmunocompromisedPopulation(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewManufacturer(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewManufacturer
		ooc.mutation.SetOutcomeOverviewManufacturer(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewConfidenceInterval(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewConfidenceInterval
		ooc.mutation.SetOutcomeOverviewConfidenceInterval(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewPercentResponders(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewPercentResponders
		ooc.mutation.SetOutcomeOverviewPercentResponders(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewTimeFrameWeeks(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewTimeFrameWeeks
		ooc.mutation.SetOutcomeOverviewTimeFrameWeeks(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewDoseDescription(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewDoseDescription
		ooc.mutation.SetOutcomeOverviewDoseDescription(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewSchedule(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewSchedule
		ooc.mutation.SetOutcomeOverviewSchedule(v)
	}
	if _, ok := ooc.mutation.OutcomeOverviewUseCaseCode(); !ok {
		v := outcomeoverview.DefaultOutcomeOverviewUseCaseCode
		ooc.mutation.SetOutcomeOverviewUseCaseCode(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ooc *OutcomeOverviewCreate) check() error {
	if _, ok := ooc.mutation.OutcomeOverviewID(); !ok {
		return &ValidationError{Name: "outcome_overview_id", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_id"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewTitle(); !ok {
		return &ValidationError{Name: "outcome_overview_title", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_title"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewDescription(); !ok {
		return &ValidationError{Name: "outcome_overview_description", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_description"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewParticipants(); !ok {
		return &ValidationError{Name: "outcome_overview_participants", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_participants"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewTimeFrame(); !ok {
		return &ValidationError{Name: "outcome_overview_time_frame", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_time_frame"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewSerotype(); !ok {
		return &ValidationError{Name: "outcome_overview_serotype", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_serotype"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewAssay(); !ok {
		return &ValidationError{Name: "outcome_overview_assay", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_assay"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewDoseNumber(); !ok {
		return &ValidationError{Name: "outcome_overview_dose_number", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_dose_number"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewValue(); !ok {
		return &ValidationError{Name: "outcome_overview_value", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_value"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewUpperLimit(); !ok {
		return &ValidationError{Name: "outcome_overview_upper_limit", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_upper_limit"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewLowerLimit(); !ok {
		return &ValidationError{Name: "outcome_overview_lower_limit", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_lower_limit"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewGroupID(); !ok {
		return &ValidationError{Name: "outcome_overview_group_id", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_group_id"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewRatio(); !ok {
		return &ValidationError{Name: "outcome_overview_ratio", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_ratio"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewMeasureTitle(); !ok {
		return &ValidationError{Name: "outcome_overview_measure_title", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_measure_title"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewVaccine(); !ok {
		return &ValidationError{Name: "outcome_overview_vaccine", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_vaccine"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewImmunocompromisedPopulation(); !ok {
		return &ValidationError{Name: "outcome_overview_immunocompromised_population", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_immunocompromised_population"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewManufacturer(); !ok {
		return &ValidationError{Name: "outcome_overview_manufacturer", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_manufacturer"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewConfidenceInterval(); !ok {
		return &ValidationError{Name: "outcome_overview_confidence_interval", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_confidence_interval"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewPercentResponders(); !ok {
		return &ValidationError{Name: "outcome_overview_percent_responders", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_percent_responders"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewTimeFrameWeeks(); !ok {
		return &ValidationError{Name: "outcome_overview_time_frame_weeks", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_time_frame_weeks"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewDoseDescription(); !ok {
		return &ValidationError{Name: "outcome_overview_dose_description", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_dose_description"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewSchedule(); !ok {
		return &ValidationError{Name: "outcome_overview_schedule", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_schedule"`)}
	}
	if _, ok := ooc.mutation.OutcomeOverviewUseCaseCode(); !ok {
		return &ValidationError{Name: "outcome_overview_use_case_code", err: errors.New(`models: missing required field "OutcomeOverview.outcome_overview_use_case_code"`)}
	}
	return nil
}

func (ooc *OutcomeOverviewCreate) sqlSave(ctx context.Context) (*OutcomeOverview, error) {
	_node, _spec := ooc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ooc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ooc *OutcomeOverviewCreate) createSpec() (*OutcomeOverview, *sqlgraph.CreateSpec) {
	var (
		_node = &OutcomeOverview{config: ooc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outcomeoverview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outcomeoverview.FieldID,
			},
		}
	)
	if value, ok := ooc.mutation.OutcomeOverviewID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewID,
		})
		_node.OutcomeOverviewID = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTitle,
		})
		_node.OutcomeOverviewTitle = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDescription,
		})
		_node.OutcomeOverviewDescription = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewParticipants(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewParticipants,
		})
		_node.OutcomeOverviewParticipants = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewTimeFrame(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrame,
		})
		_node.OutcomeOverviewTimeFrame = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewSerotype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewSerotype,
		})
		_node.OutcomeOverviewSerotype = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewAssay(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewAssay,
		})
		_node.OutcomeOverviewAssay = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewDoseNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseNumber,
		})
		_node.OutcomeOverviewDoseNumber = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewValue,
		})
		_node.OutcomeOverviewValue = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewUpperLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewUpperLimit,
		})
		_node.OutcomeOverviewUpperLimit = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewLowerLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewLowerLimit,
		})
		_node.OutcomeOverviewLowerLimit = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewGroupID,
		})
		_node.OutcomeOverviewGroupID = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewRatio(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewRatio,
		})
		_node.OutcomeOverviewRatio = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewMeasureTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewMeasureTitle,
		})
		_node.OutcomeOverviewMeasureTitle = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewVaccine(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewVaccine,
		})
		_node.OutcomeOverviewVaccine = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewImmunocompromisedPopulation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewImmunocompromisedPopulation,
		})
		_node.OutcomeOverviewImmunocompromisedPopulation = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewManufacturer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewManufacturer,
		})
		_node.OutcomeOverviewManufacturer = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewConfidenceInterval(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewConfidenceInterval,
		})
		_node.OutcomeOverviewConfidenceInterval = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewPercentResponders(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewPercentResponders,
		})
		_node.OutcomeOverviewPercentResponders = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewTimeFrameWeeks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewTimeFrameWeeks,
		})
		_node.OutcomeOverviewTimeFrameWeeks = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewDoseDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewDoseDescription,
		})
		_node.OutcomeOverviewDoseDescription = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewSchedule(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewSchedule,
		})
		_node.OutcomeOverviewSchedule = value
	}
	if value, ok := ooc.mutation.OutcomeOverviewUseCaseCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outcomeoverview.FieldOutcomeOverviewUseCaseCode,
		})
		_node.OutcomeOverviewUseCaseCode = value
	}
	if nodes := ooc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.outcome_measure_outcome_overview_list = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OutcomeOverviewCreateBulk is the builder for creating many OutcomeOverview entities in bulk.
type OutcomeOverviewCreateBulk struct {
	config
	builders []*OutcomeOverviewCreate
}

// Save creates the OutcomeOverview entities in the database.
func (oocb *OutcomeOverviewCreateBulk) Save(ctx context.Context) ([]*OutcomeOverview, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oocb.builders))
	nodes := make([]*OutcomeOverview, len(oocb.builders))
	mutators := make([]Mutator, len(oocb.builders))
	for i := range oocb.builders {
		func(i int, root context.Context) {
			builder := oocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutcomeOverviewMutation)
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
					_, err = mutators[i+1].Mutate(root, oocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oocb *OutcomeOverviewCreateBulk) SaveX(ctx context.Context) []*OutcomeOverview {
	v, err := oocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oocb *OutcomeOverviewCreateBulk) Exec(ctx context.Context) error {
	_, err := oocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oocb *OutcomeOverviewCreateBulk) ExecX(ctx context.Context) {
	if err := oocb.Exec(ctx); err != nil {
		panic(err)
	}
}
