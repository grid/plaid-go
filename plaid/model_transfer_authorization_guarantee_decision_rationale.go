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

// TransferAuthorizationGuaranteeDecisionRationale The rationale for Plaid's decision to not guarantee a transfer. Will be `null` unless `guarantee_decision` is `NOT_GUARANTEED`.
type TransferAuthorizationGuaranteeDecisionRationale struct {
	Code TransferAuthorizationGuaranteeDecisionRationaleCode `json:"code"`
	// A human-readable description of why the transfer cannot be guaranteed.
	Description string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorizationGuaranteeDecisionRationale TransferAuthorizationGuaranteeDecisionRationale

// NewTransferAuthorizationGuaranteeDecisionRationale instantiates a new TransferAuthorizationGuaranteeDecisionRationale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationGuaranteeDecisionRationale(code TransferAuthorizationGuaranteeDecisionRationaleCode, description string) *TransferAuthorizationGuaranteeDecisionRationale {
	this := TransferAuthorizationGuaranteeDecisionRationale{}
	this.Code = code
	this.Description = description
	return &this
}

// NewTransferAuthorizationGuaranteeDecisionRationaleWithDefaults instantiates a new TransferAuthorizationGuaranteeDecisionRationale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationGuaranteeDecisionRationaleWithDefaults() *TransferAuthorizationGuaranteeDecisionRationale {
	this := TransferAuthorizationGuaranteeDecisionRationale{}
	return &this
}

// GetCode returns the Code field value
func (o *TransferAuthorizationGuaranteeDecisionRationale) GetCode() TransferAuthorizationGuaranteeDecisionRationaleCode {
	if o == nil {
		var ret TransferAuthorizationGuaranteeDecisionRationaleCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationGuaranteeDecisionRationale) GetCodeOk() (*TransferAuthorizationGuaranteeDecisionRationaleCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TransferAuthorizationGuaranteeDecisionRationale) SetCode(v TransferAuthorizationGuaranteeDecisionRationaleCode) {
	o.Code = v
}

// GetDescription returns the Description field value
func (o *TransferAuthorizationGuaranteeDecisionRationale) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationGuaranteeDecisionRationale) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferAuthorizationGuaranteeDecisionRationale) SetDescription(v string) {
	o.Description = v
}

func (o TransferAuthorizationGuaranteeDecisionRationale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorizationGuaranteeDecisionRationale) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorizationGuaranteeDecisionRationale := _TransferAuthorizationGuaranteeDecisionRationale{}

	if err = json.Unmarshal(bytes, &varTransferAuthorizationGuaranteeDecisionRationale); err == nil {
		*o = TransferAuthorizationGuaranteeDecisionRationale(varTransferAuthorizationGuaranteeDecisionRationale)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorizationGuaranteeDecisionRationale struct {
	value *TransferAuthorizationGuaranteeDecisionRationale
	isSet bool
}

func (v NullableTransferAuthorizationGuaranteeDecisionRationale) Get() *TransferAuthorizationGuaranteeDecisionRationale {
	return v.value
}

func (v *NullableTransferAuthorizationGuaranteeDecisionRationale) Set(val *TransferAuthorizationGuaranteeDecisionRationale) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationGuaranteeDecisionRationale) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationGuaranteeDecisionRationale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationGuaranteeDecisionRationale(val *TransferAuthorizationGuaranteeDecisionRationale) *NullableTransferAuthorizationGuaranteeDecisionRationale {
	return &NullableTransferAuthorizationGuaranteeDecisionRationale{value: val, isSet: true}
}

func (v NullableTransferAuthorizationGuaranteeDecisionRationale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationGuaranteeDecisionRationale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


