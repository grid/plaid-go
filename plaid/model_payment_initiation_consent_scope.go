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

// PaymentInitiationConsentScope Payment consent scope. Defines possible directions for payments made with the given consent.  `ME_TO_ME`: Allows moving money between accounts owned by the same user.  `EXTERNAL`: Allows initiating payments from the user's account to third parties.
type PaymentInitiationConsentScope string

// List of PaymentInitiationConsentScope
const (
	PAYMENTINITIATIONCONSENTSCOPE_ME_TO_ME PaymentInitiationConsentScope = "ME_TO_ME"
	PAYMENTINITIATIONCONSENTSCOPE_EXTERNAL PaymentInitiationConsentScope = "EXTERNAL"
)

var allowedPaymentInitiationConsentScopeEnumValues = []PaymentInitiationConsentScope{
	"ME_TO_ME",
	"EXTERNAL",
}

func (v *PaymentInitiationConsentScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentInitiationConsentScope(value)
	for _, existing := range allowedPaymentInitiationConsentScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentInitiationConsentScope", value)
}

// NewPaymentInitiationConsentScopeFromValue returns a pointer to a valid PaymentInitiationConsentScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentInitiationConsentScopeFromValue(v string) (*PaymentInitiationConsentScope, error) {
	ev := PaymentInitiationConsentScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentInitiationConsentScope: valid values are %v", v, allowedPaymentInitiationConsentScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentInitiationConsentScope) IsValid() bool {
	for _, existing := range allowedPaymentInitiationConsentScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentInitiationConsentScope value
func (v PaymentInitiationConsentScope) Ptr() *PaymentInitiationConsentScope {
	return &v
}

type NullablePaymentInitiationConsentScope struct {
	value *PaymentInitiationConsentScope
	isSet bool
}

func (v NullablePaymentInitiationConsentScope) Get() *PaymentInitiationConsentScope {
	return v.value
}

func (v *NullablePaymentInitiationConsentScope) Set(val *PaymentInitiationConsentScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentScope(val *PaymentInitiationConsentScope) *NullablePaymentInitiationConsentScope {
	return &NullablePaymentInitiationConsentScope{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

