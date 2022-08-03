/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.161.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// UserStatedIncomeSourceFrequency The pay frequency of a specified income source
type UserStatedIncomeSourceFrequency string

// List of UserStatedIncomeSourceFrequency
const (
	USERSTATEDINCOMESOURCEFREQUENCY_UNKNOWN UserStatedIncomeSourceFrequency = "UNKNOWN"
	USERSTATEDINCOMESOURCEFREQUENCY_WEEKLY UserStatedIncomeSourceFrequency = "WEEKLY"
	USERSTATEDINCOMESOURCEFREQUENCY_BIWEEKLY UserStatedIncomeSourceFrequency = "BIWEEKLY"
	USERSTATEDINCOMESOURCEFREQUENCY_SEMI_MONTHLY UserStatedIncomeSourceFrequency = "SEMI_MONTHLY"
	USERSTATEDINCOMESOURCEFREQUENCY_MONTHLY UserStatedIncomeSourceFrequency = "MONTHLY"
)

var allowedUserStatedIncomeSourceFrequencyEnumValues = []UserStatedIncomeSourceFrequency{
	"UNKNOWN",
	"WEEKLY",
	"BIWEEKLY",
	"SEMI_MONTHLY",
	"MONTHLY",
}

func (v *UserStatedIncomeSourceFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserStatedIncomeSourceFrequency(value)
	for _, existing := range allowedUserStatedIncomeSourceFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserStatedIncomeSourceFrequency", value)
}

// NewUserStatedIncomeSourceFrequencyFromValue returns a pointer to a valid UserStatedIncomeSourceFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserStatedIncomeSourceFrequencyFromValue(v string) (*UserStatedIncomeSourceFrequency, error) {
	ev := UserStatedIncomeSourceFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserStatedIncomeSourceFrequency: valid values are %v", v, allowedUserStatedIncomeSourceFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserStatedIncomeSourceFrequency) IsValid() bool {
	for _, existing := range allowedUserStatedIncomeSourceFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserStatedIncomeSourceFrequency value
func (v UserStatedIncomeSourceFrequency) Ptr() *UserStatedIncomeSourceFrequency {
	return &v
}

type NullableUserStatedIncomeSourceFrequency struct {
	value *UserStatedIncomeSourceFrequency
	isSet bool
}

func (v NullableUserStatedIncomeSourceFrequency) Get() *UserStatedIncomeSourceFrequency {
	return v.value
}

func (v *NullableUserStatedIncomeSourceFrequency) Set(val *UserStatedIncomeSourceFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatedIncomeSourceFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatedIncomeSourceFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatedIncomeSourceFrequency(val *UserStatedIncomeSourceFrequency) *NullableUserStatedIncomeSourceFrequency {
	return &NullableUserStatedIncomeSourceFrequency{value: val, isSet: true}
}

func (v NullableUserStatedIncomeSourceFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatedIncomeSourceFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

