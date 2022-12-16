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

// PaymentInitiationConsentStatus The status of the payment consent.  `UNAUTHORISED`: Consent created, but requires user authorisation.  `REJECTED`: Consent authorisation was rejected by the user and/or the bank.  `AUTHORISED`: Consent is active and ready to be used.  `REVOKED`: Consent has been revoked and can no longer be used.  `EXPIRED`: Consent is no longer valid.
type PaymentInitiationConsentStatus string

var _ = fmt.Printf

// List of PaymentInitiationConsentStatus
const (
	PAYMENTINITIATIONCONSENTSTATUS_UNAUTHORISED PaymentInitiationConsentStatus = "UNAUTHORISED"
	PAYMENTINITIATIONCONSENTSTATUS_AUTHORISED PaymentInitiationConsentStatus = "AUTHORISED"
	PAYMENTINITIATIONCONSENTSTATUS_REVOKED PaymentInitiationConsentStatus = "REVOKED"
	PAYMENTINITIATIONCONSENTSTATUS_REJECTED PaymentInitiationConsentStatus = "REJECTED"
	PAYMENTINITIATIONCONSENTSTATUS_EXPIRED PaymentInitiationConsentStatus = "EXPIRED"
)

var allowedPaymentInitiationConsentStatusEnumValues = []PaymentInitiationConsentStatus{
	"UNAUTHORISED",
	"AUTHORISED",
	"REVOKED",
	"REJECTED",
	"EXPIRED",
}

func (v *PaymentInitiationConsentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentInitiationConsentStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentInitiationConsentStatusFromValue returns a pointer to a valid PaymentInitiationConsentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentInitiationConsentStatusFromValue(v string) (*PaymentInitiationConsentStatus, error) {
	ev := PaymentInitiationConsentStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentInitiationConsentStatus) IsValid() bool {
	for _, existing := range allowedPaymentInitiationConsentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentInitiationConsentStatus value
func (v PaymentInitiationConsentStatus) Ptr() *PaymentInitiationConsentStatus {
	return &v
}

type NullablePaymentInitiationConsentStatus struct {
	value *PaymentInitiationConsentStatus
	isSet bool
}

func (v NullablePaymentInitiationConsentStatus) Get() *PaymentInitiationConsentStatus {
	return v.value
}

func (v *NullablePaymentInitiationConsentStatus) Set(val *PaymentInitiationConsentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentStatus(val *PaymentInitiationConsentStatus) *NullablePaymentInitiationConsentStatus {
	return &NullablePaymentInitiationConsentStatus{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

