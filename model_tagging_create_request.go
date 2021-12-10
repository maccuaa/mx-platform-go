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

// TaggingCreateRequest struct for TaggingCreateRequest
type TaggingCreateRequest struct {
	TagGuid string `json:"tag_guid"`
	TransactionGuid string `json:"transaction_guid"`
}

// NewTaggingCreateRequest instantiates a new TaggingCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaggingCreateRequest(tagGuid string, transactionGuid string) *TaggingCreateRequest {
	this := TaggingCreateRequest{}
	this.TagGuid = tagGuid
	this.TransactionGuid = transactionGuid
	return &this
}

// NewTaggingCreateRequestWithDefaults instantiates a new TaggingCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggingCreateRequestWithDefaults() *TaggingCreateRequest {
	this := TaggingCreateRequest{}
	return &this
}

// GetTagGuid returns the TagGuid field value
func (o *TaggingCreateRequest) GetTagGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagGuid
}

// GetTagGuidOk returns a tuple with the TagGuid field value
// and a boolean to check if the value has been set.
func (o *TaggingCreateRequest) GetTagGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TagGuid, true
}

// SetTagGuid sets field value
func (o *TaggingCreateRequest) SetTagGuid(v string) {
	o.TagGuid = v
}

// GetTransactionGuid returns the TransactionGuid field value
func (o *TaggingCreateRequest) GetTransactionGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionGuid
}

// GetTransactionGuidOk returns a tuple with the TransactionGuid field value
// and a boolean to check if the value has been set.
func (o *TaggingCreateRequest) GetTransactionGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionGuid, true
}

// SetTransactionGuid sets field value
func (o *TaggingCreateRequest) SetTransactionGuid(v string) {
	o.TransactionGuid = v
}

func (o TaggingCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag_guid"] = o.TagGuid
	}
	if true {
		toSerialize["transaction_guid"] = o.TransactionGuid
	}
	return json.Marshal(toSerialize)
}

type NullableTaggingCreateRequest struct {
	value *TaggingCreateRequest
	isSet bool
}

func (v NullableTaggingCreateRequest) Get() *TaggingCreateRequest {
	return v.value
}

func (v *NullableTaggingCreateRequest) Set(val *TaggingCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTaggingCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTaggingCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaggingCreateRequest(val *TaggingCreateRequest) *NullableTaggingCreateRequest {
	return &NullableTaggingCreateRequest{value: val, isSet: true}
}

func (v NullableTaggingCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaggingCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


