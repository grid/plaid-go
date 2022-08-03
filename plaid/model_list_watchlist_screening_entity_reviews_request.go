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

// ListWatchlistScreeningEntityReviewsRequest Request input for listing reviews for an entity watchlist screening
type ListWatchlistScreeningEntityReviewsRequest struct {
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// ID of the associated entity screening.
	EntityWatchlistScreeningId string `json:"entity_watchlist_screening_id"`
	// An identifier that determines which page of results you receive.
	Cursor NullableString `json:"cursor,omitempty"`
}

// NewListWatchlistScreeningEntityReviewsRequest instantiates a new ListWatchlistScreeningEntityReviewsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWatchlistScreeningEntityReviewsRequest(entityWatchlistScreeningId string) *ListWatchlistScreeningEntityReviewsRequest {
	this := ListWatchlistScreeningEntityReviewsRequest{}
	this.EntityWatchlistScreeningId = entityWatchlistScreeningId
	return &this
}

// NewListWatchlistScreeningEntityReviewsRequestWithDefaults instantiates a new ListWatchlistScreeningEntityReviewsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWatchlistScreeningEntityReviewsRequestWithDefaults() *ListWatchlistScreeningEntityReviewsRequest {
	this := ListWatchlistScreeningEntityReviewsRequest{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ListWatchlistScreeningEntityReviewsRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWatchlistScreeningEntityReviewsRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ListWatchlistScreeningEntityReviewsRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ListWatchlistScreeningEntityReviewsRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ListWatchlistScreeningEntityReviewsRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWatchlistScreeningEntityReviewsRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ListWatchlistScreeningEntityReviewsRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ListWatchlistScreeningEntityReviewsRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetEntityWatchlistScreeningId returns the EntityWatchlistScreeningId field value
func (o *ListWatchlistScreeningEntityReviewsRequest) GetEntityWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistScreeningId
}

// GetEntityWatchlistScreeningIdOk returns a tuple with the EntityWatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *ListWatchlistScreeningEntityReviewsRequest) GetEntityWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningId, true
}

// SetEntityWatchlistScreeningId sets field value
func (o *ListWatchlistScreeningEntityReviewsRequest) SetEntityWatchlistScreeningId(v string) {
	o.EntityWatchlistScreeningId = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListWatchlistScreeningEntityReviewsRequest) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListWatchlistScreeningEntityReviewsRequest) GetCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *ListWatchlistScreeningEntityReviewsRequest) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *ListWatchlistScreeningEntityReviewsRequest) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *ListWatchlistScreeningEntityReviewsRequest) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *ListWatchlistScreeningEntityReviewsRequest) UnsetCursor() {
	o.Cursor.Unset()
}

func (o ListWatchlistScreeningEntityReviewsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["entity_watchlist_screening_id"] = o.EntityWatchlistScreeningId
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListWatchlistScreeningEntityReviewsRequest struct {
	value *ListWatchlistScreeningEntityReviewsRequest
	isSet bool
}

func (v NullableListWatchlistScreeningEntityReviewsRequest) Get() *ListWatchlistScreeningEntityReviewsRequest {
	return v.value
}

func (v *NullableListWatchlistScreeningEntityReviewsRequest) Set(val *ListWatchlistScreeningEntityReviewsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListWatchlistScreeningEntityReviewsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListWatchlistScreeningEntityReviewsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWatchlistScreeningEntityReviewsRequest(val *ListWatchlistScreeningEntityReviewsRequest) *NullableListWatchlistScreeningEntityReviewsRequest {
	return &NullableListWatchlistScreeningEntityReviewsRequest{value: val, isSet: true}
}

func (v NullableListWatchlistScreeningEntityReviewsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWatchlistScreeningEntityReviewsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


