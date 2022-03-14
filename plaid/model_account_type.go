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

// AccountType `investment:` Investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage` instead.  `credit:` Credit card  `depository:` Depository account  `loan:` Loan account  `other:` Non-specified account type  See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.
type AccountType string

// List of AccountType
const (
	ACCOUNTTYPE_INVESTMENT AccountType = "investment"
	ACCOUNTTYPE_CREDIT AccountType = "credit"
	ACCOUNTTYPE_DEPOSITORY AccountType = "depository"
	ACCOUNTTYPE_LOAN AccountType = "loan"
	ACCOUNTTYPE_BROKERAGE AccountType = "brokerage"
	ACCOUNTTYPE_OTHER AccountType = "other"
)

var allowedAccountTypeEnumValues = []AccountType{
	"investment",
	"credit",
	"depository",
	"loan",
	"brokerage",
	"other",
}

func (v *AccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountType(value)
	for _, existing := range allowedAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountType", value)
}

// NewAccountTypeFromValue returns a pointer to a valid AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypeFromValue(v string) (*AccountType, error) {
	ev := AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountType: valid values are %v", v, allowedAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountType) IsValid() bool {
	for _, existing := range allowedAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountType value
func (v AccountType) Ptr() *AccountType {
	return &v
}

type NullableAccountType struct {
	value *AccountType
	isSet bool
}

func (v NullableAccountType) Get() *AccountType {
	return v.value
}

func (v *NullableAccountType) Set(val *AccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountType(val *AccountType) *NullableAccountType {
	return &NullableAccountType{value: val, isSet: true}
}

func (v NullableAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

