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

// ItemPublicTokenExchangeResponse ItemPublicTokenExchangeResponse defines the response schema for `/item/public_token/exchange`
type ItemPublicTokenExchangeResponse struct {
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The `item_id` value of the Item associated with the returned `access_token`
	ItemId string `json:"item_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ItemPublicTokenExchangeResponse ItemPublicTokenExchangeResponse

// NewItemPublicTokenExchangeResponse instantiates a new ItemPublicTokenExchangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPublicTokenExchangeResponse(accessToken string, itemId string, requestId string) *ItemPublicTokenExchangeResponse {
	this := ItemPublicTokenExchangeResponse{}
	this.AccessToken = accessToken
	this.ItemId = itemId
	this.RequestId = requestId
	return &this
}

// NewItemPublicTokenExchangeResponseWithDefaults instantiates a new ItemPublicTokenExchangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPublicTokenExchangeResponseWithDefaults() *ItemPublicTokenExchangeResponse {
	this := ItemPublicTokenExchangeResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *ItemPublicTokenExchangeResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenExchangeResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ItemPublicTokenExchangeResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetItemId returns the ItemId field value
func (o *ItemPublicTokenExchangeResponse) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenExchangeResponse) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *ItemPublicTokenExchangeResponse) SetItemId(v string) {
	o.ItemId = v
}

// GetRequestId returns the RequestId field value
func (o *ItemPublicTokenExchangeResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenExchangeResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemPublicTokenExchangeResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ItemPublicTokenExchangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemPublicTokenExchangeResponse) UnmarshalJSON(bytes []byte) (err error) {
	varItemPublicTokenExchangeResponse := _ItemPublicTokenExchangeResponse{}

	if err = json.Unmarshal(bytes, &varItemPublicTokenExchangeResponse); err == nil {
		*o = ItemPublicTokenExchangeResponse(varItemPublicTokenExchangeResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemPublicTokenExchangeResponse struct {
	value *ItemPublicTokenExchangeResponse
	isSet bool
}

func (v NullableItemPublicTokenExchangeResponse) Get() *ItemPublicTokenExchangeResponse {
	return v.value
}

func (v *NullableItemPublicTokenExchangeResponse) Set(val *ItemPublicTokenExchangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPublicTokenExchangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPublicTokenExchangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPublicTokenExchangeResponse(val *ItemPublicTokenExchangeResponse) *NullableItemPublicTokenExchangeResponse {
	return &NullableItemPublicTokenExchangeResponse{value: val, isSet: true}
}

func (v NullableItemPublicTokenExchangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPublicTokenExchangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


