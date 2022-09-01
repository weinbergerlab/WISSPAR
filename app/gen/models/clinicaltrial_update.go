// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrial"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/clinicaltrialmetadata"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studyeligibility"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/studylocation"
)

// ClinicalTrialUpdate is the builder for updating ClinicalTrial entities.
type ClinicalTrialUpdate struct {
	config
	hooks    []Hook
	mutation *ClinicalTrialMutation
}

// Where appends a list predicates to the ClinicalTrialUpdate builder.
func (ctu *ClinicalTrialUpdate) Where(ps ...predicate.ClinicalTrial) *ClinicalTrialUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetStudyName sets the "study_name" field.
func (ctu *ClinicalTrialUpdate) SetStudyName(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetStudyName(s)
	return ctu
}

// SetSponsor sets the "sponsor" field.
func (ctu *ClinicalTrialUpdate) SetSponsor(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetSponsor(s)
	return ctu
}

// SetResponsibleParty sets the "responsible_party" field.
func (ctu *ClinicalTrialUpdate) SetResponsibleParty(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetResponsibleParty(s)
	return ctu
}

// SetStudyType sets the "study_type" field.
func (ctu *ClinicalTrialUpdate) SetStudyType(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetStudyType(s)
	return ctu
}

// SetPhase sets the "phase" field.
func (ctu *ClinicalTrialUpdate) SetPhase(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetPhase(s)
	return ctu
}

// SetActualEnrollment sets the "actual_enrollment" field.
func (ctu *ClinicalTrialUpdate) SetActualEnrollment(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetActualEnrollment(s)
	return ctu
}

// SetAllocation sets the "allocation" field.
func (ctu *ClinicalTrialUpdate) SetAllocation(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetAllocation(s)
	return ctu
}

// SetInterventionModel sets the "intervention_model" field.
func (ctu *ClinicalTrialUpdate) SetInterventionModel(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetInterventionModel(s)
	return ctu
}

// SetMasking sets the "masking" field.
func (ctu *ClinicalTrialUpdate) SetMasking(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetMasking(s)
	return ctu
}

// SetPrimaryPurpose sets the "primary_purpose" field.
func (ctu *ClinicalTrialUpdate) SetPrimaryPurpose(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetPrimaryPurpose(s)
	return ctu
}

// SetOfficialTitle sets the "official_title" field.
func (ctu *ClinicalTrialUpdate) SetOfficialTitle(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetOfficialTitle(s)
	return ctu
}

// SetActualStudyStartDate sets the "actual_study_start_date" field.
func (ctu *ClinicalTrialUpdate) SetActualStudyStartDate(t time.Time) *ClinicalTrialUpdate {
	ctu.mutation.SetActualStudyStartDate(t)
	return ctu
}

// SetActualPrimaryCompletionDate sets the "actual_primary_completion_date" field.
func (ctu *ClinicalTrialUpdate) SetActualPrimaryCompletionDate(t time.Time) *ClinicalTrialUpdate {
	ctu.mutation.SetActualPrimaryCompletionDate(t)
	return ctu
}

// SetActualStudyCompletionDate sets the "actual_study_completion_date" field.
func (ctu *ClinicalTrialUpdate) SetActualStudyCompletionDate(t time.Time) *ClinicalTrialUpdate {
	ctu.mutation.SetActualStudyCompletionDate(t)
	return ctu
}

// SetStudyID sets the "study_id" field.
func (ctu *ClinicalTrialUpdate) SetStudyID(s string) *ClinicalTrialUpdate {
	ctu.mutation.SetStudyID(s)
	return ctu
}

// SetResultsDefinitionID sets the "results_definition" edge to the ResultsDefinition entity by ID.
func (ctu *ClinicalTrialUpdate) SetResultsDefinitionID(id int) *ClinicalTrialUpdate {
	ctu.mutation.SetResultsDefinitionID(id)
	return ctu
}

// SetNillableResultsDefinitionID sets the "results_definition" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (ctu *ClinicalTrialUpdate) SetNillableResultsDefinitionID(id *int) *ClinicalTrialUpdate {
	if id != nil {
		ctu = ctu.SetResultsDefinitionID(*id)
	}
	return ctu
}

// SetResultsDefinition sets the "results_definition" edge to the ResultsDefinition entity.
func (ctu *ClinicalTrialUpdate) SetResultsDefinition(r *ResultsDefinition) *ClinicalTrialUpdate {
	return ctu.SetResultsDefinitionID(r.ID)
}

