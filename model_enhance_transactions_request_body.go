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

// EnhanceTransactionsRequestBody struct for EnhanceTransactionsRequestBody
type EnhanceTransactionsRequestBody struct {
	Transactions []EnhanceTransactionsRequest `json:"transactions,omitempty"`
}

// NewEnhanceTransactionsRequestBody instantiates a new EnhanceTransactionsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnhanceTransactionsRequestBody() *EnhanceTransactionsRequestBody {
	this := EnhanceTransactionsRequestBody{}
	return &this
}

// NewEnhanceTransactionsRequestBodyWithDefaults instantiates a new EnhanceTransactionsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnhanceTransactionsRequestBodyWithDefaults() *EnhanceTransactionsRequestBody {
	this := EnhanceTransactionsRequestBody{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *EnhanceTransactionsRequestBody) GetTransactions() []EnhanceTransactionsRequest {
	if o == nil || o.Transactions == nil {
		var ret []EnhanceTransactionsRequest
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhanceTransactionsRequestBody) GetTransactionsOk() ([]EnhanceTransactionsRequest, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *EnhanceTransactionsRequestBody) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []EnhanceTransactionsRequest and assigns it to the Transactions field.
func (o *EnhanceTransactionsRequestBody) SetTransactions(v []EnhanceTransactionsRequest) {
	o.Transactions = v
}

func (o EnhanceTransactionsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableEnhanceTransactionsRequestBody struct {
	value *EnhanceTransactionsRequestBody
	isSet bool
}

func (v NullableEnhanceTransactionsRequestBody) Get() *EnhanceTransactionsRequestBody {
	return v.value
}

func (v *NullableEnhanceTransactionsRequestBody) Set(val *EnhanceTransactionsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEnhanceTransactionsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEnhanceTransactionsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnhanceTransactionsRequestBody(val *EnhanceTransactionsRequestBody) *NullableEnhanceTransactionsRequestBody {
	return &NullableEnhanceTransactionsRequestBody{value: val, isSet: true}
}

func (v NullableEnhanceTransactionsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnhanceTransactionsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


