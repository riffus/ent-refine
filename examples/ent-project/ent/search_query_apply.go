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
	"github.com/entkit/entkit/examples/ent-project/ent/company"
	"github.com/entkit/entkit/examples/ent-project/ent/country"
	"github.com/entkit/entkit/examples/ent-project/ent/email"
	"github.com/entkit/entkit/examples/ent-project/ent/image"
	"github.com/entkit/entkit/examples/ent-project/ent/location"
	"github.com/entkit/entkit/examples/ent-project/ent/phone"
	"github.com/entkit/entkit/examples/ent-project/ent/predicate"
	"github.com/entkit/entkit/examples/ent-project/ent/product"
	"github.com/entkit/entkit/examples/ent-project/ent/vendor"
	"github.com/entkit/entkit/examples/ent-project/ent/warehouse"
	"github.com/entkit/entkit/examples/ent-project/ent/website"
	"github.com/google/uuid"
)

func (cwi *CompanyWhereInput) ApplySearchQuery(q *string) *CompanyWhereInput {
	if cwi == nil {
		cwi = &CompanyWhereInput{}
	}

	if q == nil {
		return cwi
	}

	var orPredicates []predicate.Company
	orPredicates = append(orPredicates, company.NameContains(*q))
	orPredicates = append(orPredicates, company.DescriptionContains(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, company.IDEQ(u))
	}
	cwi.AddPredicates(company.Or(orPredicates...))
	return cwi
}

func (cwi *CountryWhereInput) ApplySearchQuery(q *string) *CountryWhereInput {
	if cwi == nil {
		cwi = &CountryWhereInput{}
	}

	if q == nil {
		return cwi
	}

	var orPredicates []predicate.Country
	orPredicates = append(orPredicates, country.NameContains(*q))
	orPredicates = append(orPredicates, country.CodeEQ(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, country.IDEQ(u))
	}
	cwi.AddPredicates(country.Or(orPredicates...))
	return cwi
}

func (ewi *EmailWhereInput) ApplySearchQuery(q *string) *EmailWhereInput {
	if ewi == nil {
		ewi = &EmailWhereInput{}
	}

	if q == nil {
		return ewi
	}

	var orPredicates []predicate.Email
	orPredicates = append(orPredicates, email.TitleContains(*q))
	orPredicates = append(orPredicates, email.DescriptionContains(*q))
	orPredicates = append(orPredicates, email.AddressEQ(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, email.IDEQ(u))
	}
	ewi.AddPredicates(email.Or(orPredicates...))
	return ewi
}

func (iwi *ImageWhereInput) ApplySearchQuery(q *string) *ImageWhereInput {
	if iwi == nil {
		iwi = &ImageWhereInput{}
	}

	if q == nil {
		return iwi
	}

	var orPredicates []predicate.Image
	orPredicates = append(orPredicates, image.TitleContains(*q))
	orPredicates = append(orPredicates, image.OriginalURLEQ(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, image.IDEQ(u))
	}
	iwi.AddPredicates(image.Or(orPredicates...))
	return iwi
}

func (lwi *LocationWhereInput) ApplySearchQuery(q *string) *LocationWhereInput {
	if lwi == nil {
		lwi = &LocationWhereInput{}
	}

	if q == nil {
		return lwi
	}

	var orPredicates []predicate.Location
	orPredicates = append(orPredicates, location.TitleContains(*q))
	orPredicates = append(orPredicates, location.DescriptionContains(*q))
	orPredicates = append(orPredicates, location.AddressContains(*q))
	orPredicates = append(orPredicates, location.PostcodeEQ(*q))
	orPredicates = append(orPredicates, location.TypeEQ(*q))
	orPredicates = append(orPredicates, location.StateContains(*q))
	orPredicates = append(orPredicates, location.SuburbContains(*q))
	orPredicates = append(orPredicates, location.StreetTypeEQ(*q))
	orPredicates = append(orPredicates, location.StreetNameContains(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, location.IDEQ(u))
	}
	lwi.AddPredicates(location.Or(orPredicates...))
	return lwi
}

func (pwi *PhoneWhereInput) ApplySearchQuery(q *string) *PhoneWhereInput {
	if pwi == nil {
		pwi = &PhoneWhereInput{}
	}

	if q == nil {
		return pwi
	}

	var orPredicates []predicate.Phone
	orPredicates = append(orPredicates, phone.TitleContains(*q))
	orPredicates = append(orPredicates, phone.DescriptionContains(*q))
	orPredicates = append(orPredicates, phone.NumberEQ(*q))
	orPredicates = append(orPredicates, phone.TypeEQ(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, phone.IDEQ(u))
	}
	pwi.AddPredicates(phone.Or(orPredicates...))
	return pwi
}

func (pwi *ProductWhereInput) ApplySearchQuery(q *string) *ProductWhereInput {
	if pwi == nil {
		pwi = &ProductWhereInput{}
	}

	if q == nil {
		return pwi
	}

	var orPredicates []predicate.Product
	orPredicates = append(orPredicates, product.NameContains(*q))
	orPredicates = append(orPredicates, product.DescriptionContains(*q))
	orPredicates = append(orPredicates, product.ImageEQ(*q))
	orPredicates = append(orPredicates, product.URLContains(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, product.IDEQ(u))
	}
	pwi.AddPredicates(product.Or(orPredicates...))
	return pwi
}

func (vwi *VendorWhereInput) ApplySearchQuery(q *string) *VendorWhereInput {
	if vwi == nil {
		vwi = &VendorWhereInput{}
	}

	if q == nil {
		return vwi
	}

	var orPredicates []predicate.Vendor
	orPredicates = append(orPredicates, vendor.NameContains(*q))
	orPredicates = append(orPredicates, vendor.SchemaContains(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, vendor.IDEQ(u))
	}
	vwi.AddPredicates(vendor.Or(orPredicates...))
	return vwi
}

func (wwi *WarehouseWhereInput) ApplySearchQuery(q *string) *WarehouseWhereInput {
	if wwi == nil {
		wwi = &WarehouseWhereInput{}
	}

	if q == nil {
		return wwi
	}

	var orPredicates []predicate.Warehouse
	orPredicates = append(orPredicates, warehouse.NameContains(*q))
	orPredicates = append(orPredicates, warehouse.OriginalDataContains(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, warehouse.IDEQ(u))
	}
	wwi.AddPredicates(warehouse.Or(orPredicates...))
	return wwi
}

func (wwi *WebsiteWhereInput) ApplySearchQuery(q *string) *WebsiteWhereInput {
	if wwi == nil {
		wwi = &WebsiteWhereInput{}
	}

	if q == nil {
		return wwi
	}

	var orPredicates []predicate.Website
	orPredicates = append(orPredicates, website.TitleContains(*q))
	orPredicates = append(orPredicates, website.DescriptionContains(*q))
	orPredicates = append(orPredicates, website.URLContains(*q))
	// id uuid Field
	u, err := uuid.Parse(*q)
	if err == nil {
		orPredicates = append(orPredicates, website.IDEQ(u))
	}
	wwi.AddPredicates(website.Or(orPredicates...))
	return wwi
}
