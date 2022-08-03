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

// LinkTokenCreateInvestmentFilter A filter to apply to `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier).
type LinkTokenCreateInvestmentFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema). 
	AccountSubtypes *[]InvestmentAccountSubtype `json:"account_subtypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenCreateInvestmentFilter LinkTokenCreateInvestmentFilter

// NewLinkTokenCreateInvestmentFilter instantiates a new LinkTokenCreateInvestmentFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateInvestmentFilter() *LinkTokenCreateInvestmentFilter {
	this := LinkTokenCreateInvestmentFilter{}
	return &this
}

// NewLinkTokenCreateInvestmentFilterWithDefaults instantiates a new LinkTokenCreateInvestmentFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateInvestmentFilterWithDefaults() *LinkTokenCreateInvestmentFilter {
	this := LinkTokenCreateInvestmentFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value if set, zero value otherwise.
func (o *LinkTokenCreateInvestmentFilter) GetAccountSubtypes() []InvestmentAccountSubtype {
	if o == nil || o.AccountSubtypes == nil {
		var ret []InvestmentAccountSubtype
		return ret
	}
	return *o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateInvestmentFilter) GetAccountSubtypesOk() (*[]InvestmentAccountSubtype, bool) {
	if o == nil || o.AccountSubtypes == nil {
		return nil, false
	}
	return o.AccountSubtypes, true
}

// HasAccountSubtypes returns a boolean if a field has been set.
func (o *LinkTokenCreateInvestmentFilter) HasAccountSubtypes() bool {
	if o != nil && o.AccountSubtypes != nil {
		return true
	}

	return false
}

// SetAccountSubtypes gets a reference to the given []InvestmentAccountSubtype and assigns it to the AccountSubtypes field.
func (o *LinkTokenCreateInvestmentFilter) SetAccountSubtypes(v []InvestmentAccountSubtype) {
	o.AccountSubtypes = &v
}

func (o LinkTokenCreateInvestmentFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSubtypes != nil {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenCreateInvestmentFilter) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenCreateInvestmentFilter := _LinkTokenCreateInvestmentFilter{}

	if err = json.Unmarshal(bytes, &varLinkTokenCreateInvestmentFilter); err == nil {
		*o = LinkTokenCreateInvestmentFilter(varLinkTokenCreateInvestmentFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenCreateInvestmentFilter struct {
	value *LinkTokenCreateInvestmentFilter
	isSet bool
}

func (v NullableLinkTokenCreateInvestmentFilter) Get() *LinkTokenCreateInvestmentFilter {
	return v.value
}

func (v *NullableLinkTokenCreateInvestmentFilter) Set(val *LinkTokenCreateInvestmentFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateInvestmentFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateInvestmentFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateInvestmentFilter(val *LinkTokenCreateInvestmentFilter) *NullableLinkTokenCreateInvestmentFilter {
	return &NullableLinkTokenCreateInvestmentFilter{value: val, isSet: true}
}

func (v NullableLinkTokenCreateInvestmentFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateInvestmentFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


