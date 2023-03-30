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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/entkit/entkit/examples/ent-project/ent/company"
	"github.com/entkit/entkit/examples/ent-project/ent/country"
	"github.com/entkit/entkit/examples/ent-project/ent/location"
	"github.com/entkit/entkit/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetTitle sets the "title" field.
func (lu *LocationUpdate) SetTitle(s string) *LocationUpdate {
	lu.mutation.SetTitle(s)
	return lu
}

// SetDescription sets the "description" field.
func (lu *LocationUpdate) SetDescription(s string) *LocationUpdate {
	lu.mutation.SetDescription(s)
	return lu
}

// SetLatitude sets the "latitude" field.
func (lu *LocationUpdate) SetLatitude(f float64) *LocationUpdate {
	lu.mutation.ResetLatitude()
	lu.mutation.SetLatitude(f)
	return lu
}

// AddLatitude adds f to the "latitude" field.
func (lu *LocationUpdate) AddLatitude(f float64) *LocationUpdate {
	lu.mutation.AddLatitude(f)
	return lu
}

// SetLongitude sets the "longitude" field.
func (lu *LocationUpdate) SetLongitude(f float64) *LocationUpdate {
	lu.mutation.ResetLongitude()
	lu.mutation.SetLongitude(f)
	return lu
}

// AddLongitude adds f to the "longitude" field.
func (lu *LocationUpdate) AddLongitude(f float64) *LocationUpdate {
	lu.mutation.AddLongitude(f)
	return lu
}

// SetAddress sets the "address" field.
func (lu *LocationUpdate) SetAddress(s string) *LocationUpdate {
	lu.mutation.SetAddress(s)
	return lu
}

// SetPostcode sets the "postcode" field.
func (lu *LocationUpdate) SetPostcode(s string) *LocationUpdate {
	lu.mutation.SetPostcode(s)
	return lu
}

// SetType sets the "type" field.
func (lu *LocationUpdate) SetType(s string) *LocationUpdate {
	lu.mutation.SetType(s)
	return lu
}

// SetState sets the "state" field.
func (lu *LocationUpdate) SetState(s string) *LocationUpdate {
	lu.mutation.SetState(s)
	return lu
}

// SetSuburb sets the "suburb" field.
func (lu *LocationUpdate) SetSuburb(s string) *LocationUpdate {
	lu.mutation.SetSuburb(s)
	return lu
}

// SetStreetType sets the "street_type" field.
func (lu *LocationUpdate) SetStreetType(s string) *LocationUpdate {
	lu.mutation.SetStreetType(s)
	return lu
}

