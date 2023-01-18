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
)

// SandboxBankTransferSimulateResponse Defines the response schema for `/sandbox/bank_transfer/simulate`
type SandboxBankTransferSimulateResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxBankTransferSimulateResponse SandboxBankTransferSimulateResponse

// NewSandboxBankTransferSimulateResponse instantiates a new SandboxBankTransferSimulateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBankTransferSimulateResponse(requestId string) *SandboxBankTransferSimulateResponse {
	this := SandboxBankTransferSimulateResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxBankTransferSimulateResponseWithDefaults instantiates a new SandboxBankTransferSimulateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBankTransferSimulateResponseWithDefaults() *SandboxBankTransferSimulateResponse {
	this := SandboxBankTransferSimulateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxBankTransferSimulateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferSimulateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxBankTransferSimulateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxBankTransferSimulateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxBankTransferSimulateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxBankTransferSimulateResponse := _SandboxBankTransferSimulateResponse{}

	if err = json.Unmarshal(bytes, &varSandboxBankTransferSimulateResponse); err == nil {
		*o = SandboxBankTransferSimulateResponse(varSandboxBankTransferSimulateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxBankTransferSimulateResponse struct {
	value *SandboxBankTransferSimulateResponse
	isSet bool
}

func (v NullableSandboxBankTransferSimulateResponse) Get() *SandboxBankTransferSimulateResponse {
	return v.value
}

func (v *NullableSandboxBankTransferSimulateResponse) Set(val *SandboxBankTransferSimulateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBankTransferSimulateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBankTransferSimulateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBankTransferSimulateResponse(val *SandboxBankTransferSimulateResponse) *NullableSandboxBankTransferSimulateResponse {
	return &NullableSandboxBankTransferSimulateResponse{value: val, isSet: true}
}

func (v NullableSandboxBankTransferSimulateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBankTransferSimulateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


