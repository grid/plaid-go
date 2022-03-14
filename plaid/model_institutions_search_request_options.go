/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InstitutionsSearchRequestOptions An optional object to filter `/institutions/search` results.
type InstitutionsSearchRequestOptions struct {
	// Limit results to institutions with or without OAuth login flows.
	Oauth NullableBool `json:"oauth,omitempty"`
	// When true, return the institution's homepage URL, logo and primary brand color.
	IncludeOptionalMetadata *bool `json:"include_optional_metadata,omitempty"`
	// When `true`, returns metadata related to the Auth product indicating which auth methods are supported.
	IncludeAuthMetadata NullableBool `json:"include_auth_metadata,omitempty"`
	// When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.
	IncludePaymentInitiationMetadata NullableBool `json:"include_payment_initiation_metadata,omitempty"`
	PaymentInitiation NullableInstitutionsSearchPaymentInitiationOptions `json:"payment_initiation,omitempty"`
}

// NewInstitutionsSearchRequestOptions instantiates a new InstitutionsSearchRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsSearchRequestOptions() *InstitutionsSearchRequestOptions {
	this := InstitutionsSearchRequestOptions{}
	var includeAuthMetadata bool = false
	this.IncludeAuthMetadata = *NewNullableBool(&includeAuthMetadata)
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = *NewNullableBool(&includePaymentInitiationMetadata)
	return &this
}

// NewInstitutionsSearchRequestOptionsWithDefaults instantiates a new InstitutionsSearchRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsSearchRequestOptionsWithDefaults() *InstitutionsSearchRequestOptions {
	this := InstitutionsSearchRequestOptions{}
	var includeAuthMetadata bool = false
	this.IncludeAuthMetadata = *NewNullableBool(&includeAuthMetadata)
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = *NewNullableBool(&includePaymentInitiationMetadata)
	return &this
}

// GetOauth returns the Oauth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsSearchRequestOptions) GetOauth() bool {
	if o == nil || o.Oauth.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Oauth.Get()
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchRequestOptions) GetOauthOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Oauth.Get(), o.Oauth.IsSet()
}

// HasOauth returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasOauth() bool {
	if o != nil && o.Oauth.IsSet() {
		return true
	}

	return false
}

// SetOauth gets a reference to the given NullableBool and assigns it to the Oauth field.
func (o *InstitutionsSearchRequestOptions) SetOauth(v bool) {
	o.Oauth.Set(&v)
}
// SetOauthNil sets the value for Oauth to be an explicit nil
func (o *InstitutionsSearchRequestOptions) SetOauthNil() {
	o.Oauth.Set(nil)
}

// UnsetOauth ensures that no value is present for Oauth, not even an explicit nil
func (o *InstitutionsSearchRequestOptions) UnsetOauth() {
	o.Oauth.Unset()
}

// GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadata() bool {
	if o == nil || o.IncludeOptionalMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptionalMetadata
}

// GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeOptionalMetadata == nil {
		return nil, false
	}
	return o.IncludeOptionalMetadata, true
}

// HasIncludeOptionalMetadata returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasIncludeOptionalMetadata() bool {
	if o != nil && o.IncludeOptionalMetadata != nil {
		return true
	}

	return false
}

// SetIncludeOptionalMetadata gets a reference to the given bool and assigns it to the IncludeOptionalMetadata field.
func (o *InstitutionsSearchRequestOptions) SetIncludeOptionalMetadata(v bool) {
	o.IncludeOptionalMetadata = &v
}

// GetIncludeAuthMetadata returns the IncludeAuthMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsSearchRequestOptions) GetIncludeAuthMetadata() bool {
	if o == nil || o.IncludeAuthMetadata.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAuthMetadata.Get()
}

// GetIncludeAuthMetadataOk returns a tuple with the IncludeAuthMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchRequestOptions) GetIncludeAuthMetadataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeAuthMetadata.Get(), o.IncludeAuthMetadata.IsSet()
}

// HasIncludeAuthMetadata returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasIncludeAuthMetadata() bool {
	if o != nil && o.IncludeAuthMetadata.IsSet() {
		return true
	}

	return false
}

// SetIncludeAuthMetadata gets a reference to the given NullableBool and assigns it to the IncludeAuthMetadata field.
func (o *InstitutionsSearchRequestOptions) SetIncludeAuthMetadata(v bool) {
	o.IncludeAuthMetadata.Set(&v)
}
// SetIncludeAuthMetadataNil sets the value for IncludeAuthMetadata to be an explicit nil
func (o *InstitutionsSearchRequestOptions) SetIncludeAuthMetadataNil() {
	o.IncludeAuthMetadata.Set(nil)
}

// UnsetIncludeAuthMetadata ensures that no value is present for IncludeAuthMetadata, not even an explicit nil
func (o *InstitutionsSearchRequestOptions) UnsetIncludeAuthMetadata() {
	o.IncludeAuthMetadata.Unset()
}

