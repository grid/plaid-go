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
)

// PaymentInitiationRecipientListResponse PaymentInitiationRecipientListResponse defines the response schema for `/payment_initiation/recipient/list`
type PaymentInitiationRecipientListResponse struct {
	// An array of payment recipients created for Payment Initiation
	Recipients []PaymentInitiationRecipient `json:"recipients"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationRecipientListResponse PaymentInitiationRecipientListResponse

// NewPaymentInitiationRecipientListResponse instantiates a new PaymentInitiationRecipientListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientListResponse(recipients []PaymentInitiationRecipient, requestId string) *PaymentInitiationRecipientListResponse {
	this := PaymentInitiationRecipientListResponse{}
	this.Recipients = recipients
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationRecipientListResponseWithDefaults instantiates a new PaymentInitiationRecipientListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientListResponseWithDefaults() *PaymentInitiationRecipientListResponse {
	this := PaymentInitiationRecipientListResponse{}
	return &this
}

// GetRecipients returns the Recipients field value
func (o *PaymentInitiationRecipientListResponse) GetRecipients() []PaymentInitiationRecipient {
	if o == nil {
		var ret []PaymentInitiationRecipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientListResponse) GetRecipientsOk() (*[]PaymentInitiationRecipient, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *PaymentInitiationRecipientListResponse) SetRecipients(v []PaymentInitiationRecipient) {
	o.Recipients = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationRecipientListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationRecipientListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationRecipientListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationRecipientListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationRecipientListResponse := _PaymentInitiationRecipientListResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationRecipientListResponse); err == nil {
		*o = PaymentInitiationRecipientListResponse(varPaymentInitiationRecipientListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationRecipientListResponse struct {
	value *PaymentInitiationRecipientListResponse
	isSet bool
}

func (v NullablePaymentInitiationRecipientListResponse) Get() *PaymentInitiationRecipientListResponse {
	return v.value
}

func (v *NullablePaymentInitiationRecipientListResponse) Set(val *PaymentInitiationRecipientListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientListResponse(val *PaymentInitiationRecipientListResponse) *NullablePaymentInitiationRecipientListResponse {
	return &NullablePaymentInitiationRecipientListResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


