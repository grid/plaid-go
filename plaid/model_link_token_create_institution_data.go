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

// LinkTokenCreateInstitutionData A map containing data used to highlight institutions in Link.
type LinkTokenCreateInstitutionData struct {
	// The routing number of the bank to highlight.
	RoutingNumber *string `json:"routing_number,omitempty"`
}

// NewLinkTokenCreateInstitutionData instantiates a new LinkTokenCreateInstitutionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateInstitutionData() *LinkTokenCreateInstitutionData {
	this := LinkTokenCreateInstitutionData{}
	return &this
}

// NewLinkTokenCreateInstitutionDataWithDefaults instantiates a new LinkTokenCreateInstitutionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateInstitutionDataWithDefaults() *LinkTokenCreateInstitutionData {
	this := LinkTokenCreateInstitutionData{}
	return &this
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *LinkTokenCreateInstitutionData) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateInstitutionData) GetRoutingNumberOk() (*string, bool) {
	if o == nil || o.RoutingNumber == nil {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *LinkTokenCreateInstitutionData) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber != nil {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *LinkTokenCreateInstitutionData) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

func (o LinkTokenCreateInstitutionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoutingNumber != nil {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateInstitutionData struct {
	value *LinkTokenCreateInstitutionData
	isSet bool
}

func (v NullableLinkTokenCreateInstitutionData) Get() *LinkTokenCreateInstitutionData {
	return v.value
}

func (v *NullableLinkTokenCreateInstitutionData) Set(val *LinkTokenCreateInstitutionData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateInstitutionData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateInstitutionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateInstitutionData(val *LinkTokenCreateInstitutionData) *NullableLinkTokenCreateInstitutionData {
	return &NullableLinkTokenCreateInstitutionData{value: val, isSet: true}
}

func (v NullableLinkTokenCreateInstitutionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateInstitutionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


