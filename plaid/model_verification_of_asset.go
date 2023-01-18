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

// VerificationOfAsset Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type VerificationOfAsset struct {
	REPORTING_INFORMATION ReportingInformation `json:"REPORTING_INFORMATION"`
	SERVICE_PRODUCT_FULFILLMENT ServiceProductFulfillment `json:"SERVICE_PRODUCT_FULFILLMENT"`
	VERIFICATION_OF_ASSET_RESPONSE VerificationOfAssetResponse `json:"VERIFICATION_OF_ASSET_RESPONSE"`
	AdditionalProperties map[string]interface{}
}

type _VerificationOfAsset VerificationOfAsset

// NewVerificationOfAsset instantiates a new VerificationOfAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationOfAsset(rEPORTINGINFORMATION ReportingInformation, sERVICEPRODUCTFULFILLMENT ServiceProductFulfillment, vERIFICATIONOFASSETRESPONSE VerificationOfAssetResponse) *VerificationOfAsset {
	this := VerificationOfAsset{}
	this.REPORTING_INFORMATION = rEPORTINGINFORMATION
	this.SERVICE_PRODUCT_FULFILLMENT = sERVICEPRODUCTFULFILLMENT
	this.VERIFICATION_OF_ASSET_RESPONSE = vERIFICATIONOFASSETRESPONSE
	return &this
}

// NewVerificationOfAssetWithDefaults instantiates a new VerificationOfAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationOfAssetWithDefaults() *VerificationOfAsset {
	this := VerificationOfAsset{}
	return &this
}

// GetREPORTING_INFORMATION returns the REPORTING_INFORMATION field value
func (o *VerificationOfAsset) GetREPORTING_INFORMATION() ReportingInformation {
	if o == nil {
		var ret ReportingInformation
		return ret
	}

	return o.REPORTING_INFORMATION
}

// GetREPORTING_INFORMATIONOk returns a tuple with the REPORTING_INFORMATION field value
// and a boolean to check if the value has been set.
func (o *VerificationOfAsset) GetREPORTING_INFORMATIONOk() (*ReportingInformation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.REPORTING_INFORMATION, true
}

// SetREPORTING_INFORMATION sets field value
func (o *VerificationOfAsset) SetREPORTING_INFORMATION(v ReportingInformation) {
	o.REPORTING_INFORMATION = v
}

// GetSERVICE_PRODUCT_FULFILLMENT returns the SERVICE_PRODUCT_FULFILLMENT field value
func (o *VerificationOfAsset) GetSERVICE_PRODUCT_FULFILLMENT() ServiceProductFulfillment {
	if o == nil {
		var ret ServiceProductFulfillment
		return ret
	}

	return o.SERVICE_PRODUCT_FULFILLMENT
}

// GetSERVICE_PRODUCT_FULFILLMENTOk returns a tuple with the SERVICE_PRODUCT_FULFILLMENT field value
// and a boolean to check if the value has been set.
func (o *VerificationOfAsset) GetSERVICE_PRODUCT_FULFILLMENTOk() (*ServiceProductFulfillment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE_PRODUCT_FULFILLMENT, true
}

// SetSERVICE_PRODUCT_FULFILLMENT sets field value
func (o *VerificationOfAsset) SetSERVICE_PRODUCT_FULFILLMENT(v ServiceProductFulfillment) {
	o.SERVICE_PRODUCT_FULFILLMENT = v
}

// GetVERIFICATION_OF_ASSET_RESPONSE returns the VERIFICATION_OF_ASSET_RESPONSE field value
func (o *VerificationOfAsset) GetVERIFICATION_OF_ASSET_RESPONSE() VerificationOfAssetResponse {
	if o == nil {
		var ret VerificationOfAssetResponse
		return ret
	}

	return o.VERIFICATION_OF_ASSET_RESPONSE
}

// GetVERIFICATION_OF_ASSET_RESPONSEOk returns a tuple with the VERIFICATION_OF_ASSET_RESPONSE field value
// and a boolean to check if the value has been set.
func (o *VerificationOfAsset) GetVERIFICATION_OF_ASSET_RESPONSEOk() (*VerificationOfAssetResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VERIFICATION_OF_ASSET_RESPONSE, true
}

// SetVERIFICATION_OF_ASSET_RESPONSE sets field value
func (o *VerificationOfAsset) SetVERIFICATION_OF_ASSET_RESPONSE(v VerificationOfAssetResponse) {
	o.VERIFICATION_OF_ASSET_RESPONSE = v
}

func (o VerificationOfAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["REPORTING_INFORMATION"] = o.REPORTING_INFORMATION
	}
	if true {
		toSerialize["SERVICE_PRODUCT_FULFILLMENT"] = o.SERVICE_PRODUCT_FULFILLMENT
	}
	if true {
		toSerialize["VERIFICATION_OF_ASSET_RESPONSE"] = o.VERIFICATION_OF_ASSET_RESPONSE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerificationOfAsset) UnmarshalJSON(bytes []byte) (err error) {
	varVerificationOfAsset := _VerificationOfAsset{}

	if err = json.Unmarshal(bytes, &varVerificationOfAsset); err == nil {
		*o = VerificationOfAsset(varVerificationOfAsset)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "REPORTING_INFORMATION")
		delete(additionalProperties, "SERVICE_PRODUCT_FULFILLMENT")
		delete(additionalProperties, "VERIFICATION_OF_ASSET_RESPONSE")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationOfAsset struct {
	value *VerificationOfAsset
	isSet bool
}

func (v NullableVerificationOfAsset) Get() *VerificationOfAsset {
	return v.value
}

func (v *NullableVerificationOfAsset) Set(val *VerificationOfAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationOfAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationOfAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationOfAsset(val *VerificationOfAsset) *NullableVerificationOfAsset {
	return &NullableVerificationOfAsset{value: val, isSet: true}
}

func (v NullableVerificationOfAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationOfAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


