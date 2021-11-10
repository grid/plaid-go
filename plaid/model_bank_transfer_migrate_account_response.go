/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.46.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferMigrateAccountResponse Defines the response schema for `/bank_transfer/migrate_account`
type BankTransferMigrateAccountResponse struct {
	// The Plaid `access_token` for the newly created Item.
	AccessToken string `json:"access_token"`
	// The Plaid `account_id` for the newly created Item.
	AccountId string `json:"account_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferMigrateAccountResponse BankTransferMigrateAccountResponse

// NewBankTransferMigrateAccountResponse instantiates a new BankTransferMigrateAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferMigrateAccountResponse(accessToken string, accountId string, requestId string) *BankTransferMigrateAccountResponse {
	this := BankTransferMigrateAccountResponse{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	this.RequestId = requestId
	return &this
}

// NewBankTransferMigrateAccountResponseWithDefaults instantiates a new BankTransferMigrateAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferMigrateAccountResponseWithDefaults() *BankTransferMigrateAccountResponse {
	this := BankTransferMigrateAccountResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *BankTransferMigrateAccountResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *BankTransferMigrateAccountResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *BankTransferMigrateAccountResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountResponse) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BankTransferMigrateAccountResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferMigrateAccountResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferMigrateAccountResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferMigrateAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferMigrateAccountResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferMigrateAccountResponse := _BankTransferMigrateAccountResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferMigrateAccountResponse); err == nil {
		*o = BankTransferMigrateAccountResponse(varBankTransferMigrateAccountResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferMigrateAccountResponse struct {
	value *BankTransferMigrateAccountResponse
	isSet bool
}

func (v NullableBankTransferMigrateAccountResponse) Get() *BankTransferMigrateAccountResponse {
	return v.value
}

func (v *NullableBankTransferMigrateAccountResponse) Set(val *BankTransferMigrateAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferMigrateAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferMigrateAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferMigrateAccountResponse(val *BankTransferMigrateAccountResponse) *NullableBankTransferMigrateAccountResponse {
	return &NullableBankTransferMigrateAccountResponse{value: val, isSet: true}
}

func (v NullableBankTransferMigrateAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferMigrateAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


