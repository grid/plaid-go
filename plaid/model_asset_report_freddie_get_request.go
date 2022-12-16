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
)

// AssetReportFreddieGetRequest AssetReportFreddieGetResponse defines the request schema for `credit/asset_report/freddie_mac/get`
type AssetReportFreddieGetRequest struct {
	// A token that can be shared with a third party auditor to allow them to obtain access to the Asset Report. This token should be stored securely.
	AuditCopyToken string `json:"audit_copy_token"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportFreddieGetRequest AssetReportFreddieGetRequest

// NewAssetReportFreddieGetRequest instantiates a new AssetReportFreddieGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportFreddieGetRequest(auditCopyToken string) *AssetReportFreddieGetRequest {
	this := AssetReportFreddieGetRequest{}
	this.AuditCopyToken = auditCopyToken
	return &this
}

// NewAssetReportFreddieGetRequestWithDefaults instantiates a new AssetReportFreddieGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportFreddieGetRequestWithDefaults() *AssetReportFreddieGetRequest {
	this := AssetReportFreddieGetRequest{}
	return &this
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *AssetReportFreddieGetRequest) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddieGetRequest) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *AssetReportFreddieGetRequest) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportFreddieGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportFreddieGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportFreddieGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportFreddieGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportFreddieGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportFreddieGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportFreddieGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportFreddieGetRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o AssetReportFreddieGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportFreddieGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportFreddieGetRequest := _AssetReportFreddieGetRequest{}

	if err = json.Unmarshal(bytes, &varAssetReportFreddieGetRequest); err == nil {
		*o = AssetReportFreddieGetRequest(varAssetReportFreddieGetRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "audit_copy_token")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportFreddieGetRequest struct {
	value *AssetReportFreddieGetRequest
	isSet bool
}

func (v NullableAssetReportFreddieGetRequest) Get() *AssetReportFreddieGetRequest {
	return v.value
}

func (v *NullableAssetReportFreddieGetRequest) Set(val *AssetReportFreddieGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportFreddieGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportFreddieGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportFreddieGetRequest(val *AssetReportFreddieGetRequest) *NullableAssetReportFreddieGetRequest {
	return &NullableAssetReportFreddieGetRequest{value: val, isSet: true}
}

func (v NullableAssetReportFreddieGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportFreddieGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


