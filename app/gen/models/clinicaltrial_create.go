// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// ClinicalTrialCreate is the builder for creating a ClinicalTrial entity.
type ClinicalTrialCreate struct {
	config
	mutation *ClinicalTrialMutation
	hooks    []Hook
}

// SetStudyName sets the "study_name" field.
func (ctc *ClinicalTrialCreate) SetStudyName(s string) *ClinicalTrialCreate {
	ctc.mutation.SetStudyName(s)
	return ctc
}

// SetSponsor sets the "sponsor" field.
func (ctc *ClinicalTrialCreate) SetSponsor(s string) *ClinicalTrialCreate {
	ctc.mutation.SetSponsor(s)
	return ctc
}

// SetResponsibleParty sets the "responsible_party" field.
func (ctc *ClinicalTrialCreate) SetResponsibleParty(s string) *ClinicalTrialCreate {
	ctc.mutation.SetResponsibleParty(s)
	return ctc
}

// SetStudyType sets the "study_type" field.
func (ctc *ClinicalTrialCreate) SetStudyType(s string) *ClinicalTrialCreate {
	ctc.mutation.SetStudyType(s)
	return ctc
}

// SetPhase sets the "phase" field.
func (ctc *ClinicalTrialCreate) SetPhase(s string) *ClinicalTrialCreate {
	ctc.mutation.SetPhase(s)
	return ctc
}

// SetActualEnrollment sets the "actual_enrollment" field.
func (ctc *ClinicalTrialCreate) SetActualEnrollment(s string) *ClinicalTrialCreate {
	ctc.mutation.SetActualEnrollment(s)
	return ctc
}

// SetAllocation sets the "allocation" field.
func (ctc *ClinicalTrialCreate) SetAllocation(s string) *ClinicalTrialCreate {
	ctc.mutation.SetAllocation(s)
	return ctc
}

// SetInterventionModel sets the "intervention_model" field.
func (ctc *ClinicalTrialCreate) SetInterventionModel(s string) *ClinicalTrialCreate {
	ctc.mutation.SetInterventionModel(s)
	return ctc
}

// SetMasking sets the "masking" field.
func (ctc *ClinicalTrialCreate) SetMasking(s string) *ClinicalTrialCreate {
	ctc.mutation.SetMasking(s)
	return ctc
}

// SetPrimaryPurpose sets the "primary_purpose" field.
func (ctc *ClinicalTrialCreate) SetPrimaryPurpose(s string) *ClinicalTrialCreate {
	ctc.mutation.SetPrimaryPurpose(s)
	return ctc
}

// SetOfficialTitle sets the "official_title" field.
func (ctc *ClinicalTrialCreate) SetOfficialTitle(s string) *ClinicalTrialCreate {
	ctc.mutation.SetOfficialTitle(s)
	return ctc
}

// SetActualStudyStartDate sets the "actual_study_start_date" field.
func (ctc *ClinicalTrialCreate) SetActualStudyStartDate(t time.Time) *ClinicalTrialCreate {
	ctc.mutation.SetActualStudyStartDate(t)
	return ctc
}

// SetActualPrimaryCompletionDate sets the "actual_primary_completion_date" field.
func (ctc *ClinicalTrialCreate) SetActualPrimaryCompletionDate(t time.Time) *ClinicalTrialCreate {
	ctc.mutation.SetActualPrimaryCompletionDate(t)
	return ctc
}

// SetActualStudyCompletionDate sets the "actual_study_completion_date" field.
func (ctc *ClinicalTrialCreate) SetActualStudyCompletionDate(t time.Time) *ClinicalTrialCreate {
	ctc.mutation.SetActualStudyCompletionDate(t)
	return ctc
}

// SetStudyID sets the "study_id" field.
func (ctc *ClinicalTrialCreate) SetStudyID(s string) *ClinicalTrialCreate {
	ctc.mutation.SetStudyID(s)
	return ctc
}

// SetResultsDefinitionID sets the "results_definition" edge to the ResultsDefinition entity by ID.
func (ctc *ClinicalTrialCreate) SetResultsDefinitionID(id int) *ClinicalTrialCreate {
	ctc.mutation.SetResultsDefinitionID(id)
	return ctc
}

// SetNillableResultsDefinitionID sets the "results_definition" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (ctc *ClinicalTrialCreate) SetNillableResultsDefinitionID(id *int) *ClinicalTrialCreate {
	if id != nil {
		ctc = ctc.SetResultsDefinitionID(*id)
	}
	return ctc
}

