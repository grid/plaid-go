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

// PaymentInitiationRefundStatus The status of the refund.  `PROCESSING`: The refund is currently being processed. The refund will automatically exit this state when processing is complete.  `INITIATED`: The refund has been successfully initiated.  `EXECUTED`: Indicates that the refund has been successfully executed.  `FAILED`: The refund has failed to be executed. This error is retryable once the root cause is resolved.
type PaymentInitiationRefundStatus string

// List of PaymentInitiationRefundStatus
const (
	PAYMENTINITIATIONREFUNDSTATUS_PROCESSING PaymentInitiationRefundStatus = "PROCESSING"
	PAYMENTINITIATIONREFUNDSTATUS_EXECUTED PaymentInitiationRefundStatus = "EXECUTED"
	PAYMENTINITIATIONREFUNDSTATUS_INITIATED PaymentInitiationRefundStatus = "INITIATED"
	PAYMENTINITIATIONREFUNDSTATUS_FAILED PaymentInitiationRefundStatus = "FAILED"
)

var allowedPaymentInitiationRefundStatusEnumValues = []PaymentInitiationRefundStatus{
	"PROCESSING",
	"EXECUTED",
	"INITIATED",
	"FAILED",
}

func (v *PaymentInitiationRefundStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentInitiationRefundStatus(value)
	for _, existing := range allowedPaymentInitiationRefundStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentInitiationRefundStatus", value)
}

// NewPaymentInitiationRefundStatusFromValue returns a pointer to a valid PaymentInitiationRefundStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentInitiationRefundStatusFromValue(v string) (*PaymentInitiationRefundStatus, error) {
	ev := PaymentInitiationRefundStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentInitiationRefundStatus: valid values are %v", v, allowedPaymentInitiationRefundStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentInitiationRefundStatus) IsValid() bool {
	for _, existing := range allowedPaymentInitiationRefundStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentInitiationRefundStatus value
func (v PaymentInitiationRefundStatus) Ptr() *PaymentInitiationRefundStatus {
	return &v
}

type NullablePaymentInitiationRefundStatus struct {
	value *PaymentInitiationRefundStatus
	isSet bool
}

func (v NullablePaymentInitiationRefundStatus) Get() *PaymentInitiationRefundStatus {
	return v.value
}

func (v *NullablePaymentInitiationRefundStatus) Set(val *PaymentInitiationRefundStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRefundStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRefundStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRefundStatus(val *PaymentInitiationRefundStatus) *NullablePaymentInitiationRefundStatus {
	return &NullablePaymentInitiationRefundStatus{value: val, isSet: true}
}

func (v NullablePaymentInitiationRefundStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRefundStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

