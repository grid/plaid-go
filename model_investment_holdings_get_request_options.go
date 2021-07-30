/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InvestmentHoldingsGetRequestOptions An optional object to filter `/investments/holdings/get` results. If provided, must not be `null`.
type InvestmentHoldingsGetRequestOptions struct {
	// An array of `account_id`s to retrieve for the Item. An error will be returned if a provided `account_id` is not associated with the Item.
	AccountIds *[]string `json:"account_ids,omitempty"`
}

// NewInvestmentHoldingsGetRequestOptions instantiates a new InvestmentHoldingsGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentHoldingsGetRequestOptions() *InvestmentHoldingsGetRequestOptions {
	this := InvestmentHoldingsGetRequestOptions{}
	return &this
}

// NewInvestmentHoldingsGetRequestOptionsWithDefaults instantiates a new InvestmentHoldingsGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentHoldingsGetRequestOptionsWithDefaults() *InvestmentHoldingsGetRequestOptions {
	this := InvestmentHoldingsGetRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *InvestmentHoldingsGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentHoldingsGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *InvestmentHoldingsGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *InvestmentHoldingsGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o InvestmentHoldingsGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentHoldingsGetRequestOptions struct {
	value *InvestmentHoldingsGetRequestOptions
	isSet bool
}

func (v NullableInvestmentHoldingsGetRequestOptions) Get() *InvestmentHoldingsGetRequestOptions {
	return v.value
}

func (v *NullableInvestmentHoldingsGetRequestOptions) Set(val *InvestmentHoldingsGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentHoldingsGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentHoldingsGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentHoldingsGetRequestOptions(val *InvestmentHoldingsGetRequestOptions) *NullableInvestmentHoldingsGetRequestOptions {
	return &NullableInvestmentHoldingsGetRequestOptions{value: val, isSet: true}
}

func (v NullableInvestmentHoldingsGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentHoldingsGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