// AddStudyLocationIDs adds the "study_locations" edge to the StudyLocation entity by IDs.
func (ctu *ClinicalTrialUpdate) AddStudyLocationIDs(ids ...int) *ClinicalTrialUpdate {
	ctu.mutation.AddStudyLocationIDs(ids...)
	return ctu
}

// AddStudyLocations adds the "study_locations" edges to the StudyLocation entity.
func (ctu *ClinicalTrialUpdate) AddStudyLocations(s ...*StudyLocation) *ClinicalTrialUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctu.AddStudyLocationIDs(ids...)
}

// AddStudyEligibilityIDs adds the "study_eligibility" edge to the StudyEligibility entity by IDs.
func (ctu *ClinicalTrialUpdate) AddStudyEligibilityIDs(ids ...int) *ClinicalTrialUpdate {
	ctu.mutation.AddStudyEligibilityIDs(ids...)
	return ctu
}

// AddStudyEligibility adds the "study_eligibility" edges to the StudyEligibility entity.
func (ctu *ClinicalTrialUpdate) AddStudyEligibility(s ...*StudyEligibility) *ClinicalTrialUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctu.AddStudyEligibilityIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the ClinicalTrialMetadata entity by IDs.
func (ctu *ClinicalTrialUpdate) AddMetadatumIDs(ids ...int) *ClinicalTrialUpdate {
	ctu.mutation.AddMetadatumIDs(ids...)
	return ctu
}

// AddMetadata adds the "metadata" edges to the ClinicalTrialMetadata entity.
func (ctu *ClinicalTrialUpdate) AddMetadata(c ...*ClinicalTrialMetadata) *ClinicalTrialUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.AddMetadatumIDs(ids...)
}

// Mutation returns the ClinicalTrialMutation object of the builder.
func (ctu *ClinicalTrialUpdate) Mutation() *ClinicalTrialMutation {
	return ctu.mutation
}

// ClearResultsDefinition clears the "results_definition" edge to the ResultsDefinition entity.
func (ctu *ClinicalTrialUpdate) ClearResultsDefinition() *ClinicalTrialUpdate {
	ctu.mutation.ClearResultsDefinition()
	return ctu
}

// ClearStudyLocations clears all "study_locations" edges to the StudyLocation entity.
func (ctu *ClinicalTrialUpdate) ClearStudyLocations() *ClinicalTrialUpdate {
	ctu.mutation.ClearStudyLocations()
	return ctu
}

// RemoveStudyLocationIDs removes the "study_locations" edge to StudyLocation entities by IDs.
func (ctu *ClinicalTrialUpdate) RemoveStudyLocationIDs(ids ...int) *ClinicalTrialUpdate {
	ctu.mutation.RemoveStudyLocationIDs(ids...)
	return ctu
}

// RemoveStudyLocations removes "study_locations" edges to StudyLocation entities.
func (ctu *ClinicalTrialUpdate) RemoveStudyLocations(s ...*StudyLocation) *ClinicalTrialUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctu.RemoveStudyLocationIDs(ids...)
}

// ClearStudyEligibility clears all "study_eligibility" edges to the StudyEligibility entity.
func (ctu *ClinicalTrialUpdate) ClearStudyEligibility() *ClinicalTrialUpdate {
	ctu.mutation.ClearStudyEligibility()
	return ctu
}

// RemoveStudyEligibilityIDs removes the "study_eligibility" edge to StudyEligibility entities by IDs.
func (ctu *ClinicalTrialUpdate) RemoveStudyEligibilityIDs(ids ...int) *ClinicalTrialUpdate {
	ctu.mutation.RemoveStudyEligibilityIDs(ids...)
	return ctu
}

// RemoveStudyEligibility removes "study_eligibility" edges to StudyEligibility entities.
func (ctu *ClinicalTrialUpdate) RemoveStudyEligibility(s ...*StudyEligibility) *ClinicalTrialUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctu.RemoveStudyEligibilityIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the ClinicalTrialMetadata entity.
func (ctu *ClinicalTrialUpdate) ClearMetadata() *ClinicalTrialUpdate {
	ctu.mutation.ClearMetadata()
	return ctu
}

// RemoveMetadatumIDs removes the "metadata" edge to ClinicalTrialMetadata entities by IDs.
func (ctu *ClinicalTrialUpdate) RemoveMetadatumIDs(ids ...int) *ClinicalTrialUpdate {
	ctu.mutation.RemoveMetadatumIDs(ids...)
	return ctu
}

