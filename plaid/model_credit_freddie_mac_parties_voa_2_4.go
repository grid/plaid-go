/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacPartiesVOA24 A collection of objects that define specific parties to a deal. This includes the direct participating parties, such as borrower and seller and the indirect parties such as the credit report provider.
type CreditFreddieMacPartiesVOA24 struct {
	PARTY []CreditFreddieMacPartyVOA24 `json:"PARTY"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacPartiesVOA24 CreditFreddieMacPartiesVOA24

// NewCreditFreddieMacPartiesVOA24 instantiates a new CreditFreddieMacPartiesVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacPartiesVOA24(pARTY []CreditFreddieMacPartyVOA24) *CreditFreddieMacPartiesVOA24 {
	this := CreditFreddieMacPartiesVOA24{}
	this.PARTY = pARTY
	return &this
}

// NewCreditFreddieMacPartiesVOA24WithDefaults instantiates a new CreditFreddieMacPartiesVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacPartiesVOA24WithDefaults() *CreditFreddieMacPartiesVOA24 {
	this := CreditFreddieMacPartiesVOA24{}
	return &this
}

// GetPARTY returns the PARTY field value
func (o *CreditFreddieMacPartiesVOA24) GetPARTY() []CreditFreddieMacPartyVOA24 {
	if o == nil {
		var ret []CreditFreddieMacPartyVOA24
		return ret
	}

	return o.PARTY
}

// GetPARTYOk returns a tuple with the PARTY field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacPartiesVOA24) GetPARTYOk() (*[]CreditFreddieMacPartyVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PARTY, true
}

// SetPARTY sets field value
func (o *CreditFreddieMacPartiesVOA24) SetPARTY(v []CreditFreddieMacPartyVOA24) {
	o.PARTY = v
}

func (o CreditFreddieMacPartiesVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PARTY"] = o.PARTY
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacPartiesVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacPartiesVOA24 := _CreditFreddieMacPartiesVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacPartiesVOA24); err == nil {
		*o = CreditFreddieMacPartiesVOA24(varCreditFreddieMacPartiesVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "PARTY")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacPartiesVOA24 struct {
	value *CreditFreddieMacPartiesVOA24
	isSet bool
}

func (v NullableCreditFreddieMacPartiesVOA24) Get() *CreditFreddieMacPartiesVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacPartiesVOA24) Set(val *CreditFreddieMacPartiesVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacPartiesVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacPartiesVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacPartiesVOA24(val *CreditFreddieMacPartiesVOA24) *NullableCreditFreddieMacPartiesVOA24 {
	return &NullableCreditFreddieMacPartiesVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacPartiesVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacPartiesVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

