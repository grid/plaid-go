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
	"fmt"
)

// PayFrequencyValue The frequency of the pay period.
type PayFrequencyValue string

// List of PayFrequencyValue
const (
	PAYFREQUENCYVALUE_MONTHLY PayFrequencyValue = "monthly"
	PAYFREQUENCYVALUE_SEMIMONTHLY PayFrequencyValue = "semimonthly"
	PAYFREQUENCYVALUE_WEEKLY PayFrequencyValue = "weekly"
	PAYFREQUENCYVALUE_BIWEEKLY PayFrequencyValue = "biweekly"
	PAYFREQUENCYVALUE_UNKNOWN PayFrequencyValue = "unknown"
	PAYFREQUENCYVALUE_NULL PayFrequencyValue = "null"
)

var allowedPayFrequencyValueEnumValues = []PayFrequencyValue{
	"monthly",
	"semimonthly",
	"weekly",
	"biweekly",
	"unknown",
	"null",
}

func (v *PayFrequencyValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PayFrequencyValue(value)
	for _, existing := range allowedPayFrequencyValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PayFrequencyValue", value)
}

// NewPayFrequencyValueFromValue returns a pointer to a valid PayFrequencyValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPayFrequencyValueFromValue(v string) (*PayFrequencyValue, error) {
	ev := PayFrequencyValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PayFrequencyValue: valid values are %v", v, allowedPayFrequencyValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PayFrequencyValue) IsValid() bool {
	for _, existing := range allowedPayFrequencyValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PayFrequencyValue value
func (v PayFrequencyValue) Ptr() *PayFrequencyValue {
	return &v
}

type NullablePayFrequencyValue struct {
	value *PayFrequencyValue
	isSet bool
}

func (v NullablePayFrequencyValue) Get() *PayFrequencyValue {
	return v.value
}

func (v *NullablePayFrequencyValue) Set(val *PayFrequencyValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePayFrequencyValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePayFrequencyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayFrequencyValue(val *PayFrequencyValue) *NullablePayFrequencyValue {
	return &NullablePayFrequencyValue{value: val, isSet: true}
}

func (v NullablePayFrequencyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayFrequencyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

