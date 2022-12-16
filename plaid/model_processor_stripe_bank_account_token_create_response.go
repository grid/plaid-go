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

// ProcessorStripeBankAccountTokenCreateResponse ProcessorStripeBankAccountTokenCreateResponse defines the response schema for `/processor/stripe/bank_account/create`
type ProcessorStripeBankAccountTokenCreateResponse struct {
	// A token that can be sent to Stripe for use in making API calls to Plaid
	StripeBankAccountToken string `json:"stripe_bank_account_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorStripeBankAccountTokenCreateResponse ProcessorStripeBankAccountTokenCreateResponse

// NewProcessorStripeBankAccountTokenCreateResponse instantiates a new ProcessorStripeBankAccountTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorStripeBankAccountTokenCreateResponse(stripeBankAccountToken string, requestId string) *ProcessorStripeBankAccountTokenCreateResponse {
	this := ProcessorStripeBankAccountTokenCreateResponse{}
	this.StripeBankAccountToken = stripeBankAccountToken
	this.RequestId = requestId
	return &this
}

// NewProcessorStripeBankAccountTokenCreateResponseWithDefaults instantiates a new ProcessorStripeBankAccountTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorStripeBankAccountTokenCreateResponseWithDefaults() *ProcessorStripeBankAccountTokenCreateResponse {
	this := ProcessorStripeBankAccountTokenCreateResponse{}
	return &this
}

// GetStripeBankAccountToken returns the StripeBankAccountToken field value
func (o *ProcessorStripeBankAccountTokenCreateResponse) GetStripeBankAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StripeBankAccountToken
}

// GetStripeBankAccountTokenOk returns a tuple with the StripeBankAccountToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorStripeBankAccountTokenCreateResponse) GetStripeBankAccountTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StripeBankAccountToken, true
}

// SetStripeBankAccountToken sets field value
func (o *ProcessorStripeBankAccountTokenCreateResponse) SetStripeBankAccountToken(v string) {
	o.StripeBankAccountToken = v
}

// GetRequestId returns the RequestId field value
func (o *ProcessorStripeBankAccountTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorStripeBankAccountTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorStripeBankAccountTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorStripeBankAccountTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stripe_bank_account_token"] = o.StripeBankAccountToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorStripeBankAccountTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorStripeBankAccountTokenCreateResponse := _ProcessorStripeBankAccountTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varProcessorStripeBankAccountTokenCreateResponse); err == nil {
		*o = ProcessorStripeBankAccountTokenCreateResponse(varProcessorStripeBankAccountTokenCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "stripe_bank_account_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorStripeBankAccountTokenCreateResponse struct {
	value *ProcessorStripeBankAccountTokenCreateResponse
	isSet bool
}

func (v NullableProcessorStripeBankAccountTokenCreateResponse) Get() *ProcessorStripeBankAccountTokenCreateResponse {
	return v.value
}

func (v *NullableProcessorStripeBankAccountTokenCreateResponse) Set(val *ProcessorStripeBankAccountTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorStripeBankAccountTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorStripeBankAccountTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorStripeBankAccountTokenCreateResponse(val *ProcessorStripeBankAccountTokenCreateResponse) *NullableProcessorStripeBankAccountTokenCreateResponse {
	return &NullableProcessorStripeBankAccountTokenCreateResponse{value: val, isSet: true}
}

func (v NullableProcessorStripeBankAccountTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorStripeBankAccountTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


