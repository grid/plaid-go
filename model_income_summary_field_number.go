/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeSummaryFieldNumber struct for IncomeSummaryFieldNumber
type IncomeSummaryFieldNumber struct {
	// The value of the field.
	Value float32 `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
	AdditionalProperties map[string]interface{}
}

type _IncomeSummaryFieldNumber IncomeSummaryFieldNumber

// NewIncomeSummaryFieldNumber instantiates a new IncomeSummaryFieldNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeSummaryFieldNumber(value float32, verificationStatus VerificationStatus) *IncomeSummaryFieldNumber {
	this := IncomeSummaryFieldNumber{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewIncomeSummaryFieldNumberWithDefaults instantiates a new IncomeSummaryFieldNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeSummaryFieldNumberWithDefaults() *IncomeSummaryFieldNumber {
	this := IncomeSummaryFieldNumber{}
	return &this
}

// GetValue returns the Value field value
func (o *IncomeSummaryFieldNumber) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IncomeSummaryFieldNumber) GetValueOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IncomeSummaryFieldNumber) SetValue(v float32) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *IncomeSummaryFieldNumber) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *IncomeSummaryFieldNumber) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *IncomeSummaryFieldNumber) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o IncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeSummaryFieldNumber) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeSummaryFieldNumber := _IncomeSummaryFieldNumber{}

	if err = json.Unmarshal(bytes, &varIncomeSummaryFieldNumber); err == nil {
		*o = IncomeSummaryFieldNumber(varIncomeSummaryFieldNumber)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "verification_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeSummaryFieldNumber struct {
	value *IncomeSummaryFieldNumber
	isSet bool
}

func (v NullableIncomeSummaryFieldNumber) Get() *IncomeSummaryFieldNumber {
	return v.value
}

func (v *NullableIncomeSummaryFieldNumber) Set(val *IncomeSummaryFieldNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeSummaryFieldNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeSummaryFieldNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeSummaryFieldNumber(val *IncomeSummaryFieldNumber) *NullableIncomeSummaryFieldNumber {
	return &NullableIncomeSummaryFieldNumber{value: val, isSet: true}
}

func (v NullableIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeSummaryFieldNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


