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

// CreditFreddieMacServiceVOA24 A collection of details related to a fulfillment service or product in terms of request, process and result.
type CreditFreddieMacServiceVOA24 struct {
	VERIFICATION_OF_ASSET []CreditFreddieMacVerificationOfAssetVOA24 `json:"VERIFICATION_OF_ASSET"`
	STATUSES Statuses `json:"STATUSES"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacServiceVOA24 CreditFreddieMacServiceVOA24

// NewCreditFreddieMacServiceVOA24 instantiates a new CreditFreddieMacServiceVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacServiceVOA24(vERIFICATIONOFASSET []CreditFreddieMacVerificationOfAssetVOA24, sTATUSES Statuses) *CreditFreddieMacServiceVOA24 {
	this := CreditFreddieMacServiceVOA24{}
	this.VERIFICATION_OF_ASSET = vERIFICATIONOFASSET
	this.STATUSES = sTATUSES
	return &this
}

// NewCreditFreddieMacServiceVOA24WithDefaults instantiates a new CreditFreddieMacServiceVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacServiceVOA24WithDefaults() *CreditFreddieMacServiceVOA24 {
	this := CreditFreddieMacServiceVOA24{}
	return &this
}

// GetVERIFICATION_OF_ASSET returns the VERIFICATION_OF_ASSET field value
func (o *CreditFreddieMacServiceVOA24) GetVERIFICATION_OF_ASSET() []CreditFreddieMacVerificationOfAssetVOA24 {
	if o == nil {
		var ret []CreditFreddieMacVerificationOfAssetVOA24
		return ret
	}

	return o.VERIFICATION_OF_ASSET
}

// GetVERIFICATION_OF_ASSETOk returns a tuple with the VERIFICATION_OF_ASSET field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacServiceVOA24) GetVERIFICATION_OF_ASSETOk() (*[]CreditFreddieMacVerificationOfAssetVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VERIFICATION_OF_ASSET, true
}

// SetVERIFICATION_OF_ASSET sets field value
func (o *CreditFreddieMacServiceVOA24) SetVERIFICATION_OF_ASSET(v []CreditFreddieMacVerificationOfAssetVOA24) {
	o.VERIFICATION_OF_ASSET = v
}

// GetSTATUSES returns the STATUSES field value
func (o *CreditFreddieMacServiceVOA24) GetSTATUSES() Statuses {
	if o == nil {
		var ret Statuses
		return ret
	}

	return o.STATUSES
}

// GetSTATUSESOk returns a tuple with the STATUSES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacServiceVOA24) GetSTATUSESOk() (*Statuses, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.STATUSES, true
}

// SetSTATUSES sets field value
func (o *CreditFreddieMacServiceVOA24) SetSTATUSES(v Statuses) {
	o.STATUSES = v
}

func (o CreditFreddieMacServiceVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["VERIFICATION_OF_ASSET"] = o.VERIFICATION_OF_ASSET
	}
	if true {
		toSerialize["STATUSES"] = o.STATUSES
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacServiceVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacServiceVOA24 := _CreditFreddieMacServiceVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacServiceVOA24); err == nil {
		*o = CreditFreddieMacServiceVOA24(varCreditFreddieMacServiceVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "VERIFICATION_OF_ASSET")
		delete(additionalProperties, "STATUSES")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacServiceVOA24 struct {
	value *CreditFreddieMacServiceVOA24
	isSet bool
}

func (v NullableCreditFreddieMacServiceVOA24) Get() *CreditFreddieMacServiceVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacServiceVOA24) Set(val *CreditFreddieMacServiceVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacServiceVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacServiceVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacServiceVOA24(val *CreditFreddieMacServiceVOA24) *NullableCreditFreddieMacServiceVOA24 {
	return &NullableCreditFreddieMacServiceVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacServiceVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacServiceVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


