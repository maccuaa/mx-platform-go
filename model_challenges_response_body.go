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

// ChallengesResponseBody struct for ChallengesResponseBody
type ChallengesResponseBody struct {
	Challenges []ChallengeResponse `json:"challenges,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewChallengesResponseBody instantiates a new ChallengesResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengesResponseBody() *ChallengesResponseBody {
	this := ChallengesResponseBody{}
	return &this
}

// NewChallengesResponseBodyWithDefaults instantiates a new ChallengesResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengesResponseBodyWithDefaults() *ChallengesResponseBody {
	this := ChallengesResponseBody{}
	return &this
}

// GetChallenges returns the Challenges field value if set, zero value otherwise.
func (o *ChallengesResponseBody) GetChallenges() []ChallengeResponse {
	if o == nil || o.Challenges == nil {
		var ret []ChallengeResponse
		return ret
	}
	return o.Challenges
}

// GetChallengesOk returns a tuple with the Challenges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengesResponseBody) GetChallengesOk() ([]ChallengeResponse, bool) {
	if o == nil || o.Challenges == nil {
		return nil, false
	}
	return o.Challenges, true
}

// HasChallenges returns a boolean if a field has been set.
func (o *ChallengesResponseBody) HasChallenges() bool {
	if o != nil && o.Challenges != nil {
		return true
	}

	return false
}

// SetChallenges gets a reference to the given []ChallengeResponse and assigns it to the Challenges field.
func (o *ChallengesResponseBody) SetChallenges(v []ChallengeResponse) {
	o.Challenges = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ChallengesResponseBody) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengesResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ChallengesResponseBody) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *ChallengesResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o ChallengesResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Challenges != nil {
		toSerialize["challenges"] = o.Challenges
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableChallengesResponseBody struct {
	value *ChallengesResponseBody
	isSet bool
}

func (v NullableChallengesResponseBody) Get() *ChallengesResponseBody {
	return v.value
}

func (v *NullableChallengesResponseBody) Set(val *ChallengesResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengesResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengesResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengesResponseBody(val *ChallengesResponseBody) *NullableChallengesResponseBody {
	return &NullableChallengesResponseBody{value: val, isSet: true}
}

func (v NullableChallengesResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengesResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


