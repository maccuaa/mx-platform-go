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

// PaymentAccountResponseBody struct for PaymentAccountResponseBody
type PaymentAccountResponseBody struct {
	PaymentAccount *PaymentAccountResponse `json:"payment_account,omitempty"`
}

// NewPaymentAccountResponseBody instantiates a new PaymentAccountResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAccountResponseBody() *PaymentAccountResponseBody {
	this := PaymentAccountResponseBody{}
	return &this
}

// NewPaymentAccountResponseBodyWithDefaults instantiates a new PaymentAccountResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAccountResponseBodyWithDefaults() *PaymentAccountResponseBody {
	this := PaymentAccountResponseBody{}
	return &this
}

// GetPaymentAccount returns the PaymentAccount field value if set, zero value otherwise.
func (o *PaymentAccountResponseBody) GetPaymentAccount() PaymentAccountResponse {
	if o == nil || o.PaymentAccount == nil {
		var ret PaymentAccountResponse
		return ret
	}
	return *o.PaymentAccount
}

// GetPaymentAccountOk returns a tuple with the PaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountResponseBody) GetPaymentAccountOk() (*PaymentAccountResponse, bool) {
	if o == nil || o.PaymentAccount == nil {
		return nil, false
	}
	return o.PaymentAccount, true
}

// HasPaymentAccount returns a boolean if a field has been set.
func (o *PaymentAccountResponseBody) HasPaymentAccount() bool {
	if o != nil && o.PaymentAccount != nil {
		return true
	}

	return false
}

// SetPaymentAccount gets a reference to the given PaymentAccountResponse and assigns it to the PaymentAccount field.
func (o *PaymentAccountResponseBody) SetPaymentAccount(v PaymentAccountResponse) {
	o.PaymentAccount = &v
}

func (o PaymentAccountResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentAccount != nil {
		toSerialize["payment_account"] = o.PaymentAccount
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentAccountResponseBody struct {
	value *PaymentAccountResponseBody
	isSet bool
}

func (v NullablePaymentAccountResponseBody) Get() *PaymentAccountResponseBody {
	return v.value
}

func (v *NullablePaymentAccountResponseBody) Set(val *PaymentAccountResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAccountResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAccountResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAccountResponseBody(val *PaymentAccountResponseBody) *NullablePaymentAccountResponseBody {
	return &NullablePaymentAccountResponseBody{value: val, isSet: true}
}

func (v NullablePaymentAccountResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAccountResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

