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

// PartnerCustomerEnableRequest Request schema for `/partner/customer/enable`.
type PartnerCustomerEnableRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	EndCustomerClientId string `json:"end_customer_client_id"`
}

// NewPartnerCustomerEnableRequest instantiates a new PartnerCustomerEnableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCustomerEnableRequest(endCustomerClientId string) *PartnerCustomerEnableRequest {
	this := PartnerCustomerEnableRequest{}
	this.EndCustomerClientId = endCustomerClientId
	return &this
}

// NewPartnerCustomerEnableRequestWithDefaults instantiates a new PartnerCustomerEnableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCustomerEnableRequestWithDefaults() *PartnerCustomerEnableRequest {
	this := PartnerCustomerEnableRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PartnerCustomerEnableRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerEnableRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PartnerCustomerEnableRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PartnerCustomerEnableRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PartnerCustomerEnableRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerEnableRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PartnerCustomerEnableRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PartnerCustomerEnableRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetEndCustomerClientId returns the EndCustomerClientId field value
func (o *PartnerCustomerEnableRequest) GetEndCustomerClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndCustomerClientId
}

// GetEndCustomerClientIdOk returns a tuple with the EndCustomerClientId field value
// and a boolean to check if the value has been set.
func (o *PartnerCustomerEnableRequest) GetEndCustomerClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndCustomerClientId, true
}

// SetEndCustomerClientId sets field value
func (o *PartnerCustomerEnableRequest) SetEndCustomerClientId(v string) {
	o.EndCustomerClientId = v
}

func (o PartnerCustomerEnableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["end_customer_client_id"] = o.EndCustomerClientId
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerCustomerEnableRequest struct {
	value *PartnerCustomerEnableRequest
	isSet bool
}

func (v NullablePartnerCustomerEnableRequest) Get() *PartnerCustomerEnableRequest {
	return v.value
}

func (v *NullablePartnerCustomerEnableRequest) Set(val *PartnerCustomerEnableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCustomerEnableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCustomerEnableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCustomerEnableRequest(val *PartnerCustomerEnableRequest) *NullablePartnerCustomerEnableRequest {
	return &NullablePartnerCustomerEnableRequest{value: val, isSet: true}
}

func (v NullablePartnerCustomerEnableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCustomerEnableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


