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

// TransferIntentCreateRequest Defines the request schema for `/transfer/intent/create`
type TransferIntentCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId NullableString `json:"account_id,omitempty"`
	Mode TransferIntentCreateMode `json:"mode"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// A description for the underlying transfer. Maximum of 8 characters.
	Description string `json:"description"`
	AchClass ACHClass `json:"ach_class"`
	// Plaid’s unique identifier for the origination account for the intent. If not provided, the default account will be used.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
	User TransferUserInRequest `json:"user"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	// The currency of the transfer amount, e.g. \"USD\"
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
}

// NewTransferIntentCreateRequest instantiates a new TransferIntentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentCreateRequest(clientId string, secret string, mode TransferIntentCreateMode, amount string, description string, achClass ACHClass, user TransferUserInRequest) *TransferIntentCreateRequest {
	this := TransferIntentCreateRequest{}
	this.ClientId = clientId
	this.Secret = secret
	this.Mode = mode
	this.Amount = amount
	this.Description = description
	this.AchClass = achClass
	this.User = user
	return &this
}

// NewTransferIntentCreateRequestWithDefaults instantiates a new TransferIntentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentCreateRequestWithDefaults() *TransferIntentCreateRequest {
	this := TransferIntentCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *TransferIntentCreateRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *TransferIntentCreateRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *TransferIntentCreateRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *TransferIntentCreateRequest) SetSecret(v string) {
	o.Secret = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *TransferIntentCreateRequest) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *TransferIntentCreateRequest) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *TransferIntentCreateRequest) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetMode returns the Mode field value
func (o *TransferIntentCreateRequest) GetMode() TransferIntentCreateMode {
	if o == nil {
		var ret TransferIntentCreateMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetModeOk() (*TransferIntentCreateMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TransferIntentCreateRequest) SetMode(v TransferIntentCreateMode) {
	o.Mode = v
}

// GetAmount returns the Amount field value
func (o *TransferIntentCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferIntentCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *TransferIntentCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferIntentCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetAchClass returns the AchClass field value
func (o *TransferIntentCreateRequest) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *TransferIntentCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *TransferIntentCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *TransferIntentCreateRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *TransferIntentCreateRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

// GetUser returns the User field value
func (o *TransferIntentCreateRequest) GetUser() TransferUserInRequest {
	if o == nil {
		var ret TransferUserInRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetUserOk() (*TransferUserInRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferIntentCreateRequest) SetUser(v TransferUserInRequest) {
	o.User = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransferIntentCreateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *TransferIntentCreateRequest) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *TransferIntentCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

func (o TransferIntentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["ach_class"] = o.AchClass
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableTransferIntentCreateRequest struct {
	value *TransferIntentCreateRequest
	isSet bool
}

func (v NullableTransferIntentCreateRequest) Get() *TransferIntentCreateRequest {
	return v.value
}

func (v *NullableTransferIntentCreateRequest) Set(val *TransferIntentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentCreateRequest(val *TransferIntentCreateRequest) *NullableTransferIntentCreateRequest {
	return &NullableTransferIntentCreateRequest{value: val, isSet: true}
}

func (v NullableTransferIntentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


