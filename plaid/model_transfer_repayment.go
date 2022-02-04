/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// TransferRepayment A repayment is created automatically after one or more guaranteed transactions receive a return. If there are multiple eligible returns in a day, they are batched together into a single repayment.  Repayments are sent over ACH, with funds typically available on the next banking day.
type TransferRepayment struct {
	// Identifier of the repayment.
	RepaymentId string `json:"repayment_id"`
	// The datetime when the repayment occurred, in RFC 3339 format.
	Created time.Time `json:"created"`
	// Decimal amount of the repayment as it appears on your account ledger.
	Amount string `json:"amount"`
	// The currency of the repayment, e.g. \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _TransferRepayment TransferRepayment

// NewTransferRepayment instantiates a new TransferRepayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRepayment(repaymentId string, created time.Time, amount string, isoCurrencyCode string) *TransferRepayment {
	this := TransferRepayment{}
	this.RepaymentId = repaymentId
	this.Created = created
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewTransferRepaymentWithDefaults instantiates a new TransferRepayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRepaymentWithDefaults() *TransferRepayment {
	this := TransferRepayment{}
	return &this
}

// GetRepaymentId returns the RepaymentId field value
func (o *TransferRepayment) GetRepaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepaymentId
}

// GetRepaymentIdOk returns a tuple with the RepaymentId field value
// and a boolean to check if the value has been set.
func (o *TransferRepayment) GetRepaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentId, true
}

// SetRepaymentId sets field value
func (o *TransferRepayment) SetRepaymentId(v string) {
	o.RepaymentId = v
}

// GetCreated returns the Created field value
func (o *TransferRepayment) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransferRepayment) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransferRepayment) SetCreated(v time.Time) {
	o.Created = v
}

// GetAmount returns the Amount field value
func (o *TransferRepayment) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferRepayment) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferRepayment) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferRepayment) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferRepayment) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferRepayment) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o TransferRepayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["repayment_id"] = o.RepaymentId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRepayment) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRepayment := _TransferRepayment{}

	if err = json.Unmarshal(bytes, &varTransferRepayment); err == nil {
		*o = TransferRepayment(varTransferRepayment)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repayment_id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRepayment struct {
	value *TransferRepayment
	isSet bool
}

func (v NullableTransferRepayment) Get() *TransferRepayment {
	return v.value
}

func (v *NullableTransferRepayment) Set(val *TransferRepayment) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRepayment) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRepayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRepayment(val *TransferRepayment) *NullableTransferRepayment {
	return &NullableTransferRepayment{value: val, isSet: true}
}

func (v NullableTransferRepayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRepayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