// SetResultsDefinition sets the "results_definition" edge to the ResultsDefinition entity.
func (ctc *ClinicalTrialCreate) SetResultsDefinition(r *ResultsDefinition) *ClinicalTrialCreate {
	return ctc.SetResultsDefinitionID(r.ID)
}

// AddStudyLocationIDs adds the "study_locations" edge to the StudyLocation entity by IDs.
func (ctc *ClinicalTrialCreate) AddStudyLocationIDs(ids ...int) *ClinicalTrialCreate {
	ctc.mutation.AddStudyLocationIDs(ids...)
	return ctc
}

// AddStudyLocations adds the "study_locations" edges to the StudyLocation entity.
func (ctc *ClinicalTrialCreate) AddStudyLocations(s ...*StudyLocation) *ClinicalTrialCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctc.AddStudyLocationIDs(ids...)
}

// AddStudyEligibilityIDs adds the "study_eligibility" edge to the StudyEligibility entity by IDs.
func (ctc *ClinicalTrialCreate) AddStudyEligibilityIDs(ids ...int) *ClinicalTrialCreate {
	ctc.mutation.AddStudyEligibilityIDs(ids...)
	return ctc
}

// AddStudyEligibility adds the "study_eligibility" edges to the StudyEligibility entity.
func (ctc *ClinicalTrialCreate) AddStudyEligibility(s ...*StudyEligibility) *ClinicalTrialCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctc.AddStudyEligibilityIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the ClinicalTrialMetadata entity by IDs.
func (ctc *ClinicalTrialCreate) AddMetadatumIDs(ids ...int) *ClinicalTrialCreate {
	ctc.mutation.AddMetadatumIDs(ids...)
	return ctc
}

// AddMetadata adds the "metadata" edges to the ClinicalTrialMetadata entity.
func (ctc *ClinicalTrialCreate) AddMetadata(c ...*ClinicalTrialMetadata) *ClinicalTrialCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctc.AddMetadatumIDs(ids...)
}

// Mutation returns the ClinicalTrialMutation object of the builder.
func (ctc *ClinicalTrialCreate) Mutation() *ClinicalTrialMutation {
	return ctc.mutation
}

// Save creates the ClinicalTrial in the database.
func (ctc *ClinicalTrialCreate) Save(ctx context.Context) (*ClinicalTrial, error) {
	var (
		err  error
		node *ClinicalTrial
	)
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			if node, err = ctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			if ctc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *ClinicalTrialCreate) SaveX(ctx context.Context) *ClinicalTrial {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *ClinicalTrialCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *ClinicalTrialCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *ClinicalTrialCreate) check() error {
	if _, ok := ctc.mutation.StudyName(); !ok {
		return &ValidationError{Name: "study_name", err: errors.New(`models: missing required field "ClinicalTrial.study_name"`)}
	}
	if _, ok := ctc.mutation.Sponsor(); !ok {
		return &ValidationError{Name: "sponsor", err: errors.New(`models: missing required field "ClinicalTrial.sponsor"`)}
	}
	if _, ok := ctc.mutation.ResponsibleParty(); !ok {
		return &ValidationError{Name: "responsible_party", err: errors.New(`models: missing required field "ClinicalTrial.responsible_party"`)}
	}
	if _, ok := ctc.mutation.StudyType(); !ok {
		return &ValidationError{Name: "study_type", err: errors.New(`models: missing required field "ClinicalTrial.study_type"`)}
	}
	if _, ok := ctc.mutation.Phase(); !ok {
		return &ValidationError{Name: "phase", err: errors.New(`models: missing required field "ClinicalTrial.phase"`)}
	}
	if _, ok := ctc.mutation.ActualEnrollment(); !ok {
		return &ValidationError{Name: "actual_enrollment", err: errors.New(`models: missing required field "ClinicalTrial.actual_enrollment"`)}
	}
	if _, ok := ctc.mutation.Allocation(); !ok {
		return &ValidationError{Name: "allocation", err: errors.New(`models: missing required field "ClinicalTrial.allocation"`)}
	}
	if _, ok := ctc.mutation.InterventionModel(); !ok {
		return &ValidationError{Name: "intervention_model", err: errors.New(`models: missing required field "ClinicalTrial.intervention_model"`)}
	}
	if _, ok := ctc.mutation.Masking(); !ok {
		return &ValidationError{Name: "masking", err: errors.New(`models: missing required field "ClinicalTrial.masking"`)}
	}
	if _, ok := ctc.mutation.PrimaryPurpose(); !ok {
		return &ValidationError{Name: "primary_purpose", err: errors.New(`models: missing required field "ClinicalTrial.primary_purpose"`)}
	}
	if _, ok := ctc.mutation.OfficialTitle(); !ok {
		return &ValidationError{Name: "official_title", err: errors.New(`models: missing required field "ClinicalTrial.official_title"`)}
	}
	if _, ok := ctc.mutation.ActualStudyStartDate(); !ok {
		return &ValidationError{Name: "actual_study_start_date", err: errors.New(`models: missing required field "ClinicalTrial.actual_study_start_date"`)}
	}
	if _, ok := ctc.mutation.ActualPrimaryCompletionDate(); !ok {
		return &ValidationError{Name: "actual_primary_completion_date", err: errors.New(`models: missing required field "ClinicalTrial.actual_primary_completion_date"`)}
	}
	if _, ok := ctc.mutation.ActualStudyCompletionDate(); !ok {
		return &ValidationError{Name: "actual_study_completion_date", err: errors.New(`models: missing required field "ClinicalTrial.actual_study_completion_date"`)}
	}
	if _, ok := ctc.mutation.StudyID(); !ok {
		return &ValidationError{Name: "study_id", err: errors.New(`models: missing required field "ClinicalTrial.study_id"`)}
	}
	return nil
}

