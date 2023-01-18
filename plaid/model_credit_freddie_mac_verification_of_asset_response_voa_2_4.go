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

// CreditFreddieMacVerificationOfAssetResponseVOA24 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacVerificationOfAssetResponseVOA24 struct {
	ASSETS CreditFreddieMacAssetsVOA24 `json:"ASSETS"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacVerificationOfAssetResponseVOA24 CreditFreddieMacVerificationOfAssetResponseVOA24

// NewCreditFreddieMacVerificationOfAssetResponseVOA24 instantiates a new CreditFreddieMacVerificationOfAssetResponseVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacVerificationOfAssetResponseVOA24(aSSETS CreditFreddieMacAssetsVOA24) *CreditFreddieMacVerificationOfAssetResponseVOA24 {
	this := CreditFreddieMacVerificationOfAssetResponseVOA24{}
	this.ASSETS = aSSETS
	return &this
}

// NewCreditFreddieMacVerificationOfAssetResponseVOA24WithDefaults instantiates a new CreditFreddieMacVerificationOfAssetResponseVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacVerificationOfAssetResponseVOA24WithDefaults() *CreditFreddieMacVerificationOfAssetResponseVOA24 {
	this := CreditFreddieMacVerificationOfAssetResponseVOA24{}
	return &this
}

// GetASSETS returns the ASSETS field value
func (o *CreditFreddieMacVerificationOfAssetResponseVOA24) GetASSETS() CreditFreddieMacAssetsVOA24 {
	if o == nil {
		var ret CreditFreddieMacAssetsVOA24
		return ret
	}

	return o.ASSETS
}

// GetASSETSOk returns a tuple with the ASSETS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetResponseVOA24) GetASSETSOk() (*CreditFreddieMacAssetsVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSETS, true
}

// SetASSETS sets field value
func (o *CreditFreddieMacVerificationOfAssetResponseVOA24) SetASSETS(v CreditFreddieMacAssetsVOA24) {
	o.ASSETS = v
}

func (o CreditFreddieMacVerificationOfAssetResponseVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSETS"] = o.ASSETS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacVerificationOfAssetResponseVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacVerificationOfAssetResponseVOA24 := _CreditFreddieMacVerificationOfAssetResponseVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacVerificationOfAssetResponseVOA24); err == nil {
		*o = CreditFreddieMacVerificationOfAssetResponseVOA24(varCreditFreddieMacVerificationOfAssetResponseVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSETS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacVerificationOfAssetResponseVOA24 struct {
	value *CreditFreddieMacVerificationOfAssetResponseVOA24
	isSet bool
}

func (v NullableCreditFreddieMacVerificationOfAssetResponseVOA24) Get() *CreditFreddieMacVerificationOfAssetResponseVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacVerificationOfAssetResponseVOA24) Set(val *CreditFreddieMacVerificationOfAssetResponseVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacVerificationOfAssetResponseVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacVerificationOfAssetResponseVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacVerificationOfAssetResponseVOA24(val *CreditFreddieMacVerificationOfAssetResponseVOA24) *NullableCreditFreddieMacVerificationOfAssetResponseVOA24 {
	return &NullableCreditFreddieMacVerificationOfAssetResponseVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacVerificationOfAssetResponseVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacVerificationOfAssetResponseVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


