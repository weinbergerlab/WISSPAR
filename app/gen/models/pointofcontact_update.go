// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/moreinfomodule"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/pointofcontact"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models/predicate"
)

// PointOfContactUpdate is the builder for updating PointOfContact entities.
type PointOfContactUpdate struct {
	config
	hooks    []Hook
	mutation *PointOfContactMutation
}

// Where appends a list predicates to the PointOfContactUpdate builder.
func (pocu *PointOfContactUpdate) Where(ps ...predicate.PointOfContact) *PointOfContactUpdate {
	pocu.mutation.Where(ps...)
	return pocu
}

// SetPointOfContactTitle sets the "point_of_contact_title" field.
func (pocu *PointOfContactUpdate) SetPointOfContactTitle(s string) *PointOfContactUpdate {
	pocu.mutation.SetPointOfContactTitle(s)
	return pocu
}

// SetPointOfContactOrganization sets the "point_of_contact_organization" field.
func (pocu *PointOfContactUpdate) SetPointOfContactOrganization(s string) *PointOfContactUpdate {
	pocu.mutation.SetPointOfContactOrganization(s)
	return pocu
}

// SetPointOfContactEmail sets the "point_of_contact_email" field.
func (pocu *PointOfContactUpdate) SetPointOfContactEmail(s string) *PointOfContactUpdate {
	pocu.mutation.SetPointOfContactEmail(s)
	return pocu
}

// SetPointOfContactPhone sets the "point_of_contact_phone" field.
func (pocu *PointOfContactUpdate) SetPointOfContactPhone(s string) *PointOfContactUpdate {
	pocu.mutation.SetPointOfContactPhone(s)
	return pocu
}

// SetPointOfContactPhoneExt sets the "point_of_contact_phone_ext" field.
func (pocu *PointOfContactUpdate) SetPointOfContactPhoneExt(s string) *PointOfContactUpdate {
	pocu.mutation.SetPointOfContactPhoneExt(s)
	return pocu
}

// SetParentID sets the "parent" edge to the MoreInfoModule entity by ID.
func (pocu *PointOfContactUpdate) SetParentID(id int) *PointOfContactUpdate {
	pocu.mutation.SetParentID(id)
	return pocu
}

// SetParent sets the "parent" edge to the MoreInfoModule entity.
func (pocu *PointOfContactUpdate) SetParent(m *MoreInfoModule) *PointOfContactUpdate {
	return pocu.SetParentID(m.ID)
}

// Mutation returns the PointOfContactMutation object of the builder.
func (pocu *PointOfContactUpdate) Mutation() *PointOfContactMutation {
	return pocu.mutation
}

