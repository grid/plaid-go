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

// ReportingInformation Information about an report identifier and a report name.
type ReportingInformation struct {
	// No documentation available
	ReportingInformationIdentifier string `json:"ReportingInformationIdentifier"`
	AdditionalProperties map[string]interface{}
}

type _ReportingInformation ReportingInformation

// NewReportingInformation instantiates a new ReportingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingInformation(reportingInformationIdentifier string) *ReportingInformation {
	this := ReportingInformation{}
	this.ReportingInformationIdentifier = reportingInformationIdentifier
	return &this
}

// NewReportingInformationWithDefaults instantiates a new ReportingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingInformationWithDefaults() *ReportingInformation {
	this := ReportingInformation{}
	return &this
}

// GetReportingInformationIdentifier returns the ReportingInformationIdentifier field value
func (o *ReportingInformation) GetReportingInformationIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportingInformationIdentifier
}

// GetReportingInformationIdentifierOk returns a tuple with the ReportingInformationIdentifier field value
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetReportingInformationIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportingInformationIdentifier, true
}

// SetReportingInformationIdentifier sets field value
func (o *ReportingInformation) SetReportingInformationIdentifier(v string) {
	o.ReportingInformationIdentifier = v
}

func (o ReportingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ReportingInformationIdentifier"] = o.ReportingInformationIdentifier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReportingInformation) UnmarshalJSON(bytes []byte) (err error) {
	varReportingInformation := _ReportingInformation{}

	if err = json.Unmarshal(bytes, &varReportingInformation); err == nil {
		*o = ReportingInformation(varReportingInformation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ReportingInformationIdentifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportingInformation struct {
	value *ReportingInformation
	isSet bool
}

func (v NullableReportingInformation) Get() *ReportingInformation {
	return v.value
}

func (v *NullableReportingInformation) Set(val *ReportingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingInformation(val *ReportingInformation) *NullableReportingInformation {
	return &NullableReportingInformation{value: val, isSet: true}
}

func (v NullableReportingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


