// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/certainagreement"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/resultsdefinition"
)

// MoreInfoModuleUpdate is the builder for updating MoreInfoModule entities.
type MoreInfoModuleUpdate struct {
	config
	hooks    []Hook
	mutation *MoreInfoModuleMutation
}

// Where appends a list predicates to the MoreInfoModuleUpdate builder.
func (mimu *MoreInfoModuleUpdate) Where(ps ...predicate.MoreInfoModule) *MoreInfoModuleUpdate {
	mimu.mutation.Where(ps...)
	return mimu
}

// SetLimitationsAndCaveatsDescription sets the "limitations_and_caveats_description" field.
func (mimu *MoreInfoModuleUpdate) SetLimitationsAndCaveatsDescription(s string) *MoreInfoModuleUpdate {
	mimu.mutation.SetLimitationsAndCaveatsDescription(s)
	return mimu
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (mimu *MoreInfoModuleUpdate) SetParentID(id int) *MoreInfoModuleUpdate {
	mimu.mutation.SetParentID(id)
	return mimu
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (mimu *MoreInfoModuleUpdate) SetNillableParentID(id *int) *MoreInfoModuleUpdate {
	if id != nil {
		mimu = mimu.SetParentID(*id)
	}
	return mimu
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (mimu *MoreInfoModuleUpdate) SetParent(r *ResultsDefinition) *MoreInfoModuleUpdate {
	return mimu.SetParentID(r.ID)
}

// SetCertainAgreementID sets the "certain_agreement" edge to the CertainAgreement entity by ID.
func (mimu *MoreInfoModuleUpdate) SetCertainAgreementID(id int) *MoreInfoModuleUpdate {
	mimu.mutation.SetCertainAgreementID(id)
	return mimu
}

// SetNillableCertainAgreementID sets the "certain_agreement" edge to the CertainAgreement entity by ID if the given value is not nil.
func (mimu *MoreInfoModuleUpdate) SetNillableCertainAgreementID(id *int) *MoreInfoModuleUpdate {
	if id != nil {
		mimu = mimu.SetCertainAgreementID(*id)
	}
	return mimu
}

// SetCertainAgreement sets the "certain_agreement" edge to the CertainAgreement entity.
func (mimu *MoreInfoModuleUpdate) SetCertainAgreement(c *CertainAgreement) *MoreInfoModuleUpdate {
	return mimu.SetCertainAgreementID(c.ID)
}

// SetPointOfContactID sets the "point_of_contact" edge to the PointOfContact entity by ID.
func (mimu *MoreInfoModuleUpdate) SetPointOfContactID(id int) *MoreInfoModuleUpdate {
	mimu.mutation.SetPointOfContactID(id)
	return mimu
}

// SetNillablePointOfContactID sets the "point_of_contact" edge to the PointOfContact entity by ID if the given value is not nil.
func (mimu *MoreInfoModuleUpdate) SetNillablePointOfContactID(id *int) *MoreInfoModuleUpdate {
	if id != nil {
		mimu = mimu.SetPointOfContactID(*id)
	}
	return mimu
}

// SetPointOfContact sets the "point_of_contact" edge to the PointOfContact entity.
func (mimu *MoreInfoModuleUpdate) SetPointOfContact(p *PointOfContact) *MoreInfoModuleUpdate {
	return mimu.SetPointOfContactID(p.ID)
}

// Mutation returns the MoreInfoModuleMutation object of the builder.
func (mimu *MoreInfoModuleUpdate) Mutation() *MoreInfoModuleMutation {
	return mimu.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (mimu *MoreInfoModuleUpdate) ClearParent() *MoreInfoModuleUpdate {
	mimu.mutation.ClearParent()
	return mimu
}

// ClearCertainAgreement clears the "certain_agreement" edge to the CertainAgreement entity.
func (mimu *MoreInfoModuleUpdate) ClearCertainAgreement() *MoreInfoModuleUpdate {
	mimu.mutation.ClearCertainAgreement()
	return mimu
}

// ClearPointOfContact clears the "point_of_contact" edge to the PointOfContact entity.
func (mimu *MoreInfoModuleUpdate) ClearPointOfContact() *MoreInfoModuleUpdate {
	mimu.mutation.ClearPointOfContact()
	return mimu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mimu *MoreInfoModuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mimu.hooks) == 0 {
		affected, err = mimu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MoreInfoModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mimu.mutation = mutation
			affected, err = mimu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mimu.hooks) - 1; i >= 0; i-- {
			if mimu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = mimu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mimu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mimu *MoreInfoModuleUpdate) SaveX(ctx context.Context) int {
	affected, err := mimu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mimu *MoreInfoModuleUpdate) Exec(ctx context.Context) error {
	_, err := mimu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mimu *MoreInfoModuleUpdate) ExecX(ctx context.Context) {
	if err := mimu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mimu *MoreInfoModuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moreinfomodule.Table,
			Columns: moreinfomodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moreinfomodule.FieldID,
			},
		},
	}
	if ps := mimu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mimu.mutation.LimitationsAndCaveatsDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moreinfomodule.FieldLimitationsAndCaveatsDescription,
		})
	}
	if mimu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   moreinfomodule.ParentTable,
			Columns: []string{moreinfomodule.ParentColumn},
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
	if nodes := mimu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   moreinfomodule.ParentTable,
			Columns: []string{moreinfomodule.ParentColumn},
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
	if mimu.mutation.CertainAgreementCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.CertainAgreementTable,
			Columns: []string{moreinfomodule.CertainAgreementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: certainagreement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mimu.mutation.CertainAgreementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.CertainAgreementTable,
			Columns: []string{moreinfomodule.CertainAgreementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: certainagreement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mimu.mutation.PointOfContactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.PointOfContactTable,
			Columns: []string{moreinfomodule.PointOfContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pointofcontact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mimu.mutation.PointOfContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.PointOfContactTable,
			Columns: []string{moreinfomodule.PointOfContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pointofcontact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mimu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moreinfomodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MoreInfoModuleUpdateOne is the builder for updating a single MoreInfoModule entity.
type MoreInfoModuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MoreInfoModuleMutation
}

// SetLimitationsAndCaveatsDescription sets the "limitations_and_caveats_description" field.
func (mimuo *MoreInfoModuleUpdateOne) SetLimitationsAndCaveatsDescription(s string) *MoreInfoModuleUpdateOne {
	mimuo.mutation.SetLimitationsAndCaveatsDescription(s)
	return mimuo
}

// SetParentID sets the "parent" edge to the ResultsDefinition entity by ID.
func (mimuo *MoreInfoModuleUpdateOne) SetParentID(id int) *MoreInfoModuleUpdateOne {
	mimuo.mutation.SetParentID(id)
	return mimuo
}

// SetNillableParentID sets the "parent" edge to the ResultsDefinition entity by ID if the given value is not nil.
func (mimuo *MoreInfoModuleUpdateOne) SetNillableParentID(id *int) *MoreInfoModuleUpdateOne {
	if id != nil {
		mimuo = mimuo.SetParentID(*id)
	}
	return mimuo
}

// SetParent sets the "parent" edge to the ResultsDefinition entity.
func (mimuo *MoreInfoModuleUpdateOne) SetParent(r *ResultsDefinition) *MoreInfoModuleUpdateOne {
	return mimuo.SetParentID(r.ID)
}

// SetCertainAgreementID sets the "certain_agreement" edge to the CertainAgreement entity by ID.
func (mimuo *MoreInfoModuleUpdateOne) SetCertainAgreementID(id int) *MoreInfoModuleUpdateOne {
	mimuo.mutation.SetCertainAgreementID(id)
	return mimuo
}

// SetNillableCertainAgreementID sets the "certain_agreement" edge to the CertainAgreement entity by ID if the given value is not nil.
func (mimuo *MoreInfoModuleUpdateOne) SetNillableCertainAgreementID(id *int) *MoreInfoModuleUpdateOne {
	if id != nil {
		mimuo = mimuo.SetCertainAgreementID(*id)
	}
	return mimuo
}

// SetCertainAgreement sets the "certain_agreement" edge to the CertainAgreement entity.
func (mimuo *MoreInfoModuleUpdateOne) SetCertainAgreement(c *CertainAgreement) *MoreInfoModuleUpdateOne {
	return mimuo.SetCertainAgreementID(c.ID)
}

// SetPointOfContactID sets the "point_of_contact" edge to the PointOfContact entity by ID.
func (mimuo *MoreInfoModuleUpdateOne) SetPointOfContactID(id int) *MoreInfoModuleUpdateOne {
	mimuo.mutation.SetPointOfContactID(id)
	return mimuo
}

// SetNillablePointOfContactID sets the "point_of_contact" edge to the PointOfContact entity by ID if the given value is not nil.
func (mimuo *MoreInfoModuleUpdateOne) SetNillablePointOfContactID(id *int) *MoreInfoModuleUpdateOne {
	if id != nil {
		mimuo = mimuo.SetPointOfContactID(*id)
	}
	return mimuo
}

// SetPointOfContact sets the "point_of_contact" edge to the PointOfContact entity.
func (mimuo *MoreInfoModuleUpdateOne) SetPointOfContact(p *PointOfContact) *MoreInfoModuleUpdateOne {
	return mimuo.SetPointOfContactID(p.ID)
}

// Mutation returns the MoreInfoModuleMutation object of the builder.
func (mimuo *MoreInfoModuleUpdateOne) Mutation() *MoreInfoModuleMutation {
	return mimuo.mutation
}

// ClearParent clears the "parent" edge to the ResultsDefinition entity.
func (mimuo *MoreInfoModuleUpdateOne) ClearParent() *MoreInfoModuleUpdateOne {
	mimuo.mutation.ClearParent()
	return mimuo
}

// ClearCertainAgreement clears the "certain_agreement" edge to the CertainAgreement entity.
func (mimuo *MoreInfoModuleUpdateOne) ClearCertainAgreement() *MoreInfoModuleUpdateOne {
	mimuo.mutation.ClearCertainAgreement()
	return mimuo
}

// ClearPointOfContact clears the "point_of_contact" edge to the PointOfContact entity.
func (mimuo *MoreInfoModuleUpdateOne) ClearPointOfContact() *MoreInfoModuleUpdateOne {
	mimuo.mutation.ClearPointOfContact()
	return mimuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mimuo *MoreInfoModuleUpdateOne) Select(field string, fields ...string) *MoreInfoModuleUpdateOne {
	mimuo.fields = append([]string{field}, fields...)
	return mimuo
}

// Save executes the query and returns the updated MoreInfoModule entity.
func (mimuo *MoreInfoModuleUpdateOne) Save(ctx context.Context) (*MoreInfoModule, error) {
	var (
		err  error
		node *MoreInfoModule
	)
	if len(mimuo.hooks) == 0 {
		node, err = mimuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MoreInfoModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mimuo.mutation = mutation
			node, err = mimuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mimuo.hooks) - 1; i >= 0; i-- {
			if mimuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = mimuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mimuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mimuo *MoreInfoModuleUpdateOne) SaveX(ctx context.Context) *MoreInfoModule {
	node, err := mimuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mimuo *MoreInfoModuleUpdateOne) Exec(ctx context.Context) error {
	_, err := mimuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mimuo *MoreInfoModuleUpdateOne) ExecX(ctx context.Context) {
	if err := mimuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mimuo *MoreInfoModuleUpdateOne) sqlSave(ctx context.Context) (_node *MoreInfoModule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moreinfomodule.Table,
			Columns: moreinfomodule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moreinfomodule.FieldID,
			},
		},
	}
	id, ok := mimuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "MoreInfoModule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mimuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moreinfomodule.FieldID)
		for _, f := range fields {
			if !moreinfomodule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != moreinfomodule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mimuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mimuo.mutation.LimitationsAndCaveatsDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moreinfomodule.FieldLimitationsAndCaveatsDescription,
		})
	}
	if mimuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   moreinfomodule.ParentTable,
			Columns: []string{moreinfomodule.ParentColumn},
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
	if nodes := mimuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   moreinfomodule.ParentTable,
			Columns: []string{moreinfomodule.ParentColumn},
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
	if mimuo.mutation.CertainAgreementCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.CertainAgreementTable,
			Columns: []string{moreinfomodule.CertainAgreementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: certainagreement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mimuo.mutation.CertainAgreementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.CertainAgreementTable,
			Columns: []string{moreinfomodule.CertainAgreementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: certainagreement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mimuo.mutation.PointOfContactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.PointOfContactTable,
			Columns: []string{moreinfomodule.PointOfContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pointofcontact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mimuo.mutation.PointOfContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   moreinfomodule.PointOfContactTable,
			Columns: []string{moreinfomodule.PointOfContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pointofcontact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MoreInfoModule{config: mimuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mimuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moreinfomodule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
