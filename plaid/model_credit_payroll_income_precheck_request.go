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

// CreditPayrollIncomePrecheckRequest Defines the request schema for `/credit/payroll_income/precheck`.
type CreditPayrollIncomePrecheckRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken *string `json:"user_token,omitempty"`
	// An array of access tokens corresponding to Items belonging to the user whose eligibility is being checked. Note that if the Items specified here are not already initialized with `transactions`, providing them in this field will cause these Items to be initialized with (and billed for) the Transactions product.
	AccessTokens *[]string `json:"access_tokens,omitempty"`
	Employer NullableIncomeVerificationPrecheckEmployer `json:"employer,omitempty"`
	UsMilitaryInfo NullableIncomeVerificationPrecheckMilitaryInfo `json:"us_military_info,omitempty"`
}

// NewCreditPayrollIncomePrecheckRequest instantiates a new CreditPayrollIncomePrecheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomePrecheckRequest() *CreditPayrollIncomePrecheckRequest {
	this := CreditPayrollIncomePrecheckRequest{}
	return &this
}

// NewCreditPayrollIncomePrecheckRequestWithDefaults instantiates a new CreditPayrollIncomePrecheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomePrecheckRequestWithDefaults() *CreditPayrollIncomePrecheckRequest {
	this := CreditPayrollIncomePrecheckRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditPayrollIncomePrecheckRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomePrecheckRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditPayrollIncomePrecheckRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditPayrollIncomePrecheckRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditPayrollIncomePrecheckRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomePrecheckRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditPayrollIncomePrecheckRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditPayrollIncomePrecheckRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *CreditPayrollIncomePrecheckRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomePrecheckRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *CreditPayrollIncomePrecheckRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *CreditPayrollIncomePrecheckRequest) SetUserToken(v string) {
	o.UserToken = &v
}

// GetAccessTokens returns the AccessTokens field value if set, zero value otherwise.
func (o *CreditPayrollIncomePrecheckRequest) GetAccessTokens() []string {
	if o == nil || o.AccessTokens == nil {
		var ret []string
		return ret
	}
	return *o.AccessTokens
}

// GetAccessTokensOk returns a tuple with the AccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomePrecheckRequest) GetAccessTokensOk() (*[]string, bool) {
	if o == nil || o.AccessTokens == nil {
		return nil, false
	}
	return o.AccessTokens, true
}

// HasAccessTokens returns a boolean if a field has been set.
func (o *CreditPayrollIncomePrecheckRequest) HasAccessTokens() bool {
	if o != nil && o.AccessTokens != nil {
		return true
	}

	return false
}

// SetAccessTokens gets a reference to the given []string and assigns it to the AccessTokens field.
func (o *CreditPayrollIncomePrecheckRequest) SetAccessTokens(v []string) {
	o.AccessTokens = &v
}

// GetEmployer returns the Employer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditPayrollIncomePrecheckRequest) GetEmployer() IncomeVerificationPrecheckEmployer {
	if o == nil || o.Employer.Get() == nil {
		var ret IncomeVerificationPrecheckEmployer
		return ret
	}
	return *o.Employer.Get()
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayrollIncomePrecheckRequest) GetEmployerOk() (*IncomeVerificationPrecheckEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employer.Get(), o.Employer.IsSet()
}

// HasEmployer returns a boolean if a field has been set.
func (o *CreditPayrollIncomePrecheckRequest) HasEmployer() bool {
	if o != nil && o.Employer.IsSet() {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given NullableIncomeVerificationPrecheckEmployer and assigns it to the Employer field.
func (o *CreditPayrollIncomePrecheckRequest) SetEmployer(v IncomeVerificationPrecheckEmployer) {
	o.Employer.Set(&v)
}
// SetEmployerNil sets the value for Employer to be an explicit nil
func (o *CreditPayrollIncomePrecheckRequest) SetEmployerNil() {
	o.Employer.Set(nil)
}

// UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
func (o *CreditPayrollIncomePrecheckRequest) UnsetEmployer() {
	o.Employer.Unset()
}

// GetUsMilitaryInfo returns the UsMilitaryInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditPayrollIncomePrecheckRequest) GetUsMilitaryInfo() IncomeVerificationPrecheckMilitaryInfo {
	if o == nil || o.UsMilitaryInfo.Get() == nil {
		var ret IncomeVerificationPrecheckMilitaryInfo
		return ret
	}
	return *o.UsMilitaryInfo.Get()
}

// GetUsMilitaryInfoOk returns a tuple with the UsMilitaryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayrollIncomePrecheckRequest) GetUsMilitaryInfoOk() (*IncomeVerificationPrecheckMilitaryInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsMilitaryInfo.Get(), o.UsMilitaryInfo.IsSet()
}

// HasUsMilitaryInfo returns a boolean if a field has been set.
func (o *CreditPayrollIncomePrecheckRequest) HasUsMilitaryInfo() bool {
	if o != nil && o.UsMilitaryInfo.IsSet() {
		return true
	}

	return false
}

// SetUsMilitaryInfo gets a reference to the given NullableIncomeVerificationPrecheckMilitaryInfo and assigns it to the UsMilitaryInfo field.
func (o *CreditPayrollIncomePrecheckRequest) SetUsMilitaryInfo(v IncomeVerificationPrecheckMilitaryInfo) {
	o.UsMilitaryInfo.Set(&v)
}
// SetUsMilitaryInfoNil sets the value for UsMilitaryInfo to be an explicit nil
func (o *CreditPayrollIncomePrecheckRequest) SetUsMilitaryInfoNil() {
	o.UsMilitaryInfo.Set(nil)
}

// UnsetUsMilitaryInfo ensures that no value is present for UsMilitaryInfo, not even an explicit nil
func (o *CreditPayrollIncomePrecheckRequest) UnsetUsMilitaryInfo() {
	o.UsMilitaryInfo.Unset()
}

func (o CreditPayrollIncomePrecheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}
	if o.AccessTokens != nil {
		toSerialize["access_tokens"] = o.AccessTokens
	}
	if o.Employer.IsSet() {
		toSerialize["employer"] = o.Employer.Get()
	}
	if o.UsMilitaryInfo.IsSet() {
		toSerialize["us_military_info"] = o.UsMilitaryInfo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditPayrollIncomePrecheckRequest struct {
	value *CreditPayrollIncomePrecheckRequest
	isSet bool
}

func (v NullableCreditPayrollIncomePrecheckRequest) Get() *CreditPayrollIncomePrecheckRequest {
	return v.value
}

func (v *NullableCreditPayrollIncomePrecheckRequest) Set(val *CreditPayrollIncomePrecheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomePrecheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomePrecheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomePrecheckRequest(val *CreditPayrollIncomePrecheckRequest) *NullableCreditPayrollIncomePrecheckRequest {
	return &NullableCreditPayrollIncomePrecheckRequest{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomePrecheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomePrecheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


