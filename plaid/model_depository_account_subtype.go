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

// DepositoryAccountSubtype Valid account subtypes for depository accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-depository).
type DepositoryAccountSubtype string

var _ = fmt.Printf

// List of DepositoryAccountSubtype
const (
	DEPOSITORYACCOUNTSUBTYPE_CHECKING DepositoryAccountSubtype = "checking"
	DEPOSITORYACCOUNTSUBTYPE_SAVINGS DepositoryAccountSubtype = "savings"
	DEPOSITORYACCOUNTSUBTYPE_HSA DepositoryAccountSubtype = "hsa"
	DEPOSITORYACCOUNTSUBTYPE_CD DepositoryAccountSubtype = "cd"
	DEPOSITORYACCOUNTSUBTYPE_MONEY_MARKET DepositoryAccountSubtype = "money market"
	DEPOSITORYACCOUNTSUBTYPE_PAYPAL DepositoryAccountSubtype = "paypal"
	DEPOSITORYACCOUNTSUBTYPE_PREPAID DepositoryAccountSubtype = "prepaid"
	DEPOSITORYACCOUNTSUBTYPE_CASH_MANAGEMENT DepositoryAccountSubtype = "cash management"
	DEPOSITORYACCOUNTSUBTYPE_EBT DepositoryAccountSubtype = "ebt"
	DEPOSITORYACCOUNTSUBTYPE_ALL DepositoryAccountSubtype = "all"
)

var allowedDepositoryAccountSubtypeEnumValues = []DepositoryAccountSubtype{
	"checking",
	"savings",
	"hsa",
	"cd",
	"money market",
	"paypal",
	"prepaid",
	"cash management",
	"ebt",
	"all",
}

func (v *DepositoryAccountSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DepositoryAccountSubtype(value)


	*v = enumTypeValue
	return nil
}

// NewDepositoryAccountSubtypeFromValue returns a pointer to a valid DepositoryAccountSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDepositoryAccountSubtypeFromValue(v string) (*DepositoryAccountSubtype, error) {
	ev := DepositoryAccountSubtype(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DepositoryAccountSubtype) IsValid() bool {
	for _, existing := range allowedDepositoryAccountSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DepositoryAccountSubtype value
func (v DepositoryAccountSubtype) Ptr() *DepositoryAccountSubtype {
	return &v
}

type NullableDepositoryAccountSubtype struct {
	value *DepositoryAccountSubtype
	isSet bool
}

func (v NullableDepositoryAccountSubtype) Get() *DepositoryAccountSubtype {
	return v.value
}

func (v *NullableDepositoryAccountSubtype) Set(val *DepositoryAccountSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositoryAccountSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositoryAccountSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositoryAccountSubtype(val *DepositoryAccountSubtype) *NullableDepositoryAccountSubtype {
	return &NullableDepositoryAccountSubtype{value: val, isSet: true}
}

func (v NullableDepositoryAccountSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositoryAccountSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

