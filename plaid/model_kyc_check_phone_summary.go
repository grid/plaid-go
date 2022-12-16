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

// KYCCheckPhoneSummary Result summary object specifying how the `phone` field matched.
type KYCCheckPhoneSummary struct {
	Summary MatchSummaryCode `json:"summary"`
}

// NewKYCCheckPhoneSummary instantiates a new KYCCheckPhoneSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckPhoneSummary(summary MatchSummaryCode) *KYCCheckPhoneSummary {
	this := KYCCheckPhoneSummary{}
	this.Summary = summary
	return &this
}

// NewKYCCheckPhoneSummaryWithDefaults instantiates a new KYCCheckPhoneSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckPhoneSummaryWithDefaults() *KYCCheckPhoneSummary {
	this := KYCCheckPhoneSummary{}
	return &this
}

// GetSummary returns the Summary field value
func (o *KYCCheckPhoneSummary) GetSummary() MatchSummaryCode {
	if o == nil {
		var ret MatchSummaryCode
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *KYCCheckPhoneSummary) GetSummaryOk() (*MatchSummaryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *KYCCheckPhoneSummary) SetSummary(v MatchSummaryCode) {
	o.Summary = v
}

func (o KYCCheckPhoneSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableKYCCheckPhoneSummary struct {
	value *KYCCheckPhoneSummary
	isSet bool
}

func (v NullableKYCCheckPhoneSummary) Get() *KYCCheckPhoneSummary {
	return v.value
}

func (v *NullableKYCCheckPhoneSummary) Set(val *KYCCheckPhoneSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckPhoneSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckPhoneSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckPhoneSummary(val *KYCCheckPhoneSummary) *NullableKYCCheckPhoneSummary {
	return &NullableKYCCheckPhoneSummary{value: val, isSet: true}
}

func (v NullableKYCCheckPhoneSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckPhoneSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


