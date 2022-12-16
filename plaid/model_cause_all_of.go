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

// CauseAllOf struct for CauseAllOf
type CauseAllOf struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId *string `json:"item_id,omitempty"`
}

// NewCauseAllOf instantiates a new CauseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCauseAllOf() *CauseAllOf {
	this := CauseAllOf{}
	return &this
}

// NewCauseAllOfWithDefaults instantiates a new CauseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCauseAllOfWithDefaults() *CauseAllOf {
	this := CauseAllOf{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *CauseAllOf) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CauseAllOf) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *CauseAllOf) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *CauseAllOf) SetItemId(v string) {
	o.ItemId = &v
}

func (o CauseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	return json.Marshal(toSerialize)
}

type NullableCauseAllOf struct {
	value *CauseAllOf
	isSet bool
}

func (v NullableCauseAllOf) Get() *CauseAllOf {
	return v.value
}

func (v *NullableCauseAllOf) Set(val *CauseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCauseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCauseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCauseAllOf(val *CauseAllOf) *NullableCauseAllOf {
	return &NullableCauseAllOf{value: val, isSet: true}
}

func (v NullableCauseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCauseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


