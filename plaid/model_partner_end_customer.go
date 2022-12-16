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

// PartnerEndCustomer The details for an end customer.
type PartnerEndCustomer struct {
	ClientId *string `json:"client_id,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Status *PartnerEndCustomerStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomer PartnerEndCustomer

// NewPartnerEndCustomer instantiates a new PartnerEndCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomer() *PartnerEndCustomer {
	this := PartnerEndCustomer{}
	return &this
}

// NewPartnerEndCustomerWithDefaults instantiates a new PartnerEndCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerWithDefaults() *PartnerEndCustomer {
	this := PartnerEndCustomer{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PartnerEndCustomer) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomer) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PartnerEndCustomer) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PartnerEndCustomer) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *PartnerEndCustomer) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomer) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *PartnerEndCustomer) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *PartnerEndCustomer) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PartnerEndCustomer) GetStatus() PartnerEndCustomerStatus {
	if o == nil || o.Status == nil {
		var ret PartnerEndCustomerStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomer) GetStatusOk() (*PartnerEndCustomerStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PartnerEndCustomer) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PartnerEndCustomerStatus and assigns it to the Status field.
func (o *PartnerEndCustomer) SetStatus(v PartnerEndCustomerStatus) {
	o.Status = &v
}

func (o PartnerEndCustomer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CompanyName != nil {
		toSerialize["company_name"] = o.CompanyName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomer := _PartnerEndCustomer{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomer); err == nil {
		*o = PartnerEndCustomer(varPartnerEndCustomer)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomer struct {
	value *PartnerEndCustomer
	isSet bool
}

func (v NullablePartnerEndCustomer) Get() *PartnerEndCustomer {
	return v.value
}

func (v *NullablePartnerEndCustomer) Set(val *PartnerEndCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomer(val *PartnerEndCustomer) *NullablePartnerEndCustomer {
	return &NullablePartnerEndCustomer{value: val, isSet: true}
}

func (v NullablePartnerEndCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


