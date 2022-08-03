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

// PaymentInitiationPaymentTokenCreateRequest PaymentInitiationPaymentTokenCreateRequest defines the request schema for `/payment_initiation/payment/token/create`
type PaymentInitiationPaymentTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `payment_id` returned from `/payment_initiation/payment/create`.
	PaymentId string `json:"payment_id"`
}

// NewPaymentInitiationPaymentTokenCreateRequest instantiates a new PaymentInitiationPaymentTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentTokenCreateRequest(paymentId string) *PaymentInitiationPaymentTokenCreateRequest {
	this := PaymentInitiationPaymentTokenCreateRequest{}
	this.PaymentId = paymentId
	return &this
}

// NewPaymentInitiationPaymentTokenCreateRequestWithDefaults instantiates a new PaymentInitiationPaymentTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentTokenCreateRequestWithDefaults() *PaymentInitiationPaymentTokenCreateRequest {
	this := PaymentInitiationPaymentTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationPaymentTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationPaymentTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationPaymentTokenCreateRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentTokenCreateRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationPaymentTokenCreateRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o PaymentInitiationPaymentTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationPaymentTokenCreateRequest struct {
	value *PaymentInitiationPaymentTokenCreateRequest
	isSet bool
}

func (v NullablePaymentInitiationPaymentTokenCreateRequest) Get() *PaymentInitiationPaymentTokenCreateRequest {
	return v.value
}

func (v *NullablePaymentInitiationPaymentTokenCreateRequest) Set(val *PaymentInitiationPaymentTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentTokenCreateRequest(val *PaymentInitiationPaymentTokenCreateRequest) *NullablePaymentInitiationPaymentTokenCreateRequest {
	return &NullablePaymentInitiationPaymentTokenCreateRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


