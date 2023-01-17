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
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/country"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/phone"
	"github.com/google/uuid"
)

// Phone is the model entity for the Phone schema.
type Phone struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhoneQuery when eager-loading is set.
	Edges          PhoneEdges `json:"edges"`
	company_phones *uuid.UUID
	country_phones *uuid.UUID
}

// PhoneEdges holds the relations/edges for other nodes in the graph.
type PhoneEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhoneEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[0] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhoneEdges) CountryOrErr() (*Country, error) {
	if e.loadedTypes[1] {
		if e.Country == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Country, nil
	}
	return nil, &NotLoadedError{edge: "country"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Phone) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case phone.FieldTitle, phone.FieldDescription, phone.FieldNumber, phone.FieldType:
			values[i] = new(sql.NullString)
		case phone.FieldID:
			values[i] = new(uuid.UUID)
		case phone.ForeignKeys[0]: // company_phones
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case phone.ForeignKeys[1]: // country_phones
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Phone", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Phone fields.
func (ph *Phone) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case phone.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ph.ID = *value
			}
		case phone.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ph.Title = value.String
			}
		case phone.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ph.Description = value.String
			}
		case phone.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				ph.Number = value.String
			}
		case phone.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ph.Type = value.String
			}
		case phone.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field company_phones", values[i])
			} else if value.Valid {
				ph.company_phones = new(uuid.UUID)
				*ph.company_phones = *value.S.(*uuid.UUID)
			}
		case phone.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field country_phones", values[i])
			} else if value.Valid {
				ph.country_phones = new(uuid.UUID)
				*ph.country_phones = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCompany queries the "company" edge of the Phone entity.
func (ph *Phone) QueryCompany() *CompanyQuery {
	return (&PhoneClient{config: ph.config}).QueryCompany(ph)
}

// QueryCountry queries the "country" edge of the Phone entity.
func (ph *Phone) QueryCountry() *CountryQuery {
	return (&PhoneClient{config: ph.config}).QueryCountry(ph)
}

// Update returns a builder for updating this Phone.
// Note that you need to call Phone.Unwrap() before calling this method if this Phone
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *Phone) Update() *PhoneUpdateOne {
	return (&PhoneClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the Phone entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *Phone) Unwrap() *Phone {
	_tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: Phone is not a transactional entity")
	}
	ph.config.driver = _tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *Phone) String() string {
	var builder strings.Builder
	builder.WriteString("Phone(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ph.ID))
	builder.WriteString("title=")
	builder.WriteString(ph.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ph.Description)
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(ph.Number)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ph.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Phones is a parsable slice of Phone.
type Phones []*Phone

func (ph Phones) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
