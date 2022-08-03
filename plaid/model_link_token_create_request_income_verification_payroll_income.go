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

// LinkTokenCreateRequestIncomeVerificationPayrollIncome Specifies options for initializing Link for use with Payroll Income. This field is required if `income_verification` is included in the `products` array and `payroll` is specified in `income_source_types`.
type LinkTokenCreateRequestIncomeVerificationPayrollIncome struct {
	// The types of payroll income verification to enable for the link session. If none are specified, then users will see both document and digital payroll income.
	FlowTypes []IncomeVerificationPayrollFlowType `json:"flow_types,omitempty"`
	// An identifier to indicate whether the income verification link token will be used for an update or not
	IsUpdateMode *bool `json:"is_update_mode,omitempty"`
}

// NewLinkTokenCreateRequestIncomeVerificationPayrollIncome instantiates a new LinkTokenCreateRequestIncomeVerificationPayrollIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestIncomeVerificationPayrollIncome() *LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	this := LinkTokenCreateRequestIncomeVerificationPayrollIncome{}
	return &this
}

// NewLinkTokenCreateRequestIncomeVerificationPayrollIncomeWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerificationPayrollIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestIncomeVerificationPayrollIncomeWithDefaults() *LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	this := LinkTokenCreateRequestIncomeVerificationPayrollIncome{}
	return &this
}

// GetFlowTypes returns the FlowTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetFlowTypes() []IncomeVerificationPayrollFlowType {
	if o == nil  {
		var ret []IncomeVerificationPayrollFlowType
		return ret
	}
	return o.FlowTypes
}

// GetFlowTypesOk returns a tuple with the FlowTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetFlowTypesOk() (*[]IncomeVerificationPayrollFlowType, bool) {
	if o == nil || o.FlowTypes == nil {
		return nil, false
	}
	return &o.FlowTypes, true
}

// HasFlowTypes returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) HasFlowTypes() bool {
	if o != nil && o.FlowTypes != nil {
		return true
	}

	return false
}

// SetFlowTypes gets a reference to the given []IncomeVerificationPayrollFlowType and assigns it to the FlowTypes field.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) SetFlowTypes(v []IncomeVerificationPayrollFlowType) {
	o.FlowTypes = v
}

// GetIsUpdateMode returns the IsUpdateMode field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetIsUpdateMode() bool {
	if o == nil || o.IsUpdateMode == nil {
		var ret bool
		return ret
	}
	return *o.IsUpdateMode
}

// GetIsUpdateModeOk returns a tuple with the IsUpdateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetIsUpdateModeOk() (*bool, bool) {
	if o == nil || o.IsUpdateMode == nil {
		return nil, false
	}
	return o.IsUpdateMode, true
}

// HasIsUpdateMode returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) HasIsUpdateMode() bool {
	if o != nil && o.IsUpdateMode != nil {
		return true
	}

	return false
}

// SetIsUpdateMode gets a reference to the given bool and assigns it to the IsUpdateMode field.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) SetIsUpdateMode(v bool) {
	o.IsUpdateMode = &v
}

func (o LinkTokenCreateRequestIncomeVerificationPayrollIncome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowTypes != nil {
		toSerialize["flow_types"] = o.FlowTypes
	}
	if o.IsUpdateMode != nil {
		toSerialize["is_update_mode"] = o.IsUpdateMode
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome struct {
	value *LinkTokenCreateRequestIncomeVerificationPayrollIncome
	isSet bool
}

func (v NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) Get() *LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	return v.value
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) Set(val *LinkTokenCreateRequestIncomeVerificationPayrollIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestIncomeVerificationPayrollIncome(val *LinkTokenCreateRequestIncomeVerificationPayrollIncome) *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome {
	return &NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


