/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.205.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferRecurringCreateResponse Defines the response schema for `/transfer/recurring/create`
type TransferRecurringCreateResponse struct {
	RecurringTransfer RecurringTransfer `json:"recurring_transfer"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRecurringCreateResponse TransferRecurringCreateResponse

// NewTransferRecurringCreateResponse instantiates a new TransferRecurringCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRecurringCreateResponse(recurringTransfer RecurringTransfer, requestId string) *TransferRecurringCreateResponse {
	this := TransferRecurringCreateResponse{}
	this.RecurringTransfer = recurringTransfer
	this.RequestId = requestId
	return &this
}

// NewTransferRecurringCreateResponseWithDefaults instantiates a new TransferRecurringCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRecurringCreateResponseWithDefaults() *TransferRecurringCreateResponse {
	this := TransferRecurringCreateResponse{}
	return &this
}

// GetRecurringTransfer returns the RecurringTransfer field value
func (o *TransferRecurringCreateResponse) GetRecurringTransfer() RecurringTransfer {
	if o == nil {
		var ret RecurringTransfer
		return ret
	}

	return o.RecurringTransfer
}

// GetRecurringTransferOk returns a tuple with the RecurringTransfer field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateResponse) GetRecurringTransferOk() (*RecurringTransfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransfer, true
}

// SetRecurringTransfer sets field value
func (o *TransferRecurringCreateResponse) SetRecurringTransfer(v RecurringTransfer) {
	o.RecurringTransfer = v
}

// GetRequestId returns the RequestId field value
func (o *TransferRecurringCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRecurringCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRecurringCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recurring_transfer"] = o.RecurringTransfer
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRecurringCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRecurringCreateResponse := _TransferRecurringCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferRecurringCreateResponse); err == nil {
		*o = TransferRecurringCreateResponse(varTransferRecurringCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recurring_transfer")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRecurringCreateResponse struct {
	value *TransferRecurringCreateResponse
	isSet bool
}

func (v NullableTransferRecurringCreateResponse) Get() *TransferRecurringCreateResponse {
	return v.value
}

func (v *NullableTransferRecurringCreateResponse) Set(val *TransferRecurringCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRecurringCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRecurringCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRecurringCreateResponse(val *TransferRecurringCreateResponse) *NullableTransferRecurringCreateResponse {
	return &NullableTransferRecurringCreateResponse{value: val, isSet: true}
}

func (v NullableTransferRecurringCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRecurringCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

