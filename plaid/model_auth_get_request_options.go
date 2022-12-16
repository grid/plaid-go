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

// AuthGetRequestOptions An optional object to filter `/auth/get` results.
type AuthGetRequestOptions struct {
	// A list of `account_ids` to retrieve for the Item. Note: An error will be returned if a provided `account_id` is not associated with the Item.
	AccountIds *[]string `json:"account_ids,omitempty"`
}

// NewAuthGetRequestOptions instantiates a new AuthGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthGetRequestOptions() *AuthGetRequestOptions {
	this := AuthGetRequestOptions{}
	return &this
}

// NewAuthGetRequestOptionsWithDefaults instantiates a new AuthGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthGetRequestOptionsWithDefaults() *AuthGetRequestOptions {
	this := AuthGetRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *AuthGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *AuthGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *AuthGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o AuthGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableAuthGetRequestOptions struct {
	value *AuthGetRequestOptions
	isSet bool
}

func (v NullableAuthGetRequestOptions) Get() *AuthGetRequestOptions {
	return v.value
}

func (v *NullableAuthGetRequestOptions) Set(val *AuthGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGetRequestOptions(val *AuthGetRequestOptions) *NullableAuthGetRequestOptions {
	return &NullableAuthGetRequestOptions{value: val, isSet: true}
}

func (v NullableAuthGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


