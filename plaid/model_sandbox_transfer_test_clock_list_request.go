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
	"time"
)

// SandboxTransferTestClockListRequest Defines the request schema for `/sandbox/transfer/test_clock/list`
type SandboxTransferTestClockListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The start virtual timestamp of test clocks to return. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)
	StartVirtualTime NullableTime `json:"start_virtual_time,omitempty"`
	// The end virtual timestamp of test clocks to return. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)
	EndVirtualTime NullableTime `json:"end_virtual_time,omitempty"`
	// The maximum number of test clocks to return.
	Count NullableInt32 `json:"count,omitempty"`
	// The number of test clocks to skip before returning results.
	Offset *int32 `json:"offset,omitempty"`
}

// NewSandboxTransferTestClockListRequest instantiates a new SandboxTransferTestClockListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferTestClockListRequest() *SandboxTransferTestClockListRequest {
	this := SandboxTransferTestClockListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewSandboxTransferTestClockListRequestWithDefaults instantiates a new SandboxTransferTestClockListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferTestClockListRequestWithDefaults() *SandboxTransferTestClockListRequest {
	this := SandboxTransferTestClockListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxTransferTestClockListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxTransferTestClockListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxTransferTestClockListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxTransferTestClockListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxTransferTestClockListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxTransferTestClockListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetStartVirtualTime returns the StartVirtualTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxTransferTestClockListRequest) GetStartVirtualTime() time.Time {
	if o == nil || o.StartVirtualTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartVirtualTime.Get()
}

// GetStartVirtualTimeOk returns a tuple with the StartVirtualTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferTestClockListRequest) GetStartVirtualTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartVirtualTime.Get(), o.StartVirtualTime.IsSet()
}

// HasStartVirtualTime returns a boolean if a field has been set.
func (o *SandboxTransferTestClockListRequest) HasStartVirtualTime() bool {
	if o != nil && o.StartVirtualTime.IsSet() {
		return true
	}

	return false
}

// SetStartVirtualTime gets a reference to the given NullableTime and assigns it to the StartVirtualTime field.
func (o *SandboxTransferTestClockListRequest) SetStartVirtualTime(v time.Time) {
	o.StartVirtualTime.Set(&v)
}
// SetStartVirtualTimeNil sets the value for StartVirtualTime to be an explicit nil
func (o *SandboxTransferTestClockListRequest) SetStartVirtualTimeNil() {
	o.StartVirtualTime.Set(nil)
}

// UnsetStartVirtualTime ensures that no value is present for StartVirtualTime, not even an explicit nil
func (o *SandboxTransferTestClockListRequest) UnsetStartVirtualTime() {
	o.StartVirtualTime.Unset()
}

// GetEndVirtualTime returns the EndVirtualTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxTransferTestClockListRequest) GetEndVirtualTime() time.Time {
	if o == nil || o.EndVirtualTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndVirtualTime.Get()
}

// GetEndVirtualTimeOk returns a tuple with the EndVirtualTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferTestClockListRequest) GetEndVirtualTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndVirtualTime.Get(), o.EndVirtualTime.IsSet()
}

// HasEndVirtualTime returns a boolean if a field has been set.
func (o *SandboxTransferTestClockListRequest) HasEndVirtualTime() bool {
	if o != nil && o.EndVirtualTime.IsSet() {
		return true
	}

	return false
}

// SetEndVirtualTime gets a reference to the given NullableTime and assigns it to the EndVirtualTime field.
func (o *SandboxTransferTestClockListRequest) SetEndVirtualTime(v time.Time) {
	o.EndVirtualTime.Set(&v)
}
// SetEndVirtualTimeNil sets the value for EndVirtualTime to be an explicit nil
func (o *SandboxTransferTestClockListRequest) SetEndVirtualTimeNil() {
	o.EndVirtualTime.Set(nil)
}

// UnsetEndVirtualTime ensures that no value is present for EndVirtualTime, not even an explicit nil
func (o *SandboxTransferTestClockListRequest) UnsetEndVirtualTime() {
	o.EndVirtualTime.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxTransferTestClockListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferTestClockListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *SandboxTransferTestClockListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *SandboxTransferTestClockListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *SandboxTransferTestClockListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *SandboxTransferTestClockListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SandboxTransferTestClockListRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockListRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SandboxTransferTestClockListRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SandboxTransferTestClockListRequest) SetOffset(v int32) {
	o.Offset = &v
}

func (o SandboxTransferTestClockListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.StartVirtualTime.IsSet() {
		toSerialize["start_virtual_time"] = o.StartVirtualTime.Get()
	}
	if o.EndVirtualTime.IsSet() {
		toSerialize["end_virtual_time"] = o.EndVirtualTime.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferTestClockListRequest struct {
	value *SandboxTransferTestClockListRequest
	isSet bool
}

func (v NullableSandboxTransferTestClockListRequest) Get() *SandboxTransferTestClockListRequest {
	return v.value
}

func (v *NullableSandboxTransferTestClockListRequest) Set(val *SandboxTransferTestClockListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferTestClockListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferTestClockListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferTestClockListRequest(val *SandboxTransferTestClockListRequest) *NullableSandboxTransferTestClockListRequest {
	return &NullableSandboxTransferTestClockListRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferTestClockListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferTestClockListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


