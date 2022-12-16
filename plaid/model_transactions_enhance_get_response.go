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

// TransactionsEnhanceGetResponse TransactionsEnhanceGetResponse defines the response schema for `/beta/transactions/v1/enhance`.
type TransactionsEnhanceGetResponse struct {
	// An array of enhanced transactions.
	EnhancedTransactions []ClientProvidedEnhancedTransaction `json:"enhanced_transactions"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsEnhanceGetResponse TransactionsEnhanceGetResponse

// NewTransactionsEnhanceGetResponse instantiates a new TransactionsEnhanceGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsEnhanceGetResponse(enhancedTransactions []ClientProvidedEnhancedTransaction) *TransactionsEnhanceGetResponse {
	this := TransactionsEnhanceGetResponse{}
	this.EnhancedTransactions = enhancedTransactions
	return &this
}

// NewTransactionsEnhanceGetResponseWithDefaults instantiates a new TransactionsEnhanceGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsEnhanceGetResponseWithDefaults() *TransactionsEnhanceGetResponse {
	this := TransactionsEnhanceGetResponse{}
	return &this
}

// GetEnhancedTransactions returns the EnhancedTransactions field value
func (o *TransactionsEnhanceGetResponse) GetEnhancedTransactions() []ClientProvidedEnhancedTransaction {
	if o == nil {
		var ret []ClientProvidedEnhancedTransaction
		return ret
	}

	return o.EnhancedTransactions
}

// GetEnhancedTransactionsOk returns a tuple with the EnhancedTransactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsEnhanceGetResponse) GetEnhancedTransactionsOk() (*[]ClientProvidedEnhancedTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnhancedTransactions, true
}

// SetEnhancedTransactions sets field value
func (o *TransactionsEnhanceGetResponse) SetEnhancedTransactions(v []ClientProvidedEnhancedTransaction) {
	o.EnhancedTransactions = v
}

func (o TransactionsEnhanceGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enhanced_transactions"] = o.EnhancedTransactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsEnhanceGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsEnhanceGetResponse := _TransactionsEnhanceGetResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsEnhanceGetResponse); err == nil {
		*o = TransactionsEnhanceGetResponse(varTransactionsEnhanceGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enhanced_transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsEnhanceGetResponse struct {
	value *TransactionsEnhanceGetResponse
	isSet bool
}

func (v NullableTransactionsEnhanceGetResponse) Get() *TransactionsEnhanceGetResponse {
	return v.value
}

func (v *NullableTransactionsEnhanceGetResponse) Set(val *TransactionsEnhanceGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsEnhanceGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsEnhanceGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsEnhanceGetResponse(val *TransactionsEnhanceGetResponse) *NullableTransactionsEnhanceGetResponse {
	return &NullableTransactionsEnhanceGetResponse{value: val, isSet: true}
}

func (v NullableTransactionsEnhanceGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsEnhanceGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


