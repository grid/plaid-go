/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AssetReportRemoveRequest AssetReportRemoveRequest defines the request schema for `/asset_report/remove`
type AssetReportRemoveRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
}

// NewAssetReportRemoveRequest instantiates a new AssetReportRemoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRemoveRequest(assetReportToken string) *AssetReportRemoveRequest {
	this := AssetReportRemoveRequest{}
	this.AssetReportToken = assetReportToken
	return &this
}

// NewAssetReportRemoveRequestWithDefaults instantiates a new AssetReportRemoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRemoveRequestWithDefaults() *AssetReportRemoveRequest {
	this := AssetReportRemoveRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportRemoveRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRemoveRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportRemoveRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportRemoveRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportRemoveRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRemoveRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportRemoveRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportRemoveRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportRemoveRequest) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRemoveRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportRemoveRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

func (o AssetReportRemoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportRemoveRequest struct {
	value *AssetReportRemoveRequest
	isSet bool
}

func (v NullableAssetReportRemoveRequest) Get() *AssetReportRemoveRequest {
	return v.value
}

func (v *NullableAssetReportRemoveRequest) Set(val *AssetReportRemoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRemoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRemoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRemoveRequest(val *AssetReportRemoveRequest) *NullableAssetReportRemoveRequest {
	return &NullableAssetReportRemoveRequest{value: val, isSet: true}
}

func (v NullableAssetReportRemoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRemoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


