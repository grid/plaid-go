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

// TransferDiligenceStatus Originator’s diligence status.
type TransferDiligenceStatus string

var _ = fmt.Printf

// List of TransferDiligenceStatus
const (
	TRANSFERDILIGENCESTATUS_UNDER_REVIEW TransferDiligenceStatus = "under_review"
	TRANSFERDILIGENCESTATUS_APPROVED TransferDiligenceStatus = "approved"
	TRANSFERDILIGENCESTATUS_DENIED TransferDiligenceStatus = "denied"
)

var allowedTransferDiligenceStatusEnumValues = []TransferDiligenceStatus{
	"under_review",
	"approved",
	"denied",
}

func (v *TransferDiligenceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferDiligenceStatus(value)


	*v = enumTypeValue
	return nil
}

// NewTransferDiligenceStatusFromValue returns a pointer to a valid TransferDiligenceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferDiligenceStatusFromValue(v string) (*TransferDiligenceStatus, error) {
	ev := TransferDiligenceStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferDiligenceStatus) IsValid() bool {
	for _, existing := range allowedTransferDiligenceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferDiligenceStatus value
func (v TransferDiligenceStatus) Ptr() *TransferDiligenceStatus {
	return &v
}

type NullableTransferDiligenceStatus struct {
	value *TransferDiligenceStatus
	isSet bool
}

func (v NullableTransferDiligenceStatus) Get() *TransferDiligenceStatus {
	return v.value
}

func (v *NullableTransferDiligenceStatus) Set(val *TransferDiligenceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDiligenceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDiligenceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDiligenceStatus(val *TransferDiligenceStatus) *NullableTransferDiligenceStatus {
	return &NullableTransferDiligenceStatus{value: val, isSet: true}
}

func (v NullableTransferDiligenceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDiligenceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

