/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.214.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieVerificationOfEmploymentVOE25 Verification of Employment Report
type CreditFreddieVerificationOfEmploymentVOE25 struct {
	DEAL *CreditFreddieVerificationOfEmploymentDealVOE25 `json:"DEAL,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieVerificationOfEmploymentVOE25 CreditFreddieVerificationOfEmploymentVOE25

// NewCreditFreddieVerificationOfEmploymentVOE25 instantiates a new CreditFreddieVerificationOfEmploymentVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieVerificationOfEmploymentVOE25() *CreditFreddieVerificationOfEmploymentVOE25 {
	this := CreditFreddieVerificationOfEmploymentVOE25{}
	return &this
}

// NewCreditFreddieVerificationOfEmploymentVOE25WithDefaults instantiates a new CreditFreddieVerificationOfEmploymentVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieVerificationOfEmploymentVOE25WithDefaults() *CreditFreddieVerificationOfEmploymentVOE25 {
	this := CreditFreddieVerificationOfEmploymentVOE25{}
	return &this
}

// GetDEAL returns the DEAL field value if set, zero value otherwise.
func (o *CreditFreddieVerificationOfEmploymentVOE25) GetDEAL() CreditFreddieVerificationOfEmploymentDealVOE25 {
	if o == nil || o.DEAL == nil {
		var ret CreditFreddieVerificationOfEmploymentDealVOE25
		return ret
	}
	return *o.DEAL
}

// GetDEALOk returns a tuple with the DEAL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditFreddieVerificationOfEmploymentVOE25) GetDEALOk() (*CreditFreddieVerificationOfEmploymentDealVOE25, bool) {
	if o == nil || o.DEAL == nil {
		return nil, false
	}
	return o.DEAL, true
}

// HasDEAL returns a boolean if a field has been set.
func (o *CreditFreddieVerificationOfEmploymentVOE25) HasDEAL() bool {
	if o != nil && o.DEAL != nil {
		return true
	}

	return false
}

// SetDEAL gets a reference to the given CreditFreddieVerificationOfEmploymentDealVOE25 and assigns it to the DEAL field.
func (o *CreditFreddieVerificationOfEmploymentVOE25) SetDEAL(v CreditFreddieVerificationOfEmploymentDealVOE25) {
	o.DEAL = &v
}

func (o CreditFreddieVerificationOfEmploymentVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DEAL != nil {
		toSerialize["DEAL"] = o.DEAL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieVerificationOfEmploymentVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieVerificationOfEmploymentVOE25 := _CreditFreddieVerificationOfEmploymentVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieVerificationOfEmploymentVOE25); err == nil {
		*o = CreditFreddieVerificationOfEmploymentVOE25(varCreditFreddieVerificationOfEmploymentVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DEAL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieVerificationOfEmploymentVOE25 struct {
	value *CreditFreddieVerificationOfEmploymentVOE25
	isSet bool
}

func (v NullableCreditFreddieVerificationOfEmploymentVOE25) Get() *CreditFreddieVerificationOfEmploymentVOE25 {
	return v.value
}

func (v *NullableCreditFreddieVerificationOfEmploymentVOE25) Set(val *CreditFreddieVerificationOfEmploymentVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieVerificationOfEmploymentVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieVerificationOfEmploymentVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieVerificationOfEmploymentVOE25(val *CreditFreddieVerificationOfEmploymentVOE25) *NullableCreditFreddieVerificationOfEmploymentVOE25 {
	return &NullableCreditFreddieVerificationOfEmploymentVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieVerificationOfEmploymentVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieVerificationOfEmploymentVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


