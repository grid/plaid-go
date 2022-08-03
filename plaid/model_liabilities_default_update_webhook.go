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

// LiabilitiesDefaultUpdateWebhook The webhook of type `LIABILITIES` and code `DEFAULT_UPDATE` will be fired when new or updated liabilities have been detected on a liabilities item.
type LiabilitiesDefaultUpdateWebhook struct {
	// `LIABILITIES`
	WebhookType string `json:"webhook_type"`
	// `DEFAULT_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Error NullablePlaidError `json:"error"`
	// An array of `account_id`'s for accounts that contain new liabilities.'
	AccountIdsWithNewLiabilities []string `json:"account_ids_with_new_liabilities"`
	// An object with keys of `account_id`'s that are mapped to their respective liabilities fields that changed.  Example: `{ \"XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58\": [\"past_amount_due\"] }` 
	AccountIdsWithUpdatedLiabilities map[string][]string `json:"account_ids_with_updated_liabilities"`
}

// NewLiabilitiesDefaultUpdateWebhook instantiates a new LiabilitiesDefaultUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilitiesDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, error_ NullablePlaidError, accountIdsWithNewLiabilities []string, accountIdsWithUpdatedLiabilities map[string][]string) *LiabilitiesDefaultUpdateWebhook {
	this := LiabilitiesDefaultUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.Error = error_
	this.AccountIdsWithNewLiabilities = accountIdsWithNewLiabilities
	this.AccountIdsWithUpdatedLiabilities = accountIdsWithUpdatedLiabilities
	return &this
}

// NewLiabilitiesDefaultUpdateWebhookWithDefaults instantiates a new LiabilitiesDefaultUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilitiesDefaultUpdateWebhookWithDefaults() *LiabilitiesDefaultUpdateWebhook {
	this := LiabilitiesDefaultUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *LiabilitiesDefaultUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *LiabilitiesDefaultUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *LiabilitiesDefaultUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesDefaultUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *LiabilitiesDefaultUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for PlaidError will be returned
func (o *LiabilitiesDefaultUpdateWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiabilitiesDefaultUpdateWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *LiabilitiesDefaultUpdateWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}

// GetAccountIdsWithNewLiabilities returns the AccountIdsWithNewLiabilities field value
func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithNewLiabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountIdsWithNewLiabilities
}

// GetAccountIdsWithNewLiabilitiesOk returns a tuple with the AccountIdsWithNewLiabilities field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithNewLiabilitiesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIdsWithNewLiabilities, true
}

// SetAccountIdsWithNewLiabilities sets field value
func (o *LiabilitiesDefaultUpdateWebhook) SetAccountIdsWithNewLiabilities(v []string) {
	o.AccountIdsWithNewLiabilities = v
}

// GetAccountIdsWithUpdatedLiabilities returns the AccountIdsWithUpdatedLiabilities field value
func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithUpdatedLiabilities() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.AccountIdsWithUpdatedLiabilities
}

// GetAccountIdsWithUpdatedLiabilitiesOk returns a tuple with the AccountIdsWithUpdatedLiabilities field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithUpdatedLiabilitiesOk() (*map[string][]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIdsWithUpdatedLiabilities, true
}

// SetAccountIdsWithUpdatedLiabilities sets field value
func (o *LiabilitiesDefaultUpdateWebhook) SetAccountIdsWithUpdatedLiabilities(v map[string][]string) {
	o.AccountIdsWithUpdatedLiabilities = v
}

func (o LiabilitiesDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["account_ids_with_new_liabilities"] = o.AccountIdsWithNewLiabilities
	}
	if true {
		toSerialize["account_ids_with_updated_liabilities"] = o.AccountIdsWithUpdatedLiabilities
	}
	return json.Marshal(toSerialize)
}

type NullableLiabilitiesDefaultUpdateWebhook struct {
	value *LiabilitiesDefaultUpdateWebhook
	isSet bool
}

func (v NullableLiabilitiesDefaultUpdateWebhook) Get() *LiabilitiesDefaultUpdateWebhook {
	return v.value
}

func (v *NullableLiabilitiesDefaultUpdateWebhook) Set(val *LiabilitiesDefaultUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilitiesDefaultUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilitiesDefaultUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilitiesDefaultUpdateWebhook(val *LiabilitiesDefaultUpdateWebhook) *NullableLiabilitiesDefaultUpdateWebhook {
	return &NullableLiabilitiesDefaultUpdateWebhook{value: val, isSet: true}
}

func (v NullableLiabilitiesDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilitiesDefaultUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


