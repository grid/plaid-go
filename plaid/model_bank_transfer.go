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
	"time"
)

// BankTransfer Represents a bank transfer within the Bank Transfers API.
type BankTransfer struct {
	// Plaid’s unique identifier for a bank transfer.
	Id string `json:"id"`
	AchClass ACHClass `json:"ach_class"`
	// The account ID that should be credited/debited for this bank transfer.
	AccountId string `json:"account_id"`
	Type BankTransferType `json:"type"`
	User BankTransferUser `json:"user"`
	// The amount of the bank transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// The currency of the transfer amount, e.g. \"USD\"
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The description of the transfer.
	Description string `json:"description"`
	// The datetime when this bank transfer was created. This will be of the form `2006-01-02T15:04:05Z`
	Created time.Time `json:"created"`
	Status BankTransferStatus `json:"status"`
	Network BankTransferNetwork `json:"network"`
	// When `true`, you can still cancel this bank transfer.
	Cancellable bool `json:"cancellable"`
	FailureReason NullableBankTransferFailure `json:"failure_reason"`
	// A string containing the custom tag provided by the client in the create request. Will be null if not provided.
	CustomTag NullableString `json:"custom_tag"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata"`
	// Plaid’s unique identifier for the origination account that was used for this transfer.
	OriginationAccountId string `json:"origination_account_id"`
	Direction NullableBankTransferDirection `json:"direction"`
	AdditionalProperties map[string]interface{}
}

type _BankTransfer BankTransfer

// NewBankTransfer instantiates a new BankTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransfer(id string, achClass ACHClass, accountId string, type_ BankTransferType, user BankTransferUser, amount string, isoCurrencyCode string, description string, created time.Time, status BankTransferStatus, network BankTransferNetwork, cancellable bool, failureReason NullableBankTransferFailure, customTag NullableString, metadata map[string]string, originationAccountId string, direction NullableBankTransferDirection) *BankTransfer {
	this := BankTransfer{}
	this.Id = id
	this.AchClass = achClass
	this.AccountId = accountId
	this.Type = type_
	this.User = user
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.Description = description
	this.Created = created
	this.Status = status
	this.Network = network
	this.Cancellable = cancellable
	this.FailureReason = failureReason
	this.CustomTag = customTag
	this.Metadata = metadata
	this.OriginationAccountId = originationAccountId
	this.Direction = direction
	return &this
}

// NewBankTransferWithDefaults instantiates a new BankTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferWithDefaults() *BankTransfer {
	this := BankTransfer{}
	return &this
}

// GetId returns the Id field value
func (o *BankTransfer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BankTransfer) SetId(v string) {
	o.Id = v
}

// GetAchClass returns the AchClass field value
func (o *BankTransfer) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *BankTransfer) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetAccountId returns the AccountId field value
func (o *BankTransfer) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BankTransfer) SetAccountId(v string) {
	o.AccountId = v
}

// GetType returns the Type field value
func (o *BankTransfer) GetType() BankTransferType {
	if o == nil {
		var ret BankTransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetTypeOk() (*BankTransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BankTransfer) SetType(v BankTransferType) {
	o.Type = v
}

// GetUser returns the User field value
func (o *BankTransfer) GetUser() BankTransferUser {
	if o == nil {
		var ret BankTransferUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetUserOk() (*BankTransferUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BankTransfer) SetUser(v BankTransferUser) {
	o.User = v
}

// GetAmount returns the Amount field value
func (o *BankTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BankTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *BankTransfer) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *BankTransfer) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetDescription returns the Description field value
func (o *BankTransfer) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BankTransfer) SetDescription(v string) {
	o.Description = v
}

// GetCreated returns the Created field value
func (o *BankTransfer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *BankTransfer) SetCreated(v time.Time) {
	o.Created = v
}

// GetStatus returns the Status field value
func (o *BankTransfer) GetStatus() BankTransferStatus {
	if o == nil {
		var ret BankTransferStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetStatusOk() (*BankTransferStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BankTransfer) SetStatus(v BankTransferStatus) {
	o.Status = v
}

// GetNetwork returns the Network field value
func (o *BankTransfer) GetNetwork() BankTransferNetwork {
	if o == nil {
		var ret BankTransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetNetworkOk() (*BankTransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *BankTransfer) SetNetwork(v BankTransferNetwork) {
	o.Network = v
}

// GetCancellable returns the Cancellable field value
func (o *BankTransfer) GetCancellable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cancellable
}

// GetCancellableOk returns a tuple with the Cancellable field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetCancellableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cancellable, true
}

// SetCancellable sets field value
func (o *BankTransfer) SetCancellable(v bool) {
	o.Cancellable = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for BankTransferFailure will be returned
func (o *BankTransfer) GetFailureReason() BankTransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret BankTransferFailure
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransfer) GetFailureReasonOk() (*BankTransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *BankTransfer) SetFailureReason(v BankTransferFailure) {
	o.FailureReason.Set(&v)
}

// GetCustomTag returns the CustomTag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankTransfer) GetCustomTag() string {
	if o == nil || o.CustomTag.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomTag.Get()
}

// GetCustomTagOk returns a tuple with the CustomTag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransfer) GetCustomTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomTag.Get(), o.CustomTag.IsSet()
}

// SetCustomTag sets field value
func (o *BankTransfer) SetCustomTag(v string) {
	o.CustomTag.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *BankTransfer) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransfer) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *BankTransfer) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *BankTransfer) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *BankTransfer) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetDirection returns the Direction field value
// If the value is explicit nil, the zero value for BankTransferDirection will be returned
func (o *BankTransfer) GetDirection() BankTransferDirection {
	if o == nil || o.Direction.Get() == nil {
		var ret BankTransferDirection
		return ret
	}

	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransfer) GetDirectionOk() (*BankTransferDirection, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// SetDirection sets field value
func (o *BankTransfer) SetDirection(v BankTransferDirection) {
	o.Direction.Set(&v)
}

func (o BankTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["cancellable"] = o.Cancellable
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if true {
		toSerialize["custom_tag"] = o.CustomTag.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["direction"] = o.Direction.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransfer) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransfer := _BankTransfer{}

	if err = json.Unmarshal(bytes, &varBankTransfer); err == nil {
		*o = BankTransfer(varBankTransfer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "status")
		delete(additionalProperties, "network")
		delete(additionalProperties, "cancellable")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "custom_tag")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "direction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransfer struct {
	value *BankTransfer
	isSet bool
}

func (v NullableBankTransfer) Get() *BankTransfer {
	return v.value
}

func (v *NullableBankTransfer) Set(val *BankTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransfer(val *BankTransfer) *NullableBankTransfer {
	return &NullableBankTransfer{value: val, isSet: true}
}

func (v NullableBankTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


