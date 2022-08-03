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

// WalletCreateRequest WalletCreateRequest defines the request schema for `/wallet/create`
type WalletCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	IsoCurrencyCode WalletISOCurrencyCode `json:"iso_currency_code"`
}

// NewWalletCreateRequest instantiates a new WalletCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletCreateRequest(isoCurrencyCode WalletISOCurrencyCode) *WalletCreateRequest {
	this := WalletCreateRequest{}
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewWalletCreateRequestWithDefaults instantiates a new WalletCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletCreateRequestWithDefaults() *WalletCreateRequest {
	this := WalletCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WalletCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WalletCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WalletCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WalletCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WalletCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WalletCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *WalletCreateRequest) GetIsoCurrencyCode() WalletISOCurrencyCode {
	if o == nil {
		var ret WalletISOCurrencyCode
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *WalletCreateRequest) GetIsoCurrencyCodeOk() (*WalletISOCurrencyCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *WalletCreateRequest) SetIsoCurrencyCode(v WalletISOCurrencyCode) {
	o.IsoCurrencyCode = v
}

func (o WalletCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableWalletCreateRequest struct {
	value *WalletCreateRequest
	isSet bool
}

func (v NullableWalletCreateRequest) Get() *WalletCreateRequest {
	return v.value
}

func (v *NullableWalletCreateRequest) Set(val *WalletCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletCreateRequest(val *WalletCreateRequest) *NullableWalletCreateRequest {
	return &NullableWalletCreateRequest{value: val, isSet: true}
}

func (v NullableWalletCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


