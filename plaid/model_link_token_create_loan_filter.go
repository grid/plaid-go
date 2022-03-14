/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateLoanFilter A filter to apply to `loan`-type accounts
type LinkTokenCreateLoanFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema). 
	AccountSubtypes *[]LoanAccountSubtype `json:"account_subtypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenCreateLoanFilter LinkTokenCreateLoanFilter

// NewLinkTokenCreateLoanFilter instantiates a new LinkTokenCreateLoanFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateLoanFilter() *LinkTokenCreateLoanFilter {
	this := LinkTokenCreateLoanFilter{}
	return &this
}

// NewLinkTokenCreateLoanFilterWithDefaults instantiates a new LinkTokenCreateLoanFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateLoanFilterWithDefaults() *LinkTokenCreateLoanFilter {
	this := LinkTokenCreateLoanFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value if set, zero value otherwise.
func (o *LinkTokenCreateLoanFilter) GetAccountSubtypes() []LoanAccountSubtype {
	if o == nil || o.AccountSubtypes == nil {
		var ret []LoanAccountSubtype
		return ret
	}
	return *o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateLoanFilter) GetAccountSubtypesOk() (*[]LoanAccountSubtype, bool) {
	if o == nil || o.AccountSubtypes == nil {
		return nil, false
	}
	return o.AccountSubtypes, true
}

// HasAccountSubtypes returns a boolean if a field has been set.
func (o *LinkTokenCreateLoanFilter) HasAccountSubtypes() bool {
	if o != nil && o.AccountSubtypes != nil {
		return true
	}

	return false
}

// SetAccountSubtypes gets a reference to the given []LoanAccountSubtype and assigns it to the AccountSubtypes field.
func (o *LinkTokenCreateLoanFilter) SetAccountSubtypes(v []LoanAccountSubtype) {
	o.AccountSubtypes = &v
}

func (o LinkTokenCreateLoanFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSubtypes != nil {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenCreateLoanFilter) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenCreateLoanFilter := _LinkTokenCreateLoanFilter{}

	if err = json.Unmarshal(bytes, &varLinkTokenCreateLoanFilter); err == nil {
		*o = LinkTokenCreateLoanFilter(varLinkTokenCreateLoanFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenCreateLoanFilter struct {
	value *LinkTokenCreateLoanFilter
	isSet bool
}

func (v NullableLinkTokenCreateLoanFilter) Get() *LinkTokenCreateLoanFilter {
	return v.value
}

func (v *NullableLinkTokenCreateLoanFilter) Set(val *LinkTokenCreateLoanFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateLoanFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateLoanFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateLoanFilter(val *LinkTokenCreateLoanFilter) *NullableLinkTokenCreateLoanFilter {
	return &NullableLinkTokenCreateLoanFilter{value: val, isSet: true}
}

func (v NullableLinkTokenCreateLoanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateLoanFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


