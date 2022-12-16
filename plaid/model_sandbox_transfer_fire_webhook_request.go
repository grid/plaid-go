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

// SandboxTransferFireWebhookRequest Defines the request schema for `/sandbox/transfer/fire_webhook`
type SandboxTransferFireWebhookRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The URL to which the webhook should be sent.
	Webhook string `json:"webhook"`
}

// NewSandboxTransferFireWebhookRequest instantiates a new SandboxTransferFireWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferFireWebhookRequest(webhook string) *SandboxTransferFireWebhookRequest {
	this := SandboxTransferFireWebhookRequest{}
	this.Webhook = webhook
	return &this
}

// NewSandboxTransferFireWebhookRequestWithDefaults instantiates a new SandboxTransferFireWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferFireWebhookRequestWithDefaults() *SandboxTransferFireWebhookRequest {
	this := SandboxTransferFireWebhookRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxTransferFireWebhookRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferFireWebhookRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxTransferFireWebhookRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxTransferFireWebhookRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxTransferFireWebhookRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferFireWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxTransferFireWebhookRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxTransferFireWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetWebhook returns the Webhook field value
func (o *SandboxTransferFireWebhookRequest) GetWebhook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferFireWebhookRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Webhook, true
}

// SetWebhook sets field value
func (o *SandboxTransferFireWebhookRequest) SetWebhook(v string) {
	o.Webhook = v
}

func (o SandboxTransferFireWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferFireWebhookRequest struct {
	value *SandboxTransferFireWebhookRequest
	isSet bool
}

func (v NullableSandboxTransferFireWebhookRequest) Get() *SandboxTransferFireWebhookRequest {
	return v.value
}

func (v *NullableSandboxTransferFireWebhookRequest) Set(val *SandboxTransferFireWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferFireWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferFireWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferFireWebhookRequest(val *SandboxTransferFireWebhookRequest) *NullableSandboxTransferFireWebhookRequest {
	return &NullableSandboxTransferFireWebhookRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferFireWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferFireWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


