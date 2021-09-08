/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.31.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferBalanceGetResponse Defines the response schema for `/bank_transfer/balance/get`
type BankTransferBalanceGetResponse struct {
	Balance BankTransferBalance `json:"balance"`
	// The ID of the origination account that this balance belongs to.
	OriginationAccountId NullableString `json:"origination_account_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferBalanceGetResponse BankTransferBalanceGetResponse

// NewBankTransferBalanceGetResponse instantiates a new BankTransferBalanceGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferBalanceGetResponse(balance BankTransferBalance, originationAccountId NullableString, requestId string) *BankTransferBalanceGetResponse {
	this := BankTransferBalanceGetResponse{}
	this.Balance = balance
	this.OriginationAccountId = originationAccountId
	this.RequestId = requestId
	return &this
}

// NewBankTransferBalanceGetResponseWithDefaults instantiates a new BankTransferBalanceGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferBalanceGetResponseWithDefaults() *BankTransferBalanceGetResponse {
	this := BankTransferBalanceGetResponse{}
	return &this
}

// GetBalance returns the Balance field value
func (o *BankTransferBalanceGetResponse) GetBalance() BankTransferBalance {
	if o == nil {
		var ret BankTransferBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *BankTransferBalanceGetResponse) GetBalanceOk() (*BankTransferBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *BankTransferBalanceGetResponse) SetBalance(v BankTransferBalance) {
	o.Balance = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankTransferBalanceGetResponse) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferBalanceGetResponse) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// SetOriginationAccountId sets field value
func (o *BankTransferBalanceGetResponse) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *BankTransferBalanceGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferBalanceGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferBalanceGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferBalanceGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferBalanceGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferBalanceGetResponse := _BankTransferBalanceGetResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferBalanceGetResponse); err == nil {
		*o = BankTransferBalanceGetResponse(varBankTransferBalanceGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "balance")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferBalanceGetResponse struct {
	value *BankTransferBalanceGetResponse
	isSet bool
}

func (v NullableBankTransferBalanceGetResponse) Get() *BankTransferBalanceGetResponse {
	return v.value
}

func (v *NullableBankTransferBalanceGetResponse) Set(val *BankTransferBalanceGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferBalanceGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferBalanceGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferBalanceGetResponse(val *BankTransferBalanceGetResponse) *NullableBankTransferBalanceGetResponse {
	return &NullableBankTransferBalanceGetResponse{value: val, isSet: true}
}

func (v NullableBankTransferBalanceGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferBalanceGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


