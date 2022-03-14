/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateRequestIncomeVerificationBankIncome Specifies options for initializing Link for use with Bank Income. This field is required if `income_verification` is included in the `products` array and `bank` is specified in `income_source_types`.
type LinkTokenCreateRequestIncomeVerificationBankIncome struct {
	// The number of days of data to request for the Bank Income product
	DaysRequested *int32 `json:"days_requested,omitempty"`
}

// NewLinkTokenCreateRequestIncomeVerificationBankIncome instantiates a new LinkTokenCreateRequestIncomeVerificationBankIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestIncomeVerificationBankIncome() *LinkTokenCreateRequestIncomeVerificationBankIncome {
	this := LinkTokenCreateRequestIncomeVerificationBankIncome{}
	return &this
}

// NewLinkTokenCreateRequestIncomeVerificationBankIncomeWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerificationBankIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestIncomeVerificationBankIncomeWithDefaults() *LinkTokenCreateRequestIncomeVerificationBankIncome {
	this := LinkTokenCreateRequestIncomeVerificationBankIncome{}
	return &this
}

// GetDaysRequested returns the DaysRequested field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerificationBankIncome) GetDaysRequested() int32 {
	if o == nil || o.DaysRequested == nil {
		var ret int32
		return ret
	}
	return *o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerificationBankIncome) GetDaysRequestedOk() (*int32, bool) {
	if o == nil || o.DaysRequested == nil {
		return nil, false
	}
	return o.DaysRequested, true
}

// HasDaysRequested returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerificationBankIncome) HasDaysRequested() bool {
	if o != nil && o.DaysRequested != nil {
		return true
	}

	return false
}

// SetDaysRequested gets a reference to the given int32 and assigns it to the DaysRequested field.
func (o *LinkTokenCreateRequestIncomeVerificationBankIncome) SetDaysRequested(v int32) {
	o.DaysRequested = &v
}

func (o LinkTokenCreateRequestIncomeVerificationBankIncome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysRequested != nil {
		toSerialize["days_requested"] = o.DaysRequested
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestIncomeVerificationBankIncome struct {
	value *LinkTokenCreateRequestIncomeVerificationBankIncome
	isSet bool
}

func (v NullableLinkTokenCreateRequestIncomeVerificationBankIncome) Get() *LinkTokenCreateRequestIncomeVerificationBankIncome {
	return v.value
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationBankIncome) Set(val *LinkTokenCreateRequestIncomeVerificationBankIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestIncomeVerificationBankIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationBankIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestIncomeVerificationBankIncome(val *LinkTokenCreateRequestIncomeVerificationBankIncome) *NullableLinkTokenCreateRequestIncomeVerificationBankIncome {
	return &NullableLinkTokenCreateRequestIncomeVerificationBankIncome{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestIncomeVerificationBankIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationBankIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


