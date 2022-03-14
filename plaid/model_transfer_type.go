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

// TransferType The type of transfer. This will be either `debit` or `credit`.  A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.
type TransferType string

// List of TransferType
const (
	TRANSFERTYPE_DEBIT TransferType = "debit"
	TRANSFERTYPE_CREDIT TransferType = "credit"
)

var allowedTransferTypeEnumValues = []TransferType{
	"debit",
	"credit",
}

func (v *TransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferType(value)
	for _, existing := range allowedTransferTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferType", value)
}

// NewTransferTypeFromValue returns a pointer to a valid TransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferTypeFromValue(v string) (*TransferType, error) {
	ev := TransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferType: valid values are %v", v, allowedTransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferType) IsValid() bool {
	for _, existing := range allowedTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferType value
func (v TransferType) Ptr() *TransferType {
	return &v
}

type NullableTransferType struct {
	value *TransferType
	isSet bool
}

func (v NullableTransferType) Get() *TransferType {
	return v.value
}

func (v *NullableTransferType) Set(val *TransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferType(val *TransferType) *NullableTransferType {
	return &NullableTransferType{value: val, isSet: true}
}

func (v NullableTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

