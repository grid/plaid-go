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

// SandboxPublicTokenCreateRequestIncomeVerificationBankIncome Specifies options for Bank Income. This field is required if `income_verification` is included in the `initial_products` array and `bank` is specified in `income_source_types`.
type SandboxPublicTokenCreateRequestIncomeVerificationBankIncome struct {
	// The number of days of data to request for the Bank Income product
	DaysRequested *int32 `json:"days_requested,omitempty"`
}

// NewSandboxPublicTokenCreateRequestIncomeVerificationBankIncome instantiates a new SandboxPublicTokenCreateRequestIncomeVerificationBankIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxPublicTokenCreateRequestIncomeVerificationBankIncome() *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome {
	this := SandboxPublicTokenCreateRequestIncomeVerificationBankIncome{}
	return &this
}

// NewSandboxPublicTokenCreateRequestIncomeVerificationBankIncomeWithDefaults instantiates a new SandboxPublicTokenCreateRequestIncomeVerificationBankIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxPublicTokenCreateRequestIncomeVerificationBankIncomeWithDefaults() *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome {
	this := SandboxPublicTokenCreateRequestIncomeVerificationBankIncome{}
	return &this
}

// GetDaysRequested returns the DaysRequested field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) GetDaysRequested() int32 {
	if o == nil || o.DaysRequested == nil {
		var ret int32
		return ret
	}
	return *o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) GetDaysRequestedOk() (*int32, bool) {
	if o == nil || o.DaysRequested == nil {
		return nil, false
	}
	return o.DaysRequested, true
}

// HasDaysRequested returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) HasDaysRequested() bool {
	if o != nil && o.DaysRequested != nil {
		return true
	}

	return false
}

// SetDaysRequested gets a reference to the given int32 and assigns it to the DaysRequested field.
func (o *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) SetDaysRequested(v int32) {
	o.DaysRequested = &v
}

func (o SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysRequested != nil {
		toSerialize["days_requested"] = o.DaysRequested
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome struct {
	value *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome
	isSet bool
}

func (v NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome) Get() *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome {
	return v.value
}

func (v *NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome) Set(val *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome(val *SandboxPublicTokenCreateRequestIncomeVerificationBankIncome) *NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome {
	return &NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome{value: val, isSet: true}
}

func (v NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxPublicTokenCreateRequestIncomeVerificationBankIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


