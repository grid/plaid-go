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

// BankTransferSweepGetRequest Defines the request schema for `/bank_transfer/sweep/get`
type BankTransferSweepGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Identifier of the sweep.
	SweepId string `json:"sweep_id"`
}

// NewBankTransferSweepGetRequest instantiates a new BankTransferSweepGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferSweepGetRequest(sweepId string) *BankTransferSweepGetRequest {
	this := BankTransferSweepGetRequest{}
	this.SweepId = sweepId
	return &this
}

// NewBankTransferSweepGetRequestWithDefaults instantiates a new BankTransferSweepGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferSweepGetRequestWithDefaults() *BankTransferSweepGetRequest {
	this := BankTransferSweepGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransferSweepGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferSweepGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransferSweepGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransferSweepGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BankTransferSweepGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferSweepGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BankTransferSweepGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BankTransferSweepGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetSweepId returns the SweepId field value
func (o *BankTransferSweepGetRequest) GetSweepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SweepId
}

// GetSweepIdOk returns a tuple with the SweepId field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepGetRequest) GetSweepIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SweepId, true
}

// SetSweepId sets field value
func (o *BankTransferSweepGetRequest) SetSweepId(v string) {
	o.SweepId = v
}

func (o BankTransferSweepGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["sweep_id"] = o.SweepId
	}
	return json.Marshal(toSerialize)
}

type NullableBankTransferSweepGetRequest struct {
	value *BankTransferSweepGetRequest
	isSet bool
}

func (v NullableBankTransferSweepGetRequest) Get() *BankTransferSweepGetRequest {
	return v.value
}

func (v *NullableBankTransferSweepGetRequest) Set(val *BankTransferSweepGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferSweepGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferSweepGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferSweepGetRequest(val *BankTransferSweepGetRequest) *NullableBankTransferSweepGetRequest {
	return &NullableBankTransferSweepGetRequest{value: val, isSet: true}
}

func (v NullableBankTransferSweepGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferSweepGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


