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
)

// TransferRecurringGetResponse Defines the response schema for `/transfer/recurring/get`
type TransferRecurringGetResponse struct {
	RecurringTransfer RecurringTransfer `json:"recurring_transfer"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRecurringGetResponse TransferRecurringGetResponse

// NewTransferRecurringGetResponse instantiates a new TransferRecurringGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRecurringGetResponse(recurringTransfer RecurringTransfer, requestId string) *TransferRecurringGetResponse {
	this := TransferRecurringGetResponse{}
	this.RecurringTransfer = recurringTransfer
	this.RequestId = requestId
	return &this
}

// NewTransferRecurringGetResponseWithDefaults instantiates a new TransferRecurringGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRecurringGetResponseWithDefaults() *TransferRecurringGetResponse {
	this := TransferRecurringGetResponse{}
	return &this
}

// GetRecurringTransfer returns the RecurringTransfer field value
func (o *TransferRecurringGetResponse) GetRecurringTransfer() RecurringTransfer {
	if o == nil {
		var ret RecurringTransfer
		return ret
	}

	return o.RecurringTransfer
}

// GetRecurringTransferOk returns a tuple with the RecurringTransfer field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringGetResponse) GetRecurringTransferOk() (*RecurringTransfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransfer, true
}

// SetRecurringTransfer sets field value
func (o *TransferRecurringGetResponse) SetRecurringTransfer(v RecurringTransfer) {
	o.RecurringTransfer = v
}

// GetRequestId returns the RequestId field value
func (o *TransferRecurringGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRecurringGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRecurringGetResponse) MarshalJSON() ([]byte, error) {
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

func (o *TransferRecurringGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRecurringGetResponse := _TransferRecurringGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferRecurringGetResponse); err == nil {
		*o = TransferRecurringGetResponse(varTransferRecurringGetResponse)
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

type NullableTransferRecurringGetResponse struct {
	value *TransferRecurringGetResponse
	isSet bool
}

func (v NullableTransferRecurringGetResponse) Get() *TransferRecurringGetResponse {
	return v.value
}

func (v *NullableTransferRecurringGetResponse) Set(val *TransferRecurringGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRecurringGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRecurringGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRecurringGetResponse(val *TransferRecurringGetResponse) *NullableTransferRecurringGetResponse {
	return &NullableTransferRecurringGetResponse{value: val, isSet: true}
}

func (v NullableTransferRecurringGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRecurringGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


