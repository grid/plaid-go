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

// IncomeVerificationPrecheckPayrollInstitution Information about the end user's payroll institution
type IncomeVerificationPrecheckPayrollInstitution struct {
	// The name of payroll institution
	Name NullableString `json:"name,omitempty"`
}

// NewIncomeVerificationPrecheckPayrollInstitution instantiates a new IncomeVerificationPrecheckPayrollInstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckPayrollInstitution() *IncomeVerificationPrecheckPayrollInstitution {
	this := IncomeVerificationPrecheckPayrollInstitution{}
	return &this
}

// NewIncomeVerificationPrecheckPayrollInstitutionWithDefaults instantiates a new IncomeVerificationPrecheckPayrollInstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckPayrollInstitutionWithDefaults() *IncomeVerificationPrecheckPayrollInstitution {
	this := IncomeVerificationPrecheckPayrollInstitution{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckPayrollInstitution) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckPayrollInstitution) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckPayrollInstitution) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IncomeVerificationPrecheckPayrollInstitution) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IncomeVerificationPrecheckPayrollInstitution) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IncomeVerificationPrecheckPayrollInstitution) UnsetName() {
	o.Name.Unset()
}

func (o IncomeVerificationPrecheckPayrollInstitution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckPayrollInstitution struct {
	value *IncomeVerificationPrecheckPayrollInstitution
	isSet bool
}

func (v NullableIncomeVerificationPrecheckPayrollInstitution) Get() *IncomeVerificationPrecheckPayrollInstitution {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckPayrollInstitution) Set(val *IncomeVerificationPrecheckPayrollInstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckPayrollInstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckPayrollInstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckPayrollInstitution(val *IncomeVerificationPrecheckPayrollInstitution) *NullableIncomeVerificationPrecheckPayrollInstitution {
	return &NullableIncomeVerificationPrecheckPayrollInstitution{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckPayrollInstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckPayrollInstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


