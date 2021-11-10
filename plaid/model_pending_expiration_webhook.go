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
	"time"
)

// PendingExpirationWebhook Fired when an Item’s access consent is expiring in 7 days. Some Items have explicit expiration times and we try to relay this when possible to reduce service disruption. This can be resolved by having the user go through Link’s update mode.
type PendingExpirationWebhook struct {
	// `ITEM`
	WebhookType string `json:"webhook_type"`
	// `PENDING_EXPIRATION`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// The date and time at which the Item's access consent will expire, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format
	ConsentExpirationTime time.Time `json:"consent_expiration_time"`
	AdditionalProperties map[string]interface{}
}

type _PendingExpirationWebhook PendingExpirationWebhook

// NewPendingExpirationWebhook instantiates a new PendingExpirationWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingExpirationWebhook(webhookType string, webhookCode string, itemId string, consentExpirationTime time.Time) *PendingExpirationWebhook {
	this := PendingExpirationWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.ConsentExpirationTime = consentExpirationTime
	return &this
}

// NewPendingExpirationWebhookWithDefaults instantiates a new PendingExpirationWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingExpirationWebhookWithDefaults() *PendingExpirationWebhook {
	this := PendingExpirationWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *PendingExpirationWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *PendingExpirationWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *PendingExpirationWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *PendingExpirationWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *PendingExpirationWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *PendingExpirationWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *PendingExpirationWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *PendingExpirationWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *PendingExpirationWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetConsentExpirationTime returns the ConsentExpirationTime field value
func (o *PendingExpirationWebhook) GetConsentExpirationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ConsentExpirationTime
}

// GetConsentExpirationTimeOk returns a tuple with the ConsentExpirationTime field value
// and a boolean to check if the value has been set.
func (o *PendingExpirationWebhook) GetConsentExpirationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentExpirationTime, true
}

// SetConsentExpirationTime sets field value
func (o *PendingExpirationWebhook) SetConsentExpirationTime(v time.Time) {
	o.ConsentExpirationTime = v
}

func (o PendingExpirationWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["consent_expiration_time"] = o.ConsentExpirationTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PendingExpirationWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varPendingExpirationWebhook := _PendingExpirationWebhook{}

	if err = json.Unmarshal(bytes, &varPendingExpirationWebhook); err == nil {
		*o = PendingExpirationWebhook(varPendingExpirationWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "consent_expiration_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingExpirationWebhook struct {
	value *PendingExpirationWebhook
	isSet bool
}

func (v NullablePendingExpirationWebhook) Get() *PendingExpirationWebhook {
	return v.value
}

func (v *NullablePendingExpirationWebhook) Set(val *PendingExpirationWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingExpirationWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingExpirationWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingExpirationWebhook(val *PendingExpirationWebhook) *NullablePendingExpirationWebhook {
	return &NullablePendingExpirationWebhook{value: val, isSet: true}
}

func (v NullablePendingExpirationWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingExpirationWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


