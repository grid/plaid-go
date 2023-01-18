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

// CreditRelayRefreshRequest CreditRelayRefreshRequest defines the request schema for `/credit/relay/refresh`
type CreditRelayRefreshRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `relay_token` granting access to the report you would like to refresh.
	RelayToken string `json:"relay_token"`
	ReportType ReportType `json:"report_type"`
	// The URL registered to receive webhooks when the report of a relay token has been refreshed.
	Webhook NullableString `json:"webhook,omitempty"`
}

// NewCreditRelayRefreshRequest instantiates a new CreditRelayRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditRelayRefreshRequest(relayToken string, reportType ReportType) *CreditRelayRefreshRequest {
	this := CreditRelayRefreshRequest{}
	this.RelayToken = relayToken
	this.ReportType = reportType
	return &this
}

// NewCreditRelayRefreshRequestWithDefaults instantiates a new CreditRelayRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditRelayRefreshRequestWithDefaults() *CreditRelayRefreshRequest {
	this := CreditRelayRefreshRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditRelayRefreshRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditRelayRefreshRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditRelayRefreshRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditRelayRefreshRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditRelayRefreshRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditRelayRefreshRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRelayToken returns the RelayToken field value
func (o *CreditRelayRefreshRequest) GetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelayToken
}

// GetRelayTokenOk returns a tuple with the RelayToken field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshRequest) GetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelayToken, true
}

// SetRelayToken sets field value
func (o *CreditRelayRefreshRequest) SetRelayToken(v string) {
	o.RelayToken = v
}

// GetReportType returns the ReportType field value
func (o *CreditRelayRefreshRequest) GetReportType() ReportType {
	if o == nil {
		var ret ReportType
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *CreditRelayRefreshRequest) GetReportTypeOk() (*ReportType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *CreditRelayRefreshRequest) SetReportType(v ReportType) {
	o.ReportType = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditRelayRefreshRequest) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditRelayRefreshRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *CreditRelayRefreshRequest) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *CreditRelayRefreshRequest) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *CreditRelayRefreshRequest) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *CreditRelayRefreshRequest) UnsetWebhook() {
	o.Webhook.Unset()
}

func (o CreditRelayRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["relay_token"] = o.RelayToken
	}
	if true {
		toSerialize["report_type"] = o.ReportType
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditRelayRefreshRequest struct {
	value *CreditRelayRefreshRequest
	isSet bool
}

func (v NullableCreditRelayRefreshRequest) Get() *CreditRelayRefreshRequest {
	return v.value
}

func (v *NullableCreditRelayRefreshRequest) Set(val *CreditRelayRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditRelayRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditRelayRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditRelayRefreshRequest(val *CreditRelayRefreshRequest) *NullableCreditRelayRefreshRequest {
	return &NullableCreditRelayRefreshRequest{value: val, isSet: true}
}

func (v NullableCreditRelayRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditRelayRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


