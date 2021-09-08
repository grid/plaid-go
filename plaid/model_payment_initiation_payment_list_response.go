/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.31.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// PaymentInitiationPaymentListResponse PaymentInitiationPaymentListResponse defines the response schema for `/payment_initiation/payment/list`
type PaymentInitiationPaymentListResponse struct {
	// An array of payments that have been created, associated with the given `client_id`.
	Payments []PaymentInitiationPayment `json:"payments"`
	// The value that, when used as the optional `cursor` parameter to `/payment_initiation/payment/list`, will return the next unreturned payment as its first payment.
	NextCursor NullableTime `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationPaymentListResponse PaymentInitiationPaymentListResponse

// NewPaymentInitiationPaymentListResponse instantiates a new PaymentInitiationPaymentListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentListResponse(payments []PaymentInitiationPayment, nextCursor NullableTime, requestId string) *PaymentInitiationPaymentListResponse {
	this := PaymentInitiationPaymentListResponse{}
	this.Payments = payments
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationPaymentListResponseWithDefaults instantiates a new PaymentInitiationPaymentListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentListResponseWithDefaults() *PaymentInitiationPaymentListResponse {
	this := PaymentInitiationPaymentListResponse{}
	return &this
}

// GetPayments returns the Payments field value
func (o *PaymentInitiationPaymentListResponse) GetPayments() []PaymentInitiationPayment {
	if o == nil {
		var ret []PaymentInitiationPayment
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentListResponse) GetPaymentsOk() (*[]PaymentInitiationPayment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payments, true
}

// SetPayments sets field value
func (o *PaymentInitiationPaymentListResponse) SetPayments(v []PaymentInitiationPayment) {
	o.Payments = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PaymentInitiationPaymentListResponse) GetNextCursor() time.Time {
	if o == nil || o.NextCursor.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentListResponse) GetNextCursorOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaymentInitiationPaymentListResponse) SetNextCursor(v time.Time) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationPaymentListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationPaymentListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationPaymentListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payments"] = o.Payments
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationPaymentListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationPaymentListResponse := _PaymentInitiationPaymentListResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationPaymentListResponse); err == nil {
		*o = PaymentInitiationPaymentListResponse(varPaymentInitiationPaymentListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payments")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationPaymentListResponse struct {
	value *PaymentInitiationPaymentListResponse
	isSet bool
}

func (v NullablePaymentInitiationPaymentListResponse) Get() *PaymentInitiationPaymentListResponse {
	return v.value
}

func (v *NullablePaymentInitiationPaymentListResponse) Set(val *PaymentInitiationPaymentListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentListResponse(val *PaymentInitiationPaymentListResponse) *NullablePaymentInitiationPaymentListResponse {
	return &NullablePaymentInitiationPaymentListResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