// RemoveMetadata removes "metadata" edges to ClinicalTrialMetadata entities.
func (ctu *ClinicalTrialUpdate) RemoveMetadata(c ...*ClinicalTrialMetadata) *ClinicalTrialUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.RemoveMetadatumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *ClinicalTrialUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *ClinicalTrialUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *ClinicalTrialUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *ClinicalTrialUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctu *ClinicalTrialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicaltrial.Table,
			Columns: clinicaltrial.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrial.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.StudyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyName,
		})
	}
	if value, ok := ctu.mutation.Sponsor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldSponsor,
		})
	}
	if value, ok := ctu.mutation.ResponsibleParty(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldResponsibleParty,
		})
	}
	if value, ok := ctu.mutation.StudyType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyType,
		})
	}
	if value, ok := ctu.mutation.Phase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldPhase,
		})
	}
	if value, ok := ctu.mutation.ActualEnrollment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldActualEnrollment,
		})
	}
	if value, ok := ctu.mutation.Allocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldAllocation,
		})
	}
	if value, ok := ctu.mutation.InterventionModel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldInterventionModel,
		})
	}
	if value, ok := ctu.mutation.Masking(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldMasking,
		})
	}
	if value, ok := ctu.mutation.PrimaryPurpose(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldPrimaryPurpose,
		})
	}
	if value, ok := ctu.mutation.OfficialTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldOfficialTitle,
		})
	}
	if value, ok := ctu.mutation.ActualStudyStartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualStudyStartDate,
		})
	}
	if value, ok := ctu.mutation.ActualPrimaryCompletionDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualPrimaryCompletionDate,
		})
	}
	if value, ok := ctu.mutation.ActualStudyCompletionDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualStudyCompletionDate,
		})
	}
	if value, ok := ctu.mutation.StudyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyID,
		})
	}
	if ctu.mutation.ResultsDefinitionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.ResultsDefinitionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctu.mutation.StudyLocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedStudyLocationsIDs(); len(nodes) > 0 && !ctu.mutation.StudyLocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.StudyLocationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctu.mutation.StudyEligibilityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedStudyEligibilityIDs(); len(nodes) > 0 && !ctu.mutation.StudyEligibilityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.StudyEligibilityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctu.mutation.MetadataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !ctu.mutation.MetadataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.MetadataIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinicaltrial.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClinicalTrialUpdateOne is the builder for updating a single ClinicalTrial entity.
type ClinicalTrialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClinicalTrialMutation
}

// SetStudyName sets the "study_name" field.
func (ctuo *ClinicalTrialUpdateOne) SetStudyName(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetStudyName(s)
	return ctuo
}

// SetSponsor sets the "sponsor" field.
func (ctuo *ClinicalTrialUpdateOne) SetSponsor(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetSponsor(s)
	return ctuo
}

// SetResponsibleParty sets the "responsible_party" field.
func (ctuo *ClinicalTrialUpdateOne) SetResponsibleParty(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetResponsibleParty(s)
	return ctuo
}

// SetStudyType sets the "study_type" field.
func (ctuo *ClinicalTrialUpdateOne) SetStudyType(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetStudyType(s)
	return ctuo
}

// SetPhase sets the "phase" field.
func (ctuo *ClinicalTrialUpdateOne) SetPhase(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetPhase(s)
	return ctuo
}

// SetActualEnrollment sets the "actual_enrollment" field.
func (ctuo *ClinicalTrialUpdateOne) SetActualEnrollment(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetActualEnrollment(s)
	return ctuo
}

// SetAllocation sets the "allocation" field.
func (ctuo *ClinicalTrialUpdateOne) SetAllocation(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetAllocation(s)
	return ctuo
}

// SetInterventionModel sets the "intervention_model" field.
func (ctuo *ClinicalTrialUpdateOne) SetInterventionModel(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetInterventionModel(s)
	return ctuo
}

// SetMasking sets the "masking" field.
func (ctuo *ClinicalTrialUpdateOne) SetMasking(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetMasking(s)
	return ctuo
}

// SetPrimaryPurpose sets the "primary_purpose" field.
func (ctuo *ClinicalTrialUpdateOne) SetPrimaryPurpose(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetPrimaryPurpose(s)
	return ctuo
}

// SetOfficialTitle sets the "official_title" field.
func (ctuo *ClinicalTrialUpdateOne) SetOfficialTitle(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetOfficialTitle(s)
	return ctuo
}

// SetActualStudyStartDate sets the "actual_study_start_date" field.
func (ctuo *ClinicalTrialUpdateOne) SetActualStudyStartDate(t time.Time) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetActualStudyStartDate(t)
	return ctuo
}

// SetActualPrimaryCompletionDate sets the "actual_primary_completion_date" field.
func (ctuo *ClinicalTrialUpdateOne) SetActualPrimaryCompletionDate(t time.Time) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetActualPrimaryCompletionDate(t)
	return ctuo
}

