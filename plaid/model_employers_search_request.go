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

// EmployersSearchRequest EmployersSearchRequest defines the request schema for `/employers/search`.
type EmployersSearchRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The employer name to be searched for.
	Query string `json:"query"`
	// The Plaid products the returned employers should support. Currently, this field must be set to `\"deposit_switch\"`.
	Products []string `json:"products"`
}

// NewEmployersSearchRequest instantiates a new EmployersSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployersSearchRequest(query string, products []string) *EmployersSearchRequest {
	this := EmployersSearchRequest{}
	this.Query = query
	this.Products = products
	return &this
}

// NewEmployersSearchRequestWithDefaults instantiates a new EmployersSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployersSearchRequestWithDefaults() *EmployersSearchRequest {
	this := EmployersSearchRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *EmployersSearchRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployersSearchRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *EmployersSearchRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *EmployersSearchRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *EmployersSearchRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployersSearchRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *EmployersSearchRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *EmployersSearchRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetQuery returns the Query field value
func (o *EmployersSearchRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EmployersSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *EmployersSearchRequest) SetQuery(v string) {
	o.Query = v
}

// GetProducts returns the Products field value
func (o *EmployersSearchRequest) GetProducts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *EmployersSearchRequest) GetProductsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Products, true
}

// SetProducts sets field value
func (o *EmployersSearchRequest) SetProducts(v []string) {
	o.Products = v
}

func (o EmployersSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["products"] = o.Products
	}
	return json.Marshal(toSerialize)
}

type NullableEmployersSearchRequest struct {
	value *EmployersSearchRequest
	isSet bool
}

func (v NullableEmployersSearchRequest) Get() *EmployersSearchRequest {
	return v.value
}

func (v *NullableEmployersSearchRequest) Set(val *EmployersSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployersSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployersSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployersSearchRequest(val *EmployersSearchRequest) *NullableEmployersSearchRequest {
	return &NullableEmployersSearchRequest{value: val, isSet: true}
}

func (v NullableEmployersSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployersSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


