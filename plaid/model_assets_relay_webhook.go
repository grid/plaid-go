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

// AssetsRelayWebhook Fired when the Secondary Client successfully retrieves an Asset Report by calling `asset_report/relay/get`.
type AssetsRelayWebhook struct {
	// `ASSETS`
	WebhookType string `json:"webhook_type"`
	// `RELAY_EVENT`
	WebhookCode string `json:"webhook_code"`
	RelayEvent RelayEvent `json:"relay_event"`
	// The id of the client with whom the Asset Report is being shared.
	SecondaryClientId string `json:"secondary_client_id"`
	// The `asset_relay_token` that was created by calling `/asset_report/relay/create.
	AssetRelayToken string `json:"asset_relay_token"`
	// The `asset_report_id` that can be provided to `/asset_report/relay/get` to retrieve the Asset Report.
	AssetReportId string `json:"asset_report_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetsRelayWebhook AssetsRelayWebhook

// NewAssetsRelayWebhook instantiates a new AssetsRelayWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsRelayWebhook(webhookType string, webhookCode string, relayEvent RelayEvent, secondaryClientId string, assetRelayToken string, assetReportId string) *AssetsRelayWebhook {
	this := AssetsRelayWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.RelayEvent = relayEvent
	this.SecondaryClientId = secondaryClientId
	this.AssetRelayToken = assetRelayToken
	this.AssetReportId = assetReportId
	return &this
}

// NewAssetsRelayWebhookWithDefaults instantiates a new AssetsRelayWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsRelayWebhookWithDefaults() *AssetsRelayWebhook {
	this := AssetsRelayWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *AssetsRelayWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *AssetsRelayWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *AssetsRelayWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *AssetsRelayWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *AssetsRelayWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *AssetsRelayWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetRelayEvent returns the RelayEvent field value
func (o *AssetsRelayWebhook) GetRelayEvent() RelayEvent {
	if o == nil {
		var ret RelayEvent
		return ret
	}

	return o.RelayEvent
}

// GetRelayEventOk returns a tuple with the RelayEvent field value
// and a boolean to check if the value has been set.
func (o *AssetsRelayWebhook) GetRelayEventOk() (*RelayEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelayEvent, true
}

// SetRelayEvent sets field value
func (o *AssetsRelayWebhook) SetRelayEvent(v RelayEvent) {
	o.RelayEvent = v
}

// GetSecondaryClientId returns the SecondaryClientId field value
func (o *AssetsRelayWebhook) GetSecondaryClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryClientId
}

// GetSecondaryClientIdOk returns a tuple with the SecondaryClientId field value
// and a boolean to check if the value has been set.
func (o *AssetsRelayWebhook) GetSecondaryClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecondaryClientId, true
}

// SetSecondaryClientId sets field value
func (o *AssetsRelayWebhook) SetSecondaryClientId(v string) {
	o.SecondaryClientId = v
}

// GetAssetRelayToken returns the AssetRelayToken field value
func (o *AssetsRelayWebhook) GetAssetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetRelayToken
}

// GetAssetRelayTokenOk returns a tuple with the AssetRelayToken field value
// and a boolean to check if the value has been set.
func (o *AssetsRelayWebhook) GetAssetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetRelayToken, true
}

// SetAssetRelayToken sets field value
func (o *AssetsRelayWebhook) SetAssetRelayToken(v string) {
	o.AssetRelayToken = v
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetsRelayWebhook) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetsRelayWebhook) GetAssetReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetsRelayWebhook) SetAssetReportId(v string) {
	o.AssetReportId = v
}

func (o AssetsRelayWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["relay_event"] = o.RelayEvent
	}
	if true {
		toSerialize["secondary_client_id"] = o.SecondaryClientId
	}
	if true {
		toSerialize["asset_relay_token"] = o.AssetRelayToken
	}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetsRelayWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varAssetsRelayWebhook := _AssetsRelayWebhook{}

	if err = json.Unmarshal(bytes, &varAssetsRelayWebhook); err == nil {
		*o = AssetsRelayWebhook(varAssetsRelayWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "relay_event")
		delete(additionalProperties, "secondary_client_id")
		delete(additionalProperties, "asset_relay_token")
		delete(additionalProperties, "asset_report_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetsRelayWebhook struct {
	value *AssetsRelayWebhook
	isSet bool
}

func (v NullableAssetsRelayWebhook) Get() *AssetsRelayWebhook {
	return v.value
}

func (v *NullableAssetsRelayWebhook) Set(val *AssetsRelayWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsRelayWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsRelayWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsRelayWebhook(val *AssetsRelayWebhook) *NullableAssetsRelayWebhook {
	return &NullableAssetsRelayWebhook{value: val, isSet: true}
}

func (v NullableAssetsRelayWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsRelayWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


