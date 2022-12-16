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

// CreditEmployerVerification An object containing employer data.
type CreditEmployerVerification struct {
	// Name of employer.
	Name NullableString `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreditEmployerVerification CreditEmployerVerification

// NewCreditEmployerVerification instantiates a new CreditEmployerVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditEmployerVerification(name NullableString) *CreditEmployerVerification {
	this := CreditEmployerVerification{}
	this.Name = name
	return &this
}

// NewCreditEmployerVerificationWithDefaults instantiates a new CreditEmployerVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditEmployerVerificationWithDefaults() *CreditEmployerVerification {
	this := CreditEmployerVerification{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmployerVerification) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmployerVerification) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CreditEmployerVerification) SetName(v string) {
	o.Name.Set(&v)
}

func (o CreditEmployerVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditEmployerVerification) UnmarshalJSON(bytes []byte) (err error) {
	varCreditEmployerVerification := _CreditEmployerVerification{}

	if err = json.Unmarshal(bytes, &varCreditEmployerVerification); err == nil {
		*o = CreditEmployerVerification(varCreditEmployerVerification)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditEmployerVerification struct {
	value *CreditEmployerVerification
	isSet bool
}

func (v NullableCreditEmployerVerification) Get() *CreditEmployerVerification {
	return v.value
}

func (v *NullableCreditEmployerVerification) Set(val *CreditEmployerVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditEmployerVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditEmployerVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditEmployerVerification(val *CreditEmployerVerification) *NullableCreditEmployerVerification {
	return &NullableCreditEmployerVerification{value: val, isSet: true}
}

func (v NullableCreditEmployerVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditEmployerVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


