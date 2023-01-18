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

// WalletListResponse WalletListResponse defines the response schema for `/wallet/list`
type WalletListResponse struct {
	// An array of e-wallets
	Wallets []Wallet `json:"wallets"`
	// Cursor used for fetching e-wallets created before the latest e-wallet provided in this response
	NextCursor *string `json:"next_cursor,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletListResponse WalletListResponse

// NewWalletListResponse instantiates a new WalletListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletListResponse(wallets []Wallet, requestId string) *WalletListResponse {
	this := WalletListResponse{}
	this.Wallets = wallets
	this.RequestId = requestId
	return &this
}

// NewWalletListResponseWithDefaults instantiates a new WalletListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletListResponseWithDefaults() *WalletListResponse {
	this := WalletListResponse{}
	return &this
}

// GetWallets returns the Wallets field value
func (o *WalletListResponse) GetWallets() []Wallet {
	if o == nil {
		var ret []Wallet
		return ret
	}

	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value
// and a boolean to check if the value has been set.
func (o *WalletListResponse) GetWalletsOk() (*[]Wallet, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Wallets, true
}

// SetWallets sets field value
func (o *WalletListResponse) SetWallets(v []Wallet) {
	o.Wallets = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *WalletListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *WalletListResponse) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *WalletListResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetRequestId returns the RequestId field value
func (o *WalletListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallets"] = o.Wallets
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletListResponse := _WalletListResponse{}

	if err = json.Unmarshal(bytes, &varWalletListResponse); err == nil {
		*o = WalletListResponse(varWalletListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "wallets")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletListResponse struct {
	value *WalletListResponse
	isSet bool
}

func (v NullableWalletListResponse) Get() *WalletListResponse {
	return v.value
}

func (v *NullableWalletListResponse) Set(val *WalletListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletListResponse(val *WalletListResponse) *NullableWalletListResponse {
	return &NullableWalletListResponse{value: val, isSet: true}
}

func (v NullableWalletListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


