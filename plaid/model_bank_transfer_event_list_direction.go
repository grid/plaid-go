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
	"fmt"
)

// BankTransferEventListDirection Indicates the direction of the transfer: `outbound`: for API-initiated transfers `inbound`: for payments received by the FBO account.
type BankTransferEventListDirection string

var _ = fmt.Printf

// List of BankTransferEventListDirection
const (
	BANKTRANSFEREVENTLISTDIRECTION_INBOUND BankTransferEventListDirection = "inbound"
	BANKTRANSFEREVENTLISTDIRECTION_OUTBOUND BankTransferEventListDirection = "outbound"
	BANKTRANSFEREVENTLISTDIRECTION_NULL BankTransferEventListDirection = "null"
)

var allowedBankTransferEventListDirectionEnumValues = []BankTransferEventListDirection{
	"inbound",
	"outbound",
	"null",
}

func (v *BankTransferEventListDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BankTransferEventListDirection(value)


	*v = enumTypeValue
	return nil
}

// NewBankTransferEventListDirectionFromValue returns a pointer to a valid BankTransferEventListDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankTransferEventListDirectionFromValue(v string) (*BankTransferEventListDirection, error) {
	ev := BankTransferEventListDirection(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankTransferEventListDirection) IsValid() bool {
	for _, existing := range allowedBankTransferEventListDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankTransferEventListDirection value
func (v BankTransferEventListDirection) Ptr() *BankTransferEventListDirection {
	return &v
}

type NullableBankTransferEventListDirection struct {
	value *BankTransferEventListDirection
	isSet bool
}

func (v NullableBankTransferEventListDirection) Get() *BankTransferEventListDirection {
	return v.value
}

func (v *NullableBankTransferEventListDirection) Set(val *BankTransferEventListDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEventListDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEventListDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEventListDirection(val *BankTransferEventListDirection) *NullableBankTransferEventListDirection {
	return &NullableBankTransferEventListDirection{value: val, isSet: true}
}

func (v NullableBankTransferEventListDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEventListDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

