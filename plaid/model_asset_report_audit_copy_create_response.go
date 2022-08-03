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

// AssetReportAuditCopyCreateResponse AssetReportAuditCopyCreateResponse defines the response schema for `/asset_report/audit_copy/get`
type AssetReportAuditCopyCreateResponse struct {
	// A token that can be shared with a third party auditor to allow them to obtain access to the Asset Report. This token should be stored securely.
	AuditCopyToken string `json:"audit_copy_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportAuditCopyCreateResponse AssetReportAuditCopyCreateResponse

// NewAssetReportAuditCopyCreateResponse instantiates a new AssetReportAuditCopyCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportAuditCopyCreateResponse(auditCopyToken string, requestId string) *AssetReportAuditCopyCreateResponse {
	this := AssetReportAuditCopyCreateResponse{}
	this.AuditCopyToken = auditCopyToken
	this.RequestId = requestId
	return &this
}

// NewAssetReportAuditCopyCreateResponseWithDefaults instantiates a new AssetReportAuditCopyCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportAuditCopyCreateResponseWithDefaults() *AssetReportAuditCopyCreateResponse {
	this := AssetReportAuditCopyCreateResponse{}
	return &this
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *AssetReportAuditCopyCreateResponse) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyCreateResponse) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *AssetReportAuditCopyCreateResponse) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportAuditCopyCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportAuditCopyCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportAuditCopyCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportAuditCopyCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportAuditCopyCreateResponse := _AssetReportAuditCopyCreateResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportAuditCopyCreateResponse); err == nil {
		*o = AssetReportAuditCopyCreateResponse(varAssetReportAuditCopyCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "audit_copy_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportAuditCopyCreateResponse struct {
	value *AssetReportAuditCopyCreateResponse
	isSet bool
}

func (v NullableAssetReportAuditCopyCreateResponse) Get() *AssetReportAuditCopyCreateResponse {
	return v.value
}

func (v *NullableAssetReportAuditCopyCreateResponse) Set(val *AssetReportAuditCopyCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportAuditCopyCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportAuditCopyCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportAuditCopyCreateResponse(val *AssetReportAuditCopyCreateResponse) *NullableAssetReportAuditCopyCreateResponse {
	return &NullableAssetReportAuditCopyCreateResponse{value: val, isSet: true}
}

func (v NullableAssetReportAuditCopyCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportAuditCopyCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


