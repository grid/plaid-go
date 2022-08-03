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

// TransferCreateRequest Defines the request schema for `/transfer/create`
type TransferCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Deprecated. `authorization_id` is now used as idempotency instead.  A random key provided by the client, per unique transfer. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a transfer fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single transfer is created.
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// The Plaid `access_token` for the account that will be debited or credited.
	AccessToken *string `json:"access_token,omitempty"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId *string `json:"account_id,omitempty"`
	// Plaid’s unique identifier for a transfer authorization. This parameter also serves the purpose of acting as an idempotency identifier.
	AuthorizationId string `json:"authorization_id"`
	Type TransferType `json:"type"`
	Network TransferNetwork `json:"network"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// The transfer description. Maximum of 10 characters.
	Description string `json:"description"`
	AchClass ACHClass `json:"ach_class"`
	User TransferUserInRequest `json:"user"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	// Plaid’s unique identifier for the origination account for this transfer. If you have more than one origination account, this value must be specified. Otherwise, this field should be left blank.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
	// The currency of the transfer amount. The default value is \"USD\".
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
	// Plaid’s unique identifier for a payment profile.
	PaymentProfileId *string `json:"payment_profile_id,omitempty"`
}

// NewTransferCreateRequest instantiates a new TransferCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCreateRequest(authorizationId string, type_ TransferType, network TransferNetwork, amount string, description string, achClass ACHClass, user TransferUserInRequest) *TransferCreateRequest {
	this := TransferCreateRequest{}
	this.AuthorizationId = authorizationId
	this.Type = type_
	this.Network = network
	this.Amount = amount
	this.Description = description
	this.AchClass = achClass
	this.User = user
	return &this
}

// NewTransferCreateRequestWithDefaults instantiates a new TransferCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCreateRequestWithDefaults() *TransferCreateRequest {
	this := TransferCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetIdempotencyKey() string {
	if o == nil || o.IdempotencyKey == nil {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || o.IdempotencyKey == nil {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasIdempotencyKey() bool {
	if o != nil && o.IdempotencyKey != nil {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *TransferCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TransferCreateRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransferCreateRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAuthorizationId returns the AuthorizationId field value
func (o *TransferCreateRequest) GetAuthorizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationId
}

// GetAuthorizationIdOk returns a tuple with the AuthorizationId field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetAuthorizationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthorizationId, true
}

// SetAuthorizationId sets field value
func (o *TransferCreateRequest) SetAuthorizationId(v string) {
	o.AuthorizationId = v
}

// GetType returns the Type field value
func (o *TransferCreateRequest) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferCreateRequest) SetType(v TransferType) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *TransferCreateRequest) GetNetwork() TransferNetwork {
	if o == nil {
		var ret TransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetNetworkOk() (*TransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransferCreateRequest) SetNetwork(v TransferNetwork) {
	o.Network = v
}

// GetAmount returns the Amount field value
func (o *TransferCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *TransferCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetAchClass returns the AchClass field value
func (o *TransferCreateRequest) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *TransferCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetUser returns the User field value
func (o *TransferCreateRequest) GetUser() TransferUserInRequest {
	if o == nil {
		var ret TransferUserInRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetUserOk() (*TransferUserInRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferCreateRequest) SetUser(v TransferUserInRequest) {
	o.User = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferCreateRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferCreateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransferCreateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *TransferCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *TransferCreateRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *TransferCreateRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *TransferCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

// GetPaymentProfileId returns the PaymentProfileId field value if set, zero value otherwise.
func (o *TransferCreateRequest) GetPaymentProfileId() string {
	if o == nil || o.PaymentProfileId == nil {
		var ret string
		return ret
	}
	return *o.PaymentProfileId
}

// GetPaymentProfileIdOk returns a tuple with the PaymentProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCreateRequest) GetPaymentProfileIdOk() (*string, bool) {
	if o == nil || o.PaymentProfileId == nil {
		return nil, false
	}
	return o.PaymentProfileId, true
}

// HasPaymentProfileId returns a boolean if a field has been set.
func (o *TransferCreateRequest) HasPaymentProfileId() bool {
	if o != nil && o.PaymentProfileId != nil {
		return true
	}

	return false
}

// SetPaymentProfileId gets a reference to the given string and assigns it to the PaymentProfileId field.
func (o *TransferCreateRequest) SetPaymentProfileId(v string) {
	o.PaymentProfileId = &v
}

func (o TransferCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.IdempotencyKey != nil {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["authorization_id"] = o.AuthorizationId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["network"] = o.Network
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
	if true {
		toSerialize["user"] = o.User
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if o.PaymentProfileId != nil {
		toSerialize["payment_profile_id"] = o.PaymentProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableTransferCreateRequest struct {
	value *TransferCreateRequest
	isSet bool
}

func (v NullableTransferCreateRequest) Get() *TransferCreateRequest {
	return v.value
}

func (v *NullableTransferCreateRequest) Set(val *TransferCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCreateRequest(val *TransferCreateRequest) *NullableTransferCreateRequest {
	return &NullableTransferCreateRequest{value: val, isSet: true}
}

func (v NullableTransferCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


