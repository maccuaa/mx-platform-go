/*
MX Platform API

The MX Platform API is a powerful, fully-featured API designed to make aggregating and enhancing financial data easy and reliable. It can seamlessly connect your app or website to tens of thousands of financial institutions.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mxplatformgo

import (
	"encoding/json"
)

// SpendingPlansResponseBody struct for SpendingPlansResponseBody
type SpendingPlansResponseBody struct {
	IterationItems []SpendingPlanResponse `json:"iteration_items,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewSpendingPlansResponseBody instantiates a new SpendingPlansResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingPlansResponseBody() *SpendingPlansResponseBody {
	this := SpendingPlansResponseBody{}
	return &this
}

// NewSpendingPlansResponseBodyWithDefaults instantiates a new SpendingPlansResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingPlansResponseBodyWithDefaults() *SpendingPlansResponseBody {
	this := SpendingPlansResponseBody{}
	return &this
}

// GetIterationItems returns the IterationItems field value if set, zero value otherwise.
func (o *SpendingPlansResponseBody) GetIterationItems() []SpendingPlanResponse {
	if o == nil || o.IterationItems == nil {
		var ret []SpendingPlanResponse
		return ret
	}
	return o.IterationItems
}

// GetIterationItemsOk returns a tuple with the IterationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlansResponseBody) GetIterationItemsOk() ([]SpendingPlanResponse, bool) {
	if o == nil || o.IterationItems == nil {
		return nil, false
	}
	return o.IterationItems, true
}

// HasIterationItems returns a boolean if a field has been set.
func (o *SpendingPlansResponseBody) HasIterationItems() bool {
	if o != nil && o.IterationItems != nil {
		return true
	}

	return false
}

// SetIterationItems gets a reference to the given []SpendingPlanResponse and assigns it to the IterationItems field.
func (o *SpendingPlansResponseBody) SetIterationItems(v []SpendingPlanResponse) {
	o.IterationItems = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SpendingPlansResponseBody) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlansResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SpendingPlansResponseBody) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *SpendingPlansResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o SpendingPlansResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IterationItems != nil {
		toSerialize["iteration_items"] = o.IterationItems
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableSpendingPlansResponseBody struct {
	value *SpendingPlansResponseBody
	isSet bool
}

func (v NullableSpendingPlansResponseBody) Get() *SpendingPlansResponseBody {
	return v.value
}

func (v *NullableSpendingPlansResponseBody) Set(val *SpendingPlansResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingPlansResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingPlansResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingPlansResponseBody(val *SpendingPlansResponseBody) *NullableSpendingPlansResponseBody {
	return &NullableSpendingPlansResponseBody{value: val, isSet: true}
}

func (v NullableSpendingPlansResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingPlansResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