// ClearParent clears the "parent" edge to the MoreInfoModule entity.
func (pocu *PointOfContactUpdate) ClearParent() *PointOfContactUpdate {
	pocu.mutation.ClearParent()
	return pocu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pocu *PointOfContactUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pocu.hooks) == 0 {
		if err = pocu.check(); err != nil {
			return 0, err
		}
		affected, err = pocu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PointOfContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pocu.check(); err != nil {
				return 0, err
			}
			pocu.mutation = mutation
			affected, err = pocu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pocu.hooks) - 1; i >= 0; i-- {
			if pocu.hooks[i] == nil {
				return 0, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pocu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pocu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pocu *PointOfContactUpdate) SaveX(ctx context.Context) int {
	affected, err := pocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pocu *PointOfContactUpdate) Exec(ctx context.Context) error {
	_, err := pocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocu *PointOfContactUpdate) ExecX(ctx context.Context) {
	if err := pocu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pocu *PointOfContactUpdate) check() error {
	if _, ok := pocu.mutation.ParentID(); pocu.mutation.ParentCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "PointOfContact.parent"`)
	}
	return nil
}

func (pocu *PointOfContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pointofcontact.Table,
			Columns: pointofcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pointofcontact.FieldID,
			},
		},
	}
	if ps := pocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pocu.mutation.PointOfContactTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactTitle,
		})
	}
	if value, ok := pocu.mutation.PointOfContactOrganization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactOrganization,
		})
	}
	if value, ok := pocu.mutation.PointOfContactEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactEmail,
		})
	}
	if value, ok := pocu.mutation.PointOfContactPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactPhone,
		})
	}
	if value, ok := pocu.mutation.PointOfContactPhoneExt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactPhoneExt,
		})
	}
	if pocu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pointofcontact.ParentTable,
			Columns: []string{pointofcontact.ParentColumn},
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
	if nodes := pocu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pointofcontact.ParentTable,
			Columns: []string{pointofcontact.ParentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pointofcontact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PointOfContactUpdateOne is the builder for updating a single PointOfContact entity.
type PointOfContactUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PointOfContactMutation
}

// SetPointOfContactTitle sets the "point_of_contact_title" field.
func (pocuo *PointOfContactUpdateOne) SetPointOfContactTitle(s string) *PointOfContactUpdateOne {
	pocuo.mutation.SetPointOfContactTitle(s)
	return pocuo
}

// SetPointOfContactOrganization sets the "point_of_contact_organization" field.
func (pocuo *PointOfContactUpdateOne) SetPointOfContactOrganization(s string) *PointOfContactUpdateOne {
	pocuo.mutation.SetPointOfContactOrganization(s)
	return pocuo
}

// SetPointOfContactEmail sets the "point_of_contact_email" field.
func (pocuo *PointOfContactUpdateOne) SetPointOfContactEmail(s string) *PointOfContactUpdateOne {
	pocuo.mutation.SetPointOfContactEmail(s)
	return pocuo
}

// SetPointOfContactPhone sets the "point_of_contact_phone" field.
func (pocuo *PointOfContactUpdateOne) SetPointOfContactPhone(s string) *PointOfContactUpdateOne {
	pocuo.mutation.SetPointOfContactPhone(s)
	return pocuo
}

// SetPointOfContactPhoneExt sets the "point_of_contact_phone_ext" field.
func (pocuo *PointOfContactUpdateOne) SetPointOfContactPhoneExt(s string) *PointOfContactUpdateOne {
	pocuo.mutation.SetPointOfContactPhoneExt(s)
	return pocuo
}

// SetParentID sets the "parent" edge to the MoreInfoModule entity by ID.
func (pocuo *PointOfContactUpdateOne) SetParentID(id int) *PointOfContactUpdateOne {
	pocuo.mutation.SetParentID(id)
	return pocuo
}

// SetParent sets the "parent" edge to the MoreInfoModule entity.
func (pocuo *PointOfContactUpdateOne) SetParent(m *MoreInfoModule) *PointOfContactUpdateOne {
	return pocuo.SetParentID(m.ID)
}

// Mutation returns the PointOfContactMutation object of the builder.
func (pocuo *PointOfContactUpdateOne) Mutation() *PointOfContactMutation {
	return pocuo.mutation
}

// ClearParent clears the "parent" edge to the MoreInfoModule entity.
func (pocuo *PointOfContactUpdateOne) ClearParent() *PointOfContactUpdateOne {
	pocuo.mutation.ClearParent()
	return pocuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pocuo *PointOfContactUpdateOne) Select(field string, fields ...string) *PointOfContactUpdateOne {
	pocuo.fields = append([]string{field}, fields...)
	return pocuo
}

// Save executes the query and returns the updated PointOfContact entity.
func (pocuo *PointOfContactUpdateOne) Save(ctx context.Context) (*PointOfContact, error) {
	var (
		err  error
		node *PointOfContact
	)
	if len(pocuo.hooks) == 0 {
		if err = pocuo.check(); err != nil {
			return nil, err
		}
		node, err = pocuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PointOfContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pocuo.check(); err != nil {
				return nil, err
			}
			pocuo.mutation = mutation
			node, err = pocuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pocuo.hooks) - 1; i >= 0; i-- {
			if pocuo.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = pocuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pocuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pocuo *PointOfContactUpdateOne) SaveX(ctx context.Context) *PointOfContact {
	node, err := pocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pocuo *PointOfContactUpdateOne) Exec(ctx context.Context) error {
	_, err := pocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocuo *PointOfContactUpdateOne) ExecX(ctx context.Context) {
	if err := pocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pocuo *PointOfContactUpdateOne) check() error {
	if _, ok := pocuo.mutation.ParentID(); pocuo.mutation.ParentCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "PointOfContact.parent"`)
	}
	return nil
}

func (pocuo *PointOfContactUpdateOne) sqlSave(ctx context.Context) (_node *PointOfContact, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pointofcontact.Table,
			Columns: pointofcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pointofcontact.FieldID,
			},
		},
	}
	id, ok := pocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "PointOfContact.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pointofcontact.FieldID)
		for _, f := range fields {
			if !pointofcontact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != pointofcontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pocuo.mutation.PointOfContactTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactTitle,
		})
	}
	if value, ok := pocuo.mutation.PointOfContactOrganization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactOrganization,
		})
	}
	if value, ok := pocuo.mutation.PointOfContactEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactEmail,
		})
	}
	if value, ok := pocuo.mutation.PointOfContactPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactPhone,
		})
	}
	if value, ok := pocuo.mutation.PointOfContactPhoneExt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pointofcontact.FieldPointOfContactPhoneExt,
		})
	}
	if pocuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pointofcontact.ParentTable,
			Columns: []string{pointofcontact.ParentColumn},
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
	if nodes := pocuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pointofcontact.ParentTable,
			Columns: []string{pointofcontact.ParentColumn},
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
	_node = &PointOfContact{config: pocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pointofcontact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
