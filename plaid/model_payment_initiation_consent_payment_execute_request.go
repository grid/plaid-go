/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationConsentPaymentExecuteRequest PaymentInitiationConsentPaymentExecuteRequest defines the request schema for `/payment_initiation/consent/payment/execute`
type PaymentInitiationConsentPaymentExecuteRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The consent ID.
	ConsentId string `json:"consent_id"`
	Amount PaymentAmount `json:"amount"`
	// A random key provided by the client, per unique consent payment. Maximum of 128 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a consent payment fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single payment is created. If the request was successfully processed, it will prevent any payment that uses the same idempotency key, and was received within 24 hours of the first request, from being processed.
	IdempotencyKey string `json:"idempotency_key"`
}

// NewPaymentInitiationConsentPaymentExecuteRequest instantiates a new PaymentInitiationConsentPaymentExecuteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentPaymentExecuteRequest(consentId string, amount PaymentAmount, idempotencyKey string) *PaymentInitiationConsentPaymentExecuteRequest {
	this := PaymentInitiationConsentPaymentExecuteRequest{}
	this.ConsentId = consentId
	this.Amount = amount
	this.IdempotencyKey = idempotencyKey
	return &this
}

// NewPaymentInitiationConsentPaymentExecuteRequestWithDefaults instantiates a new PaymentInitiationConsentPaymentExecuteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentPaymentExecuteRequestWithDefaults() *PaymentInitiationConsentPaymentExecuteRequest {
	this := PaymentInitiationConsentPaymentExecuteRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationConsentPaymentExecuteRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationConsentPaymentExecuteRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetConsentId returns the ConsentId field value
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetConsentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetConsentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentId, true
}

// SetConsentId sets field value
func (o *PaymentInitiationConsentPaymentExecuteRequest) SetConsentId(v string) {
	o.ConsentId = v
}

// GetAmount returns the Amount field value
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetAmount() PaymentAmount {
	if o == nil {
		var ret PaymentAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetAmountOk() (*PaymentAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentInitiationConsentPaymentExecuteRequest) SetAmount(v PaymentAmount) {
	o.Amount = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentPaymentExecuteRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *PaymentInitiationConsentPaymentExecuteRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

func (o PaymentInitiationConsentPaymentExecuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["consent_id"] = o.ConsentId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationConsentPaymentExecuteRequest struct {
	value *PaymentInitiationConsentPaymentExecuteRequest
	isSet bool
}

func (v NullablePaymentInitiationConsentPaymentExecuteRequest) Get() *PaymentInitiationConsentPaymentExecuteRequest {
	return v.value
}

func (v *NullablePaymentInitiationConsentPaymentExecuteRequest) Set(val *PaymentInitiationConsentPaymentExecuteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentPaymentExecuteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentPaymentExecuteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentPaymentExecuteRequest(val *PaymentInitiationConsentPaymentExecuteRequest) *NullablePaymentInitiationConsentPaymentExecuteRequest {
	return &NullablePaymentInitiationConsentPaymentExecuteRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentPaymentExecuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentPaymentExecuteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


