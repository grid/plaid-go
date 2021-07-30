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

// WebhookVerificationKeyGetResponse WebhookVerificationKeyGetResponse defines the response schema for `/webhook_verification_key/get`
type WebhookVerificationKeyGetResponse struct {
	Key JWKPublicKey `json:"key"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WebhookVerificationKeyGetResponse WebhookVerificationKeyGetResponse

// NewWebhookVerificationKeyGetResponse instantiates a new WebhookVerificationKeyGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookVerificationKeyGetResponse(key JWKPublicKey, requestId string) *WebhookVerificationKeyGetResponse {
	this := WebhookVerificationKeyGetResponse{}
	this.Key = key
	this.RequestId = requestId
	return &this
}

// NewWebhookVerificationKeyGetResponseWithDefaults instantiates a new WebhookVerificationKeyGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookVerificationKeyGetResponseWithDefaults() *WebhookVerificationKeyGetResponse {
	this := WebhookVerificationKeyGetResponse{}
	return &this
}

// GetKey returns the Key field value
func (o *WebhookVerificationKeyGetResponse) GetKey() JWKPublicKey {
	if o == nil {
		var ret JWKPublicKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *WebhookVerificationKeyGetResponse) GetKeyOk() (*JWKPublicKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *WebhookVerificationKeyGetResponse) SetKey(v JWKPublicKey) {
	o.Key = v
}

// GetRequestId returns the RequestId field value
func (o *WebhookVerificationKeyGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WebhookVerificationKeyGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WebhookVerificationKeyGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WebhookVerificationKeyGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebhookVerificationKeyGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWebhookVerificationKeyGetResponse := _WebhookVerificationKeyGetResponse{}

	if err = json.Unmarshal(bytes, &varWebhookVerificationKeyGetResponse); err == nil {
		*o = WebhookVerificationKeyGetResponse(varWebhookVerificationKeyGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookVerificationKeyGetResponse struct {
	value *WebhookVerificationKeyGetResponse
	isSet bool
}

func (v NullableWebhookVerificationKeyGetResponse) Get() *WebhookVerificationKeyGetResponse {
	return v.value
}

func (v *NullableWebhookVerificationKeyGetResponse) Set(val *WebhookVerificationKeyGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookVerificationKeyGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookVerificationKeyGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookVerificationKeyGetResponse(val *WebhookVerificationKeyGetResponse) *NullableWebhookVerificationKeyGetResponse {
	return &NullableWebhookVerificationKeyGetResponse{value: val, isSet: true}
}

func (v NullableWebhookVerificationKeyGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookVerificationKeyGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


