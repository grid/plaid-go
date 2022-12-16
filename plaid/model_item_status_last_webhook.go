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
	"time"
)

// ItemStatusLastWebhook Information about the last webhook fired for the Item.
type ItemStatusLastWebhook struct {
	// [ISO 8601](https://wikipedia.org/wiki/ISO_8601) timestamp of when the webhook was fired. 
	SentAt NullableTime `json:"sent_at,omitempty"`
	// The last webhook code sent.
	CodeSent NullableString `json:"code_sent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemStatusLastWebhook ItemStatusLastWebhook

// NewItemStatusLastWebhook instantiates a new ItemStatusLastWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemStatusLastWebhook() *ItemStatusLastWebhook {
	this := ItemStatusLastWebhook{}
	return &this
}

// NewItemStatusLastWebhookWithDefaults instantiates a new ItemStatusLastWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemStatusLastWebhookWithDefaults() *ItemStatusLastWebhook {
	this := ItemStatusLastWebhook{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusLastWebhook) GetSentAt() time.Time {
	if o == nil || o.SentAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusLastWebhook) GetSentAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *ItemStatusLastWebhook) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableTime and assigns it to the SentAt field.
func (o *ItemStatusLastWebhook) SetSentAt(v time.Time) {
	o.SentAt.Set(&v)
}
// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *ItemStatusLastWebhook) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *ItemStatusLastWebhook) UnsetSentAt() {
	o.SentAt.Unset()
}

// GetCodeSent returns the CodeSent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusLastWebhook) GetCodeSent() string {
	if o == nil || o.CodeSent.Get() == nil {
		var ret string
		return ret
	}
	return *o.CodeSent.Get()
}

// GetCodeSentOk returns a tuple with the CodeSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusLastWebhook) GetCodeSentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CodeSent.Get(), o.CodeSent.IsSet()
}

// HasCodeSent returns a boolean if a field has been set.
func (o *ItemStatusLastWebhook) HasCodeSent() bool {
	if o != nil && o.CodeSent.IsSet() {
		return true
	}

	return false
}

// SetCodeSent gets a reference to the given NullableString and assigns it to the CodeSent field.
func (o *ItemStatusLastWebhook) SetCodeSent(v string) {
	o.CodeSent.Set(&v)
}
// SetCodeSentNil sets the value for CodeSent to be an explicit nil
func (o *ItemStatusLastWebhook) SetCodeSentNil() {
	o.CodeSent.Set(nil)
}

// UnsetCodeSent ensures that no value is present for CodeSent, not even an explicit nil
func (o *ItemStatusLastWebhook) UnsetCodeSent() {
	o.CodeSent.Unset()
}

func (o ItemStatusLastWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SentAt.IsSet() {
		toSerialize["sent_at"] = o.SentAt.Get()
	}
	if o.CodeSent.IsSet() {
		toSerialize["code_sent"] = o.CodeSent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemStatusLastWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varItemStatusLastWebhook := _ItemStatusLastWebhook{}

	if err = json.Unmarshal(bytes, &varItemStatusLastWebhook); err == nil {
		*o = ItemStatusLastWebhook(varItemStatusLastWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sent_at")
		delete(additionalProperties, "code_sent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemStatusLastWebhook struct {
	value *ItemStatusLastWebhook
	isSet bool
}

func (v NullableItemStatusLastWebhook) Get() *ItemStatusLastWebhook {
	return v.value
}

func (v *NullableItemStatusLastWebhook) Set(val *ItemStatusLastWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStatusLastWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStatusLastWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStatusLastWebhook(val *ItemStatusLastWebhook) *NullableItemStatusLastWebhook {
	return &NullableItemStatusLastWebhook{value: val, isSet: true}
}

func (v NullableItemStatusLastWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStatusLastWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


