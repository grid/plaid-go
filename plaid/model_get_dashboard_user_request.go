/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// GetDashboardUserRequest Request input for fetching a dashboard user
type GetDashboardUserRequest struct {
	// ID of the associated user.
	DashboardUserId string `json:"dashboard_user_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
}

// NewGetDashboardUserRequest instantiates a new GetDashboardUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDashboardUserRequest(dashboardUserId string) *GetDashboardUserRequest {
	this := GetDashboardUserRequest{}
	this.DashboardUserId = dashboardUserId
	return &this
}

// NewGetDashboardUserRequestWithDefaults instantiates a new GetDashboardUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDashboardUserRequestWithDefaults() *GetDashboardUserRequest {
	this := GetDashboardUserRequest{}
	return &this
}

// GetDashboardUserId returns the DashboardUserId field value
func (o *GetDashboardUserRequest) GetDashboardUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardUserId
}

// GetDashboardUserIdOk returns a tuple with the DashboardUserId field value
// and a boolean to check if the value has been set.
func (o *GetDashboardUserRequest) GetDashboardUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DashboardUserId, true
}

// SetDashboardUserId sets field value
func (o *GetDashboardUserRequest) SetDashboardUserId(v string) {
	o.DashboardUserId = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *GetDashboardUserRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDashboardUserRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *GetDashboardUserRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *GetDashboardUserRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetDashboardUserRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDashboardUserRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetDashboardUserRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetDashboardUserRequest) SetClientId(v string) {
	o.ClientId = &v
}

func (o GetDashboardUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dashboard_user_id"] = o.DashboardUserId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	return json.Marshal(toSerialize)
}

type NullableGetDashboardUserRequest struct {
	value *GetDashboardUserRequest
	isSet bool
}

func (v NullableGetDashboardUserRequest) Get() *GetDashboardUserRequest {
	return v.value
}

func (v *NullableGetDashboardUserRequest) Set(val *GetDashboardUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDashboardUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDashboardUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDashboardUserRequest(val *GetDashboardUserRequest) *NullableGetDashboardUserRequest {
	return &NullableGetDashboardUserRequest{value: val, isSet: true}
}

func (v NullableGetDashboardUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDashboardUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

