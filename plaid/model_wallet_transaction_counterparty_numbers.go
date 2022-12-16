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

// WalletTransactionCounterpartyNumbers The counterparty's bank account numbers. Exactly one of IBAN or BACS data is required.
type WalletTransactionCounterpartyNumbers struct {
	Bacs *WalletTransactionCounterpartyBACS `json:"bacs,omitempty"`
	International NullableWalletTransactionCounterpartyInternational `json:"international,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionCounterpartyNumbers WalletTransactionCounterpartyNumbers

// NewWalletTransactionCounterpartyNumbers instantiates a new WalletTransactionCounterpartyNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCounterpartyNumbers() *WalletTransactionCounterpartyNumbers {
	this := WalletTransactionCounterpartyNumbers{}
	return &this
}

// NewWalletTransactionCounterpartyNumbersWithDefaults instantiates a new WalletTransactionCounterpartyNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCounterpartyNumbersWithDefaults() *WalletTransactionCounterpartyNumbers {
	this := WalletTransactionCounterpartyNumbers{}
	return &this
}

// GetBacs returns the Bacs field value if set, zero value otherwise.
func (o *WalletTransactionCounterpartyNumbers) GetBacs() WalletTransactionCounterpartyBACS {
	if o == nil || o.Bacs == nil {
		var ret WalletTransactionCounterpartyBACS
		return ret
	}
	return *o.Bacs
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterpartyNumbers) GetBacsOk() (*WalletTransactionCounterpartyBACS, bool) {
	if o == nil || o.Bacs == nil {
		return nil, false
	}
	return o.Bacs, true
}

// HasBacs returns a boolean if a field has been set.
func (o *WalletTransactionCounterpartyNumbers) HasBacs() bool {
	if o != nil && o.Bacs != nil {
		return true
	}

	return false
}

// SetBacs gets a reference to the given WalletTransactionCounterpartyBACS and assigns it to the Bacs field.
func (o *WalletTransactionCounterpartyNumbers) SetBacs(v WalletTransactionCounterpartyBACS) {
	o.Bacs = &v
}

// GetInternational returns the International field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletTransactionCounterpartyNumbers) GetInternational() WalletTransactionCounterpartyInternational {
	if o == nil || o.International.Get() == nil {
		var ret WalletTransactionCounterpartyInternational
		return ret
	}
	return *o.International.Get()
}

// GetInternationalOk returns a tuple with the International field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletTransactionCounterpartyNumbers) GetInternationalOk() (*WalletTransactionCounterpartyInternational, bool) {
	if o == nil  {
		return nil, false
	}
	return o.International.Get(), o.International.IsSet()
}

// HasInternational returns a boolean if a field has been set.
func (o *WalletTransactionCounterpartyNumbers) HasInternational() bool {
	if o != nil && o.International.IsSet() {
		return true
	}

	return false
}

// SetInternational gets a reference to the given NullableWalletTransactionCounterpartyInternational and assigns it to the International field.
func (o *WalletTransactionCounterpartyNumbers) SetInternational(v WalletTransactionCounterpartyInternational) {
	o.International.Set(&v)
}
// SetInternationalNil sets the value for International to be an explicit nil
func (o *WalletTransactionCounterpartyNumbers) SetInternationalNil() {
	o.International.Set(nil)
}

// UnsetInternational ensures that no value is present for International, not even an explicit nil
func (o *WalletTransactionCounterpartyNumbers) UnsetInternational() {
	o.International.Unset()
}

func (o WalletTransactionCounterpartyNumbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bacs != nil {
		toSerialize["bacs"] = o.Bacs
	}
	if o.International.IsSet() {
		toSerialize["international"] = o.International.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionCounterpartyNumbers) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionCounterpartyNumbers := _WalletTransactionCounterpartyNumbers{}

	if err = json.Unmarshal(bytes, &varWalletTransactionCounterpartyNumbers); err == nil {
		*o = WalletTransactionCounterpartyNumbers(varWalletTransactionCounterpartyNumbers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bacs")
		delete(additionalProperties, "international")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionCounterpartyNumbers struct {
	value *WalletTransactionCounterpartyNumbers
	isSet bool
}

func (v NullableWalletTransactionCounterpartyNumbers) Get() *WalletTransactionCounterpartyNumbers {
	return v.value
}

func (v *NullableWalletTransactionCounterpartyNumbers) Set(val *WalletTransactionCounterpartyNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCounterpartyNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCounterpartyNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCounterpartyNumbers(val *WalletTransactionCounterpartyNumbers) *NullableWalletTransactionCounterpartyNumbers {
	return &NullableWalletTransactionCounterpartyNumbers{value: val, isSet: true}
}

func (v NullableWalletTransactionCounterpartyNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCounterpartyNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