// GetIncludePaymentInitiationMetadata returns the IncludePaymentInitiationMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsSearchRequestOptions) GetIncludePaymentInitiationMetadata() bool {
	if o == nil || o.IncludePaymentInitiationMetadata.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludePaymentInitiationMetadata.Get()
}

// GetIncludePaymentInitiationMetadataOk returns a tuple with the IncludePaymentInitiationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchRequestOptions) GetIncludePaymentInitiationMetadataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludePaymentInitiationMetadata.Get(), o.IncludePaymentInitiationMetadata.IsSet()
}

// HasIncludePaymentInitiationMetadata returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasIncludePaymentInitiationMetadata() bool {
	if o != nil && o.IncludePaymentInitiationMetadata.IsSet() {
		return true
	}

	return false
}

// SetIncludePaymentInitiationMetadata gets a reference to the given NullableBool and assigns it to the IncludePaymentInitiationMetadata field.
func (o *InstitutionsSearchRequestOptions) SetIncludePaymentInitiationMetadata(v bool) {
	o.IncludePaymentInitiationMetadata.Set(&v)
}
// SetIncludePaymentInitiationMetadataNil sets the value for IncludePaymentInitiationMetadata to be an explicit nil
func (o *InstitutionsSearchRequestOptions) SetIncludePaymentInitiationMetadataNil() {
	o.IncludePaymentInitiationMetadata.Set(nil)
}

// UnsetIncludePaymentInitiationMetadata ensures that no value is present for IncludePaymentInitiationMetadata, not even an explicit nil
func (o *InstitutionsSearchRequestOptions) UnsetIncludePaymentInitiationMetadata() {
	o.IncludePaymentInitiationMetadata.Unset()
}

// GetPaymentInitiation returns the PaymentInitiation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsSearchRequestOptions) GetPaymentInitiation() InstitutionsSearchPaymentInitiationOptions {
	if o == nil || o.PaymentInitiation.Get() == nil {
		var ret InstitutionsSearchPaymentInitiationOptions
		return ret
	}
	return *o.PaymentInitiation.Get()
}

// GetPaymentInitiationOk returns a tuple with the PaymentInitiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchRequestOptions) GetPaymentInitiationOk() (*InstitutionsSearchPaymentInitiationOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentInitiation.Get(), o.PaymentInitiation.IsSet()
}

// HasPaymentInitiation returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasPaymentInitiation() bool {
	if o != nil && o.PaymentInitiation.IsSet() {
		return true
	}

	return false
}

// SetPaymentInitiation gets a reference to the given NullableInstitutionsSearchPaymentInitiationOptions and assigns it to the PaymentInitiation field.
func (o *InstitutionsSearchRequestOptions) SetPaymentInitiation(v InstitutionsSearchPaymentInitiationOptions) {
	o.PaymentInitiation.Set(&v)
}
// SetPaymentInitiationNil sets the value for PaymentInitiation to be an explicit nil
func (o *InstitutionsSearchRequestOptions) SetPaymentInitiationNil() {
	o.PaymentInitiation.Set(nil)
}

// UnsetPaymentInitiation ensures that no value is present for PaymentInitiation, not even an explicit nil
func (o *InstitutionsSearchRequestOptions) UnsetPaymentInitiation() {
	o.PaymentInitiation.Unset()
}

func (o InstitutionsSearchRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Oauth.IsSet() {
		toSerialize["oauth"] = o.Oauth.Get()
	}
	if o.IncludeOptionalMetadata != nil {
		toSerialize["include_optional_metadata"] = o.IncludeOptionalMetadata
	}
	if o.IncludeAuthMetadata.IsSet() {
		toSerialize["include_auth_metadata"] = o.IncludeAuthMetadata.Get()
	}
	if o.IncludePaymentInitiationMetadata.IsSet() {
		toSerialize["include_payment_initiation_metadata"] = o.IncludePaymentInitiationMetadata.Get()
	}
	if o.PaymentInitiation.IsSet() {
		toSerialize["payment_initiation"] = o.PaymentInitiation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsSearchRequestOptions struct {
	value *InstitutionsSearchRequestOptions
	isSet bool
}

func (v NullableInstitutionsSearchRequestOptions) Get() *InstitutionsSearchRequestOptions {
	return v.value
}

func (v *NullableInstitutionsSearchRequestOptions) Set(val *InstitutionsSearchRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsSearchRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsSearchRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsSearchRequestOptions(val *InstitutionsSearchRequestOptions) *NullableInstitutionsSearchRequestOptions {
	return &NullableInstitutionsSearchRequestOptions{value: val, isSet: true}
}

func (v NullableInstitutionsSearchRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsSearchRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


