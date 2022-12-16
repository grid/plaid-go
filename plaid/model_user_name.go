/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.210.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// UserName The full name provided by the user. If the user has not submitted their name, this field will be null. Otherwise, both fields are guaranteed to be filled.
type UserName struct {
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	GivenName string `json:"given_name"`
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	FamilyName string `json:"family_name"`
}

// NewUserName instantiates a new UserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserName(givenName string, familyName string) *UserName {
	this := UserName{}
	this.GivenName = givenName
	this.FamilyName = familyName
	return &this
}

// NewUserNameWithDefaults instantiates a new UserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNameWithDefaults() *UserName {
	this := UserName{}
	return &this
}

// GetGivenName returns the GivenName field value
func (o *UserName) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *UserName) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *UserName) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *UserName) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *UserName) GetFamilyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *UserName) SetFamilyName(v string) {
	o.FamilyName = v
}

func (o UserName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["given_name"] = o.GivenName
	}
	if true {
		toSerialize["family_name"] = o.FamilyName
	}
	return json.Marshal(toSerialize)
}

type NullableUserName struct {
	value *UserName
	isSet bool
}

func (v NullableUserName) Get() *UserName {
	return v.value
}

func (v *NullableUserName) Set(val *UserName) {
	v.value = val
	v.isSet = true
}

func (v NullableUserName) IsSet() bool {
	return v.isSet
}

func (v *NullableUserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserName(val *UserName) *NullableUserName {
	return &NullableUserName{value: val, isSet: true}
}

func (v NullableUserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


