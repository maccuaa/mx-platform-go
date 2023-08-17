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

// MemberCreateRequestBody struct for MemberCreateRequestBody
type MemberCreateRequestBody struct {
	ClientRedirectUrl *string `json:"client_redirect_url,omitempty"`
	EnableApp2app *bool `json:"enable_app2app,omitempty"`
	Member *MemberCreateRequest `json:"member,omitempty"`
	ReferralSource *string `json:"referral_source,omitempty"`
	UiMessageWebviewUrlScheme *string `json:"ui_message_webview_url_scheme,omitempty"`
}

// NewMemberCreateRequestBody instantiates a new MemberCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberCreateRequestBody() *MemberCreateRequestBody {
	this := MemberCreateRequestBody{}
	return &this
}

// NewMemberCreateRequestBodyWithDefaults instantiates a new MemberCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberCreateRequestBodyWithDefaults() *MemberCreateRequestBody {
	this := MemberCreateRequestBody{}
	return &this
}

// GetClientRedirectUrl returns the ClientRedirectUrl field value if set, zero value otherwise.
func (o *MemberCreateRequestBody) GetClientRedirectUrl() string {
	if o == nil || o.ClientRedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.ClientRedirectUrl
}

// GetClientRedirectUrlOk returns a tuple with the ClientRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCreateRequestBody) GetClientRedirectUrlOk() (*string, bool) {
	if o == nil || o.ClientRedirectUrl == nil {
		return nil, false
	}
	return o.ClientRedirectUrl, true
}

// HasClientRedirectUrl returns a boolean if a field has been set.
func (o *MemberCreateRequestBody) HasClientRedirectUrl() bool {
	if o != nil && o.ClientRedirectUrl != nil {
		return true
	}

	return false
}

// SetClientRedirectUrl gets a reference to the given string and assigns it to the ClientRedirectUrl field.
func (o *MemberCreateRequestBody) SetClientRedirectUrl(v string) {
	o.ClientRedirectUrl = &v
}

// GetEnableApp2app returns the EnableApp2app field value if set, zero value otherwise.
func (o *MemberCreateRequestBody) GetEnableApp2app() bool {
	if o == nil || o.EnableApp2app == nil {
		var ret bool
		return ret
	}
	return *o.EnableApp2app
}

// GetEnableApp2appOk returns a tuple with the EnableApp2app field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCreateRequestBody) GetEnableApp2appOk() (*bool, bool) {
	if o == nil || o.EnableApp2app == nil {
		return nil, false
	}
	return o.EnableApp2app, true
}

// HasEnableApp2app returns a boolean if a field has been set.
func (o *MemberCreateRequestBody) HasEnableApp2app() bool {
	if o != nil && o.EnableApp2app != nil {
		return true
	}

	return false
}

// SetEnableApp2app gets a reference to the given bool and assigns it to the EnableApp2app field.
func (o *MemberCreateRequestBody) SetEnableApp2app(v bool) {
	o.EnableApp2app = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *MemberCreateRequestBody) GetMember() MemberCreateRequest {
	if o == nil || o.Member == nil {
		var ret MemberCreateRequest
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCreateRequestBody) GetMemberOk() (*MemberCreateRequest, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *MemberCreateRequestBody) HasMember() bool {
	if o != nil && o.Member != nil {
		return true
	}

	return false
}

// SetMember gets a reference to the given MemberCreateRequest and assigns it to the Member field.
func (o *MemberCreateRequestBody) SetMember(v MemberCreateRequest) {
	o.Member = &v
}

// GetReferralSource returns the ReferralSource field value if set, zero value otherwise.
func (o *MemberCreateRequestBody) GetReferralSource() string {
	if o == nil || o.ReferralSource == nil {
		var ret string
		return ret
	}
	return *o.ReferralSource
}

// GetReferralSourceOk returns a tuple with the ReferralSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCreateRequestBody) GetReferralSourceOk() (*string, bool) {
	if o == nil || o.ReferralSource == nil {
		return nil, false
	}
	return o.ReferralSource, true
}

// HasReferralSource returns a boolean if a field has been set.
func (o *MemberCreateRequestBody) HasReferralSource() bool {
	if o != nil && o.ReferralSource != nil {
		return true
	}

	return false
}

// SetReferralSource gets a reference to the given string and assigns it to the ReferralSource field.
func (o *MemberCreateRequestBody) SetReferralSource(v string) {
	o.ReferralSource = &v
}

// GetUiMessageWebviewUrlScheme returns the UiMessageWebviewUrlScheme field value if set, zero value otherwise.
func (o *MemberCreateRequestBody) GetUiMessageWebviewUrlScheme() string {
	if o == nil || o.UiMessageWebviewUrlScheme == nil {
		var ret string
		return ret
	}
	return *o.UiMessageWebviewUrlScheme
}

// GetUiMessageWebviewUrlSchemeOk returns a tuple with the UiMessageWebviewUrlScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberCreateRequestBody) GetUiMessageWebviewUrlSchemeOk() (*string, bool) {
	if o == nil || o.UiMessageWebviewUrlScheme == nil {
		return nil, false
	}
	return o.UiMessageWebviewUrlScheme, true
}

// HasUiMessageWebviewUrlScheme returns a boolean if a field has been set.
func (o *MemberCreateRequestBody) HasUiMessageWebviewUrlScheme() bool {
	if o != nil && o.UiMessageWebviewUrlScheme != nil {
		return true
	}

	return false
}

// SetUiMessageWebviewUrlScheme gets a reference to the given string and assigns it to the UiMessageWebviewUrlScheme field.
func (o *MemberCreateRequestBody) SetUiMessageWebviewUrlScheme(v string) {
	o.UiMessageWebviewUrlScheme = &v
}

func (o MemberCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientRedirectUrl != nil {
		toSerialize["client_redirect_url"] = o.ClientRedirectUrl
	}
	if o.EnableApp2app != nil {
		toSerialize["enable_app2app"] = o.EnableApp2app
	}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	if o.ReferralSource != nil {
		toSerialize["referral_source"] = o.ReferralSource
	}
	if o.UiMessageWebviewUrlScheme != nil {
		toSerialize["ui_message_webview_url_scheme"] = o.UiMessageWebviewUrlScheme
	}
	return json.Marshal(toSerialize)
}

type NullableMemberCreateRequestBody struct {
	value *MemberCreateRequestBody
	isSet bool
}

func (v NullableMemberCreateRequestBody) Get() *MemberCreateRequestBody {
	return v.value
}

func (v *NullableMemberCreateRequestBody) Set(val *MemberCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberCreateRequestBody(val *MemberCreateRequestBody) *NullableMemberCreateRequestBody {
	return &NullableMemberCreateRequestBody{value: val, isSet: true}
}

func (v NullableMemberCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


