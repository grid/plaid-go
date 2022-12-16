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

// CreditBankIncomeWarning The warning associated with the data that was unavailable for the Bank Income Report.
type CreditBankIncomeWarning struct {
	WarningType *CreditBankIncomeWarningType `json:"warning_type,omitempty"`
	WarningCode *CreditBankIncomeWarningCode `json:"warning_code,omitempty"`
	Cause *CreditBankIncomeCause `json:"cause,omitempty"`
}

// NewCreditBankIncomeWarning instantiates a new CreditBankIncomeWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeWarning() *CreditBankIncomeWarning {
	this := CreditBankIncomeWarning{}
	return &this
}

// NewCreditBankIncomeWarningWithDefaults instantiates a new CreditBankIncomeWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeWarningWithDefaults() *CreditBankIncomeWarning {
	this := CreditBankIncomeWarning{}
	return &this
}

// GetWarningType returns the WarningType field value if set, zero value otherwise.
func (o *CreditBankIncomeWarning) GetWarningType() CreditBankIncomeWarningType {
	if o == nil || o.WarningType == nil {
		var ret CreditBankIncomeWarningType
		return ret
	}
	return *o.WarningType
}

// GetWarningTypeOk returns a tuple with the WarningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWarning) GetWarningTypeOk() (*CreditBankIncomeWarningType, bool) {
	if o == nil || o.WarningType == nil {
		return nil, false
	}
	return o.WarningType, true
}

// HasWarningType returns a boolean if a field has been set.
func (o *CreditBankIncomeWarning) HasWarningType() bool {
	if o != nil && o.WarningType != nil {
		return true
	}

	return false
}

// SetWarningType gets a reference to the given CreditBankIncomeWarningType and assigns it to the WarningType field.
func (o *CreditBankIncomeWarning) SetWarningType(v CreditBankIncomeWarningType) {
	o.WarningType = &v
}

// GetWarningCode returns the WarningCode field value if set, zero value otherwise.
func (o *CreditBankIncomeWarning) GetWarningCode() CreditBankIncomeWarningCode {
	if o == nil || o.WarningCode == nil {
		var ret CreditBankIncomeWarningCode
		return ret
	}
	return *o.WarningCode
}

// GetWarningCodeOk returns a tuple with the WarningCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWarning) GetWarningCodeOk() (*CreditBankIncomeWarningCode, bool) {
	if o == nil || o.WarningCode == nil {
		return nil, false
	}
	return o.WarningCode, true
}

// HasWarningCode returns a boolean if a field has been set.
func (o *CreditBankIncomeWarning) HasWarningCode() bool {
	if o != nil && o.WarningCode != nil {
		return true
	}

	return false
}

// SetWarningCode gets a reference to the given CreditBankIncomeWarningCode and assigns it to the WarningCode field.
func (o *CreditBankIncomeWarning) SetWarningCode(v CreditBankIncomeWarningCode) {
	o.WarningCode = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *CreditBankIncomeWarning) GetCause() CreditBankIncomeCause {
	if o == nil || o.Cause == nil {
		var ret CreditBankIncomeCause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeWarning) GetCauseOk() (*CreditBankIncomeCause, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *CreditBankIncomeWarning) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given CreditBankIncomeCause and assigns it to the Cause field.
func (o *CreditBankIncomeWarning) SetCause(v CreditBankIncomeCause) {
	o.Cause = &v
}

func (o CreditBankIncomeWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WarningType != nil {
		toSerialize["warning_type"] = o.WarningType
	}
	if o.WarningCode != nil {
		toSerialize["warning_code"] = o.WarningCode
	}
	if o.Cause != nil {
		toSerialize["cause"] = o.Cause
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeWarning struct {
	value *CreditBankIncomeWarning
	isSet bool
}

func (v NullableCreditBankIncomeWarning) Get() *CreditBankIncomeWarning {
	return v.value
}

func (v *NullableCreditBankIncomeWarning) Set(val *CreditBankIncomeWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeWarning(val *CreditBankIncomeWarning) *NullableCreditBankIncomeWarning {
	return &NullableCreditBankIncomeWarning{value: val, isSet: true}
}

func (v NullableCreditBankIncomeWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


