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

// MatchSummary Summary object reflecting the match result of the associated data
type MatchSummary struct {
	Summary MatchSummaryCode `json:"summary"`
}

// NewMatchSummary instantiates a new MatchSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchSummary(summary MatchSummaryCode) *MatchSummary {
	this := MatchSummary{}
	this.Summary = summary
	return &this
}

// NewMatchSummaryWithDefaults instantiates a new MatchSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchSummaryWithDefaults() *MatchSummary {
	this := MatchSummary{}
	return &this
}

// GetSummary returns the Summary field value
func (o *MatchSummary) GetSummary() MatchSummaryCode {
	if o == nil {
		var ret MatchSummaryCode
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *MatchSummary) GetSummaryOk() (*MatchSummaryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *MatchSummary) SetSummary(v MatchSummaryCode) {
	o.Summary = v
}

func (o MatchSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableMatchSummary struct {
	value *MatchSummary
	isSet bool
}

func (v NullableMatchSummary) Get() *MatchSummary {
	return v.value
}

func (v *NullableMatchSummary) Set(val *MatchSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchSummary(val *MatchSummary) *NullableMatchSummary {
	return &NullableMatchSummary{value: val, isSet: true}
}

func (v NullableMatchSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


