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

// PhoneNumberMatchScore Score found by matching phone number provided by the API with the phone number on the account at the financial institution. If the account contains multiple owners, the maximum match score is filled.
type PhoneNumberMatchScore struct {
	// Match score for normalized phone number. 100 is a perfect match and 0 is a no match. If the phone number is missing from either the API or financial institution, this is empty.
	Score NullableInt32 `json:"score,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PhoneNumberMatchScore PhoneNumberMatchScore

// NewPhoneNumberMatchScore instantiates a new PhoneNumberMatchScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumberMatchScore() *PhoneNumberMatchScore {
	this := PhoneNumberMatchScore{}
	return &this
}

// NewPhoneNumberMatchScoreWithDefaults instantiates a new PhoneNumberMatchScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberMatchScoreWithDefaults() *PhoneNumberMatchScore {
	this := PhoneNumberMatchScore{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumberMatchScore) GetScore() int32 {
	if o == nil || o.Score.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumberMatchScore) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *PhoneNumberMatchScore) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableInt32 and assigns it to the Score field.
func (o *PhoneNumberMatchScore) SetScore(v int32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *PhoneNumberMatchScore) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *PhoneNumberMatchScore) UnsetScore() {
	o.Score.Unset()
}

func (o PhoneNumberMatchScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PhoneNumberMatchScore) UnmarshalJSON(bytes []byte) (err error) {
	varPhoneNumberMatchScore := _PhoneNumberMatchScore{}

	if err = json.Unmarshal(bytes, &varPhoneNumberMatchScore); err == nil {
		*o = PhoneNumberMatchScore(varPhoneNumberMatchScore)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePhoneNumberMatchScore struct {
	value *PhoneNumberMatchScore
	isSet bool
}

func (v NullablePhoneNumberMatchScore) Get() *PhoneNumberMatchScore {
	return v.value
}

func (v *NullablePhoneNumberMatchScore) Set(val *PhoneNumberMatchScore) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumberMatchScore) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumberMatchScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumberMatchScore(val *PhoneNumberMatchScore) *NullablePhoneNumberMatchScore {
	return &NullablePhoneNumberMatchScore{value: val, isSet: true}
}

func (v NullablePhoneNumberMatchScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumberMatchScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


