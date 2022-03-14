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

// DepositSwitchGetResponse DepositSwitchGetResponse defines the response schema for `/deposit_switch/get`
type DepositSwitchGetResponse struct {
	// The ID of the deposit switch.
	DepositSwitchId string `json:"deposit_switch_id"`
	// The ID of the bank account the direct deposit was switched to.
	TargetAccountId NullableString `json:"target_account_id"`
	// The ID of the Item the direct deposit was switched to.
	TargetItemId NullableString `json:"target_item_id"`
	//  The state, or status, of the deposit switch.  - `initialized` – The deposit switch has been initialized with the user entering the information required to submit the deposit switch request.  - `processing` – The deposit switch request has been submitted and is being processed.  - `completed` – The user's employer has fulfilled the deposit switch request.  - `error` – There was an error processing the deposit switch request.
	State string `json:"state"`
	// The method used to make the deposit switch.  - `instant` – User instantly switched their direct deposit to a new or existing bank account by connecting their payroll or employer account.  - `mail` – User requested that Plaid contact their employer by mail to make the direct deposit switch.  - `pdf` – User generated a PDF or email to be sent to their employer with the information necessary to make the deposit switch.'
	SwitchMethod NullableString `json:"switch_method,omitempty"`
	// When `true`, user’s direct deposit goes to multiple banks. When false, user’s direct deposit only goes to the target account. Always `null` if the deposit switch has not been completed.
	AccountHasMultipleAllocations NullableBool `json:"account_has_multiple_allocations"`
	// When `true`, the target account is allocated the remainder of direct deposit after all other allocations have been deducted. When `false`, user’s direct deposit is allocated as a percent or amount. Always `null` if the deposit switch has not been completed.
	IsAllocatedRemainder NullableBool `json:"is_allocated_remainder"`
	// The percentage of direct deposit allocated to the target account. Always `null` if the target account is not allocated a percentage or if the deposit switch has not been completed or if `is_allocated_remainder` is true.
	PercentAllocated NullableFloat64 `json:"percent_allocated"`
	// The dollar amount of direct deposit allocated to the target account. Always `null` if the target account is not allocated an amount or if the deposit switch has not been completed.
	AmountAllocated NullableFloat64 `json:"amount_allocated"`
	// The name of the employer selected by the user. If the user did not select an employer, the value returned is `null`.
	EmployerName NullableString `json:"employer_name,omitempty"`
	// The ID of the employer selected by the user. If the user did not select an employer, the value returned is `null`.
	EmployerId NullableString `json:"employer_id,omitempty"`
	// The name of the institution selected by the user. If the user did not select an institution, the value returned is `null`.
	InstitutionName NullableString `json:"institution_name,omitempty"`
	// The ID of the institution selected by the user. If the user did not select an institution, the value returned is `null`.
	InstitutionId NullableString `json:"institution_id,omitempty"`
	// [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date the deposit switch was created. 
	DateCreated string `json:"date_created"`
	// [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date the deposit switch was completed. Always `null` if the deposit switch has not been completed. 
	DateCompleted NullableString `json:"date_completed"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchGetResponse DepositSwitchGetResponse

// NewDepositSwitchGetResponse instantiates a new DepositSwitchGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchGetResponse(depositSwitchId string, targetAccountId NullableString, targetItemId NullableString, state string, accountHasMultipleAllocations NullableBool, isAllocatedRemainder NullableBool, percentAllocated NullableFloat64, amountAllocated NullableFloat64, dateCreated string, dateCompleted NullableString, requestId string) *DepositSwitchGetResponse {
	this := DepositSwitchGetResponse{}
	this.DepositSwitchId = depositSwitchId
	this.TargetAccountId = targetAccountId
	this.TargetItemId = targetItemId
	this.State = state
	this.AccountHasMultipleAllocations = accountHasMultipleAllocations
	this.IsAllocatedRemainder = isAllocatedRemainder
	this.PercentAllocated = percentAllocated
	this.AmountAllocated = amountAllocated
	this.DateCreated = dateCreated
	this.DateCompleted = dateCompleted
	this.RequestId = requestId
	return &this
}

// NewDepositSwitchGetResponseWithDefaults instantiates a new DepositSwitchGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchGetResponseWithDefaults() *DepositSwitchGetResponse {
	this := DepositSwitchGetResponse{}
	return &this
}

// GetDepositSwitchId returns the DepositSwitchId field value
func (o *DepositSwitchGetResponse) GetDepositSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchId, true
}

// SetDepositSwitchId sets field value
func (o *DepositSwitchGetResponse) SetDepositSwitchId(v string) {
	o.DepositSwitchId = v
}

// GetTargetAccountId returns the TargetAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DepositSwitchGetResponse) GetTargetAccountId() string {
	if o == nil || o.TargetAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetAccountId.Get()
}

// GetTargetAccountIdOk returns a tuple with the TargetAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetTargetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetAccountId.Get(), o.TargetAccountId.IsSet()
}

// SetTargetAccountId sets field value
func (o *DepositSwitchGetResponse) SetTargetAccountId(v string) {
	o.TargetAccountId.Set(&v)
}

// GetTargetItemId returns the TargetItemId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DepositSwitchGetResponse) GetTargetItemId() string {
	if o == nil || o.TargetItemId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetItemId.Get()
}

// GetTargetItemIdOk returns a tuple with the TargetItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetTargetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetItemId.Get(), o.TargetItemId.IsSet()
}

// SetTargetItemId sets field value
func (o *DepositSwitchGetResponse) SetTargetItemId(v string) {
	o.TargetItemId.Set(&v)
}

// GetState returns the State field value
func (o *DepositSwitchGetResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DepositSwitchGetResponse) SetState(v string) {
	o.State = v
}

// GetSwitchMethod returns the SwitchMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchGetResponse) GetSwitchMethod() string {
	if o == nil || o.SwitchMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.SwitchMethod.Get()
}

// GetSwitchMethodOk returns a tuple with the SwitchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetSwitchMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SwitchMethod.Get(), o.SwitchMethod.IsSet()
}

// HasSwitchMethod returns a boolean if a field has been set.
func (o *DepositSwitchGetResponse) HasSwitchMethod() bool {
	if o != nil && o.SwitchMethod.IsSet() {
		return true
	}

	return false
}

// SetSwitchMethod gets a reference to the given NullableString and assigns it to the SwitchMethod field.
func (o *DepositSwitchGetResponse) SetSwitchMethod(v string) {
	o.SwitchMethod.Set(&v)
}
// SetSwitchMethodNil sets the value for SwitchMethod to be an explicit nil
func (o *DepositSwitchGetResponse) SetSwitchMethodNil() {
	o.SwitchMethod.Set(nil)
}

// UnsetSwitchMethod ensures that no value is present for SwitchMethod, not even an explicit nil
func (o *DepositSwitchGetResponse) UnsetSwitchMethod() {
	o.SwitchMethod.Unset()
}

// GetAccountHasMultipleAllocations returns the AccountHasMultipleAllocations field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DepositSwitchGetResponse) GetAccountHasMultipleAllocations() bool {
	if o == nil || o.AccountHasMultipleAllocations.Get() == nil {
		var ret bool
		return ret
	}

	return *o.AccountHasMultipleAllocations.Get()
}

// GetAccountHasMultipleAllocationsOk returns a tuple with the AccountHasMultipleAllocations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetAccountHasMultipleAllocationsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountHasMultipleAllocations.Get(), o.AccountHasMultipleAllocations.IsSet()
}

// SetAccountHasMultipleAllocations sets field value
func (o *DepositSwitchGetResponse) SetAccountHasMultipleAllocations(v bool) {
	o.AccountHasMultipleAllocations.Set(&v)
}

// GetIsAllocatedRemainder returns the IsAllocatedRemainder field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DepositSwitchGetResponse) GetIsAllocatedRemainder() bool {
	if o == nil || o.IsAllocatedRemainder.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsAllocatedRemainder.Get()
}

// GetIsAllocatedRemainderOk returns a tuple with the IsAllocatedRemainder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetIsAllocatedRemainderOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsAllocatedRemainder.Get(), o.IsAllocatedRemainder.IsSet()
}

// SetIsAllocatedRemainder sets field value
func (o *DepositSwitchGetResponse) SetIsAllocatedRemainder(v bool) {
	o.IsAllocatedRemainder.Set(&v)
}

// GetPercentAllocated returns the PercentAllocated field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *DepositSwitchGetResponse) GetPercentAllocated() float64 {
	if o == nil || o.PercentAllocated.Get() == nil {
		var ret float64
		return ret
	}

	return *o.PercentAllocated.Get()
}

// GetPercentAllocatedOk returns a tuple with the PercentAllocated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetPercentAllocatedOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PercentAllocated.Get(), o.PercentAllocated.IsSet()
}

// SetPercentAllocated sets field value
func (o *DepositSwitchGetResponse) SetPercentAllocated(v float64) {
	o.PercentAllocated.Set(&v)
}

// GetAmountAllocated returns the AmountAllocated field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *DepositSwitchGetResponse) GetAmountAllocated() float64 {
	if o == nil || o.AmountAllocated.Get() == nil {
		var ret float64
		return ret
	}

	return *o.AmountAllocated.Get()
}

// GetAmountAllocatedOk returns a tuple with the AmountAllocated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetAmountAllocatedOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AmountAllocated.Get(), o.AmountAllocated.IsSet()
}

// SetAmountAllocated sets field value
func (o *DepositSwitchGetResponse) SetAmountAllocated(v float64) {
	o.AmountAllocated.Set(&v)
}

// GetEmployerName returns the EmployerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchGetResponse) GetEmployerName() string {
	if o == nil || o.EmployerName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployerName.Get()
}

// GetEmployerNameOk returns a tuple with the EmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetEmployerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployerName.Get(), o.EmployerName.IsSet()
}

// HasEmployerName returns a boolean if a field has been set.
func (o *DepositSwitchGetResponse) HasEmployerName() bool {
	if o != nil && o.EmployerName.IsSet() {
		return true
	}

	return false
}

// SetEmployerName gets a reference to the given NullableString and assigns it to the EmployerName field.
func (o *DepositSwitchGetResponse) SetEmployerName(v string) {
	o.EmployerName.Set(&v)
}
// SetEmployerNameNil sets the value for EmployerName to be an explicit nil
func (o *DepositSwitchGetResponse) SetEmployerNameNil() {
	o.EmployerName.Set(nil)
}

// UnsetEmployerName ensures that no value is present for EmployerName, not even an explicit nil
func (o *DepositSwitchGetResponse) UnsetEmployerName() {
	o.EmployerName.Unset()
}

// GetEmployerId returns the EmployerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchGetResponse) GetEmployerId() string {
	if o == nil || o.EmployerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployerId.Get()
}

// GetEmployerIdOk returns a tuple with the EmployerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetEmployerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployerId.Get(), o.EmployerId.IsSet()
}

// HasEmployerId returns a boolean if a field has been set.
func (o *DepositSwitchGetResponse) HasEmployerId() bool {
	if o != nil && o.EmployerId.IsSet() {
		return true
	}

	return false
}

// SetEmployerId gets a reference to the given NullableString and assigns it to the EmployerId field.
func (o *DepositSwitchGetResponse) SetEmployerId(v string) {
	o.EmployerId.Set(&v)
}
// SetEmployerIdNil sets the value for EmployerId to be an explicit nil
func (o *DepositSwitchGetResponse) SetEmployerIdNil() {
	o.EmployerId.Set(nil)
}

// UnsetEmployerId ensures that no value is present for EmployerId, not even an explicit nil
func (o *DepositSwitchGetResponse) UnsetEmployerId() {
	o.EmployerId.Unset()
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchGetResponse) GetInstitutionName() string {
	if o == nil || o.InstitutionName.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstitutionName.Get()
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetInstitutionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionName.Get(), o.InstitutionName.IsSet()
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *DepositSwitchGetResponse) HasInstitutionName() bool {
	if o != nil && o.InstitutionName.IsSet() {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given NullableString and assigns it to the InstitutionName field.
func (o *DepositSwitchGetResponse) SetInstitutionName(v string) {
	o.InstitutionName.Set(&v)
}
// SetInstitutionNameNil sets the value for InstitutionName to be an explicit nil
func (o *DepositSwitchGetResponse) SetInstitutionNameNil() {
	o.InstitutionName.Set(nil)
}

// UnsetInstitutionName ensures that no value is present for InstitutionName, not even an explicit nil
func (o *DepositSwitchGetResponse) UnsetInstitutionName() {
	o.InstitutionName.Unset()
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchGetResponse) GetInstitutionId() string {
	if o == nil || o.InstitutionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId.Get()
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionId.Get(), o.InstitutionId.IsSet()
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *DepositSwitchGetResponse) HasInstitutionId() bool {
	if o != nil && o.InstitutionId.IsSet() {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given NullableString and assigns it to the InstitutionId field.
func (o *DepositSwitchGetResponse) SetInstitutionId(v string) {
	o.InstitutionId.Set(&v)
}
// SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil
func (o *DepositSwitchGetResponse) SetInstitutionIdNil() {
	o.InstitutionId.Set(nil)
}

// UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
func (o *DepositSwitchGetResponse) UnsetInstitutionId() {
	o.InstitutionId.Unset()
}

// GetDateCreated returns the DateCreated field value
func (o *DepositSwitchGetResponse) GetDateCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetDateCreatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *DepositSwitchGetResponse) SetDateCreated(v string) {
	o.DateCreated = v
}

// GetDateCompleted returns the DateCompleted field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DepositSwitchGetResponse) GetDateCompleted() string {
	if o == nil || o.DateCompleted.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateCompleted.Get()
}

// GetDateCompletedOk returns a tuple with the DateCompleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetDateCompletedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateCompleted.Get(), o.DateCompleted.IsSet()
}

// SetDateCompleted sets field value
func (o *DepositSwitchGetResponse) SetDateCompleted(v string) {
	o.DateCompleted.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *DepositSwitchGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DepositSwitchGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o DepositSwitchGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}
	if true {
		toSerialize["target_account_id"] = o.TargetAccountId.Get()
	}
	if true {
		toSerialize["target_item_id"] = o.TargetItemId.Get()
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.SwitchMethod.IsSet() {
		toSerialize["switch_method"] = o.SwitchMethod.Get()
	}
	if true {
		toSerialize["account_has_multiple_allocations"] = o.AccountHasMultipleAllocations.Get()
	}
	if true {
		toSerialize["is_allocated_remainder"] = o.IsAllocatedRemainder.Get()
	}
	if true {
		toSerialize["percent_allocated"] = o.PercentAllocated.Get()
	}
	if true {
		toSerialize["amount_allocated"] = o.AmountAllocated.Get()
	}
	if o.EmployerName.IsSet() {
		toSerialize["employer_name"] = o.EmployerName.Get()
	}
	if o.EmployerId.IsSet() {
		toSerialize["employer_id"] = o.EmployerId.Get()
	}
	if o.InstitutionName.IsSet() {
		toSerialize["institution_name"] = o.InstitutionName.Get()
	}
	if o.InstitutionId.IsSet() {
		toSerialize["institution_id"] = o.InstitutionId.Get()
	}
	if true {
		toSerialize["date_created"] = o.DateCreated
	}
	if true {
		toSerialize["date_completed"] = o.DateCompleted.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchGetResponse := _DepositSwitchGetResponse{}

	if err = json.Unmarshal(bytes, &varDepositSwitchGetResponse); err == nil {
		*o = DepositSwitchGetResponse(varDepositSwitchGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deposit_switch_id")
		delete(additionalProperties, "target_account_id")
		delete(additionalProperties, "target_item_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "switch_method")
		delete(additionalProperties, "account_has_multiple_allocations")
		delete(additionalProperties, "is_allocated_remainder")
		delete(additionalProperties, "percent_allocated")
		delete(additionalProperties, "amount_allocated")
		delete(additionalProperties, "employer_name")
		delete(additionalProperties, "employer_id")
		delete(additionalProperties, "institution_name")
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "date_created")
		delete(additionalProperties, "date_completed")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchGetResponse struct {
	value *DepositSwitchGetResponse
	isSet bool
}

func (v NullableDepositSwitchGetResponse) Get() *DepositSwitchGetResponse {
	return v.value
}

func (v *NullableDepositSwitchGetResponse) Set(val *DepositSwitchGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchGetResponse(val *DepositSwitchGetResponse) *NullableDepositSwitchGetResponse {
	return &NullableDepositSwitchGetResponse{value: val, isSet: true}
}

func (v NullableDepositSwitchGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


