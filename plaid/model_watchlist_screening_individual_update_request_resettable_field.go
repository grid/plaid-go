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
	"fmt"
)

// WatchlistScreeningIndividualUpdateRequestResettableField The name of a field that can be reset back to null
type WatchlistScreeningIndividualUpdateRequestResettableField string

var _ = fmt.Printf

// List of WatchlistScreeningIndividualUpdateRequestResettableField
const (
	WATCHLISTSCREENINGINDIVIDUALUPDATEREQUESTRESETTABLEFIELD_ASSIGNEE WatchlistScreeningIndividualUpdateRequestResettableField = "assignee"
)

var allowedWatchlistScreeningIndividualUpdateRequestResettableFieldEnumValues = []WatchlistScreeningIndividualUpdateRequestResettableField{
	"assignee",
}

func (v *WatchlistScreeningIndividualUpdateRequestResettableField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WatchlistScreeningIndividualUpdateRequestResettableField(value)


	*v = enumTypeValue
	return nil
}

// NewWatchlistScreeningIndividualUpdateRequestResettableFieldFromValue returns a pointer to a valid WatchlistScreeningIndividualUpdateRequestResettableField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatchlistScreeningIndividualUpdateRequestResettableFieldFromValue(v string) (*WatchlistScreeningIndividualUpdateRequestResettableField, error) {
	ev := WatchlistScreeningIndividualUpdateRequestResettableField(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatchlistScreeningIndividualUpdateRequestResettableField) IsValid() bool {
	for _, existing := range allowedWatchlistScreeningIndividualUpdateRequestResettableFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchlistScreeningIndividualUpdateRequestResettableField value
func (v WatchlistScreeningIndividualUpdateRequestResettableField) Ptr() *WatchlistScreeningIndividualUpdateRequestResettableField {
	return &v
}

type NullableWatchlistScreeningIndividualUpdateRequestResettableField struct {
	value *WatchlistScreeningIndividualUpdateRequestResettableField
	isSet bool
}

func (v NullableWatchlistScreeningIndividualUpdateRequestResettableField) Get() *WatchlistScreeningIndividualUpdateRequestResettableField {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualUpdateRequestResettableField) Set(val *WatchlistScreeningIndividualUpdateRequestResettableField) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualUpdateRequestResettableField) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualUpdateRequestResettableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualUpdateRequestResettableField(val *WatchlistScreeningIndividualUpdateRequestResettableField) *NullableWatchlistScreeningIndividualUpdateRequestResettableField {
	return &NullableWatchlistScreeningIndividualUpdateRequestResettableField{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualUpdateRequestResettableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualUpdateRequestResettableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

