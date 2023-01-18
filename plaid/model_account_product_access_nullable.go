/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.214.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AccountProductAccessNullable Allow the application to access specific products on this account
type AccountProductAccessNullable struct {
	// Allow the application to access account data. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	AccountData NullableBool `json:"account_data,omitempty"`
	// Allow the application to access bank statements. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	Statements NullableBool `json:"statements,omitempty"`
	// Allow the application to access tax documents. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	TaxDocuments NullableBool `json:"tax_documents,omitempty"`
}

// NewAccountProductAccessNullable instantiates a new AccountProductAccessNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountProductAccessNullable() *AccountProductAccessNullable {
	this := AccountProductAccessNullable{}
	var accountData bool = true
	this.AccountData = *NewNullableBool(&accountData)
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var taxDocuments bool = true
	this.TaxDocuments = *NewNullableBool(&taxDocuments)
	return &this
}

// NewAccountProductAccessNullableWithDefaults instantiates a new AccountProductAccessNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountProductAccessNullableWithDefaults() *AccountProductAccessNullable {
	this := AccountProductAccessNullable{}
	var accountData bool = true
	this.AccountData = *NewNullableBool(&accountData)
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var taxDocuments bool = true
	this.TaxDocuments = *NewNullableBool(&taxDocuments)
	return &this
}

// GetAccountData returns the AccountData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountProductAccessNullable) GetAccountData() bool {
	if o == nil || o.AccountData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccountData.Get()
}

// GetAccountDataOk returns a tuple with the AccountData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountProductAccessNullable) GetAccountDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountData.Get(), o.AccountData.IsSet()
}

// HasAccountData returns a boolean if a field has been set.
func (o *AccountProductAccessNullable) HasAccountData() bool {
	if o != nil && o.AccountData.IsSet() {
		return true
	}

	return false
}

// SetAccountData gets a reference to the given NullableBool and assigns it to the AccountData field.
func (o *AccountProductAccessNullable) SetAccountData(v bool) {
	o.AccountData.Set(&v)
}
// SetAccountDataNil sets the value for AccountData to be an explicit nil
func (o *AccountProductAccessNullable) SetAccountDataNil() {
	o.AccountData.Set(nil)
}

// UnsetAccountData ensures that no value is present for AccountData, not even an explicit nil
func (o *AccountProductAccessNullable) UnsetAccountData() {
	o.AccountData.Unset()
}

// GetStatements returns the Statements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountProductAccessNullable) GetStatements() bool {
	if o == nil || o.Statements.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Statements.Get()
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountProductAccessNullable) GetStatementsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Statements.Get(), o.Statements.IsSet()
}

// HasStatements returns a boolean if a field has been set.
func (o *AccountProductAccessNullable) HasStatements() bool {
	if o != nil && o.Statements.IsSet() {
		return true
	}

	return false
}

// SetStatements gets a reference to the given NullableBool and assigns it to the Statements field.
func (o *AccountProductAccessNullable) SetStatements(v bool) {
	o.Statements.Set(&v)
}
// SetStatementsNil sets the value for Statements to be an explicit nil
func (o *AccountProductAccessNullable) SetStatementsNil() {
	o.Statements.Set(nil)
}

// UnsetStatements ensures that no value is present for Statements, not even an explicit nil
func (o *AccountProductAccessNullable) UnsetStatements() {
	o.Statements.Unset()
}

// GetTaxDocuments returns the TaxDocuments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountProductAccessNullable) GetTaxDocuments() bool {
	if o == nil || o.TaxDocuments.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TaxDocuments.Get()
}

// GetTaxDocumentsOk returns a tuple with the TaxDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountProductAccessNullable) GetTaxDocumentsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxDocuments.Get(), o.TaxDocuments.IsSet()
}

// HasTaxDocuments returns a boolean if a field has been set.
func (o *AccountProductAccessNullable) HasTaxDocuments() bool {
	if o != nil && o.TaxDocuments.IsSet() {
		return true
	}

	return false
}

// SetTaxDocuments gets a reference to the given NullableBool and assigns it to the TaxDocuments field.
func (o *AccountProductAccessNullable) SetTaxDocuments(v bool) {
	o.TaxDocuments.Set(&v)
}
// SetTaxDocumentsNil sets the value for TaxDocuments to be an explicit nil
func (o *AccountProductAccessNullable) SetTaxDocumentsNil() {
	o.TaxDocuments.Set(nil)
}

// UnsetTaxDocuments ensures that no value is present for TaxDocuments, not even an explicit nil
func (o *AccountProductAccessNullable) UnsetTaxDocuments() {
	o.TaxDocuments.Unset()
}

func (o AccountProductAccessNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountData.IsSet() {
		toSerialize["account_data"] = o.AccountData.Get()
	}
	if o.Statements.IsSet() {
		toSerialize["statements"] = o.Statements.Get()
	}
	if o.TaxDocuments.IsSet() {
		toSerialize["tax_documents"] = o.TaxDocuments.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAccountProductAccessNullable struct {
	value *AccountProductAccessNullable
	isSet bool
}

func (v NullableAccountProductAccessNullable) Get() *AccountProductAccessNullable {
	return v.value
}

func (v *NullableAccountProductAccessNullable) Set(val *AccountProductAccessNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountProductAccessNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountProductAccessNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountProductAccessNullable(val *AccountProductAccessNullable) *NullableAccountProductAccessNullable {
	return &NullableAccountProductAccessNullable{value: val, isSet: true}
}

func (v NullableAccountProductAccessNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountProductAccessNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