// SetActualStudyCompletionDate sets the "actual_study_completion_date" field.
func (ctuo *ClinicalTrialUpdateOne) SetActualStudyCompletionDate(t time.Time) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetActualStudyCompletionDate(t)
	return ctuo
}

// SetStudyID sets the "study_id" field.
func (ctuo *ClinicalTrialUpdateOne) SetStudyID(s string) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetStudyID(s)
	return ctuo
}

// SetResultsDefinitionID sets the "results_definition" edge to the ResultsDefinition entity by ID.
func (ctuo *ClinicalTrialUpdateOne) SetResultsDefinitionID(id int) *ClinicalTrialUpdateOne {
	ctuo.mutation.SetResultsDefinitionID(id)
	return ctuo
}

// SetNillableResultsDefinitionID sets the "results_definition" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (ctuo *ClinicalTrialUpdateOne) SetNillableResultsDefinitionID(id *int) *ClinicalTrialUpdateOne {
	if id != nil {
		ctuo = ctuo.SetResultsDefinitionID(*id)
	}
	return ctuo
}

// SetResultsDefinition sets the "results_definition" edge to the ResultsDefinition entity.
func (ctuo *ClinicalTrialUpdateOne) SetResultsDefinition(r *ResultsDefinition) *ClinicalTrialUpdateOne {
	return ctuo.SetResultsDefinitionID(r.ID)
}

// AddStudyLocationIDs adds the "study_locations" edge to the StudyLocation entity by IDs.
func (ctuo *ClinicalTrialUpdateOne) AddStudyLocationIDs(ids ...int) *ClinicalTrialUpdateOne {
	ctuo.mutation.AddStudyLocationIDs(ids...)
	return ctuo
}

// AddStudyLocations adds the "study_locations" edges to the StudyLocation entity.
func (ctuo *ClinicalTrialUpdateOne) AddStudyLocations(s ...*StudyLocation) *ClinicalTrialUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctuo.AddStudyLocationIDs(ids...)
}

// AddStudyEligibilityIDs adds the "study_eligibility" edge to the StudyEligibility entity by IDs.
func (ctuo *ClinicalTrialUpdateOne) AddStudyEligibilityIDs(ids ...int) *ClinicalTrialUpdateOne {
	ctuo.mutation.AddStudyEligibilityIDs(ids...)
	return ctuo
}

// AddStudyEligibility adds the "study_eligibility" edges to the StudyEligibility entity.
func (ctuo *ClinicalTrialUpdateOne) AddStudyEligibility(s ...*StudyEligibility) *ClinicalTrialUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctuo.AddStudyEligibilityIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the ClinicalTrialMetadata entity by IDs.
func (ctuo *ClinicalTrialUpdateOne) AddMetadatumIDs(ids ...int) *ClinicalTrialUpdateOne {
	ctuo.mutation.AddMetadatumIDs(ids...)
	return ctuo
}

// AddMetadata adds the "metadata" edges to the ClinicalTrialMetadata entity.
func (ctuo *ClinicalTrialUpdateOne) AddMetadata(c ...*ClinicalTrialMetadata) *ClinicalTrialUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.AddMetadatumIDs(ids...)
}

// Mutation returns the ClinicalTrialMutation object of the builder.
func (ctuo *ClinicalTrialUpdateOne) Mutation() *ClinicalTrialMutation {
	return ctuo.mutation
}

// ClearResultsDefinition clears the "results_definition" edge to the ResultsDefinition entity.
func (ctuo *ClinicalTrialUpdateOne) ClearResultsDefinition() *ClinicalTrialUpdateOne {
	ctuo.mutation.ClearResultsDefinition()
	return ctuo
}

// ClearStudyLocations clears all "study_locations" edges to the StudyLocation entity.
func (ctuo *ClinicalTrialUpdateOne) ClearStudyLocations() *ClinicalTrialUpdateOne {
	ctuo.mutation.ClearStudyLocations()
	return ctuo
}

// RemoveStudyLocationIDs removes the "study_locations" edge to StudyLocation entities by IDs.
func (ctuo *ClinicalTrialUpdateOne) RemoveStudyLocationIDs(ids ...int) *ClinicalTrialUpdateOne {
	ctuo.mutation.RemoveStudyLocationIDs(ids...)
	return ctuo
}

// RemoveStudyLocations removes "study_locations" edges to StudyLocation entities.
func (ctuo *ClinicalTrialUpdateOne) RemoveStudyLocations(s ...*StudyLocation) *ClinicalTrialUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctuo.RemoveStudyLocationIDs(ids...)
}