func (ctc *ClinicalTrialCreate) sqlSave(ctx context.Context) (*ClinicalTrial, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctc *ClinicalTrialCreate) createSpec() (*ClinicalTrial, *sqlgraph.CreateSpec) {
	var (
		_node = &ClinicalTrial{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: clinicaltrial.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrial.FieldID,
			},
		}
	)
	if value, ok := ctc.mutation.StudyName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyName,
		})
		_node.StudyName = value
	}
	if value, ok := ctc.mutation.Sponsor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldSponsor,
		})
		_node.Sponsor = value
	}
	if value, ok := ctc.mutation.ResponsibleParty(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldResponsibleParty,
		})
		_node.ResponsibleParty = value
	}
	if value, ok := ctc.mutation.StudyType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyType,
		})
		_node.StudyType = value
	}
	if value, ok := ctc.mutation.Phase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldPhase,
		})
		_node.Phase = value
	}
	if value, ok := ctc.mutation.ActualEnrollment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldActualEnrollment,
		})
		_node.ActualEnrollment = value
	}
	if value, ok := ctc.mutation.Allocation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldAllocation,
		})
		_node.Allocation = value
	}
	if value, ok := ctc.mutation.InterventionModel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldInterventionModel,
		})
		_node.InterventionModel = value
	}
	if value, ok := ctc.mutation.Masking(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldMasking,
		})
		_node.Masking = value
	}
	if value, ok := ctc.mutation.PrimaryPurpose(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldPrimaryPurpose,
		})
		_node.PrimaryPurpose = value
	}
	if value, ok := ctc.mutation.OfficialTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldOfficialTitle,
		})
		_node.OfficialTitle = value
	}
	if value, ok := ctc.mutation.ActualStudyStartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualStudyStartDate,
		})
		_node.ActualStudyStartDate = value
	}
	if value, ok := ctc.mutation.ActualPrimaryCompletionDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualPrimaryCompletionDate,
		})
		_node.ActualPrimaryCompletionDate = value
	}
	if value, ok := ctc.mutation.ActualStudyCompletionDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualStudyCompletionDate,
		})
		_node.ActualStudyCompletionDate = value
	}
	if value, ok := ctc.mutation.StudyID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyID,
		})
		_node.StudyID = value
	}
	if nodes := ctc.mutation.ResultsDefinitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   clinicaltrial.ResultsDefinitionTable,
			Columns: []string{clinicaltrial.ResultsDefinitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resultsdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ctc.mutation.StudyLocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinicaltrial.StudyLocationsTable,
			Columns: []string{clinicaltrial.StudyLocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studylocation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ctc.mutation.StudyEligibilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinicaltrial.StudyEligibilityTable,
			Columns: []string{clinicaltrial.StudyEligibilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studyeligibility.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ctc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clinicaltrial.MetadataTable,
			Columns: []string{clinicaltrial.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clinicaltrialmetadata.FieldID,
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

// ClinicalTrialCreateBulk is the builder for creating many ClinicalTrial entities in bulk.
type ClinicalTrialCreateBulk struct {
	config
	builders []*ClinicalTrialCreate
}

// Save creates the ClinicalTrial entities in the database.
func (ctcb *ClinicalTrialCreateBulk) Save(ctx context.Context) ([]*ClinicalTrial, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*ClinicalTrial, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClinicalTrialMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *ClinicalTrialCreateBulk) SaveX(ctx context.Context) []*ClinicalTrial {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *ClinicalTrialCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *ClinicalTrialCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}
