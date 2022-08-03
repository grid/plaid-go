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
)

// TransferSweepGetResponse Defines the response schema for `/transfer/sweep/get`
type TransferSweepGetResponse struct {
	Sweep TransferSweep `json:"sweep"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferSweepGetResponse TransferSweepGetResponse

// NewTransferSweepGetResponse instantiates a new TransferSweepGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferSweepGetResponse(sweep TransferSweep, requestId string) *TransferSweepGetResponse {
	this := TransferSweepGetResponse{}
	this.Sweep = sweep
	this.RequestId = requestId
	return &this
}

// NewTransferSweepGetResponseWithDefaults instantiates a new TransferSweepGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferSweepGetResponseWithDefaults() *TransferSweepGetResponse {
	this := TransferSweepGetResponse{}
	return &this
}

// GetSweep returns the Sweep field value
func (o *TransferSweepGetResponse) GetSweep() TransferSweep {
	if o == nil {
		var ret TransferSweep
		return ret
	}

	return o.Sweep
}

// GetSweepOk returns a tuple with the Sweep field value
// and a boolean to check if the value has been set.
func (o *TransferSweepGetResponse) GetSweepOk() (*TransferSweep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sweep, true
}

// SetSweep sets field value
func (o *TransferSweepGetResponse) SetSweep(v TransferSweep) {
	o.Sweep = v
}

// GetRequestId returns the RequestId field value
func (o *TransferSweepGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferSweepGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferSweepGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferSweepGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sweep"] = o.Sweep
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferSweepGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferSweepGetResponse := _TransferSweepGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferSweepGetResponse); err == nil {
		*o = TransferSweepGetResponse(varTransferSweepGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sweep")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferSweepGetResponse struct {
	value *TransferSweepGetResponse
	isSet bool
}

func (v NullableTransferSweepGetResponse) Get() *TransferSweepGetResponse {
	return v.value
}

func (v *NullableTransferSweepGetResponse) Set(val *TransferSweepGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferSweepGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferSweepGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferSweepGetResponse(val *TransferSweepGetResponse) *NullableTransferSweepGetResponse {
	return &NullableTransferSweepGetResponse{value: val, isSet: true}
}

func (v NullableTransferSweepGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferSweepGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


