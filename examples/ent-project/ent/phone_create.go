// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/country"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/phone"
	"github.com/google/uuid"
)

// PhoneCreate is the builder for creating a Phone entity.
type PhoneCreate struct {
	config
	mutation *PhoneMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (pc *PhoneCreate) SetTitle(s string) *PhoneCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PhoneCreate) SetDescription(s string) *PhoneCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNumber sets the "number" field.
func (pc *PhoneCreate) SetNumber(s string) *PhoneCreate {
	pc.mutation.SetNumber(s)
	return pc
}

// SetType sets the "type" field.
func (pc *PhoneCreate) SetType(s string) *PhoneCreate {
	pc.mutation.SetType(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PhoneCreate) SetID(u uuid.UUID) *PhoneCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PhoneCreate) SetNillableID(u *uuid.UUID) *PhoneCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pc *PhoneCreate) SetCompanyID(id uuid.UUID) *PhoneCreate {
	pc.mutation.SetCompanyID(id)
	return pc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pc *PhoneCreate) SetNillableCompanyID(id *uuid.UUID) *PhoneCreate {
	if id != nil {
		pc = pc.SetCompanyID(*id)
	}
	return pc
}

// SetCompany sets the "company" edge to the Company entity.
func (pc *PhoneCreate) SetCompany(c *Company) *PhoneCreate {
	return pc.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (pc *PhoneCreate) SetCountryID(id uuid.UUID) *PhoneCreate {
	pc.mutation.SetCountryID(id)
	return pc
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (pc *PhoneCreate) SetNillableCountryID(id *uuid.UUID) *PhoneCreate {
	if id != nil {
		pc = pc.SetCountryID(*id)
	}
	return pc
}

// SetCountry sets the "country" edge to the Country entity.
func (pc *PhoneCreate) SetCountry(c *Country) *PhoneCreate {
	return pc.SetCountryID(c.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (pc *PhoneCreate) Mutation() *PhoneMutation {
	return pc.mutation
}

// Save creates the Phone in the database.
func (pc *PhoneCreate) Save(ctx context.Context) (*Phone, error) {
	var (
		err  error
		node *Phone
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Phone)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PhoneMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PhoneCreate) SaveX(ctx context.Context) *Phone {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PhoneCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PhoneCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PhoneCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := phone.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PhoneCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Phone.title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := phone.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Phone.title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Phone.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := phone.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Phone.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Phone.number"`)}
	}
	if v, ok := pc.mutation.Number(); ok {
		if err := phone.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Phone.number": %w`, err)}
		}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Phone.type"`)}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := phone.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Phone.type": %w`, err)}
		}
	}
	return nil
}

func (pc *PhoneCreate) sqlSave(ctx context.Context) (*Phone, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *PhoneCreate) createSpec() (*Phone, *sqlgraph.CreateSpec) {
	var (
		_node = &Phone{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: phone.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: phone.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(phone.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(phone.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Number(); ok {
		_spec.SetField(phone.FieldNumber, field.TypeString, value)
		_node.Number = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(phone.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := pc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CompanyTable,
			Columns: []string{phone.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_phones = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phone.CountryTable,
			Columns: []string{phone.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.country_phones = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PhoneCreateBulk is the builder for creating many Phone entities in bulk.
type PhoneCreateBulk struct {
	config
	builders []*PhoneCreate
}

// Save creates the Phone entities in the database.
func (pcb *PhoneCreateBulk) Save(ctx context.Context) ([]*Phone, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Phone, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PhoneMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PhoneCreateBulk) SaveX(ctx context.Context) []*Phone {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PhoneCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PhoneCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
