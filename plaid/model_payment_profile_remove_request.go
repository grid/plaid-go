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

// PaymentProfileRemoveRequest PaymentProfileRemoveRequest defines the request schema for `/payment_profile/remove`
type PaymentProfileRemoveRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A payment profile token associated with the Payment Profile data that is being requested.
	PaymentProfileToken string `json:"payment_profile_token"`
}

// NewPaymentProfileRemoveRequest instantiates a new PaymentProfileRemoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProfileRemoveRequest(paymentProfileToken string) *PaymentProfileRemoveRequest {
	this := PaymentProfileRemoveRequest{}
	this.PaymentProfileToken = paymentProfileToken
	return &this
}

// NewPaymentProfileRemoveRequestWithDefaults instantiates a new PaymentProfileRemoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProfileRemoveRequestWithDefaults() *PaymentProfileRemoveRequest {
	this := PaymentProfileRemoveRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentProfileRemoveRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentProfileRemoveRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentProfileRemoveRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentProfileRemoveRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentProfileRemoveRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentProfileRemoveRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentProfileRemoveRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentProfileRemoveRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPaymentProfileToken returns the PaymentProfileToken field value
func (o *PaymentProfileRemoveRequest) GetPaymentProfileToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentProfileToken
}

// GetPaymentProfileTokenOk returns a tuple with the PaymentProfileToken field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileRemoveRequest) GetPaymentProfileTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentProfileToken, true
}

// SetPaymentProfileToken sets field value
func (o *PaymentProfileRemoveRequest) SetPaymentProfileToken(v string) {
	o.PaymentProfileToken = v
}

func (o PaymentProfileRemoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["payment_profile_token"] = o.PaymentProfileToken
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentProfileRemoveRequest struct {
	value *PaymentProfileRemoveRequest
	isSet bool
}

func (v NullablePaymentProfileRemoveRequest) Get() *PaymentProfileRemoveRequest {
	return v.value
}

func (v *NullablePaymentProfileRemoveRequest) Set(val *PaymentProfileRemoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileRemoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileRemoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileRemoveRequest(val *PaymentProfileRemoveRequest) *NullablePaymentProfileRemoveRequest {
	return &NullablePaymentProfileRemoveRequest{value: val, isSet: true}
}

func (v NullablePaymentProfileRemoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileRemoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


