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

// AssetReportRemoveResponse AssetReportRemoveResponse defines the response schema for `/asset_report/remove`
type AssetReportRemoveResponse struct {
	// `true` if the Asset Report was successfully removed.
	Removed bool `json:"removed"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportRemoveResponse AssetReportRemoveResponse

// NewAssetReportRemoveResponse instantiates a new AssetReportRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRemoveResponse(removed bool, requestId string) *AssetReportRemoveResponse {
	this := AssetReportRemoveResponse{}
	this.Removed = removed
	this.RequestId = requestId
	return &this
}

// NewAssetReportRemoveResponseWithDefaults instantiates a new AssetReportRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRemoveResponseWithDefaults() *AssetReportRemoveResponse {
	this := AssetReportRemoveResponse{}
	return &this
}

// GetRemoved returns the Removed field value
func (o *AssetReportRemoveResponse) GetRemoved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *AssetReportRemoveResponse) GetRemovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *AssetReportRemoveResponse) SetRemoved(v bool) {
	o.Removed = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["removed"] = o.Removed
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportRemoveResponse := _AssetReportRemoveResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportRemoveResponse); err == nil {
		*o = AssetReportRemoveResponse(varAssetReportRemoveResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "removed")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportRemoveResponse struct {
	value *AssetReportRemoveResponse
	isSet bool
}

func (v NullableAssetReportRemoveResponse) Get() *AssetReportRemoveResponse {
	return v.value
}

func (v *NullableAssetReportRemoveResponse) Set(val *AssetReportRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRemoveResponse(val *AssetReportRemoveResponse) *NullableAssetReportRemoveResponse {
	return &NullableAssetReportRemoveResponse{value: val, isSet: true}
}

func (v NullableAssetReportRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


