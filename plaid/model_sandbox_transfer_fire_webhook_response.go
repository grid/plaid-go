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

// SandboxTransferFireWebhookResponse Defines the response schema for `/sandbox/transfer/fire_webhook`
type SandboxTransferFireWebhookResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxTransferFireWebhookResponse SandboxTransferFireWebhookResponse

// NewSandboxTransferFireWebhookResponse instantiates a new SandboxTransferFireWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferFireWebhookResponse(requestId string) *SandboxTransferFireWebhookResponse {
	this := SandboxTransferFireWebhookResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxTransferFireWebhookResponseWithDefaults instantiates a new SandboxTransferFireWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferFireWebhookResponseWithDefaults() *SandboxTransferFireWebhookResponse {
	this := SandboxTransferFireWebhookResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxTransferFireWebhookResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferFireWebhookResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxTransferFireWebhookResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxTransferFireWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxTransferFireWebhookResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxTransferFireWebhookResponse := _SandboxTransferFireWebhookResponse{}

	if err = json.Unmarshal(bytes, &varSandboxTransferFireWebhookResponse); err == nil {
		*o = SandboxTransferFireWebhookResponse(varSandboxTransferFireWebhookResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxTransferFireWebhookResponse struct {
	value *SandboxTransferFireWebhookResponse
	isSet bool
}

func (v NullableSandboxTransferFireWebhookResponse) Get() *SandboxTransferFireWebhookResponse {
	return v.value
}

func (v *NullableSandboxTransferFireWebhookResponse) Set(val *SandboxTransferFireWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferFireWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferFireWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferFireWebhookResponse(val *SandboxTransferFireWebhookResponse) *NullableSandboxTransferFireWebhookResponse {
	return &NullableSandboxTransferFireWebhookResponse{value: val, isSet: true}
}

func (v NullableSandboxTransferFireWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferFireWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


