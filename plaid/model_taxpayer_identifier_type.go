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

// TaxpayerIdentifierType A value from a MISMO prescribed list that classifies identification numbers used by the Internal Revenue Service (IRS) in the administration of tax laws. A Social Security number (SSN) is issued by the SSA; all other taxpayer identification numbers are issued by the IRS.
type TaxpayerIdentifierType string

var _ = fmt.Printf

// List of TaxpayerIdentifierType
const (
	TAXPAYERIDENTIFIERTYPE_INDIVIDUAL_TAXPAYER_IDENTIFICATION_NUMBER TaxpayerIdentifierType = "IndividualTaxpayerIdentificationNumber"
	TAXPAYERIDENTIFIERTYPE_SOCIAL_SECURITY_NUMBER TaxpayerIdentifierType = "SocialSecurityNumber"
)

var allowedTaxpayerIdentifierTypeEnumValues = []TaxpayerIdentifierType{
	"IndividualTaxpayerIdentificationNumber",
	"SocialSecurityNumber",
}

func (v *TaxpayerIdentifierType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TaxpayerIdentifierType(value)


	*v = enumTypeValue
	return nil
}

// NewTaxpayerIdentifierTypeFromValue returns a pointer to a valid TaxpayerIdentifierType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxpayerIdentifierTypeFromValue(v string) (*TaxpayerIdentifierType, error) {
	ev := TaxpayerIdentifierType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxpayerIdentifierType) IsValid() bool {
	for _, existing := range allowedTaxpayerIdentifierTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaxpayerIdentifierType value
func (v TaxpayerIdentifierType) Ptr() *TaxpayerIdentifierType {
	return &v
}

type NullableTaxpayerIdentifierType struct {
	value *TaxpayerIdentifierType
	isSet bool
}

func (v NullableTaxpayerIdentifierType) Get() *TaxpayerIdentifierType {
	return v.value
}

func (v *NullableTaxpayerIdentifierType) Set(val *TaxpayerIdentifierType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxpayerIdentifierType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxpayerIdentifierType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxpayerIdentifierType(val *TaxpayerIdentifierType) *NullableTaxpayerIdentifierType {
	return &NullableTaxpayerIdentifierType{value: val, isSet: true}
}

func (v NullableTaxpayerIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxpayerIdentifierType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

