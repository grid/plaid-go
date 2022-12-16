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

// NumbersEFT Identifying information for transferring money to or from a Canadian bank account via EFT.
type NumbersEFT struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The EFT account number for the account
	Account string `json:"account"`
	// The EFT institution number for the account
	Institution string `json:"institution"`
	// The EFT branch number for the account
	Branch string `json:"branch"`
	AdditionalProperties map[string]interface{}
}

type _NumbersEFT NumbersEFT

// NewNumbersEFT instantiates a new NumbersEFT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersEFT(accountId string, account string, institution string, branch string) *NumbersEFT {
	this := NumbersEFT{}
	this.AccountId = accountId
	this.Account = account
	this.Institution = institution
	this.Branch = branch
	return &this
}

// NewNumbersEFTWithDefaults instantiates a new NumbersEFT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersEFTWithDefaults() *NumbersEFT {
	this := NumbersEFT{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersEFT) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersEFT) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersEFT) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersEFT) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersEFT) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersEFT) SetAccount(v string) {
	o.Account = v
}

// GetInstitution returns the Institution field value
func (o *NumbersEFT) GetInstitution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value
// and a boolean to check if the value has been set.
func (o *NumbersEFT) GetInstitutionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Institution, true
}

// SetInstitution sets field value
func (o *NumbersEFT) SetInstitution(v string) {
	o.Institution = v
}

// GetBranch returns the Branch field value
func (o *NumbersEFT) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *NumbersEFT) GetBranchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *NumbersEFT) SetBranch(v string) {
	o.Branch = v
}

func (o NumbersEFT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["institution"] = o.Institution
	}
	if true {
		toSerialize["branch"] = o.Branch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NumbersEFT) UnmarshalJSON(bytes []byte) (err error) {
	varNumbersEFT := _NumbersEFT{}

	if err = json.Unmarshal(bytes, &varNumbersEFT); err == nil {
		*o = NumbersEFT(varNumbersEFT)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account")
		delete(additionalProperties, "institution")
		delete(additionalProperties, "branch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumbersEFT struct {
	value *NumbersEFT
	isSet bool
}

func (v NullableNumbersEFT) Get() *NumbersEFT {
	return v.value
}

func (v *NullableNumbersEFT) Set(val *NumbersEFT) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersEFT) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersEFT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersEFT(val *NumbersEFT) *NullableNumbersEFT {
	return &NullableNumbersEFT{value: val, isSet: true}
}

func (v NullableNumbersEFT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersEFT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


