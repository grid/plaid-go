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

// ScreeningHitDateOfBirthItem Analyzed date of birth for the associated hit
type ScreeningHitDateOfBirthItem struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *DateRange `json:"data,omitempty"`
}

// NewScreeningHitDateOfBirthItem instantiates a new ScreeningHitDateOfBirthItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreeningHitDateOfBirthItem() *ScreeningHitDateOfBirthItem {
	this := ScreeningHitDateOfBirthItem{}
	return &this
}

// NewScreeningHitDateOfBirthItemWithDefaults instantiates a new ScreeningHitDateOfBirthItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreeningHitDateOfBirthItemWithDefaults() *ScreeningHitDateOfBirthItem {
	this := ScreeningHitDateOfBirthItem{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *ScreeningHitDateOfBirthItem) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitDateOfBirthItem) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *ScreeningHitDateOfBirthItem) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *ScreeningHitDateOfBirthItem) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScreeningHitDateOfBirthItem) GetData() DateRange {
	if o == nil || o.Data == nil {
		var ret DateRange
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitDateOfBirthItem) GetDataOk() (*DateRange, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScreeningHitDateOfBirthItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given DateRange and assigns it to the Data field.
func (o *ScreeningHitDateOfBirthItem) SetData(v DateRange) {
	o.Data = &v
}

func (o ScreeningHitDateOfBirthItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableScreeningHitDateOfBirthItem struct {
	value *ScreeningHitDateOfBirthItem
	isSet bool
}

func (v NullableScreeningHitDateOfBirthItem) Get() *ScreeningHitDateOfBirthItem {
	return v.value
}

func (v *NullableScreeningHitDateOfBirthItem) Set(val *ScreeningHitDateOfBirthItem) {
	v.value = val
	v.isSet = true
}

func (v NullableScreeningHitDateOfBirthItem) IsSet() bool {
	return v.isSet
}

func (v *NullableScreeningHitDateOfBirthItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreeningHitDateOfBirthItem(val *ScreeningHitDateOfBirthItem) *NullableScreeningHitDateOfBirthItem {
	return &NullableScreeningHitDateOfBirthItem{value: val, isSet: true}
}

func (v NullableScreeningHitDateOfBirthItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreeningHitDateOfBirthItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


