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

// InstitutionsGetByIdResponse InstitutionsGetByIdResponse defines the response schema for `/institutions/get_by_id`
type InstitutionsGetByIdResponse struct {
	Institution Institution `json:"institution"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionsGetByIdResponse InstitutionsGetByIdResponse

// NewInstitutionsGetByIdResponse instantiates a new InstitutionsGetByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsGetByIdResponse(institution Institution, requestId string) *InstitutionsGetByIdResponse {
	this := InstitutionsGetByIdResponse{}
	this.Institution = institution
	this.RequestId = requestId
	return &this
}

// NewInstitutionsGetByIdResponseWithDefaults instantiates a new InstitutionsGetByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsGetByIdResponseWithDefaults() *InstitutionsGetByIdResponse {
	this := InstitutionsGetByIdResponse{}
	return &this
}

// GetInstitution returns the Institution field value
func (o *InstitutionsGetByIdResponse) GetInstitution() Institution {
	if o == nil {
		var ret Institution
		return ret
	}

	return o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdResponse) GetInstitutionOk() (*Institution, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Institution, true
}

// SetInstitution sets field value
func (o *InstitutionsGetByIdResponse) SetInstitution(v Institution) {
	o.Institution = v
}

// GetRequestId returns the RequestId field value
func (o *InstitutionsGetByIdResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InstitutionsGetByIdResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InstitutionsGetByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institution"] = o.Institution
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionsGetByIdResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionsGetByIdResponse := _InstitutionsGetByIdResponse{}

	if err = json.Unmarshal(bytes, &varInstitutionsGetByIdResponse); err == nil {
		*o = InstitutionsGetByIdResponse(varInstitutionsGetByIdResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "institution")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionsGetByIdResponse struct {
	value *InstitutionsGetByIdResponse
	isSet bool
}

func (v NullableInstitutionsGetByIdResponse) Get() *InstitutionsGetByIdResponse {
	return v.value
}

func (v *NullableInstitutionsGetByIdResponse) Set(val *InstitutionsGetByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsGetByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsGetByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsGetByIdResponse(val *InstitutionsGetByIdResponse) *NullableInstitutionsGetByIdResponse {
	return &NullableInstitutionsGetByIdResponse{value: val, isSet: true}
}

func (v NullableInstitutionsGetByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsGetByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


