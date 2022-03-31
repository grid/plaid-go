/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.94.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PayrollItem An object containing information about the payroll item.
type PayrollItem struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	PayrollIncome []PayrollIncomeObject `json:"payroll_income"`
}

// NewPayrollItem instantiates a new PayrollItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayrollItem(itemId string, payrollIncome []PayrollIncomeObject) *PayrollItem {
	this := PayrollItem{}
	this.ItemId = itemId
	this.PayrollIncome = payrollIncome
	return &this
}

// NewPayrollItemWithDefaults instantiates a new PayrollItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayrollItemWithDefaults() *PayrollItem {
	this := PayrollItem{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *PayrollItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *PayrollItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *PayrollItem) SetItemId(v string) {
	o.ItemId = v
}

// GetPayrollIncome returns the PayrollIncome field value
func (o *PayrollItem) GetPayrollIncome() []PayrollIncomeObject {
	if o == nil {
		var ret []PayrollIncomeObject
		return ret
	}

	return o.PayrollIncome
}

// GetPayrollIncomeOk returns a tuple with the PayrollIncome field value
// and a boolean to check if the value has been set.
func (o *PayrollItem) GetPayrollIncomeOk() (*[]PayrollIncomeObject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayrollIncome, true
}

// SetPayrollIncome sets field value
func (o *PayrollItem) SetPayrollIncome(v []PayrollIncomeObject) {
	o.PayrollIncome = v
}

func (o PayrollItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["payroll_income"] = o.PayrollIncome
	}
	return json.Marshal(toSerialize)
}

type NullablePayrollItem struct {
	value *PayrollItem
	isSet bool
}

func (v NullablePayrollItem) Get() *PayrollItem {
	return v.value
}

func (v *NullablePayrollItem) Set(val *PayrollItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePayrollItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePayrollItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayrollItem(val *PayrollItem) *NullablePayrollItem {
	return &NullablePayrollItem{value: val, isSet: true}
}

func (v NullablePayrollItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayrollItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

