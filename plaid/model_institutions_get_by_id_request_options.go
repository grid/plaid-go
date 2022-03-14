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

// InstitutionsGetByIdRequestOptions Specifies optional parameters for `/institutions/get_by_id`. If provided, must not be `null`.
type InstitutionsGetByIdRequestOptions struct {
	// When `true`, return an institution's logo, brand color, and URL. When available, the bank's logo is returned as a base64 encoded 152x152 PNG, the brand color is in hexadecimal format. The default value is `false`.  Note that Plaid does not own any of the logos shared by the API and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos.
	IncludeOptionalMetadata *bool `json:"include_optional_metadata,omitempty"`
	// If `true`, the response will include status information about the institution. Default value is `false`.
	IncludeStatus *bool `json:"include_status,omitempty"`
	// When `true`, returns metadata related to the Auth product indicating which auth methods are supported.
	IncludeAuthMetadata *bool `json:"include_auth_metadata,omitempty"`
	// When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.
	IncludePaymentInitiationMetadata *bool `json:"include_payment_initiation_metadata,omitempty"`
}

// NewInstitutionsGetByIdRequestOptions instantiates a new InstitutionsGetByIdRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsGetByIdRequestOptions() *InstitutionsGetByIdRequestOptions {
	this := InstitutionsGetByIdRequestOptions{}
	var includeOptionalMetadata bool = false
	this.IncludeOptionalMetadata = &includeOptionalMetadata
	var includeStatus bool = false
	this.IncludeStatus = &includeStatus
	var includeAuthMetadata bool = false
	this.IncludeAuthMetadata = &includeAuthMetadata
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = &includePaymentInitiationMetadata
	return &this
}

// NewInstitutionsGetByIdRequestOptionsWithDefaults instantiates a new InstitutionsGetByIdRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsGetByIdRequestOptionsWithDefaults() *InstitutionsGetByIdRequestOptions {
	this := InstitutionsGetByIdRequestOptions{}
	var includeOptionalMetadata bool = false
	this.IncludeOptionalMetadata = &includeOptionalMetadata
	var includeStatus bool = false
	this.IncludeStatus = &includeStatus
	var includeAuthMetadata bool = false
	this.IncludeAuthMetadata = &includeAuthMetadata
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = &includePaymentInitiationMetadata
	return &this
}

// GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequestOptions) GetIncludeOptionalMetadata() bool {
	if o == nil || o.IncludeOptionalMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptionalMetadata
}

// GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeOptionalMetadata == nil {
		return nil, false
	}
	return o.IncludeOptionalMetadata, true
}

// HasIncludeOptionalMetadata returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequestOptions) HasIncludeOptionalMetadata() bool {
	if o != nil && o.IncludeOptionalMetadata != nil {
		return true
	}

	return false
}

// SetIncludeOptionalMetadata gets a reference to the given bool and assigns it to the IncludeOptionalMetadata field.
func (o *InstitutionsGetByIdRequestOptions) SetIncludeOptionalMetadata(v bool) {
	o.IncludeOptionalMetadata = &v
}

// GetIncludeStatus returns the IncludeStatus field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequestOptions) GetIncludeStatus() bool {
	if o == nil || o.IncludeStatus == nil {
		var ret bool
		return ret
	}
	return *o.IncludeStatus
}

// GetIncludeStatusOk returns a tuple with the IncludeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequestOptions) GetIncludeStatusOk() (*bool, bool) {
	if o == nil || o.IncludeStatus == nil {
		return nil, false
	}
	return o.IncludeStatus, true
}

// HasIncludeStatus returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequestOptions) HasIncludeStatus() bool {
	if o != nil && o.IncludeStatus != nil {
		return true
	}

	return false
}

// SetIncludeStatus gets a reference to the given bool and assigns it to the IncludeStatus field.
func (o *InstitutionsGetByIdRequestOptions) SetIncludeStatus(v bool) {
	o.IncludeStatus = &v
}

// GetIncludeAuthMetadata returns the IncludeAuthMetadata field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequestOptions) GetIncludeAuthMetadata() bool {
	if o == nil || o.IncludeAuthMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAuthMetadata
}

// GetIncludeAuthMetadataOk returns a tuple with the IncludeAuthMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequestOptions) GetIncludeAuthMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeAuthMetadata == nil {
		return nil, false
	}
	return o.IncludeAuthMetadata, true
}

// HasIncludeAuthMetadata returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequestOptions) HasIncludeAuthMetadata() bool {
	if o != nil && o.IncludeAuthMetadata != nil {
		return true
	}

	return false
}

// SetIncludeAuthMetadata gets a reference to the given bool and assigns it to the IncludeAuthMetadata field.
func (o *InstitutionsGetByIdRequestOptions) SetIncludeAuthMetadata(v bool) {
	o.IncludeAuthMetadata = &v
}

// GetIncludePaymentInitiationMetadata returns the IncludePaymentInitiationMetadata field value if set, zero value otherwise.
func (o *InstitutionsGetByIdRequestOptions) GetIncludePaymentInitiationMetadata() bool {
	if o == nil || o.IncludePaymentInitiationMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludePaymentInitiationMetadata
}

// GetIncludePaymentInitiationMetadataOk returns a tuple with the IncludePaymentInitiationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetByIdRequestOptions) GetIncludePaymentInitiationMetadataOk() (*bool, bool) {
	if o == nil || o.IncludePaymentInitiationMetadata == nil {
		return nil, false
	}
	return o.IncludePaymentInitiationMetadata, true
}

// HasIncludePaymentInitiationMetadata returns a boolean if a field has been set.
func (o *InstitutionsGetByIdRequestOptions) HasIncludePaymentInitiationMetadata() bool {
	if o != nil && o.IncludePaymentInitiationMetadata != nil {
		return true
	}

	return false
}

// SetIncludePaymentInitiationMetadata gets a reference to the given bool and assigns it to the IncludePaymentInitiationMetadata field.
func (o *InstitutionsGetByIdRequestOptions) SetIncludePaymentInitiationMetadata(v bool) {
	o.IncludePaymentInitiationMetadata = &v
}

func (o InstitutionsGetByIdRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeOptionalMetadata != nil {
		toSerialize["include_optional_metadata"] = o.IncludeOptionalMetadata
	}
	if o.IncludeStatus != nil {
		toSerialize["include_status"] = o.IncludeStatus
	}
	if o.IncludeAuthMetadata != nil {
		toSerialize["include_auth_metadata"] = o.IncludeAuthMetadata
	}
	if o.IncludePaymentInitiationMetadata != nil {
		toSerialize["include_payment_initiation_metadata"] = o.IncludePaymentInitiationMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsGetByIdRequestOptions struct {
	value *InstitutionsGetByIdRequestOptions
	isSet bool
}

func (v NullableInstitutionsGetByIdRequestOptions) Get() *InstitutionsGetByIdRequestOptions {
	return v.value
}

func (v *NullableInstitutionsGetByIdRequestOptions) Set(val *InstitutionsGetByIdRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsGetByIdRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsGetByIdRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsGetByIdRequestOptions(val *InstitutionsGetByIdRequestOptions) *NullableInstitutionsGetByIdRequestOptions {
	return &NullableInstitutionsGetByIdRequestOptions{value: val, isSet: true}
}

func (v NullableInstitutionsGetByIdRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsGetByIdRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


