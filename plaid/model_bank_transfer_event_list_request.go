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

// BankTransferEventListRequest Defines the request schema for `/bank_transfer/event/list`
type BankTransferEventListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The start datetime of bank transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)
	StartDate NullableTime `json:"start_date,omitempty"`
	// The end datetime of bank transfers to list. This should be in RFC 3339 format (i.e. `2019-12-06T22:35:49Z`)
	EndDate NullableTime `json:"end_date,omitempty"`
	// Plaid’s unique identifier for a bank transfer.
	BankTransferId NullableString `json:"bank_transfer_id,omitempty"`
	// The account ID to get events for all transactions to/from an account.
	AccountId NullableString `json:"account_id,omitempty"`
	BankTransferType NullableBankTransferEventListBankTransferType `json:"bank_transfer_type,omitempty"`
	// Filter events by event type.
	EventTypes *[]BankTransferEventType `json:"event_types,omitempty"`
	// The maximum number of bank transfer events to return. If the number of events matching the above parameters is greater than `count`, the most recent events will be returned.
	Count NullableInt32 `json:"count,omitempty"`
	// The offset into the list of bank transfer events. When `count`=25 and `offset`=0, the first 25 events will be returned. When `count`=25 and `offset`=25, the next 25 bank transfer events will be returned.
	Offset NullableInt32 `json:"offset,omitempty"`
	// The origination account ID to get events for transfers from a specific origination account.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
	Direction NullableBankTransferEventListDirection `json:"direction,omitempty"`
}

// NewBankTransferEventListRequest instantiates a new BankTransferEventListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferEventListRequest() *BankTransferEventListRequest {
	this := BankTransferEventListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = *NewNullableInt32(&offset)
	return &this
}

// NewBankTransferEventListRequestWithDefaults instantiates a new BankTransferEventListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferEventListRequestWithDefaults() *BankTransferEventListRequest {
	this := BankTransferEventListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = *NewNullableInt32(&offset)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransferEventListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferEventListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransferEventListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BankTransferEventListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferEventListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BankTransferEventListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *BankTransferEventListRequest) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *BankTransferEventListRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *BankTransferEventListRequest) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *BankTransferEventListRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetBankTransferId returns the BankTransferId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetBankTransferId() string {
	if o == nil || o.BankTransferId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankTransferId.Get()
}

// GetBankTransferIdOk returns a tuple with the BankTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetBankTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankTransferId.Get(), o.BankTransferId.IsSet()
}

// HasBankTransferId returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasBankTransferId() bool {
	if o != nil && o.BankTransferId.IsSet() {
		return true
	}

	return false
}

// SetBankTransferId gets a reference to the given NullableString and assigns it to the BankTransferId field.
func (o *BankTransferEventListRequest) SetBankTransferId(v string) {
	o.BankTransferId.Set(&v)
}
// SetBankTransferIdNil sets the value for BankTransferId to be an explicit nil
func (o *BankTransferEventListRequest) SetBankTransferIdNil() {
	o.BankTransferId.Set(nil)
}

// UnsetBankTransferId ensures that no value is present for BankTransferId, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetBankTransferId() {
	o.BankTransferId.Unset()
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *BankTransferEventListRequest) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *BankTransferEventListRequest) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetBankTransferType returns the BankTransferType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetBankTransferType() BankTransferEventListBankTransferType {
	if o == nil || o.BankTransferType.Get() == nil {
		var ret BankTransferEventListBankTransferType
		return ret
	}
	return *o.BankTransferType.Get()
}

// GetBankTransferTypeOk returns a tuple with the BankTransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetBankTransferTypeOk() (*BankTransferEventListBankTransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankTransferType.Get(), o.BankTransferType.IsSet()
}

// HasBankTransferType returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasBankTransferType() bool {
	if o != nil && o.BankTransferType.IsSet() {
		return true
	}

	return false
}

// SetBankTransferType gets a reference to the given NullableBankTransferEventListBankTransferType and assigns it to the BankTransferType field.
func (o *BankTransferEventListRequest) SetBankTransferType(v BankTransferEventListBankTransferType) {
	o.BankTransferType.Set(&v)
}
// SetBankTransferTypeNil sets the value for BankTransferType to be an explicit nil
func (o *BankTransferEventListRequest) SetBankTransferTypeNil() {
	o.BankTransferType.Set(nil)
}

// UnsetBankTransferType ensures that no value is present for BankTransferType, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetBankTransferType() {
	o.BankTransferType.Unset()
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *BankTransferEventListRequest) GetEventTypes() []BankTransferEventType {
	if o == nil || o.EventTypes == nil {
		var ret []BankTransferEventType
		return ret
	}
	return *o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferEventListRequest) GetEventTypesOk() (*[]BankTransferEventType, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []BankTransferEventType and assigns it to the EventTypes field.
func (o *BankTransferEventListRequest) SetEventTypes(v []BankTransferEventType) {
	o.EventTypes = &v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *BankTransferEventListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *BankTransferEventListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetOffset() int32 {
	if o == nil || o.Offset.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *BankTransferEventListRequest) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *BankTransferEventListRequest) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetOffset() {
	o.Offset.Unset()
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *BankTransferEventListRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *BankTransferEventListRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferEventListRequest) GetDirection() BankTransferEventListDirection {
	if o == nil || o.Direction.Get() == nil {
		var ret BankTransferEventListDirection
		return ret
	}
	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEventListRequest) GetDirectionOk() (*BankTransferEventListDirection, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// HasDirection returns a boolean if a field has been set.
func (o *BankTransferEventListRequest) HasDirection() bool {
	if o != nil && o.Direction.IsSet() {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NullableBankTransferEventListDirection and assigns it to the Direction field.
func (o *BankTransferEventListRequest) SetDirection(v BankTransferEventListDirection) {
	o.Direction.Set(&v)
}
// SetDirectionNil sets the value for Direction to be an explicit nil
func (o *BankTransferEventListRequest) SetDirectionNil() {
	o.Direction.Set(nil)
}

// UnsetDirection ensures that no value is present for Direction, not even an explicit nil
func (o *BankTransferEventListRequest) UnsetDirection() {
	o.Direction.Unset()
}

func (o BankTransferEventListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.BankTransferId.IsSet() {
		toSerialize["bank_transfer_id"] = o.BankTransferId.Get()
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.BankTransferType.IsSet() {
		toSerialize["bank_transfer_type"] = o.BankTransferType.Get()
	}
	if o.EventTypes != nil {
		toSerialize["event_types"] = o.EventTypes
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if o.Direction.IsSet() {
		toSerialize["direction"] = o.Direction.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBankTransferEventListRequest struct {
	value *BankTransferEventListRequest
	isSet bool
}

func (v NullableBankTransferEventListRequest) Get() *BankTransferEventListRequest {
	return v.value
}

func (v *NullableBankTransferEventListRequest) Set(val *BankTransferEventListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEventListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEventListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEventListRequest(val *BankTransferEventListRequest) *NullableBankTransferEventListRequest {
	return &NullableBankTransferEventListRequest{value: val, isSet: true}
}

func (v NullableBankTransferEventListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEventListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


