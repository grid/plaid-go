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

// ProductStatusBreakdown A detailed breakdown of the institution's performance for a request type. The values for `success`, `error_plaid`, and `error_institution` sum to 1. The time range used for calculating the breakdown may range from the most recent few minutes to the past six hours. In general, smaller institutions will show status that was calculated over a longer period of time. For Investment updates, which are refreshed less frequently, the period assessed may be 24 hours or more. For more details, see [Institution status details](https://plaid.com/docs/account/activity/#institution-status-details).
type ProductStatusBreakdown struct {
	// The percentage of login attempts that are successful, expressed as a decimal.
	Success float64 `json:"success"`
	// The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal. 
	ErrorPlaid float64 `json:"error_plaid"`
	// The percentage of logins that are failing due to an issue in the institution's system, expressed as a decimal.
	ErrorInstitution float64 `json:"error_institution"`
	// The `refresh_interval` may be `DELAYED` or `STOPPED` even when the success rate is high. This value is only returned for Transactions status breakdowns.
	RefreshInterval *string `json:"refresh_interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductStatusBreakdown ProductStatusBreakdown

// NewProductStatusBreakdown instantiates a new ProductStatusBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductStatusBreakdown(success float64, errorPlaid float64, errorInstitution float64) *ProductStatusBreakdown {
	this := ProductStatusBreakdown{}
	this.Success = success
	this.ErrorPlaid = errorPlaid
	this.ErrorInstitution = errorInstitution
	return &this
}

// NewProductStatusBreakdownWithDefaults instantiates a new ProductStatusBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductStatusBreakdownWithDefaults() *ProductStatusBreakdown {
	this := ProductStatusBreakdown{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ProductStatusBreakdown) GetSuccess() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ProductStatusBreakdown) GetSuccessOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ProductStatusBreakdown) SetSuccess(v float64) {
	o.Success = v
}

// GetErrorPlaid returns the ErrorPlaid field value
func (o *ProductStatusBreakdown) GetErrorPlaid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ErrorPlaid
}

// GetErrorPlaidOk returns a tuple with the ErrorPlaid field value
// and a boolean to check if the value has been set.
func (o *ProductStatusBreakdown) GetErrorPlaidOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorPlaid, true
}

// SetErrorPlaid sets field value
func (o *ProductStatusBreakdown) SetErrorPlaid(v float64) {
	o.ErrorPlaid = v
}

// GetErrorInstitution returns the ErrorInstitution field value
func (o *ProductStatusBreakdown) GetErrorInstitution() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ErrorInstitution
}

// GetErrorInstitutionOk returns a tuple with the ErrorInstitution field value
// and a boolean to check if the value has been set.
func (o *ProductStatusBreakdown) GetErrorInstitutionOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorInstitution, true
}

// SetErrorInstitution sets field value
func (o *ProductStatusBreakdown) SetErrorInstitution(v float64) {
	o.ErrorInstitution = v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *ProductStatusBreakdown) GetRefreshInterval() string {
	if o == nil || o.RefreshInterval == nil {
		var ret string
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductStatusBreakdown) GetRefreshIntervalOk() (*string, bool) {
	if o == nil || o.RefreshInterval == nil {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *ProductStatusBreakdown) HasRefreshInterval() bool {
	if o != nil && o.RefreshInterval != nil {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given string and assigns it to the RefreshInterval field.
func (o *ProductStatusBreakdown) SetRefreshInterval(v string) {
	o.RefreshInterval = &v
}

func (o ProductStatusBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["success"] = o.Success
	}
	if true {
		toSerialize["error_plaid"] = o.ErrorPlaid
	}
	if true {
		toSerialize["error_institution"] = o.ErrorInstitution
	}
	if o.RefreshInterval != nil {
		toSerialize["refresh_interval"] = o.RefreshInterval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProductStatusBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	varProductStatusBreakdown := _ProductStatusBreakdown{}

	if err = json.Unmarshal(bytes, &varProductStatusBreakdown); err == nil {
		*o = ProductStatusBreakdown(varProductStatusBreakdown)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "error_plaid")
		delete(additionalProperties, "error_institution")
		delete(additionalProperties, "refresh_interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductStatusBreakdown struct {
	value *ProductStatusBreakdown
	isSet bool
}

func (v NullableProductStatusBreakdown) Get() *ProductStatusBreakdown {
	return v.value
}

func (v *NullableProductStatusBreakdown) Set(val *ProductStatusBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableProductStatusBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableProductStatusBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductStatusBreakdown(val *ProductStatusBreakdown) *NullableProductStatusBreakdown {
	return &NullableProductStatusBreakdown{value: val, isSet: true}
}

func (v NullableProductStatusBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductStatusBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


