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

// SpendingPlanIterationsResponse struct for SpendingPlanIterationsResponse
type SpendingPlanIterationsResponse struct {
	Iterations []SpendingPlanIterationResponse `json:"iterations,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewSpendingPlanIterationsResponse instantiates a new SpendingPlanIterationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingPlanIterationsResponse() *SpendingPlanIterationsResponse {
	this := SpendingPlanIterationsResponse{}
	return &this
}

// NewSpendingPlanIterationsResponseWithDefaults instantiates a new SpendingPlanIterationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingPlanIterationsResponseWithDefaults() *SpendingPlanIterationsResponse {
	this := SpendingPlanIterationsResponse{}
	return &this
}

// GetIterations returns the Iterations field value if set, zero value otherwise.
func (o *SpendingPlanIterationsResponse) GetIterations() []SpendingPlanIterationResponse {
	if o == nil || o.Iterations == nil {
		var ret []SpendingPlanIterationResponse
		return ret
	}
	return o.Iterations
}

// GetIterationsOk returns a tuple with the Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlanIterationsResponse) GetIterationsOk() ([]SpendingPlanIterationResponse, bool) {
	if o == nil || o.Iterations == nil {
		return nil, false
	}
	return o.Iterations, true
}

// HasIterations returns a boolean if a field has been set.
func (o *SpendingPlanIterationsResponse) HasIterations() bool {
	if o != nil && o.Iterations != nil {
		return true
	}

	return false
}

// SetIterations gets a reference to the given []SpendingPlanIterationResponse and assigns it to the Iterations field.
func (o *SpendingPlanIterationsResponse) SetIterations(v []SpendingPlanIterationResponse) {
	o.Iterations = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SpendingPlanIterationsResponse) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlanIterationsResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SpendingPlanIterationsResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *SpendingPlanIterationsResponse) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o SpendingPlanIterationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iterations != nil {
		toSerialize["iterations"] = o.Iterations
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableSpendingPlanIterationsResponse struct {
	value *SpendingPlanIterationsResponse
	isSet bool
}

func (v NullableSpendingPlanIterationsResponse) Get() *SpendingPlanIterationsResponse {
	return v.value
}

func (v *NullableSpendingPlanIterationsResponse) Set(val *SpendingPlanIterationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingPlanIterationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingPlanIterationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingPlanIterationsResponse(val *SpendingPlanIterationsResponse) *NullableSpendingPlanIterationsResponse {
	return &NullableSpendingPlanIterationsResponse{value: val, isSet: true}
}

func (v NullableSpendingPlanIterationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingPlanIterationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


