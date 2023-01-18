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

// AssetReportRefreshRequestOptions An optional object to filter `/asset_report/refresh` results. If provided, cannot be `null`. If not specified, the `options` from the original call to `/asset_report/create` will be used.
type AssetReportRefreshRequestOptions struct {
	// Client-generated identifier, which can be used by lenders to track loan applications.
	ClientReportId NullableString `json:"client_report_id,omitempty"`
	// URL to which Plaid will send Assets webhooks, for example when the requested Asset Report is ready.
	Webhook NullableString `json:"webhook,omitempty"`
	User *AssetReportUser `json:"user,omitempty"`
}

// NewAssetReportRefreshRequestOptions instantiates a new AssetReportRefreshRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRefreshRequestOptions() *AssetReportRefreshRequestOptions {
	this := AssetReportRefreshRequestOptions{}
	return &this
}

// NewAssetReportRefreshRequestOptionsWithDefaults instantiates a new AssetReportRefreshRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRefreshRequestOptionsWithDefaults() *AssetReportRefreshRequestOptions {
	this := AssetReportRefreshRequestOptions{}
	return &this
}

// GetClientReportId returns the ClientReportId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportRefreshRequestOptions) GetClientReportId() string {
	if o == nil || o.ClientReportId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientReportId.Get()
}

// GetClientReportIdOk returns a tuple with the ClientReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportRefreshRequestOptions) GetClientReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientReportId.Get(), o.ClientReportId.IsSet()
}

// HasClientReportId returns a boolean if a field has been set.
func (o *AssetReportRefreshRequestOptions) HasClientReportId() bool {
	if o != nil && o.ClientReportId.IsSet() {
		return true
	}

	return false
}

// SetClientReportId gets a reference to the given NullableString and assigns it to the ClientReportId field.
func (o *AssetReportRefreshRequestOptions) SetClientReportId(v string) {
	o.ClientReportId.Set(&v)
}
// SetClientReportIdNil sets the value for ClientReportId to be an explicit nil
func (o *AssetReportRefreshRequestOptions) SetClientReportIdNil() {
	o.ClientReportId.Set(nil)
}

// UnsetClientReportId ensures that no value is present for ClientReportId, not even an explicit nil
func (o *AssetReportRefreshRequestOptions) UnsetClientReportId() {
	o.ClientReportId.Unset()
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportRefreshRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportRefreshRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *AssetReportRefreshRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *AssetReportRefreshRequestOptions) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *AssetReportRefreshRequestOptions) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *AssetReportRefreshRequestOptions) UnsetWebhook() {
	o.Webhook.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AssetReportRefreshRequestOptions) GetUser() AssetReportUser {
	if o == nil || o.User == nil {
		var ret AssetReportUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequestOptions) GetUserOk() (*AssetReportUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AssetReportRefreshRequestOptions) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AssetReportUser and assigns it to the User field.
func (o *AssetReportRefreshRequestOptions) SetUser(v AssetReportUser) {
	o.User = &v
}

func (o AssetReportRefreshRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientReportId.IsSet() {
		toSerialize["client_report_id"] = o.ClientReportId.Get()
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportRefreshRequestOptions struct {
	value *AssetReportRefreshRequestOptions
	isSet bool
}

func (v NullableAssetReportRefreshRequestOptions) Get() *AssetReportRefreshRequestOptions {
	return v.value
}

func (v *NullableAssetReportRefreshRequestOptions) Set(val *AssetReportRefreshRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRefreshRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRefreshRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRefreshRequestOptions(val *AssetReportRefreshRequestOptions) *NullableAssetReportRefreshRequestOptions {
	return &NullableAssetReportRefreshRequestOptions{value: val, isSet: true}
}

func (v NullableAssetReportRefreshRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRefreshRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


