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

// WalletNumbers An object representing the e-wallet account numbers
type WalletNumbers struct {
	Bacs NullableRecipientBACS `json:"bacs,omitempty"`
	International NullableNumbersInternationalIBAN `json:"international,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletNumbers WalletNumbers

// NewWalletNumbers instantiates a new WalletNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletNumbers() *WalletNumbers {
	this := WalletNumbers{}
	return &this
}

// NewWalletNumbersWithDefaults instantiates a new WalletNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletNumbersWithDefaults() *WalletNumbers {
	this := WalletNumbers{}
	return &this
}

// GetBacs returns the Bacs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletNumbers) GetBacs() RecipientBACS {
	if o == nil || o.Bacs.Get() == nil {
		var ret RecipientBACS
		return ret
	}
	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletNumbers) GetBacsOk() (*RecipientBACS, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// HasBacs returns a boolean if a field has been set.
func (o *WalletNumbers) HasBacs() bool {
	if o != nil && o.Bacs.IsSet() {
		return true
	}

	return false
}

// SetBacs gets a reference to the given NullableRecipientBACS and assigns it to the Bacs field.
func (o *WalletNumbers) SetBacs(v RecipientBACS) {
	o.Bacs.Set(&v)
}
// SetBacsNil sets the value for Bacs to be an explicit nil
func (o *WalletNumbers) SetBacsNil() {
	o.Bacs.Set(nil)
}

// UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
func (o *WalletNumbers) UnsetBacs() {
	o.Bacs.Unset()
}

// GetInternational returns the International field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletNumbers) GetInternational() NumbersInternationalIBAN {
	if o == nil || o.International.Get() == nil {
		var ret NumbersInternationalIBAN
		return ret
	}
	return *o.International.Get()
}

// GetInternationalOk returns a tuple with the International field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletNumbers) GetInternationalOk() (*NumbersInternationalIBAN, bool) {
	if o == nil  {
		return nil, false
	}
	return o.International.Get(), o.International.IsSet()
}

// HasInternational returns a boolean if a field has been set.
func (o *WalletNumbers) HasInternational() bool {
	if o != nil && o.International.IsSet() {
		return true
	}

	return false
}

// SetInternational gets a reference to the given NullableNumbersInternationalIBAN and assigns it to the International field.
func (o *WalletNumbers) SetInternational(v NumbersInternationalIBAN) {
	o.International.Set(&v)
}
// SetInternationalNil sets the value for International to be an explicit nil
func (o *WalletNumbers) SetInternationalNil() {
	o.International.Set(nil)
}

// UnsetInternational ensures that no value is present for International, not even an explicit nil
func (o *WalletNumbers) UnsetInternational() {
	o.International.Unset()
}

func (o WalletNumbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bacs.IsSet() {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	if o.International.IsSet() {
		toSerialize["international"] = o.International.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletNumbers) UnmarshalJSON(bytes []byte) (err error) {
	varWalletNumbers := _WalletNumbers{}

	if err = json.Unmarshal(bytes, &varWalletNumbers); err == nil {
		*o = WalletNumbers(varWalletNumbers)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bacs")
		delete(additionalProperties, "international")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletNumbers struct {
	value *WalletNumbers
	isSet bool
}

func (v NullableWalletNumbers) Get() *WalletNumbers {
	return v.value
}

func (v *NullableWalletNumbers) Set(val *WalletNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletNumbers(val *WalletNumbers) *NullableWalletNumbers {
	return &NullableWalletNumbers{value: val, isSet: true}
}

func (v NullableWalletNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


