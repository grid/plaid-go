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

// InstitutionsGetResponse InstitutionsGetResponse defines the response schema for `/institutions/get`
type InstitutionsGetResponse struct {
	// A list of Plaid institutions
	Institutions []Institution `json:"institutions"`
	// The total number of institutions available via this endpoint
	Total int32 `json:"total"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionsGetResponse InstitutionsGetResponse

// NewInstitutionsGetResponse instantiates a new InstitutionsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsGetResponse(institutions []Institution, total int32, requestId string) *InstitutionsGetResponse {
	this := InstitutionsGetResponse{}
	this.Institutions = institutions
	this.Total = total
	this.RequestId = requestId
	return &this
}

// NewInstitutionsGetResponseWithDefaults instantiates a new InstitutionsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsGetResponseWithDefaults() *InstitutionsGetResponse {
	this := InstitutionsGetResponse{}
	return &this
}

// GetInstitutions returns the Institutions field value
func (o *InstitutionsGetResponse) GetInstitutions() []Institution {
	if o == nil {
		var ret []Institution
		return ret
	}

	return o.Institutions
}

// GetInstitutionsOk returns a tuple with the Institutions field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetResponse) GetInstitutionsOk() (*[]Institution, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Institutions, true
}

// SetInstitutions sets field value
func (o *InstitutionsGetResponse) SetInstitutions(v []Institution) {
	o.Institutions = v
}

// GetTotal returns the Total field value
func (o *InstitutionsGetResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetResponse) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InstitutionsGetResponse) SetTotal(v int32) {
	o.Total = v
}

// GetRequestId returns the RequestId field value
func (o *InstitutionsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InstitutionsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InstitutionsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institutions"] = o.Institutions
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionsGetResponse := _InstitutionsGetResponse{}

	if err = json.Unmarshal(bytes, &varInstitutionsGetResponse); err == nil {
		*o = InstitutionsGetResponse(varInstitutionsGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "institutions")
		delete(additionalProperties, "total")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionsGetResponse struct {
	value *InstitutionsGetResponse
	isSet bool
}

func (v NullableInstitutionsGetResponse) Get() *InstitutionsGetResponse {
	return v.value
}

func (v *NullableInstitutionsGetResponse) Set(val *InstitutionsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsGetResponse(val *InstitutionsGetResponse) *NullableInstitutionsGetResponse {
	return &NullableInstitutionsGetResponse{value: val, isSet: true}
}

func (v NullableInstitutionsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


