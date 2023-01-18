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
	"time"
)

// SandboxTransferTestClockCreateRequest Defines the request schema for `/sandbox/transfer/test_clock/create`
type SandboxTransferTestClockCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The virtual timestamp on the test clock. This will be of the form `2006-01-02T15:04:05Z`.
	VirtualTime time.Time `json:"virtual_time"`
}

// NewSandboxTransferTestClockCreateRequest instantiates a new SandboxTransferTestClockCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferTestClockCreateRequest(virtualTime time.Time) *SandboxTransferTestClockCreateRequest {
	this := SandboxTransferTestClockCreateRequest{}
	this.VirtualTime = virtualTime
	return &this
}

// NewSandboxTransferTestClockCreateRequestWithDefaults instantiates a new SandboxTransferTestClockCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferTestClockCreateRequestWithDefaults() *SandboxTransferTestClockCreateRequest {
	this := SandboxTransferTestClockCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxTransferTestClockCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxTransferTestClockCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxTransferTestClockCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxTransferTestClockCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxTransferTestClockCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxTransferTestClockCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetVirtualTime returns the VirtualTime field value
func (o *SandboxTransferTestClockCreateRequest) GetVirtualTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.VirtualTime
}

// GetVirtualTimeOk returns a tuple with the VirtualTime field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockCreateRequest) GetVirtualTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualTime, true
}

// SetVirtualTime sets field value
func (o *SandboxTransferTestClockCreateRequest) SetVirtualTime(v time.Time) {
	o.VirtualTime = v
}

func (o SandboxTransferTestClockCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["virtual_time"] = o.VirtualTime
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferTestClockCreateRequest struct {
	value *SandboxTransferTestClockCreateRequest
	isSet bool
}

func (v NullableSandboxTransferTestClockCreateRequest) Get() *SandboxTransferTestClockCreateRequest {
	return v.value
}

func (v *NullableSandboxTransferTestClockCreateRequest) Set(val *SandboxTransferTestClockCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferTestClockCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferTestClockCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferTestClockCreateRequest(val *SandboxTransferTestClockCreateRequest) *NullableSandboxTransferTestClockCreateRequest {
	return &NullableSandboxTransferTestClockCreateRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferTestClockCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferTestClockCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


