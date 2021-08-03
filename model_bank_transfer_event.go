/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.12
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferEvent Represents an event in the Bank Transfers API.
type BankTransferEvent struct {
	// Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers.
	EventId int32 `json:"event_id"`
	// The datetime when this event occurred. This will be of the form `2006-01-02T15:04:05Z`.
	Timestamp string `json:"timestamp"`
	EventType BankTransferEventType `json:"event_type"`
	// The account ID associated with the bank transfer.
	AccountId string `json:"account_id"`
	// Plaid’s unique identifier for a bank transfer.
	BankTransferId string `json:"bank_transfer_id"`
	// The ID of the origination account that this balance belongs to.
	OriginationAccountId NullableString `json:"origination_account_id"`
	BankTransferType BankTransferType `json:"bank_transfer_type"`
	// The bank transfer amount.
	BankTransferAmount string `json:"bank_transfer_amount"`
	// The currency of the bank transfer amount.
	BankTransferIsoCurrencyCode string `json:"bank_transfer_iso_currency_code"`
	FailureReason NullableBankTransferFailure `json:"failure_reason"`
	Direction NullableBankTransferDirection `json:"direction"`
	ReceiverDetails NullableBankTransferReceiverDetails `json:"receiver_details"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferEvent BankTransferEvent

// NewBankTransferEvent instantiates a new BankTransferEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferEvent(eventId int32, timestamp string, eventType BankTransferEventType, accountId string, bankTransferId string, originationAccountId NullableString, bankTransferType BankTransferType, bankTransferAmount string, bankTransferIsoCurrencyCode string, failureReason NullableBankTransferFailure, direction NullableBankTransferDirection, receiverDetails NullableBankTransferReceiverDetails) *BankTransferEvent {
	this := BankTransferEvent{}
	this.EventId = eventId
	this.Timestamp = timestamp
	this.EventType = eventType
	this.AccountId = accountId
	this.BankTransferId = bankTransferId
	this.OriginationAccountId = originationAccountId
	this.BankTransferType = bankTransferType
	this.BankTransferAmount = bankTransferAmount
	this.BankTransferIsoCurrencyCode = bankTransferIsoCurrencyCode
	this.FailureReason = failureReason
	this.Direction = direction
	this.ReceiverDetails = receiverDetails
	return &this
}

// NewBankTransferEventWithDefaults instantiates a new BankTransferEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferEventWithDefaults() *BankTransferEvent {
	this := BankTransferEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *BankTransferEvent) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetEventIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *BankTransferEvent) SetEventId(v int32) {
	o.EventId = v
}

// GetTimestamp returns the Timestamp field value
func (o *BankTransferEvent) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BankTransferEvent) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetEventType returns the EventType field value
func (o *BankTransferEvent) GetEventType() BankTransferEventType {
	if o == nil {
		var ret BankTransferEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetEventTypeOk() (*BankTransferEventType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *BankTransferEvent) SetEventType(v BankTransferEventType) {
	o.EventType = v
}

// GetAccountId returns the AccountId field value
func (o *BankTransferEvent) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BankTransferEvent) SetAccountId(v string) {
	o.AccountId = v
}

// GetBankTransferId returns the BankTransferId field value
func (o *BankTransferEvent) GetBankTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankTransferId
}

// GetBankTransferIdOk returns a tuple with the BankTransferId field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetBankTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferId, true
}

// SetBankTransferId sets field value
func (o *BankTransferEvent) SetBankTransferId(v string) {
	o.BankTransferId = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankTransferEvent) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEvent) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// SetOriginationAccountId sets field value
func (o *BankTransferEvent) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}

// GetBankTransferType returns the BankTransferType field value
func (o *BankTransferEvent) GetBankTransferType() BankTransferType {
	if o == nil {
		var ret BankTransferType
		return ret
	}

	return o.BankTransferType
}

// GetBankTransferTypeOk returns a tuple with the BankTransferType field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetBankTransferTypeOk() (*BankTransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferType, true
}

// SetBankTransferType sets field value
func (o *BankTransferEvent) SetBankTransferType(v BankTransferType) {
	o.BankTransferType = v
}

// GetBankTransferAmount returns the BankTransferAmount field value
func (o *BankTransferEvent) GetBankTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankTransferAmount
}

// GetBankTransferAmountOk returns a tuple with the BankTransferAmount field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetBankTransferAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferAmount, true
}

// SetBankTransferAmount sets field value
func (o *BankTransferEvent) SetBankTransferAmount(v string) {
	o.BankTransferAmount = v
}

// GetBankTransferIsoCurrencyCode returns the BankTransferIsoCurrencyCode field value
func (o *BankTransferEvent) GetBankTransferIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankTransferIsoCurrencyCode
}

// GetBankTransferIsoCurrencyCodeOk returns a tuple with the BankTransferIsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *BankTransferEvent) GetBankTransferIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferIsoCurrencyCode, true
}

// SetBankTransferIsoCurrencyCode sets field value
func (o *BankTransferEvent) SetBankTransferIsoCurrencyCode(v string) {
	o.BankTransferIsoCurrencyCode = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for BankTransferFailure will be returned
func (o *BankTransferEvent) GetFailureReason() BankTransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret BankTransferFailure
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEvent) GetFailureReasonOk() (*BankTransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *BankTransferEvent) SetFailureReason(v BankTransferFailure) {
	o.FailureReason.Set(&v)
}

// GetDirection returns the Direction field value
// If the value is explicit nil, the zero value for BankTransferDirection will be returned
func (o *BankTransferEvent) GetDirection() BankTransferDirection {
	if o == nil || o.Direction.Get() == nil {
		var ret BankTransferDirection
		return ret
	}

	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEvent) GetDirectionOk() (*BankTransferDirection, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// SetDirection sets field value
func (o *BankTransferEvent) SetDirection(v BankTransferDirection) {
	o.Direction.Set(&v)
}

// GetReceiverDetails returns the ReceiverDetails field value
// If the value is explicit nil, the zero value for BankTransferReceiverDetails will be returned
func (o *BankTransferEvent) GetReceiverDetails() BankTransferReceiverDetails {
	if o == nil || o.ReceiverDetails.Get() == nil {
		var ret BankTransferReceiverDetails
		return ret
	}

	return *o.ReceiverDetails.Get()
}

// GetReceiverDetailsOk returns a tuple with the ReceiverDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferEvent) GetReceiverDetailsOk() (*BankTransferReceiverDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReceiverDetails.Get(), o.ReceiverDetails.IsSet()
}

// SetReceiverDetails sets field value
func (o *BankTransferEvent) SetReceiverDetails(v BankTransferReceiverDetails) {
	o.ReceiverDetails.Set(&v)
}

func (o BankTransferEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["bank_transfer_id"] = o.BankTransferId
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if true {
		toSerialize["bank_transfer_type"] = o.BankTransferType
	}
	if true {
		toSerialize["bank_transfer_amount"] = o.BankTransferAmount
	}
	if true {
		toSerialize["bank_transfer_iso_currency_code"] = o.BankTransferIsoCurrencyCode
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if true {
		toSerialize["direction"] = o.Direction.Get()
	}
	if true {
		toSerialize["receiver_details"] = o.ReceiverDetails.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferEvent) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferEvent := _BankTransferEvent{}

	if err = json.Unmarshal(bytes, &varBankTransferEvent); err == nil {
		*o = BankTransferEvent(varBankTransferEvent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "event_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "event_type")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "bank_transfer_id")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "bank_transfer_type")
		delete(additionalProperties, "bank_transfer_amount")
		delete(additionalProperties, "bank_transfer_iso_currency_code")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "receiver_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferEvent struct {
	value *BankTransferEvent
	isSet bool
}

func (v NullableBankTransferEvent) Get() *BankTransferEvent {
	return v.value
}

func (v *NullableBankTransferEvent) Set(val *BankTransferEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEvent(val *BankTransferEvent) *NullableBankTransferEvent {
	return &NullableBankTransferEvent{value: val, isSet: true}
}

func (v NullableBankTransferEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