// ClearStudyEligibility clears all "study_eligibility" edges to the StudyEligibility entity.
func (ctuo *ClinicalTrialUpdateOne) ClearStudyEligibility() *ClinicalTrialUpdateOne {
	ctuo.mutation.ClearStudyEligibility()
	return ctuo
}

// RemoveStudyEligibilityIDs removes the "study_eligibility" edge to StudyEligibility entities by IDs.
func (ctuo *ClinicalTrialUpdateOne) RemoveStudyEligibilityIDs(ids ...int) *ClinicalTrialUpdateOne {
	ctuo.mutation.RemoveStudyEligibilityIDs(ids...)
	return ctuo
}

// RemoveStudyEligibility removes "study_eligibility" edges to StudyEligibility entities.
func (ctuo *ClinicalTrialUpdateOne) RemoveStudyEligibility(s ...*StudyEligibility) *ClinicalTrialUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ctuo.RemoveStudyEligibilityIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the ClinicalTrialMetadata entity.
func (ctuo *ClinicalTrialUpdateOne) ClearMetadata() *ClinicalTrialUpdateOne {
	ctuo.mutation.ClearMetadata()
	return ctuo
}

// RemoveMetadatumIDs removes the "metadata" edge to ClinicalTrialMetadata entities by IDs.
func (ctuo *ClinicalTrialUpdateOne) RemoveMetadatumIDs(ids ...int) *ClinicalTrialUpdateOne {
	ctuo.mutation.RemoveMetadatumIDs(ids...)
	return ctuo
}

// RemoveMetadata removes "metadata" edges to ClinicalTrialMetadata entities.
func (ctuo *ClinicalTrialUpdateOne) RemoveMetadata(c ...*ClinicalTrialMetadata) *ClinicalTrialUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.RemoveMetadatumIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *ClinicalTrialUpdateOne) Select(field string, fields ...string) *ClinicalTrialUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated ClinicalTrial entity.
func (ctuo *ClinicalTrialUpdateOne) Save(ctx context.Context) (*ClinicalTrial, error) {
	var (
		err  error
		node *ClinicalTrial
	)
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClinicalTrialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *ClinicalTrialUpdateOne) SaveX(ctx context.Context) *ClinicalTrial {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *ClinicalTrialUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *ClinicalTrialUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctuo *ClinicalTrialUpdateOne) sqlSave(ctx context.Context) (_node *ClinicalTrial, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicaltrial.Table,
			Columns: clinicaltrial.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clinicaltrial.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "ClinicalTrial.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clinicaltrial.FieldID)
		for _, f := range fields {
			if !clinicaltrial.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != clinicaltrial.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.StudyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyName,
		})
	}
	if value, ok := ctuo.mutation.Sponsor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldSponsor,
		})
	}
	if value, ok := ctuo.mutation.ResponsibleParty(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldResponsibleParty,
		})
	}
	if value, ok := ctuo.mutation.StudyType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyType,
		})
	}
	if value, ok := ctuo.mutation.Phase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldPhase,
		})
	}
	if value, ok := ctuo.mutation.ActualEnrollment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldActualEnrollment,
		})
	}
	if value, ok := ctuo.mutation.Allocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldAllocation,
		})
	}
	if value, ok := ctuo.mutation.InterventionModel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldInterventionModel,
		})
	}
	if value, ok := ctuo.mutation.Masking(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldMasking,
		})
	}
	if value, ok := ctuo.mutation.PrimaryPurpose(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldPrimaryPurpose,
		})
	}
	if value, ok := ctuo.mutation.OfficialTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldOfficialTitle,
		})
	}
	if value, ok := ctuo.mutation.ActualStudyStartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualStudyStartDate,
		})
	}
	if value, ok := ctuo.mutation.ActualPrimaryCompletionDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualPrimaryCompletionDate,
		})
	}
	if value, ok := ctuo.mutation.ActualStudyCompletionDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clinicaltrial.FieldActualStudyCompletionDate,
		})
	}
	if value, ok := ctuo.mutation.StudyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clinicaltrial.FieldStudyID,
		})
	}
	if ctuo.mutation.ResultsDefinitionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.ResultsDefinitionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctuo.mutation.StudyLocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedStudyLocationsIDs(); len(nodes) > 0 && !ctuo.mutation.StudyLocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.StudyLocationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctuo.mutation.StudyEligibilityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedStudyEligibilityIDs(); len(nodes) > 0 && !ctuo.mutation.StudyEligibilityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.StudyEligibilityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctuo.mutation.MetadataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !ctuo.mutation.MetadataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.MetadataIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ClinicalTrial{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinicaltrial.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
