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

// SandboxTransferTestClockGetRequest Defines the request schema for `/sandbox/transfer/test_clock/get`
type SandboxTransferTestClockGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
	// Plaid’s unique identifier for a test clock.
	TestClockId NullableString `json:"test_clock_id"`
}

// NewSandboxTransferTestClockGetRequest instantiates a new SandboxTransferTestClockGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferTestClockGetRequest(clientId string, secret string, testClockId NullableString) *SandboxTransferTestClockGetRequest {
	this := SandboxTransferTestClockGetRequest{}
	this.ClientId = clientId
	this.Secret = secret
	this.TestClockId = testClockId
	return &this
}

// NewSandboxTransferTestClockGetRequestWithDefaults instantiates a new SandboxTransferTestClockGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferTestClockGetRequestWithDefaults() *SandboxTransferTestClockGetRequest {
	this := SandboxTransferTestClockGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *SandboxTransferTestClockGetRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *SandboxTransferTestClockGetRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *SandboxTransferTestClockGetRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockGetRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SandboxTransferTestClockGetRequest) SetSecret(v string) {
	o.Secret = v
}

// GetTestClockId returns the TestClockId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SandboxTransferTestClockGetRequest) GetTestClockId() string {
	if o == nil || o.TestClockId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TestClockId.Get()
}

// GetTestClockIdOk returns a tuple with the TestClockId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferTestClockGetRequest) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TestClockId.Get(), o.TestClockId.IsSet()
}

// SetTestClockId sets field value
func (o *SandboxTransferTestClockGetRequest) SetTestClockId(v string) {
	o.TestClockId.Set(&v)
}

func (o SandboxTransferTestClockGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["test_clock_id"] = o.TestClockId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferTestClockGetRequest struct {
	value *SandboxTransferTestClockGetRequest
	isSet bool
}

func (v NullableSandboxTransferTestClockGetRequest) Get() *SandboxTransferTestClockGetRequest {
	return v.value
}

func (v *NullableSandboxTransferTestClockGetRequest) Set(val *SandboxTransferTestClockGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferTestClockGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferTestClockGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferTestClockGetRequest(val *SandboxTransferTestClockGetRequest) *NullableSandboxTransferTestClockGetRequest {
	return &NullableSandboxTransferTestClockGetRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferTestClockGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferTestClockGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

