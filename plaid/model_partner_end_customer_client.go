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

// PartnerEndCustomerClient The end customer details for the newly-created customer client.
type PartnerEndCustomerClient struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
}

// NewPartnerEndCustomerClient instantiates a new PartnerEndCustomerClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerClient() *PartnerEndCustomerClient {
	this := PartnerEndCustomerClient{}
	return &this
}

// NewPartnerEndCustomerClientWithDefaults instantiates a new PartnerEndCustomerClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerClientWithDefaults() *PartnerEndCustomerClient {
	this := PartnerEndCustomerClient{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PartnerEndCustomerClient) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerClient) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PartnerEndCustomerClient) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PartnerEndCustomerClient) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *PartnerEndCustomerClient) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerClient) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *PartnerEndCustomerClient) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *PartnerEndCustomerClient) SetCompanyName(v string) {
	o.CompanyName = &v
}

func (o PartnerEndCustomerClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CompanyName != nil {
		toSerialize["company_name"] = o.CompanyName
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerEndCustomerClient struct {
	value *PartnerEndCustomerClient
	isSet bool
}

func (v NullablePartnerEndCustomerClient) Get() *PartnerEndCustomerClient {
	return v.value
}

func (v *NullablePartnerEndCustomerClient) Set(val *PartnerEndCustomerClient) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerClient) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerClient(val *PartnerEndCustomerClient) *NullablePartnerEndCustomerClient {
	return &NullablePartnerEndCustomerClient{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


