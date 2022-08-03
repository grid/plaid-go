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

// IncomeVerificationPaystubsGetRequest IncomeVerificationPaystubsGetRequest defines the request schema for `/income/verification/paystubs/get`.
type IncomeVerificationPaystubsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the verification for which to get paystub information.
	IncomeVerificationId NullableString `json:"income_verification_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken NullableString `json:"access_token,omitempty"`
}

// NewIncomeVerificationPaystubsGetRequest instantiates a new IncomeVerificationPaystubsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPaystubsGetRequest() *IncomeVerificationPaystubsGetRequest {
	this := IncomeVerificationPaystubsGetRequest{}
	return &this
}

// NewIncomeVerificationPaystubsGetRequestWithDefaults instantiates a new IncomeVerificationPaystubsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPaystubsGetRequestWithDefaults() *IncomeVerificationPaystubsGetRequest {
	this := IncomeVerificationPaystubsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IncomeVerificationPaystubsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IncomeVerificationPaystubsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IncomeVerificationPaystubsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IncomeVerificationPaystubsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIncomeVerificationId returns the IncomeVerificationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPaystubsGetRequest) GetIncomeVerificationId() string {
	if o == nil || o.IncomeVerificationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IncomeVerificationId.Get()
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPaystubsGetRequest) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncomeVerificationId.Get(), o.IncomeVerificationId.IsSet()
}

// HasIncomeVerificationId returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubsGetRequest) HasIncomeVerificationId() bool {
	if o != nil && o.IncomeVerificationId.IsSet() {
		return true
	}

	return false
}

// SetIncomeVerificationId gets a reference to the given NullableString and assigns it to the IncomeVerificationId field.
func (o *IncomeVerificationPaystubsGetRequest) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId.Set(&v)
}
// SetIncomeVerificationIdNil sets the value for IncomeVerificationId to be an explicit nil
func (o *IncomeVerificationPaystubsGetRequest) SetIncomeVerificationIdNil() {
	o.IncomeVerificationId.Set(nil)
}

// UnsetIncomeVerificationId ensures that no value is present for IncomeVerificationId, not even an explicit nil
func (o *IncomeVerificationPaystubsGetRequest) UnsetIncomeVerificationId() {
	o.IncomeVerificationId.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPaystubsGetRequest) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPaystubsGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubsGetRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *IncomeVerificationPaystubsGetRequest) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *IncomeVerificationPaystubsGetRequest) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *IncomeVerificationPaystubsGetRequest) UnsetAccessToken() {
	o.AccessToken.Unset()
}

func (o IncomeVerificationPaystubsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.IncomeVerificationId.IsSet() {
		toSerialize["income_verification_id"] = o.IncomeVerificationId.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPaystubsGetRequest struct {
	value *IncomeVerificationPaystubsGetRequest
	isSet bool
}

func (v NullableIncomeVerificationPaystubsGetRequest) Get() *IncomeVerificationPaystubsGetRequest {
	return v.value
}

func (v *NullableIncomeVerificationPaystubsGetRequest) Set(val *IncomeVerificationPaystubsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPaystubsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPaystubsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPaystubsGetRequest(val *IncomeVerificationPaystubsGetRequest) *NullableIncomeVerificationPaystubsGetRequest {
	return &NullableIncomeVerificationPaystubsGetRequest{value: val, isSet: true}
}

func (v NullableIncomeVerificationPaystubsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPaystubsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


