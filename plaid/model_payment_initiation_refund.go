/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.46.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// PaymentInitiationRefund PaymentInitiationRefund defines a payment initiation refund
type PaymentInitiationRefund struct {
	// The ID of the refund. Like all Plaid identifiers, the `refund_id` is case sensitive.
	RefundId string `json:"refund_id"`
	Amount PaymentAmount `json:"amount"`
	// The status of the refund.  `PROCESSING`: The refund is currently being processed. The refund will automatically exit this state when processing is complete.  `INITIATED`: The refund has been successfully initiated.  `EXECUTED`: Indicates that the refund has been successfully executed.  `FAILED`: The refund has failed to be executed. This error is retryable once the root cause is resolved.
	Status string `json:"status"`
	// The date and time of the last time the `status` was updated, in IS0 8601 format
	LastStatusUpdate time.Time `json:"last_status_update"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationRefund PaymentInitiationRefund

// NewPaymentInitiationRefund instantiates a new PaymentInitiationRefund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRefund(refundId string, amount PaymentAmount, status string, lastStatusUpdate time.Time) *PaymentInitiationRefund {
	this := PaymentInitiationRefund{}
	this.RefundId = refundId
	this.Amount = amount
	this.Status = status
	this.LastStatusUpdate = lastStatusUpdate
	return &this
}

// NewPaymentInitiationRefundWithDefaults instantiates a new PaymentInitiationRefund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRefundWithDefaults() *PaymentInitiationRefund {
	this := PaymentInitiationRefund{}
	return &this
}

// GetRefundId returns the RefundId field value
func (o *PaymentInitiationRefund) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRefund) GetRefundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *PaymentInitiationRefund) SetRefundId(v string) {
	o.RefundId = v
}

// GetAmount returns the Amount field value
func (o *PaymentInitiationRefund) GetAmount() PaymentAmount {
	if o == nil {
		var ret PaymentAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRefund) GetAmountOk() (*PaymentAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentInitiationRefund) SetAmount(v PaymentAmount) {
	o.Amount = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationRefund) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRefund) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationRefund) SetStatus(v string) {
	o.Status = v
}

// GetLastStatusUpdate returns the LastStatusUpdate field value
func (o *PaymentInitiationRefund) GetLastStatusUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastStatusUpdate
}

// GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRefund) GetLastStatusUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatusUpdate, true
}

// SetLastStatusUpdate sets field value
func (o *PaymentInitiationRefund) SetLastStatusUpdate(v time.Time) {
	o.LastStatusUpdate = v
}

func (o PaymentInitiationRefund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refund_id"] = o.RefundId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["last_status_update"] = o.LastStatusUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationRefund) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationRefund := _PaymentInitiationRefund{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationRefund); err == nil {
		*o = PaymentInitiationRefund(varPaymentInitiationRefund)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "refund_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "last_status_update")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationRefund struct {
	value *PaymentInitiationRefund
	isSet bool
}

func (v NullablePaymentInitiationRefund) Get() *PaymentInitiationRefund {
	return v.value
}

func (v *NullablePaymentInitiationRefund) Set(val *PaymentInitiationRefund) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRefund) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRefund(val *PaymentInitiationRefund) *NullablePaymentInitiationRefund {
	return &NullablePaymentInitiationRefund{value: val, isSet: true}
}

func (v NullablePaymentInitiationRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