// SetStreetName sets the "street_name" field.
func (lu *LocationUpdate) SetStreetName(s string) *LocationUpdate {
	lu.mutation.SetStreetName(s)
	return lu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (lu *LocationUpdate) SetCompanyID(id uuid.UUID) *LocationUpdate {
	lu.mutation.SetCompanyID(id)
	return lu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (lu *LocationUpdate) SetNillableCompanyID(id *uuid.UUID) *LocationUpdate {
	if id != nil {
		lu = lu.SetCompanyID(*id)
	}
	return lu
}

// SetCompany sets the "company" edge to the Company entity.
func (lu *LocationUpdate) SetCompany(c *Company) *LocationUpdate {
	return lu.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (lu *LocationUpdate) SetCountryID(id uuid.UUID) *LocationUpdate {
	lu.mutation.SetCountryID(id)
	return lu
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (lu *LocationUpdate) SetNillableCountryID(id *uuid.UUID) *LocationUpdate {
	if id != nil {
		lu = lu.SetCountryID(*id)
	}
	return lu
}

// SetCountry sets the "country" edge to the Country entity.
func (lu *LocationUpdate) SetCountry(c *Country) *LocationUpdate {
	return lu.SetCountryID(c.ID)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (lu *LocationUpdate) ClearCompany() *LocationUpdate {
	lu.mutation.ClearCompany()
	return lu
}

// ClearCountry clears the "country" edge to the Country entity.
func (lu *LocationUpdate) ClearCountry() *LocationUpdate {
	lu.mutation.ClearCountry()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LocationMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LocationUpdate) check() error {
	if v, ok := lu.mutation.Title(); ok {
		if err := location.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Location.title": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Description(); ok {
		if err := location.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Location.description": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Address(); ok {
		if err := location.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Location.address": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Postcode(); ok {
		if err := location.PostcodeValidator(v); err != nil {
			return &ValidationError{Name: "postcode", err: fmt.Errorf(`ent: validator failed for field "Location.postcode": %w`, err)}
		}
	}
	if v, ok := lu.mutation.GetType(); ok {
		if err := location.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Location.type": %w`, err)}
		}
	}
	if v, ok := lu.mutation.State(); ok {
		if err := location.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Location.state": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Suburb(); ok {
		if err := location.SuburbValidator(v); err != nil {
			return &ValidationError{Name: "suburb", err: fmt.Errorf(`ent: validator failed for field "Location.suburb": %w`, err)}
		}
	}
	if v, ok := lu.mutation.StreetType(); ok {
		if err := location.StreetTypeValidator(v); err != nil {
			return &ValidationError{Name: "street_type", err: fmt.Errorf(`ent: validator failed for field "Location.street_type": %w`, err)}
		}
	}
	if v, ok := lu.mutation.StreetName(); ok {
		if err := location.StreetNameValidator(v); err != nil {
			return &ValidationError{Name: "street_name", err: fmt.Errorf(`ent: validator failed for field "Location.street_name": %w`, err)}
		}
	}
	return nil
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Title(); ok {
		_spec.SetField(location.FieldTitle, field.TypeString, value)
	}
	if value, ok := lu.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
	}
	if value, ok := lu.mutation.Latitude(); ok {
		_spec.SetField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.AddedLatitude(); ok {
		_spec.AddField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.Longitude(); ok {
		_spec.SetField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.AddedLongitude(); ok {
		_spec.AddField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.Address(); ok {
		_spec.SetField(location.FieldAddress, field.TypeString, value)
	}
	if value, ok := lu.mutation.Postcode(); ok {
		_spec.SetField(location.FieldPostcode, field.TypeString, value)
	}
	if value, ok := lu.mutation.GetType(); ok {
		_spec.SetField(location.FieldType, field.TypeString, value)
	}
	if value, ok := lu.mutation.State(); ok {
		_spec.SetField(location.FieldState, field.TypeString, value)
	}
	if value, ok := lu.mutation.Suburb(); ok {
		_spec.SetField(location.FieldSuburb, field.TypeString, value)
	}
	if value, ok := lu.mutation.StreetType(); ok {
		_spec.SetField(location.FieldStreetType, field.TypeString, value)
	}
	if value, ok := lu.mutation.StreetName(); ok {
		_spec.SetField(location.FieldStreetName, field.TypeString, value)
	}
	if lu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CompanyTable,
			Columns: []string{location.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CompanyTable,
			Columns: []string{location.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CountryTable,
			Columns: []string{location.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CountryTable,
			Columns: []string{location.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetTitle sets the "title" field.
func (luo *LocationUpdateOne) SetTitle(s string) *LocationUpdateOne {
	luo.mutation.SetTitle(s)
	return luo
}

// SetDescription sets the "description" field.
func (luo *LocationUpdateOne) SetDescription(s string) *LocationUpdateOne {
	luo.mutation.SetDescription(s)
	return luo
}

// SetLatitude sets the "latitude" field.
func (luo *LocationUpdateOne) SetLatitude(f float64) *LocationUpdateOne {
	luo.mutation.ResetLatitude()
	luo.mutation.SetLatitude(f)
	return luo
}

// AddLatitude adds f to the "latitude" field.
func (luo *LocationUpdateOne) AddLatitude(f float64) *LocationUpdateOne {
	luo.mutation.AddLatitude(f)
	return luo
}

// SetLongitude sets the "longitude" field.
func (luo *LocationUpdateOne) SetLongitude(f float64) *LocationUpdateOne {
	luo.mutation.ResetLongitude()
	luo.mutation.SetLongitude(f)
	return luo
}

// AddLongitude adds f to the "longitude" field.
func (luo *LocationUpdateOne) AddLongitude(f float64) *LocationUpdateOne {
	luo.mutation.AddLongitude(f)
	return luo
}

// SetAddress sets the "address" field.
func (luo *LocationUpdateOne) SetAddress(s string) *LocationUpdateOne {
	luo.mutation.SetAddress(s)
	return luo
}

// SetPostcode sets the "postcode" field.
func (luo *LocationUpdateOne) SetPostcode(s string) *LocationUpdateOne {
	luo.mutation.SetPostcode(s)
	return luo
}

// SetType sets the "type" field.
func (luo *LocationUpdateOne) SetType(s string) *LocationUpdateOne {
	luo.mutation.SetType(s)
	return luo
}

// SetState sets the "state" field.
func (luo *LocationUpdateOne) SetState(s string) *LocationUpdateOne {
	luo.mutation.SetState(s)
	return luo
}

// SetSuburb sets the "suburb" field.
func (luo *LocationUpdateOne) SetSuburb(s string) *LocationUpdateOne {
	luo.mutation.SetSuburb(s)
	return luo
}

// SetStreetType sets the "street_type" field.
func (luo *LocationUpdateOne) SetStreetType(s string) *LocationUpdateOne {
	luo.mutation.SetStreetType(s)
	return luo
}

// SetStreetName sets the "street_name" field.
func (luo *LocationUpdateOne) SetStreetName(s string) *LocationUpdateOne {
	luo.mutation.SetStreetName(s)
	return luo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (luo *LocationUpdateOne) SetCompanyID(id uuid.UUID) *LocationUpdateOne {
	luo.mutation.SetCompanyID(id)
	return luo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableCompanyID(id *uuid.UUID) *LocationUpdateOne {
	if id != nil {
		luo = luo.SetCompanyID(*id)
	}
	return luo
}

// SetCompany sets the "company" edge to the Company entity.
func (luo *LocationUpdateOne) SetCompany(c *Company) *LocationUpdateOne {
	return luo.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (luo *LocationUpdateOne) SetCountryID(id uuid.UUID) *LocationUpdateOne {
	luo.mutation.SetCountryID(id)
	return luo
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableCountryID(id *uuid.UUID) *LocationUpdateOne {
	if id != nil {
		luo = luo.SetCountryID(*id)
	}
	return luo
}

// SetCountry sets the "country" edge to the Country entity.
func (luo *LocationUpdateOne) SetCountry(c *Country) *LocationUpdateOne {
	return luo.SetCountryID(c.ID)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (luo *LocationUpdateOne) ClearCompany() *LocationUpdateOne {
	luo.mutation.ClearCompany()
	return luo
}

// ClearCountry clears the "country" edge to the Country entity.
func (luo *LocationUpdateOne) ClearCountry() *LocationUpdateOne {
	luo.mutation.ClearCountry()
	return luo
}

// Where appends a list predicates to the LocationUpdate builder.
func (luo *LocationUpdateOne) Where(ps ...predicate.Location) *LocationUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	return withHooks[*Location, LocationMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LocationUpdateOne) check() error {
	if v, ok := luo.mutation.Title(); ok {
		if err := location.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Location.title": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Description(); ok {
		if err := location.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Location.description": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Address(); ok {
		if err := location.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Location.address": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Postcode(); ok {
		if err := location.PostcodeValidator(v); err != nil {
			return &ValidationError{Name: "postcode", err: fmt.Errorf(`ent: validator failed for field "Location.postcode": %w`, err)}
		}
	}
	if v, ok := luo.mutation.GetType(); ok {
		if err := location.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Location.type": %w`, err)}
		}
	}
	if v, ok := luo.mutation.State(); ok {
		if err := location.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Location.state": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Suburb(); ok {
		if err := location.SuburbValidator(v); err != nil {
			return &ValidationError{Name: "suburb", err: fmt.Errorf(`ent: validator failed for field "Location.suburb": %w`, err)}
		}
	}
	if v, ok := luo.mutation.StreetType(); ok {
		if err := location.StreetTypeValidator(v); err != nil {
			return &ValidationError{Name: "street_type", err: fmt.Errorf(`ent: validator failed for field "Location.street_type": %w`, err)}
		}
	}
	if v, ok := luo.mutation.StreetName(); ok {
		if err := location.StreetNameValidator(v); err != nil {
			return &ValidationError{Name: "street_name", err: fmt.Errorf(`ent: validator failed for field "Location.street_name": %w`, err)}
		}
	}
	return nil
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Title(); ok {
		_spec.SetField(location.FieldTitle, field.TypeString, value)
	}
	if value, ok := luo.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
	}
	if value, ok := luo.mutation.Latitude(); ok {
		_spec.SetField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.AddedLatitude(); ok {
		_spec.AddField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.Longitude(); ok {
		_spec.SetField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.AddedLongitude(); ok {
		_spec.AddField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.Address(); ok {
		_spec.SetField(location.FieldAddress, field.TypeString, value)
	}
	if value, ok := luo.mutation.Postcode(); ok {
		_spec.SetField(location.FieldPostcode, field.TypeString, value)
	}
	if value, ok := luo.mutation.GetType(); ok {
		_spec.SetField(location.FieldType, field.TypeString, value)
	}
	if value, ok := luo.mutation.State(); ok {
		_spec.SetField(location.FieldState, field.TypeString, value)
	}
	if value, ok := luo.mutation.Suburb(); ok {
		_spec.SetField(location.FieldSuburb, field.TypeString, value)
	}
	if value, ok := luo.mutation.StreetType(); ok {
		_spec.SetField(location.FieldStreetType, field.TypeString, value)
	}
	if value, ok := luo.mutation.StreetName(); ok {
		_spec.SetField(location.FieldStreetName, field.TypeString, value)
	}
	if luo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CompanyTable,
			Columns: []string{location.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CompanyTable,
			Columns: []string{location.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CountryTable,
			Columns: []string{location.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CountryTable,
			Columns: []string{location.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
