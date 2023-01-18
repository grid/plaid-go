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

// PaymentConsentPeriodicAmountAmount Maximum cumulative amount for all payments in the specified interval.
type PaymentConsentPeriodicAmountAmount struct {
	Currency PaymentAmountCurrency `json:"currency"`
	// The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.
	Value float64 `json:"value"`
}

// NewPaymentConsentPeriodicAmountAmount instantiates a new PaymentConsentPeriodicAmountAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConsentPeriodicAmountAmount(currency PaymentAmountCurrency, value float64) *PaymentConsentPeriodicAmountAmount {
	this := PaymentConsentPeriodicAmountAmount{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewPaymentConsentPeriodicAmountAmountWithDefaults instantiates a new PaymentConsentPeriodicAmountAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConsentPeriodicAmountAmountWithDefaults() *PaymentConsentPeriodicAmountAmount {
	this := PaymentConsentPeriodicAmountAmount{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *PaymentConsentPeriodicAmountAmount) GetCurrency() PaymentAmountCurrency {
	if o == nil {
		var ret PaymentAmountCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentPeriodicAmountAmount) GetCurrencyOk() (*PaymentAmountCurrency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentConsentPeriodicAmountAmount) SetCurrency(v PaymentAmountCurrency) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *PaymentConsentPeriodicAmountAmount) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentPeriodicAmountAmount) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PaymentConsentPeriodicAmountAmount) SetValue(v float64) {
	o.Value = v
}

func (o PaymentConsentPeriodicAmountAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentConsentPeriodicAmountAmount struct {
	value *PaymentConsentPeriodicAmountAmount
	isSet bool
}

func (v NullablePaymentConsentPeriodicAmountAmount) Get() *PaymentConsentPeriodicAmountAmount {
	return v.value
}

func (v *NullablePaymentConsentPeriodicAmountAmount) Set(val *PaymentConsentPeriodicAmountAmount) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConsentPeriodicAmountAmount) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConsentPeriodicAmountAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConsentPeriodicAmountAmount(val *PaymentConsentPeriodicAmountAmount) *NullablePaymentConsentPeriodicAmountAmount {
	return &NullablePaymentConsentPeriodicAmountAmount{value: val, isSet: true}
}

func (v NullablePaymentConsentPeriodicAmountAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConsentPeriodicAmountAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


