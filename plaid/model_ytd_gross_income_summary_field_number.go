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

// YTDGrossIncomeSummaryFieldNumber Year-to-date pre-tax earnings, as reported on the paystub.
type YTDGrossIncomeSummaryFieldNumber struct {
	// The value of the field.
	Value float64 `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
}

// NewYTDGrossIncomeSummaryFieldNumber instantiates a new YTDGrossIncomeSummaryFieldNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYTDGrossIncomeSummaryFieldNumber(value float64, verificationStatus VerificationStatus) *YTDGrossIncomeSummaryFieldNumber {
	this := YTDGrossIncomeSummaryFieldNumber{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewYTDGrossIncomeSummaryFieldNumberWithDefaults instantiates a new YTDGrossIncomeSummaryFieldNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYTDGrossIncomeSummaryFieldNumberWithDefaults() *YTDGrossIncomeSummaryFieldNumber {
	this := YTDGrossIncomeSummaryFieldNumber{}
	return &this
}

// GetValue returns the Value field value
func (o *YTDGrossIncomeSummaryFieldNumber) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *YTDGrossIncomeSummaryFieldNumber) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *YTDGrossIncomeSummaryFieldNumber) SetValue(v float64) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *YTDGrossIncomeSummaryFieldNumber) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *YTDGrossIncomeSummaryFieldNumber) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *YTDGrossIncomeSummaryFieldNumber) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o YTDGrossIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableYTDGrossIncomeSummaryFieldNumber struct {
	value *YTDGrossIncomeSummaryFieldNumber
	isSet bool
}

func (v NullableYTDGrossIncomeSummaryFieldNumber) Get() *YTDGrossIncomeSummaryFieldNumber {
	return v.value
}

func (v *NullableYTDGrossIncomeSummaryFieldNumber) Set(val *YTDGrossIncomeSummaryFieldNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableYTDGrossIncomeSummaryFieldNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableYTDGrossIncomeSummaryFieldNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYTDGrossIncomeSummaryFieldNumber(val *YTDGrossIncomeSummaryFieldNumber) *NullableYTDGrossIncomeSummaryFieldNumber {
	return &NullableYTDGrossIncomeSummaryFieldNumber{value: val, isSet: true}
}

func (v NullableYTDGrossIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYTDGrossIncomeSummaryFieldNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


