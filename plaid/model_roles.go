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

// Roles No documentation available
type Roles struct {
	ROLE Role `json:"ROLE"`
	AdditionalProperties map[string]interface{}
}

type _Roles Roles

// NewRoles instantiates a new Roles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoles(rOLE Role) *Roles {
	this := Roles{}
	this.ROLE = rOLE
	return &this
}

// NewRolesWithDefaults instantiates a new Roles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesWithDefaults() *Roles {
	this := Roles{}
	return &this
}

// GetROLE returns the ROLE field value
func (o *Roles) GetROLE() Role {
	if o == nil {
		var ret Role
		return ret
	}

	return o.ROLE
}

// GetROLEOk returns a tuple with the ROLE field value
// and a boolean to check if the value has been set.
func (o *Roles) GetROLEOk() (*Role, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ROLE, true
}

// SetROLE sets field value
func (o *Roles) SetROLE(v Role) {
	o.ROLE = v
}

func (o Roles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ROLE"] = o.ROLE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Roles) UnmarshalJSON(bytes []byte) (err error) {
	varRoles := _Roles{}

	if err = json.Unmarshal(bytes, &varRoles); err == nil {
		*o = Roles(varRoles)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ROLE")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoles struct {
	value *Roles
	isSet bool
}

func (v NullableRoles) Get() *Roles {
	return v.value
}

func (v *NullableRoles) Set(val *Roles) {
	v.value = val
	v.isSet = true
}

func (v NullableRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoles(val *Roles) *NullableRoles {
	return &NullableRoles{value: val, isSet: true}
}

func (v NullableRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


